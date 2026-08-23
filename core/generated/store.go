package generated

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

type Limits struct {
	MaxEntries    int
	MaxFileBytes  int64
	MaxTotalBytes int64
}

func DefaultLimits() Limits {
	return Limits{
		MaxEntries:    10_000,
		MaxFileBytes:  16 << 20,
		MaxTotalBytes: 256 << 20,
	}
}

type Archive struct {
	Hash    string                     `json:"hash"`
	Root    model.LocalPath            `json:"root"`
	Entries map[string]model.LocalPath `json:"entries"`
}

type Report struct {
	Archive     *Archive           `json:"archive,omitempty"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

type manifest struct {
	Hash    string   `json:"hash"`
	Entries []string `json:"entries"`
}

func ExtractSrcJar(
	archivePath string,
	cacheRoot string,
	limits Limits,
) Report {
	report := Report{Diagnostics: make([]model.Diagnostic, 0)}
	if limits.MaxEntries <= 0 || limits.MaxFileBytes <= 0 ||
		limits.MaxTotalBytes <= 0 {
		limits = DefaultLimits()
	}
	hash, err := fileHash(archivePath)
	if err != nil {
		return archiveError(
			report,
			"GENERATED_ARCHIVE_OPEN_FAILED",
			err,
			archivePath,
		)
	}
	base := filepath.Join(cacheRoot, "srcjar")
	finalRoot := filepath.Join(base, hash)
	if archive, ok := loadArchive(finalRoot, hash); ok {
		report.Archive = archive
		return report
	}
	if err := os.MkdirAll(base, 0o755); err != nil {
		return archiveError(
			report,
			"GENERATED_ENTRY_WRITE_FAILED",
			err,
			base,
		)
	}
	temporary, err := os.MkdirTemp(base, ".extract-")
	if err != nil {
		return archiveError(
			report,
			"GENERATED_ENTRY_WRITE_FAILED",
			err,
			base,
		)
	}
	published := false
	defer func() {
		if !published {
			_ = os.RemoveAll(temporary)
		}
	}()
	entriesRoot := filepath.Join(temporary, "entries")
	reader, err := zip.OpenReader(archivePath)
	if err != nil {
		return archiveError(
			report,
			"GENERATED_ARCHIVE_OPEN_FAILED",
			err,
			archivePath,
		)
	}
	defer reader.Close()
	entries := make([]string, 0)
	var total int64
	for _, file := range reader.File {
		name, safe := safeEntryName(file.Name)
		if !safe || file.Mode()&os.ModeSymlink != 0 {
			return archiveMessage(
				report,
				"GENERATED_ENTRY_UNSAFE",
				"unsafe srcjar entry: "+file.Name,
				archivePath,
			)
		}
		if file.FileInfo().IsDir() {
			continue
		}
		if !file.Mode().IsRegular() || len(entries) >= limits.MaxEntries ||
			int64(file.UncompressedSize64) > limits.MaxFileBytes ||
			total+int64(file.UncompressedSize64) > limits.MaxTotalBytes {
			return archiveMessage(
				report,
				"GENERATED_ARCHIVE_LIMIT",
				"srcjar entry or size limit exceeded",
				archivePath,
			)
		}
		target := filepath.Join(entriesRoot, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
			return archiveError(
				report,
				"GENERATED_ENTRY_WRITE_FAILED",
				err,
				target,
			)
		}
		written, err := extractEntry(file, target, limits.MaxFileBytes)
		if err != nil {
			return archiveError(
				report,
				"GENERATED_ENTRY_WRITE_FAILED",
				err,
				target,
			)
		}
		if total+written > limits.MaxTotalBytes {
			return archiveMessage(
				report,
				"GENERATED_ARCHIVE_LIMIT",
				"srcjar extracted size limit exceeded",
				archivePath,
			)
		}
		total += written
		entries = append(entries, name)
	}
	sort.Strings(entries)
	data, err := json.MarshalIndent(manifest{Hash: hash, Entries: entries}, "", "  ")
	if err != nil {
		return archiveError(report, "GENERATED_ENTRY_WRITE_FAILED", err, temporary)
	}
	if err := os.WriteFile(
		filepath.Join(temporary, "manifest.json"),
		append(data, '\n'),
		0o644,
	); err != nil {
		return archiveError(report, "GENERATED_ENTRY_WRITE_FAILED", err, temporary)
	}
	if err := os.Rename(temporary, finalRoot); err != nil {
		if archive, ok := loadArchive(finalRoot, hash); ok {
			report.Archive = archive
			return report
		}
		return archiveError(report, "GENERATED_ENTRY_WRITE_FAILED", err, finalRoot)
	}
	published = true
	archive, ok := loadArchive(finalRoot, hash)
	if !ok {
		return archiveMessage(
			report,
			"GENERATED_ENTRY_WRITE_FAILED",
			"published srcjar cache is incomplete",
			finalRoot,
		)
	}
	report.Archive = archive
	return report
}

func (archive Archive) Entry(name string) (model.LocalPath, bool) {
	value, ok := archive.Entries[path.Clean(strings.ReplaceAll(name, "\\", "/"))]
	return value, ok
}

func extractEntry(
	file *zip.File,
	target string,
	limit int64,
) (int64, error) {
	input, err := file.Open()
	if err != nil {
		return 0, err
	}
	defer input.Close()
	output, err := os.OpenFile(target, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		return 0, err
	}
	written, copyErr := io.Copy(output, io.LimitReader(input, limit+1))
	closeErr := output.Close()
	if copyErr != nil {
		return written, copyErr
	}
	if closeErr != nil {
		return written, closeErr
	}
	if written > limit {
		return written, fmt.Errorf("entry exceeds %d bytes", limit)
	}
	return written, nil
}

func safeEntryName(value string) (string, bool) {
	value = strings.ReplaceAll(value, "\\", "/")
	if value == "" || strings.HasPrefix(value, "/") ||
		len(value) >= 2 && value[1] == ':' {
		return "", false
	}
	clean := path.Clean(value)
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return "", false
	}
	return clean, true
}

func fileHash(value string) (string, error) {
	file, err := os.Open(value)
	if err != nil {
		return "", err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func loadArchive(root, hash string) (*Archive, bool) {
	data, err := os.ReadFile(filepath.Join(root, "manifest.json"))
	if err != nil {
		return nil, false
	}
	var value manifest
	if json.Unmarshal(data, &value) != nil || value.Hash != hash {
		return nil, false
	}
	archive := &Archive{
		Hash:    hash,
		Root:    model.LocalPath(root),
		Entries: make(map[string]model.LocalPath, len(value.Entries)),
	}
	for _, entry := range value.Entries {
		entryPath := model.LocalPath(
			filepath.Join(root, "entries", filepath.FromSlash(entry)),
		)
		if info, err := os.Stat(string(entryPath)); err != nil || !info.Mode().IsRegular() {
			return nil, false
		}
		archive.Entries[entry] = entryPath
	}
	return archive, true
}

func archiveError(
	report Report,
	code string,
	err error,
	localPath string,
) Report {
	return archiveMessage(report, code, err.Error(), localPath)
}

func archiveMessage(
	report Report,
	code string,
	message string,
	localPath string,
) Report {
	report.Diagnostics = append(report.Diagnostics, model.Diagnostic{
		Code:      code,
		Severity:  model.SeverityWarning,
		Stage:     "generated-store",
		Message:   message,
		LocalPath: localPath,
	})
	return report
}

package pack

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/jni"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

type Diagnostic struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Path    string `json:"path,omitempty"`
}

type Result struct {
	Valid       bool         `json:"valid"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

func Validate(root string) Result {
	result := Result{Valid: false, Diagnostics: []Diagnostic{}}
	required := []string{"manifest.json", "checksums.sha256", filepath.Join("generated", "manifest.json"), filepath.Join("modules", "catalog.json"), filepath.Join("modules", "dependency-graph.json"), filepath.Join("artifacts", "catalog.json"), filepath.Join("aidl", "symbols.json"), filepath.Join("aidl", "index.json"), filepath.Join("aidl", "modules.json")}
	for _, file := range required {
		if _, err := os.Stat(filepath.Join(root, file)); err != nil {
			result.Diagnostics = append(result.Diagnostics, Diagnostic{Code: "PACK_FILE_MISSING", Message: err.Error(), Path: file})
		}
	}
	if len(result.Diagnostics) > 0 {
		return result
	}
	var manifest struct {
		SchemaVersion int      `json:"schemaVersion"`
		Packs         []string `json:"packs"`
	}
	if err := readJSON(root, "manifest.json", &manifest); err != nil {
		return invalid("PACK_MANIFEST_INVALID", err.Error(), "manifest.json")
	}
	if manifest.SchemaVersion != 1 {
		return invalid("PACK_SCHEMA_UNSUPPORTED", fmt.Sprintf("unsupported schema: %d", manifest.SchemaVersion), "manifest.json")
	}
	if _, err := modules.LoadPackCatalog(root); err != nil {
		return invalid("PACK_MODULE_CATALOG_INVALID", err.Error(), filepath.Join("modules", "catalog.json"))
	}
	var artifactsCatalog map[string]any
	if err := readJSON(root, filepath.Join("artifacts", "catalog.json"), &artifactsCatalog); err != nil || artifactsCatalog == nil {
		return invalid("PACK_ARTIFACT_CATALOG_INVALID", "artifact catalog is invalid", filepath.Join("artifacts", "catalog.json"))
	}
	var aidlIndex aidl.Index
	if err := readJSON(root, filepath.Join("aidl", "index.json"), &aidlIndex); err != nil || aidlIndex.Documents == nil || aidlIndex.Symbols == nil {
		return invalid("PACK_AIDL_INDEX_INVALID", "AIDL index is invalid", filepath.Join("aidl", "index.json"))
	}
	var symbols map[string][]aidl.Symbol
	if err := readJSON(root, filepath.Join("aidl", "symbols.json"), &symbols); err != nil || symbols == nil {
		return invalid("PACK_AIDL_SYMBOLS_INVALID", "AIDL symbols are invalid", filepath.Join("aidl", "symbols.json"))
	}
	var generated []struct {
		Path    string   `json:"path"`
		Entries []string `json:"entries"`
	}
	if err := readJSON(root, filepath.Join("generated", "manifest.json"), &generated); err != nil || generated == nil {
		return invalid("PACK_GENERATED_MANIFEST_INVALID", "generated manifest is invalid", filepath.Join("generated", "manifest.json"))
	}
	for _, entry := range generated {
		name := filepath.ToSlash(entry.Path)
		if !strings.HasPrefix(name, "generated/srcjars/") || strings.Contains(name, "../") {
			return invalid("PACK_GENERATED_PATH_INVALID", "generated path is invalid", name)
		}
		if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(name))); err != nil {
			return invalid("PACK_GENERATED_ARCHIVE_MISSING", err.Error(), name)
		}
		actual, err := archiveEntries(filepath.Join(root, filepath.FromSlash(name)))
		if err != nil {
			return invalid("PACK_GENERATED_ARCHIVE_INVALID", err.Error(), name)
		}
		if !sameStrings(actual, entry.Entries) {
			return invalid("PACK_GENERATED_ENTRIES_MISMATCH", "generated manifest entries do not match archive", name)
		}
	}
	if contains(manifest.Packs, "native") {
		var value jni.Index
		if err := readJSON(root, filepath.Join("jni", "index.json"), &value); err != nil || value.Links == nil {
			return invalid("PACK_JNI_INVALID", err.Error(), filepath.Join("jni", "index.json"))
		}
	}
	checksums, err := os.ReadFile(filepath.Join(root, "checksums.sha256"))
	if err != nil {
		return invalid("PACK_CHECKSUMS_READ_FAILED", err.Error(), "checksums.sha256")
	}
	covered := map[string]bool{}
	for _, line := range strings.Split(strings.TrimSpace(string(checksums)), "\n") {
		parts := strings.SplitN(line, "  ", 2)
		if len(parts) != 2 {
			return invalid("PACK_CHECKSUM_LINE_INVALID", "invalid checksum line", "checksums.sha256")
		}
		name := filepath.ToSlash(parts[1])
		if name == "" || strings.HasPrefix(name, "/") || strings.Contains(name, "../") || covered[name] {
			return invalid("PACK_CHECKSUM_PATH_INVALID", "invalid checksum path", name)
		}
		covered[name] = true
		data, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(name)))
		if err != nil {
			return invalid("PACK_CHECKSUM_FILE_MISSING", err.Error(), name)
		}
		sum := sha256.Sum256(data)
		if hex.EncodeToString(sum[:]) != parts[0] {
			return invalid("PACK_CHECKSUM_MISMATCH", "checksum mismatch", name)
		}
	}
	files, err := filesUnder(root)
	if err != nil {
		return invalid("PACK_FILE_SCAN_FAILED", err.Error(), root)
	}
	for _, file := range files {
		if !covered[file] {
			return invalid("PACK_CHECKSUM_ENTRY_MISSING", "checksum entry missing", file)
		}
	}
	for _, file := range files {
		if strings.HasPrefix(file, "out/") && strings.HasSuffix(strings.ToLower(file), ".jar") {
			reader, err := zip.OpenReader(filepath.Join(root, filepath.FromSlash(file)))
			if err != nil {
				return invalid("PACK_JAVA_JAR_INVALID", err.Error(), file)
			}
			_ = reader.Close()
		}
	}
	result.Valid = true
	return result
}

func archiveEntries(pathName string) ([]string, error) {
	reader, err := zip.OpenReader(pathName)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	entries := make([]string, 0, len(reader.File))
	for _, file := range reader.File {
		name := strings.TrimPrefix(strings.ReplaceAll(file.Name, "\\", "/"), "/")
		if name == "" || name == "." || strings.Contains(name, "../") || strings.HasPrefix(name, "../") {
			return nil, fmt.Errorf("unsafe archive entry: %s", file.Name)
		}
		if file.FileInfo().IsDir() {
			continue
		}
		entries = append(entries, name)
	}
	sort.Strings(entries)
	return entries, nil
}

func sameStrings(left, right []string) bool {
	copyLeft := append([]string(nil), left...)
	copyRight := append([]string(nil), right...)
	sort.Strings(copyLeft)
	sort.Strings(copyRight)
	return reflect.DeepEqual(copyLeft, copyRight)
}

func invalid(code, message, path string) Result {
	return Result{Diagnostics: []Diagnostic{{Code: code, Message: message, Path: path}}}
}
func contains(values []string, target string) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}
func readJSON(root, name string, destination any) error {
	data, err := os.ReadFile(filepath.Join(root, name))
	if err != nil {
		return err
	}
	return json.Unmarshal(data, destination)
}
func filesUnder(root string) ([]string, error) {
	result := []string{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || entry.Name() == "checksums.sha256" {
			return nil
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		result = append(result, filepath.ToSlash(relative))
		return nil
	})
	return result, err
}

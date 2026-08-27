package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/generated"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type GeneratedMaterializeParams struct {
	MetadataPackRoot string   `json:"metadataPackRoot"`
	MetadataRoot     string   `json:"metadataRoot,omitempty"`
	Archives         []string `json:"archives,omitempty"`
	CacheRoot        string   `json:"cacheRoot"`
}

type GeneratedSourceRoot struct {
	Archive string `json:"archive"`
	Root    string `json:"root"`
	Hash    string `json:"hash"`
	Entries int    `json:"entries"`
}

type GeneratedMaterializeResult struct {
	SourceRoots []GeneratedSourceRoot `json:"sourceRoots"`
	Diagnostics []model.Diagnostic    `json:"diagnostics"`
}

type generatedManifestEntry struct {
	Path string `json:"path"`
}

func (server *Server) generatedMaterialize(request jsonrpc.Request) jsonrpc.Response {
	var params GeneratedMaterializeParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	result := GeneratedMaterializeResult{SourceRoots: []GeneratedSourceRoot{}, Diagnostics: []model.Diagnostic{}}
	packRoot := strings.TrimSpace(params.MetadataPackRoot)
	metadataRoot := strings.TrimSpace(params.MetadataRoot)
	cacheRoot := strings.TrimSpace(params.CacheRoot)
	if (packRoot == "" && metadataRoot == "") || cacheRoot == "" {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{
			Code: "GENERATED_ROOT_MISSING", Severity: model.SeverityError,
			Stage: "generated-rpc", Message: "metadataPackRoot or metadataRoot and cacheRoot are required",
		})
		return jsonrpc.Success(request.ID, result)
	}
	if packRoot != "" {
		result.Diagnostics = append(result.Diagnostics, diagnosticsToModel(server.validatePackIfDeclared(packRoot))...)
		if len(result.Diagnostics) > 0 {
			return jsonrpc.Success(request.ID, result)
		}
	} else {
		roots, diagnostics := materializeLocalGeneratedSources(metadataRoot, params.Archives, cacheRoot)
		result.SourceRoots = roots
		result.Diagnostics = append(result.Diagnostics, diagnostics...)
		return jsonrpc.Success(request.ID, result)
	}
	roots, _, diagnostics := materializeGeneratedSources(packRoot, cacheRoot)
	result.SourceRoots = roots
	result.Diagnostics = append(result.Diagnostics, diagnostics...)
	return jsonrpc.Success(request.ID, result)
}

func materializeGeneratedSources(packRoot, cacheRoot string) ([]GeneratedSourceRoot, []aidl.ArchiveEntry, []model.Diagnostic) {
	roots := []GeneratedSourceRoot{}
	archiveEntries := []aidl.ArchiveEntry{}
	diagnostics := []model.Diagnostic{}
	data, err := os.ReadFile(filepath.Join(packRoot, "generated", "manifest.json"))
	if err != nil {
		diagnostics = append(diagnostics, generatedDiagnostic(
			"GENERATED_MANIFEST_READ_FAILED", err.Error(), filepath.Join(packRoot, "generated", "manifest.json"),
		))
		return roots, archiveEntries, diagnostics
	}
	var entries []generatedManifestEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		diagnostics = append(diagnostics, generatedDiagnostic(
			"GENERATED_MANIFEST_INVALID", err.Error(), filepath.Join(packRoot, "generated", "manifest.json"),
		))
		return roots, archiveEntries, diagnostics
	}
	for _, entry := range entries {
		relative := filepath.ToSlash(strings.TrimSpace(entry.Path))
		if relative == "" || !strings.HasPrefix(relative, "generated/srcjars/") || strings.Contains(relative, "..") {
			diagnostics = append(diagnostics, generatedDiagnostic(
				"GENERATED_ARCHIVE_PATH_INVALID", "manifest archive path is outside generated/srcjars", relative,
			))
			continue
		}
		archive := filepath.Join(packRoot, filepath.FromSlash(relative))
		report := generated.ExtractSrcJar(archive, cacheRoot, generated.DefaultLimits())
		diagnostics = append(diagnostics, report.Diagnostics...)
		if report.Archive == nil {
			continue
		}
		roots = append(roots, GeneratedSourceRoot{
			Archive: relative, Root: string(report.Archive.Root), Hash: report.Archive.Hash,
			Entries: len(report.Archive.Entries),
		})
		archiveRef := metadata.NewPathRef(relative)
		for name, localPath := range report.Archive.Entries {
			archiveEntries = append(archiveEntries, aidl.ArchiveEntry{Archive: archiveRef, Entry: name, LocalPath: localPath})
		}
	}
	return roots, archiveEntries, diagnostics
}

func materializeLocalGeneratedSources(metadataRoot string, archives []string, cacheRoot string) ([]GeneratedSourceRoot, []model.Diagnostic) {
	roots := []GeneratedSourceRoot{}
	diagnostics := []model.Diagnostic{}
	root, err := filepath.Abs(metadataRoot)
	if err != nil {
		return roots, []model.Diagnostic{generatedDiagnostic("GENERATED_OUT_ROOT_INVALID", err.Error(), metadataRoot)}
	}
	root = filepath.Clean(root)
	seen := map[string]bool{}
	for _, raw := range archives {
		archive, ok := localGeneratedArchive(root, raw)
		if !ok {
			diagnostics = append(diagnostics, generatedDiagnostic(
				"GENERATED_ARCHIVE_PATH_INVALID", "archive is outside metadataRoot", raw,
			))
			continue
		}
		if seen[archive] {
			continue
		}
		seen[archive] = true
		report := generated.ExtractSrcJar(archive, cacheRoot, generated.DefaultLimits())
		diagnostics = append(diagnostics, report.Diagnostics...)
		if report.Archive == nil {
			continue
		}
		roots = append(roots, GeneratedSourceRoot{
			Archive: archive, Root: string(report.Archive.Root), Hash: report.Archive.Hash,
			Entries: len(report.Archive.Entries),
		})
	}
	return roots, diagnostics
}

func localGeneratedArchive(metadataRoot, raw string) (string, bool) {
	value := strings.TrimSpace(raw)
	if value == "" || strings.Contains(filepath.ToSlash(value), "..") || !strings.HasSuffix(strings.ToLower(value), ".srcjar") {
		return "", false
	}
	var candidate string
	if filepath.IsAbs(value) {
		candidate = filepath.Clean(value)
	} else {
		normalized := filepath.ToSlash(value)
		normalized = strings.TrimPrefix(normalized, "out/")
		candidate = filepath.Join(metadataRoot, filepath.FromSlash(normalized))
	}
	relative, err := filepath.Rel(metadataRoot, candidate)
	if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) || filepath.IsAbs(relative) {
		return "", false
	}
	return candidate, true
}

func generatedDiagnostic(code, message, pathName string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "generated-rpc", Message: message, LocalPath: pathName}
}

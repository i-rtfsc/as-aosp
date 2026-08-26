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
	MetadataPackRoot string `json:"metadataPackRoot"`
	CacheRoot        string `json:"cacheRoot"`
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
	cacheRoot := strings.TrimSpace(params.CacheRoot)
	if packRoot == "" || cacheRoot == "" {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{
			Code: "GENERATED_ROOT_MISSING", Severity: model.SeverityError,
			Stage: "generated-rpc", Message: "metadataPackRoot and cacheRoot are required",
		})
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

func generatedDiagnostic(code, message, pathName string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "generated-rpc", Message: message, LocalPath: pathName}
}

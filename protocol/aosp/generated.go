package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/generated"
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
	data, err := os.ReadFile(filepath.Join(packRoot, "generated", "manifest.json"))
	if err != nil {
		result.Diagnostics = append(result.Diagnostics, generatedDiagnostic(
			"GENERATED_MANIFEST_READ_FAILED", err.Error(), filepath.Join(packRoot, "generated", "manifest.json"),
		))
		return jsonrpc.Success(request.ID, result)
	}
	var entries []generatedManifestEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		result.Diagnostics = append(result.Diagnostics, generatedDiagnostic(
			"GENERATED_MANIFEST_INVALID", err.Error(), filepath.Join(packRoot, "generated", "manifest.json"),
		))
		return jsonrpc.Success(request.ID, result)
	}
	for _, entry := range entries {
		relative := filepath.ToSlash(strings.TrimSpace(entry.Path))
		if relative == "" || !strings.HasPrefix(relative, "generated/srcjars/") || strings.Contains(relative, "..") {
			result.Diagnostics = append(result.Diagnostics, generatedDiagnostic(
				"GENERATED_ARCHIVE_PATH_INVALID", "manifest archive path is outside generated/srcjars", relative,
			))
			continue
		}
		archive := filepath.Join(packRoot, filepath.FromSlash(relative))
		report := generated.ExtractSrcJar(archive, cacheRoot, generated.DefaultLimits())
		result.Diagnostics = append(result.Diagnostics, report.Diagnostics...)
		if report.Archive == nil {
			continue
		}
		result.SourceRoots = append(result.SourceRoots, GeneratedSourceRoot{
			Archive: relative, Root: string(report.Archive.Root), Hash: report.Archive.Hash,
			Entries: len(report.Archive.Entries),
		})
	}
	return jsonrpc.Success(request.ID, result)
}

func generatedDiagnostic(code, message, pathName string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "generated-rpc", Message: message, LocalPath: pathName}
}

package aosp

import (
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/generated"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type GeneratedMaterializeParams struct {
	OutRoot   string   `json:"outRoot"`
	Archives  []string `json:"archives,omitempty"`
	CacheRoot string   `json:"cacheRoot"`
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

func (server *Server) generatedMaterialize(request jsonrpc.Request) jsonrpc.Response {
	var params GeneratedMaterializeParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	result := GeneratedMaterializeResult{SourceRoots: []GeneratedSourceRoot{}, Diagnostics: []model.Diagnostic{}}
	outRoot := strings.TrimSpace(params.OutRoot)
	cacheRoot := strings.TrimSpace(params.CacheRoot)
	if outRoot == "" || cacheRoot == "" {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{
			Code: "GENERATED_ROOT_MISSING", Severity: model.SeverityError,
			Stage: "generated-rpc", Message: "outRoot and cacheRoot are required",
		})
		return jsonrpc.Success(request.ID, result)
	}
	roots, diagnostics := materializeLocalGeneratedSources(outRoot, params.Archives, cacheRoot)
	result.SourceRoots = roots
	result.Diagnostics = append(result.Diagnostics, diagnostics...)
	return jsonrpc.Success(request.ID, result)
}

func materializeLocalGeneratedSources(outRoot string, archives []string, cacheRoot string) ([]GeneratedSourceRoot, []model.Diagnostic) {
	roots := []GeneratedSourceRoot{}
	diagnostics := []model.Diagnostic{}
	root, err := filepath.Abs(outRoot)
	if err != nil {
		return roots, []model.Diagnostic{generatedDiagnostic("GENERATED_OUT_ROOT_INVALID", err.Error(), outRoot)}
	}
	root = filepath.Clean(root)
	seen := map[string]bool{}
	for _, raw := range archives {
		archive, ok := localGeneratedArchive(root, raw)
		if !ok {
			diagnostics = append(diagnostics, generatedDiagnostic(
				"GENERATED_ARCHIVE_PATH_INVALID", "archive is outside outRoot", raw,
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

func localGeneratedArchive(outRoot, raw string) (string, bool) {
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
		candidate = filepath.Join(outRoot, filepath.FromSlash(normalized))
	}
	relative, err := filepath.Rel(outRoot, candidate)
	if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) || filepath.IsAbs(relative) {
		return "", false
	}
	return candidate, true
}

func generatedDiagnostic(code, message, pathName string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "generated-rpc", Message: message, LocalPath: pathName}
}

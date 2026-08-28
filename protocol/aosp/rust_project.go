package aosp

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/rustproject"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type RustProjectParams struct {
	WorkspaceViewRoot string              `json:"workspaceViewRoot"`
	OutRoot           string              `json:"outRoot"`
	CacheRoot         string              `json:"cacheRoot"`
	PathMappings      []model.PathMapping `json:"pathMappings,omitempty"`
}

type RustProjectResult struct {
	ProjectPath string             `json:"projectPath,omitempty"`
	Crates      int                `json:"crates"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

func (server *Server) rustProject(request jsonrpc.Request) jsonrpc.Response {
	var params RustProjectParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	if strings.TrimSpace(params.OutRoot) == "" || strings.TrimSpace(params.CacheRoot) == "" {
		return jsonrpc.Success(request.ID, RustProjectResult{Diagnostics: []model.Diagnostic{{
			Code: "RUST_PROJECT_ROOT_MISSING", Severity: model.SeverityError,
			Stage: "rust-project", Message: "outRoot and cacheRoot are required",
		}}})
	}
	source := filepath.Join(params.OutRoot, "soong", "rust-project.json")
	data, err := os.ReadFile(source)
	if err != nil {
		return jsonrpc.Success(request.ID, RustProjectResult{Diagnostics: []model.Diagnostic{{
			Code: "RUST_PROJECT_READ_FAILED", Severity: model.SeverityWarning,
			Stage: "rust-project", Message: err.Error(), LocalPath: source,
		}}})
	}
	normalized := rustproject.Normalize(
		data, params.WorkspaceViewRoot, params.OutRoot, params.PathMappings,
	)
	result := RustProjectResult{
		Crates: normalized.Crates,
		Diagnostics: append(
			make([]model.Diagnostic, 0, len(normalized.Diagnostics)),
			normalized.Diagnostics...,
		),
	}
	if len(normalized.Data) == 0 {
		return jsonrpc.Success(request.ID, result)
	}
	directory := filepath.Join(params.CacheRoot, "rust")
	if err := os.MkdirAll(directory, 0o755); err != nil {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{
			Code: "RUST_PROJECT_WRITE_FAILED", Severity: model.SeverityWarning,
			Stage: "rust-project", Message: err.Error(), LocalPath: directory,
		})
		return jsonrpc.Success(request.ID, result)
	}
	target := filepath.Join(directory, "rust-project.json")
	temporary := target + ".tmp"
	if err := os.WriteFile(temporary, normalized.Data, 0o644); err != nil {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{
			Code: "RUST_PROJECT_WRITE_FAILED", Severity: model.SeverityWarning,
			Stage: "rust-project", Message: err.Error(), LocalPath: target,
		})
		return jsonrpc.Success(request.ID, result)
	}
	if err := os.Rename(temporary, target); err != nil {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{
			Code: "RUST_PROJECT_WRITE_FAILED", Severity: model.SeverityWarning,
			Stage: "rust-project", Message: err.Error(), LocalPath: target,
		})
		return jsonrpc.Success(request.ID, result)
	}
	result.ProjectPath = target
	return jsonrpc.Success(request.ID, result)
}

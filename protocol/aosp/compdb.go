package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type CompDBOverlayParams struct {
	MetadataRoot      string              `json:"metadataRoot"`
	WorkspaceViewRoot string              `json:"workspaceViewRoot"`
	PathMappings      []model.PathMapping `json:"pathMappings,omitempty"`
}
type CompDBOverlayResult struct {
	Entries     []CompDBOverlayEntry `json:"entries"`
	Diagnostics []model.Diagnostic   `json:"diagnostics"`
}
type CompDBOverlayEntry struct {
	Directory string   `json:"directory"`
	Arguments []string `json:"arguments"`
	File      string   `json:"file"`
	Output    string   `json:"output,omitempty"`
}

func (server *Server) compdbOverlay(request jsonrpc.Request) jsonrpc.Response {
	var params CompDBOverlayParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	root, err := filepath.Abs(params.MetadataRoot)
	if err != nil {
		return jsonrpc.Success(request.ID, CompDBOverlayResult{Diagnostics: []model.Diagnostic{{Code: "COMPDB_ROOT_INVALID", Severity: model.SeverityError, Stage: "compdb-rpc", Message: err.Error()}}})
	}
	data, err := readCompDB(root)
	if err != nil {
		return jsonrpc.Success(request.ID, CompDBOverlayResult{Diagnostics: []model.Diagnostic{{Code: "COMPDB_READ_FAILED", Severity: model.SeverityError, Stage: "compdb-rpc", Message: err.Error()}}})
	}
	var entries []CompDBOverlayEntry
	if err := json.Unmarshal(data, &entries); err != nil {
		return jsonrpc.Success(request.ID, CompDBOverlayResult{Diagnostics: []model.Diagnostic{{Code: "COMPDB_PARSE_FAILED", Severity: model.SeverityError, Stage: "compdb-rpc", Message: err.Error()}}})
	}
	diagnostics := make([]model.Diagnostic, 0)
	for index := range entries {
		entries[index].Directory, diagnostics = mapCompDBPath(entries[index].Directory, params.PathMappings, diagnostics)
		entries[index].File, diagnostics = mapCompDBPath(entries[index].File, params.PathMappings, diagnostics)
		entries[index].Output, diagnostics = mapCompDBPath(entries[index].Output, params.PathMappings, diagnostics)
		for arg := range entries[index].Arguments {
			entries[index].Arguments[arg], diagnostics = mapCompDBArgument(entries[index].Arguments[arg], params.PathMappings, diagnostics)
		}
	}
	return jsonrpc.Success(request.ID, CompDBOverlayResult{Entries: entries, Diagnostics: diagnostics})
}

func readCompDB(root string) ([]byte, error) {
	paths := []string{filepath.Join(root, "compile_commands.json"), filepath.Join(root, "development", "ide", "compdb", "compile_commands.json"), filepath.Join(root, "soong", "development", "ide", "compdb", "compile_commands.json")}
	for _, path := range paths {
		if data, err := os.ReadFile(path); err == nil {
			return data, nil
		}
	}
	return nil, os.ErrNotExist
}

func mapCompDBPath(value string, mappings []model.PathMapping, diagnostics []model.Diagnostic) (string, []model.Diagnostic) {
	if value == "" || !filepath.IsAbs(value) {
		return value, diagnostics
	}
	for _, mapping := range mappings {
		external := filepath.Clean(mapping.ExternalPrefix)
		if value == external || strings.HasPrefix(value, external+string(os.PathSeparator)) || strings.HasPrefix(strings.ReplaceAll(value, "\\", "/"), strings.TrimRight(strings.ReplaceAll(external, "\\", "/"), "/")+"/") {
			return filepath.Clean(filepath.Join(mapping.LocalPrefix, strings.TrimPrefix(strings.TrimPrefix(value, external), string(os.PathSeparator)))), diagnostics
		}
	}
	diagnostics = append(diagnostics, model.Diagnostic{Code: "COMPDB_PATH_UNMAPPED", Severity: model.SeverityWarning, Stage: "compdb-rpc", Message: "compdb path was not mapped: " + value})
	return value, diagnostics
}

func mapCompDBArgument(value string, mappings []model.PathMapping, diagnostics []model.Diagnostic) (string, []model.Diagnostic) {
	for _, prefix := range []string{"-I", "-isystem", "--sysroot", "-include"} {
		if strings.HasPrefix(value, prefix) {
			mapped, next := mapCompDBPath(strings.TrimPrefix(value, prefix), mappings, diagnostics)
			return prefix + mapped, next
		}
	}
	return value, diagnostics
}

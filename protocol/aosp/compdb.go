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
	Entries       []CompDBOverlayEntry `json:"entries"`
	ResponseFiles map[string]string    `json:"responseFiles,omitempty"`
	Diagnostics   []model.Diagnostic   `json:"diagnostics"`
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
	responseFiles := make(map[string]string)
	for index := range entries {
		entries[index].Directory, diagnostics = mapCompDBPath(entries[index].Directory, params.PathMappings, diagnostics)
		entries[index].File, diagnostics = mapCompDBPath(entries[index].File, params.PathMappings, diagnostics)
		entries[index].Output, diagnostics = mapCompDBPath(entries[index].Output, params.PathMappings, diagnostics)
		entries[index].Arguments, responseFiles, diagnostics = mapCompDBArguments(entries[index].Arguments, responseFiles, params.PathMappings, diagnostics)
	}
	return jsonrpc.Success(request.ID, CompDBOverlayResult{Entries: entries, ResponseFiles: responseFiles, Diagnostics: diagnostics})
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
	normalizedValue := strings.ReplaceAll(value, "\\", "/")
	if value == "" || (!filepath.IsAbs(value) && !isCompDBAbsolute(normalizedValue)) {
		return value, diagnostics
	}
	selected := -1
	selectedPrefix := ""
	for index, mapping := range mappings {
		external := strings.TrimRight(strings.ReplaceAll(mapping.ExternalPrefix, "\\", "/"), "/")
		if normalizedValue == external || strings.HasPrefix(normalizedValue, external+"/") {
			if len(external) > len(selectedPrefix) {
				selected = index
				selectedPrefix = external
			}
		}
	}
	if selected >= 0 {
		suffix := strings.TrimPrefix(normalizedValue, selectedPrefix)
		local := strings.ReplaceAll(mappings[selected].LocalPrefix, "\\", "/") + suffix
		return filepath.Clean(filepath.FromSlash(local)), diagnostics
	}
	diagnostics = append(diagnostics, model.Diagnostic{Code: "COMPDB_PATH_UNMAPPED", Severity: model.SeverityWarning, Stage: "compdb-rpc", Message: "compdb path was not mapped: " + value})
	return value, diagnostics
}

func isCompDBAbsolute(value string) bool {
	return strings.HasPrefix(value, "/") || strings.HasPrefix(value, "//") ||
		(len(value) >= 3 && value[1] == ':' && value[2] == '/')
}

func mapCompDBArguments(arguments []string, responseFiles map[string]string, mappings []model.PathMapping, diagnostics []model.Diagnostic) ([]string, map[string]string, []model.Diagnostic) {
	result := append([]string(nil), arguments...)
	separate := map[string]bool{"-I": true, "-isystem": true, "--sysroot": true, "-include": true}
	for index := 0; index < len(result); index++ {
		value := result[index]
		if separate[value] && index+1 < len(result) {
			result[index+1], diagnostics = mapCompDBPath(result[index+1], mappings, diagnostics)
			index++
			continue
		}
		if strings.HasPrefix(value, "@") {
			mapped, next := mapCompDBPath(strings.TrimPrefix(value, "@"), mappings, diagnostics)
			if data, err := os.ReadFile(strings.TrimPrefix(value, "@")); err == nil {
				parts := tokenizeResponseFile(string(data))
				mappedParts, _, nestedDiagnostics := mapCompDBArguments(parts, responseFiles, mappings, next)
				next = nestedDiagnostics
				responseFiles[mapped] = strings.Join(mappedParts, "\n") + "\n"
			} else {
				next = append(next, model.Diagnostic{Code: "COMPDB_RESPONSE_READ_FAILED", Severity: model.SeverityWarning, Stage: "compdb-rpc", Message: "response file could not be read: " + strings.TrimPrefix(value, "@")})
			}
			result[index], diagnostics = "@"+mapped, next
			continue
		}
		for _, prefix := range []string{"-I", "-isystem", "--sysroot=", "-include"} {
			if strings.HasPrefix(value, prefix) && len(value) > len(prefix) {
				mapped, next := mapCompDBPath(strings.TrimPrefix(value, prefix), mappings, diagnostics)
				result[index], diagnostics = prefix+mapped, next
				break
			}
		}
	}
	return result, responseFiles, diagnostics
}

func tokenizeResponseFile(value string) []string {
	result := make([]string, 0)
	current := strings.Builder{}
	quoted := byte(0)
	escaped := false
	flush := func() {
		if current.Len() > 0 {
			result = append(result, current.String())
			current.Reset()
		}
	}
	for index := 0; index < len(value); index++ {
		ch := value[index]
		if escaped {
			current.WriteByte(ch)
			escaped = false
			continue
		}
		if ch == '\\' {
			escaped = true
			continue
		}
		if quoted != 0 {
			if ch == quoted {
				quoted = 0
			} else {
				current.WriteByte(ch)
			}
			continue
		}
		if ch == '\'' || ch == '"' {
			quoted = ch
			continue
		}
		if ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' {
			flush()
			continue
		}
		current.WriteByte(ch)
	}
	if escaped {
		current.WriteByte('\\')
	}
	flush()
	return result
}

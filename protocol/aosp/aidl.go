package aosp

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type AIDLResolveParams struct {
	WorkspaceViewRoot string `json:"workspaceViewRoot"`
	CurrentPath       string `json:"currentPath"`
	Name              string `json:"name"`
}

type AIDLResolveResult struct {
	Resolution  aidl.Resolution    `json:"resolution"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

func (server *Server) aidlResolve(request jsonrpc.Request) jsonrpc.Response {
	var params AIDLResolveParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	current, err := model.NormalizeRepoPath(params.CurrentPath)
	if err != nil {
		return jsonrpc.Success(request.ID, AIDLResolveResult{
			Diagnostics: []model.Diagnostic{{
				Code: "AIDL_INVALID_REPO_PATH", Severity: model.SeverityError,
				Stage: "aidl-rpc", Message: err.Error(),
			}},
		})
	}
	index, diagnostics := scanAIDLWorkspace(params.WorkspaceViewRoot)
	document, ok := index.Documents[current]
	if !ok {
		diagnostics = append(diagnostics, model.Diagnostic{
			Code: "AIDL_CURRENT_FILE_NOT_FOUND", Severity: model.SeverityWarning,
			Stage: "aidl-rpc", Message: "current AIDL file is not visible: " + string(current),
			RepoPath: current,
		})
		return jsonrpc.Success(request.ID, AIDLResolveResult{
			Resolution:  aidl.Resolution{Name: params.Name, State: aidl.ResolutionNotFound},
			Diagnostics: diagnostics,
		})
	}
	resolution := index.Resolve(document, params.Name)
	diagnostics = append(diagnostics, index.Validate(document)...)
	return jsonrpc.Success(request.ID, AIDLResolveResult{Resolution: resolution, Diagnostics: diagnostics})
}

func scanAIDLWorkspace(root string) (aidl.Index, []model.Diagnostic) {
	index := aidl.NewIndex()
	diagnostics := make([]model.Diagnostic, 0)
	absolute, err := filepath.Abs(root)
	if err != nil {
		return index, []model.Diagnostic{{Code: "AIDL_WORKSPACE_INVALID", Severity: model.SeverityError, Stage: "aidl-workspace", Message: err.Error()}}
	}
	info, err := os.Stat(absolute)
	if err != nil || !info.IsDir() {
		message := "workspace view root is not a directory"
		if err != nil {
			message = err.Error()
		}
		return index, []model.Diagnostic{{Code: "AIDL_WORKSPACE_INVALID", Severity: model.SeverityError, Stage: "aidl-workspace", Message: message}}
	}
	var walk func(string, string, map[string]bool) error
	walk = func(current, relative string, ancestors map[string]bool) error {
		real, evalErr := filepath.EvalSymlinks(current)
		if evalErr == nil {
			real = filepath.Clean(real)
			if ancestors[real] {
				return nil
			}
			ancestors = cloneStringSet(ancestors)
			ancestors[real] = true
		}
		entries, readErr := os.ReadDir(current)
		if readErr != nil {
			return readErr
		}
		for _, entry := range entries {
			childRelative := filepath.Join(relative, entry.Name())
			if isJNIExcluded(childRelative) {
				continue
			}
			child := filepath.Join(current, entry.Name())
			info, statErr := os.Stat(child)
			if statErr != nil {
				diagnostics = append(diagnostics, aidlWorkspaceDiagnostic("AIDL_WORKSPACE_READ", statErr.Error()))
				continue
			}
			if info.IsDir() {
				if err := walk(child, childRelative, ancestors); err != nil {
					diagnostics = append(diagnostics, aidlWorkspaceDiagnostic("AIDL_WORKSPACE_READ", err.Error()))
				}
				continue
			}
			if strings.ToLower(filepath.Ext(entry.Name())) != ".aidl" {
				continue
			}
			source, readErr := os.ReadFile(child)
			if readErr != nil {
				diagnostics = append(diagnostics, aidlWorkspaceDiagnostic("AIDL_WORKSPACE_READ", readErr.Error()))
				continue
			}
			repoPath, pathErr := model.NormalizeRepoPath(filepath.ToSlash(childRelative))
			if pathErr != nil {
				diagnostics = append(diagnostics, aidlWorkspaceDiagnostic("AIDL_INVALID_REPO_PATH", pathErr.Error()))
				continue
			}
			document := aidl.Parse(repoPath, string(source))
			for index := range document.Diagnostics {
				if document.Diagnostics[index].RepoPath == "" {
					document.Diagnostics[index].RepoPath = repoPath
				}
			}
			diagnostics = append(diagnostics, document.Diagnostics...)
			index.Add(document)
		}
		return nil
	}
	if err := walk(absolute, "", map[string]bool{}); err != nil {
		diagnostics = append(diagnostics, aidlWorkspaceDiagnostic("AIDL_WORKSPACE_READ", err.Error()))
	}
	return index, diagnostics
}

func aidlWorkspaceDiagnostic(code, message string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "aidl-workspace", Message: message}
}

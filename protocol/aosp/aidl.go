package aosp

import (
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
	scan := scanWorkspaceSources(
		root,
		"aidl-workspace",
		"AIDL_WORKSPACE_INVALID",
		"AIDL_WORKSPACE_READ",
		func(extension string) bool { return extension == ".aidl" },
	)
	diagnostics := scan.Diagnostics
	for repoPath, source := range scan.Files {
		document := aidl.Parse(repoPath, source)
		for index := range document.Diagnostics {
			if document.Diagnostics[index].RepoPath == "" {
				document.Diagnostics[index].RepoPath = repoPath
			}
		}
		diagnostics = append(diagnostics, document.Diagnostics...)
		index.Add(document)
	}
	return index, diagnostics
}

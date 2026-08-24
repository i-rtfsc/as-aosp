package aosp

import (
	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type AIDLReferencesParams struct {
	WorkspaceViewRoot string `json:"workspaceViewRoot"`
	CurrentPath       string `json:"currentPath"`
	Name              string `json:"name"`
}
type AIDLReferencesResult struct {
	References  []aidl.Reference   `json:"references"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

func (server *Server) aidlReferences(request jsonrpc.Request) jsonrpc.Response {
	var params AIDLReferencesParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	current, err := model.NormalizeRepoPath(params.CurrentPath)
	if err != nil {
		return jsonrpc.Success(request.ID, AIDLReferencesResult{Diagnostics: []model.Diagnostic{{Code: "AIDL_INVALID_REPO_PATH", Severity: model.SeverityError, Stage: "aidl-rpc", Message: err.Error()}}})
	}
	index, diagnostics := server.buildAIDLIndex(params.WorkspaceViewRoot)
	document, ok := index.Documents[current]
	if !ok {
		diagnostics = append(diagnostics, model.Diagnostic{Code: "AIDL_CURRENT_FILE_NOT_FOUND", Severity: model.SeverityWarning, Stage: "aidl-rpc", Message: "current AIDL file is not visible: " + string(current), RepoPath: current})
		return jsonrpc.Success(request.ID, AIDLReferencesResult{Diagnostics: diagnostics})
	}
	resolution := index.Resolve(document, params.Name)
	if resolution.State != aidl.ResolutionFound || len(resolution.Symbols) != 1 {
		return jsonrpc.Success(request.ID, AIDLReferencesResult{Diagnostics: append(diagnostics, model.Diagnostic{Code: "AIDL_TYPE_NOT_FOUND", Severity: model.SeverityWarning, Stage: "aidl-rpc", Message: "AIDL type is not uniquely resolved: " + params.Name, RepoPath: current})})
	}
	return jsonrpc.Success(request.ID, AIDLReferencesResult{References: index.References(resolution.Name), Diagnostics: diagnostics})
}

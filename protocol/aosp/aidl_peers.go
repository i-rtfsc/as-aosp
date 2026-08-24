package aosp

import (
	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type AIDLPeersParams struct {
	WorkspaceViewRoot string `json:"workspaceViewRoot"`
	CurrentPath       string `json:"currentPath"`
	Name              string `json:"name"`
}

type AIDLPeersResult struct {
	Peers       []aidl.GeneratedPeer `json:"peers"`
	Diagnostics []model.Diagnostic   `json:"diagnostics"`
}

func (server *Server) aidlPeers(request jsonrpc.Request) jsonrpc.Response {
	var params AIDLPeersParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	current, err := model.NormalizeRepoPath(params.CurrentPath)
	if err != nil {
		return jsonrpc.Success(request.ID, AIDLPeersResult{Diagnostics: []model.Diagnostic{{
			Code: "AIDL_INVALID_REPO_PATH", Severity: model.SeverityError,
			Stage: "aidl-rpc", Message: err.Error(),
		}}})
	}
	index, diagnostics := server.buildAIDLIndex(params.WorkspaceViewRoot)
	document, ok := index.Documents[current]
	if !ok {
		diagnostics = append(diagnostics, model.Diagnostic{
			Code: "AIDL_CURRENT_FILE_NOT_FOUND", Severity: model.SeverityWarning,
			Stage: "aidl-rpc", Message: "current AIDL file is not visible: " + string(current),
			RepoPath: current,
		})
		return jsonrpc.Success(request.ID, AIDLPeersResult{Diagnostics: diagnostics})
	}
	resolution := index.Resolve(document, params.Name)
	if resolution.State != aidl.ResolutionFound || len(resolution.Symbols) != 1 {
		return jsonrpc.Success(request.ID, AIDLPeersResult{
			Diagnostics: append(diagnostics, model.Diagnostic{
				Code: "AIDL_TYPE_NOT_FOUND", Severity: model.SeverityWarning,
				Stage: "aidl-generated", Message: "AIDL type is not uniquely resolved: " + params.Name,
				RepoPath: current,
			}),
		})
	}
	allSources := scanWorkspaceSources(
		params.WorkspaceViewRoot,
		"aidl-generated",
		"AIDL_WORKSPACE_INVALID",
		"AIDL_WORKSPACE_READ",
		isJNISourceExtension,
	)
	generatedPaths := make([]metadata.PathRef, 0, len(allSources.Files))
	for path := range allSources.Files {
		generatedPaths = append(generatedPaths, metadata.NewPathRef(string(path)))
	}
	peerReport := index.GeneratedPeers(document, resolution.Name, aidl.PeerOptions{
		GeneratedPaths: generatedPaths,
		Exists:         func(metadata.PathRef) bool { return true },
	})
	diagnostics = append(diagnostics, allSources.Diagnostics...)
	diagnostics = append(diagnostics, peerReport.Diagnostics...)
	return jsonrpc.Success(request.ID, AIDLPeersResult{Peers: peerReport.Peers, Diagnostics: diagnostics})
}

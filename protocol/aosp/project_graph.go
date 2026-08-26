package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/artifacts"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
	"github.com/i-rtfsc/as-aosp/core/projects"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type ProjectGraphParams struct {
	MetadataPackRoot string             `json:"metadataPackRoot"`
	Targets          []model.TargetSpec `json:"targets"`
}

type ProjectGraphResult struct {
	Graph       projects.ProjectGraph `json:"graph"`
	Diagnostics []model.Diagnostic    `json:"diagnostics"`
}

func (server *Server) projectGraph(request jsonrpc.Request) jsonrpc.Response {
	var params ProjectGraphParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	root := strings.TrimSpace(params.MetadataPackRoot)
	if root == "" {
		return jsonrpc.Success(request.ID, ProjectGraphResult{Diagnostics: []model.Diagnostic{{Code: "PROJECT_GRAPH_PACK_MISSING", Severity: model.SeverityError, Stage: "project-graph-rpc", Message: "metadataPackRoot is required"}}})
	}
	if len(params.Targets) == 0 {
		return jsonrpc.Success(request.ID, ProjectGraphResult{Diagnostics: []model.Diagnostic{{Code: "PROJECT_GRAPH_TARGETS_MISSING", Severity: model.SeverityError, Stage: "project-graph-rpc", Message: "at least one target is required"}}})
	}
	if diagnostics := diagnosticsToModel(server.validatePackIfDeclared(root)); len(diagnostics) > 0 {
		return jsonrpc.Success(request.ID, ProjectGraphResult{Diagnostics: diagnostics})
	}
	moduleCatalog, err := modules.LoadPackCatalog(root)
	if err != nil {
		return jsonrpc.Success(request.ID, ProjectGraphResult{Diagnostics: []model.Diagnostic{{Code: "PROJECT_GRAPH_MODULE_CATALOG_INVALID", Severity: model.SeverityError, Stage: "project-graph-rpc", Message: err.Error()}}})
	}
	data, err := os.ReadFile(filepath.Join(root, "artifacts", "catalog.json"))
	if err != nil {
		return jsonrpc.Success(request.ID, ProjectGraphResult{Diagnostics: []model.Diagnostic{{Code: "PROJECT_GRAPH_ARTIFACT_CATALOG_READ_FAILED", Severity: model.SeverityError, Stage: "project-graph-rpc", Message: err.Error()}}})
	}
	var artifactCatalog artifacts.Catalog
	if err := json.Unmarshal(data, &artifactCatalog); err != nil {
		return jsonrpc.Success(request.ID, ProjectGraphResult{Diagnostics: []model.Diagnostic{{Code: "PROJECT_GRAPH_ARTIFACT_CATALOG_INVALID", Severity: model.SeverityError, Stage: "project-graph-rpc", Message: err.Error()}}})
	}
	graph := projects.Build(moduleCatalog, artifactCatalog, params.Targets)
	return jsonrpc.Success(request.ID, ProjectGraphResult{Graph: graph, Diagnostics: graph.Diagnostics})
}

package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
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

type ArtifactsCatalogParams struct {
	MetadataPackRoot string `json:"metadataPackRoot"`
	Kind             string `json:"kind,omitempty"`
	Query            string `json:"query,omitempty"`
}

type ArtifactsCatalogResult struct {
	Artifacts   []artifacts.Artifact `json:"artifacts"`
	Diagnostics []model.Diagnostic   `json:"diagnostics"`
}

func (server *Server) artifactsCatalog(request jsonrpc.Request) jsonrpc.Response {
	var params ArtifactsCatalogParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	root := strings.TrimSpace(params.MetadataPackRoot)
	result := ArtifactsCatalogResult{Artifacts: []artifacts.Artifact{}, Diagnostics: []model.Diagnostic{}}
	if root == "" {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{Code: "ARTIFACT_CATALOG_ROOT_MISSING", Severity: model.SeverityError, Stage: "artifact-rpc", Message: "metadataPackRoot is required"})
		return jsonrpc.Success(request.ID, result)
	}
	if diagnostics := diagnosticsToModel(server.validatePackIfDeclared(root)); len(diagnostics) > 0 {
		result.Diagnostics = diagnostics
		return jsonrpc.Success(request.ID, result)
	}
	data, err := os.ReadFile(filepath.Join(root, "artifacts", "catalog.json"))
	if err != nil {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{Code: "ARTIFACT_CATALOG_READ_FAILED", Severity: model.SeverityError, Stage: "artifact-rpc", Message: err.Error()})
		return jsonrpc.Success(request.ID, result)
	}
	var catalog artifacts.Catalog
	if err := json.Unmarshal(data, &catalog); err != nil || catalog.Artifacts == nil {
		result.Diagnostics = append(result.Diagnostics, model.Diagnostic{Code: "ARTIFACT_CATALOG_INVALID", Severity: model.SeverityError, Stage: "artifact-rpc", Message: "artifact catalog is invalid"})
		return jsonrpc.Success(request.ID, result)
	}
	query := strings.ToLower(strings.TrimSpace(params.Query))
	for _, artifact := range catalog.Artifacts {
		if params.Kind != "" && string(artifact.Kind) != params.Kind {
			continue
		}
		if query != "" && !strings.Contains(strings.ToLower(artifact.Name), query) && !strings.Contains(strings.ToLower(artifact.Path.Raw), query) {
			continue
		}
		result.Artifacts = append(result.Artifacts, artifact)
	}
	sort.Slice(result.Artifacts, func(i, j int) bool { return result.Artifacts[i].ID < result.Artifacts[j].ID })
	return jsonrpc.Success(request.ID, result)
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

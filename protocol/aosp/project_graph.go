package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/artifacts"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/projects"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type ProjectGraphParams struct {
	MetadataPackRoot string             `json:"metadataPackRoot"`
	MetadataRoot     string             `json:"metadataRoot,omitempty"`
	Targets          []model.TargetSpec `json:"targets"`
}

type ProjectGraphResult struct {
	Graph       projects.ProjectGraph `json:"graph"`
	Diagnostics []model.Diagnostic    `json:"diagnostics"`
}

type ArtifactsCatalogParams struct {
	MetadataPackRoot string `json:"metadataPackRoot"`
	MetadataRoot     string `json:"metadataRoot,omitempty"`
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
	result := ArtifactsCatalogResult{Artifacts: []artifacts.Artifact{}, Diagnostics: []model.Diagnostic{}}
	var catalog artifacts.Catalog
	if strings.TrimSpace(params.MetadataPackRoot) != "" {
		root := strings.TrimSpace(params.MetadataPackRoot)
		if diagnostics := diagnosticsToModel(server.validatePackIfDeclared(root)); len(diagnostics) > 0 {
			result.Diagnostics = diagnostics
			return jsonrpc.Success(request.ID, result)
		}
		data, err := os.ReadFile(filepath.Join(root, "artifacts", "catalog.json"))
		if err != nil {
			result.Diagnostics = []model.Diagnostic{buildFactsDiagnostic("ARTIFACT_CATALOG_READ_FAILED", err.Error(), filepath.Join(root, "artifacts", "catalog.json"))}
			return jsonrpc.Success(request.ID, result)
		}
		if err := json.Unmarshal(data, &catalog); err != nil || catalog.Artifacts == nil {
			result.Diagnostics = []model.Diagnostic{buildFactsDiagnostic("ARTIFACT_CATALOG_INVALID", "artifact catalog is invalid", filepath.Join(root, "artifacts", "catalog.json"))}
			return jsonrpc.Success(request.ID, result)
		}
	} else {
		_, loaded, diagnostics := server.loadBuildFactCatalogs("", params.MetadataRoot)
		if len(diagnostics) > 0 {
			result.Diagnostics = diagnostics
			return jsonrpc.Success(request.ID, result)
		}
		catalog = loaded
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
	if len(params.Targets) == 0 {
		return jsonrpc.Success(request.ID, ProjectGraphResult{Diagnostics: []model.Diagnostic{{Code: "PROJECT_GRAPH_TARGETS_MISSING", Severity: model.SeverityError, Stage: "project-graph-rpc", Message: "at least one target is required"}}})
	}
	moduleCatalog, artifactCatalog, diagnostics := server.loadBuildFactCatalogs(params.MetadataPackRoot, params.MetadataRoot)
	if len(diagnostics) > 0 {
		return jsonrpc.Success(request.ID, ProjectGraphResult{Diagnostics: diagnostics})
	}
	graph := projects.Build(moduleCatalog, artifactCatalog, params.Targets)
	return jsonrpc.Success(request.ID, ProjectGraphResult{Graph: graph, Diagnostics: graph.Diagnostics})
}

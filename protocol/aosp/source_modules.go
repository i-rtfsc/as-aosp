package aosp

import (
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/sourcecatalog"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type SourceModulesParams struct {
	WorkspaceViewRoot string `json:"workspaceViewRoot"`
	CacheRoot         string `json:"cacheRoot,omitempty"`
	Query             string `json:"query,omitempty"`
}

type SourceModuleResult struct {
	Modules      []sourcecatalog.Module `json:"modules"`
	Diagnostics  []model.Diagnostic     `json:"diagnostics"`
	FilesScanned int                    `json:"filesScanned"`
	CacheHit     bool                   `json:"cacheHit"`
}

func (server *Server) sourceModules(request jsonrpc.Request) jsonrpc.Response {
	var params SourceModulesParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	root := strings.TrimSpace(params.WorkspaceViewRoot)
	if root == "" {
		return jsonrpc.Success(request.ID, SourceModuleResult{
			Modules: []sourcecatalog.Module{},
			Diagnostics: []model.Diagnostic{{
				Code: "SOURCE_ROOT_MISSING", Severity: model.SeverityError,
				Stage: "source-catalog", Message: "workspaceViewRoot is required",
			}},
		})
	}
	report := sourcecatalog.Discover(root, params.CacheRoot)
	if query := strings.ToLower(strings.TrimSpace(params.Query)); query != "" {
		filtered := make([]sourcecatalog.Module, 0)
		for _, module := range report.Catalog.Modules {
			if strings.Contains(strings.ToLower(module.Name), query) ||
				strings.Contains(strings.ToLower(string(module.Blueprint)), query) ||
				strings.Contains(strings.ToLower(string(module.Directory)), query) {
				filtered = append(filtered, module)
			}
		}
		report.Catalog.Modules = filtered
	}
	return jsonrpc.Success(request.ID, SourceModuleResult{
		Modules: report.Catalog.Modules, Diagnostics: report.Diagnostics,
		FilesScanned: report.FilesScanned, CacheHit: report.CacheHit,
	})
}

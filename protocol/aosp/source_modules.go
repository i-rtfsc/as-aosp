package aosp

import (
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/sourcecatalog"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type SourceModulesParams struct {
	WorkspaceViewRoot string `json:"workspaceViewRoot"`
	OutRoot           string `json:"outRoot,omitempty"`
	CacheRoot         string `json:"cacheRoot,omitempty"`
	Query             string `json:"query,omitempty"`
	Limit             int    `json:"limit,omitempty"`
}

type SourceModuleResult struct {
	Modules      []sourcecatalog.Module `json:"modules"`
	Diagnostics  []model.Diagnostic     `json:"diagnostics"`
	FilesScanned int                    `json:"filesScanned"`
	CacheHit     bool                   `json:"cacheHit"`
	BuildBacked  int                    `json:"buildBacked"`
	Total        int                    `json:"total"`
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
	diagnostics := append([]model.Diagnostic(nil), report.Diagnostics...)
	if outRoot := strings.TrimSpace(params.OutRoot); outRoot != "" {
		buildCatalog, buildDiagnostics := server.loadBuildFactModuleCatalogCached(
			outRoot, params.CacheRoot,
		)
		diagnostics = append(diagnostics, buildDiagnostics...)
		if len(buildCatalog.Modules) > 0 {
			report.Catalog.Modules = sourcecatalog.MergeBuildFacts(
				report.Catalog.Modules, buildCatalog,
			)
		}
	}
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
	sort.SliceStable(report.Catalog.Modules, func(i, j int) bool {
		if report.Catalog.Modules[i].Recommended != report.Catalog.Modules[j].Recommended {
			return report.Catalog.Modules[i].Recommended
		}
		if report.Catalog.Modules[i].Name != report.Catalog.Modules[j].Name {
			return report.Catalog.Modules[i].Name < report.Catalog.Modules[j].Name
		}
		return report.Catalog.Modules[i].Blueprint < report.Catalog.Modules[j].Blueprint
	})
	total := len(report.Catalog.Modules)
	buildBacked := 0
	for _, module := range report.Catalog.Modules {
		if module.Status == sourcecatalog.StatusBuildBacked {
			buildBacked++
		}
	}
	limit := params.Limit
	if limit <= 0 || limit > 1000 {
		limit = 500
	}
	if len(report.Catalog.Modules) > limit {
		report.Catalog.Modules = report.Catalog.Modules[:limit]
	}
	return jsonrpc.Success(request.ID, SourceModuleResult{
		Modules: report.Catalog.Modules, Diagnostics: diagnostics,
		FilesScanned: report.FilesScanned, CacheHit: report.CacheHit,
		BuildBacked: buildBacked, Total: total,
	})
}

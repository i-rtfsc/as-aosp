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
	Modules      []SourceModuleSummary `json:"modules"`
	Diagnostics  []model.Diagnostic    `json:"diagnostics"`
	FilesScanned int                   `json:"filesScanned"`
	CacheHit     bool                  `json:"cacheHit"`
	BuildBacked  int                   `json:"buildBacked"`
	Total        int                   `json:"total"`
}

type SourceModuleSummary struct {
	Name           string                     `json:"name"`
	Type           string                     `json:"type"`
	Blueprint      model.RepoPath             `json:"blueprint,omitempty"`
	Directory      model.RepoPath             `json:"directory,omitempty"`
	Status         sourcecatalog.ModuleStatus `json:"status"`
	Capabilities   []string                   `json:"capabilities,omitempty"`
	Recommended    bool                       `json:"recommended"`
	BuildRole      sourcecatalog.BuildRole    `json:"buildRole"`
	OutputKinds    []string                   `json:"outputKinds,omitempty"`
	DefaultVisible bool                       `json:"defaultVisible"`
}

func (server *Server) sourceModules(request jsonrpc.Request) jsonrpc.Response {
	var params SourceModulesParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	root := strings.TrimSpace(params.WorkspaceViewRoot)
	if root == "" {
		return jsonrpc.Success(request.ID, SourceModuleResult{
			Modules: []SourceModuleSummary{},
			Diagnostics: []model.Diagnostic{{
				Code: "SOURCE_ROOT_MISSING", Severity: model.SeverityError,
				Stage: "source-catalog", Message: "workspaceViewRoot is required",
			}},
		})
	}
	report := sourcecatalog.Discover(root, params.CacheRoot)
	diagnostics := make([]model.Diagnostic, 0, len(report.Diagnostics))
	diagnostics = append(diagnostics, report.Diagnostics...)
	if outRoot := strings.TrimSpace(params.OutRoot); outRoot != "" {
		buildCatalog, buildDiagnostics := server.loadBuildFactModuleCatalogCached(
			outRoot, params.CacheRoot,
		)
		diagnostics = append(diagnostics, buildDiagnostics...)
		if len(buildCatalog.Modules) > 0 {
			report.Catalog.Modules = sourcecatalog.MergeBuildFacts(
				report.Catalog.Modules, buildCatalog,
			)
			sourceBacked := make([]sourcecatalog.Module, 0, len(report.Catalog.Modules))
			for _, module := range report.Catalog.Modules {
				if module.Status != sourcecatalog.StatusBuildOnly {
					sourceBacked = append(sourceBacked, module)
				}
			}
			report.Catalog.Modules = sourceBacked
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
	modules := report.Catalog.Modules
	if len(modules) > limit {
		modules = modules[:limit]
	}
	summaries := make([]SourceModuleSummary, 0, len(modules))
	for _, module := range modules {
		summaries = append(summaries, SourceModuleSummary{
			Name: module.Name, Type: module.Type, Blueprint: module.Blueprint,
			Directory: module.Directory, Status: module.Status,
			Capabilities: module.Capabilities, Recommended: module.Recommended,
			BuildRole: module.BuildRole, OutputKinds: module.OutputKinds,
			DefaultVisible: module.DefaultVisible,
		})
	}
	return jsonrpc.Success(request.ID, SourceModuleResult{
		Modules: summaries, Diagnostics: diagnostics,
		FilesScanned: report.FilesScanned, CacheHit: report.CacheHit,
		BuildBacked: buildBacked, Total: total,
	})
}

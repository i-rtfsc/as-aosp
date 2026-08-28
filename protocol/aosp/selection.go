package aosp

import (
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/selection"
	"github.com/i-rtfsc/as-aosp/core/sourcecatalog"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type SelectionCatalogResult struct {
	Presets          []selection.Preset `json:"presets"`
	DefaultPresetIDs []string           `json:"defaultPresetIds"`
}

type SelectionPlanParams struct {
	WorkspaceViewRoot string   `json:"workspaceViewRoot"`
	OutRoot           string   `json:"outRoot,omitempty"`
	CacheRoot         string   `json:"cacheRoot,omitempty"`
	PresetIDs         []string `json:"presetIds,omitempty"`
	ManualModules     []string `json:"manualModules,omitempty"`
}

func (server *Server) selectionCatalog(request jsonrpc.Request) jsonrpc.Response {
	return jsonrpc.Success(request.ID, SelectionCatalogResult{
		Presets: selection.Catalog(), DefaultPresetIDs: selection.DefaultPresetIDs(),
	})
}

func (server *Server) selectionPlan(request jsonrpc.Request) jsonrpc.Response {
	var params SelectionPlanParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	moduleKinds := map[string]model.TargetKind{}
	root := strings.TrimSpace(params.WorkspaceViewRoot)
	if root != "" && len(params.ManualModules) > 0 {
		report := sourcecatalog.Discover(root, params.CacheRoot)
		if outRoot := strings.TrimSpace(params.OutRoot); outRoot != "" {
			buildCatalog, _ := server.loadBuildFactModuleCatalogCached(
				outRoot, params.CacheRoot,
			)
			report.Catalog.Modules = sourcecatalog.MergeBuildFacts(
				report.Catalog.Modules, buildCatalog,
			)
		}
		for _, module := range report.Catalog.Modules {
			moduleKinds[module.Name] = selectionModuleKind(module)
		}
	}
	plan := selection.Build(params.PresetIDs, params.ManualModules, moduleKinds)
	return jsonrpc.Success(request.ID, plan)
}

func selectionModuleKind(module sourcecatalog.Module) model.TargetKind {
	typeName := strings.ToLower(module.Type)
	switch {
	case strings.Contains(typeName, "android_app") ||
		strings.Contains(typeName, "android_library"):
		return model.TargetAndroid
	case strings.Contains(typeName, "cc_") ||
		strings.Contains(typeName, "rust_") ||
		strings.Contains(typeName, "native"):
		return model.TargetNative
	default:
		return model.TargetJava
	}
}

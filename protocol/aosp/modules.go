package aosp

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
	corepack "github.com/i-rtfsc/as-aosp/core/pack"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type ModulesParams struct {
	MetadataPackRoot string   `json:"metadataPackRoot"`
	MetadataRoot     string   `json:"metadataRoot,omitempty"`
	Names            []string `json:"names,omitempty"`
	Name             string   `json:"name,omitempty"`
}

type ModuleSummary struct {
	Name            string `json:"name"`
	Types           int    `json:"types"`
	SourceRoots     int    `json:"sourceRoots"`
	Sources         int    `json:"sources"`
	Generated       int    `json:"generated"`
	Jars            int    `json:"jars"`
	SrcJars         int    `json:"srcjars"`
	Dependencies    int    `json:"dependencies"`
	AidlIncludeDirs int    `json:"aidlIncludeDirs"`
}

type ModulesCatalogResult struct {
	Modules     []ModuleSummary    `json:"modules"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

type ModuleResult struct {
	Module      *modules.Module    `json:"module,omitempty"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

type ModulesClosureResult struct {
	Modules     []modules.Module   `json:"modules"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

type moduleCatalogSnapshot struct {
	path    string
	size    int64
	modTime int64
	catalog modules.Catalog
}

func (server *Server) modulesCatalog(request jsonrpc.Request) jsonrpc.Response {
	var params ModulesParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	catalog, diagnostics := server.loadBuildFactModuleCatalog(params.MetadataPackRoot, params.MetadataRoot)
	result := ModulesCatalogResult{Modules: []ModuleSummary{}, Diagnostics: diagnostics}
	if len(diagnostics) != 0 {
		return jsonrpc.Success(request.ID, result)
	}
	names := sortedModuleNames(catalog.Modules)
	for _, name := range names {
		module := catalog.Modules[name]
		result.Modules = append(result.Modules, ModuleSummary{
			Name: name, Types: len(module.Types), SourceRoots: len(module.SourceRoots),
			Sources: len(module.Sources), Generated: len(module.Generated),
			Jars: len(module.Jars), SrcJars: len(module.SrcJars),
			Dependencies:    len(module.Dependencies),
			AidlIncludeDirs: len(module.AidlIncludeDirs),
		})
	}
	return jsonrpc.Success(request.ID, result)
}

func (server *Server) moduleGet(request jsonrpc.Request) jsonrpc.Response {
	var params ModulesParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	catalog, diagnostics := server.loadBuildFactModuleCatalog(params.MetadataPackRoot, params.MetadataRoot)
	result := ModuleResult{Diagnostics: diagnostics}
	if len(diagnostics) != 0 {
		return jsonrpc.Success(request.ID, result)
	}
	name := strings.TrimSpace(params.Name)
	module, ok := catalog.Modules[name]
	if !ok || name == "" {
		result.Diagnostics = []model.Diagnostic{moduleDiagnostic(
			"MODULE_NOT_FOUND", "module is not present in the metadata pack: "+name,
			params.MetadataPackRoot,
		)}
		return jsonrpc.Success(request.ID, result)
	}
	result.Module = &module
	return jsonrpc.Success(request.ID, result)
}

func (server *Server) modulesClosure(request jsonrpc.Request) jsonrpc.Response {
	var params ModulesParams
	if err := decodeParams(request, &params); err != nil {
		return jsonrpc.Failure(request.ID, jsonrpc.InvalidParams(err.Error()))
	}
	catalog, diagnostics := server.loadBuildFactModuleCatalog(params.MetadataPackRoot, params.MetadataRoot)
	result := ModulesClosureResult{Modules: []modules.Module{}, Diagnostics: diagnostics}
	if len(diagnostics) != 0 {
		return jsonrpc.Success(request.ID, result)
	}
	requested := uniqueModuleNames(params.Names)
	for _, name := range requested {
		if _, ok := catalog.Modules[name]; !ok {
			result.Diagnostics = append(result.Diagnostics, moduleDiagnostic(
				"MODULE_NOT_FOUND", "module is not present in the metadata pack: "+name,
				params.MetadataPackRoot,
			))
		}
	}
	for _, name := range catalog.Closure(requested) {
		if module, ok := catalog.Modules[name]; ok {
			result.Modules = append(result.Modules, module)
		}
	}
	return jsonrpc.Success(request.ID, result)
}

func (server *Server) loadModuleCatalog(root string) (modules.Catalog, []model.Diagnostic) {
	root = strings.TrimSpace(root)
	if root == "" {
		return modules.Catalog{}, []model.Diagnostic{moduleDiagnostic(
			"METADATA_PACK_ROOT_MISSING", "metadataPackRoot is required", "",
		)}
	}
	if diagnostics := server.validatePackIfDeclared(root); len(diagnostics) > 0 {
		return modules.Catalog{}, diagnosticsToModel(diagnostics)
	}
	pathName := filepath.Join(root, "modules", "catalog.json")
	info, err := os.Stat(pathName)
	if err != nil {
		return modules.Catalog{}, []model.Diagnostic{moduleDiagnostic(
			"MODULE_CATALOG_READ_FAILED", fmt.Sprintf("read module catalog: %v", err),
			pathName,
		)}
	}
	if server.moduleCatalogSnapshot != nil && server.moduleCatalogSnapshot.path == pathName &&
		server.moduleCatalogSnapshot.size == info.Size() &&
		server.moduleCatalogSnapshot.modTime == info.ModTime().UnixNano() {
		return server.moduleCatalogSnapshot.catalog, nil
	}
	catalog, err := modules.LoadPackCatalog(root)
	if err != nil {
		return modules.Catalog{}, []model.Diagnostic{moduleDiagnostic(
			"MODULE_CATALOG_INVALID", err.Error(), pathName,
		)}
	}
	server.moduleCatalogSnapshot = &moduleCatalogSnapshot{
		path: pathName, size: info.Size(), modTime: info.ModTime().UnixNano(), catalog: catalog,
	}
	return catalog, nil
}

func diagnosticsToModel(values []corepack.Diagnostic) []model.Diagnostic {
	result := make([]model.Diagnostic, 0, len(values))
	for _, value := range values {
		result = append(result, model.Diagnostic{
			Code: value.Code, Severity: model.SeverityError, Stage: "pack-validation",
			Message: value.Message, LocalPath: value.Path,
		})
	}
	return result
}

func moduleDiagnostic(code, message, pathName string) model.Diagnostic {
	return model.Diagnostic{
		Code: code, Severity: model.SeverityWarning, Stage: "modules-rpc",
		Message: message, LocalPath: pathName,
	}
}

func sortedModuleNames(values map[string]modules.Module) []string {
	names := make([]string, 0, len(values))
	for name := range values {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func uniqueModuleNames(values []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	sort.Strings(result)
	return result
}

package aosp

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/artifacts"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

func (server *Server) loadBuildFactCatalogs(outRoot string) (modules.Catalog, artifacts.Catalog, []model.Diagnostic) {
	moduleCatalog, diagnostics := server.loadBuildFactModuleCatalog(outRoot)
	if len(diagnostics) > 0 {
		return modules.Catalog{}, artifacts.Catalog{}, diagnostics
	}
	return moduleCatalog, artifacts.Build(moduleCatalog), nil
}

func (server *Server) loadBuildFactModuleCatalog(outRoot string) (modules.Catalog, []model.Diagnostic) {
	return server.loadBuildFactModuleCatalogCached(outRoot, "")
}

func (server *Server) loadBuildFactModuleCatalogCached(outRoot, cacheRoot string) (modules.Catalog, []model.Diagnostic) {
	outRoot = strings.TrimSpace(outRoot)
	if outRoot == "" {
		return modules.Catalog{}, []model.Diagnostic{buildFactsDiagnostic(
			"BUILD_FACTS_ROOT_MISSING", "outRoot is required", "",
		)}
	}
	absolute, err := filepath.Abs(outRoot)
	if err != nil {
		return modules.Catalog{}, []model.Diagnostic{buildFactsDiagnostic(
			"BUILD_FACTS_ROOT_INVALID", err.Error(), outRoot,
		)}
	}
	signature := moduleFactsSignature(absolute)
	server.moduleCatalogMu.Lock()
	defer server.moduleCatalogMu.Unlock()
	if snapshot := server.moduleCatalogSnapshot; snapshot != nil &&
		snapshot.root == absolute && snapshot.signature == signature {
		return snapshot.catalog, append([]model.Diagnostic(nil), snapshot.diagnostics...)
	}
	if cached, ok := loadBuildFactDiskCache(cacheRoot, absolute, signature); ok {
		server.moduleCatalogSnapshot = &moduleCatalogSnapshot{
			root: absolute, signature: signature, catalog: cached,
			diagnostics: []model.Diagnostic{},
		}
		return cached, nil
	}
	report := metadata.LoadModuleDirectory(absolute)
	diagnostics := make([]model.Diagnostic, 0)
	for _, diagnostic := range report.Diagnostics {
		if diagnostic.Code == "METADATA_READ_FAILED" || diagnostic.Code == "METADATA_PARSE_FAILED" {
			diagnostics = append(diagnostics, diagnostic)
		}
	}
	moduleCatalog := modules.Build(report.Snapshot)
	if len(moduleCatalog.Modules) == 0 {
		diagnostics = append(diagnostics, buildFactsDiagnostic(
			"BUILD_FACTS_EMPTY", "no module facts were found under outRoot", absolute,
		))
	}
	server.moduleCatalogSnapshot = &moduleCatalogSnapshot{
		root: absolute, signature: signature, catalog: moduleCatalog,
		diagnostics: append([]model.Diagnostic(nil), diagnostics...),
	}
	saveBuildFactDiskCache(cacheRoot, absolute, signature, moduleCatalog)
	return moduleCatalog, diagnostics
}

type buildFactCacheEnvelope struct {
	Schema    int             `json:"schema"`
	OutRoot   string          `json:"outRoot"`
	Signature string          `json:"signature"`
	Catalog   modules.Catalog `json:"catalog"`
}

func loadBuildFactDiskCache(cacheRoot, outRoot, signature string) (modules.Catalog, bool) {
	if strings.TrimSpace(cacheRoot) == "" {
		return modules.Catalog{}, false
	}
	data, err := os.ReadFile(filepath.Join(cacheRoot, "build-facts", "catalog.json"))
	if err != nil {
		return modules.Catalog{}, false
	}
	var envelope buildFactCacheEnvelope
	if json.Unmarshal(data, &envelope) != nil || envelope.Schema != 1 ||
		envelope.OutRoot != outRoot || envelope.Signature != signature ||
		len(envelope.Catalog.Modules) == 0 {
		return modules.Catalog{}, false
	}
	return envelope.Catalog, true
}

func saveBuildFactDiskCache(cacheRoot, outRoot, signature string, catalog modules.Catalog) {
	if strings.TrimSpace(cacheRoot) == "" || len(catalog.Modules) == 0 {
		return
	}
	directory := filepath.Join(cacheRoot, "build-facts")
	if os.MkdirAll(directory, 0o755) != nil {
		return
	}
	data, err := json.Marshal(buildFactCacheEnvelope{
		Schema: 1, OutRoot: outRoot, Signature: signature, Catalog: catalog,
	})
	if err != nil {
		return
	}
	temporary := filepath.Join(directory, "catalog.json.tmp")
	if os.WriteFile(temporary, append(data, '\n'), 0o644) == nil {
		_ = os.Rename(temporary, filepath.Join(directory, "catalog.json"))
	}
}

func moduleFactsSignature(root string) string {
	candidates := []string{
		filepath.Join(root, "module-info.json"),
		filepath.Join(root, "soong", "module-info.json"),
		filepath.Join(root, "module_bp_java_deps.json"),
		filepath.Join(root, "soong", "module_bp_java_deps.json"),
		filepath.Join(root, "module_bp_cc_deps.json"),
		filepath.Join(root, "soong", "module_bp_cc_deps.json"),
	}
	for _, pattern := range []string{
		filepath.Join(root, "soong", "module-info-*.json"),
		filepath.Join(root, "target", "product", "*", "module-info.json"),
	} {
		matches, _ := filepath.Glob(pattern)
		candidates = append(candidates, matches...)
	}
	parts := make([]string, 0, len(candidates))
	for _, pathName := range candidates {
		if info, err := os.Stat(pathName); err == nil && !info.IsDir() {
			parts = append(parts, fmt.Sprintf(
				"%s:%d:%d", filepath.Clean(pathName), info.Size(),
				info.ModTime().UnixNano(),
			))
		}
	}
	sort.Strings(parts)
	return strings.Join(parts, "|")
}

func buildFactsDiagnostic(code, message, pathName string) model.Diagnostic {
	return model.Diagnostic{
		Code: code, Severity: model.SeverityError, Stage: "build-facts-rpc",
		Message: message, LocalPath: pathName,
	}
}

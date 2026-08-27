package aosp

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/artifacts"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

func (server *Server) loadBuildFactCatalogs(packRoot, metadataRoot string) (modules.Catalog, artifacts.Catalog, []model.Diagnostic) {
	packRoot = strings.TrimSpace(packRoot)
	metadataRoot = strings.TrimSpace(metadataRoot)
	if packRoot != "" {
		if diagnostics := diagnosticsToModel(server.validatePackIfDeclared(packRoot)); len(diagnostics) > 0 {
			return modules.Catalog{}, artifacts.Catalog{}, diagnostics
		}
		moduleCatalog, err := modules.LoadPackCatalog(packRoot)
		if err != nil {
			return modules.Catalog{}, artifacts.Catalog{}, []model.Diagnostic{buildFactsDiagnostic(
				"BUILD_FACTS_MODULE_CATALOG_INVALID", err.Error(), filepath.Join(packRoot, "modules", "catalog.json"),
			)}
		}
		data, err := os.ReadFile(filepath.Join(packRoot, "artifacts", "catalog.json"))
		if err != nil {
			return modules.Catalog{}, artifacts.Catalog{}, []model.Diagnostic{buildFactsDiagnostic(
				"BUILD_FACTS_ARTIFACT_CATALOG_READ_FAILED", err.Error(), filepath.Join(packRoot, "artifacts", "catalog.json"),
			)}
		}
		var artifactCatalog artifacts.Catalog
		if err := json.Unmarshal(data, &artifactCatalog); err != nil || artifactCatalog.Artifacts == nil {
			return modules.Catalog{}, artifacts.Catalog{}, []model.Diagnostic{buildFactsDiagnostic(
				"BUILD_FACTS_ARTIFACT_CATALOG_INVALID", "artifact catalog is invalid", filepath.Join(packRoot, "artifacts", "catalog.json"),
			)}
		}
		return moduleCatalog, artifactCatalog, nil
	}
	moduleCatalog, diagnostics := server.loadBuildFactModuleCatalog("", metadataRoot)
	if len(diagnostics) > 0 {
		return modules.Catalog{}, artifacts.Catalog{}, diagnostics
	}
	return moduleCatalog, artifacts.Build(moduleCatalog), nil
}

func (server *Server) loadBuildFactModuleCatalog(packRoot, metadataRoot string) (modules.Catalog, []model.Diagnostic) {
	packRoot = strings.TrimSpace(packRoot)
	metadataRoot = strings.TrimSpace(metadataRoot)
	if packRoot != "" {
		return server.loadModuleCatalog(packRoot)
	}
	if metadataRoot == "" {
		return modules.Catalog{}, []model.Diagnostic{buildFactsDiagnostic(
			"BUILD_FACTS_ROOT_MISSING", "metadataPackRoot or metadataRoot is required", "",
		)}
	}
	report := metadata.LoadDirectory(metadataRoot)
	diagnostics := make([]model.Diagnostic, 0)
	for _, diagnostic := range report.Diagnostics {
		if diagnostic.Code == "METADATA_READ_FAILED" || diagnostic.Code == "METADATA_PARSE_FAILED" {
			diagnostics = append(diagnostics, diagnostic)
		}
	}
	moduleCatalog := modules.Build(report.Snapshot)
	if len(moduleCatalog.Modules) == 0 {
		diagnostics = append(diagnostics, buildFactsDiagnostic(
			"BUILD_FACTS_EMPTY", "no module facts were found under metadataRoot", metadataRoot,
		))
	}
	return moduleCatalog, diagnostics
}

func buildFactsDiagnostic(code, message, pathName string) model.Diagnostic {
	return model.Diagnostic{
		Code: code, Severity: model.SeverityError, Stage: "build-facts-rpc",
		Message: message, LocalPath: pathName,
	}
}

package aosp

import (
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
	outRoot = strings.TrimSpace(outRoot)
	if outRoot == "" {
		return modules.Catalog{}, []model.Diagnostic{buildFactsDiagnostic(
			"BUILD_FACTS_ROOT_MISSING", "outRoot is required", "",
		)}
	}
	report := metadata.LoadDirectory(outRoot)
	diagnostics := make([]model.Diagnostic, 0)
	for _, diagnostic := range report.Diagnostics {
		if diagnostic.Code == "METADATA_READ_FAILED" || diagnostic.Code == "METADATA_PARSE_FAILED" {
			diagnostics = append(diagnostics, diagnostic)
		}
	}
	moduleCatalog := modules.Build(report.Snapshot)
	if len(moduleCatalog.Modules) == 0 {
		diagnostics = append(diagnostics, buildFactsDiagnostic(
			"BUILD_FACTS_EMPTY", "no module facts were found under outRoot", outRoot,
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

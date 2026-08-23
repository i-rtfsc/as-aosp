package model

import "fmt"

func (p Profile) Validate() []Diagnostic {
	diagnostics := make([]Diagnostic, 0)
	if p.SchemaVersion == 0 {
		diagnostics = append(diagnostics, Diagnostic{Code: "PROFILE_SCHEMA_VERSION_MISSING", Severity: SeverityError, Stage: "profile", Message: "schemaVersion is required"})
	}
	if p.ID == "" {
		diagnostics = append(diagnostics, Diagnostic{Code: "PROFILE_ID_EMPTY", Severity: SeverityError, Stage: "profile", Message: "profile id is required"})
	}
	if p.WorkspaceViewRoot == "" {
		diagnostics = append(diagnostics, Diagnostic{Code: "PROFILE_VIEW_ROOT_EMPTY", Severity: SeverityError, Stage: "profile", Message: "workspaceViewRoot is required"})
	}
	if p.VisibilityPolicy.Mode != "" && p.VisibilityPolicy.Mode != "strict-view" && p.VisibilityPolicy.Mode != "open-external-on-demand" && p.VisibilityPolicy.Mode != "auto-expand" {
		diagnostics = append(diagnostics, Diagnostic{Code: "PROFILE_VISIBILITY_MODE_UNKNOWN", Severity: SeverityError, Stage: "profile", Message: fmt.Sprintf("unknown visibility mode %q", p.VisibilityPolicy.Mode)})
	}
	return diagnostics
}

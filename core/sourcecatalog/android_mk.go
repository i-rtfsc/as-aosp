package sourcecatalog

import "github.com/i-rtfsc/as-aosp/core/model"

// Android.mk is intentionally discovered as a source marker in the first slice.
// Soong facts remain authoritative for its build metadata until a dedicated
// Make variable parser is added.
func parseMakefile(file string) []Module {
	return []Module{{
		Name: "Android.mk:" + string(model.RepoPath(file)), Type: "make-module",
		Blueprint: model.RepoPath(file), Directory: moduleDirectory(file),
		Status: StatusSourceOnly, Capabilities: []string{"source"},
	}}
}

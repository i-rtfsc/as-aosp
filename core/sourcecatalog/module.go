package sourcecatalog

import "github.com/i-rtfsc/as-aosp/core/model"

type ModuleStatus string

const (
	StatusSourceOnly  ModuleStatus = "source-only"
	StatusBuildBacked ModuleStatus = "build-backed"
	StatusBuildOnly   ModuleStatus = "build-only"
	StatusPartial     ModuleStatus = "partial"
	StatusError       ModuleStatus = "error"
)

type BuildRole string

const (
	RoleProduct      BuildRole = "product"
	RoleLibrary      BuildRole = "library"
	RoleGenerator    BuildRole = "generator"
	RoleBuildSupport BuildRole = "build-support"
	RoleTest         BuildRole = "test"
)

type Module struct {
	Name           string           `json:"name"`
	Type           string           `json:"type"`
	Blueprint      model.RepoPath   `json:"blueprint"`
	Directory      model.RepoPath   `json:"directory"`
	SourceRoots    []model.RepoPath `json:"sourceRoots,omitempty"`
	Sources        []model.RepoPath `json:"sources,omitempty"`
	Dependencies   []string         `json:"dependencies,omitempty"`
	Status         ModuleStatus     `json:"status"`
	Capabilities   []string         `json:"capabilities,omitempty"`
	Recommended    bool             `json:"recommended"`
	BuildRole      BuildRole        `json:"buildRole"`
	OutputKinds    []string         `json:"outputKinds,omitempty"`
	DefaultVisible bool             `json:"defaultVisible"`
}

type Catalog struct {
	Modules      []Module          `json:"modules"`
	Fingerprints map[string]string `json:"fingerprints,omitempty"`
}

type Report struct {
	Catalog      Catalog            `json:"catalog"`
	Diagnostics  []model.Diagnostic `json:"diagnostics"`
	FilesScanned int                `json:"filesScanned"`
	CacheHit     bool               `json:"cacheHit"`
}

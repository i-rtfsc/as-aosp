package model

type Profile struct {
	SchemaVersion     int                  `json:"schemaVersion"`
	ID                string               `json:"id"`
	WorkspaceViewRoot string               `json:"workspaceViewRoot"`
	CanonicalRepoRoot string               `json:"canonicalRepoRoot,omitempty"`
	OutRoot           string               `json:"outRoot,omitempty"`
	MetadataPackRoot  string               `json:"metadataPackRoot,omitempty"`
	BuildTarget       BuildTarget          `json:"buildTarget,omitempty"`
	Scope             Scope                `json:"scope,omitempty"`
	SourceMounts      []ProfileSourceMount `json:"sourceMounts,omitempty"`
	VisibilityPolicy  VisibilityPolicy     `json:"visibilityPolicy,omitempty"`
	PathMappings      []PathMapping        `json:"pathMappings,omitempty"`
	BuildIntegration  BuildIntegration     `json:"buildIntegration,omitempty"`
}

type BuildTarget struct {
	Product string `json:"product,omitempty"`
	Release string `json:"release,omitempty"`
	Variant string `json:"variant,omitempty"`
}

type Scope struct {
	Mode    string   `json:"mode,omitempty"`
	Include []string `json:"include,omitempty"`
	Exclude []string `json:"exclude,omitempty"`
}

type VisibilityPolicy struct {
	Mode                   string `json:"mode,omitempty"`
	AllowOpenExternalFile  bool   `json:"allowOpenExternalFile,omitempty"`
	ConfirmBeforeExpanding bool   `json:"confirmBeforeExpanding,omitempty"`
}

type PathMapping struct {
	Name           string `json:"name,omitempty"`
	ExternalPrefix string `json:"externalPrefix"`
	LocalPrefix    string `json:"localPrefix"`
	Kind           string `json:"kind,omitempty"`
}

type BuildIntegration struct {
	Mode string `json:"mode,omitempty"`
}

type ProfileSourceMount struct {
	Name       string `json:"name"`
	RepoPrefix string `json:"repoPrefix"`
	ViewPath   string `json:"viewPath"`
	LocalPath  string `json:"localPath"`
	Origin     string `json:"origin"`
	Priority   int    `json:"priority"`
	ReadOnly   bool   `json:"readOnly"`
}

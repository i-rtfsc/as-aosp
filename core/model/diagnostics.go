package model

type Severity string

const (
	SeverityInfo    Severity = "info"
	SeverityWarning Severity = "warning"
	SeverityError   Severity = "error"
)

type Diagnostic struct {
	Code            string   `json:"code"`
	Severity        Severity `json:"severity"`
	Stage           string   `json:"stage"`
	Message         string   `json:"message"`
	SuggestedAction string   `json:"suggestedAction,omitempty"`
	RepoPath        RepoPath `json:"repoPath,omitempty"`
	ViewPath        ViewPath `json:"viewPath,omitempty"`
	LocalPath       string   `json:"localPath,omitempty"`
	ExternalPath    string   `json:"externalPath,omitempty"`
}

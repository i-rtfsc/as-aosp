package topology

import (
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

type MountOrigin string

const (
	OriginDirectory    MountOrigin = "directory"
	OriginSymlink      MountOrigin = "symlink"
	OriginJunction     MountOrigin = "junction"
	OriginBindMount    MountOrigin = "bind-mount"
	OriginNetworkMount MountOrigin = "network-mount"
	OriginCopiedTree   MountOrigin = "copied-tree"
	OriginOverlay      MountOrigin = "overlay"
	OriginCanonical    MountOrigin = "canonical-fallback"
)

type SourceMount struct {
	Name       string          `json:"name"`
	RepoPrefix model.RepoPath  `json:"repoPrefix"`
	ViewPath   model.ViewPath  `json:"viewPath"`
	LocalPath  model.LocalPath `json:"localPath"`
	Origin     MountOrigin     `json:"origin"`
	Priority   int             `json:"priority"`
	ReadOnly   bool            `json:"readOnly"`
}

type WorkspaceTopology struct {
	WorkspaceViewRoot model.LocalPath `json:"workspaceViewRoot"`
	CanonicalRepoRoot model.LocalPath `json:"canonicalRepoRoot,omitempty"`
	OutRoot           model.LocalPath `json:"outRoot,omitempty"`
	SourceMounts      []SourceMount   `json:"sourceMounts"`
}

type ProviderState string

const (
	ProviderActive   ProviderState = "active"
	ProviderShadowed ProviderState = "shadowed"
	ProviderConflict ProviderState = "conflict"
	ProviderMissing  ProviderState = "missing"
)

type ProviderResolution struct {
	RepoPath   model.RepoPath `json:"repoPath"`
	Mount      *SourceMount   `json:"mount,omitempty"`
	Candidates []SourceMount  `json:"candidates,omitempty"`
	State      ProviderState  `json:"state"`
}

func (t WorkspaceTopology) Validate() []model.Diagnostic {
	diagnostics := make([]model.Diagnostic, 0)
	for i := range t.SourceMounts {
		mount := t.SourceMounts[i]
		if mount.Name == "" {
			diagnostics = append(diagnostics, model.Diagnostic{Code: "TOPOLOGY_MOUNT_NAME_EMPTY", Severity: model.SeverityError, Stage: "topology", Message: "source mount name is empty"})
		}
		if mount.RepoPrefix == "" {
			diagnostics = append(diagnostics, model.Diagnostic{Code: "TOPOLOGY_REPO_PREFIX_EMPTY", Severity: model.SeverityError, Stage: "topology", Message: "source mount repoPrefix is empty", ViewPath: mount.ViewPath})
		}
		if mount.ViewPath == "" {
			diagnostics = append(diagnostics, model.Diagnostic{Code: "TOPOLOGY_VIEW_PATH_EMPTY", Severity: model.SeverityError, Stage: "topology", Message: "source mount viewPath is empty", RepoPath: mount.RepoPrefix})
		}
		if mount.LocalPath == "" {
			diagnostics = append(diagnostics, model.Diagnostic{Code: "TOPOLOGY_LOCAL_PATH_EMPTY", Severity: model.SeverityError, Stage: "topology", Message: "source mount localPath is empty", RepoPath: mount.RepoPrefix})
		}
	}
	for i := range t.SourceMounts {
		for j := i + 1; j < len(t.SourceMounts); j++ {
			a, b := t.SourceMounts[i], t.SourceMounts[j]
			if a.RepoPrefix == b.RepoPrefix && a.Priority == b.Priority && normalizeLocal(a.LocalPath) != normalizeLocal(b.LocalPath) {
				diagnostics = append(diagnostics, model.Diagnostic{Code: "TOPOLOGY_PROVIDER_CONFLICT", Severity: model.SeverityError, Stage: "topology", Message: "same repo prefix has different providers at the same priority", RepoPath: a.RepoPrefix})
			}
		}
	}
	return diagnostics
}

func (t WorkspaceTopology) Resolve(repoPath model.RepoPath) ProviderResolution {
	var candidates []SourceMount
	for _, mount := range t.SourceMounts {
		if repoPath == mount.RepoPrefix || strings.HasPrefix(string(repoPath), string(mount.RepoPrefix)+"/") {
			candidates = append(candidates, mount)
		}
	}
	sort.SliceStable(candidates, func(i, j int) bool {
		if len(candidates[i].RepoPrefix) != len(candidates[j].RepoPrefix) {
			return len(candidates[i].RepoPrefix) > len(candidates[j].RepoPrefix)
		}
		return candidates[i].Priority > candidates[j].Priority
	})
	if len(candidates) == 0 {
		return ProviderResolution{RepoPath: repoPath, State: ProviderMissing}
	}
	if len(candidates) > 1 && candidates[0].RepoPrefix == candidates[1].RepoPrefix && candidates[0].Priority == candidates[1].Priority && normalizeLocal(candidates[0].LocalPath) != normalizeLocal(candidates[1].LocalPath) {
		return ProviderResolution{RepoPath: repoPath, Candidates: candidates, State: ProviderConflict}
	}
	return ProviderResolution{RepoPath: repoPath, Mount: &candidates[0], Candidates: candidates, State: ProviderActive}
}

func normalizeLocal(value model.LocalPath) string {
	normalized := strings.ReplaceAll(string(value), "\\", "/")
	if len(normalized) >= 2 && normalized[1] == ':' || strings.HasPrefix(string(value), `\\`) {
		return strings.ToLower(normalized)
	}
	return normalized
}

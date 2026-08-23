package topology

import (
	"github.com/i-rtfsc/as-aosp/core/model"
)

func FromProfile(profile model.Profile) WorkspaceTopology {
	topology := WorkspaceTopology{
		WorkspaceViewRoot: model.LocalPath(profile.WorkspaceViewRoot),
		CanonicalRepoRoot: model.LocalPath(profile.CanonicalRepoRoot),
		OutRoot:           model.LocalPath(profile.OutRoot),
	}
	for _, source := range profile.SourceMounts {
		repoPrefix, _ := model.NormalizeRepoPath(source.RepoPrefix)
		viewPath, _ := model.NormalizeRelativePath(source.ViewPath)
		origin := MountOrigin(source.Origin)
		if origin == "" {
			origin = OriginDirectory
		}
		topology.SourceMounts = append(topology.SourceMounts, SourceMount{
			Name: source.Name, RepoPrefix: repoPrefix, ViewPath: model.ViewPath(viewPath), LocalPath: model.LocalPath(source.LocalPath), Origin: origin, Priority: source.Priority, ReadOnly: source.ReadOnly,
		})
	}
	return topology
}

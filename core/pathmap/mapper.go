package pathmap

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/topology"
)

type Mapper struct {
	Topology topology.WorkspaceTopology
	Mappings []model.PathMapping
}

func (m Mapper) RepoToLocal(repoPath model.RepoPath) (model.LocalPath, error) {
	resolution := m.Topology.Resolve(repoPath)
	if resolution.State != topology.ProviderActive || resolution.Mount == nil {
		return "", fmt.Errorf("repo path %q provider state is %s", repoPath, resolution.State)
	}
	suffix := strings.TrimPrefix(string(repoPath), string(resolution.Mount.RepoPrefix))
	suffix = strings.TrimPrefix(suffix, "/")
	if suffix == "" {
		return resolution.Mount.LocalPath, nil
	}
	return model.LocalPath(joinLocal(string(resolution.Mount.LocalPath), suffix)), nil
}

func (m Mapper) ViewToRepo(viewPath model.ViewPath) (model.RepoPath, error) {
	clean, err := model.NormalizeRelativePath(string(viewPath))
	if err != nil {
		return "", err
	}
	var selected *topology.SourceMount
	for i := range m.Topology.SourceMounts {
		mount := &m.Topology.SourceMounts[i]
		prefix := strings.Trim(string(mount.ViewPath), "/")
		if clean == prefix || strings.HasPrefix(clean, prefix+"/") {
			if selected == nil || len(prefix) > len(string(selected.ViewPath)) {
				selected = mount
			}
		}
	}
	if selected == nil {
		return "", fmt.Errorf("view path %q has no source mount", viewPath)
	}
	suffix := strings.TrimPrefix(clean, strings.Trim(string(selected.ViewPath), "/"))
	suffix = strings.TrimPrefix(suffix, "/")
	if suffix == "" {
		return selected.RepoPrefix, nil
	}
	return model.NormalizeRepoPath(string(selected.RepoPrefix) + "/" + suffix)
}

func (m Mapper) ExternalToLocal(external model.ExternalPath) (model.LocalPath, bool) {
	input := normalizeAbsolute(string(external))
	var selected *model.PathMapping
	for _, mapping := range m.Mappings {
		prefix := normalizeAbsolute(mapping.ExternalPrefix)
		if input == prefix || strings.HasPrefix(input, prefix+"/") {
			if selected == nil || len(prefix) > len(normalizeAbsolute(selected.ExternalPrefix)) {
				candidate := mapping
				selected = &candidate
			}
		}
	}
	if selected != nil {
		prefix := normalizeAbsolute(selected.ExternalPrefix)
		suffix := strings.TrimPrefix(strings.TrimPrefix(input, prefix), "/")
		return model.LocalPath(joinLocal(selected.LocalPrefix, suffix)), true
	}
	return "", false
}

func (m Mapper) LocalToRepo(local model.LocalPath) (model.RepoPath, error) {
	input := localKey(string(local))
	var selected *topology.SourceMount
	for i := range m.Topology.SourceMounts {
		mount := &m.Topology.SourceMounts[i]
		prefix := localKey(string(mount.LocalPath))
		if input == prefix || strings.HasPrefix(input, prefix+"/") {
			if selected == nil || len(prefix) > len(localKey(string(selected.LocalPath))) {
				selected = mount
			}
		}
	}
	if selected == nil {
		return "", fmt.Errorf("local path %q has no source mount", local)
	}
	suffix := strings.TrimPrefix(strings.TrimPrefix(input, localKey(string(selected.LocalPath))), "/")
	if suffix == "" {
		return selected.RepoPrefix, nil
	}
	return model.NormalizeRepoPath(string(selected.RepoPrefix) + "/" + suffix)
}

func normalizeAbsolute(value string) string {
	return strings.TrimRight(strings.ReplaceAll(value, "\\", "/"), "/")
}

func joinLocal(prefix, suffix string) string {
	if suffix == "" {
		return prefix
	}
	if isWindowsPath(prefix) {
		return strings.TrimRight(strings.ReplaceAll(prefix, "/", "\\"), "\\") + "\\" + strings.ReplaceAll(suffix, "/", "\\")
	}
	return filepath.Join(prefix, filepath.FromSlash(suffix))
}

func localKey(value string) string {
	key := normalizeAbsolute(value)
	if isWindowsPath(value) {
		return strings.ToLower(key)
	}
	return key
}

func isWindowsPath(value string) bool {
	return len(value) >= 2 && value[1] == ':' || strings.HasPrefix(value, `\\`)
}

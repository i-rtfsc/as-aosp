package topology

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

type DiscoverOptions struct {
	WorkspaceViewRoot string
	CanonicalRepoRoot string
	MaxDepth          int
}

func Discover(options DiscoverOptions) (WorkspaceTopology, []model.Diagnostic, error) {
	root, err := filepath.Abs(options.WorkspaceViewRoot)
	if err != nil {
		return WorkspaceTopology{}, nil, err
	}
	info, err := os.Stat(root)
	if err != nil {
		return WorkspaceTopology{}, nil, err
	}
	if !info.IsDir() {
		return WorkspaceTopology{}, nil, fmt.Errorf("workspace view root is not a directory: %s", root)
	}
	if options.MaxDepth <= 0 {
		options.MaxDepth = 1
	}
	if options.CanonicalRepoRoot != "" {
		canonical, canonicalErr := filepath.Abs(options.CanonicalRepoRoot)
		if canonicalErr == nil {
			if resolved, resolveErr := filepath.EvalSymlinks(canonical); resolveErr == nil {
				canonical = resolved
			}
			options.CanonicalRepoRoot = canonical
		}
	}
	topology := WorkspaceTopology{WorkspaceViewRoot: model.LocalPath(root), CanonicalRepoRoot: model.LocalPath(options.CanonicalRepoRoot)}
	diagnostics := make([]model.Diagnostic, 0)
	err = discoverDir(root, root, options, 0, map[string]bool{}, map[string]bool{}, &topology, &diagnostics)
	if err != nil {
		return topology, diagnostics, err
	}
	diagnostics = append(diagnostics, topology.Validate()...)
	return topology, diagnostics, nil
}

func discoverDir(root, current string, options DiscoverOptions, depth int, active, seen map[string]bool, topology *WorkspaceTopology, diagnostics *[]model.Diagnostic) error {
	identity := filepath.Clean(current)
	if active[identity] {
		*diagnostics = append(*diagnostics, model.Diagnostic{Code: "TOPOLOGY_LINK_CYCLE", Severity: model.SeverityError, Stage: "discover", Message: "directory traversal cycle detected", LocalPath: current})
		return nil
	}
	if seen[identity] {
		*diagnostics = append(*diagnostics, model.Diagnostic{Code: "TOPOLOGY_DUPLICATE_TARGET", Severity: model.SeverityWarning, Stage: "discover", Message: "directory target is exposed by multiple view paths", LocalPath: current})
		return nil
	}
	active[identity] = true
	seen[identity] = true
	defer delete(active, identity)
	entries, err := os.ReadDir(current)
	if err != nil {
		*diagnostics = append(*diagnostics, model.Diagnostic{Code: "TOPOLOGY_READ_DIR_FAILED", Severity: model.SeverityWarning, Stage: "discover", Message: err.Error(), LocalPath: current})
		return nil
	}
	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), ".") {
			continue
		}
		viewPath := filepath.Join(current, entry.Name())
		fileInfo, infoErr := os.Lstat(viewPath)
		if infoErr != nil {
			*diagnostics = append(*diagnostics, model.Diagnostic{Code: "TOPOLOGY_LSTAT_FAILED", Severity: model.SeverityWarning, Stage: "discover", Message: infoErr.Error(), LocalPath: viewPath})
			continue
		}
		if fileInfo.Mode()&os.ModeSymlink != 0 {
			target, targetErr := filepath.EvalSymlinks(viewPath)
			if targetErr != nil {
				*diagnostics = append(*diagnostics, model.Diagnostic{Code: "TOPOLOGY_DANGLING_LINK", Severity: model.SeverityError, Stage: "discover", Message: targetErr.Error(), LocalPath: viewPath})
				continue
			}
			repoPrefix := inferRepoPrefix(root, options.CanonicalRepoRoot, viewPath, target)
			if repoPrefix == "" {
				*diagnostics = append(*diagnostics, model.Diagnostic{Code: "TOPOLOGY_REPO_PREFIX_UNKNOWN", Severity: model.SeverityWarning, Stage: "discover", Message: "cannot infer RepoPath prefix; add an explicit source mount", LocalPath: viewPath})
				continue
			}
			mount := SourceMount{Name: entry.Name(), RepoPrefix: repoPrefix, ViewPath: model.ViewPath(relativeSlash(root, viewPath)), LocalPath: model.LocalPath(target), Origin: OriginSymlink, Priority: 100, ReadOnly: true}
			topology.SourceMounts = append(topology.SourceMounts, mount)
			if depth+1 < options.MaxDepth {
				_ = discoverDir(root, target, options, depth+1, active, seen, topology, diagnostics)
			}
			continue
		}
		if fileInfo.IsDir() && depth+1 < options.MaxDepth {
			_ = discoverDir(root, viewPath, options, depth+1, active, seen, topology, diagnostics)
		}
	}
	return nil
}

func inferRepoPrefix(root, canonical, viewPath, target string) model.RepoPath {
	if canonical != "" {
		if relative, err := filepath.Rel(canonical, target); err == nil && relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
			if normalized, err := model.NormalizeRepoPath(filepath.ToSlash(relative)); err == nil {
				return normalized
			}
		}
	}
	relative, err := filepath.Rel(root, viewPath)
	if err != nil {
		return ""
	}
	normalized, err := model.NormalizeRepoPath(filepath.ToSlash(relative))
	if err != nil {
		return ""
	}
	return normalized
}

func relativeSlash(root, value string) string {
	relative, err := filepath.Rel(root, value)
	if err != nil {
		return filepath.ToSlash(value)
	}
	return filepath.ToSlash(relative)
}

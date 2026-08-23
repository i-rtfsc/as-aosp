package model

import (
	"fmt"
	"path"
	"strings"
)

// RepoPath is a slash-separated path relative to the canonical AOSP root.
type RepoPath string

// ViewPath is a path relative to the workspace view root.
type ViewPath string

// LocalPath is a path understood by the current host filesystem.
type LocalPath string

// ExternalPath is an absolute path recorded by a build host or metadata file.
type ExternalPath string

func NormalizeRepoPath(value string) (RepoPath, error) {
	value = strings.TrimSpace(strings.ReplaceAll(value, "\\", "/"))
	if value == "" {
		return "", fmt.Errorf("repo path is empty")
	}
	if strings.HasPrefix(value, "/") || hasVolumePrefix(value) {
		return "", fmt.Errorf("repo path must be relative: %q", value)
	}
	clean := path.Clean(value)
	if clean == "." || clean == ".." || strings.HasPrefix(clean, "../") {
		return "", fmt.Errorf("repo path escapes root: %q", value)
	}
	return RepoPath(clean), nil
}

func NormalizeRelativePath(value string) (string, error) {
	value = strings.TrimSpace(strings.ReplaceAll(value, "\\", "/"))
	if value == "" {
		return ".", nil
	}
	if strings.HasPrefix(value, "/") || hasVolumePrefix(value) {
		return "", fmt.Errorf("relative path expected: %q", value)
	}
	clean := path.Clean(value)
	if clean == ".." || strings.HasPrefix(clean, "../") {
		return "", fmt.Errorf("relative path escapes root: %q", value)
	}
	return clean, nil
}

func hasVolumePrefix(value string) bool {
	return len(value) >= 2 && ((value[0] >= 'A' && value[0] <= 'Z') || (value[0] >= 'a' && value[0] <= 'z')) && value[1] == ':'
}

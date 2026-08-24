package aosp

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

type workspaceSourceScan struct {
	Files       map[model.RepoPath]string
	Diagnostics []model.Diagnostic
}

func scanWorkspaceSources(
	root string,
	stage string,
	invalidCode string,
	readCode string,
	accept func(extension string) bool,
) workspaceSourceScan {
	result := workspaceSourceScan{
		Files:       make(map[model.RepoPath]string),
		Diagnostics: make([]model.Diagnostic, 0),
	}
	absolute, err := filepath.Abs(root)
	if err != nil {
		result.Diagnostics = append(result.Diagnostics,
			workspaceDiagnostic(invalidCode, stage, err.Error()))
		return result
	}
	info, err := os.Stat(absolute)
	if err != nil || !info.IsDir() {
		message := "workspace view root is not a directory"
		if err != nil {
			message = err.Error()
		}
		result.Diagnostics = append(result.Diagnostics,
			workspaceDiagnostic(invalidCode, stage, message))
		return result
	}
	var walk func(string, string, map[string]bool) error
	walk = func(current, relative string, ancestors map[string]bool) error {
		real, evalErr := filepath.EvalSymlinks(current)
		if evalErr == nil {
			real = filepath.Clean(real)
			if ancestors[real] {
				return nil
			}
			ancestors = cloneStringSet(ancestors)
			ancestors[real] = true
		}
		entries, readErr := os.ReadDir(current)
		if readErr != nil {
			return readErr
		}
		for _, entry := range entries {
			childRelative := filepath.Join(relative, entry.Name())
			if isWorkspaceExcluded(childRelative) {
				continue
			}
			child := filepath.Join(current, entry.Name())
			childInfo, statErr := os.Stat(child)
			if statErr != nil {
				result.Diagnostics = append(result.Diagnostics,
					workspaceDiagnostic(readCode, stage, statErr.Error()))
				continue
			}
			if childInfo.IsDir() {
				if err := walk(child, childRelative, ancestors); err != nil {
					result.Diagnostics = append(result.Diagnostics,
						workspaceDiagnostic(readCode, stage, err.Error()))
				}
				continue
			}
			extension := strings.ToLower(filepath.Ext(entry.Name()))
			if !accept(extension) {
				continue
			}
			source, readErr := os.ReadFile(child)
			if readErr != nil {
				result.Diagnostics = append(result.Diagnostics,
					workspaceDiagnostic(readCode, stage, readErr.Error()))
				continue
			}
			repoPath, pathErr := model.NormalizeRepoPath(
				filepath.ToSlash(childRelative),
			)
			if pathErr != nil {
				result.Diagnostics = append(result.Diagnostics,
					workspaceDiagnostic(invalidCode, stage, pathErr.Error()))
				continue
			}
			result.Files[repoPath] = string(source)
		}
		return nil
	}
	if err := walk(absolute, "", map[string]bool{}); err != nil {
		result.Diagnostics = append(result.Diagnostics,
			workspaceDiagnostic(readCode, stage, err.Error()))
	}
	return result
}

func isWorkspaceExcluded(relative string) bool {
	for _, part := range strings.Split(filepath.ToSlash(relative), "/") {
		if part == ".git" || part == ".repo" ||
			part == "out" || part == ".aosp-ide" {
			return true
		}
	}
	return false
}

func cloneStringSet(values map[string]bool) map[string]bool {
	clone := make(map[string]bool, len(values)+1)
	for value := range values {
		clone[value] = true
	}
	return clone
}

func workspaceDiagnostic(code, stage, message string) model.Diagnostic {
	return model.Diagnostic{
		Code: code, Severity: model.SeverityWarning,
		Stage: stage, Message: message,
	}
}

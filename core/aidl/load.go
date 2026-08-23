package aidl

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

type LoadReport struct {
	Index       Index              `json:"index"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

func LoadTree(root string) (LoadReport, error) {
	absolute, err := filepath.Abs(root)
	if err != nil {
		return LoadReport{}, err
	}
	report := LoadReport{Index: NewIndex(), Diagnostics: make([]model.Diagnostic, 0)}
	var files []string
	err = filepath.WalkDir(absolute, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			report.Diagnostics = append(report.Diagnostics, model.Diagnostic{Code: "AIDL_READ_FAILED", Severity: model.SeverityWarning, Stage: "aidl-load", Message: walkErr.Error(), LocalPath: path})
			return nil
		}
		if entry.IsDir() && strings.HasPrefix(entry.Name(), ".") && path != absolute {
			return filepath.SkipDir
		}
		if !entry.IsDir() && strings.EqualFold(filepath.Ext(entry.Name()), ".aidl") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return report, err
	}
	sort.Strings(files)
	for _, path := range files {
		content, readErr := os.ReadFile(path)
		if readErr != nil {
			report.Diagnostics = append(report.Diagnostics, model.Diagnostic{Code: "AIDL_READ_FAILED", Severity: model.SeverityWarning, Stage: "aidl-load", Message: readErr.Error(), LocalPath: path})
			continue
		}
		relative, relErr := filepath.Rel(absolute, path)
		if relErr != nil {
			return report, fmt.Errorf("relative AIDL path: %w", relErr)
		}
		repoPath, normalizeErr := model.NormalizeRepoPath(filepath.ToSlash(relative))
		if normalizeErr != nil {
			report.Diagnostics = append(report.Diagnostics, model.Diagnostic{Code: "AIDL_PATH_INVALID", Severity: model.SeverityWarning, Stage: "aidl-load", Message: normalizeErr.Error(), LocalPath: path})
			continue
		}
		document := Parse(repoPath, string(content))
		for index := range document.Diagnostics {
			if document.Diagnostics[index].RepoPath == "" {
				document.Diagnostics[index].RepoPath = repoPath
			}
		}
		report.Diagnostics = append(report.Diagnostics, document.Diagnostics...)
		report.Index.Add(document)
	}
	for _, document := range report.Index.Documents {
		report.Diagnostics = append(report.Diagnostics, report.Index.Validate(document)...)
	}
	return report, nil
}

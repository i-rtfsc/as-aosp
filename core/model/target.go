package model

import (
	"fmt"
	"sort"
	"strings"
)

type TargetKind string

const (
	TargetJava    TargetKind = "java"
	TargetAndroid TargetKind = "android"
	TargetNative  TargetKind = "native"
	TargetMixed   TargetKind = "mixed"
)

type ArtifactKind string

const (
	ArtifactJar    ArtifactKind = "jar"
	ArtifactAPK    ArtifactKind = "apk"
	ArtifactSrcJar ArtifactKind = "srcjar"
	ArtifactAAR    ArtifactKind = "aar"
	ArtifactOutput ArtifactKind = "output"
)

type ArtifactSelector struct {
	Kind         ArtifactKind `json:"kind"`
	Name         string       `json:"name,omitempty"`
	RelativePath string       `json:"relativePath,omitempty"`
	SHA256       string       `json:"sha256,omitempty"`
}

type TargetSelectors struct {
	Modules     []string           `json:"modules,omitempty"`
	Artifacts   []ArtifactSelector `json:"artifacts,omitempty"`
	SourceRoots []RepoPath         `json:"sourceRoots,omitempty"`
	Paths       []string           `json:"paths,omitempty"`
}

type TargetSpec struct {
	ID                      string          `json:"id"`
	DisplayName             string          `json:"displayName,omitempty"`
	Kind                    TargetKind      `json:"kind"`
	Selectors               TargetSelectors `json:"selectors"`
	IncludeDependencies     bool            `json:"includeDependencies,omitempty"`
	IncludeGeneratedSources bool            `json:"includeGeneratedSources,omitempty"`
	IncludeResources        bool            `json:"includeResources,omitempty"`
	IncludeManifest         bool            `json:"includeManifest,omitempty"`
	Priority                int             `json:"priority,omitempty"`
}

func (profile Profile) EffectiveTargets() []TargetSpec {
	if len(profile.Targets) > 0 {
		return normalizeTargets(profile.Targets)
	}
	modules := uniqueTargetStrings(profile.SelectedModules)
	targets := make([]TargetSpec, 0, len(modules))
	for _, module := range modules {
		targets = append(targets, TargetSpec{
			ID: "module:" + module, DisplayName: module, Kind: TargetMixed,
			Selectors:           TargetSelectors{Modules: []string{module}},
			IncludeDependencies: true, IncludeGeneratedSources: true,
		})
	}
	return targets
}

func (target TargetSpec) Validate() []Diagnostic {
	diagnostics := make([]Diagnostic, 0)
	if strings.TrimSpace(target.ID) == "" {
		diagnostics = append(diagnostics, targetDiagnostic("TARGET_ID_EMPTY", "target id is required"))
	}
	switch target.Kind {
	case TargetJava, TargetAndroid, TargetNative, TargetMixed:
	default:
		diagnostics = append(diagnostics, targetDiagnostic("TARGET_KIND_UNKNOWN", fmt.Sprintf("unknown target kind %q", target.Kind)))
	}
	if len(target.Selectors.Modules) == 0 && len(target.Selectors.Artifacts) == 0 &&
		len(target.Selectors.SourceRoots) == 0 && len(target.Selectors.Paths) == 0 {
		diagnostics = append(diagnostics, targetDiagnostic("TARGET_SELECTORS_EMPTY", "target requires at least one selector"))
	}
	for _, artifact := range target.Selectors.Artifacts {
		switch artifact.Kind {
		case ArtifactJar, ArtifactAPK, ArtifactSrcJar, ArtifactAAR, ArtifactOutput:
		default:
			diagnostics = append(diagnostics, targetDiagnostic("TARGET_ARTIFACT_KIND_UNKNOWN", fmt.Sprintf("unknown artifact kind %q", artifact.Kind)))
		}
		if strings.TrimSpace(artifact.Name) == "" && strings.TrimSpace(artifact.RelativePath) == "" && strings.TrimSpace(artifact.SHA256) == "" {
			diagnostics = append(diagnostics, targetDiagnostic("TARGET_ARTIFACT_IDENTITY_EMPTY", "artifact selector requires name, relativePath, or sha256"))
		}
	}
	return diagnostics
}

func normalizeTargets(values []TargetSpec) []TargetSpec {
	result := append([]TargetSpec(nil), values...)
	for index := range result {
		result[index].ID = strings.TrimSpace(result[index].ID)
		result[index].DisplayName = strings.TrimSpace(result[index].DisplayName)
		result[index].Selectors.Modules = uniqueTargetStrings(result[index].Selectors.Modules)
	}
	sort.SliceStable(result, func(i, j int) bool {
		if result[i].Priority != result[j].Priority {
			return result[i].Priority > result[j].Priority
		}
		return result[i].ID < result[j].ID
	})
	return result
}

func uniqueTargetStrings(values []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(values))
	for _, value := range values {
		value = strings.TrimSpace(value)
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	sort.Strings(result)
	return result
}

func targetDiagnostic(code, message string) Diagnostic {
	return Diagnostic{Code: code, Severity: SeverityError, Stage: "target", Message: message}
}

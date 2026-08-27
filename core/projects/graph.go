package projects

import (
	"sort"

	"github.com/i-rtfsc/as-aosp/core/artifacts"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

type TargetProject struct {
	Spec              model.TargetSpec     `json:"spec"`
	ResolvedModules   []string             `json:"resolvedModules,omitempty"`
	ResolvedArtifacts []artifacts.Artifact `json:"resolvedArtifacts,omitempty"`
	SourceRoots       []metadata.PathRef   `json:"sourceRoots,omitempty"`
	GeneratedRoots    []metadata.PathRef   `json:"generatedRoots,omitempty"`
	Libraries         []metadata.PathRef   `json:"libraries,omitempty"`
	Dependencies      []string             `json:"dependencies,omitempty"`
	Diagnostics       []model.Diagnostic   `json:"diagnostics"`
}

type Ownership struct {
	RepoPath string `json:"repoPath"`
	TargetID string `json:"targetId"`
}

type ProjectGraph struct {
	Targets                  []TargetProject    `json:"targets"`
	SharedLibraries          []metadata.PathRef `json:"sharedLibraries,omitempty"`
	Ownership                []Ownership        `json:"ownership,omitempty"`
	DuplicateSourceDecisions []string           `json:"duplicateSourceDecisions,omitempty"`
	Diagnostics              []model.Diagnostic `json:"diagnostics"`
}

func Build(catalog modules.Catalog, artifactCatalog artifacts.Catalog, targets []model.TargetSpec) ProjectGraph {
	graph := ProjectGraph{Targets: []TargetProject{}, Ownership: []Ownership{}, Diagnostics: []model.Diagnostic{}}
	ownerByPath := map[string]string{}
	for _, spec := range targets {
		project := TargetProject{Spec: spec, Diagnostics: append([]model.Diagnostic(nil), spec.Validate()...)}
		if len(project.Diagnostics) > 0 {
			graph.Targets = append(graph.Targets, project)
			graph.Diagnostics = append(graph.Diagnostics, project.Diagnostics...)
			continue
		}
		moduleNames := append([]string(nil), spec.Selectors.Modules...)
		for _, name := range spec.Selectors.Modules {
			if _, ok := catalog.Modules[name]; !ok {
				project.Diagnostics = append(project.Diagnostics, projectDiagnostic("TARGET_MODULE_NOT_FOUND", "module is not present: "+name))
			}
		}
		for _, selector := range spec.Selectors.Artifacts {
			resolution := artifactCatalog.Resolve(selector)
			project.ResolvedArtifacts = append(project.ResolvedArtifacts, resolution.Artifacts...)
			project.Diagnostics = append(project.Diagnostics, resolution.Diagnostics...)
			for _, artifact := range resolution.Artifacts {
				moduleNames = append(moduleNames, artifact.OwnerModules...)
				project.SourceRoots = append(project.SourceRoots, artifact.SourceAttachments...)
				project.GeneratedRoots = append(project.GeneratedRoots, artifact.GeneratedRoots...)
				if artifact.Kind == model.ArtifactJar || artifact.Kind == model.ArtifactAAR {
					project.Libraries = append(project.Libraries, artifact.Path)
				}
			}
		}
		if spec.IncludeDependencies {
			moduleNames = catalog.Closure(moduleNames)
		} else {
			moduleNames = unique(moduleNames)
		}
		project.ResolvedModules = moduleNames
		for _, name := range moduleNames {
			module, ok := catalog.Modules[name]
			if !ok {
				continue
			}
			project.SourceRoots = append(project.SourceRoots, module.SourceRoots...)
			project.GeneratedRoots = append(project.GeneratedRoots, module.Generated...)
			project.Libraries = append(project.Libraries, append(module.Jars, module.SrcJars...)...)
			project.Dependencies = append(project.Dependencies, module.Dependencies...)
		}
		for _, path := range spec.Selectors.SourceRoots {
			project.SourceRoots = append(project.SourceRoots, metadata.NewPathRef(string(path)))
		}
		project.ResolvedModules = unique(project.ResolvedModules)
		project.Dependencies = unique(project.Dependencies)
		project.SourceRoots = uniquePaths(project.SourceRoots)
		project.GeneratedRoots = uniquePaths(project.GeneratedRoots)
		project.Libraries = uniquePaths(project.Libraries)
		for _, root := range project.SourceRoots {
			path := string(root.RepoPath)
			if path == "" {
				continue
			}
			if existing, ok := ownerByPath[path]; ok && existing != spec.ID {
				graph.DuplicateSourceDecisions = append(graph.DuplicateSourceDecisions, path+": "+existing+" > "+spec.ID)
				continue
			}
			ownerByPath[path] = spec.ID
			graph.Ownership = append(graph.Ownership, Ownership{RepoPath: path, TargetID: spec.ID})
		}
		if len(project.Diagnostics) > 0 {
			graph.Diagnostics = append(graph.Diagnostics, project.Diagnostics...)
		}
		graph.Targets = append(graph.Targets, project)
	}
	sort.Slice(graph.Targets, func(i, j int) bool { return graph.Targets[i].Spec.ID < graph.Targets[j].Spec.ID })
	sort.Slice(graph.Ownership, func(i, j int) bool { return graph.Ownership[i].RepoPath < graph.Ownership[j].RepoPath })
	sort.Strings(graph.DuplicateSourceDecisions)
	return graph
}

func unique(values []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for _, value := range values {
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	sort.Strings(result)
	return result
}
func uniquePaths(values []metadata.PathRef) []metadata.PathRef {
	seen := map[string]metadata.PathRef{}
	for _, value := range values {
		seen[value.Raw] = value
	}
	result := make([]metadata.PathRef, 0, len(seen))
	for _, value := range seen {
		result = append(result, value)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Raw < result[j].Raw })
	return result
}
func projectDiagnostic(code, message string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityError, Stage: "project-graph", Message: message}
}

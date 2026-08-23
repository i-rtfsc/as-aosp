package graph

import (
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
)

type ModuleOrigin string

const (
	OriginJavaDeps   ModuleOrigin = "java-deps"
	OriginCCDeps     ModuleOrigin = "cc-deps"
	OriginModuleInfo ModuleOrigin = "module-info"
)

type Module struct {
	Name            string             `json:"name"`
	Origins         []ModuleOrigin     `json:"origins"`
	ModuleTypes     []string           `json:"moduleTypes,omitempty"`
	Classes         []string           `json:"classes,omitempty"`
	Tags            []string           `json:"tags,omitempty"`
	Sources         []metadata.PathRef `json:"sources,omitempty"`
	Generated       []metadata.PathRef `json:"generated,omitempty"`
	Jars            []metadata.PathRef `json:"jars,omitempty"`
	SrcJars         []metadata.PathRef `json:"srcjars,omitempty"`
	Dependencies    []string           `json:"dependencies,omitempty"`
	AidlIncludeDirs []metadata.PathRef `json:"aidlIncludeDirs,omitempty"`
	AidlSources     []metadata.PathRef `json:"aidlSources,omitempty"`
	CC              *metadata.CCModule `json:"cc,omitempty"`
}

type ModuleGraph struct {
	Modules      map[string]Module                 `json:"modules"`
	CompDBByFile map[string][]metadata.CompDBEntry `json:"compdbByFile"`
	Diagnostics  []model.Diagnostic                `json:"diagnostics"`
}

func (g ModuleGraph) GeneratedPaths(moduleName string) []metadata.PathRef {
	module, ok := g.Modules[moduleName]
	if !ok {
		return []metadata.PathRef{}
	}
	result := addPathRefs(nil, module.Generated)
	for _, source := range module.Sources {
		if source.Kind == metadata.PathGenerated || source.Kind == metadata.PathExternalAbsolute && strings.Contains(strings.ReplaceAll(source.Raw, "\\", "/"), "/out/") {
			result = addPathRefs(result, []metadata.PathRef{source})
		}
	}
	return result
}

func (g ModuleGraph) AidlIncludeDirs(moduleName string) []metadata.PathRef {
	module, ok := g.Modules[moduleName]
	if !ok {
		return []metadata.PathRef{}
	}
	return addPathRefs(nil, module.AidlIncludeDirs)
}

func (g ModuleGraph) AidlOwner(
	repoPath model.RepoPath,
) (string, bool, bool) {
	owners := make([]string, 0)
	for name, module := range g.Modules {
		for _, source := range module.AidlSources {
			if source.RepoPath == repoPath {
				owners = append(owners, name)
				break
			}
		}
	}
	sort.Strings(owners)
	if len(owners) == 1 {
		return owners[0], true, false
	}
	return "", false, len(owners) > 1
}

func (g ModuleGraph) SourceArchives(moduleName string) []metadata.PathRef {
	module, ok := g.Modules[moduleName]
	if !ok {
		return []metadata.PathRef{}
	}
	return addPathRefs(nil, module.SrcJars)
}

func Build(snapshot metadata.Snapshot) ModuleGraph {
	graph := ModuleGraph{Modules: make(map[string]Module), CompDBByFile: make(map[string][]metadata.CompDBEntry), Diagnostics: make([]model.Diagnostic, 0)}
	for name, javaModule := range snapshot.Java {
		module := graph.Modules[name]
		module.Name = name
		module.Origins = addOrigin(module.Origins, OriginJavaDeps)
		module.ModuleTypes = addString(module.ModuleTypes, javaModule.ModuleType)
		module.Sources = addPathRefs(module.Sources, javaModule.Sources)
		module.Jars = addPathRefs(module.Jars, javaModule.Jars)
		module.SrcJars = addPathRefs(module.SrcJars, javaModule.SrcJars)
		module.Dependencies = addStrings(module.Dependencies, javaModule.Dependencies)
		module.AidlIncludeDirs = addPathRefs(module.AidlIncludeDirs, javaModule.AidlIncludeDirs)
		module.AidlSources = addPathRefs(module.AidlSources, javaModule.AidlSources)
		graph.Modules[name] = module
	}
	for name, ccModule := range snapshot.CC {
		module := graph.Modules[name]
		module.Name = name
		module.Origins = addOrigin(module.Origins, OriginCCDeps)
		module.Sources = addPathRefs(module.Sources, ccModule.Sources)
		copy := ccModule
		module.CC = &copy
		graph.Modules[name] = module
	}
	for name, info := range snapshot.ModuleInfo {
		module := graph.Modules[name]
		module.Name = name
		module.Origins = addOrigin(module.Origins, OriginModuleInfo)
		module.Classes = addStrings(module.Classes, info.Classes)
		module.Tags = addStrings(module.Tags, info.Tags)
		module.Sources = addPathRefs(module.Sources, info.Sources)
		module.Generated = addPathRefs(module.Generated, info.Generated)
		module.Dependencies = addStrings(module.Dependencies, info.Dependencies)
		graph.Modules[name] = module
	}
	for _, entry := range snapshot.CompDB {
		ref := metadata.NewPathRef(entry.File)
		key := entry.File
		if ref.RepoPath != "" {
			key = string(ref.RepoPath)
		}
		graph.CompDBByFile[key] = append(graph.CompDBByFile[key], entry)
	}
	return graph
}

func addOrigin(values []ModuleOrigin, value ModuleOrigin) []ModuleOrigin {
	for _, current := range values {
		if current == value {
			return values
		}
	}
	values = append(values, value)
	sort.Slice(values, func(i, j int) bool { return values[i] < values[j] })
	return values
}

func addString(values []string, value string) []string {
	if value == "" {
		return values
	}
	return addStrings(values, []string{value})
}

func addStrings(values, additions []string) []string {
	seen := make(map[string]bool, len(values)+len(additions))
	for _, value := range values {
		if value != "" {
			seen[value] = true
		}
	}
	for _, value := range additions {
		if value != "" {
			seen[value] = true
		}
	}
	result := make([]string, 0, len(seen))
	for value := range seen {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func addPathRefs(values, additions []metadata.PathRef) []metadata.PathRef {
	seen := make(map[string]metadata.PathRef, len(values)+len(additions))
	for _, value := range values {
		seen[string(value.Kind)+"\x00"+value.Raw] = value
	}
	for _, value := range additions {
		seen[string(value.Kind)+"\x00"+value.Raw] = value
	}
	result := make([]metadata.PathRef, 0, len(seen))
	for _, value := range seen {
		result = append(result, value)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Raw < result[j].Raw })
	return result
}

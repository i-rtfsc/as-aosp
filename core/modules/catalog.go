package modules

import (
	"sort"

	"github.com/i-rtfsc/as-aosp/core/metadata"
)

type Module struct {
	Name            string             `json:"name"`
	Origins         []string           `json:"origins"`
	Types           []string           `json:"types,omitempty"`
	Paths           []metadata.PathRef `json:"paths,omitempty"`
	Sources         []metadata.PathRef `json:"sources,omitempty"`
	SourceRoots     []metadata.PathRef `json:"sourceRoots,omitempty"`
	Generated       []metadata.PathRef `json:"generated,omitempty"`
	Jars            []metadata.PathRef `json:"jars,omitempty"`
	SrcJars         []metadata.PathRef `json:"srcjars,omitempty"`
	Dependencies    []string           `json:"dependencies,omitempty"`
	AidlIncludeDirs []metadata.PathRef `json:"aidlIncludeDirs,omitempty"`
	CompilerFlags   []string           `json:"compilerFlags,omitempty"`
}

type Catalog struct {
	Modules map[string]Module `json:"modules"`
}

func Build(snapshot metadata.Snapshot) Catalog {
	catalog := Catalog{Modules: make(map[string]Module)}
	for name, java := range snapshot.Java {
		module := catalog.Modules[name]
		module.Name = name
		module.Origins = appendUnique(module.Origins, "java-deps")
		module.Types = appendUnique(module.Types, java.ModuleType)
		module.Paths = appendPathRefs(module.Paths, java.Paths...)
		module.Sources = appendPathRefs(module.Sources, java.Sources...)
		module.SourceRoots = appendPathRefs(module.SourceRoots, java.Paths...)
		module.Jars = appendPathRefs(module.Jars, java.Jars...)
		module.SrcJars = appendPathRefs(module.SrcJars, java.SrcJars...)
		module.Dependencies = appendUnique(module.Dependencies, java.Dependencies...)
		module.AidlIncludeDirs = appendPathRefs(module.AidlIncludeDirs, java.AidlIncludeDirs...)
		module.CompilerFlags = appendUnique(module.CompilerFlags, java.CompilerFlags...)
		catalog.Modules[name] = normalize(module)
	}
	for name, cc := range snapshot.CC {
		module := catalog.Modules[name]
		module.Name = name
		module.Origins = appendUnique(module.Origins, "cc-deps")
		module.Paths = appendPathRefs(module.Paths, cc.Paths...)
		module.Sources = appendPathRefs(module.Sources, cc.Sources...)
		module.SourceRoots = appendPathRefs(module.SourceRoots, cc.Paths...)
		module.CompilerFlags = appendUnique(module.CompilerFlags, cc.Flags...)
		catalog.Modules[name] = normalize(module)
	}
	for name, info := range snapshot.ModuleInfo {
		module := catalog.Modules[name]
		module.Name = name
		module.Origins = appendUnique(module.Origins, "module-info")
		module.Types = appendUnique(module.Types, info.Classes...)
		module.Paths = appendPathRefs(module.Paths, info.Paths...)
		module.Sources = appendPathRefs(module.Sources, info.Sources...)
		module.SourceRoots = appendPathRefs(module.SourceRoots, info.Paths...)
		module.Generated = appendPathRefs(module.Generated, info.Generated...)
		for _, installed := range info.Installed {
			module.Generated = appendPathRefs(module.Generated, installed)
		}
		module.Dependencies = appendUnique(module.Dependencies, info.Dependencies...)
		catalog.Modules[name] = normalize(module)
	}
	return catalog
}

func (catalog Catalog) Closure(names []string) []string {
	seen := make(map[string]bool)
	queue := append([]string(nil), names...)
	for len(queue) > 0 {
		name := queue[0]
		queue = queue[1:]
		if seen[name] {
			continue
		}
		seen[name] = true
		if module, ok := catalog.Modules[name]; ok {
			queue = append(queue, module.Dependencies...)
		}
	}
	result := make([]string, 0, len(seen))
	for name := range seen {
		result = append(result, name)
	}
	sort.Strings(result)
	return result
}

func normalize(module Module) Module {
	sort.Strings(module.Origins)
	sort.Strings(module.Types)
	sort.Strings(module.Dependencies)
	sort.Strings(module.CompilerFlags)
	sortPathRefs(module.Paths)
	sortPathRefs(module.Sources)
	sortPathRefs(module.SourceRoots)
	sortPathRefs(module.Generated)
	sortPathRefs(module.Jars)
	sortPathRefs(module.SrcJars)
	sortPathRefs(module.AidlIncludeDirs)
	return module
}
func appendUnique(values []string, additions ...string) []string {
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
	return result
}
func appendPathRefs(values []metadata.PathRef, additions ...metadata.PathRef) []metadata.PathRef {
	seen := make(map[string]metadata.PathRef)
	for _, value := range append(values, additions...) {
		seen[value.Raw] = value
	}
	result := make([]metadata.PathRef, 0, len(seen))
	for _, value := range seen {
		result = append(result, value)
	}
	return result
}
func sortPathRefs(values []metadata.PathRef) {
	sort.Slice(values, func(i, j int) bool { return values[i].Raw < values[j].Raw })
}

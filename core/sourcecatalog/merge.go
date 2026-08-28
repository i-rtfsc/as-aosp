package sourcecatalog

import (
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

func MergeBuildFacts(source []Module, build modules.Catalog) []Module {
	result := append([]Module(nil), source...)
	for index := range result {
		facts, found := build.Modules[result[index].Name]
		if !found {
			continue
		}
		result[index].Status = StatusBuildBacked
		result[index].Dependencies = uniqueStrings(append(
			result[index].Dependencies, facts.Dependencies...,
		))
		result[index].SourceRoots = uniqueRepoPaths(append(
			result[index].SourceRoots, repoPaths(facts.SourceRoots)...,
		))
		result[index].Sources = uniqueRepoPaths(append(
			result[index].Sources, repoPaths(facts.Sources)...,
		))
		result[index].Capabilities = mergedCapabilities(
			result[index].Capabilities, facts,
		)
		result[index].Recommended = result[index].Recommended ||
			isBuildRecommended(result[index], facts)
	}
	return result
}

func mergedCapabilities(values []string, facts modules.Module) []string {
	result := append([]string(nil), values...)
	for _, origin := range facts.Origins {
		switch origin {
		case "java-deps":
			result = append(result, "java")
		case "cc-deps":
			result = append(result, "native")
		}
	}
	if len(facts.Generated) > 0 || len(facts.SrcJars) > 0 {
		result = append(result, "generated")
	}
	if len(facts.AidlIncludeDirs) > 0 {
		result = append(result, "aidl")
	}
	if len(facts.Jars) > 0 || len(facts.SrcJars) > 0 {
		result = append(result, "artifact")
	}
	return uniqueStrings(result)
}

func repoPaths(values []metadata.PathRef) []model.RepoPath {
	result := make([]model.RepoPath, 0, len(values))
	for _, value := range values {
		if value.RepoPath != "" && value.Kind == metadata.PathRepoRelative {
			result = append(result, value.RepoPath)
		}
	}
	return result
}

func uniqueRepoPaths(values []model.RepoPath) []model.RepoPath {
	seen := map[model.RepoPath]bool{}
	result := make([]model.RepoPath, 0, len(values))
	for _, value := range values {
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })
	return result
}

func uniqueStrings(values []string) []string {
	seen := map[string]bool{}
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

func isBuildRecommended(module Module, facts modules.Module) bool {
	_ = facts
	return isRecommended(module)
}

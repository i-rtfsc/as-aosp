package profiles

import "sort"

func ApplyDelta(base Catalog, delta ReleaseDelta) Catalog {
	result := base
	for index := range result.Profiles {
		profile := &result.Profiles[index]
		change, ok := delta.Profiles[profile.ID]
		if !ok {
			continue
		}
		profile.ModuleNames = applyList(profile.ModuleNames, change.ModuleNames)
		profile.ModulePatterns = applyList(profile.ModulePatterns, change.ModulePatterns)
		profile.SourceRoots = applyList(profile.SourceRoots, change.SourceRoots)
		profile.SourcePatterns = applyList(profile.SourcePatterns, change.SourcePatterns)
		profile.NativeRoots = applyList(profile.NativeRoots, change.NativeRoots)
		profile.IncludeRoots = applyList(profile.IncludeRoots, change.IncludeRoots)
		profile.Dependencies = applyList(profile.Dependencies, change.Dependencies)
	}
	result.Baseline.AndroidRelease = delta.AndroidRelease
	return result
}

func applyList(values []string, delta ListDelta) []string {
	set := make(map[string]bool, len(values)+len(delta.Add))
	for _, value := range values {
		if value != "" {
			set[value] = true
		}
	}
	for _, value := range delta.Remove {
		delete(set, value)
	}
	for _, value := range delta.Add {
		if value != "" {
			set[value] = true
		}
	}
	result := make([]string, 0, len(set))
	for value := range set {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

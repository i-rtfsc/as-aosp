package selection

import (
	"path"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/profiles"
)

type Preset struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Kind        model.TargetKind `json:"kind"`
	Modules     []string         `json:"modules"`
	Patterns    []string         `json:"patterns,omitempty"`
	Default     bool             `json:"default"`
	Recommended bool             `json:"recommended"`
	Priority    int              `json:"priority"`
}

type Plan struct {
	Targets         []model.TargetSpec `json:"targets"`
	SelectedModules []string           `json:"selectedModules"`
	PresetIDs       []string           `json:"presetIds"`
	Diagnostics     []model.Diagnostic `json:"diagnostics"`
}

func Catalog() []Preset {
	catalog := profiles.RuntimeCatalog("android-17", "android-17")
	result := make([]Preset, 0, len(catalog.Profiles))
	for _, profile := range catalog.Profiles {
		result = append(result, Preset{
			ID: profile.ID, Name: profile.Name, Kind: profileKind(profile.Kind),
			Modules: profile.ModuleNames, Default: profile.Default,
			Recommended: profile.Recommended, Priority: profile.Priority,
			Patterns: profile.ModulePatterns,
		})
	}
	return result
}

func profileKind(value string) model.TargetKind {
	switch strings.ToLower(strings.TrimSpace(value)) {
	case "android":
		return model.TargetAndroid
	case "native":
		return model.TargetNative
	case "mixed":
		return model.TargetMixed
	default:
		return model.TargetJava
	}
}

func Build(presetIDs, manualModules []string, moduleKinds map[string]model.TargetKind) Plan {
	result := Plan{Targets: []model.TargetSpec{}, Diagnostics: []model.Diagnostic{}}
	presetSet := stringSet(presetIDs)
	knownPresets := map[string]bool{}
	assigned := map[string]bool{}
	for _, preset := range Catalog() {
		knownPresets[preset.ID] = true
		if !presetSet[preset.ID] {
			continue
		}
		modules := unassigned(expand(preset, moduleKinds), assigned)
		if len(modules) == 0 {
			continue
		}
		result.PresetIDs = append(result.PresetIDs, preset.ID)
		result.Targets = append(result.Targets, target(
			preset.ID, preset.Name, preset.Kind, modules, preset.Priority,
		))
	}
	for id := range presetSet {
		if !knownPresets[id] {
			result.Diagnostics = append(result.Diagnostics, model.Diagnostic{
				Code: "SELECTION_PRESET_UNKNOWN", Severity: model.SeverityWarning,
				Stage: "selection", Message: "unknown preset: " + id,
			})
		}
	}
	buckets := map[model.TargetKind][]string{}
	for _, name := range uniqueStrings(manualModules) {
		if assigned[name] {
			continue
		}
		kind := moduleKinds[name]
		switch kind {
		case model.TargetJava, model.TargetAndroid, model.TargetNative, model.TargetMixed:
		default:
			kind = model.TargetJava
			result.Diagnostics = append(result.Diagnostics, model.Diagnostic{
				Code: "SELECTION_MODULE_KIND_UNKNOWN", Severity: model.SeverityWarning,
				Stage: "selection", Message: "defaulted module to Java: " + name,
			})
		}
		buckets[kind] = append(buckets[kind], name)
		assigned[name] = true
	}
	for _, item := range []struct {
		kind model.TargetKind
		id   string
		name string
	}{
		{model.TargetJava, "custom-java", "Custom Java"},
		{model.TargetAndroid, "custom-android", "Custom Android"},
		{model.TargetNative, "custom-native", "Custom Native"},
		{model.TargetMixed, "custom-mixed", "Custom Mixed"},
	} {
		modules := uniqueStrings(buckets[item.kind])
		if len(modules) > 0 {
			result.Targets = append(result.Targets, target(
				item.id, item.name, item.kind, modules, 50,
			))
		}
	}
	for name := range assigned {
		result.SelectedModules = append(result.SelectedModules, name)
	}
	sort.Strings(result.PresetIDs)
	sort.Strings(result.SelectedModules)
	return result
}

func expand(preset Preset, available map[string]model.TargetKind) []string {
	values := append([]string(nil), preset.Modules...)
	if len(available) == 0 || len(preset.Patterns) == 0 {
		return uniqueStrings(values)
	}
	for name := range available {
		for _, pattern := range preset.Patterns {
			if matched, err := path.Match(pattern, name); err == nil && matched {
				values = append(values, name)
				break
			}
		}
	}
	return uniqueStrings(values)
}

func DefaultPresetIDs() []string {
	result := []string{}
	for _, preset := range Catalog() {
		if preset.Default {
			result = append(result, preset.ID)
		}
	}
	return result
}

func target(id, name string, kind model.TargetKind, modules []string, priority int) model.TargetSpec {
	return model.TargetSpec{
		ID: id, DisplayName: name, Kind: kind,
		Selectors:           model.TargetSelectors{Modules: uniqueStrings(modules)},
		IncludeDependencies: true, IncludeGeneratedSources: true,
		IncludeResources: kind == model.TargetAndroid || kind == model.TargetMixed,
		IncludeManifest:  kind == model.TargetAndroid || kind == model.TargetMixed,
		Priority:         priority,
	}
}

func unassigned(values []string, assigned map[string]bool) []string {
	result := []string{}
	for _, value := range uniqueStrings(values) {
		if !assigned[value] {
			assigned[value] = true
			result = append(result, value)
		}
	}
	return result
}

func stringSet(values []string) map[string]bool {
	result := map[string]bool{}
	for _, value := range values {
		if value = strings.TrimSpace(value); value != "" {
			result[value] = true
		}
	}
	return result
}

func uniqueStrings(values []string) []string {
	set := stringSet(values)
	result := make([]string, 0, len(set))
	for value := range set {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

package selection

import (
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

type Preset struct {
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Kind        model.TargetKind `json:"kind"`
	Modules     []string         `json:"modules"`
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
	return []Preset{
		{ID: "java-foundation", Name: "Java Foundation", Kind: model.TargetJava,
			Modules: []string{"core-oj", "core-libart", "ext", "framework-minus-apex", "framework"},
			Default: true, Recommended: true, Priority: 700},
		{ID: "system-server-wms", Name: "System Server & WMS", Kind: model.TargetJava,
			Modules: []string{"services", "services.core"},
			Default: true, Recommended: true, Priority: 600},
		{ID: "system-ui", Name: "System UI", Kind: model.TargetAndroid,
			Modules: []string{"SystemUI"}, Default: true, Recommended: true, Priority: 500},
		{ID: "native-foundation", Name: "Native Foundation", Kind: model.TargetNative,
			Modules: []string{"libbase", "libutils", "liblog", "libcutils", "libbinder", "libbinder_ndk", "libnativehelper"},
			Default: true, Recommended: true, Priority: 400},
		{ID: "graphics-surfaceflinger", Name: "Graphics & SurfaceFlinger", Kind: model.TargetNative,
			Modules: []string{"libui", "libgui", "libnativewindow", "surfaceflinger"},
			Default: true, Recommended: true, Priority: 300},
		{ID: "input-stack", Name: "Input Stack", Kind: model.TargetNative,
			Modules: []string{"libinput", "libinputflinger", "libinputdispatcher", "libinputreader", "libinputreporter"},
			Default: true, Recommended: true, Priority: 200},
		{ID: "automotive", Name: "Automotive", Kind: model.TargetMixed,
			Modules:     []string{"android.car", "CarService", "CarServiceUpdatable", "CarSystemUI"},
			Recommended: true, Priority: 100},
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
		modules := unassigned(preset.Modules, assigned)
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

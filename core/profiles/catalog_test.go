package profiles

import "testing"

func TestApplyDeltaKeepsCommonData(t *testing.T) {
	base := DefaultCatalog("android-17", "android-17")
	delta := ReleaseDelta{
		AndroidRelease: "android-18",
		Profiles: map[string]ProfileDelta{
			"system-server-wms": {ModuleNames: ListDelta{Add: []string{"services.new"}, Remove: []string{"services"}}},
		},
	}
	result := ApplyDelta(base, delta)
	for _, profile := range result.Profiles {
		if profile.ID != "system-server-wms" {
			continue
		}
		if contains(profile.ModuleNames, "services") || !contains(profile.ModuleNames, "services.new") {
			t.Fatalf("profile=%+v", profile)
		}
		if len(profile.SourceRoots) == 0 {
			t.Fatal("common source roots were lost")
		}
	}
	if result.Baseline.AndroidRelease != "android-18" {
		t.Fatalf("baseline=%+v", result.Baseline)
	}
}

func TestBuiltInCatalogUsesEmbeddedData(t *testing.T) {
	catalog, err := BuiltInCatalog("android-17")
	if err != nil {
		t.Fatal(err)
	}
	if len(catalog.Profiles) != 28 || catalog.Baseline.AndroidRelease != "android-17" {
		t.Fatalf("catalog=%+v", catalog.Baseline)
	}
}

func TestDefaultCatalogHasExpectedProfiles(t *testing.T) {
	catalog := DefaultCatalog("android-17", "android-17")
	if len(catalog.Profiles) != 28 {
		t.Fatalf("profiles=%d", len(catalog.Profiles))
	}
	if catalog.Baseline.AndroidRelease != "android-17" || catalog.Baseline.AndroidTag != "android-17" {
		t.Fatalf("baseline=%+v", catalog.Baseline)
	}
}

func contains(values []string, expected string) bool {
	for _, value := range values {
		if value == expected {
			return true
		}
	}
	return false
}

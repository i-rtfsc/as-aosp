package selection

import (
	"testing"

	"github.com/i-rtfsc/as-aosp/core/model"
)

func TestBuildDeduplicatesPresetAndManualModules(t *testing.T) {
	plan := Build(
		[]string{"java-foundation", "system-server-wms"},
		[]string{"framework", "services", "CarService"},
		map[string]model.TargetKind{"CarService": model.TargetAndroid},
	)
	if len(plan.Targets) != 3 {
		t.Fatalf("targets=%+v", plan.Targets)
	}
	for _, target := range plan.Targets {
		seen := map[string]bool{}
		for _, module := range target.Selectors.Modules {
			if seen[module] {
				t.Fatalf("duplicate module %s", module)
			}
			seen[module] = true
		}
	}
	if plan.Targets[2].ID != "custom-android" ||
		plan.Targets[2].Selectors.Modules[0] != "CarService" {
		t.Fatalf("custom=%+v", plan.Targets[2])
	}
}

func TestDefaultPresetIDsUsePlatformScope(t *testing.T) {
	ids := DefaultPresetIDs()
	if len(ids) != 6 {
		t.Fatalf("ids=%+v", ids)
	}
	for _, id := range ids {
		if id == "automotive" {
			t.Fatal("Automotive must not default on phone products")
		}
	}
}

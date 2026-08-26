package model

import "testing"

func TestEffectiveTargetsPreservesExplicitArtifactTargets(t *testing.T) {
	profile := Profile{Targets: []TargetSpec{
		{ID: "framework", Kind: TargetJava, Selectors: TargetSelectors{Artifacts: []ArtifactSelector{{Kind: ArtifactJar, Name: "framework.jar"}}}},
		{ID: "services", Kind: TargetJava, Selectors: TargetSelectors{Artifacts: []ArtifactSelector{{Kind: ArtifactJar, Name: "services.jar"}}}},
		{ID: "systemui", Kind: TargetAndroid, Selectors: TargetSelectors{Artifacts: []ArtifactSelector{{Kind: ArtifactAPK, Name: "SystemUI.apk"}}}},
	}}
	targets := profile.EffectiveTargets()
	if len(targets) != 3 || targets[0].ID != "framework" || targets[2].Selectors.Artifacts[0].Kind != ArtifactAPK {
		t.Fatalf("targets=%+v", targets)
	}
}

func TestEffectiveTargetsTranslatesLegacySelectedModules(t *testing.T) {
	targets := (Profile{SelectedModules: []string{" services ", "framework", "services"}}).EffectiveTargets()
	if len(targets) != 2 || targets[0].ID != "module:framework" || targets[1].Selectors.Modules[0] != "services" {
		t.Fatalf("targets=%+v", targets)
	}
	if !targets[0].IncludeDependencies || !targets[0].IncludeGeneratedSources {
		t.Fatalf("legacy target defaults=%+v", targets[0])
	}
}

func TestTargetValidationRejectsEmptyAndMalformedSelectors(t *testing.T) {
	target := TargetSpec{ID: "systemui", Kind: TargetAndroid, Selectors: TargetSelectors{Artifacts: []ArtifactSelector{{Kind: ArtifactAPK}}}}
	diagnostics := target.Validate()
	if len(diagnostics) != 1 || diagnostics[0].Code != "TARGET_ARTIFACT_IDENTITY_EMPTY" {
		t.Fatalf("diagnostics=%+v", diagnostics)
	}
}

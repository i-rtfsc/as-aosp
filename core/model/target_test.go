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

func TestTargetValidationRejectsEmptyAndMalformedSelectors(t *testing.T) {
	target := TargetSpec{ID: "systemui", Kind: TargetAndroid, Selectors: TargetSelectors{Artifacts: []ArtifactSelector{{Kind: ArtifactAPK}}}}
	diagnostics := target.Validate()
	if len(diagnostics) != 1 || diagnostics[0].Code != "TARGET_ARTIFACT_IDENTITY_EMPTY" {
		t.Fatalf("diagnostics=%+v", diagnostics)
	}
}

func TestProfileValidationRequiresExplicitTargets(t *testing.T) {
	diagnostics := (Profile{SchemaVersion: 1, ID: "aosp", WorkspaceViewRoot: "/workspace"}).Validate()
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == "PROFILE_TARGETS_EMPTY" {
			return
		}
	}
	t.Fatalf("missing target diagnostic: %+v", diagnostics)
}

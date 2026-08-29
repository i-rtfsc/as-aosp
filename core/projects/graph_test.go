package projects

import (
	"testing"

	"github.com/i-rtfsc/as-aosp/core/artifacts"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

func TestBuildResolvesFrameworkServicesAndSystemUIProjects(t *testing.T) {
	catalog := modules.Catalog{Modules: map[string]modules.Module{
		"framework": {Name: "framework", SourceRoots: []metadata.PathRef{metadata.NewPathRef("frameworks/base/core/java")}, Dependencies: []string{"base"}},
		"base":      {Name: "base", SourceRoots: []metadata.PathRef{metadata.NewPathRef("frameworks/base")}},
		"services":  {Name: "services", SourceRoots: []metadata.PathRef{metadata.NewPathRef("frameworks/base/services")}, Dependencies: []string{"framework"}},
		"SystemUI":  {Name: "SystemUI", SourceRoots: []metadata.PathRef{metadata.NewPathRef("frameworks/base/packages/SystemUI")}},
	}}
	artifactsCatalog := artifacts.Catalog{Artifacts: map[string]artifacts.Artifact{
		"jar:framework.jar": {ID: "jar:framework.jar", Kind: model.ArtifactJar, Name: "framework.jar", OwnerModules: []string{"framework"}},
		"jar:services.jar":  {ID: "jar:services.jar", Kind: model.ArtifactJar, Name: "services.jar", OwnerModules: []string{"services"}},
		"apk:SystemUI.apk":  {ID: "apk:SystemUI.apk", Kind: model.ArtifactAPK, Name: "SystemUI.apk", OwnerModules: []string{"SystemUI"}},
	}}
	graph := Build(catalog, artifactsCatalog, []model.TargetSpec{
		{ID: "framework", Kind: model.TargetJava, Selectors: model.TargetSelectors{Artifacts: []model.ArtifactSelector{{Kind: model.ArtifactJar, Name: "framework.jar"}}}, IncludeDependencies: true},
		{ID: "services", Kind: model.TargetJava, Selectors: model.TargetSelectors{Artifacts: []model.ArtifactSelector{{Kind: model.ArtifactJar, Name: "services.jar"}}}},
		{ID: "systemui", Kind: model.TargetAndroid, Selectors: model.TargetSelectors{Artifacts: []model.ArtifactSelector{{Kind: model.ArtifactAPK, Name: "SystemUI.apk"}}}},
	})
	if len(graph.Targets) != 3 || len(graph.Diagnostics) != 0 || len(graph.Targets[0].ResolvedArtifacts) != 1 {
		t.Fatalf("graph=%+v", graph)
	}
	systemUI := graph.Targets[2]
	if len(systemUI.SourceRoots) != 1 || systemUI.SourceRoots[0].Raw != "frameworks/base/packages/SystemUI" {
		t.Fatalf("SystemUI APK did not project its owner source root: %+v", systemUI)
	}
	if len(systemUI.Libraries) != 0 {
		t.Fatalf("SystemUI APK must not become a Java library: %+v", systemUI.Libraries)
	}
	services := graph.Targets[1]
	if len(services.TargetDependencies) != 1 || services.TargetDependencies[0] != "framework" {
		t.Fatalf("services target dependencies=%v", services.TargetDependencies)
	}
	framework := graph.Targets[0]
	if len(framework.OwnedSourceRoots) != 1 ||
		framework.OwnedSourceRoots[0].Raw != "frameworks/base/core/java" {
		t.Fatalf("framework owned roots=%v", framework.OwnedSourceRoots)
	}
	if len(framework.ExcludedSourceRoots) != 2 ||
		framework.ExcludedSourceRoots[0].Raw != "frameworks/base/packages/SystemUI" ||
		framework.ExcludedSourceRoots[1].Raw != "frameworks/base/services" {
		t.Fatalf("framework excluded roots=%v", framework.ExcludedSourceRoots)
	}
}

func TestBuildReportsMissingArtifact(t *testing.T) {
	graph := Build(
		modules.Catalog{Modules: map[string]modules.Module{}},
		artifacts.Catalog{Artifacts: map[string]artifacts.Artifact{}},
		[]model.TargetSpec{{
			ID: "framework", Kind: model.TargetJava,
			Selectors: model.TargetSelectors{Artifacts: []model.ArtifactSelector{{Kind: model.ArtifactJar, Name: "framework.jar"}}},
		}},
	)
	if len(graph.Targets) != 1 || len(graph.Targets[0].Diagnostics) != 1 ||
		graph.Targets[0].Diagnostics[0].Code != "ARTIFACT_NOT_FOUND" {
		t.Fatalf("missing artifact diagnostics=%+v", graph)
	}
}

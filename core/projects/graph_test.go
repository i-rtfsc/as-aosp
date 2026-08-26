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
		"services":  {Name: "services", SourceRoots: []metadata.PathRef{metadata.NewPathRef("frameworks/base/services")}},
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
}

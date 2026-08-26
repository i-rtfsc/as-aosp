package artifacts

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

func TestCatalogResolvesFrameworkJarAndOwners(t *testing.T) {
	catalog := Build(modules.Catalog{Modules: map[string]modules.Module{
		"framework-minus-apex": {Name: "framework-minus-apex", Jars: []metadata.PathRef{metadata.NewPathRef("out/soong/framework.jar")}},
		"framework":            {Name: "framework", Jars: []metadata.PathRef{metadata.NewPathRef("out/soong/framework.jar")}},
	}})
	resolution := catalog.Resolve(model.ArtifactSelector{Kind: model.ArtifactJar, RelativePath: "out/soong/framework.jar"})
	if resolution.Match != MatchExact || len(resolution.Artifacts) != 1 || len(resolution.Artifacts[0].OwnerModules) != 2 {
		t.Fatalf("resolution=%+v", resolution)
	}
}

func TestCatalogRejectsAmbiguousBasename(t *testing.T) {
	catalog := Build(modules.Catalog{Modules: map[string]modules.Module{
		"one": {Name: "one", Jars: []metadata.PathRef{metadata.NewPathRef("out/one/services.jar")}},
		"two": {Name: "two", Jars: []metadata.PathRef{metadata.NewPathRef("out/two/services.jar")}},
	}})
	resolution := catalog.Resolve(model.ArtifactSelector{Kind: model.ArtifactJar, Name: "services.jar"})
	if resolution.Match != MatchMissing || len(resolution.Diagnostics) != 1 || resolution.Diagnostics[0].Code != "ARTIFACT_AMBIGUOUS" {
		t.Fatalf("resolution=%+v", resolution)
	}
}

func TestCatalogSupportsExplicitAPKFacts(t *testing.T) {
	root := t.TempDir()
	pathName := filepath.Join(root, "SystemUI.apk")
	if err := os.WriteFile(pathName, []byte("apk"), 0o644); err != nil {
		t.Fatal(err)
	}
	catalog := Catalog{Artifacts: map[string]Artifact{}}
	catalog.AddExternal(model.ArtifactAPK, pathName, []string{"SystemUI"})
	resolution := catalog.Resolve(model.ArtifactSelector{Kind: model.ArtifactAPK, Name: "SystemUI.apk"})
	if resolution.Match != MatchExact || resolution.Artifacts[0].OwnerModules[0] != "SystemUI" {
		t.Fatalf("resolution=%+v", resolution)
	}
}

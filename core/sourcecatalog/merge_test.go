package sourcecatalog

import (
	"testing"

	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

func TestMergeBuildFactsUpgradesSourceModule(t *testing.T) {
	source := []Module{{
		Name: "framework-minus-apex", Type: "java_library",
		Directory: "frameworks/base", Status: StatusSourceOnly,
	}}
	build := modules.Catalog{Modules: map[string]modules.Module{
		"framework-minus-apex": {
			Name: "framework-minus-apex", Origins: []string{"java-deps"},
			SourceRoots: []metadata.PathRef{
				metadata.NewPathRef("frameworks/base/core/java"),
			},
			SrcJars: []metadata.PathRef{
				metadata.NewPathRef("out/soong/framework.srcjar"),
			},
		},
	}}
	merged := MergeBuildFacts(source, build)
	if merged[0].Status != StatusBuildBacked ||
		!contains(merged[0].Capabilities, "generated") ||
		!contains(merged[0].Capabilities, "artifact") ||
		merged[0].BuildRole != RoleLibrary || !contains(merged[0].OutputKinds, "jar") {
		t.Fatalf("merged=%+v", merged[0])
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

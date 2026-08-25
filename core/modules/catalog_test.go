package modules

import (
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"testing"
)

func TestBuildMergesJavaAndModuleInfoFacts(t *testing.T) {
	snapshot := metadata.Snapshot{Java: map[string]metadata.JavaModule{"framework": {Name: "framework", Paths: []metadata.PathRef{metadata.NewPathRef("frameworks/base/core/java")}, Jars: []metadata.PathRef{metadata.NewPathRef("out/framework.jar")}, SrcJars: []metadata.PathRef{metadata.NewPathRef("out/framework.srcjar")}, Dependencies: []string{"core"}}}, ModuleInfo: map[string]metadata.ModuleInfo{"framework": {Name: "framework", Sources: []metadata.PathRef{metadata.NewPathRef("frameworks/base/core/java/android/os/Binder.java")}}}}
	catalog := Build(snapshot)
	module := catalog.Modules["framework"]
	if len(module.Origins) != 2 || len(module.SourceRoots) != 1 || len(module.Jars) != 1 || len(module.SrcJars) != 1 {
		t.Fatalf("module=%+v", module)
	}
}

func TestClosureIncludesDependencies(t *testing.T) {
	catalog := Catalog{Modules: map[string]Module{"framework": {Name: "framework", Dependencies: []string{"core"}}, "core": {Name: "core"}}}
	closure := catalog.Closure([]string{"framework"})
	if len(closure) != 2 || closure[0] != "core" || closure[1] != "framework" {
		t.Fatalf("closure=%v", closure)
	}
}

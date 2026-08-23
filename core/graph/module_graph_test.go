package graph

import (
	"testing"

	"github.com/i-rtfsc/as-aosp/core/metadata"
)

func TestBuildMergesJavaCCAndModuleInfo(t *testing.T) {
	snapshot := metadata.Snapshot{
		Java:       map[string]metadata.JavaModule{"foo": {Name: "foo", ModuleType: "java_library", Dependencies: []string{"bar"}, Sources: []metadata.PathRef{metadata.NewPathRef("frameworks/base/Foo.java")}}},
		CC:         map[string]metadata.CCModule{"foo": {Name: "foo", Sources: []metadata.PathRef{metadata.NewPathRef("frameworks/base/Foo.cpp")}}},
		ModuleInfo: map[string]metadata.ModuleInfo{"foo": {Name: "foo", Classes: []string{"JAVA_LIBRARIES"}, Tags: []string{"optional"}, Dependencies: []string{"bar", "baz"}}},
		CompDB:     []metadata.CompDBEntry{{File: "frameworks/base/Foo.cpp", Arguments: []string{"clang++"}}},
	}
	result := Build(snapshot)
	module := result.Modules["foo"]
	if len(module.Origins) != 3 || len(module.Sources) != 2 || len(module.Dependencies) != 2 || module.CC == nil {
		t.Fatalf("module: %+v", module)
	}
	if len(result.CompDBByFile["frameworks/base/Foo.cpp"]) != 1 {
		t.Fatalf("compdb: %+v", result.CompDBByFile)
	}
}

func TestGeneratedPathsIncludesOutBackedSources(t *testing.T) {
	result := Build(metadata.Snapshot{Java: map[string]metadata.JavaModule{"foo": {
		Name:    "foo",
		Sources: []metadata.PathRef{metadata.NewPathRef("/srv/aosp/out/generated/android/demo/IFoo.java")},
		SrcJars: []metadata.PathRef{metadata.NewPathRef("/srv/aosp/out/generated/foo.srcjar")},
	}}})
	if paths := result.GeneratedPaths("foo"); len(paths) != 1 {
		t.Fatalf("paths: %+v", paths)
	}
}

func TestAidlOwnerUsesExactMetadataSource(t *testing.T) {
	result := Build(metadata.Snapshot{Java: map[string]metadata.JavaModule{
		"framework": {
			Name: "framework",
			AidlSources: []metadata.PathRef{
				metadata.NewPathRef("android/demo/IFoo.aidl"),
			},
		},
	}})
	owner, found, ambiguous := result.AidlOwner("android/demo/IFoo.aidl")
	if !found || ambiguous || owner != "framework" {
		t.Fatalf(
			"owner=%q found=%v ambiguous=%v",
			owner,
			found,
			ambiguous,
		)
	}
}

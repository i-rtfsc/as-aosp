package metadata

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDirectoryReadsAllSupportedMetadata(t *testing.T) {
	report := LoadDirectory(filepath.Join("..", "..", "testdata", "metadata", "out"))
	if len(report.Diagnostics) != 0 {
		t.Fatalf("diagnostics: %+v", report.Diagnostics)
	}
	if len(report.Snapshot.Java) != 1 || len(report.Snapshot.CC) != 1 || len(report.Snapshot.ModuleInfo) != 1 || len(report.Snapshot.CompDB) != 1 {
		t.Fatalf("snapshot: %+v", report.Snapshot)
	}
	java := report.Snapshot.Java["framework-minus-apex"]
	if !hasRepoPath(
		java.Sources,
		"frameworks/base/core/java/android/os/Bundle.java",
	) {
		t.Fatalf("java source: %+v", java.Sources)
	}
	if java.Jars[0].Kind != PathExternalAbsolute {
		t.Fatalf("jar path kind: %+v", java.Jars[0])
	}
	if len(java.AidlSources) != 1 ||
		java.AidlSources[0].RepoPath != "android/demo/IFoo.aidl" {
		t.Fatalf("AIDL sources: %+v", java.AidlSources)
	}
}

func hasRepoPath(values []PathRef, expected string) bool {
	for _, value := range values {
		if string(value.RepoPath) == expected {
			return true
		}
	}
	return false
}

func TestLoadDirectoryIsolatesMissingReader(t *testing.T) {
	root := t.TempDir()
	if err := os.WriteFile(filepath.Join(root, "module_bp_java_deps.json"), []byte(`{"java":{"srcs":["x.java"]}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	report := LoadDirectory(root)
	if len(report.Snapshot.Java) != 1 {
		t.Fatalf("java snapshot: %+v", report.Snapshot.Java)
	}
	if len(report.Diagnostics) != 3 {
		t.Fatalf("expected three missing reader diagnostics, got %+v", report.Diagnostics)
	}
}

func TestJavaDepsKeepsStaticLibrariesAsOwnershipDependencies(t *testing.T) {
	modules := map[string]JavaModule{}
	err := decodeJavaDeps([]byte(`{
  "services":{"module_type":"system_java_library","static_libs":["services.impl"]},
  "combiner":{"module_type":"java_genrule_combiner","dependencies":["part.one"]}
}`), &modules)
	if err != nil {
		t.Fatal(err)
	}
	if len(modules["services"].DirectDependencies) != 1 ||
		modules["services"].DirectDependencies[0] != "services.impl" {
		t.Fatalf("services=%+v", modules["services"])
	}
	if len(modules["combiner"].DirectDependencies) != 1 ||
		modules["combiner"].DirectDependencies[0] != "part.one" {
		t.Fatalf("combiner=%+v", modules["combiner"])
	}
}

func TestNewPathRefKeepsExternalAbsolutePath(t *testing.T) {
	ref := NewPathRef(`/srv/aosp/frameworks/base/core/java/Foo.java`)
	if ref.Kind != PathExternalAbsolute || ref.RepoPath != "" || ref.ExternalPath == "" {
		t.Fatalf("ref: %+v", ref)
	}
}

func TestNewPathRefMarksRelativeOutPathGenerated(t *testing.T) {
	ref := NewPathRef(`out/soong/.intermediates/foo/generated.srcjar`)
	if ref.Kind != PathGenerated || ref.RepoPath == "" {
		t.Fatalf("ref: %+v", ref)
	}
}

func TestLoadDirectoryFindsProductModuleInfo(t *testing.T) {
	root := t.TempDir()
	productDir := filepath.Join(root, "target", "product", "demo")
	if err := os.MkdirAll(productDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(productDir, "module-info.json"), []byte(`{"demo":{"path":["packages/apps/Demo"]}}`), 0o644); err != nil {
		t.Fatal(err)
	}
	report := LoadDirectory(root)
	if _, ok := report.Snapshot.ModuleInfo["demo"]; !ok {
		t.Fatalf("module info: %+v", report.Snapshot.ModuleInfo)
	}
}

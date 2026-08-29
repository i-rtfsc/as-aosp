package sourcecatalog

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscoverParsesBlueprintModulesAndRecommendations(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "frameworks", "base", "Android.bp")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	source := `java_library { name: "framework-minus-apex", srcs: ["core/java/**/*.java"], static_libs: ["core-oj"], }
android_app { name: "SystemUI", srcs: ["src/**/*.java"], }`
	if err := os.WriteFile(path, []byte(source), 0o644); err != nil {
		t.Fatal(err)
	}
	report := Discover(root, filepath.Join(root, ".cache"))
	if len(report.Catalog.Modules) != 2 || report.CacheHit {
		t.Fatalf("report=%+v", report)
	}
	if !report.Catalog.Modules[0].Recommended || report.Catalog.Modules[0].Name != "SystemUI" && report.Catalog.Modules[1].Name != "SystemUI" {
		t.Fatalf("modules=%+v", report.Catalog.Modules)
	}
	for _, module := range report.Catalog.Modules {
		if module.Name == "SystemUI" && (module.BuildRole != RoleProduct || !module.DefaultVisible ||
			!contains(module.OutputKinds, "apk")) {
			t.Fatalf("SystemUI classification=%+v", module)
		}
	}
	second := Discover(root, filepath.Join(root, ".cache"))
	if !second.CacheHit || len(second.Catalog.Modules) != 2 {
		t.Fatalf("second=%+v", second)
	}
}

func TestDiscoverKeepsOtherFilesWhenOneBlueprintIsInvalid(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "good"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Join(root, "bad"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "good", "Android.bp"), []byte(`cc_library { name: "libbase", }`), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "bad", "Android.bp"), []byte(`java_library { name: "broken",`), 0o644); err != nil {
		t.Fatal(err)
	}
	report := Discover(root, "")
	if len(report.Catalog.Modules) != 1 || report.Catalog.Modules[0].Name != "libbase" {
		t.Fatalf("report=%+v", report)
	}
	if len(report.Diagnostics) == 0 {
		t.Fatal("expected parse diagnostic")
	}
}

func TestDiscoverClassifiesTestsAndBuildSupport(t *testing.T) {
	root := t.TempDir()
	path := filepath.Join(root, "Android.bp")
	source := `
java_test { name: "framework-tests" }
se_build_files { name: "202604.board.ignore.map" }
cc_library_shared { name: "libbinder" }
`
	if err := os.WriteFile(path, []byte(source), 0o644); err != nil {
		t.Fatal(err)
	}
	report := Discover(root, "")
	byName := make(map[string]Module, len(report.Catalog.Modules))
	for _, module := range report.Catalog.Modules {
		byName[module.Name] = module
	}
	if byName["framework-tests"].BuildRole != RoleTest || byName["framework-tests"].DefaultVisible {
		t.Fatalf("test classification=%+v", byName["framework-tests"])
	}
	if byName["202604.board.ignore.map"].BuildRole != RoleBuildSupport || byName["202604.board.ignore.map"].DefaultVisible {
		t.Fatalf("support classification=%+v", byName["202604.board.ignore.map"])
	}
	if byName["libbinder"].BuildRole != RoleLibrary || !byName["libbinder"].DefaultVisible ||
		!contains(byName["libbinder"].OutputKinds, "so") {
		t.Fatalf("library classification=%+v", byName["libbinder"])
	}
}

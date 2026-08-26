package main

import (
	"archive/zip"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestExportPackIncludesCoreMetadataAndSrcjarsByDefault(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "module-info.json"), `{}`)
	writePackFixture(t, filepath.Join(out, "module_bp_java_deps.json"), `{}`)
	writeSrcJar(t, filepath.Join(out, "soong", ".intermediates", "framework", "gen", "framework.srcjar"))
	destination := filepath.Join(root, "pack")
	if err := exportPack(root, out, destination, false, nil); err != nil {
		t.Fatal(err)
	}
	for _, relative := range []string{"manifest.json", "checksums.sha256", "module-info.json", "module_bp_java_deps.json", filepath.Join("generated", "manifest.json"), filepath.Join("generated", "srcjars", "framework", "gen", "framework.srcjar"), filepath.Join("modules", "catalog.json"), filepath.Join("modules", "dependency-graph.json"), filepath.Join("aidl", "symbols.json"), filepath.Join("aidl", "index.json"), filepath.Join("aidl", "modules.json")} {
		if _, err := os.Stat(filepath.Join(destination, relative)); err != nil {
			t.Fatalf("missing %s: %v", relative, err)
		}
	}
	var manifest packManifest
	data, _ := os.ReadFile(filepath.Join(destination, "manifest.json"))
	if err := json.Unmarshal(data, &manifest); err != nil || len(manifest.Packs) != 1 || manifest.Packs[0] != "core" {
		t.Fatalf("manifest=%+v err=%v", manifest, err)
	}
	var generated []struct {
		Entries []string `json:"entries"`
	}
	data, _ = os.ReadFile(filepath.Join(destination, "generated", "manifest.json"))
	if err := json.Unmarshal(data, &generated); err != nil || len(generated) != 1 || len(generated[0].Entries) != 1 {
		t.Fatalf("generated=%+v err=%v", generated, err)
	}
}

func TestExportPackReadsSoongMetadataCandidate(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "soong", "module_bp_java_deps.json"), `{}`)
	destination := filepath.Join(root, "pack")
	if err := exportPack(root, out, destination, false, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(destination, "module_bp_java_deps.json")); err != nil {
		t.Fatal(err)
	}
}

func TestExportPackCopiesReferencedJarsIntoPack(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "module-info.json"), `{}`)
	writePackFixture(t, filepath.Join(out, "module_bp_java_deps.json"), `{"framework":{"jars":["/srv/aosp/out/soong/framework.jar"]}}`)
	jar := filepath.Join(out, "soong", "framework.jar")
	writePackFixture(t, jar, "jar-bytes")
	// The fixture uses a different recorded build root, so no copy is expected.
	destination := filepath.Join(root, "pack")
	if err := exportPack(root, out, destination, false, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(destination, "out", "soong", "framework.jar")); !os.IsNotExist(err) {
		t.Fatalf("unrelated absolute jar was copied: %v", err)
	}
}

func TestExportPackCopiesRelativeReferencedJars(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "module-info.json"), `{}`)
	writePackFixture(t, filepath.Join(out, "module_bp_java_deps.json"), `{"framework":{"jars":["soong/framework.jar"]}}`)
	writeJar(t, filepath.Join(out, "soong", "framework.jar"))
	destination := filepath.Join(root, "pack")
	if err := exportPack(root, out, destination, false, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(destination, "out", "soong", "framework.jar")); err != nil {
		t.Fatalf("relative jar missing: %v", err)
	}
}

func TestExportPackCopiesOutPrefixedReferencedJars(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "module-info.json"), `{}`)
	writePackFixture(t, filepath.Join(out, "module_bp_java_deps.json"), `{"framework":{"jars":["out/soong/framework.jar"]}}`)
	writeJar(t, filepath.Join(out, "soong", "framework.jar"))
	destination := filepath.Join(root, "pack")
	if err := exportPack(root, out, destination, false, nil); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(destination, "out", "soong", "framework.jar")); err != nil {
		t.Fatalf("out-prefixed jar missing: %v", err)
	}
}

func writeJar(t *testing.T, pathName string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(pathName), 0o755); err != nil {
		t.Fatal(err)
	}
	file, err := os.Create(pathName)
	if err != nil {
		t.Fatal(err)
	}
	writer := zip.NewWriter(file)
	entry, err := writer.Create("META-INF/MANIFEST.MF")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = entry.Write([]byte("Manifest-Version: 1.0\n"))
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestValidatePackRejectsChecksumMismatch(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	destination := filepath.Join(root, "pack")
	writePackFixture(t, filepath.Join(out, "module-info.json"), `{}`)
	if err := exportPack(root, out, destination, false, nil); err != nil {
		t.Fatal(err)
	}
	writePackFixture(t, filepath.Join(destination, "module-info.json"), `{"tampered":true}`)
	if err := validatePack(destination); err == nil {
		t.Fatal("tampered pack was accepted")
	}
}

func TestValidatePackRejectsUntrackedFile(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	destination := filepath.Join(root, "pack")
	writePackFixture(t, filepath.Join(out, "module-info.json"), `{}`)
	if err := exportPack(root, out, destination, false, nil); err != nil {
		t.Fatal(err)
	}
	writePackFixture(t, filepath.Join(destination, "unexpected.json"), `{}`)
	if err := validatePack(destination); err == nil {
		t.Fatal("pack with untracked file was accepted")
	}
}

func TestValidateNativePackRequiresJNIIndex(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	destination := filepath.Join(root, "pack")
	if err := exportPack(root, out, destination, true, nil); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(filepath.Join(destination, "jni", "index.json")); err != nil {
		t.Fatal(err)
	}
	if err := validatePack(destination); err == nil {
		t.Fatal("native pack without JNI index was accepted")
	}
}

func TestExportPackAddsNativeFilesOnlyWhenRequested(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "module_bp_cc_deps.json"), `{}`)
	writePackFixture(t, filepath.Join(out, "development", "ide", "compdb", "compile_commands.json"), `[]`)
	destination := filepath.Join(root, "pack")
	if err := exportPack(root, out, destination, true, nil); err != nil {
		t.Fatal(err)
	}
	for _, relative := range []string{"module_bp_cc_deps.json", filepath.Join("development", "ide", "compdb", "compile_commands.json")} {
		if _, err := os.Stat(filepath.Join(destination, relative)); err != nil {
			t.Fatalf("missing %s", relative)
		}
	}
	if _, err := os.Stat(filepath.Join(destination, "jni", "index.json")); err != nil {
		t.Fatalf("missing JNI index: %v", err)
	}
}

func TestExportPackFiltersAIDLBySelectedModuleClosure(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "module-info.json"), `{
  "framework": {"path":["frameworks/base"],"dependencies":["core"]},
  "core": {"path":["system/core"]},
  "other": {"path":["packages/apps/Other"]}
}`)
	writePackFixture(t, filepath.Join(root, "frameworks", "base", "demo", "IFoo.aidl"), "package demo; interface IFoo {}")
	writePackFixture(t, filepath.Join(root, "system", "core", "demo", "IBar.aidl"), "package demo; interface IBar {}")
	writePackFixture(t, filepath.Join(root, "packages", "apps", "Other", "demo", "IOther.aidl"), "package demo; interface IOther {}")
	destination := filepath.Join(root, "pack")
	if err := exportPack(root, out, destination, false, []string{"framework"}); err != nil {
		t.Fatal(err)
	}
	data, _ := os.ReadFile(filepath.Join(destination, "aidl", "index.json"))
	var index struct {
		Documents map[string]any `json:"documents"`
	}
	if err := json.Unmarshal(data, &index); err != nil {
		t.Fatal(err)
	}
	if len(index.Documents) != 2 || index.Documents["packages/apps/Other/demo/IOther.aidl"] != nil {
		t.Fatalf("documents=%v", index.Documents)
	}
}

func TestExportPackFiltersSrcJarsBySelectedModuleFacts(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "module-info.json"), `{"framework":{"path":["frameworks/base"]}}`)
	writePackFixture(t, filepath.Join(out, "module_bp_java_deps.json"), `{"framework":{"srcjars":["/srv/aosp/out/soong/.intermediates/framework/gen/framework.srcjar"]}}`)
	writeSrcJar(t, filepath.Join(out, "soong", ".intermediates", "framework", "gen", "framework.srcjar"))
	writeSrcJar(t, filepath.Join(out, "soong", ".intermediates", "other", "gen", "other.srcjar"))
	destination := filepath.Join(root, "pack")
	if err := exportPack(root, out, destination, false, []string{"framework"}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(destination, "generated", "srcjars", "framework", "gen", "framework.srcjar")); err != nil {
		t.Fatalf("selected archive missing: %v", err)
	}
	if _, err := os.Stat(filepath.Join(destination, "generated", "srcjars", "other", "gen", "other.srcjar")); !os.IsNotExist(err) {
		t.Fatalf("unselected archive was exported: %v", err)
	}
}

func TestExportPackRejectsUnknownSelectedModule(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "module-info.json"), `{}`)
	if err := exportPack(root, out, filepath.Join(root, "pack"), false, []string{"missing"}); err == nil {
		t.Fatal("unknown selected module was accepted")
	}
}

func writePackFixture(t *testing.T, pathName, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(pathName), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(pathName, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
func writeSrcJar(t *testing.T, pathName string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(pathName), 0o755); err != nil {
		t.Fatal(err)
	}
	file, err := os.Create(pathName)
	if err != nil {
		t.Fatal(err)
	}
	writer := zip.NewWriter(file)
	entry, err := writer.Create("android/demo/IFoo.java")
	if err != nil {
		t.Fatal(err)
	}
	_, _ = entry.Write([]byte("package android.demo; class IFoo {}"))
	_ = writer.Close()
	_ = file.Close()
}

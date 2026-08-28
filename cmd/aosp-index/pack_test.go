package main

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func captureStdout(t *testing.T, action func()) string {
	t.Helper()
	old := os.Stdout
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = writer
	action()
	_ = writer.Close()
	os.Stdout = old
	var buffer bytes.Buffer
	_, _ = buffer.ReadFrom(reader)
	_ = reader.Close()
	return buffer.String()
}

func TestExportPortableOutPreservesRelativePaths(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "target", "product", "demo", "module-info.json"), `{"framework":{"path":["frameworks/base"]}}`)
	writePackFixture(t, filepath.Join(out, "soong", "module_bp_java_deps.json"), `{"framework":{"jars":["out/soong/framework.jar"],"srcjars":["out/soong/.intermediates/framework/gen/framework.srcjar"]}}`)
	writePackFixture(t, filepath.Join(out, "soong", "module_bp_cc_deps.json"), `{"modules":{}}`)
	writePackFixture(t, filepath.Join(out, "soong", "development", "ide", "compdb", "compile_commands.json"), `[]`)
	writePackFixture(t, filepath.Join(out, "soong", "rust-project.json"), `{"crates":[]}`)
	writePackFixture(t, filepath.Join(out, "development", "ide", "clion", "frameworks", "base", "libdemo", "CMakeLists.txt"), `project(libdemo)`)
	writeJar(t, filepath.Join(out, "soong", "framework.jar"))
	writeSrcJar(t, filepath.Join(out, "soong", ".intermediates", "framework", "gen", "framework.srcjar"))
	destination := filepath.Join(root, "portable-out")
	if err := exportPack(root, out, destination, true, true, nil); err != nil {
		t.Fatal(err)
	}
	for _, relative := range []string{
		"target/product/demo/module-info.json",
		"soong/module_bp_java_deps.json",
		"soong/module_bp_cc_deps.json",
		"soong/development/ide/compdb/compile_commands.json",
		"soong/rust-project.json",
		"development/ide/clion/frameworks/base/libdemo/CMakeLists.txt",
		"soong/framework.jar",
		"soong/.intermediates/framework/gen/framework.srcjar",
	} {
		if _, err := os.Stat(filepath.Join(destination, filepath.FromSlash(relative))); err != nil {
			t.Fatalf("missing portable out file %s: %v", relative, err)
		}
	}
	for _, forbidden := range []string{"manifest.json", "checksums.sha256", "modules", "artifacts", "aidl", "jni", "generated"} {
		if _, err := os.Stat(filepath.Join(destination, forbidden)); !os.IsNotExist(err) {
			t.Fatalf("portable out contains private path %s: %v", forbidden, err)
		}
	}
}

func TestExportPortableOutRequiresCoreFacts(t *testing.T) {
	root := t.TempDir()
	out := filepath.Join(root, "out")
	writePackFixture(t, filepath.Join(out, "soong", "module_bp_java_deps.json"), `{}`)
	if err := exportPack(root, out, filepath.Join(root, "portable-out"), false, false, nil); err == nil || !strings.Contains(err.Error(), "module-info") {
		t.Fatalf("error=%v", err)
	}
}

func TestPackInspectReadsPortableOutMetadata(t *testing.T) {
	out := t.TempDir()
	writePackFixture(t, filepath.Join(out, "target", "product", "demo", "module-info.json"), `{"framework":{"path":["frameworks/base"]}}`)
	writePackFixture(t, filepath.Join(out, "soong", "module_bp_java_deps.json"), `{"framework":{"jars":["soong/framework.jar"]}}`)
	output := captureStdout(t, func() { runPackInspect([]string{"--root", out, "--module", "framework"}) })
	var result struct {
		Modules int            `json:"modules"`
		Jars    map[string]any `json:"jars"`
	}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		t.Fatal(err)
	}
	if result.Modules != 1 {
		t.Fatalf("result=%+v", result)
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
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
}

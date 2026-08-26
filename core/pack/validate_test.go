package pack

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestValidateRejectsUntrackedFile(t *testing.T) {
	root := t.TempDir()
	for _, file := range []string{"manifest.json", "checksums.sha256", "generated/manifest.json", "modules/catalog.json", "modules/dependency-graph.json", "aidl/symbols.json", "aidl/index.json", "aidl/modules.json"} {
		if err := os.MkdirAll(filepath.Dir(filepath.Join(root, file)), 0o755); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(root, file), []byte("{}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.WriteFile(filepath.Join(root, "extra"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	if Validate(root).Valid {
		t.Fatal("accepted untracked file")
	}
}

func TestValidateRejectsGeneratedEntryDrift(t *testing.T) {
	root := minimalPack(t)
	archive := filepath.Join(root, "generated", "srcjars", "demo.srcjar")
	writeZip(t, archive, "demo/Foo.java")
	manifest := []map[string]any{{"path": "generated/srcjars/demo.srcjar", "entries": []string{"demo/Other.java"}}}
	writeJSON(t, filepath.Join(root, "generated", "manifest.json"), manifest)
	refreshChecksums(t, root)
	result := Validate(root)
	if result.Valid || result.Diagnostics[0].Code != "PACK_GENERATED_ENTRIES_MISMATCH" {
		t.Fatalf("result=%+v", result)
	}
}

func TestValidateRejectsUnsafeGeneratedEntry(t *testing.T) {
	root := minimalPack(t)
	archive := filepath.Join(root, "generated", "srcjars", "demo.srcjar")
	writeZip(t, archive, "../escape.java")
	manifest := []map[string]any{{"path": "generated/srcjars/demo.srcjar", "entries": []string{"../escape.java"}}}
	writeJSON(t, filepath.Join(root, "generated", "manifest.json"), manifest)
	refreshChecksums(t, root)
	result := Validate(root)
	if result.Valid || result.Diagnostics[0].Code != "PACK_GENERATED_ARCHIVE_INVALID" {
		t.Fatalf("result=%+v", result)
	}
}

func TestValidateRejectsInvalidJavaJar(t *testing.T) {
	root := minimalPack(t)
	writeFile(t, filepath.Join(root, "out", "soong", "framework.jar"), []byte("not-a-jar"))
	refreshChecksums(t, root)
	result := Validate(root)
	if result.Valid || result.Diagnostics[0].Code != "PACK_JAVA_JAR_INVALID" {
		t.Fatalf("result=%+v", result)
	}
}

func minimalPack(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	files := map[string]string{
		"manifest.json":           `{"schemaVersion":1,"packs":["core"]}`,
		"generated/manifest.json": `[]`, "modules/catalog.json": `{"modules":{}}`,
		"modules/dependency-graph.json": `{}`, "aidl/symbols.json": `{}`,
		"aidl/index.json": `{"documents":{},"symbols":{}}`, "aidl/modules.json": `[]`,
		"artifacts/catalog.json": `{"artifacts":{}}`,
	}
	for name, data := range files {
		writeFile(t, filepath.Join(root, name), []byte(data))
	}
	refreshChecksums(t, root)
	return root
}

func writeFile(t *testing.T, pathName string, data []byte) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(pathName), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(pathName, data, 0o644); err != nil {
		t.Fatal(err)
	}
}
func writeJSON(t *testing.T, pathName string, value any) {
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	writeFile(t, pathName, data)
}
func writeZip(t *testing.T, pathName, entry string) {
	if err := os.MkdirAll(filepath.Dir(pathName), 0o755); err != nil {
		t.Fatal(err)
	}
	file, err := os.Create(pathName)
	if err != nil {
		t.Fatal(err)
	}
	writer := zip.NewWriter(file)
	item, err := writer.Create(entry)
	if err != nil {
		t.Fatal(err)
	}
	_, _ = item.Write([]byte("class Demo {}"))
	_ = writer.Close()
	_ = file.Close()
}
func refreshChecksums(t *testing.T, root string) {
	t.Helper()
	names := []string{}
	_ = filepath.WalkDir(root, func(pathName string, entry os.DirEntry, err error) error {
		if err != nil || entry.IsDir() || entry.Name() == "checksums.sha256" {
			return err
		}
		relative, _ := filepath.Rel(root, pathName)
		names = append(names, filepath.ToSlash(relative))
		return nil
	})
	sort.Strings(names)
	output := ""
	for _, name := range names {
		data, _ := os.ReadFile(filepath.Join(root, filepath.FromSlash(name)))
		sum := sha256.Sum256(data)
		output += hex.EncodeToString(sum[:]) + "  " + name + "\n"
	}
	writeFile(t, filepath.Join(root, "checksums.sha256"), []byte(output))
}

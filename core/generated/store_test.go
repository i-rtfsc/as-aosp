package generated

import (
	"archive/zip"
	"os"
	"path/filepath"
	"testing"
)

func TestExtractSrcJarCachesByContent(t *testing.T) {
	root := t.TempDir()
	archivePath := filepath.Join(root, "framework.srcjar")
	writeZip(t, archivePath, map[string]string{
		"android/demo/IFoo.java": "package android.demo;",
	})
	first := ExtractSrcJar(archivePath, filepath.Join(root, "cache"), DefaultLimits())
	second := ExtractSrcJar(archivePath, filepath.Join(root, "cache"), DefaultLimits())
	if first.Archive == nil || second.Archive == nil {
		t.Fatalf("reports: %+v %+v", first, second)
	}
	if first.Archive.Root != second.Archive.Root {
		t.Fatalf("roots: %q %q", first.Archive.Root, second.Archive.Root)
	}
	entry, ok := first.Archive.Entry("android/demo/IFoo.java")
	if !ok {
		t.Fatal("entry not found")
	}
	if _, err := os.Stat(string(entry)); err != nil {
		t.Fatal(err)
	}
}

func TestExtractSrcJarRejectsEscapingEntry(t *testing.T) {
	root := t.TempDir()
	archivePath := filepath.Join(root, "unsafe.srcjar")
	writeZip(t, archivePath, map[string]string{
		"../outside.java": "unsafe",
	})
	report := ExtractSrcJar(
		archivePath,
		filepath.Join(root, "cache"),
		DefaultLimits(),
	)
	if len(report.Diagnostics) != 1 ||
		report.Diagnostics[0].Code != "GENERATED_ENTRY_UNSAFE" {
		t.Fatalf("report: %+v", report)
	}
	if _, err := os.Stat(filepath.Join(root, "outside.java")); !os.IsNotExist(err) {
		t.Fatalf("outside file exists: %v", err)
	}
}

func TestExtractSrcJarEnforcesFileLimit(t *testing.T) {
	root := t.TempDir()
	archivePath := filepath.Join(root, "large.srcjar")
	writeZip(t, archivePath, map[string]string{
		"android/demo/IFoo.java": "12345",
	})
	report := ExtractSrcJar(
		archivePath,
		filepath.Join(root, "cache"),
		Limits{MaxEntries: 10, MaxFileBytes: 4, MaxTotalBytes: 10},
	)
	if len(report.Diagnostics) != 1 ||
		report.Diagnostics[0].Code != "GENERATED_ARCHIVE_LIMIT" {
		t.Fatalf("report: %+v", report)
	}
}

func writeZip(
	t *testing.T,
	archivePath string,
	entries map[string]string,
) {
	t.Helper()
	file, err := os.Create(archivePath)
	if err != nil {
		t.Fatal(err)
	}
	writer := zip.NewWriter(file)
	for name, content := range entries {
		entry, err := writer.Create(name)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := entry.Write([]byte(content)); err != nil {
			t.Fatal(err)
		}
	}
	if err := writer.Close(); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}
}

package topology

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestDiscoverInfersSymlinkRepoPrefix(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("symlink fixture requires platform permission on Windows")
	}
	root := t.TempDir()
	canonical := filepath.Join(root, "full")
	view := filepath.Join(root, "view")
	if err := os.MkdirAll(filepath.Join(canonical, "frameworks", "base"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(view, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(canonical, "frameworks", "base"), filepath.Join(view, "frameworks-base")); err != nil {
		t.Fatal(err)
	}
	topology, diagnostics, err := Discover(DiscoverOptions{WorkspaceViewRoot: view, CanonicalRepoRoot: canonical, MaxDepth: 1})
	if err != nil {
		t.Fatal(err)
	}
	if len(diagnostics) != 0 {
		t.Fatalf("diagnostics: %+v", diagnostics)
	}
	if len(topology.SourceMounts) != 1 || topology.SourceMounts[0].RepoPrefix != "frameworks/base" {
		t.Fatalf("mounts: %+v", topology.SourceMounts)
	}
}

func TestDiscoverReportsDanglingSymlink(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("symlink fixture requires platform permission on Windows")
	}
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "view"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(root, "missing"), filepath.Join(root, "view", "broken")); err != nil {
		t.Fatal(err)
	}
	_, diagnostics, err := Discover(DiscoverOptions{WorkspaceViewRoot: filepath.Join(root, "view"), MaxDepth: 1})
	if err != nil {
		t.Fatal(err)
	}
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == "TOPOLOGY_DANGLING_LINK" {
			return
		}
	}
	t.Fatalf("expected dangling-link diagnostic: %+v", diagnostics)
}

func TestDiscoverReportsDuplicateTargetWithoutCycle(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("symlink fixture requires platform permission on Windows")
	}
	root := t.TempDir()
	canonical := filepath.Join(root, "full")
	view := filepath.Join(root, "view")
	if err := os.MkdirAll(filepath.Join(canonical, "frameworks", "base"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(view, 0o755); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"base-a", "base-b"} {
		if err := os.Symlink(filepath.Join(canonical, "frameworks", "base"), filepath.Join(view, name)); err != nil {
			t.Fatal(err)
		}
	}
	_, diagnostics, err := Discover(DiscoverOptions{WorkspaceViewRoot: view, CanonicalRepoRoot: canonical, MaxDepth: 2})
	if err != nil {
		t.Fatal(err)
	}
	for _, diagnostic := range diagnostics {
		if diagnostic.Code == "TOPOLOGY_DUPLICATE_TARGET" {
			return
		}
		if diagnostic.Code == "TOPOLOGY_LINK_CYCLE" {
			t.Fatalf("duplicate target was classified as a cycle: %+v", diagnostics)
		}
	}
	t.Fatalf("expected duplicate-target diagnostic: %+v", diagnostics)
}

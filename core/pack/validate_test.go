package pack

import (
	"os"
	"path/filepath"
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

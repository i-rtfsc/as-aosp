package modules

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadPackCatalog(t *testing.T) {
	root := t.TempDir()
	pathName := filepath.Join(root, "modules", "catalog.json")
	if err := os.MkdirAll(filepath.Dir(pathName), 0o755); err != nil {
		t.Fatal(err)
	}
	data, _ := json.Marshal(Catalog{Modules: map[string]Module{"framework": {Name: "framework"}}})
	if err := os.WriteFile(pathName, data, 0o644); err != nil {
		t.Fatal(err)
	}
	catalog, err := LoadPackCatalog(root)
	if err != nil || catalog.Modules["framework"].Name != "framework" {
		t.Fatalf("catalog=%+v err=%v", catalog, err)
	}
}

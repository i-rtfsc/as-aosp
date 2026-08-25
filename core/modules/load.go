package modules

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func LoadPackCatalog(root string) (Catalog, error) {
	data, err := os.ReadFile(filepath.Join(root, "modules", "catalog.json"))
	if err != nil {
		return Catalog{}, err
	}
	var catalog Catalog
	if err := json.Unmarshal(data, &catalog); err != nil {
		return Catalog{}, fmt.Errorf("parse module catalog: %w", err)
	}
	if catalog.Modules == nil {
		return Catalog{}, fmt.Errorf("module catalog has no modules")
	}
	return catalog, nil
}

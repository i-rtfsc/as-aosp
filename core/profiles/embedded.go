package profiles

import (
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

// The readable files under profiles/ are the source-of-truth. These copies are
// generated into this package so released engines do not depend on repo paths.
//
//go:generate sh -c "mkdir -p data && cp ../../profiles/common.json data/common.json && cp ../../profiles/android-17.json data/android-17.json"
//go:embed data/common.json data/android-17.json
var embeddedCatalogFS embed.FS

func BuiltInCatalog(release string) (Catalog, error) {
	if strings.TrimSpace(release) == "" {
		release = "android-17"
	}
	commonData, err := embeddedCatalogFS.ReadFile("data/common.json")
	if err != nil {
		return Catalog{}, fmt.Errorf("read embedded common profile: %w", err)
	}
	var catalog Catalog
	if err := json.Unmarshal(commonData, &catalog); err != nil {
		return Catalog{}, fmt.Errorf("decode embedded common profile: %w", err)
	}
	if release == catalog.Baseline.AndroidRelease {
		deltaData, err := embeddedCatalogFS.ReadFile("data/" + release + ".json")
		if err != nil {
			return Catalog{}, fmt.Errorf("read embedded release profile %q: %w", release, err)
		}
		var delta ReleaseDelta
		if err := json.Unmarshal(deltaData, &delta); err != nil {
			return Catalog{}, fmt.Errorf("decode embedded release profile %q: %w", release, err)
		}
		return ApplyDelta(catalog, delta), nil
	}
	return Catalog{}, fmt.Errorf("unsupported built-in Android release %q", release)
}

func RuntimeCatalog(release, tag string) Catalog {
	if strings.TrimSpace(release) == "" {
		release = "android-17"
	}
	if strings.TrimSpace(tag) == "" {
		tag = release
	}
	catalog, err := BuiltInCatalog(release)
	if err != nil {
		return DefaultCatalog(release, tag)
	}
	catalog.Baseline.AndroidTag = tag
	return catalog
}

package sourcecatalog

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

const cacheSchema = 5

type cacheEnvelope struct {
	Schema       int               `json:"schema"`
	Fingerprints map[string]string `json:"fingerprints"`
	Catalog      Catalog           `json:"catalog"`
}

func Discover(root, cacheRoot string) Report {
	report := Report{Catalog: Catalog{Modules: []Module{}}, Diagnostics: []model.Diagnostic{}}
	absolute, err := filepath.Abs(strings.TrimSpace(root))
	if err != nil {
		report.Diagnostics = append(report.Diagnostics, diagnostic("SOURCE_ROOT_INVALID", err.Error(), root))
		return report
	}
	info, err := os.Stat(absolute)
	if err != nil || !info.IsDir() {
		message := "source root is not a directory"
		if err != nil {
			message = err.Error()
		}
		report.Diagnostics = append(report.Diagnostics, diagnostic("SOURCE_ROOT_INVALID", message, absolute))
		return report
	}
	files, fingerprints, diagnostics := discoverFiles(absolute)
	report.FilesScanned = len(files)
	report.Diagnostics = append(report.Diagnostics, diagnostics...)
	if cacheRoot != "" {
		if cached, ok := loadCache(cacheRoot, fingerprints); ok {
			cached.CacheHit = true
			cached.FilesScanned = len(files)
			cached.Diagnostics = append(cached.Diagnostics, diagnostics...)
			return cached
		}
	}
	modules := make([]Module, 0)
	for _, file := range files {
		data, readErr := os.ReadFile(filepath.Join(absolute, filepath.FromSlash(file)))
		if readErr != nil {
			report.Diagnostics = append(report.Diagnostics, diagnostic("SOURCE_FILE_READ_FAILED", readErr.Error(), file))
			continue
		}
		var parsed []Module
		var parseErr error
		if strings.HasSuffix(strings.ToLower(file), ".mk") {
			parsed = parseMakefile(file, string(data))
		} else {
			parsed, parseErr = parseBlueprint(file, string(data))
		}
		if parseErr != nil {
			report.Diagnostics = append(report.Diagnostics, diagnostic("ANDROID_BP_PARSE_FAILED", parseErr.Error(), file))
			continue
		}
		modules = append(modules, parsed...)
	}
	sort.Slice(modules, func(i, j int) bool {
		if modules[i].Name != modules[j].Name {
			return modules[i].Name < modules[j].Name
		}
		return modules[i].Blueprint < modules[j].Blueprint
	})
	for i := range modules {
		modules[i] = classifySourceModule(modules[i])
		modules[i].Recommended = isRecommended(modules[i])
	}
	report.Catalog = Catalog{Modules: modules, Fingerprints: fingerprints}
	if cacheRoot != "" {
		saveCache(cacheRoot, report.Catalog)
	}
	return report
}

func discoverFiles(root string) ([]string, map[string]string, []model.Diagnostic) {
	files := []string{}
	fingerprints := map[string]string{}
	diagnostics := []model.Diagnostic{}
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			diagnostics = append(diagnostics, diagnostic("SOURCE_SCAN_FAILED", walkErr.Error(), path))
			return nil
		}
		if entry.IsDir() {
			if path != root && excluded(filepath.ToSlash(path[len(root):])) {
				return filepath.SkipDir
			}
			return nil
		}
		if entry.Name() != "Android.bp" && entry.Name() != "Android.mk" {
			return nil
		}
		rel, err := filepath.Rel(root, path)
		if err != nil {
			return nil
		}
		rel = filepath.ToSlash(rel)
		info, err := entry.Info()
		if err != nil {
			diagnostics = append(diagnostics, diagnostic("SOURCE_FILE_STAT_FAILED", err.Error(), rel))
			return nil
		}
		fingerprints[rel] = fmt.Sprintf("%d:%d", info.Size(), info.ModTime().UnixNano())
		files = append(files, rel)
		return nil
	})
	if err != nil {
		diagnostics = append(diagnostics, diagnostic("SOURCE_SCAN_FAILED", err.Error(), root))
	}
	sort.Strings(files)
	return files, fingerprints, diagnostics
}

func excluded(relative string) bool {
	for _, part := range strings.Split(strings.Trim(relative, "/"), "/") {
		if part == ".git" || part == ".repo" || part == "out" || part == ".aosp-ide" {
			return true
		}
	}
	return false
}

func diagnostic(code, message, path string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: "source-catalog", Message: message, LocalPath: path}
}

func loadCache(root string, fingerprints map[string]string) (Report, bool) {
	data, err := os.ReadFile(filepath.Join(root, "source-catalog", "catalog.json"))
	if err != nil {
		return Report{}, false
	}
	var envelope cacheEnvelope
	if json.Unmarshal(data, &envelope) != nil || envelope.Schema != cacheSchema || !sameFingerprints(envelope.Fingerprints, fingerprints) {
		return Report{}, false
	}
	return Report{Catalog: envelope.Catalog, Diagnostics: []model.Diagnostic{}}, true
}

func saveCache(root string, catalog Catalog) {
	dir := filepath.Join(root, "source-catalog")
	if os.MkdirAll(dir, 0o755) != nil {
		return
	}
	envelope := cacheEnvelope{Schema: cacheSchema, Fingerprints: catalog.Fingerprints, Catalog: catalog}
	data, err := json.MarshalIndent(envelope, "", "  ")
	if err != nil {
		return
	}
	temporary := filepath.Join(dir, "catalog.json.tmp")
	if os.WriteFile(temporary, append(data, '\n'), 0o644) == nil {
		_ = os.Rename(temporary, filepath.Join(dir, "catalog.json"))
	}
}

func sameFingerprints(a, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for key, value := range b {
		if a[key] != value {
			return false
		}
	}
	return true
}

func isRecommended(module Module) bool {
	name := strings.ToLower(module.Name)
	switch name {
	case "framework-minus-apex", "services", "systemui", "framework-res",
		"libbinder", "libbase", "libutils", "surfaceflinger":
		return true
	default:
		return false
	}
}

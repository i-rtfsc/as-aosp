package main

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/artifacts"
	"github.com/i-rtfsc/as-aosp/core/jni"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	coremodules "github.com/i-rtfsc/as-aosp/core/modules"
	corepack "github.com/i-rtfsc/as-aosp/core/pack"
)

type packManifest struct {
	SchemaVersion int      `json:"schemaVersion"`
	AOSPRoot      string   `json:"aospRoot,omitempty"`
	OutRoot       string   `json:"outRoot,omitempty"`
	Packs         []string `json:"packs"`
	Files         int      `json:"files"`
	Bytes         int64    `json:"bytes"`
	Modules       []string `json:"modules,omitempty"`
}

func runPack(args []string) {
	if len(args) == 0 || (args[0] != "export" && args[0] != "validate" && args[0] != "inspect" && args[0] != "preflight") {
		fatal(errors.New("usage: aosp-index pack <export|validate|inspect|preflight> ..."))
	}
	if args[0] == "validate" {
		runPackValidate(args[1:])
		return
	}
	if args[0] == "inspect" {
		runPackInspect(args[1:])
		return
	}
	if args[0] == "preflight" {
		runPackPreflight(args[1:])
		return
	}
	flags := flag.NewFlagSet("pack export", flag.ExitOnError)
	aospRoot := flags.String("aosp-root", "", "AOSP source root")
	outRoot := flags.String("out-root", "", "AOSP out root")
	destination := flags.String("destination", "", "Metadata pack destination")
	native := flags.Bool("native", false, "include native metadata and compdb")
	modules := stringFlags{}
	flags.Var(&modules, "module", "selected module (repeatable)")
	_ = flags.Parse(args[1:])
	if *aospRoot == "" || *outRoot == "" || *destination == "" {
		fatal(errors.New("--aosp-root, --out-root and --destination are required"))
	}
	if err := exportPack(*aospRoot, *outRoot, *destination, *native, []string(modules)); err != nil {
		fatal(err)
	}
}

func runPackPreflight(args []string) {
	flags := flag.NewFlagSet("pack preflight", flag.ExitOnError)
	outRoot := flags.String("out-root", "", "AOSP out root")
	var requestedModules stringFlags
	flags.Var(&requestedModules, "module", "Soong module to verify (repeatable)")
	_ = flags.Parse(args)
	if *outRoot == "" {
		fatal(errors.New("--out-root is required"))
	}
	outRootAbs, err := filepath.Abs(*outRoot)
	if err != nil {
		fatal(err)
	}
	checks := []map[string]any{}
	for _, item := range []struct {
		name       string
		candidates []string
	}{
		{"module-info", []string{"module-info.json", filepath.Join("soong", "module-info.json")}},
		{"java-deps", []string{"module_bp_java_deps.json", filepath.Join("soong", "module_bp_java_deps.json")}},
		{"cc-deps", []string{"module_bp_cc_deps.json", filepath.Join("soong", "module_bp_cc_deps.json")}},
		{"compdb", []string{filepath.Join("development", "ide", "compdb", "compile_commands.json"), filepath.Join("soong", "development", "ide", "compdb", "compile_commands.json")}},
	} {
		found := ""
		for _, candidate := range item.candidates {
			pathName := filepath.Join(outRootAbs, candidate)
			if info, statErr := os.Stat(pathName); statErr == nil && !info.IsDir() {
				found = pathName
				break
			}
		}
		checks = append(checks, map[string]any{"name": item.name, "found": found != "", "path": found})
	}
	report := metadata.LoadDirectory(outRootAbs)
	catalog := coremodules.Build(report.Snapshot)
	artifactCatalog := artifacts.Build(catalog)
	moduleChecks := make([]map[string]any, 0, len(requestedModules))
	for _, name := range requestedModules {
		module, found := catalog.Modules[name]
		artifactNames := make([]string, 0)
		for _, artifact := range artifactCatalog.Artifacts {
			for _, owner := range artifact.OwnerModules {
				if owner == name {
					artifactNames = append(artifactNames, artifact.Name)
					break
				}
			}
		}
		sort.Strings(artifactNames)
		moduleChecks = append(moduleChecks, map[string]any{
			"name": name, "found": found, "sourceRoots": len(module.SourceRoots),
			"generated": len(module.Generated), "jars": len(module.Jars),
			"srcjars": len(module.SrcJars), "artifacts": artifactNames,
		})
	}
	srcjars, jars, apks := 0, 0, 0
	_ = filepath.WalkDir(outRootAbs, func(pathName string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil || entry.IsDir() {
			return nil
		}
		switch strings.ToLower(filepath.Ext(entry.Name())) {
		case ".srcjar":
			srcjars++
		case ".jar":
			jars++
		case ".apk":
			apks++
		}
		return nil
	})
	writeJSON(map[string]any{
		"outRoot":     outRootAbs,
		"checks":      checks,
		"srcjars":     srcjars,
		"jars":        jars,
		"apks":        apks,
		"modules":     moduleChecks,
		"diagnostics": report.Diagnostics,
		"suggestions": []string{"If compdb is missing, run SOONG_GEN_COMPDB=1 m nothing."},
	})
}

func runPackInspect(args []string) {
	flags := flag.NewFlagSet("pack inspect", flag.ExitOnError)
	root := flags.String("root", "", "Metadata pack root")
	moduleName := flags.String("module", "", "module name")
	_ = flags.Parse(args)
	if *root == "" {
		fatal(errors.New("--root is required"))
	}
	catalog, err := coremodules.LoadPackCatalog(*root)
	if err != nil {
		fatal(err)
	}
	output := map[string]any{"modules": len(catalog.Modules)}
	jarFiles, jarBytes := packArtifactStats(*root, filepath.Join("out"), ".jar")
	srcjarFiles, srcjarBytes := packArtifactStats(*root, filepath.Join("generated", "srcjars"), ".srcjar")
	output["jars"] = map[string]any{"files": jarFiles, "bytes": jarBytes}
	output["srcjars"] = map[string]any{"files": srcjarFiles, "bytes": srcjarBytes}
	if *moduleName != "" {
		module, ok := catalog.Modules[*moduleName]
		if !ok {
			fatal(fmt.Errorf("module not found: %s", *moduleName))
		}
		output["module"] = module
		output["closure"] = catalog.Closure([]string{*moduleName})
	}
	writeJSON(output)
}

func packArtifactStats(root, relativeRoot, extension string) (int, int64) {
	files := 0
	var bytes int64
	_ = filepath.WalkDir(filepath.Join(root, relativeRoot), func(pathName string, entry os.DirEntry, err error) error {
		if err != nil || entry.IsDir() || strings.ToLower(filepath.Ext(entry.Name())) != extension {
			return nil
		}
		if info, infoErr := entry.Info(); infoErr == nil {
			files++
			bytes += info.Size()
		}
		return nil
	})
	return files, bytes
}

type stringFlags []string

func (s *stringFlags) String() string         { return strings.Join(*s, ",") }
func (s *stringFlags) Set(value string) error { *s = append(*s, value); return nil }

func exportPack(aospRoot, outRoot, destination string, native bool, modules []string) error {
	var err error
	aospRoot, err = filepath.Abs(aospRoot)
	if err != nil {
		return fmt.Errorf("resolve AOSP root: %w", err)
	}
	outRoot, err = filepath.Abs(outRoot)
	if err != nil {
		return fmt.Errorf("resolve out root: %w", err)
	}
	destination, err = filepath.Abs(destination)
	if err != nil {
		return fmt.Errorf("resolve pack destination: %w", err)
	}
	tmp, err := os.MkdirTemp(filepath.Dir(destination), ".aosp-pack-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)
	if err := os.MkdirAll(filepath.Join(tmp, "generated", "srcjars"), 0o755); err != nil {
		return err
	}
	manifest := packManifest{SchemaVersion: 1, AOSPRoot: aospRoot, OutRoot: outRoot, Packs: []string{"core"}, Modules: modules}
	coreFiles := []string{"module-info.json", "module_bp_java_deps.json"}
	if native {
		manifest.Packs = append(manifest.Packs, "native")
		coreFiles = append(coreFiles, "module_bp_cc_deps.json", filepath.Join("development", "ide", "compdb", "compile_commands.json"))
	}
	for _, relative := range coreFiles {
		source, ok := findPackMetadata(outRoot, relative)
		if !ok {
			continue
		}
		target := filepath.Join(tmp, relative)
		if err := copyFile(source, target); err != nil {
			return err
		}
	}
	if err := writePackCatalog(tmp); err != nil {
		return err
	}
	if err := writeArtifactCatalog(tmp); err != nil {
		return err
	}
	if err := copyPackLibraries(tmp, outRoot, modules); err != nil {
		return err
	}
	selectedPrefixes, err := selectedRepoPrefixes(tmp, modules)
	if err != nil {
		return err
	}
	if err := writeAIDLPack(aospRoot, tmp, selectedPrefixes); err != nil {
		return err
	}
	if native {
		if err := writeJNIPack(aospRoot, tmp, selectedPrefixes); err != nil {
			return err
		}
	}
	selectedArchives, err := selectedSrcJars(tmp, modules)
	if err != nil {
		return err
	}
	generatedManifest := make([]map[string]any, 0)
	err = filepath.WalkDir(filepath.Join(outRoot, "soong", ".intermediates"), func(pathName string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return nil
		}
		if entry.IsDir() || strings.ToLower(filepath.Ext(entry.Name())) != ".srcjar" {
			return nil
		}
		relative, relErr := filepath.Rel(outRoot, pathName)
		if relErr != nil {
			return relErr
		}
		intermediateRelative, relErr := filepath.Rel(filepath.Join(outRoot, "soong", ".intermediates"), pathName)
		if relErr != nil {
			return relErr
		}
		if len(modules) > 0 && !selectedArchives[filepath.ToSlash(intermediateRelative)] {
			return nil
		}
		packPath := filepath.Join("generated", "srcjars", intermediateRelative)
		target := filepath.Join(tmp, packPath)
		if err := copyFile(pathName, target); err != nil {
			return err
		}
		entries, zipErr := srcJarEntries(pathName)
		if zipErr != nil {
			return zipErr
		}
		generatedManifest = append(generatedManifest, map[string]any{"source": filepath.ToSlash(relative), "path": filepath.ToSlash(packPath), "entries": entries})
		return nil
	})
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	manifestData, _ := json.MarshalIndent(manifest, "", "  ")
	if err := os.WriteFile(filepath.Join(tmp, "manifest.json"), append(manifestData, '\n'), 0o644); err != nil {
		return err
	}
	generatedData, _ := json.MarshalIndent(generatedManifest, "", "  ")
	if err := os.WriteFile(filepath.Join(tmp, "generated", "manifest.json"), append(generatedData, '\n'), 0o644); err != nil {
		return err
	}
	files, err := packFiles(tmp)
	if err != nil {
		return err
	}
	manifest.Files = len(files)
	for _, file := range files {
		if file == "manifest.json" {
			continue
		}
		info, _ := os.Stat(filepath.Join(tmp, file))
		if info != nil {
			manifest.Bytes += info.Size()
		}
	}
	manifestData, _ = json.MarshalIndent(manifest, "", "  ")
	if err := os.WriteFile(filepath.Join(tmp, "manifest.json"), append(manifestData, '\n'), 0o644); err != nil {
		return err
	}
	checksums := make([]string, 0, len(files))
	for _, file := range files {
		sum, err := fileSHA256(filepath.Join(tmp, file))
		if err != nil {
			return err
		}
		checksums = append(checksums, sum+"  "+file)
	}
	sort.Strings(checksums)
	if err := os.WriteFile(filepath.Join(tmp, "checksums.sha256"), []byte(strings.Join(checksums, "\n")+"\n"), 0o644); err != nil {
		return err
	}
	if result := corepack.Validate(tmp); !result.Valid {
		if len(result.Diagnostics) == 0 {
			return fmt.Errorf("exported metadata pack failed validation")
		}
		diagnostic := result.Diagnostics[0]
		return fmt.Errorf("exported pack invalid: %s: %s: %s", diagnostic.Code, diagnostic.Path, diagnostic.Message)
	}
	_ = os.RemoveAll(destination)
	return os.Rename(tmp, destination)
}

func writeArtifactCatalog(root string) error {
	catalog, err := coremodules.LoadPackCatalog(root)
	if err != nil {
		return err
	}
	artifactCatalog := artifacts.Build(catalog)
	if err := os.MkdirAll(filepath.Join(root, "artifacts"), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(artifactCatalog, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(root, "artifacts", "catalog.json"), append(data, '\n'), 0o644)
}

func copyPackLibraries(root, outRoot string, names []string) error {
	catalog, err := coremodules.LoadPackCatalog(root)
	if err != nil {
		return err
	}
	selected := make([]string, 0)
	if len(names) == 0 {
		for name := range catalog.Modules {
			selected = append(selected, name)
		}
	} else {
		selected = catalog.Closure(names)
	}
	seen := map[string]bool{}
	for _, name := range selected {
		for _, ref := range catalog.Modules[name].Jars {
			source, destination, ok := resolvePackArtifact(outRoot, ref.Raw)
			if !ok || seen[destination] {
				continue
			}
			if err := copyFile(source, filepath.Join(root, destination)); err != nil {
				return err
			}
			seen[destination] = true
		}
	}
	return nil
}

func resolvePackArtifact(outRoot, raw string) (string, string, bool) {
	normalized := filepath.ToSlash(strings.TrimSpace(raw))
	if normalized == "" || strings.Contains(normalized, "..") {
		return "", "", false
	}
	if filepath.IsAbs(raw) {
		base := filepath.ToSlash(filepath.Clean(outRoot))
		if !strings.HasPrefix(normalized, base+"/") {
			return "", "", false
		}
		suffix := strings.TrimPrefix(normalized, base+"/")
		if !strings.HasSuffix(strings.ToLower(suffix), ".jar") {
			return "", "", false
		}
		if _, err := os.Stat(raw); err != nil {
			return "", "", false
		}
		return raw, filepath.ToSlash(filepath.Join("out", filepath.FromSlash(suffix))), true
	}
	if !strings.HasSuffix(strings.ToLower(normalized), ".jar") {
		return "", "", false
	}
	candidates := []struct{ source, suffix string }{{filepath.Join(outRoot, filepath.FromSlash(normalized)), normalized}}
	outName := filepath.Base(outRoot)
	if strings.HasPrefix(normalized, outName+"/") {
		suffix := strings.TrimPrefix(normalized, outName+"/")
		candidates = append(candidates, struct{ source, suffix string }{filepath.Join(outRoot, filepath.FromSlash(suffix)), suffix})
	}
	for _, candidate := range candidates {
		if _, err := os.Stat(candidate.source); err == nil {
			return candidate.source, filepath.ToSlash(filepath.Join("out", filepath.FromSlash(candidate.suffix))), true
		}
	}
	return "", "", false
}

func selectedSrcJars(root string, names []string) (map[string]bool, error) {
	result := make(map[string]bool)
	if len(names) == 0 {
		return result, nil
	}
	catalog, err := coremodules.LoadPackCatalog(root)
	if err != nil {
		return nil, err
	}
	for _, name := range names {
		if _, ok := catalog.Modules[name]; !ok {
			return nil, fmt.Errorf("selected module not found: %s", name)
		}
	}
	for _, name := range catalog.Closure(names) {
		for _, ref := range append(catalog.Modules[name].SrcJars, catalog.Modules[name].Generated...) {
			raw := filepath.ToSlash(ref.Raw)
			if marker := strings.Index(raw, ".intermediates/"); marker >= 0 {
				result[strings.TrimPrefix(raw[marker+len(".intermediates/"):], "/")] = true
			}
		}
	}
	return result, nil
}

func selectedRepoPrefixes(root string, names []string) ([]string, error) {
	if len(names) == 0 {
		return nil, nil
	}
	catalog, err := coremodules.LoadPackCatalog(root)
	if err != nil {
		return nil, err
	}
	for _, name := range names {
		if _, ok := catalog.Modules[name]; !ok {
			return nil, fmt.Errorf("selected module not found: %s", name)
		}
	}
	prefixes := make(map[string]bool)
	for _, name := range catalog.Closure(names) {
		module := catalog.Modules[name]
		for _, ref := range append(module.SourceRoots, module.Paths...) {
			if ref.RepoPath != "" {
				prefixes[string(ref.RepoPath)] = true
			}
		}
	}
	result := make([]string, 0, len(prefixes))
	for prefix := range prefixes {
		result = append(result, prefix)
	}
	sort.Strings(result)
	return result, nil
}

func writeJNIPack(aospRoot, destination string, prefixes []string) error {
	javaFiles, cppFiles := map[model.RepoPath]string{}, map[model.RepoPath]string{}
	err := filepath.WalkDir(aospRoot, func(pathName string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil || entry.IsDir() || strings.HasPrefix(filepath.ToSlash(pathName), filepath.ToSlash(filepath.Join(aospRoot, ".repo"))) || strings.Contains(filepath.ToSlash(pathName), "/.git/") || strings.Contains(filepath.ToSlash(pathName), "/out/") {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if ext != ".java" && ext != ".kt" && ext != ".c" && ext != ".cc" && ext != ".cpp" && ext != ".cxx" && ext != ".h" && ext != ".hh" && ext != ".hpp" {
			return nil
		}
		relative, err := filepath.Rel(aospRoot, pathName)
		if err != nil {
			return err
		}
		repoPath, err := model.NormalizeRepoPath(filepath.ToSlash(relative))
		if err != nil {
			return nil
		}
		if !repoPathAllowed(repoPath, prefixes) {
			return nil
		}
		data, err := os.ReadFile(pathName)
		if err != nil {
			return nil
		}
		if ext == ".java" || ext == ".kt" {
			javaFiles[repoPath] = string(data)
		} else {
			cppFiles[repoPath] = string(data)
		}
		return nil
	})
	if err != nil {
		return err
	}
	index := jni.BuildIndex(javaFiles, cppFiles)
	if err := os.MkdirAll(filepath.Join(destination, "jni"), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(destination, "jni", "index.json"), append(data, '\n'), 0o644)
}

func writeAIDLPack(aospRoot, destination string, prefixes []string) error {
	index := aidl.NewIndex()
	err := filepath.WalkDir(aospRoot, func(pathName string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil || entry.IsDir() {
			return nil
		}
		if strings.ToLower(filepath.Ext(entry.Name())) != ".aidl" {
			return nil
		}
		relative, err := filepath.Rel(aospRoot, pathName)
		if err != nil {
			return err
		}
		repoPath, err := model.NormalizeRepoPath(filepath.ToSlash(relative))
		if err != nil {
			return nil
		}
		if !repoPathAllowed(repoPath, prefixes) {
			return nil
		}
		data, err := os.ReadFile(pathName)
		if err != nil {
			return nil
		}
		index.Add(aidl.Parse(repoPath, string(data)))
		return nil
	})
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(destination, "aidl"), 0o755); err != nil {
		return err
	}
	symbols, err := json.MarshalIndent(index.Symbols, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(destination, "aidl", "symbols.json"), append(symbols, '\n'), 0o644); err != nil {
		return err
	}
	indexData, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(destination, "aidl", "index.json"), append(indexData, '\n'), 0o644); err != nil {
		return err
	}
	documents := make([]string, 0, len(index.Documents))
	for path := range index.Documents {
		documents = append(documents, string(path))
	}
	sort.Strings(documents)
	data, err := json.MarshalIndent(documents, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(destination, "aidl", "modules.json"), append(data, '\n'), 0o644)
}

func repoPathAllowed(path model.RepoPath, prefixes []string) bool {
	if len(prefixes) == 0 {
		return true
	}
	value := string(path)
	for _, prefix := range prefixes {
		if value == prefix || strings.HasPrefix(value, prefix+"/") {
			return true
		}
	}
	return false
}

func writePackCatalog(root string) error {
	report := metadata.LoadDirectory(root)
	catalog := coremodules.Build(report.Snapshot)
	if err := os.MkdirAll(filepath.Join(root, "modules"), 0o755); err != nil {
		return err
	}
	data, err := json.MarshalIndent(catalog, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(root, "modules", "catalog.json"), append(data, '\n'), 0o644); err != nil {
		return err
	}
	graph := make(map[string][]string, len(catalog.Modules))
	for name, module := range catalog.Modules {
		graph[name] = module.Dependencies
	}
	data, err = json.MarshalIndent(graph, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filepath.Join(root, "modules", "dependency-graph.json"), append(data, '\n'), 0o644)
}

func runPackValidate(args []string) {
	flags := flag.NewFlagSet("pack validate", flag.ExitOnError)
	root := flags.String("root", "", "Metadata pack root")
	_ = flags.Parse(args)
	if *root == "" {
		fatal(errors.New("--root is required"))
	}
	if err := validatePack(*root); err != nil {
		fatal(err)
	}
}

func validatePack(root string) error {
	result := corepack.Validate(root)
	if result.Valid {
		return nil
	}
	if len(result.Diagnostics) == 0 {
		return fmt.Errorf("metadata pack validation failed")
	}
	diagnostic := result.Diagnostics[0]
	return fmt.Errorf("%s: %s: %s", diagnostic.Code, diagnostic.Path, diagnostic.Message)
}

func findPackMetadata(outRoot, relative string) (string, bool) {
	candidates := []string{filepath.Join(outRoot, relative), filepath.Join(outRoot, "soong", relative)}
	if relative == "module-info.json" {
		matches, _ := filepath.Glob(filepath.Join(outRoot, "target", "product", "*", "module-info.json"))
		candidates = append(candidates, matches...)
	}
	for _, candidate := range candidates {
		if info, err := os.Stat(candidate); err == nil && !info.IsDir() {
			return candidate, true
		}
	}
	return "", false
}

func srcJarEntries(pathName string) ([]string, error) {
	reader, err := zip.OpenReader(pathName)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	entries := make([]string, 0, len(reader.File))
	for _, file := range reader.File {
		name := strings.TrimPrefix(strings.ReplaceAll(file.Name, "\\", "/"), "/")
		if name == "" || strings.Contains(name, "../") {
			continue
		}
		entries = append(entries, name)
	}
	sort.Strings(entries)
	return entries, nil
}

func copyFile(source, target string) error {
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	in, err := os.Open(source)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(target)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	return err
}
func packFiles(root string) ([]string, error) {
	result := []string{}
	err := filepath.WalkDir(root, func(pathName string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || entry.Name() == "checksums.sha256" {
			return nil
		}
		relative, err := filepath.Rel(root, pathName)
		if err != nil {
			return err
		}
		result = append(result, filepath.ToSlash(relative))
		return nil
	})
	sort.Strings(result)
	return result, err
}
func fileSHA256(pathName string) (string, error) {
	data, err := os.ReadFile(pathName)
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

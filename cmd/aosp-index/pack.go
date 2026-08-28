package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/artifacts"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	coremodules "github.com/i-rtfsc/as-aosp/core/modules"
)

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
	destination := flags.String("destination", "", "Portable Out destination")
	native := flags.Bool("native", false, "include native metadata and compdb")
	ide := flags.Bool("ide", false, "include generated CMake and Rust project metadata")
	modules := stringFlags{}
	flags.Var(&modules, "module", "selected module (repeatable)")
	_ = flags.Parse(args[1:])
	if *aospRoot == "" || *outRoot == "" || *destination == "" {
		fatal(errors.New("--aosp-root, --out-root and --destination are required"))
	}
	if err := exportPack(*aospRoot, *outRoot, *destination, *native, *ide, []string(modules)); err != nil {
		fatal(err)
	}
}

func runPackPreflight(args []string) {
	flags := flag.NewFlagSet("pack preflight", flag.ExitOnError)
	outRoot := flags.String("out-root", "", "AOSP out root")
	native := flags.Bool("native", false, "require native metadata and compdb")
	ide := flags.Bool("ide", false, "require generated CMake and Rust metadata")
	scanAllArtifacts := flags.Bool("scan-all-artifacts", false, "scan all jar, srcjar, and apk files")
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
	blocking := []string{}
	for _, item := range []struct {
		name     string
		relative string
	}{
		{"module-info", "module-info.json"},
		{"java-deps", "module_bp_java_deps.json"},
		{"cc-deps", "module_bp_cc_deps.json"},
		{"compdb", filepath.Join("development", "ide", "compdb", "compile_commands.json")},
	} {
		found, ok := findPackMetadata(outRootAbs, item.relative)
		if !ok {
			found = ""
		}
		checks = append(checks, map[string]any{"name": item.name, "found": found != "", "path": found})
		required := item.name == "module-info" || item.name == "java-deps" || (*native && (item.name == "cc-deps" || item.name == "compdb"))
		if required && found == "" {
			blocking = append(blocking, "missing required "+item.name)
		}
	}
	if *ide {
		if _, ok := findPackMetadata(outRootAbs, filepath.Join("soong", "rust-project.json")); !ok {
			blocking = append(blocking, "missing required IDE metadata: soong/rust-project.json")
		}
		if !hasRegularFileUnder(filepath.Join(outRootAbs, "development", "ide", "clion")) {
			blocking = append(blocking, "missing required IDE metadata: development/ide/clion")
		}
	}
	report := metadata.LoadModuleDirectory(outRootAbs)
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
		if !found {
			blocking = append(blocking, "module not found: "+name)
		}
	}
	srcjars, jars, apks := 0, 0, 0
	if *scanAllArtifacts {
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
	}
	writeJSON(map[string]any{
		"ready":               len(blocking) == 0,
		"outRoot":             outRootAbs,
		"checks":              checks,
		"srcjars":             srcjars,
		"jars":                jars,
		"apks":                apks,
		"modules":             moduleChecks,
		"diagnostics":         report.Diagnostics,
		"blockingDiagnostics": blocking,
		"scannedAllArtifacts": *scanAllArtifacts,
		"suggestions": []string{
			"If module-info is missing, run m module-info.",
			"If compdb is missing, run SOONG_GEN_COMPDB=1 m nothing.",
		},
	})
}

func hasRegularFileUnder(root string) bool {
	found := false
	_ = filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err == nil && !entry.IsDir() {
			found = true
			return filepath.SkipDir
		}
		return nil
	})
	return found
}

func runPackInspect(args []string) {
	flags := flag.NewFlagSet("pack inspect", flag.ExitOnError)
	root := flags.String("root", "", "Portable Out root")
	moduleName := flags.String("module", "", "module name")
	_ = flags.Parse(args)
	if *root == "" {
		fatal(errors.New("--root is required"))
	}
	report := metadata.LoadModuleDirectory(*root)
	catalog := coremodules.Build(report.Snapshot)
	artifactCatalog := artifacts.Build(catalog)
	output := map[string]any{"modules": len(catalog.Modules)}
	jarFiles, jarBytes := packArtifactStats(*root, filepath.Join("soong", ".intermediates"), ".jar")
	srcjarFiles, srcjarBytes := packArtifactStats(*root, filepath.Join("soong", ".intermediates"), ".srcjar")
	output["jars"] = map[string]any{"files": jarFiles, "bytes": jarBytes}
	output["srcjars"] = map[string]any{"files": srcjarFiles, "bytes": srcjarBytes}
	output["artifacts"] = len(artifactCatalog.Artifacts)
	output["diagnostics"] = report.Diagnostics
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

func exportPack(aospRoot, outRoot, destination string, native, ide bool, modules []string) error {
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
	if err := os.MkdirAll(tmp, 0o755); err != nil {
		return err
	}
	coreFiles := []string{"module-info.json", "module_bp_java_deps.json"}
	if native {
		coreFiles = append(coreFiles, "module_bp_cc_deps.json", filepath.Join("development", "ide", "compdb", "compile_commands.json"))
	}
	if ide {
		if err := copyOptionalIDERoots(tmp, outRoot); err != nil {
			return err
		}
	}
	for _, relative := range coreFiles {
		source, ok := findPackMetadata(outRoot, relative)
		if !ok {
			if relative == "module-info.json" {
				return fmt.Errorf("required metadata missing: module-info.json; run `m module-info` in the lunch shell")
			}
			if relative == "module_bp_java_deps.json" {
				return fmt.Errorf("required metadata missing: module_bp_java_deps.json; run `m nothing` in the lunch shell")
			}
			continue
		}
		outRelative, relErr := filepath.Rel(outRoot, source)
		if relErr != nil || strings.HasPrefix(outRelative, ".."+string(filepath.Separator)) {
			return fmt.Errorf("metadata is outside out root: %s", source)
		}
		target := filepath.Join(tmp, outRelative)
		if err := copyFile(source, target); err != nil {
			return err
		}
	}
	if err := copyPackLibraries(tmp, outRoot, modules); err != nil {
		return err
	}
	selectedArchives := map[string]bool{}
	if len(modules) > 0 {
		catalog := coremodules.Build(metadata.LoadModuleDirectory(outRoot).Snapshot)
		for _, name := range catalog.Closure(modules) {
			for _, ref := range append(catalog.Modules[name].SrcJars, catalog.Modules[name].Generated...) {
				raw := filepath.ToSlash(ref.Raw)
				if marker := strings.Index(raw, ".intermediates/"); marker >= 0 {
					selectedArchives[strings.TrimPrefix(raw[marker+len(".intermediates/"):], "/")] = true
				}
			}
		}
	}
	copySrcJar := func(pathName, intermediateRelative string) error {
		if info, statErr := os.Stat(pathName); statErr != nil || info.IsDir() || strings.ToLower(filepath.Ext(pathName)) != ".srcjar" {
			return nil
		}
		packPath := filepath.Join("soong", ".intermediates", intermediateRelative)
		target := filepath.Join(tmp, packPath)
		if err := copyFile(pathName, target); err != nil {
			return err
		}
		return nil
	}
	if len(modules) > 0 {
		for intermediateRelative := range selectedArchives {
			pathName := filepath.Join(outRoot, "soong", ".intermediates", filepath.FromSlash(intermediateRelative))
			if err := copySrcJar(pathName, intermediateRelative); err != nil {
				return err
			}
		}
	} else {
		err = filepath.WalkDir(filepath.Join(outRoot, "soong", ".intermediates"), func(pathName string, entry os.DirEntry, walkErr error) error {
			if walkErr != nil {
				return nil
			}
			if entry.IsDir() || strings.ToLower(filepath.Ext(entry.Name())) != ".srcjar" {
				return nil
			}
			intermediateRelative, relErr := filepath.Rel(filepath.Join(outRoot, "soong", ".intermediates"), pathName)
			if relErr != nil {
				return relErr
			}
			return copySrcJar(pathName, filepath.ToSlash(intermediateRelative))
		})
		if err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	if err := validatePortableOut(tmp, native, modules); err != nil {
		return fmt.Errorf("exported portable out invalid: %w", err)
	}
	_ = os.RemoveAll(destination)
	return os.Rename(tmp, destination)
}

func copyOptionalIDERoots(destination, outRoot string) error {
	for _, relative := range []string{"soong/rust-project.json", "soong/rust-project.kzip"} {
		source := filepath.Join(outRoot, filepath.FromSlash(relative))
		if info, err := os.Stat(source); err == nil && !info.IsDir() {
			if err := copyFile(source, filepath.Join(destination, filepath.FromSlash(relative))); err != nil {
				return err
			}
		}
	}
	root := filepath.Join(outRoot, "development", "ide", "clion")
	return filepath.WalkDir(root, func(pathName string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			if os.IsNotExist(walkErr) {
				return nil
			}
			return walkErr
		}
		if entry.IsDir() {
			return nil
		}
		relative, err := filepath.Rel(outRoot, pathName)
		if err != nil {
			return err
		}
		return copyFile(pathName, filepath.Join(destination, relative))
	})
}

func copyPackLibraries(root, outRoot string, names []string) error {
	catalog := coremodules.Build(metadata.LoadModuleDirectory(outRoot).Snapshot)
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
		return raw, filepath.ToSlash(filepath.FromSlash(suffix)), true
	}
	if !strings.HasSuffix(strings.ToLower(normalized), ".jar") {
		return "", "", false
	}
	suffix := strings.TrimPrefix(normalized, "out/")
	candidates := []struct{ source, suffix string }{{filepath.Join(outRoot, filepath.FromSlash(suffix)), suffix}}
	outName := filepath.Base(outRoot)
	if strings.HasPrefix(normalized, outName+"/") {
		suffix := strings.TrimPrefix(normalized, outName+"/")
		candidates = append(candidates, struct{ source, suffix string }{filepath.Join(outRoot, filepath.FromSlash(suffix)), suffix})
	}
	for _, candidate := range candidates {
		if _, err := os.Stat(candidate.source); err == nil {
			return candidate.source, filepath.ToSlash(filepath.FromSlash(candidate.suffix)), true
		}
	}
	return "", "", false
}

func runPackValidate(args []string) {
	flags := flag.NewFlagSet("pack validate", flag.ExitOnError)
	root := flags.String("root", "", "Portable Out root")
	native := flags.Bool("native", false, "require native metadata and compdb")
	ide := flags.Bool("ide", false, "require generated CMake and Rust metadata")
	_ = flags.Parse(args)
	if *root == "" {
		fatal(errors.New("--root is required"))
	}
	if err := validatePortableOut(*root, *native, nil); err != nil {
		fatal(err)
	}
	if *ide {
		if err := validateIDERoot(*root); err != nil {
			fatal(err)
		}
	}
}

func validateIDERoot(root string) error {
	if _, ok := findPackMetadata(root, filepath.Join("soong", "rust-project.json")); !ok {
		return errors.New("required IDE metadata is missing: soong/rust-project.json")
	}
	if !hasRegularFileUnder(filepath.Join(root, "development", "ide", "clion")) {
		return errors.New("required IDE metadata is missing: development/ide/clion")
	}
	return nil
}

func validatePortableOut(root string, native bool, selectedModules []string) error {
	for _, relative := range []string{"module-info.json", "module_bp_java_deps.json"} {
		if _, ok := findPackMetadata(root, relative); !ok {
			return fmt.Errorf("required build fact is missing: %s", relative)
		}
	}
	if native {
		for _, relative := range []string{"module_bp_cc_deps.json", filepath.Join("development", "ide", "compdb", "compile_commands.json")} {
			if _, ok := findPackMetadata(root, relative); !ok {
				return fmt.Errorf("required native build fact is missing: %s", relative)
			}
		}
	}
	report := metadata.LoadModuleDirectory(root)
	for _, diagnostic := range report.Diagnostics {
		if diagnostic.Code == "METADATA_READ_FAILED" || diagnostic.Code == "METADATA_PARSE_FAILED" {
			return fmt.Errorf("%s: %s", diagnostic.Code, diagnostic.Message)
		}
	}
	catalog := coremodules.Build(report.Snapshot)
	if len(catalog.Modules) == 0 {
		return fmt.Errorf("module catalog is empty")
	}
	for _, name := range selectedModules {
		if _, ok := catalog.Modules[name]; !ok {
			return fmt.Errorf("selected module not found: %s", name)
		}
	}
	return nil
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

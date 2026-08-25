package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
	if len(args) == 0 || args[0] != "export" {
		fatal(errors.New("usage: aosp-index pack export --aosp-root <root> --out-root <out> --destination <dir> [--native] [--module name]"))
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

type stringFlags []string

func (s *stringFlags) String() string         { return strings.Join(*s, ",") }
func (s *stringFlags) Set(value string) error { *s = append(*s, value); return nil }

func exportPack(aospRoot, outRoot, destination string, native bool, modules []string) error {
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
		source := filepath.Join(outRoot, relative)
		if _, err := os.Stat(source); err != nil {
			continue
		}
		target := filepath.Join(tmp, relative)
		if err := copyFile(source, target); err != nil {
			return err
		}
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
		target := filepath.Join(tmp, "generated", "srcjars", filepath.Base(pathName))
		if err := copyFile(pathName, target); err != nil {
			return err
		}
		generatedManifest = append(generatedManifest, map[string]any{"source": filepath.ToSlash(relative), "path": filepath.ToSlash(filepath.Join("generated", "srcjars", filepath.Base(pathName)))})
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
	_ = os.RemoveAll(destination)
	return os.Rename(tmp, destination)
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

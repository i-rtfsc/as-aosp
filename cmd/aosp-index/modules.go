package main

import (
	"errors"
	"flag"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/metadata"
	coremodules "github.com/i-rtfsc/as-aosp/core/modules"
	"github.com/i-rtfsc/as-aosp/core/sourcecatalog"
)

func runModules(args []string) {
	flags := flag.NewFlagSet("modules", flag.ExitOnError)
	root := flags.String("root", "", "workspace source root")
	outRoot := flags.String("out-root", "", "optional AOSP Out Root")
	cacheRoot := flags.String("cache-root", "", "workspace module cache root")
	query := flags.String("query", "", "module name, type, or path filter")
	summary := flags.Bool("summary", false, "print catalog summary")
	_ = flags.Parse(args)
	if strings.TrimSpace(*root) == "" {
		fatal(errors.New("--root is required"))
	}
	report := sourcecatalog.Discover(*root, *cacheRoot)
	if strings.TrimSpace(*outRoot) != "" {
		buildReport := metadata.LoadModuleDirectory(*outRoot)
		report.Diagnostics = append(report.Diagnostics, buildReport.Diagnostics...)
		report.Catalog.Modules = sourcecatalog.MergeBuildFacts(
			report.Catalog.Modules,
			coremodules.Build(buildReport.Snapshot),
		)
	}
	modules := report.Catalog.Modules
	if value := strings.ToLower(strings.TrimSpace(*query)); value != "" {
		filtered := make([]sourcecatalog.Module, 0)
		for _, module := range modules {
			if strings.Contains(strings.ToLower(module.Name), value) ||
				strings.Contains(strings.ToLower(module.Type), value) ||
				strings.Contains(strings.ToLower(string(module.Blueprint)), value) {
				filtered = append(filtered, module)
			}
		}
		modules = filtered
	}
	if !*summary {
		writeJSON(map[string]any{
			"modules": modules, "filesScanned": report.FilesScanned,
			"cacheHit": report.CacheHit, "diagnostics": report.Diagnostics,
		})
		return
	}
	buildBacked, recommended := 0, []string{}
	for _, module := range modules {
		if module.Status == sourcecatalog.StatusBuildBacked {
			buildBacked++
		}
		if module.Recommended {
			recommended = append(recommended, module.Name)
		}
	}
	writeJSON(map[string]any{
		"modules": len(modules), "buildBacked": buildBacked,
		"recommended": recommended, "filesScanned": report.FilesScanned,
		"cacheHit": report.CacheHit, "diagnostics": report.Diagnostics,
	})
}

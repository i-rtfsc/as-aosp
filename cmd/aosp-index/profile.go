package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/i-rtfsc/as-aosp/core/profiles"
)

func runProfile(args []string) {
	if len(args) == 0 || (args[0] != "compile" && args[0] != "validate") {
		fatal(errors.New("usage: aosp-index profile <compile|validate> ..."))
	}
	if args[0] == "validate" {
		runProfileValidate(args[1:])
		return
	}
	flags := flag.NewFlagSet("profile compile", flag.ExitOnError)
	outRoot := flags.String("out-root", "", "development AOSP Out root used for validation")
	release := flags.String("release", "android-17", "Android release baseline")
	tag := flags.String("tag", "", "validated Android source tag")
	output := flags.String("output", "profiles/common.json", "compact common profile output")
	_ = flags.Parse(args[1:])
	report, err := profiles.Compile(*outRoot, *release, *tag)
	if err != nil {
		fatal(err)
	}
	if err := profiles.WriteCatalog(*output, report.Catalog); err != nil {
		fatal(err)
	}
	writeJSON(map[string]any{
		"output": *output, "release": report.Catalog.Baseline.AndroidRelease,
		"tag":      report.Catalog.Baseline.AndroidTag,
		"profiles": len(report.Catalog.Profiles), "anchorsChecked": report.AnchorsChecked,
		"anchorsFound": report.AnchorsFound, "diagnostics": report.Diagnostics,
	})
}

func runProfileValidate(args []string) {
	flags := flag.NewFlagSet("profile validate", flag.ExitOnError)
	input := flags.String("input", "profiles/common.json", "compact profile catalog")
	_ = flags.Parse(args)
	catalog, err := profiles.LoadCatalog(*input)
	if err != nil {
		fatal(err)
	}
	diagnostics := profiles.ValidateCatalog(catalog)
	writeJSON(map[string]any{
		"valid": len(diagnostics) == 0, "input": *input,
		"profiles": len(catalog.Profiles), "diagnostics": diagnostics,
	})
	if len(diagnostics) > 0 {
		os.Exit(1)
	}
}

func profileFatal(format string, args ...any) {
	fatal(fmt.Errorf(format, args...))
}

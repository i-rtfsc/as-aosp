package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/i-rtfsc/as-aosp/core/graph"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/topology"
)

type scanOutput struct {
	Profile     model.Profile              `json:"profile"`
	Topology    topology.WorkspaceTopology `json:"topology"`
	Diagnostics []model.Diagnostic         `json:"diagnostics"`
}

func main() {
	if len(os.Args) < 2 {
		fatal(errors.New("usage: aosp-index <doctor|scan|metadata> [options]"))
	}
	switch os.Args[1] {
	case "doctor":
		runDoctor(os.Args[2:])
	case "scan":
		runScan(os.Args[2:])
	case "metadata":
		runMetadata(os.Args[2:])
	default:
		fatal(fmt.Errorf("unknown command %q", os.Args[1]))
	}
}

type metadataOutput struct {
	Report  metadata.ReadReport `json:"report"`
	Graph   graph.ModuleGraph   `json:"graph"`
	Summary metadata.Summary    `json:"summary"`
}

func runMetadata(args []string) {
	flags := flag.NewFlagSet("metadata", flag.ExitOnError)
	root := flags.String("root", "", "metadata root")
	_ = flags.Parse(args)
	if *root == "" {
		fatal(errors.New("--root is required"))
	}
	report := metadata.LoadDirectory(*root)
	moduleGraph := graph.Build(report.Snapshot)
	writeJSON(metadataOutput{Report: report, Graph: moduleGraph, Summary: report.Snapshot.Summary(len(moduleGraph.Modules))})
}

func runDoctor(args []string) {
	profile := loadProfile(args)
	diagnostics := profile.Validate()
	writeJSON(map[string]any{"profile": profile, "diagnostics": diagnostics})
	if hasError(diagnostics) {
		os.Exit(2)
	}
}

func runScan(args []string) {
	profile := loadProfile(args)
	diagnostics := profile.Validate()
	if hasError(diagnostics) {
		writeJSON(map[string]any{"profile": profile, "diagnostics": diagnostics})
		os.Exit(2)
	}
	view := topology.FromProfile(profile)
	if len(view.SourceMounts) == 0 {
		discovered, discoveryDiagnostics, err := topology.Discover(topology.DiscoverOptions{WorkspaceViewRoot: profile.WorkspaceViewRoot, CanonicalRepoRoot: profile.CanonicalRepoRoot, MaxDepth: 4})
		if err != nil {
			fatal(err)
		}
		view = discovered
		diagnostics = append(diagnostics, discoveryDiagnostics...)
	} else {
		diagnostics = append(diagnostics, view.Validate()...)
	}
	writeJSON(scanOutput{Profile: profile, Topology: view, Diagnostics: diagnostics})
	if hasError(diagnostics) {
		os.Exit(2)
	}
}

func loadProfile(args []string) model.Profile {
	flags := flag.NewFlagSet("profile", flag.ExitOnError)
	profilePath := flags.String("profile", "", "profile JSON path")
	_ = flags.Parse(args)
	if *profilePath == "" {
		fatal(errors.New("--profile is required"))
	}
	data, err := os.ReadFile(*profilePath)
	if err != nil {
		fatal(err)
	}
	var profile model.Profile
	if err := json.Unmarshal(data, &profile); err != nil {
		fatal(fmt.Errorf("parse profile: %w", err))
	}
	return profile
}

func hasError(diagnostics []model.Diagnostic) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Severity == model.SeverityError {
			return true
		}
	}
	return false
}

func writeJSON(value any) {
	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		fatal(err)
	}
}

func fatal(err error) {
	_, _ = fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

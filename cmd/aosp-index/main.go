package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/i-rtfsc/as-aosp/core/aidl"
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
		fatal(errors.New("usage: aosp-index <doctor|scan|metadata|aidl> [options]"))
	}
	switch os.Args[1] {
	case "doctor":
		runDoctor(os.Args[2:])
	case "scan":
		runScan(os.Args[2:])
	case "metadata":
		runMetadata(os.Args[2:])
	case "aidl":
		runAidl(os.Args[2:])
	default:
		fatal(fmt.Errorf("unknown command %q", os.Args[1]))
	}
}

type aidlOutput struct {
	Report     aidl.LoadReport  `json:"report"`
	Resolution *aidl.Resolution `json:"resolution,omitempty"`
}

type aidlSummaryOutput struct {
	Summary     aidl.Summary       `json:"summary"`
	Resolution  *aidl.Resolution   `json:"resolution,omitempty"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

func runAidl(args []string) {
	flags := flag.NewFlagSet("aidl", flag.ExitOnError)
	root := flags.String("root", "", "AIDL source root")
	file := flags.String("file", "", "RepoPath of the AIDL document")
	typeName := flags.String("type", "", "canonical or imported type to resolve")
	summaryOnly := flags.Bool("summary", false, "print only summary and resolution")
	_ = flags.Parse(args)
	if *root == "" {
		fatal(errors.New("--root is required"))
	}
	report, err := aidl.LoadTree(*root)
	if err != nil {
		fatal(err)
	}
	output := aidlOutput{Report: report}
	if *file != "" && *typeName != "" {
		repoPath, err := model.NormalizeRepoPath(*file)
		if err != nil {
			fatal(err)
		}
		document, ok := report.Index.Documents[repoPath]
		if !ok {
			fatal(fmt.Errorf("AIDL document not found: %s", repoPath))
		}
		resolution := report.Index.Resolve(document, *typeName)
		output.Resolution = &resolution
	}
	if *summaryOnly {
		writeJSON(aidlSummaryOutput{Summary: report.Index.Summary(len(report.Diagnostics)), Resolution: output.Resolution, Diagnostics: report.Diagnostics})
		return
	}
	writeJSON(output)
}

type metadataOutput struct {
	Report  metadata.ReadReport `json:"report"`
	Graph   graph.ModuleGraph   `json:"graph"`
	Summary metadata.Summary    `json:"summary"`
}

type metadataSummaryOutput struct {
	Summary     metadata.Summary   `json:"summary"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

func runMetadata(args []string) {
	flags := flag.NewFlagSet("metadata", flag.ExitOnError)
	root := flags.String("root", "", "metadata root")
	summaryOnly := flags.Bool("summary", false, "print only summary and diagnostics")
	_ = flags.Parse(args)
	if *root == "" {
		fatal(errors.New("--root is required"))
	}
	report := metadata.LoadDirectory(*root)
	moduleGraph := graph.Build(report.Snapshot)
	if *summaryOnly {
		writeJSON(metadataSummaryOutput{Summary: report.Snapshot.Summary(len(moduleGraph.Modules)), Diagnostics: report.Diagnostics})
		return
	}
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

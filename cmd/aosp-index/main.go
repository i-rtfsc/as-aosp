package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/generated"
	"github.com/i-rtfsc/as-aosp/core/graph"
	coreindex "github.com/i-rtfsc/as-aosp/core/index"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/pathmap"
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
	Peers      *aidl.PeerReport `json:"peers,omitempty"`
}

type aidlSummaryOutput struct {
	Summary     aidl.Summary       `json:"summary"`
	Resolution  *aidl.Resolution   `json:"resolution,omitempty"`
	Peers       *aidl.PeerReport   `json:"peers,omitempty"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

type mappingFlags []string

func (values *mappingFlags) String() string {
	return strings.Join(*values, ",")
}

func (values *mappingFlags) Set(value string) error {
	if !strings.Contains(value, "=") {
		return fmt.Errorf("path mapping must use external=local")
	}
	*values = append(*values, value)
	return nil
}

func runAidl(args []string) {
	flags := flag.NewFlagSet("aidl", flag.ExitOnError)
	root := flags.String("root", "", "AIDL source root")
	file := flags.String("file", "", "RepoPath of the AIDL document")
	typeName := flags.String("type", "", "canonical or imported type to resolve")
	metadataRoot := flags.String("metadata-root", "", "Soong metadata root")
	moduleName := flags.String("module", "", "Soong module for generated peers")
	cacheRoot := flags.String("cache-root", "", "generated source cache root")
	summaryOnly := flags.Bool("summary", false, "print only summary and resolution")
	var rawMappings mappingFlags
	flags.Var(&rawMappings, "path-map", "external=local generated path mapping")
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
		if *metadataRoot != "" && *moduleName != "" {
			metadataReport := metadata.LoadDirectory(*metadataRoot)
			moduleGraph := graph.Build(metadataReport.Snapshot)
			peers := buildAidlPeers(
				document,
				*typeName,
				moduleGraph,
				*moduleName,
				*cacheRoot,
				rawMappings,
			)
			output.Peers = &peers
		} else if *metadataRoot != "" {
			metadataReport := metadata.LoadDirectory(*metadataRoot)
			moduleGraph := graph.Build(metadataReport.Snapshot)
			owner, found, ambiguous := moduleGraph.AidlOwner(document.Path)
			if !found || ambiguous {
				peers := coreindex.AidlGeneratedPeersAutoOwner(
					document,
					*typeName,
					moduleGraph,
					func(metadata.PathRef) bool { return false },
				)
				output.Peers = &peers
			} else {
				peers := buildAidlPeers(
					document,
					*typeName,
					moduleGraph,
					owner,
					*cacheRoot,
					rawMappings,
				)
				output.Peers = &peers
			}
		}
	}
	if *summaryOnly {
		writeJSON(aidlSummaryOutput{Summary: report.Index.Summary(len(report.Diagnostics)), Resolution: output.Resolution, Peers: output.Peers, Diagnostics: report.Diagnostics})
		return
	}
	writeJSON(output)
}

func buildAidlPeers(
	document aidl.Document,
	typeName string,
	moduleGraph graph.ModuleGraph,
	moduleName string,
	cacheRoot string,
	rawMappings []string,
) aidl.PeerReport {
	mapper := pathmap.Mapper{Mappings: parsePathMappings(rawMappings)}
	exists := func(ref metadata.PathRef) bool {
		local, ok := localMetadataPath(mapper, ref)
		if !ok {
			return false
		}
		_, err := os.Stat(string(local))
		return err == nil
	}
	archiveEntries := make([]aidl.ArchiveEntry, 0)
	diagnostics := make([]model.Diagnostic, 0)
	if cacheRoot != "" {
		for _, archiveRef := range moduleGraph.SourceArchives(moduleName) {
			local, ok := localMetadataPath(mapper, archiveRef)
			if !ok {
				continue
			}
			extracted := generated.ExtractSrcJar(
				string(local),
				cacheRoot,
				generated.DefaultLimits(),
			)
			diagnostics = append(diagnostics, extracted.Diagnostics...)
			if extracted.Archive == nil {
				continue
			}
			for entry, entryPath := range extracted.Archive.Entries {
				archiveEntries = append(archiveEntries, aidl.ArchiveEntry{
					Archive:   archiveRef,
					Entry:     entry,
					LocalPath: entryPath,
				})
			}
		}
	}
	peers := coreindex.AidlGeneratedPeersWithArchives(
		document,
		typeName,
		moduleGraph,
		moduleName,
		exists,
		archiveEntries,
	)
	peers.Diagnostics = append(peers.Diagnostics, diagnostics...)
	return peers
}

func localMetadataPath(
	mapper pathmap.Mapper,
	ref metadata.PathRef,
) (model.LocalPath, bool) {
	if ref.ExternalPath != "" {
		return mapper.ExternalToLocal(ref.ExternalPath)
	}
	return "", false
}

func parsePathMappings(values []string) []model.PathMapping {
	mappings := make([]model.PathMapping, 0, len(values))
	for _, value := range values {
		parts := strings.SplitN(value, "=", 2)
		mappings = append(mappings, model.PathMapping{
			ExternalPrefix: parts[0],
			LocalPrefix:    parts[1],
			Kind:           "generated",
		})
	}
	return mappings
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

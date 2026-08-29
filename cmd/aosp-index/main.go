package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/aidl"
	"github.com/i-rtfsc/as-aosp/core/generated"
	"github.com/i-rtfsc/as-aosp/core/graph"
	coreindex "github.com/i-rtfsc/as-aosp/core/index"
	"github.com/i-rtfsc/as-aosp/core/jni"
	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/pathmap"
	"github.com/i-rtfsc/as-aosp/core/topology"
	"github.com/i-rtfsc/as-aosp/protocol/aosp"
	"github.com/i-rtfsc/as-aosp/protocol/jsonrpc"
)

type scanOutput struct {
	Profile     model.Profile              `json:"profile"`
	Topology    topology.WorkspaceTopology `json:"topology"`
	Diagnostics []model.Diagnostic         `json:"diagnostics"`
}

func main() {
	if len(os.Args) < 2 {
		fatal(errors.New("usage: aosp-index <doctor|scan|metadata|modules|aidl|jni|profile|pack|serve> [options]"))
	}
	switch os.Args[1] {
	case "doctor":
		runDoctor(os.Args[2:])
	case "scan":
		runScan(os.Args[2:])
	case "metadata":
		runMetadata(os.Args[2:])
	case "modules":
		runModules(os.Args[2:])
	case "aidl":
		runAidl(os.Args[2:])
	case "jni":
		runJni(os.Args[2:])
	case "profile":
		runProfile(os.Args[2:])
	case "pack":
		runPack(os.Args[2:])
	case "serve":
		runServe(os.Args[2:])
	default:
		fatal(fmt.Errorf("unknown command %q", os.Args[1]))
	}
}

func runServe(args []string) {
	flags := flag.NewFlagSet("serve", flag.ExitOnError)
	stdio := flags.Bool("stdio", false, "serve JSON-RPC over stdin/stdout")
	_ = flags.Parse(args)
	if !*stdio || flags.NArg() != 0 {
		fatal(errors.New("usage: aosp-index serve --stdio"))
	}
	if err := runJSONRPC(os.Stdin, os.Stdout, log.New(os.Stderr, "", 0)); err != nil {
		fatal(err)
	}
}

func runJSONRPC(reader io.Reader, writer io.Writer, logger *log.Logger) error {
	framer := jsonrpc.NewFramer(reader)
	server := &aosp.Server{}
	for {
		payload, err := framer.Read()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		request, parseErr := jsonrpc.ParseRequest(payload)
		if parseErr != nil {
			rpcError, ok := parseErr.(*jsonrpc.Error)
			if !ok {
				return parseErr
			}
			if err := writeRPCResponse(writer, jsonrpc.Failure(nil, rpcError)); err != nil {
				return err
			}
			return nil
		}
		response, reply := server.Handle(request)
		if reply {
			if err := writeRPCResponse(writer, response); err != nil {
				return err
			}
		}
		if server.IsShutdown() {
			logger.Print("shutdown complete")
			return nil
		}
	}
}

func writeRPCResponse(writer io.Writer, response jsonrpc.Response) error {
	payload, err := json.Marshal(response)
	if err != nil {
		return err
	}
	return jsonrpc.Write(writer, payload)
}

type jniOutput struct {
	Index jni.Index `json:"index"`
}

type jniSummaryOutput struct {
	JavaSymbols   int                `json:"javaSymbols"`
	NativeSymbols int                `json:"nativeSymbols"`
	Registrations int                `json:"registrations"`
	Links         int                `json:"links"`
	ExactLinks    int                `json:"exactLinks"`
	InferredLinks int                `json:"inferredLinks"`
	Diagnostics   []model.Diagnostic `json:"diagnostics"`
	Peers         []jni.Link         `json:"peers,omitempty"`
}

func runJni(args []string) {
	flags := flag.NewFlagSet("jni", flag.ExitOnError)
	javaRoot := flags.String("java-root", "", "Java/Kotlin source root")
	cppRoot := flags.String("cpp-root", "", "C/C++ source root")
	className := flags.String("class", "", "Java class name to inspect")
	methodName := flags.String("method", "", "native method name to inspect")
	summaryOnly := flags.Bool("summary", false, "print only index counts")
	_ = flags.Parse(args)
	if *javaRoot == "" || *cppRoot == "" {
		fatal(errors.New("--java-root and --cpp-root are required"))
	}
	javaFiles, javaDiagnostics := loadSourceFiles(*javaRoot, map[string]bool{
		".java": true,
		".kt":   true,
	})
	cppFiles, cppDiagnostics := loadSourceFiles(*cppRoot, map[string]bool{
		".c":   true,
		".cc":  true,
		".cpp": true,
		".cxx": true,
		".h":   true,
		".hh":  true,
		".hpp": true,
	})
	index := jni.BuildIndex(javaFiles, cppFiles)
	index.Diagnostics = append(index.Diagnostics, javaDiagnostics...)
	index.Diagnostics = append(index.Diagnostics, cppDiagnostics...)
	if *summaryOnly {
		output := jniSummaryOutput{
			JavaSymbols:   len(index.JavaSymbols),
			NativeSymbols: len(index.NativeSymbols),
			Registrations: len(index.Registrations),
			Links:         len(index.Links),
			Diagnostics:   index.Diagnostics,
		}
		for _, link := range index.Links {
			switch link.Confidence {
			case jni.ConfidenceExact:
				output.ExactLinks++
			case jni.ConfidenceInferred:
				output.InferredLinks++
			}
		}
		if *className != "" && *methodName != "" {
			output.Peers = index.JavaPeers(*className, *methodName)
		}
		writeJSON(output)
		return
	}
	writeJSON(jniOutput{Index: index})
}

func loadSourceFiles(root string, extensions map[string]bool) (map[model.RepoPath]string, []model.Diagnostic) {
	files := make(map[model.RepoPath]string)
	diagnostics := make([]model.Diagnostic, 0)
	err := filepath.WalkDir(root, func(pathName string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			diagnostics = append(diagnostics, model.Diagnostic{
				Code:      "JNI_SOURCE_READ_FAILED",
				Severity:  model.SeverityWarning,
				Stage:     "jni-load",
				Message:   walkErr.Error(),
				LocalPath: pathName,
			})
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		if !extensions[strings.ToLower(filepath.Ext(pathName))] {
			return nil
		}
		relative, err := filepath.Rel(root, pathName)
		if err != nil {
			return err
		}
		repoPath, err := model.NormalizeRepoPath(relative)
		if err != nil {
			diagnostics = append(diagnostics, model.Diagnostic{
				Code:      "JNI_SOURCE_PATH_INVALID",
				Severity:  model.SeverityWarning,
				Stage:     "jni-load",
				Message:   err.Error(),
				LocalPath: pathName,
			})
			return nil
		}
		data, err := os.ReadFile(pathName)
		if err != nil {
			diagnostics = append(diagnostics, model.Diagnostic{
				Code:      "JNI_SOURCE_READ_FAILED",
				Severity:  model.SeverityWarning,
				Stage:     "jni-load",
				Message:   err.Error(),
				LocalPath: pathName,
			})
			return nil
		}
		files[repoPath] = string(data)
		return nil
	})
	if err != nil {
		diagnostics = append(diagnostics, model.Diagnostic{
			Code:      "JNI_SOURCE_WALK_FAILED",
			Severity:  model.SeverityWarning,
			Stage:     "jni-load",
			Message:   err.Error(),
			LocalPath: root,
		})
	}
	return files, diagnostics
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
	outRoot := flags.String("out-root", "", "AOSP out root")
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
		if *outRoot != "" && *moduleName != "" {
			metadataReport := metadata.LoadModuleDirectory(*outRoot)
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
		} else if *outRoot != "" {
			metadataReport := metadata.LoadModuleDirectory(*outRoot)
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

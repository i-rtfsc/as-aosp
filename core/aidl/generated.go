package aidl

import (
	"path"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
)

type PeerLanguage string

const (
	PeerJava PeerLanguage = "java"
	PeerCpp  PeerLanguage = "cpp"
	PeerNdk  PeerLanguage = "ndk"
)

type PeerKind string

const (
	PeerSource    PeerKind = "source"
	PeerCppSource PeerKind = "cpp-source"
	PeerHeader    PeerKind = "header"
)

type GeneratedPeer struct {
	Source     model.RepoPath   `json:"source"`
	Type       string           `json:"type"`
	Language   PeerLanguage     `json:"language"`
	Kind       PeerKind         `json:"kind"`
	Path       metadata.PathRef `json:"path"`
	Origin     string           `json:"origin"`
	Confidence string           `json:"confidence"`
	Exists     bool             `json:"exists"`
	LocalPath  model.LocalPath  `json:"localPath,omitempty"`
}

type ArchiveEntry struct {
	Archive   metadata.PathRef `json:"archive"`
	Entry     string           `json:"entry"`
	LocalPath model.LocalPath  `json:"localPath"`
}

type PeerOptions struct {
	GeneratedPaths []metadata.PathRef
	IncludeRoots   []metadata.PathRef
	SourceArchives []metadata.PathRef
	ArchiveEntries []ArchiveEntry
	Exists         func(metadata.PathRef) bool
}

type PeerReport struct {
	Peers          []GeneratedPeer    `json:"peers"`
	IncludeRoots   []metadata.PathRef `json:"includeRoots,omitempty"`
	SourceArchives []metadata.PathRef `json:"sourceArchives,omitempty"`
	Diagnostics    []model.Diagnostic `json:"diagnostics"`
}

func (index Index) GeneratedPeers(document Document, typeName string, options PeerOptions) PeerReport {
	report := PeerReport{
		Peers:          make([]GeneratedPeer, 0),
		IncludeRoots:   options.IncludeRoots,
		SourceArchives: options.SourceArchives,
		Diagnostics:    make([]model.Diagnostic, 0),
	}
	declaration, ownerName, nested, ok := findDeclaration(document.Declarations, typeName)
	if !ok {
		report.Diagnostics = append(report.Diagnostics, model.Diagnostic{Code: "AIDL_TYPE_NOT_FOUND", Severity: model.SeverityWarning, Stage: "aidl-generated", Message: "declaration was not found: " + typeName, RepoPath: document.Path})
		return report
	}
	packagePath := strings.ReplaceAll(document.Package, ".", "/")
	basePath := path.Join(packagePath, ownerName)
	baseName := declaration.Name
	peers := make([]GeneratedPeer, 0)
	add := func(language PeerLanguage, kind PeerKind, raw string) {
		ref := metadata.NewPathRef(raw)
		ref.Kind = metadata.PathGenerated
		peers = append(peers, GeneratedPeer{Source: document.Path, Type: declaration.CanonicalName, Language: language, Kind: kind, Path: ref, Origin: "naming-rule", Confidence: "inferred"})
	}
	add(PeerJava, PeerSource, basePath+".java")
	if nested {
		report.Peers = peers
		report = applyGeneratedPaths(report, options)
		return applyArchiveEntries(report, options)
	}
	if declaration.Kind == DeclarationInterface {
		add(PeerCpp, PeerCppSource, basePath+".cpp")
		cppBase := interfaceBaseName(baseName)
		add(PeerCpp, PeerHeader, path.Join(packagePath, "Bp"+cppBase+".h"))
		add(PeerCpp, PeerHeader, path.Join(packagePath, "Bn"+cppBase+".h"))
		add(PeerNdk, PeerHeader, path.Join("aidl", packagePath, "Bp"+cppBase+".h"))
		add(PeerNdk, PeerHeader, path.Join("aidl", packagePath, "Bn"+cppBase+".h"))
	} else if declaration.Nested == nil || len(declaration.Nested) == 0 {
		add(PeerCpp, PeerCppSource, basePath+".cpp")
		add(PeerCpp, PeerHeader, basePath+".h")
		add(PeerNdk, PeerHeader, path.Join("aidl", packagePath, baseName+".h"))
	}
	report.Peers = peers
	report = applyGeneratedPaths(report, options)
	return applyArchiveEntries(report, options)
}

func applyGeneratedPaths(report PeerReport, options PeerOptions) PeerReport {
	for index := range report.Peers {
		peer := &report.Peers[index]
		for _, generated := range options.GeneratedPaths {
			if !matchesGeneratedPeer(*peer, generated.Raw) {
				continue
			}
			if options.Exists == nil || !options.Exists(generated) {
				report.Diagnostics = append(report.Diagnostics, model.Diagnostic{Code: "AIDL_GENERATED_MISSING", Severity: model.SeverityWarning, Stage: "aidl-generated", Message: "generated peer path is referenced but not visible: " + generated.Raw, RepoPath: peer.Source})
				continue
			}
			peer.Path = generated
			peer.Origin = "build-metadata"
			peer.Confidence = "exact"
			peer.Exists = true
			break
		}
	}
	return report
}

func matchesGeneratedPeer(peer GeneratedPeer, actual string) bool {
	actual = strings.TrimPrefix(
		path.Clean(strings.ReplaceAll(actual, "\\", "/")),
		"./",
	)
	if !sameGeneratedSuffix(peer.Path.Raw, actual) {
		return false
	}
	if peer.Language == PeerCpp && peer.Kind == PeerHeader {
		expected := strings.TrimPrefix(peer.Path.Raw, "./")
		if strings.HasSuffix(actual, "/aidl/"+expected) {
			return false
		}
	}
	if peer.Language == PeerNdk {
		return strings.Contains(actual, "/aidl/") ||
			strings.HasPrefix(actual, "aidl/")
	}
	return true
}

func applyArchiveEntries(
	report PeerReport,
	options PeerOptions,
) PeerReport {
	foundJava := false
	for index := range report.Peers {
		peer := &report.Peers[index]
		if peer.Language != PeerJava || peer.Exists {
			if peer.Language == PeerJava && peer.Exists {
				foundJava = true
			}
			continue
		}
		for _, entry := range options.ArchiveEntries {
			if !sameGeneratedSuffix(peer.Path.Raw, entry.Entry) {
				continue
			}
			peer.Path = metadata.PathRef{
				Raw:  entry.Archive.Raw + "!/" + entry.Entry,
				Kind: metadata.PathGenerated,
			}
			peer.Origin = "build-srcjar"
			peer.Confidence = "exact"
			peer.Exists = true
			peer.LocalPath = entry.LocalPath
			foundJava = true
			break
		}
	}
	if len(options.SourceArchives) > 0 && !foundJava {
		report.Diagnostics = append(
			report.Diagnostics,
			model.Diagnostic{
				Code:     "GENERATED_ENTRY_NOT_FOUND",
				Severity: model.SeverityWarning,
				Stage:    "aidl-generated",
				Message:  "generated Java entry was not found in source archives",
			},
		)
	}
	return report
}

func sameGeneratedSuffix(expected, actual string) bool {
	expected = strings.TrimPrefix(path.Clean(strings.ReplaceAll(expected, "\\", "/")), "./")
	actual = strings.TrimPrefix(path.Clean(strings.ReplaceAll(actual, "\\", "/")), "./")
	return actual == expected || strings.HasSuffix(actual, "/"+expected)
}

func interfaceBaseName(name string) string {
	if len(name) >= 2 && name[0] == 'I' && name[1] >= 'A' && name[1] <= 'Z' {
		return name[1:]
	}
	return name
}

func findDeclaration(declarations []Declaration, typeName string) (Declaration, string, bool, bool) {
	for index := range declarations {
		declaration := declarations[index]
		if declaration.CanonicalName == typeName || declaration.Name == typeName {
			return declaration, declaration.Name, false, true
		}
		if nested, ok := findNested(declaration, typeName); ok {
			return nested, declaration.Name, true, true
		}
	}
	return Declaration{}, "", false, false
}

func findNested(parent Declaration, typeName string) (Declaration, bool) {
	for index := range parent.Nested {
		nested := parent.Nested[index]
		if nested.CanonicalName == typeName || nested.Name == typeName {
			return nested, true
		}
		if found, ok := findNested(nested, typeName); ok {
			return found, ok
		}
	}
	return Declaration{}, false
}

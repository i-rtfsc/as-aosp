package aidl

import (
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

type SymbolKind string

const (
	SymbolInterface  SymbolKind = "interface"
	SymbolParcelable SymbolKind = "parcelable"
	SymbolEnum       SymbolKind = "enum"
	SymbolUnion      SymbolKind = "union"
	SymbolMethod     SymbolKind = "method"
	SymbolField      SymbolKind = "field"
	SymbolValue      SymbolKind = "value"
)

type Symbol struct {
	CanonicalName string         `json:"canonicalName"`
	Kind          SymbolKind     `json:"kind"`
	Path          model.RepoPath `json:"path"`
	Range         SourceRange    `json:"range"`
	Origin        string         `json:"origin"`
	Confidence    string         `json:"confidence"`
}

type Index struct {
	Documents map[model.RepoPath]Document `json:"documents"`
	Symbols   map[string][]Symbol         `json:"symbols"`
}

type Summary struct {
	Documents   int `json:"documents"`
	Symbols     int `json:"symbols"`
	Diagnostics int `json:"diagnostics"`
}

func (index Index) Summary(diagnostics int) Summary {
	return Summary{Documents: len(index.Documents), Symbols: len(index.Symbols), Diagnostics: diagnostics}
}

func NewIndex() Index {
	return Index{Documents: make(map[model.RepoPath]Document), Symbols: make(map[string][]Symbol)}
}

func (index *Index) Add(document Document) {
	index.Documents[document.Path] = document
	for _, declaration := range document.Declarations {
		index.addDeclaration(document.Path, declaration)
	}
}

func (index *Index) addDeclaration(path model.RepoPath, declaration Declaration) {
	index.add(Symbol{CanonicalName: declaration.CanonicalName, Kind: SymbolKind(declaration.Kind), Path: path, Range: declaration.Range, Origin: "parser", Confidence: "exact"})
	for _, member := range declaration.Members {
		index.add(Symbol{CanonicalName: declaration.CanonicalName + "." + member.Name, Kind: SymbolKind(member.Kind), Path: path, Range: member.Range, Origin: "parser", Confidence: "exact"})
	}
	for _, nested := range declaration.Nested {
		index.addDeclaration(path, nested)
	}
}

func (index *Index) Resolve(document Document, name string) Resolution {
	canonical := canonicalCandidates(document, name)
	for _, candidate := range canonical {
		if symbols := index.Symbols[candidate]; len(symbols) > 0 {
			if len(symbols) == 1 {
				return Resolution{Name: candidate, Symbols: symbols, State: ResolutionFound}
			}
			return Resolution{Name: candidate, Symbols: symbols, State: ResolutionAmbiguous}
		}
	}
	return Resolution{Name: name, State: ResolutionNotFound}
}

func (index *Index) Validate(document Document) []model.Diagnostic {
	diagnostics := make([]model.Diagnostic, 0)
	for _, item := range document.Imports {
		symbols := index.Symbols[item.Name]
		switch len(symbols) {
		case 0:
			diagnostics = append(diagnostics, model.Diagnostic{Code: "AIDL_IMPORT_NOT_FOUND", Severity: model.SeverityWarning, Stage: "aidl-resolver", Message: "import was not found: " + item.Name, RepoPath: document.Path})
		case 1:
		default:
			diagnostics = append(diagnostics, model.Diagnostic{Code: "AIDL_TYPE_AMBIGUOUS", Severity: model.SeverityWarning, Stage: "aidl-resolver", Message: "import resolves to multiple files: " + item.Name, RepoPath: document.Path})
		}
	}
	return diagnostics
}

type ResolutionState string

const (
	ResolutionFound     ResolutionState = "found"
	ResolutionAmbiguous ResolutionState = "ambiguous"
	ResolutionNotFound  ResolutionState = "not-found"
)

type Resolution struct {
	Name    string          `json:"name"`
	State   ResolutionState `json:"state"`
	Symbols []Symbol        `json:"symbols,omitempty"`
}

type Reference struct {
	Name       string         `json:"name"`
	Path       model.RepoPath `json:"path"`
	Range      SourceRange    `json:"range"`
	Origin     string         `json:"origin"`
	Confidence string         `json:"confidence"`
}

func (index *Index) References(name string) []Reference {
	result := make([]Reference, 0)
	for _, document := range index.Documents {
		for _, item := range document.Imports {
			if item.Name == name || strings.HasSuffix(item.Name, "."+name) {
				result = append(result, Reference{Name: name, Path: document.Path, Range: item.Range, Origin: "import", Confidence: "inferred"})
			}
		}
		for _, declaration := range document.Declarations {
			result = append(result, declarationReferences(document.Path, declaration, name)...)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Path != result[j].Path {
			return result[i].Path < result[j].Path
		}
		return result[i].Range.Start.Offset < result[j].Range.Start.Offset
	})
	return result
}

func declarationReferences(path model.RepoPath, declaration Declaration, name string) []Reference {
	result := make([]Reference, 0)
	add := func(typeName string, sourceRange SourceRange) {
		shortName := typeName
		if index := strings.LastIndexByte(shortName, '.'); index >= 0 {
			shortName = shortName[index+1:]
		}
		if typeName == name || shortName == name || strings.HasSuffix(typeName, "."+name) ||
			strings.HasSuffix(name, "."+shortName) {
			result = append(result, Reference{Name: name, Path: path, Range: sourceRange, Origin: "member", Confidence: "inferred"})
		}
	}
	for _, member := range declaration.Members {
		add(member.Type, member.Range)
		for _, parameter := range member.Parameters {
			add(parameter.Type, parameter.Range)
		}
	}
	for _, nested := range declaration.Nested {
		result = append(result, declarationReferences(path, nested, name)...)
	}
	return result
}

func (index *Index) add(symbol Symbol) {
	values := index.Symbols[symbol.CanonicalName]
	for _, existing := range values {
		if existing.Path == symbol.Path && existing.Range == symbol.Range {
			return
		}
	}
	index.Symbols[symbol.CanonicalName] = append(values, symbol)
	sort.Slice(index.Symbols[symbol.CanonicalName], func(i, j int) bool {
		return index.Symbols[symbol.CanonicalName][i].Path < index.Symbols[symbol.CanonicalName][j].Path
	})
}

func canonicalCandidates(document Document, name string) []string {
	name = strings.TrimSpace(strings.TrimPrefix(name, "."))
	if name == "" {
		return nil
	}
	seen := make(map[string]bool)
	result := make([]string, 0, len(document.Imports)+2)
	add := func(value string) {
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	if strings.Contains(name, ".") {
		add(name)
	} else {
		if document.Package != "" {
			add(document.Package + "." + name)
		}
		for _, item := range document.Imports {
			if strings.HasSuffix(item.Name, "."+name) || item.Name == name {
				add(item.Name)
			}
		}
	}
	return result
}

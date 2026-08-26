package artifacts

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/metadata"
	"github.com/i-rtfsc/as-aosp/core/model"
	"github.com/i-rtfsc/as-aosp/core/modules"
)

type MatchLevel string

const (
	MatchExact       MatchLevel = "exact"
	MatchCompatible  MatchLevel = "compatible"
	MatchApproximate MatchLevel = "approximate"
	MatchMissing     MatchLevel = "missing"
)

type Artifact struct {
	ID                string             `json:"id"`
	Kind              model.ArtifactKind `json:"kind"`
	Name              string             `json:"name"`
	Path              metadata.PathRef   `json:"path"`
	SHA256            string             `json:"sha256,omitempty"`
	Size              int64              `json:"size,omitempty"`
	OwnerModules      []string           `json:"ownerModules,omitempty"`
	SourceAttachments []metadata.PathRef `json:"sourceAttachments,omitempty"`
	GeneratedRoots    []metadata.PathRef `json:"generatedRoots,omitempty"`
}

type Catalog struct {
	Artifacts map[string]Artifact `json:"artifacts"`
}

type Resolution struct {
	Selector    model.ArtifactSelector `json:"selector"`
	Match       MatchLevel             `json:"match"`
	Artifacts   []Artifact             `json:"artifacts"`
	Diagnostics []model.Diagnostic     `json:"diagnostics"`
}

func Build(moduleCatalog modules.Catalog) Catalog {
	catalog := Catalog{Artifacts: make(map[string]Artifact)}
	for moduleName, module := range moduleCatalog.Modules {
		for _, ref := range module.Jars {
			catalog.add(moduleName, artifactFromRef(model.ArtifactJar, ref))
		}
		for _, ref := range module.SrcJars {
			catalog.add(moduleName, artifactFromRef(model.ArtifactSrcJar, ref))
		}
	}
	return catalog
}

func (catalog *Catalog) AddExternal(kind model.ArtifactKind, pathName string, owners []string) Artifact {
	ref := metadata.NewPathRef(pathName)
	artifact := artifactFromRef(kind, ref)
	if info, err := os.Stat(pathName); err == nil && !info.IsDir() {
		artifact.Size = info.Size()
		if data, readErr := os.ReadFile(pathName); readErr == nil {
			sum := sha256.Sum256(data)
			artifact.SHA256 = hex.EncodeToString(sum[:])
		}
	}
	for _, owner := range owners {
		catalog.add(owner, artifact)
	}
	if len(owners) == 0 {
		catalog.Artifacts[artifact.ID] = artifact
	}
	return catalog.Artifacts[artifact.ID]
}

func (catalog Catalog) Resolve(selector model.ArtifactSelector) Resolution {
	result := Resolution{Selector: selector, Match: MatchMissing, Artifacts: []Artifact{}, Diagnostics: []model.Diagnostic{}}
	candidates := make([]Artifact, 0)
	for _, artifact := range catalog.Artifacts {
		if artifact.Kind != selector.Kind {
			continue
		}
		if selector.SHA256 != "" && strings.EqualFold(selector.SHA256, artifact.SHA256) {
			candidates = append(candidates, artifact)
			continue
		}
		if selector.RelativePath != "" && normalizePath(selector.RelativePath) == normalizePath(artifact.Path.Raw) {
			candidates = append(candidates, artifact)
			continue
		}
		if selector.Name != "" && selector.Name == artifact.Name {
			candidates = append(candidates, artifact)
		}
	}
	sort.Slice(candidates, func(i, j int) bool { return candidates[i].ID < candidates[j].ID })
	result.Artifacts = candidates
	switch len(candidates) {
	case 1:
		result.Match = MatchExact
	case 0:
		result.Diagnostics = append(result.Diagnostics, artifactDiagnostic("ARTIFACT_NOT_FOUND", "artifact selector did not match"))
	default:
		result.Diagnostics = append(result.Diagnostics, artifactDiagnostic("ARTIFACT_AMBIGUOUS", "artifact selector matched multiple artifacts"))
	}
	return result
}

func (catalog *Catalog) add(owner string, artifact Artifact) {
	existing := catalog.Artifacts[artifact.ID]
	if existing.ID == "" {
		existing = artifact
	}
	existing.OwnerModules = uniqueStrings(append(existing.OwnerModules, owner))
	catalog.Artifacts[artifact.ID] = existing
}

func artifactFromRef(kind model.ArtifactKind, ref metadata.PathRef) Artifact {
	pathName := normalizePath(ref.Raw)
	return Artifact{ID: string(kind) + ":" + pathName, Kind: kind, Name: filepath.Base(pathName), Path: ref}
}

func normalizePath(value string) string {
	return strings.TrimPrefix(filepath.ToSlash(filepath.Clean(value)), "./")
}
func uniqueStrings(values []string) []string {
	seen := map[string]bool{}
	result := []string{}
	for _, value := range values {
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	sort.Strings(result)
	return result
}
func artifactDiagnostic(code, message string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityError, Stage: "artifact-resolver", Message: message}
}

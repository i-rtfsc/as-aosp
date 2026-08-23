package metadata

import (
	"path/filepath"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

type PathKind string

const (
	PathRepoRelative     PathKind = "repo-relative"
	PathExternalAbsolute PathKind = "external-absolute"
	PathGenerated        PathKind = "generated"
	PathUnknown          PathKind = "unknown"
)

type PathRef struct {
	Raw          string             `json:"raw"`
	Kind         PathKind           `json:"kind"`
	RepoPath     model.RepoPath     `json:"repoPath,omitempty"`
	ExternalPath model.ExternalPath `json:"externalPath,omitempty"`
}

func NewPathRef(raw string) PathRef {
	clean := strings.TrimSpace(strings.ReplaceAll(raw, "\\", "/"))
	if clean == "" {
		return PathRef{Raw: raw, Kind: PathUnknown}
	}
	if isAbsolute(clean) {
		return PathRef{Raw: raw, Kind: PathExternalAbsolute, ExternalPath: model.ExternalPath(clean)}
	}
	if repoPath, err := model.NormalizeRepoPath(clean); err == nil {
		if repoPath == "out" || strings.HasPrefix(string(repoPath), "out/") {
			return PathRef{Raw: raw, Kind: PathGenerated, RepoPath: repoPath}
		}
		return PathRef{Raw: raw, Kind: PathRepoRelative, RepoPath: repoPath}
	}
	return PathRef{Raw: raw, Kind: PathUnknown}
}

func isAbsolute(value string) bool {
	return filepath.IsAbs(value) || strings.HasPrefix(value, "//") ||
		(len(value) >= 2 && value[1] == ':')
}

type JavaModule struct {
	Name            string    `json:"name"`
	ModuleType      string    `json:"module_type,omitempty"`
	Paths           []PathRef `json:"paths,omitempty"`
	Sources         []PathRef `json:"sources,omitempty"`
	Generated       []PathRef `json:"generated,omitempty"`
	Dependencies    []string  `json:"dependencies,omitempty"`
	Jars            []PathRef `json:"jars,omitempty"`
	SrcJars         []PathRef `json:"srcjars,omitempty"`
	AidlIncludeDirs []PathRef `json:"aidlIncludeDirs,omitempty"`
	ResourceDirs    []PathRef `json:"resourceDirs,omitempty"`
	CompilerFlags   []string  `json:"compilerFlags,omitempty"`
}

type CCModule struct {
	Name           string    `json:"name"`
	Paths          []PathRef `json:"paths,omitempty"`
	Sources        []PathRef `json:"sources,omitempty"`
	IncludeDirs    []PathRef `json:"includeDirs,omitempty"`
	SystemIncludes []PathRef `json:"systemIncludes,omitempty"`
	Flags          []string  `json:"flags,omitempty"`
	Sysroot        *PathRef  `json:"sysroot,omitempty"`
	Compiler       string    `json:"compiler,omitempty"`
}

type ModuleInfo struct {
	Name         string    `json:"name"`
	Paths        []PathRef `json:"paths,omitempty"`
	Sources      []PathRef `json:"sources,omitempty"`
	Generated    []PathRef `json:"generated,omitempty"`
	Dependencies []string  `json:"dependencies,omitempty"`
	Classes      []string  `json:"classes,omitempty"`
	Tags         []string  `json:"tags,omitempty"`
	Installed    []PathRef `json:"installed,omitempty"`
}

type CompDBEntry struct {
	Directory string   `json:"directory"`
	Arguments []string `json:"arguments"`
	File      string   `json:"file"`
	Output    string   `json:"output,omitempty"`
}

type Snapshot struct {
	Java       map[string]JavaModule `json:"java,omitempty"`
	CC         map[string]CCModule   `json:"cc,omitempty"`
	ModuleInfo map[string]ModuleInfo `json:"moduleInfo,omitempty"`
	CompDB     []CompDBEntry         `json:"compdb,omitempty"`
}

type ReadReport struct {
	Snapshot    Snapshot           `json:"snapshot"`
	Diagnostics []model.Diagnostic `json:"diagnostics"`
}

type Summary struct {
	JavaModules   int `json:"javaModules"`
	CCModules     int `json:"ccModules"`
	ModuleInfo    int `json:"moduleInfoModules"`
	CompDBEntries int `json:"compdbEntries"`
	GraphModules  int `json:"graphModules"`
}

func (s Snapshot) Summary(graphModules int) Summary {
	return Summary{JavaModules: len(s.Java), CCModules: len(s.CC), ModuleInfo: len(s.ModuleInfo), CompDBEntries: len(s.CompDB), GraphModules: graphModules}
}

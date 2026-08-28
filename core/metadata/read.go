package metadata

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/i-rtfsc/as-aosp/core/model"
)

func LoadDirectory(root string) ReadReport {
	report := LoadModuleDirectory(root)
	readJSONCandidate(root, "compdb", []string{
		filepath.Join("development", "ide", "compdb", "compile_commands.json"),
		filepath.Join("soong", "development", "ide", "compdb", "compile_commands.json"),
		"compile_commands.json",
	}, &report, func(data []byte) error {
		return json.Unmarshal(data, &report.Snapshot.CompDB)
	})
	return report
}

func LoadModuleDirectory(root string) ReadReport {
	report := ReadReport{Snapshot: Snapshot{
		Java:       make(map[string]JavaModule),
		CC:         make(map[string]CCModule),
		ModuleInfo: make(map[string]ModuleInfo),
		CompDB:     make([]CompDBEntry, 0),
	}, Diagnostics: make([]model.Diagnostic, 0)}
	readJSONCandidate(root, "java-deps", []string{"module_bp_java_deps.json", filepath.Join("soong", "module_bp_java_deps.json")}, &report, func(data []byte) error {
		return decodeJavaDeps(data, &report.Snapshot.Java)
	})
	readJSONCandidate(root, "cc-deps", []string{"module_bp_cc_deps.json", filepath.Join("soong", "module_bp_cc_deps.json")}, &report, func(data []byte) error {
		return decodeCCDeps(data, &report.Snapshot.CC)
	})
	moduleInfoCandidates := []string{"module-info.json", filepath.Join("soong", "module-info.json")}
	moduleInfoCandidates = append(moduleInfoCandidates, globCandidates(root, filepath.Join("soong", "module-info-*.json"))...)
	moduleInfoCandidates = append(moduleInfoCandidates, globCandidates(root, filepath.Join("target", "product", "*", "module-info.json"))...)
	readJSONCandidate(root, "module-info", moduleInfoCandidates, &report, func(data []byte) error {
		return decodeModuleInfo(data, &report.Snapshot.ModuleInfo)
	})
	return report
}

func readJSONCandidate(root, kind string, candidates []string, report *ReadReport, decode func([]byte) error) {
	for _, candidate := range candidates {
		path := filepath.Join(root, candidate)
		data, err := os.ReadFile(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			report.Diagnostics = append(report.Diagnostics, diagnostic("METADATA_READ_FAILED", "metadata", fmt.Sprintf("read %s: %v", kind, err), path))
			return
		}
		if err := decode(data); err != nil {
			report.Diagnostics = append(report.Diagnostics, diagnostic("METADATA_PARSE_FAILED", "metadata", fmt.Sprintf("parse %s: %v", kind, err), path))
		}
		return
	}
	report.Diagnostics = append(report.Diagnostics, diagnostic("METADATA_FILE_MISSING", "metadata", fmt.Sprintf("%s metadata file was not found", kind), root))
}

func diagnostic(code, stage, message, localPath string) model.Diagnostic {
	return model.Diagnostic{Code: code, Severity: model.SeverityWarning, Stage: stage, Message: message, LocalPath: localPath}
}

type rawJavaInfo struct {
	ModuleType      string   `json:"module_type,omitempty"`
	Manifest        string   `json:"manifest,omitempty"`
	Deps            []string `json:"dependencies,omitempty"`
	Srcs            []string `json:"srcs,omitempty"`
	AidlIncludeDirs []string `json:"aidl_include_dirs,omitempty"`
	Jars            []string `json:"jars,omitempty"`
	SrcJars         []string `json:"srcjars,omitempty"`
	Paths           []string `json:"path,omitempty"`
	ResourceDirs    []string `json:"resource_dirs,omitempty"`
	Javacflags      []string `json:"javacflags,omitempty"`
	Kotlincflags    []string `json:"kotlincflags,omitempty"`
	Aidl            *struct {
		Srcs []string `json:"srcs,omitempty"`
	} `json:"aidl,omitempty"`
}

func decodeJavaDeps(data []byte, destination *map[string]JavaModule) error {
	var raw map[string]rawJavaInfo
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	for name, item := range raw {
		module := JavaModule{Name: name, ModuleType: item.ModuleType, Dependencies: uniqueStrings(item.Deps), CompilerFlags: uniqueStrings(append(item.Javacflags, item.Kotlincflags...))}
		module.Paths = pathRefs(item.Paths)
		module.Sources = pathRefs(item.Srcs)
		module.Jars = pathRefs(item.Jars)
		module.SrcJars = pathRefs(item.SrcJars)
		module.AidlIncludeDirs = pathRefs(item.AidlIncludeDirs)
		if item.Aidl != nil {
			module.AidlSources = pathRefs(item.Aidl.Srcs)
		}
		module.ResourceDirs = pathRefs(item.ResourceDirs)
		(*destination)[name] = module
	}
	return nil
}

type rawCCParams struct {
	HeaderSearchPath       []string `json:"header_search_path,omitempty"`
	SystemHeaderSearchPath []string `json:"system_search_path,omitempty"`
	FlagParameters         []string `json:"flag,omitempty"`
	SysRoot                string   `json:"system_root,omitempty"`
}

type rawCCInfo struct {
	Path               []string    `json:"path,omitempty"`
	Srcs               []string    `json:"srcs,omitempty"`
	GlobalCommonFlags  rawCCParams `json:"global_common_flags,omitempty"`
	LocalCommonFlags   rawCCParams `json:"local_common_flags,omitempty"`
	GlobalCFlags       rawCCParams `json:"global_c_flags,omitempty"`
	LocalCFlags        rawCCParams `json:"local_c_flags,omitempty"`
	GlobalCppFlags     rawCCParams `json:"global_cpp_flags,omitempty"`
	LocalCppFlags      rawCCParams `json:"local_cpp_flags,omitempty"`
	SystemIncludeFlags rawCCParams `json:"system_include_flags,omitempty"`
	ModuleName         string      `json:"module_name,omitempty"`
}

type rawCCDeps struct {
	Clang    string               `json:"clang,omitempty"`
	CppClang string               `json:"clang++,omitempty"`
	Modules  map[string]rawCCInfo `json:"modules,omitempty"`
}

func decodeCCDeps(data []byte, destination *map[string]CCModule) error {
	var raw rawCCDeps
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	for name, item := range raw.Modules {
		module := CCModule{Name: name, Paths: pathRefs(item.Path), Sources: pathRefs(item.Srcs)}
		parameters := []rawCCParams{item.GlobalCommonFlags, item.LocalCommonFlags, item.GlobalCFlags, item.LocalCFlags, item.GlobalCppFlags, item.LocalCppFlags}
		var includes, systemIncludes, flags []string
		for _, parameters := range parameters {
			includes = append(includes, parameters.HeaderSearchPath...)
			systemIncludes = append(systemIncludes, parameters.SystemHeaderSearchPath...)
			flags = append(flags, parameters.FlagParameters...)
		}
		module.IncludeDirs = pathRefs(includes)
		module.SystemIncludes = pathRefs(append(item.SystemIncludeFlags.HeaderSearchPath, item.SystemIncludeFlags.SystemHeaderSearchPath...))
		module.SystemIncludes = pathRefs(append(pathRefRaw(module.SystemIncludes), systemIncludes...))
		module.Flags = uniqueStrings(flags)
		if sysroot := firstNonEmpty(item.SystemIncludeFlags.SysRoot, item.GlobalCommonFlags.SysRoot); sysroot != "" {
			ref := NewPathRef(sysroot)
			module.Sysroot = &ref
		}
		module.Compiler = firstNonEmpty(raw.CppClang, raw.Clang)
		if item.ModuleName != "" {
			module.Name = item.ModuleName
		}
		(*destination)[module.Name] = module
	}
	return nil
}

type rawModuleInfo struct {
	Path         []string `json:"path,omitempty"`
	Installed    []string `json:"installed,omitempty"`
	Dependencies []string `json:"dependencies,omitempty"`
	Srcs         []string `json:"srcs,omitempty"`
	SrcJars      []string `json:"srcjars,omitempty"`
	Classes      []string `json:"class,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

func decodeModuleInfo(data []byte, destination *map[string]ModuleInfo) error {
	var raw map[string]rawModuleInfo
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	for name, item := range raw {
		(*destination)[name] = ModuleInfo{Name: name, Paths: pathRefs(item.Path), Installed: pathRefs(item.Installed), Sources: pathRefs(item.Srcs), Generated: pathRefs(item.SrcJars), Dependencies: uniqueStrings(item.Dependencies), Classes: uniqueStrings(item.Classes), Tags: uniqueStrings(item.Tags)}
	}
	return nil
}

func pathRefs(values []string) []PathRef {
	refs := make([]PathRef, 0, len(values))
	seen := make(map[string]bool)
	for _, value := range values {
		ref := NewPathRef(value)
		key := string(ref.Kind) + "\x00" + ref.Raw
		if !seen[key] {
			seen[key] = true
			refs = append(refs, ref)
		}
	}
	sort.Slice(refs, func(i, j int) bool { return refs[i].Raw < refs[j].Raw })
	return refs
}

func uniqueStrings(values []string) []string {
	seen := make(map[string]bool)
	result := make([]string, 0, len(values))
	for _, value := range values {
		if value != "" && !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	sort.Strings(result)
	return result
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func globCandidates(root, pattern string) []string {
	matches, err := filepath.Glob(filepath.Join(root, pattern))
	if err != nil {
		return nil
	}
	sort.Strings(matches)
	result := make([]string, 0, len(matches))
	for _, match := range matches {
		relative, err := filepath.Rel(root, match)
		if err == nil {
			result = append(result, relative)
		}
	}
	return result
}

func pathRefRaw(values []PathRef) []string {
	result := make([]string, 0, len(values))
	for _, value := range values {
		result = append(result, value.Raw)
	}
	return result
}

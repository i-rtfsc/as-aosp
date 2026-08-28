package sourcecatalog

import (
	"regexp"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

var makeModule = regexp.MustCompile(`(?m)^\s*LOCAL_MODULE\s*:?=\s*([^\s#]+)`)
var makeSources = regexp.MustCompile(`(?ms)^\s*LOCAL_SRC_FILES\s*:?=\s*(.*?)(?:\n\s*[A-Z_][A-Z0-9_]*\s*:?=|\z)`)
var makeLibraries = regexp.MustCompile(`(?ms)^\s*LOCAL_(?:JAVA_LIBRARIES|STATIC_JAVA_LIBRARIES|SHARED_LIBRARIES|STATIC_LIBRARIES)\s*:?=\s*(.*?)(?:\n\s*[A-Z_][A-Z0-9_]*\s*:?=|\z)`)

func parseMakefile(file, source string) []Module {
	matches := makeModule.FindAllStringSubmatchIndex(source, -1)
	result := make([]Module, 0, len(matches))
	for index, match := range matches {
		end := len(source)
		if index+1 < len(matches) {
			end = matches[index+1][0]
		}
		block := source[match[0]:end]
		name := strings.TrimSpace(source[match[2]:match[3]])
		if name == "" || strings.Contains(name, "$(") {
			continue
		}
		module := Module{
			Name: name, Type: makeModuleType(block), Blueprint: model.RepoPath(file),
			Directory: moduleDirectory(file), Status: StatusSourceOnly,
		}
		module.Sources = makeRepoPaths(makeValues(makeSources, block))
		module.Dependencies = makeValues(makeLibraries, block)
		module.Capabilities = capabilities(module)
		result = append(result, module)
	}
	return result
}

func makeModuleType(block string) string {
	switch {
	case strings.Contains(block, "BUILD_PACKAGE"):
		return "android_app"
	case strings.Contains(block, "BUILD_JAVA_LIBRARY"):
		return "java_library"
	case strings.Contains(block, "BUILD_SHARED_LIBRARY"):
		return "cc_library_shared"
	case strings.Contains(block, "BUILD_STATIC_LIBRARY"):
		return "cc_library_static"
	case strings.Contains(block, "BUILD_EXECUTABLE"):
		return "cc_binary"
	default:
		return "make_module"
	}
}

func makeValues(pattern *regexp.Regexp, block string) []string {
	match := pattern.FindStringSubmatch(block)
	if len(match) < 2 {
		return nil
	}
	clean := strings.ReplaceAll(match[1], "\\", " ")
	values := strings.Fields(clean)
	result := make([]string, 0, len(values))
	for _, value := range values {
		if !strings.HasPrefix(value, "$(") && !strings.HasPrefix(value, "#") {
			result = append(result, value)
		}
	}
	return result
}

func makeRepoPaths(values []string) []model.RepoPath {
	result := make([]model.RepoPath, 0, len(values))
	for _, value := range values {
		result = append(result, model.RepoPath(value))
	}
	return result
}

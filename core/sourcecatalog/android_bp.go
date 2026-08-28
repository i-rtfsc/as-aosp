package sourcecatalog

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

var moduleHeader = regexp.MustCompile(`(?m)([A-Za-z_][A-Za-z0-9_]*)\s*\{`)
var nameField = regexp.MustCompile(`(?m)\bname\s*:\s*"([^"]+)"`)
var listField = regexp.MustCompile(`(?m)\b(srcs|path|static_libs|libs|defaults|generated_sources|generated_headers)\s*:\s*\[([^\]]*)\]`)
var quoted = regexp.MustCompile(`"([^"]+)"`)

func parseBlueprint(file, source string) ([]Module, error) {
	if strings.TrimSpace(source) == "" {
		return nil, nil
	}
	result := []Module{}
	matches := moduleHeader.FindAllStringSubmatchIndex(source, -1)
	for _, match := range matches {
		typeName := source[match[2]:match[3]]
		if typeName == "package" || typeName == "soong_namespace" || typeName == "include" {
			continue
		}
		start := match[1]
		end := balancedObjectEnd(source, start-1)
		if end < 0 {
			return result, fmt.Errorf("unclosed module %q", typeName)
		}
		body := source[start:end]
		nameMatch := nameField.FindStringSubmatch(body)
		if len(nameMatch) == 0 {
			continue
		}
		module := Module{Name: nameMatch[1], Type: typeName, Blueprint: model.RepoPath(file), Directory: moduleDirectory(file), Status: StatusSourceOnly}
		for _, field := range listField.FindAllStringSubmatch(body, -1) {
			values := quoted.FindAllStringSubmatch(field[2], -1)
			for _, value := range values {
				switch field[1] {
				case "srcs":
					module.Sources = append(module.Sources, model.RepoPath(value[1]))
				case "path":
					module.SourceRoots = append(module.SourceRoots, model.RepoPath(value[1]))
				case "static_libs", "libs", "defaults", "generated_sources", "generated_headers":
					module.Dependencies = append(module.Dependencies, value[1])
				}
			}
		}
		module.Capabilities = capabilities(module)
		result = append(result, module)
	}
	return result, nil
}

func balancedObjectEnd(source string, open int) int {
	depth := 0
	for i := open; i < len(source); i++ {
		switch source[i] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return i
			}
		}
	}
	return -1
}

func moduleDirectory(file string) model.RepoPath {
	index := strings.LastIndex(file, "/")
	if index < 0 {
		return ""
	}
	return model.RepoPath(file[:index])
}

func capabilities(module Module) []string {
	values := []string{}
	lower := strings.ToLower(module.Type)
	if strings.Contains(lower, "java") || strings.Contains(lower, "android") {
		values = append(values, "java")
	}
	if strings.Contains(lower, "cc") || strings.Contains(lower, "rust") {
		values = append(values, "native")
	}
	if len(module.Dependencies) > 0 {
		values = append(values, "generated")
	}
	return values
}

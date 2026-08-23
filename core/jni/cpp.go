package jni

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/i-rtfsc/as-aosp/core/model"
)

var staticExportPattern = regexp.MustCompile(
	`\bJava_([A-Za-z0-9_]+)\s*\(`,
)
var nativeTablePattern = regexp.MustCompile(
	`(?s)(?:static\s+)?(?:const\s+)?JNINativeMethod\s+` +
		`([A-Za-z_][A-Za-z0-9_]*)\s*\[\s*\]\s*=\s*\{(.*?)\};`,
)
var nativeEntryPattern = regexp.MustCompile(
	`\{\s*"([^"]+)"\s*,\s*"([^"]*)"\s*,\s*` +
		`\(\s*void\s*\*\s*\)\s*([A-Za-z_][A-Za-z0-9_]*)`,
)
var nativeMacroPattern = regexp.MustCompile(
	`MAKE_JNI_NATIVE_METHOD\s*\(\s*"([^"]+)"\s*,\s*` +
		`"([^"]*)"\s*,\s*([A-Za-z_][A-Za-z0-9_]*)\s*\)`,
)
var registrationPattern = regexp.MustCompile(
	`(?:jniRegisterNativeMethods|RegisterMethodsOrDie|` +
		`(?:AndroidRuntime::)?registerNativeMethods)\s*\(` +
		`[^,]+,\s*"([^"]+)"\s*,\s*` +
		`([A-Za-z_][A-Za-z0-9_]*)`,
)

type tableEntry struct {
	MethodName   string
	Descriptor   string
	FunctionName string
	Range        Range
}

func ScanCpp(
	path model.RepoPath,
	source string,
) ([]NativeSymbol, []Registration) {
	natives := scanStaticExports(path, source)
	tables := scanNativeTables(path, source, &natives)
	registrations := make([]Registration, 0)
	for _, match := range registrationPattern.FindAllStringSubmatchIndex(
		source,
		-1,
	) {
		className := strings.ReplaceAll(source[match[2]:match[3]], "/", ".")
		tableName := source[match[4]:match[5]]
		entries := tables[tableName]
		if len(entries) == 0 {
			registrations = append(registrations, Registration{
				ClassName:    className,
				FunctionName: tableName,
				Path:         path,
				Range:        Range{Start: match[0], End: match[1]},
				Origin:       "registration-call",
				Confidence:   ConfidenceUnresolved,
			})
			continue
		}
		for _, entry := range entries {
			registrations = append(registrations, Registration{
				ClassName:    className,
				MethodName:   entry.MethodName,
				Descriptor:   entry.Descriptor,
				FunctionName: entry.FunctionName,
				Path:         path,
				Range:        entry.Range,
				Origin:       "native-method-table",
				Confidence:   descriptorConfidence(entry.Descriptor),
			})
		}
	}
	return natives, registrations
}

func scanStaticExports(
	path model.RepoPath,
	source string,
) []NativeSymbol {
	result := make([]NativeSymbol, 0)
	for _, match := range staticExportPattern.FindAllStringSubmatchIndex(
		source,
		-1,
	) {
		encoded := source[match[2]:match[3]]
		className, method, descriptor, ok := decodeExport(encoded)
		result = append(result, NativeSymbol{
			FunctionName: "Java_" + encoded,
			ClassName:    className,
			MethodName:   method,
			Descriptor:   descriptor,
			Path:         path,
			Range:        Range{Start: match[0], End: match[1]},
			Origin:       "static-export",
			Confidence:   confidence(ok),
		})
	}
	return result
}

func scanNativeTables(
	path model.RepoPath,
	source string,
	natives *[]NativeSymbol,
) map[string][]tableEntry {
	result := make(map[string][]tableEntry)
	for _, table := range nativeTablePattern.FindAllStringSubmatchIndex(
		source,
		-1,
	) {
		name := source[table[2]:table[3]]
		bodyStart := table[4]
		body := source[table[4]:table[5]]
		entries := parseTableEntries(body, bodyStart)
		result[name] = entries
		for _, entry := range entries {
			*natives = append(*natives, NativeSymbol{
				FunctionName: entry.FunctionName,
				MethodName:   entry.MethodName,
				Descriptor:   entry.Descriptor,
				Path:         path,
				Range:        entry.Range,
				Origin:       "native-method-table",
				Confidence:   descriptorConfidence(entry.Descriptor),
			})
		}
	}
	return result
}

func parseTableEntries(body string, offset int) []tableEntry {
	result := make([]tableEntry, 0)
	patterns := []*regexp.Regexp{nativeEntryPattern, nativeMacroPattern}
	for _, pattern := range patterns {
		for _, match := range pattern.FindAllStringSubmatchIndex(body, -1) {
			result = append(result, tableEntry{
				MethodName:   body[match[2]:match[3]],
				Descriptor:   body[match[4]:match[5]],
				FunctionName: body[match[6]:match[7]],
				Range:        Range{Start: offset + match[0], End: offset + match[1]},
			})
		}
	}
	return result
}

func descriptorConfidence(value string) Confidence {
	if _, ok := NormalizeDescriptor(value); ok {
		return ConfidenceExact
	}
	return ConfidenceInferred
}

func decodeExport(encoded string) (string, string, string, bool) {
	parts := strings.Split(encoded, "__")
	segments := splitJNIComponents(parts[0])
	if len(segments) < 3 {
		return "", "", "", false
	}
	method := segments[len(segments)-1]
	class := segments[len(segments)-2]
	packageParts := segments[:len(segments)-2]
	return strings.Join(packageParts, ".") + "." + class,
		method,
		"",
		len(parts) == 1
}

func splitJNIComponents(value string) []string {
	components := []string{""}
	for index := 0; index < len(value); {
		if value[index] != '_' {
			components[len(components)-1] += string(value[index])
			index++
			continue
		}
		if index+1 >= len(value) {
			components = append(components, "")
			break
		}
		switch value[index+1] {
		case '1':
			components[len(components)-1] += "_"
			index += 2
		case '2':
			components[len(components)-1] += ";"
			index += 2
		case '3':
			components[len(components)-1] += "["
			index += 2
		case '0':
			if index+6 <= len(value) {
				code, err := strconv.ParseInt(value[index+2:index+6], 16, 32)
				if err == nil && utf8.ValidRune(rune(code)) {
					components[len(components)-1] += string(rune(code))
					index += 6
					continue
				}
			}
			components[len(components)-1] += "_"
			index++
		default:
			components = append(components, "")
			index++
		}
	}
	return components
}

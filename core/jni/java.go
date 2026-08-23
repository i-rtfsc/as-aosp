package jni

import (
	"regexp"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

var packagePattern = regexp.MustCompile(
	`(?m)\bpackage\s+([A-Za-z_$][\w$]*(?:\.[A-Za-z_$][\w$]*)*)\s*;?`,
)
var importPattern = regexp.MustCompile(
	`(?m)\bimport\s+([A-Za-z_$][\w$]*(?:\.[A-Za-z_$][\w$]*)*)\s*;?`,
)
var classPattern = regexp.MustCompile(
	`\b(?:class|interface|object|enum)\s+([A-Za-z_$][\w$]*)`,
)
var nativeJavaPattern = regexp.MustCompile(
	`(?m)\bnative\s+([A-Za-z_$][\w$<>\[\].?]*)\s+` +
		`([A-Za-z_$][\w$]*)\s*\(([^)]*)\)\s*;`,
)
var externalKotlinPattern = regexp.MustCompile(
	`(?m)\bexternal\s+fun\s+([A-Za-z_$][\w$]*)\s*` +
		`\(([^)]*)\)\s*(?::\s*([A-Za-z_$][\w$<>\[\].?]*))?`,
)

func ScanJava(path model.RepoPath, source string) []Symbol {
	packageName := firstGroup(packagePattern, source)
	spans := scanClassSpans(source)
	imports := scanImports(source)
	result := make([]Symbol, 0)
	for _, match := range nativeJavaPattern.FindAllStringSubmatchIndex(
		source,
		-1,
	) {
		returnType := source[match[2]:match[3]]
		method := source[match[4]:match[5]]
		parameters := source[match[6]:match[7]]
		descriptor, exact := javaDescriptor(
			parameters,
			returnType,
			packageName,
			imports,
		)
		result = append(result, Symbol{
			ClassName:  joinClass(packageName, classNameAt(spans, match[0])),
			MethodName: method,
			Descriptor: descriptor,
			Path:       path,
			Range:      Range{Start: match[0], End: match[1]},
			Origin:     "java-native-declaration",
			Confidence: confidence(exact),
		})
	}
	for _, match := range externalKotlinPattern.FindAllStringSubmatchIndex(
		source,
		-1,
	) {
		method := source[match[2]:match[3]]
		parameters := source[match[4]:match[5]]
		returnType := "Unit"
		if match[6] >= 0 {
			returnType = source[match[6]:match[7]]
		}
		descriptor, exact := kotlinDescriptor(
			parameters,
			returnType,
			packageName,
			imports,
		)
		result = append(result, Symbol{
			ClassName:  joinClass(packageName, classNameAt(spans, match[0])),
			MethodName: method,
			Descriptor: descriptor,
			Path:       path,
			Range:      Range{Start: match[0], End: match[1]},
			Origin:     "kotlin-external-declaration",
			Confidence: confidence(exact),
		})
	}
	return result
}

func javaDescriptor(
	parameters string,
	returnType string,
	packageName string,
	imports map[string]string,
) (string, bool) {
	types := make([]string, 0)
	for _, parameter := range splitParameters(parameters) {
		fields := strings.Fields(stripAnnotations(parameter))
		if len(fields) < 2 {
			return "", false
		}
		types = append(types, strings.Join(fields[:len(fields)-1], ""))
	}
	return buildDescriptor(types, returnType, packageName, imports, false)
}

func kotlinDescriptor(
	parameters string,
	returnType string,
	packageName string,
	imports map[string]string,
) (string, bool) {
	types := make([]string, 0)
	for _, parameter := range splitParameters(parameters) {
		parts := strings.SplitN(stripAnnotations(parameter), ":", 2)
		if len(parts) != 2 {
			return "", false
		}
		types = append(types, strings.TrimSpace(parts[1]))
	}
	return buildDescriptor(types, returnType, packageName, imports, true)
}

func buildDescriptor(
	parameters []string,
	returnType string,
	packageName string,
	imports map[string]string,
	kotlin bool,
) (string, bool) {
	var builder strings.Builder
	builder.WriteByte('(')
	exact := true
	for _, value := range parameters {
		descriptor, ok := sourceTypeDescriptor(
			value,
			packageName,
			imports,
			kotlin,
		)
		if !ok {
			exact = false
		}
		builder.WriteString(descriptor)
	}
	builder.WriteByte(')')
	result, ok := sourceTypeDescriptor(
		returnType,
		packageName,
		imports,
		kotlin,
	)
	if !ok {
		exact = false
	}
	builder.WriteString(result)
	return builder.String(), exact
}

func sourceTypeDescriptor(
	value string,
	packageName string,
	imports map[string]string,
	kotlin bool,
) (string, bool) {
	value = strings.TrimSpace(strings.TrimSuffix(value, "?"))
	arrayDepth := 0
	for strings.HasSuffix(value, "[]") {
		arrayDepth++
		value = strings.TrimSuffix(value, "[]")
	}
	if kotlin && strings.HasPrefix(value, "Array<") &&
		strings.HasSuffix(value, ">") {
		arrayDepth++
		value = strings.TrimSuffix(strings.TrimPrefix(value, "Array<"), ">")
	}
	value = strings.TrimSpace(strings.SplitN(value, "<", 2)[0])
	primitive, ok := primitiveSourceDescriptor(value, kotlin)
	if ok {
		return strings.Repeat("[", arrayDepth) + primitive, true
	}
	fqcn, exact := resolveSourceType(value, packageName, imports)
	return strings.Repeat("[", arrayDepth) + "L" +
		strings.ReplaceAll(fqcn, ".", "/") + ";", exact
}

func primitiveSourceDescriptor(value string, kotlin bool) (string, bool) {
	if kotlin {
		switch value {
		case "Unit":
			return "V", true
		case "Boolean":
			return "Z", true
		case "Byte":
			return "B", true
		case "Char":
			return "C", true
		case "Short":
			return "S", true
		case "Int":
			return "I", true
		case "Long":
			return "J", true
		case "Float":
			return "F", true
		case "Double":
			return "D", true
		case "String":
			return "Ljava/lang/String;", true
		case "BooleanArray":
			return "[Z", true
		case "ByteArray":
			return "[B", true
		case "CharArray":
			return "[C", true
		case "ShortArray":
			return "[S", true
		case "IntArray":
			return "[I", true
		case "LongArray":
			return "[J", true
		case "FloatArray":
			return "[F", true
		case "DoubleArray":
			return "[D", true
		}
	}
	switch value {
	case "void":
		return "V", true
	case "boolean":
		return "Z", true
	case "byte":
		return "B", true
	case "char":
		return "C", true
	case "short":
		return "S", true
	case "int":
		return "I", true
	case "long":
		return "J", true
	case "float":
		return "F", true
	case "double":
		return "D", true
	case "String":
		return "Ljava/lang/String;", true
	default:
		return "", false
	}
}

func resolveSourceType(
	value string,
	packageName string,
	imports map[string]string,
) (string, bool) {
	if strings.Contains(value, ".") {
		return value, true
	}
	if imported := imports[value]; imported != "" {
		return imported, true
	}
	if packageName != "" {
		return packageName + "." + value, false
	}
	return value, false
}

func scanImports(source string) map[string]string {
	result := make(map[string]string)
	for _, match := range importPattern.FindAllStringSubmatch(source, -1) {
		parts := strings.Split(match[1], ".")
		result[parts[len(parts)-1]] = match[1]
	}
	return result
}

func firstGroup(pattern *regexp.Regexp, source string) string {
	match := pattern.FindStringSubmatch(source)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func splitParameters(value string) []string {
	value = strings.TrimSpace(value)
	if value == "" {
		return nil
	}
	var result []string
	depth := 0
	start := 0
	for index, ch := range value {
		switch ch {
		case '<', '(', '[':
			depth++
		case '>', ')', ']':
			if depth > 0 {
				depth--
			}
		case ',':
			if depth == 0 {
				result = append(
					result,
					strings.TrimSpace(value[start:index]),
				)
				start = index + 1
			}
		}
	}
	result = append(result, strings.TrimSpace(value[start:]))
	return result
}

func stripAnnotations(value string) string {
	fields := strings.Fields(value)
	result := make([]string, 0, len(fields))
	for _, field := range fields {
		if !strings.HasPrefix(field, "@") && field != "final" {
			result = append(result, field)
		}
	}
	return strings.Join(result, " ")
}

func joinClass(packageName, className string) string {
	if packageName == "" {
		return className
	}
	return packageName + "." + className
}

func confidence(exact bool) Confidence {
	if exact {
		return ConfidenceExact
	}
	return ConfidenceInferred
}

type classSpan struct {
	Name  string
	Start int
	End   int
}

func scanClassSpans(source string) []classSpan {
	result := make([]classSpan, 0)
	for _, match := range classPattern.FindAllStringSubmatchIndex(source, -1) {
		open := strings.IndexByte(source[match[1]:], '{')
		if open < 0 {
			continue
		}
		open += match[1]
		close := matchingBrace(source, open)
		if close < 0 {
			continue
		}
		result = append(result, classSpan{
			Name:  source[match[2]:match[3]],
			Start: open,
			End:   close,
		})
	}
	return result
}

func classNameAt(spans []classSpan, offset int) string {
	parts := make([]string, 0)
	for _, span := range spans {
		if offset >= span.Start && offset <= span.End {
			parts = append(parts, span.Name)
		}
	}
	return strings.Join(parts, "$")
}

func matchingBrace(source string, open int) int {
	depth := 0
	inString := byte(0)
	inLineComment := false
	inBlockComment := false
	for index := open; index < len(source); index++ {
		ch := source[index]
		next := byte(0)
		if index+1 < len(source) {
			next = source[index+1]
		}
		if inLineComment {
			if ch == '\n' {
				inLineComment = false
			}
			continue
		}
		if inBlockComment {
			if ch == '*' && next == '/' {
				inBlockComment = false
				index++
			}
			continue
		}
		if inString != 0 {
			if ch == '\\' {
				index++
				continue
			}
			if ch == inString {
				inString = 0
			}
			continue
		}
		if ch == '/' && next == '/' {
			inLineComment = true
			index++
			continue
		}
		if ch == '/' && next == '*' {
			inBlockComment = true
			index++
			continue
		}
		if ch == '"' || ch == '\'' {
			inString = ch
			continue
		}
		switch ch {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return index
			}
		}
	}
	return -1
}

package jni

import "strings"

func NormalizeDescriptor(value string) (string, bool) {
	value = strings.TrimSpace(value)
	if value == "" || value[0] != '(' {
		return value, false
	}
	close := strings.IndexByte(value, ')')
	if close < 1 || close == len(value)-1 {
		return value, false
	}
	if !validDescriptorList(value[1:close], false) {
		return value, false
	}
	if !validDescriptorList(value[close+1:], true) {
		return value, false
	}
	return value, true
}

func validDescriptorList(value string, allowVoid bool) bool {
	count := 0
	for index := 0; index < len(value); {
		next, ok := descriptorTypeEnd(value, index, allowVoid)
		if !ok {
			return false
		}
		index = next
		count++
	}
	if allowVoid {
		return count == 1
	}
	return true
}

func descriptorTypeEnd(
	value string,
	start int,
	allowVoid bool,
) (int, bool) {
	if start >= len(value) {
		return start, false
	}
	index := start
	for index < len(value) && value[index] == '[' {
		index++
	}
	if index >= len(value) {
		return start, false
	}
	switch value[index] {
	case 'Z', 'B', 'C', 'S', 'I', 'J', 'F', 'D':
		return index + 1, true
	case 'V':
		return index + 1, allowVoid && index == start
	case 'L':
		end := strings.IndexByte(value[index:], ';')
		if end < 2 {
			return start, false
		}
		return index + end + 1, true
	default:
		return start, false
	}
}

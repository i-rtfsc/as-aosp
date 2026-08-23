package jni

import (
	"sort"
	"strings"

	"github.com/i-rtfsc/as-aosp/core/model"
)

func BuildIndex(
	javaFiles map[model.RepoPath]string,
	cppFiles map[model.RepoPath]string,
) Index {
	index := NewIndex()
	for path, source := range javaFiles {
		index.JavaSymbols = append(index.JavaSymbols, ScanJava(path, source)...)
	}
	for path, source := range cppFiles {
		natives, registrations := ScanCpp(path, source)
		index.NativeSymbols = append(index.NativeSymbols, natives...)
		index.Registrations = append(index.Registrations, registrations...)
	}
	index.sort()
	index.link()
	return index
}

func (index *Index) sort() {
	sort.Slice(index.JavaSymbols, func(i, j int) bool {
		return symbolKey(index.JavaSymbols[i]) < symbolKey(index.JavaSymbols[j])
	})
	sort.Slice(index.NativeSymbols, func(i, j int) bool {
		return index.NativeSymbols[i].FunctionName < index.NativeSymbols[j].FunctionName
	})
	sort.Slice(index.Registrations, func(i, j int) bool {
		return registrationKey(index.Registrations[i]) < registrationKey(index.Registrations[j])
	})
}

func (index *Index) link() {
	for _, registration := range index.Registrations {
		for _, java := range index.JavaSymbols {
			if java.ClassName != registration.ClassName ||
				java.MethodName != registration.MethodName {
				continue
			}
			confidence := ConfidenceInferred
			if registration.Confidence == ConfidenceExact &&
				java.Confidence == ConfidenceExact &&
				java.Descriptor != "" &&
				java.Descriptor == registration.Descriptor {
				confidence = ConfidenceExact
			}
			index.Links = append(index.Links, Link{
				Java: registrationToJava(java, confidence),
				Native: NativeSymbol{
					FunctionName: registration.FunctionName,
					ClassName:    registration.ClassName,
					MethodName:   registration.MethodName,
					Descriptor:   registration.Descriptor,
					Path:         registration.Path,
					Range:        registration.Range,
					Origin:       registration.Origin,
					Confidence:   confidence,
				},
				Registration: &registration,
				Confidence:   confidence,
			})
		}
	}
	for _, native := range index.NativeSymbols {
		if native.ClassName == "" || native.MethodName == "" {
			continue
		}
		for _, java := range index.JavaSymbols {
			if java.ClassName == native.ClassName && java.MethodName == native.MethodName {
				confidence := ConfidenceInferred
				if native.Confidence == ConfidenceExact && java.Confidence == ConfidenceExact {
					confidence = ConfidenceExact
				}
				index.Links = append(index.Links, Link{Java: java, Native: native, Confidence: confidence})
			}
		}
	}
	sort.Slice(index.Links, func(i, j int) bool {
		return symbolKey(index.Links[i].Java)+"\x00"+index.Links[i].Native.FunctionName < symbolKey(index.Links[j].Java)+"\x00"+index.Links[j].Native.FunctionName
	})
}

func registrationToJava(java Symbol, confidence Confidence) Symbol {
	java.Confidence = confidence
	return java
}

func symbolKey(symbol Symbol) string {
	return symbol.ClassName + "\x00" + symbol.MethodName + "\x00" + symbol.Descriptor + "\x00" + string(symbol.Path)
}

func registrationKey(registration Registration) string {
	return registration.ClassName + "\x00" + registration.MethodName + "\x00" + registration.FunctionName + "\x00" + string(registration.Path)
}

func (index Index) JavaPeers(className, methodName string) []Link {
	result := make([]Link, 0)
	for _, link := range index.Links {
		if link.Java.ClassName == className && link.Java.MethodName == methodName {
			result = append(result, link)
		}
	}
	return result
}

func (index Index) NativePeers(functionName string) []Link {
	result := make([]Link, 0)
	for _, link := range index.Links {
		if strings.TrimSpace(link.Native.FunctionName) == functionName {
			result = append(result, link)
		}
	}
	return result
}

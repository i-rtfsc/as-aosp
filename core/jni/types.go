package jni

import "github.com/i-rtfsc/as-aosp/core/model"

type Confidence string

const (
	ConfidenceExact      Confidence = "exact"
	ConfidenceInferred   Confidence = "inferred"
	ConfidenceUnresolved Confidence = "unresolved"
)

type Symbol struct {
	ClassName  string         `json:"className"`
	MethodName string         `json:"methodName"`
	Descriptor string         `json:"descriptor,omitempty"`
	Path       model.RepoPath `json:"path"`
	Range      Range          `json:"range"`
	Origin     string         `json:"origin"`
	Confidence Confidence     `json:"confidence"`
}

type NativeSymbol struct {
	FunctionName string         `json:"functionName"`
	ClassName    string         `json:"className,omitempty"`
	MethodName   string         `json:"methodName,omitempty"`
	Descriptor   string         `json:"descriptor,omitempty"`
	Path         model.RepoPath `json:"path"`
	Range        Range          `json:"range"`
	Origin       string         `json:"origin"`
	Confidence   Confidence     `json:"confidence"`
}

type Registration struct {
	ClassName    string         `json:"className"`
	MethodName   string         `json:"methodName"`
	Descriptor   string         `json:"descriptor,omitempty"`
	FunctionName string         `json:"functionName"`
	Path         model.RepoPath `json:"path"`
	Range        Range          `json:"range"`
	Origin       string         `json:"origin"`
	Confidence   Confidence     `json:"confidence"`
}

type Range struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

type Link struct {
	Java         Symbol        `json:"java"`
	Native       NativeSymbol  `json:"native"`
	Registration *Registration `json:"registration,omitempty"`
	Confidence   Confidence    `json:"confidence"`
}

type Index struct {
	JavaSymbols   []Symbol           `json:"javaSymbols"`
	NativeSymbols []NativeSymbol     `json:"nativeSymbols"`
	Registrations []Registration     `json:"registrations"`
	Links         []Link             `json:"links"`
	Diagnostics   []model.Diagnostic `json:"diagnostics"`
}

func NewIndex() Index {
	return Index{
		JavaSymbols:   make([]Symbol, 0),
		NativeSymbols: make([]NativeSymbol, 0),
		Registrations: make([]Registration, 0),
		Links:         make([]Link, 0),
		Diagnostics:   make([]model.Diagnostic, 0),
	}
}

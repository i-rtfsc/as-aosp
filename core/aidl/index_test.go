package aidl

import "testing"

func TestIndexResolvesPackageAndImportTypes(t *testing.T) {
	index := NewIndex()
	bar := Parse("android/other/IBar.aidl", "package android.other; interface IBar {}")
	foo := Parse("android/demo/IFoo.aidl", "package android.demo; import android.other.IBar; interface IFoo { IBar get(); }")
	index.Add(bar)
	index.Add(foo)
	resolution := index.Resolve(foo, "IBar")
	if resolution.State != ResolutionFound || len(resolution.Symbols) != 1 || resolution.Symbols[0].Path != "android/other/IBar.aidl" {
		t.Fatalf("resolution: %+v", resolution)
	}
}

func TestIndexReportsAmbiguousType(t *testing.T) {
	index := NewIndex()
	first := Parse("one/Foo.aidl", "package demo; interface Foo {}")
	second := Parse("two/Foo.aidl", "package demo; interface Foo {}")
	use := Parse("use/Use.aidl", "package use; import demo.Foo; interface Use {}")
	index.Add(first)
	index.Add(second)
	index.Add(use)
	resolution := index.Resolve(use, "demo.Foo")
	if resolution.State != ResolutionAmbiguous || len(resolution.Symbols) != 2 {
		t.Fatalf("resolution: %+v", resolution)
	}
}

func TestIndexAddsNestedDeclarationSymbols(t *testing.T) {
	index := NewIndex()
	document := Parse("android/demo/Outer.aidl", "package android.demo; parcelable Outer { enum Kind { ONE } }")
	index.Add(document)
	resolution := index.Resolve(document, "android.demo.Outer.Kind")
	if resolution.State != ResolutionFound || resolution.Symbols[0].Kind != SymbolEnum {
		t.Fatalf("resolution: %+v", resolution)
	}
}

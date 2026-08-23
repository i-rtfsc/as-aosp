package aidl

import "testing"

func TestParseDeclarationsAndMembers(t *testing.T) {
	document := Parse("android/demo/IFoo.aidl", `package android.demo;
import android.other.IBar;
interface IFoo {
  void ping(in int value);
  IBar getBar();
}
parcelable Data { int value; }
enum Mode { ONE, TWO }
union Payload { int number; String text; }
`)
	if len(document.Diagnostics) != 0 {
		t.Fatalf("diagnostics: %+v", document.Diagnostics)
	}
	if document.Package != "android.demo" || len(document.Imports) != 1 || len(document.Declarations) != 4 {
		t.Fatalf("document: %+v", document)
	}
	if document.Declarations[0].CanonicalName != "android.demo.IFoo" || len(document.Declarations[0].Members) != 2 {
		t.Fatalf("interface: %+v", document.Declarations[0])
	}
	if document.Declarations[0].Members[0].Parameters[0].Direction != "in" {
		t.Fatalf("parameter: %+v", document.Declarations[0].Members[0])
	}
}

func TestParseReportsMissingPackage(t *testing.T) {
	document := Parse("Foo.aidl", "interface Foo {}")
	for _, diagnostic := range document.Diagnostics {
		if diagnostic.Code == "AIDL_PACKAGE_MISSING" {
			return
		}
	}
	t.Fatalf("expected package diagnostic: %+v", document.Diagnostics)
}

func TestParseOnewayInterface(t *testing.T) {
	document := Parse("demo/IFoo.aidl", "package demo; oneway interface IFoo { void ping(); }")
	if len(document.Diagnostics) != 0 || len(document.Declarations) != 1 || document.Declarations[0].Name != "IFoo" {
		t.Fatalf("document: %+v", document)
	}
}

func TestParseNestedDeclarations(t *testing.T) {
	document := Parse("android/demo/Outer.aidl", `package android.demo;
parcelable Outer {
  @Backing(type="int")
  enum Kind { ONE, TWO }
  union Value { int number; String text; }
}
`)
	if len(document.Diagnostics) != 0 || len(document.Declarations) != 1 || len(document.Declarations[0].Nested) != 2 {
		t.Fatalf("document: %+v", document)
	}
	if document.Declarations[0].Nested[0].CanonicalName != "android.demo.Outer.Kind" {
		t.Fatalf("nested: %+v", document.Declarations[0].Nested)
	}
}

func TestParseQualifiedParcelableAndAnnotatedParameter(t *testing.T) {
	document := Parse("android/os/Debug.aidl", `package android.os;
parcelable Debug.MemoryInfo;
interface Factory { void create(in @nullable Debug.MemoryInfo value) = 1; }
interface Constants { const @utf8InCpp String NAME = "demo"; }
`)
	if len(document.Diagnostics) != 0 || len(document.Declarations) != 3 {
		t.Fatalf("document: %+v", document)
	}
	if document.Declarations[0].CanonicalName != "android.os.Debug.MemoryInfo" {
		t.Fatalf("declaration: %+v", document.Declarations[0])
	}
}

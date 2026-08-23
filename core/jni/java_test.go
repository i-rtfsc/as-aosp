package jni

import "testing"

func TestScanJavaNativeDeclarations(t *testing.T) {
	source := `package android.security;
class Scrypt {
  private static native byte[] nativeScrypt(
      byte[] password, byte[] salt, int n, int r, int p, int keyLen);
}`
	symbols := ScanJava("android/security/Scrypt.java", source)
	if len(symbols) != 1 {
		t.Fatalf("symbols: %+v", symbols)
	}
	symbol := symbols[0]
	if symbol.ClassName != "android.security.Scrypt" ||
		symbol.MethodName != "nativeScrypt" ||
		symbol.Descriptor != "([B[BIIII)[B" ||
		symbol.Confidence != ConfidenceExact {
		t.Fatalf("symbol: %+v", symbol)
	}
}

func TestScanKotlinExternalDeclaration(t *testing.T) {
	source := `package demo
object NativeBridge {
  external fun add(left: Int, right: Int): Int
}`
	symbols := ScanJava("demo/NativeBridge.kt", source)
	if len(symbols) != 1 ||
		symbols[0].ClassName != "demo.NativeBridge" ||
		symbols[0].Descriptor != "(II)I" {
		t.Fatalf("symbols: %+v", symbols)
	}
}

func TestScanKotlinPrimitiveArrays(t *testing.T) {
	source := `package demo
object NativeBridge {
  external fun transform(input: ByteArray, values: IntArray): LongArray
}`
	symbols := ScanJava("demo/NativeBridge.kt", source)
	if len(symbols) != 1 || symbols[0].Descriptor != "([B[I)[J" {
		t.Fatalf("symbols: %+v", symbols)
	}
}

func TestScanJavaUsesNestedClassName(t *testing.T) {
	source := `package demo;
class Outer {
  class Inner {
    private native void close();
  }
}`
	symbols := ScanJava("demo/Outer.java", source)
	if len(symbols) != 1 || symbols[0].ClassName != "demo.Outer$Inner" {
		t.Fatalf("symbols: %+v", symbols)
	}
}

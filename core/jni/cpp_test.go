package jni

import "testing"

func TestScanCppDynamicRegistration(t *testing.T) {
	source := `
static const JNINativeMethod sMethods[] = {
  {"nativeScrypt", "([B[BIIII)[B", (void*)nativeScrypt},
};
int registerMethods(JNIEnv* env) {
  return jniRegisterNativeMethods(
      env, "android/security/Scrypt", sMethods, 1);
}`
	natives, registrations := ScanCpp("android_security_Scrypt.cpp", source)
	if len(natives) != 1 || len(registrations) != 1 {
		t.Fatalf("natives=%+v registrations=%+v", natives, registrations)
	}
	registration := registrations[0]
	if registration.ClassName != "android.security.Scrypt" ||
		registration.MethodName != "nativeScrypt" ||
		registration.FunctionName != "nativeScrypt" ||
		registration.Confidence != ConfidenceExact {
		t.Fatalf("registration: %+v", registration)
	}
}

func TestScanCppStaticExport(t *testing.T) {
	source := `void Java_android_filterfw_core_NativeFrame_nativeAllocate(
    JNIEnv* env, jobject object) {}`
	natives, _ := ScanCpp("jni_native_frame.cpp", source)
	if len(natives) != 1 ||
		natives[0].ClassName != "android.filterfw.core.NativeFrame" ||
		natives[0].MethodName != "nativeAllocate" ||
		natives[0].Confidence != ConfidenceExact {
		t.Fatalf("natives: %+v", natives)
	}
}

func TestScanCppOverloadedStaticExportIsInferred(t *testing.T) {
	source := `void Java_demo_Bridge_close__I(JNIEnv*, jobject, jint) {}`
	natives, _ := ScanCpp("bridge.cpp", source)
	if len(natives) != 1 ||
		natives[0].ClassName != "demo.Bridge" ||
		natives[0].MethodName != "close" ||
		natives[0].Confidence != ConfidenceInferred {
		t.Fatalf("natives: %+v", natives)
	}
}

func TestScanCppMacroEntry(t *testing.T) {
	source := `
static const JNINativeMethod methods[] = {
  MAKE_JNI_NATIVE_METHOD("close", "()V", closeNative),
};
int registerMethods(JNIEnv* env) {
  return RegisterMethodsOrDie(env, "demo/Bridge", methods, 1);
}`
	_, registrations := ScanCpp("bridge.cpp", source)
	if len(registrations) != 1 ||
		registrations[0].FunctionName != "closeNative" {
		t.Fatalf("registrations: %+v", registrations)
	}
}

func TestScanCppBindsOnlyReferencedTable(t *testing.T) {
	source := `
static const JNINativeMethod first[] = {
  {"open", "()V", (void*)openNative},
};
static const JNINativeMethod second[] = {
  {"close", "()V", (void*)closeNative},
};
int registerMethods(JNIEnv* env) {
  return RegisterMethodsOrDie(env, "demo/Bridge", second, 1);
}`
	_, registrations := ScanCpp("bridge.cpp", source)
	if len(registrations) != 1 ||
		registrations[0].MethodName != "close" ||
		registrations[0].FunctionName != "closeNative" {
		t.Fatalf("registrations: %+v", registrations)
	}
}

func TestScanCppReportsUnknownTableAsUnresolved(t *testing.T) {
	source := `int registerMethods(JNIEnv* env) {
  return RegisterMethodsOrDie(env, "demo/Bridge", unknownMethods, 1);
}`
	_, registrations := ScanCpp("bridge.cpp", source)
	if len(registrations) != 1 ||
		registrations[0].ClassName != "demo.Bridge" ||
		registrations[0].FunctionName != "unknownMethods" ||
		registrations[0].Confidence != ConfidenceUnresolved {
		t.Fatalf("registrations: %+v", registrations)
	}
}

package jni

import (
	"testing"

	"github.com/i-rtfsc/as-aosp/core/model"
)

func TestBuildIndexLinksDynamicRegistration(t *testing.T) {
	java := map[model.RepoPath]string{
		"android/security/Scrypt.java": `package android.security;
class Scrypt {
  private static native byte[] nativeScrypt(
      byte[] password, byte[] salt, int n, int r, int p, int keyLen);
}`,
	}
	cpp := map[model.RepoPath]string{
		"android_security_Scrypt.cpp": `
static const JNINativeMethod sMethods[] = {
  {"nativeScrypt", "([B[BIIII)[B", (void*)nativeScrypt},
};
int registerMethods(JNIEnv* env) {
  return jniRegisterNativeMethods(
      env, "android/security/Scrypt", sMethods, 1);
}`,
	}
	index := BuildIndex(java, cpp)
	links := index.JavaPeers("android.security.Scrypt", "nativeScrypt")
	if len(links) != 1 || links[0].Confidence != ConfidenceExact {
		t.Fatalf("index: %+v", index)
	}
	if len(index.NativePeers("nativeScrypt")) != 1 {
		t.Fatalf("native peers: %+v", index.NativePeers("nativeScrypt"))
	}
}

func TestBuildIndexLinksStaticExport(t *testing.T) {
	java := map[model.RepoPath]string{
		"android/filterfw/core/NativeFrame.java": `package android.filterfw.core;
class NativeFrame { private native void nativeAllocate(); }`,
	}
	cpp := map[model.RepoPath]string{
		"jni_native_frame.cpp": `void
Java_android_filterfw_core_NativeFrame_nativeAllocate(
    JNIEnv* env, jobject object) {}`,
	}
	index := BuildIndex(java, cpp)
	if links := index.JavaPeers(
		"android.filterfw.core.NativeFrame",
		"nativeAllocate",
	); len(links) != 1 {
		t.Fatalf("index: %+v", index)
	}
}

func TestBuildIndexDoesNotMarkMismatchedDescriptorExact(t *testing.T) {
	java := map[model.RepoPath]string{
		"demo/Bridge.java": `package demo;
class Bridge { private native void close(int value); }`,
	}
	cpp := map[model.RepoPath]string{
		"bridge.cpp": `
static const JNINativeMethod methods[] = {
  {"close", "()V", (void*)closeNative},
};
int registerMethods(JNIEnv* env) {
  return RegisterMethodsOrDie(env, "demo/Bridge", methods, 1);
}`,
	}
	links := BuildIndex(java, cpp).JavaPeers("demo.Bridge", "close")
	if len(links) != 1 || links[0].Confidence != ConfidenceInferred {
		t.Fatalf("links: %+v", links)
	}
}

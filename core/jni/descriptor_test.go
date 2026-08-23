package jni

import "testing"

func TestNormalizeDescriptor(t *testing.T) {
	for _, value := range []string{
		"()V",
		"(JIZ)V",
		"([B[BIIII)[B",
		"(Ljava/lang/String;[I)J",
	} {
		if _, ok := NormalizeDescriptor(value); !ok {
			t.Fatalf("descriptor rejected: %s", value)
		}
	}
}

func TestNormalizeDescriptorRejectsInvalidValue(t *testing.T) {
	for _, value := range []string{"V", "(V)V", "()", "(Lfoo)V"} {
		if _, ok := NormalizeDescriptor(value); ok {
			t.Fatalf("descriptor accepted: %s", value)
		}
	}
}

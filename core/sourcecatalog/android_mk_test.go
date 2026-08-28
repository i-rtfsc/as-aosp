package sourcecatalog

import "testing"

func TestParseMakefileFindsLocalModules(t *testing.T) {
	modules := parseMakefile("legacy/Android.mk", `
include $(CLEAR_VARS)
LOCAL_MODULE := legacy-java
LOCAL_SRC_FILES := src/Legacy.java
LOCAL_JAVA_LIBRARIES := framework
include $(BUILD_JAVA_LIBRARY)

include $(CLEAR_VARS)
LOCAL_MODULE := legacy-native
LOCAL_SRC_FILES := native.cpp
include $(BUILD_SHARED_LIBRARY)
`)
	if len(modules) != 2 || modules[0].Name != "legacy-java" ||
		modules[0].Type != "java_library" ||
		modules[1].Type != "cc_library_shared" {
		t.Fatalf("modules=%+v", modules)
	}
}

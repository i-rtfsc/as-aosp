cmake_minimum_required(VERSION 3.5)
project(libandroid_runtime)

set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/frameworks/base/core/jni/*.c
        ${ANDROID_ROOT}/frameworks/base/core/jni/*.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android/opengl/*.c
        ${ANDROID_ROOT}/frameworks/base/core/jni/android/opengl/*.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/hwbinder/*.c
        ${ANDROID_ROOT}/frameworks/base/core/jni/hwbinder/*.cpp
)

include_directories(
        # frameworks/base
        "${ANDROID_ROOT}/frameworks/base/core/jni"
        "${ANDROID_ROOT}/frameworks/base/core/jni/include"
        "${ANDROID_ROOT}/frameworks/base/core/jni/include_vm"
        "${ANDROID_ROOT}/frameworks/base/media/jni"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/include"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/include"
        # frameworks/native
        "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
        "${ANDROID_ROOT}/frameworks/native/libs/battery"
        "${ANDROID_ROOT}/frameworks/native/libs/binderthreadstate/include"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private"
        "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include"
        "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include"
        "${ANDROID_ROOT}/frameworks/native/libs/math/include"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types"
        "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include"
        "${ANDROID_ROOT}/frameworks/native/libs/graphicsenv/include"
        "${ANDROID_ROOT}/frameworks/native/libs/gui/include"
        "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include"
        "${ANDROID_ROOT}/frameworks/native/opengl/include"
        "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private"
        "${ANDROID_ROOT}/frameworks/native/libs/permission/include"
        "${ANDROID_ROOT}/frameworks/native/libs/sensor/include"
        "${ANDROID_ROOT}/frameworks/native/headers/media_plugin"
        "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/media/openmax"
        "${ANDROID_ROOT}/frameworks/native/libs/nativedisplay/include"
        "${ANDROID_ROOT}/frameworks/native/libs/nativedisplay/include-private"
        "${ANDROID_ROOT}/frameworks/native/libs/cputimeinstate"
        # frameworks/av
        "${ANDROID_ROOT}/frameworks/av/media/ndk/include"
        "${ANDROID_ROOT}/frameworks/av/media/ndk/include_platform"
        "${ANDROID_ROOT}/frameworks/av/camera/include"
        "${ANDROID_ROOT}/frameworks/av/camera/include/camera"
        "${ANDROID_ROOT}/frameworks/av/media/libmedia/include"
        "${ANDROID_ROOT}/frameworks/av/media/libstagefright/include"
        "${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include"
        "${ANDROID_ROOT}/frameworks/av/media/libaudiofoundation/include"
        "${ANDROID_ROOT}/frameworks/av/media/liberror/include"
        "${ANDROID_ROOT}/frameworks/av/media/libmediahelper/include"
        "${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include"
        "${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl"
        "${ANDROID_ROOT}/frameworks/av/media/libstagefright/foundation/include"
        "${ANDROID_ROOT}/frameworks/av/media/img_utils/include"
)

file (GLOB_RECURSE HEADERS
        # frameworks/base
        "${ANDROID_ROOT}/frameworks/base/core/jni/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/core/jni/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/core/jni/include_vm/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/media/jni/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/include/**/*.h"
        # frameworks/native
        "${ANDROID_ROOT}/frameworks/native/libs/binderthreadstate/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/battery/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/math/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/graphicsenv/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/gui/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/permission/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/sensor/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/media/openmax/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/nativedisplay/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/nativedisplay/include-private/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/cputimeinstate/**/*.h"
        # frameworks/av
        "${ANDROID_ROOT}/frameworks/av/media/ndk/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/ndk/include_platform/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/camera/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/camera/include/camera/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/libmedia/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/libstagefright/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/libaudiofoundation/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/liberror/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/libmediahelper/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/libstagefright/foundation/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/av/media/img_utils/include/**/*.h"
)

add_executable(
        libandroid_runtime
        ${SOURCE_FILES}
        ${HEADERS}
)
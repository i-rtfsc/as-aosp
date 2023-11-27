cmake_minimum_required(VERSION 3.5)
project(AndroidRuntime)

# 判断主 CMakeLists.txt 是否已经设置正确的 ANDROID_ROOT 目录
if(IS_DIRECTORY ${ANDROID_ROOT})
    message(STATUS "${PROJECT_NAME}, has been set android root")
else()
    if(IS_DIRECTORY ${BUILD_NATIVE_ROOT})
        message(WARNING "${PROJECT_NAME}, set ANDROID_ROOT via BUILD_NATIVE_ROOT=${BUILD_NATIVE_ROOT}")
        set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})
    else()
        message(FATAL_ERROR "${PROJECT_NAME}, ANDROID_ROOT was NULL")
    endif()
endif()

set(AOSP_FWK_BASE true)
set(AOSP_FWK_NATIVE true)
set(AOSP_SYSTEM_COMMON true)
set(AOSP_FWK_AV false)
set(AOSP_SYSTEM false)
set(AOSP_ART false)
set(AOSP_BIONIC false)
set(AOSP_EXTERNAL false)
set(AOSP_PACKAGES false)
set(AOSP_BOOTABLE false)
set(AOSP_HARDWARE false)
set(AOSP_VENDOR false)
set(AOSP_OUT false)
set(AOSP_OTHER false)

# 设置 cflags
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -DFLYME_GRAPHICS_EXTEND")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -DFLYME_REFLECT_CHECK")
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -DFLYME_GRAPHICS_EXTEND")
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -DFLYME_REFLECT_CHECK")

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/frameworks/base/cmds/uinput/jni/*.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/*.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android/opengl/*.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/hwbinder/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/fuzz/cursorwindow_fuzzer/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/fuzz/resourcefile_fuzzer/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/tests/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/canvas/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/effects/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/flyme/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/hwui/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/fonts/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/pdf/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/text/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/pipeline/skia/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderstate/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/renderthread/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/service/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/tests/common/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/tests/common/scenes/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/tests/macrobench/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/tests/microbench/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/tests/unit/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/thread/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/hwui/utils/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/input/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/input/tests/*.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/storage/*.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/*.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/audioeffect/*.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/soundpool/*.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/soundpool/tests/*.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/tuner/*.cpp
        ${ANDROID_ROOT}/frameworks/base/media/native/midi/*.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/*.c
        ${ANDROID_ROOT}/frameworks/base/native/android/*.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/tests/activitymanager/nativeTests/src/*.cpp
        ${ANDROID_ROOT}/frameworks/base/native/android/tests/performance_hint/*.cpp
        )

include_directories(
        "${ANDROID_ROOT}/frameworks/base/cmds/uinput/jni"
        "${ANDROID_ROOT}/frameworks/base/core/jni"
        "${ANDROID_ROOT}/frameworks/base/core/jni/include"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/fuzz/cursorwindow_fuzzer"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/fuzz/resourcefile_fuzzer"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/include"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/include"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui/jni"
        "${ANDROID_ROOT}/frameworks/base/libs/input"
        "${ANDROID_ROOT}/frameworks/base/libs/input/tests"
        "${ANDROID_ROOT}/frameworks/base/libs/storage"
        "${ANDROID_ROOT}/frameworks/base/libs/storage/include"
        "${ANDROID_ROOT}/frameworks/base/media/jni"
        "${ANDROID_ROOT}/frameworks/base/media/jni/audioeffect"
        "${ANDROID_ROOT}/frameworks/base/media/jni/soundpool"
        "${ANDROID_ROOT}/frameworks/base/media/jni/soundpool/tests"
        "${ANDROID_ROOT}/frameworks/base/media/native/midi"
        "${ANDROID_ROOT}/frameworks/base/media/native/midi/include"
        "${ANDROID_ROOT}/frameworks/base/native/android"
        "${ANDROID_ROOT}/frameworks/base/native/android/aidl"
        "${ANDROID_ROOT}/frameworks/base/native/android/include_platform"
        "${ANDROID_ROOT}/frameworks/base/native/android/tests/activitymanager/nativeTests"
        "${ANDROID_ROOT}/frameworks/base/native/android/tests/performance_hint"
)

file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/frameworks/base/cmds/uinput/jni/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/core/jni/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/core/jni/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/fuzz/cursorwindow_fuzzer/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/fuzz/resourcefile_fuzzer/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui/jni/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/input/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/input/tests/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/storage/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/storage/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/media/jni/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/media/jni/audioeffect/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/media/jni/soundpool/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/media/jni/soundpool/tests/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/media/native/midi/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/media/native/midi/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/native/android/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/native/android/aidl/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/native/android/include_platform/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/native/android/tests/activitymanager/nativeTests/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/native/android/tests/performance_hint/**/*.h"
        )


# add aosp fwk_base
if (${AOSP_FWK_BASE})
    message("${PROJECT_NAME}, enable aosp fwk_base")

    # add aosp fwk_base header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/base/libs"
            "${ANDROID_ROOT}/frameworks/base/libs/protoutil/include"
    )

    # add aosp fwk_base header file
    file(GLOB_RECURSE HEADERS_FWK_BASE
            "${ANDROID_ROOT}/frameworks/base/libs/**/*.h"
            "${ANDROID_ROOT}/frameworks/base/libs/protoutil/include/**/*.h"
            )
    # append aosp fwk_base header
    list(APPEND HEADERS ${HEADERS_FWK_BASE})
endif ()


# add aosp fwk_native
if (${AOSP_FWK_NATIVE})
    message("${PROJECT_NAME}, enable aosp fwk_native")

    # add aosp fwk_native header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/native/headers/media_plugin"
            "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/media/openmax"
            "${ANDROID_ROOT}/frameworks/native/include"
            "${ANDROID_ROOT}/frameworks/native/include/media/openmax"
            "${ANDROID_ROOT}/frameworks/native/include/private"
            "${ANDROID_ROOT}/frameworks/native/libs/android_runtime_lazy/include"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
            "${ANDROID_ROOT}/frameworks/native/libs/battery"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_activitymanager"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform"
            "${ANDROID_ROOT}/frameworks/native/libs/binderthreadstate/include"
            "${ANDROID_ROOT}/frameworks/native/libs/cputimeinstate"
            "${ANDROID_ROOT}/frameworks/native/libs/em/aidl"
            "${ANDROID_ROOT}/frameworks/native/libs/em/include"
            "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include"
            "${ANDROID_ROOT}/frameworks/native/libs/graphicsenv/include"
            "${ANDROID_ROOT}/frameworks/native/libs/gui/include"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativedisplay/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativedisplay/include-private"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private"
            "${ANDROID_ROOT}/frameworks/native/libs/permission/include"
            "${ANDROID_ROOT}/frameworks/native/libs/sensor/include"
            "${ANDROID_ROOT}/frameworks/native/libs/shaders/include"
            "${ANDROID_ROOT}/frameworks/native/libs/tonemap/include"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types"
            "${ANDROID_ROOT}/frameworks/native/libs/ultrahdr/include"
            "${ANDROID_ROOT}/frameworks/native/opengl/include"
            "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL"
            "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include"
            "${ANDROID_ROOT}/frameworks/native/services"
            "${ANDROID_ROOT}/frameworks/native/services/inputflinger/include"
            "${ANDROID_ROOT}/frameworks/native/services/powermanager/include"
    )

    # add aosp fwk_native header file
    file(GLOB_RECURSE HEADERS_FWK_NATIVE
            "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/media/openmax/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/include/media/openmax/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/include/private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/android_runtime_lazy/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/battery/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_activitymanager/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binderthreadstate/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/cputimeinstate/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/em/aidl/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/em/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/graphicsenv/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/gui/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativedisplay/include-private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativedisplay/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/permission/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/sensor/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/shaders/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/tonemap/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ultrahdr/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/services/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/services/inputflinger/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/services/powermanager/include/**/*.h"
            )
    # append aosp fwk_native header
    list(APPEND HEADERS ${HEADERS_FWK_NATIVE})
endif ()


# add aosp system_common
if (${AOSP_SYSTEM_COMMON})
    message("${PROJECT_NAME}, enable aosp system_common")

    # add aosp system_common header dir
    include_directories(
            "${ANDROID_ROOT}/system/core/debuggerd/common/include"
            "${ANDROID_ROOT}/system/core/debuggerd/include"
            "${ANDROID_ROOT}/system/core/include"
            "${ANDROID_ROOT}/system/core/libappfuse/include"
            "${ANDROID_ROOT}/system/core/libcutils/include"
            "${ANDROID_ROOT}/system/core/libgrallocusage/include"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include"
            "${ANDROID_ROOT}/system/core/libsync/include"
            "${ANDROID_ROOT}/system/core/libsystem/include"
            "${ANDROID_ROOT}/system/core/libusbhost/include"
            "${ANDROID_ROOT}/system/core/libutils/include"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
            "${ANDROID_ROOT}/system/libbase/include"
            "${ANDROID_ROOT}/system/libfmq/base"
            "${ANDROID_ROOT}/system/libfmq/include"
            "${ANDROID_ROOT}/system/libhidl/base/include"
            "${ANDROID_ROOT}/system/libhidl/libhidlmemory/include"
            "${ANDROID_ROOT}/system/libhidl/transport/allocator/1.0/utils/include"
            "${ANDROID_ROOT}/system/libhidl/transport/include"
            "${ANDROID_ROOT}/system/libhidl/transport/token/1.0/utils/include"
            "${ANDROID_ROOT}/system/libhwbinder/include"
            "${ANDROID_ROOT}/system/logging/liblog/include"
            "${ANDROID_ROOT}/system/tools/aidl/metadata/include"
            "${ANDROID_ROOT}/system/tools/hidl/metadata/include"
    )

    # add aosp system_common header file
    file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
            "${ANDROID_ROOT}/system/core/debuggerd/common/include/**/*.h"
            "${ANDROID_ROOT}/system/core/debuggerd/include/**/*.h"
            "${ANDROID_ROOT}/system/core/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libappfuse/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libgrallocusage/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libsync/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libsystem/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libusbhost/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
            "${ANDROID_ROOT}/system/libbase/include/**/*.h"
            "${ANDROID_ROOT}/system/libfmq/base/**/*.h"
            "${ANDROID_ROOT}/system/libfmq/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/base/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/libhidlmemory/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/allocator/1.0/utils/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/token/1.0/utils/include/**/*.h"
            "${ANDROID_ROOT}/system/libhwbinder/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
            "${ANDROID_ROOT}/system/tools/aidl/metadata/include/**/*.h"
            "${ANDROID_ROOT}/system/tools/hidl/metadata/include/**/*.h"
            )
    # append aosp system_common header
    list(APPEND HEADERS ${HEADERS_SYSTEM_COMMON})
endif ()


# add aosp fwk_av
if (${AOSP_FWK_AV})
    message("${PROJECT_NAME}, enable aosp fwk_av")

    # add aosp fwk_av header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/av/camera/include"
            "${ANDROID_ROOT}/frameworks/av/camera/include/camera"
            "${ANDROID_ROOT}/frameworks/av/drm/libmediadrm/include"
            "${ANDROID_ROOT}/frameworks/av/drm/libmediadrm/interface"
            "${ANDROID_ROOT}/frameworks/av/include"
            "${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include"
            "${ANDROID_ROOT}/frameworks/av/media/codec2/core/include"
            "${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/include"
            "${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/internal"
            "${ANDROID_ROOT}/frameworks/av/media/img_utils/include"
            "${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiofoundation/include"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiohal/include"
            "${ANDROID_ROOT}/frameworks/av/media/liberror/include"
            "${ANDROID_ROOT}/frameworks/av/media/libheif/include"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/include"
            "${ANDROID_ROOT}/frameworks/av/media/libmediahelper/include"
            "${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include"
            "${ANDROID_ROOT}/frameworks/av/media/libnbaio/include_mono"
            "${ANDROID_ROOT}/frameworks/av/media/libstagefright/include"
            "${ANDROID_ROOT}/frameworks/av/media/module/foundation/include"
            "${ANDROID_ROOT}/frameworks/av/media/mtp"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include_platform"
    )

    # add aosp fwk_av header file
    file(GLOB_RECURSE HEADERS_FWK_AV
            "${ANDROID_ROOT}/frameworks/av/camera/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/camera/include/camera/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/drm/libmediadrm/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/drm/libmediadrm/interface/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/codec2/core/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/internal/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/img_utils/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiofoundation/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiohal/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/liberror/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libheif/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmediahelper/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libnbaio/include_mono/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libstagefright/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/module/foundation/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/mtp/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include_platform/**/*.h"
            )
    # append aosp fwk_av header
    list(APPEND HEADERS ${HEADERS_FWK_AV})
endif ()


# add aosp system
if (${AOSP_SYSTEM})
    message("${PROJECT_NAME}, enable aosp system")

    # add aosp system header dir
    include_directories(
            "${ANDROID_ROOT}/system/extras/root_feature/user"
            "${ANDROID_ROOT}/system/incremental_delivery/incfs/include"
            "${ANDROID_ROOT}/system/incremental_delivery/incfs/kernel-headers"
            "${ANDROID_ROOT}/system/incremental_delivery/incfs/util/include"
            "${ANDROID_ROOT}/system/incremental_delivery/libdataloader/include"
            "${ANDROID_ROOT}/system/libprocinfo/include"
            "${ANDROID_ROOT}/system/libvintf"
            "${ANDROID_ROOT}/system/libvintf/include"
            "${ANDROID_ROOT}/system/libziparchive/incfs_support/include"
            "${ANDROID_ROOT}/system/libziparchive/include"
            "${ANDROID_ROOT}/system/media/audio/include"
            "${ANDROID_ROOT}/system/media/audio_utils/include"
            "${ANDROID_ROOT}/system/media/camera/include"
            "${ANDROID_ROOT}/system/media/private/camera/include"
            "${ANDROID_ROOT}/system/memory/libdmabufheap/include"
            "${ANDROID_ROOT}/system/memory/libion/include"
            "${ANDROID_ROOT}/system/memory/libion/kernel-headers"
            "${ANDROID_ROOT}/system/memory/libmeminfo/include"
            "${ANDROID_ROOT}/system/memory/libmeminfo/libdmabufinfo/include"
            "${ANDROID_ROOT}/system/memory/libmemtrack/include"
            "${ANDROID_ROOT}/system/memory/libmemunreachable/include"
            "${ANDROID_ROOT}/system/netd/include"
            "${ANDROID_ROOT}/system/server_configurable_flags/libflags/include"
    )

    # add aosp system header file
    file(GLOB_RECURSE HEADERS_SYSTEM
            "${ANDROID_ROOT}/system/extras/root_feature/user/**/*.h"
            "${ANDROID_ROOT}/system/incremental_delivery/incfs/include/**/*.h"
            "${ANDROID_ROOT}/system/incremental_delivery/incfs/kernel-headers/**/*.h"
            "${ANDROID_ROOT}/system/incremental_delivery/incfs/util/include/**/*.h"
            "${ANDROID_ROOT}/system/incremental_delivery/libdataloader/include/**/*.h"
            "${ANDROID_ROOT}/system/libprocinfo/include/**/*.h"
            "${ANDROID_ROOT}/system/libvintf/**/*.h"
            "${ANDROID_ROOT}/system/libvintf/include/**/*.h"
            "${ANDROID_ROOT}/system/libziparchive/incfs_support/include/**/*.h"
            "${ANDROID_ROOT}/system/libziparchive/include/**/*.h"
            "${ANDROID_ROOT}/system/media/audio/include/**/*.h"
            "${ANDROID_ROOT}/system/media/audio_utils/include/**/*.h"
            "${ANDROID_ROOT}/system/media/camera/include/**/*.h"
            "${ANDROID_ROOT}/system/media/private/camera/include/**/*.h"
            "${ANDROID_ROOT}/system/memory/libdmabufheap/include/**/*.h"
            "${ANDROID_ROOT}/system/memory/libion/include/**/*.h"
            "${ANDROID_ROOT}/system/memory/libion/kernel-headers/**/*.h"
            "${ANDROID_ROOT}/system/memory/libmeminfo/include/**/*.h"
            "${ANDROID_ROOT}/system/memory/libmeminfo/libdmabufinfo/include/**/*.h"
            "${ANDROID_ROOT}/system/memory/libmemtrack/include/**/*.h"
            "${ANDROID_ROOT}/system/memory/libmemunreachable/include/**/*.h"
            "${ANDROID_ROOT}/system/netd/include/**/*.h"
            "${ANDROID_ROOT}/system/server_configurable_flags/libflags/include/**/*.h"
            )
    # append aosp system header
    list(APPEND HEADERS ${HEADERS_SYSTEM})
endif ()


# add aosp art
if (${AOSP_ART})
    message("${PROJECT_NAME}, enable aosp art")

    # add aosp art header dir
    include_directories(
            "${ANDROID_ROOT}/art/libnativebridge/include"
            "${ANDROID_ROOT}/art/libnativeloader/include"
    )

    # add aosp art header file
    file(GLOB_RECURSE HEADERS_ART
            "${ANDROID_ROOT}/art/libnativebridge/include/**/*.h"
            "${ANDROID_ROOT}/art/libnativeloader/include/**/*.h"
            )
    # append aosp art header
    list(APPEND HEADERS ${HEADERS_ART})
endif ()


# add aosp bionic
if (${AOSP_BIONIC})
    message("${PROJECT_NAME}, enable aosp bionic")

    # add aosp bionic header dir
    include_directories(
            "${ANDROID_ROOT}/bionic/libc/async_safe/include"
            "${ANDROID_ROOT}/bionic/libc/dns/include"
            "${ANDROID_ROOT}/bionic/libc/include"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/scsi"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/uapi"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-arm64"
            "${ANDROID_ROOT}/bionic/libc/platform"
            "${ANDROID_ROOT}/bionic/libc/seccomp/include"
            "${ANDROID_ROOT}/bionic/libc/system_properties/include"
    )

    # add aosp bionic header file
    file(GLOB_RECURSE HEADERS_BIONIC
            "${ANDROID_ROOT}/bionic/libc/async_safe/include/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/dns/include/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/include/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/scsi/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/uapi/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-arm64/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/platform/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/seccomp/include/**/*.h"
            "${ANDROID_ROOT}/bionic/libc/system_properties/include/**/*.h"
            )
    # append aosp bionic header
    list(APPEND HEADERS ${HEADERS_BIONIC})
endif ()


# add aosp external
if (${AOSP_EXTERNAL})
    message("${PROJECT_NAME}, enable aosp external")

    # add aosp external header dir
    include_directories(
            "${ANDROID_ROOT}/external/boringssl/src/include"
            "${ANDROID_ROOT}/external/dng_sdk/source"
            "${ANDROID_ROOT}/external/expat/lib"
            "${ANDROID_ROOT}/external/flatbuffers/include"
            "${ANDROID_ROOT}/external/fmtlib/include"
            "${ANDROID_ROOT}/external/freetype/include"
            "${ANDROID_ROOT}/external/giflib"
            "${ANDROID_ROOT}/external/google-benchmark/include"
            "${ANDROID_ROOT}/external/googletest/googlemock/include"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/harfbuzz_ng/src"
            "${ANDROID_ROOT}/external/icu/libicu/ndk_headers"
            "${ANDROID_ROOT}/external/image_io/includes"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/libjpeg-turbo"
            "${ANDROID_ROOT}/external/libpng"
            "${ANDROID_ROOT}/external/libtextclassifier/native"
            "${ANDROID_ROOT}/external/libxml2/include"
            "${ANDROID_ROOT}/external/pdfium/public"
            "${ANDROID_ROOT}/external/perfetto/include"
            "${ANDROID_ROOT}/external/perfetto/include/perfetto/base/build_configs/android_tree"
            "${ANDROID_ROOT}/external/piex"
            "${ANDROID_ROOT}/external/protobuf/src"
            "${ANDROID_ROOT}/external/scrypt/lib/crypto"
            "${ANDROID_ROOT}/external/selinux/libselinux/include"
            "${ANDROID_ROOT}/external/sfntly/cpp/src"
            "${ANDROID_ROOT}/external/skia"
            "${ANDROID_ROOT}/external/skia/android"
            "${ANDROID_ROOT}/external/skia/client_utils/android"
            "${ANDROID_ROOT}/external/skia/include/android"
            "${ANDROID_ROOT}/external/skia/include/codec"
            "${ANDROID_ROOT}/external/skia/include/config"
            "${ANDROID_ROOT}/external/skia/include/core"
            "${ANDROID_ROOT}/external/skia/include/docs"
            "${ANDROID_ROOT}/external/skia/include/effects"
            "${ANDROID_ROOT}/external/skia/include/encode"
            "${ANDROID_ROOT}/external/skia/include/gpu"
            "${ANDROID_ROOT}/external/skia/include/pathops"
            "${ANDROID_ROOT}/external/skia/include/ports"
            "${ANDROID_ROOT}/external/skia/include/private"
            "${ANDROID_ROOT}/external/skia/include/svg"
            "${ANDROID_ROOT}/external/skia/include/utils"
            "${ANDROID_ROOT}/external/skia/include/utils/mac"
            "${ANDROID_ROOT}/external/skia/modules/skottie/include"
            "${ANDROID_ROOT}/external/skia/modules/skparagraph/include"
            "${ANDROID_ROOT}/external/skia/modules/skshaper/include"
            "${ANDROID_ROOT}/external/skia/modules/svg/include"
            "${ANDROID_ROOT}/external/skia/src/codec"
            "${ANDROID_ROOT}/external/skia/src/core"
            "${ANDROID_ROOT}/external/skia/src/gpu"
            "${ANDROID_ROOT}/external/skia/src/image"
            "${ANDROID_ROOT}/external/skia/src/shaders"
            "${ANDROID_ROOT}/external/skia/src/utils"
            "${ANDROID_ROOT}/external/skia/third_party/libtiff"
            "${ANDROID_ROOT}/external/sonivox/arm-wt-22k/include"
            "${ANDROID_ROOT}/external/sqlite/android"
            "${ANDROID_ROOT}/external/sqlite/dist"
            "${ANDROID_ROOT}/external/tensorflow"
            "${ANDROID_ROOT}/external/vulkan-headers/include"
            "${ANDROID_ROOT}/external/webp/src"
            "${ANDROID_ROOT}/external/wuffs-mirror-release-c/release/c"
            "${ANDROID_ROOT}/external/zlib"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/boringssl/src/include/**/*.h"
            "${ANDROID_ROOT}/external/dng_sdk/source/**/*.h"
            "${ANDROID_ROOT}/external/expat/lib/**/*.h"
            "${ANDROID_ROOT}/external/flatbuffers/include/**/*.h"
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/freetype/include/**/*.h"
            "${ANDROID_ROOT}/external/giflib/**/*.h"
            "${ANDROID_ROOT}/external/google-benchmark/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/harfbuzz_ng/src/**/*.h"
            "${ANDROID_ROOT}/external/icu/libicu/ndk_headers/**/*.h"
            "${ANDROID_ROOT}/external/image_io/includes/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/libjpeg-turbo/**/*.h"
            "${ANDROID_ROOT}/external/libpng/**/*.h"
            "${ANDROID_ROOT}/external/libtextclassifier/native/**/*.h"
            "${ANDROID_ROOT}/external/libxml2/include/**/*.h"
            "${ANDROID_ROOT}/external/pdfium/public/**/*.h"
            "${ANDROID_ROOT}/external/perfetto/include/**/*.h"
            "${ANDROID_ROOT}/external/perfetto/include/perfetto/base/build_configs/android_tree/**/*.h"
            "${ANDROID_ROOT}/external/piex/**/*.h"
            "${ANDROID_ROOT}/external/protobuf/src/**/*.h"
            "${ANDROID_ROOT}/external/scrypt/lib/crypto/**/*.h"
            "${ANDROID_ROOT}/external/selinux/libselinux/include/**/*.h"
            "${ANDROID_ROOT}/external/sfntly/cpp/src/**/*.h"
            "${ANDROID_ROOT}/external/skia/**/*.h"
            "${ANDROID_ROOT}/external/skia/android/**/*.h"
            "${ANDROID_ROOT}/external/skia/client_utils/android/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/android/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/codec/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/config/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/core/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/docs/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/effects/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/encode/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/gpu/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/pathops/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/ports/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/private/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/svg/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/utils/**/*.h"
            "${ANDROID_ROOT}/external/skia/include/utils/mac/**/*.h"
            "${ANDROID_ROOT}/external/skia/modules/skottie/include/**/*.h"
            "${ANDROID_ROOT}/external/skia/modules/skparagraph/include/**/*.h"
            "${ANDROID_ROOT}/external/skia/modules/skshaper/include/**/*.h"
            "${ANDROID_ROOT}/external/skia/modules/svg/include/**/*.h"
            "${ANDROID_ROOT}/external/skia/src/codec/**/*.h"
            "${ANDROID_ROOT}/external/skia/src/core/**/*.h"
            "${ANDROID_ROOT}/external/skia/src/gpu/**/*.h"
            "${ANDROID_ROOT}/external/skia/src/image/**/*.h"
            "${ANDROID_ROOT}/external/skia/src/shaders/**/*.h"
            "${ANDROID_ROOT}/external/skia/src/utils/**/*.h"
            "${ANDROID_ROOT}/external/skia/third_party/libtiff/**/*.h"
            "${ANDROID_ROOT}/external/sonivox/arm-wt-22k/include/**/*.h"
            "${ANDROID_ROOT}/external/sqlite/android/**/*.h"
            "${ANDROID_ROOT}/external/sqlite/dist/**/*.h"
            "${ANDROID_ROOT}/external/tensorflow/**/*.h"
            "${ANDROID_ROOT}/external/vulkan-headers/include/**/*.h"
            "${ANDROID_ROOT}/external/webp/src/**/*.h"
            "${ANDROID_ROOT}/external/wuffs-mirror-release-c/release/c/**/*.h"
            "${ANDROID_ROOT}/external/zlib/**/*.h"
            )
    # append aosp external header
    list(APPEND HEADERS ${HEADERS_EXTERNAL})
endif ()


# add aosp packages
if (${AOSP_PACKAGES})
    message("${PROJECT_NAME}, enable aosp packages")

    # add aosp packages header dir
    include_directories(
            "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types"
            "${ANDROID_ROOT}/packages/modules/DnsResolver/include/dnsproxyd_protocol"
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatspull/include"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatssocket/include"
    )

    # add aosp packages header file
    file(GLOB_RECURSE HEADERS_PACKAGES
            "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types/**/*.h"
            "${ANDROID_ROOT}/packages/modules/DnsResolver/include/dnsproxyd_protocol/**/*.h"
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include/**/*.h"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatspull/include/**/*.h"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatssocket/include/**/*.h"
            )
    # append aosp packages header
    list(APPEND HEADERS ${HEADERS_PACKAGES})
endif ()


# add aosp hardware
if (${AOSP_HARDWARE})
    message("${PROJECT_NAME}, enable aosp hardware")

    # add aosp hardware header dir
    include_directories(
            "${ANDROID_ROOT}/hardware/interfaces/common/support/include"
            "${ANDROID_ROOT}/hardware/interfaces/configstore/utils/include"
            "${ANDROID_ROOT}/hardware/libhardware/include"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include"
            "${ANDROID_ROOT}/hardware/ril/include"
    )

    # add aosp hardware header file
    file(GLOB_RECURSE HEADERS_HARDWARE
            "${ANDROID_ROOT}/hardware/interfaces/common/support/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/configstore/utils/include/**/*.h"
            "${ANDROID_ROOT}/hardware/libhardware/include/**/*.h"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include/**/*.h"
            "${ANDROID_ROOT}/hardware/ril/include/**/*.h"
            )
    # append aosp hardware header
    list(APPEND HEADERS ${HEADERS_HARDWARE})
endif ()


# add aosp vendor
if (${AOSP_VENDOR})
    message("${PROJECT_NAME}, enable aosp vendor")

    # add aosp vendor header dir
    include_directories(
            "${ANDROID_ROOT}/device/meizu/common"
    )

    # add aosp vendor header file
    file(GLOB_RECURSE HEADERS_VENDOR
            "${ANDROID_ROOT}/device/meizu/common/**/*.h"
            )
    # append aosp vendor header
    list(APPEND HEADERS ${HEADERS_VENDOR})
endif ()


# add aosp out
if (${AOSP_OUT})
    message("${PROJECT_NAME}, enable aosp out")

    # add aosp out src
    file(GLOB SOURCE_FILES_OUT
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/cmds/uinput/jni/libuinputcommand_jni/android_arm64_armv8-a-branchprot_shared/gen/aidl/src/com/android/commands/uinput/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/core/jni/libandroid_runtime/android_arm64_armv8-a-branchprot_shared/gen/aidl/android/hardware/display/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/*.cpp/gen/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static/gen/proto/frameworks/base/libs/hwui/protos/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static/gen/sysprop/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui_static/android_arm64_armv8-a-branchprot_static/gen/proto/frameworks/base/libs/hwui/protos/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui_static/android_arm64_armv8-a-branchprot_static/gen/sysprop/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/media/jni/libmedia_jni_utils/android_arm64_armv8-a-branchprot_shared/gen/aidl/android/gui/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/media/native/midi/libamidi/android_arm64_armv8-a-branchprot_shared/gen/aidl/android/media/midi/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/native/android/lib-platform-compat-native-api/android_arm64_armv8-a-branchprot_shared/gen/aidl/com/android/internal/compat/*.cpp
            )
    # append aosp out src
    list(APPEND SOURCE_FILES ${SOURCE_FILES_OUT})

    # add aosp out header dir
    include_directories(
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_ipc_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_ipc_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_wire_protocol_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_wire_protocol_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_src_base_version_gen_h/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/av-types-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/camera/libcamera_client/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/libaudioclient/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia_omx/android_arm64_armv8-a-branchprot_shared_cfi/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libshmem/shared-file-region-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/cmds/uinput/jni/libuinputcommand_jni/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/core/jni/libandroid_runtime/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/proto/frameworks/base/libs/hwui"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static/gen/proto/frameworks/base/libs/hwui"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/proto/frameworks/base/libs/hwui"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui_static/android_arm64_armv8-a-branchprot_static/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui_static/android_arm64_armv8-a-branchprot_static/gen/proto/frameworks/base/libs/hwui"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui_static/android_arm64_armv8-a-branchprot_static/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/statslog_hwui.cpp/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/statslog_hwui.h/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/media/jni/libmedia_jni_utils/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/media/native/midi/libamidi/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/native/android/lib-platform-compat-native-api/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libactivitymanager_aidl/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/em/libem/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/services/powermanager/libpowermanager/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/camera/common/1.0/android.hardware.camera.common@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/camera/device/3.2/android.hardware.camera.device@3.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/cas/1.0/android.hardware.cas@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/cas/native/1.0/android.hardware.cas.native@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/fmq/aidl/android.hardware.common.fmq-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/configstore/1.0/android.hardware.configstore@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/configstore/1.1/android.hardware.configstore@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.0/android.hardware.drm@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.1/android.hardware.drm@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.2/android.hardware.drm@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.3/android.hardware.drm@1.3_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.4/android.hardware.drm@1.4_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/composer/aidl/android.hardware.graphics.composer3-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/bufferpool/2.0/android.hardware.media.bufferpool@2.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/omx/1.0/android.hardware.media.omx@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/1.0/android.hardware.power@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/1.1/android.hardware.power@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/1.2/android.hardware.power@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/1.3/android.hardware.power@1.3_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/aidl/android.hardware.power-V4-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/aidl/android.hardware.power-V4-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/allocator/1.0/android.hidl.allocator@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/memory/1.0/android.hidl.memory@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/memory/token/1.0/android.hidl.memory.token@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_shared/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_static/gen/sysprop/include"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_ipc_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_ipc_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_wire_protocol_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_ipc_wire_protocol_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_src_base_version_gen_h/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/av-types-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/camera/libcamera_client/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/libaudioclient/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia/android_arm64_armv8-a-branchprot_static_cfi/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libmedia/libmedia_omx/android_arm64_armv8-a-branchprot_shared_cfi/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/media/libshmem/shared-file-region-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/services/mediaresourcemanager/resourcemanager_aidl_interface-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/av/services/tuner/tv_tuner_aidl_interface-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/cmds/uinput/jni/libuinputcommand_jni/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/core/jni/libandroid_runtime/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/proto/frameworks/base/libs/hwui/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static/gen/proto/frameworks/base/libs/hwui/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/proto/frameworks/base/libs/hwui/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui_static/android_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui_static/android_arm64_armv8-a-branchprot_static/gen/proto/frameworks/base/libs/hwui/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui_static/android_arm64_armv8-a-branchprot_static/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/statslog_hwui.cpp/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/statslog_hwui.h/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/media/jni/libmedia_jni_utils/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/media/native/midi/libamidi/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/native/android/lib-platform-compat-native-api/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libactivitymanager_aidl/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/em/libem/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/services/powermanager/libpowermanager/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/camera/common/1.0/android.hardware.camera.common@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/camera/device/3.2/android.hardware.camera.device@3.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/cas/1.0/android.hardware.cas@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/cas/native/1.0/android.hardware.cas.native@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/fmq/aidl/android.hardware.common.fmq-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/configstore/1.0/android.hardware.configstore@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/configstore/1.1/android.hardware.configstore@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.0/android.hardware.drm@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.1/android.hardware.drm@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.2/android.hardware.drm@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.3/android.hardware.drm@1.3_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/1.4/android.hardware.drm@1.4_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/drm/aidl/android.hardware.drm-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/composer/aidl/android.hardware.graphics.composer3-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/bufferpool/2.0/android.hardware.media.bufferpool@2.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/omx/1.0/android.hardware.media.omx@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/1.0/android.hardware.power@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/1.1/android.hardware.power@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/1.2/android.hardware.power@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/1.3/android.hardware.power@1.3_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/aidl/android.hardware.power-V4-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/power/aidl/android.hardware.power-V4-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/allocator/1.0/android.hidl.allocator@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/memory/1.0/android.hidl.memory@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/memory/token/1.0/android.hidl.memory.token@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_shared/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_static/gen/sysprop/include/**/*.h"
            )
    # append aosp out header
    list(APPEND HEADERS ${HEADERS_OUT})
endif ()


add_executable(
        AndroidRuntime
        ${SOURCE_FILES}
        ${HEADERS}
)

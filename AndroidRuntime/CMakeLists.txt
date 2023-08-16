cmake_minimum_required(VERSION 3.5)
project(android_runtime)

# set(CMAKE_CXX_STANDARD 20)

#set(ANDROID_ROOT "/home/solo/code/aosp")
set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

set(AOSP_SYSTEM_COMMON true)
set(AOSP_OUT false)
set(AOSP_AV false)
set(AOSP_ART false)
set(AOSP_BIONIC false)
set(AOSP_SYSTEM false)
set(AOSP_EXTERNAL false)
set(AOSP_PACKAGES false)
set(AOSP_HARDWARE false)

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/frameworks/base/core/jni/*.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/android/opengl/*.cpp
        ${ANDROID_ROOT}/frameworks/base/core/jni/hwbinder/*.cpp
        )

include_directories(
        "${ANDROID_ROOT}/frameworks/base/core/jni"
        "${ANDROID_ROOT}/frameworks/base/core/jni/include"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/include"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/include"
        "${ANDROID_ROOT}/frameworks/base/media/jni"
        "${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/include"
        "${ANDROID_ROOT}/frameworks/minikin/include"
        "${ANDROID_ROOT}/frameworks/native/headers/media_plugin"
        "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/media/openmax"
        "${ANDROID_ROOT}/frameworks/native/include"
        "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
        "${ANDROID_ROOT}/frameworks/native/libs/battery"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform"
        "${ANDROID_ROOT}/frameworks/native/libs/binderthreadstate/include"
        "${ANDROID_ROOT}/frameworks/native/libs/cputimeinstate"
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
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types"
        "${ANDROID_ROOT}/frameworks/native/libs/ultrahdr/include"
        "${ANDROID_ROOT}/frameworks/native/opengl/include"
        "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include"
)

file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/frameworks/base/core/jni/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/core/jni/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/androidfw/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/libs/hwui/apex/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/base/media/jni/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/minikin/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/headers/media_plugin/media/openmax/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/battery/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/binderthreadstate/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/cputimeinstate/**/*.h"
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
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/libs/ultrahdr/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include/**/*.h"
        )

# add aosp system common
if (${AOSP_SYSTEM_COMMON})
    message("enable aosp system common")

    # add aosp system common header dir
    include_directories(
            "${ANDROID_ROOT}/libnativehelper/header_only_include"
            "${ANDROID_ROOT}/libnativehelper/include"
            "${ANDROID_ROOT}/libnativehelper/include_jni"
            "${ANDROID_ROOT}/libnativehelper/include_platform"
            "${ANDROID_ROOT}/libnativehelper/include_platform_header_only"
            "${ANDROID_ROOT}/system/core/debuggerd/common/include"
            "${ANDROID_ROOT}/system/core/debuggerd/include"
            "${ANDROID_ROOT}/system/core/include"
            "${ANDROID_ROOT}/system/core/libappfuse/include"
            "${ANDROID_ROOT}/system/core/libcutils/include"
            "${ANDROID_ROOT}/system/core/libgrallocusage/include"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include"
            "${ANDROID_ROOT}/system/core/libsystem/include"
            "${ANDROID_ROOT}/system/core/libusbhost/include"
            "${ANDROID_ROOT}/system/core/libutils/include"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
            "${ANDROID_ROOT}/system/libbase/include"
            "${ANDROID_ROOT}/system/libfmq/base"
            "${ANDROID_ROOT}/system/libhidl/base/include"
            "${ANDROID_ROOT}/system/libhidl/transport/include"
            "${ANDROID_ROOT}/system/libhidl/transport/token/1.0/utils/include"
            "${ANDROID_ROOT}/system/libhwbinder/include"
            "${ANDROID_ROOT}/system/logging/liblog/include"
            "${ANDROID_ROOT}/system/tools/aidl/metadata/include"
            "${ANDROID_ROOT}/system/tools/hidl/metadata/include"
    )

    # add aosp system common header file
    file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
            "${ANDROID_ROOT}/system/core/debuggerd/common/include/**/*.h"
            "${ANDROID_ROOT}/system/core/debuggerd/include/**/*.h"
            "${ANDROID_ROOT}/system/core/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libappfuse/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libgrallocusage/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libsystem/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libusbhost/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
            "${ANDROID_ROOT}/system/libbase/include/**/*.h"
            "${ANDROID_ROOT}/system/libfmq/base/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/base/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/token/1.0/utils/include/**/*.h"
            "${ANDROID_ROOT}/system/libhwbinder/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
            "${ANDROID_ROOT}/system/tools/aidl/metadata/include/**/*.h"
            "${ANDROID_ROOT}/system/tools/hidl/metadata/include/**/*.h"
            )
    # append aosp system common header
    list(APPEND HEADERS ${HEADERS_SYSTEM_COMMON})
endif ()

# add aosp system
if (${AOSP_SYSTEM})
    message("enable aosp system")

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
            "${ANDROID_ROOT}/system/media/camera/include"
            "${ANDROID_ROOT}/system/media/private/camera/include"
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
            "${ANDROID_ROOT}/system/media/camera/include/**/*.h"
            "${ANDROID_ROOT}/system/media/private/camera/include/**/*.h"
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

# add aosp av
if (${AOSP_AV})
    message("enable aosp av")

    # add aosp av header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/av/camera/include"
            "${ANDROID_ROOT}/frameworks/av/camera/include/camera"
            "${ANDROID_ROOT}/frameworks/av/include"
            "${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include"
            "${ANDROID_ROOT}/frameworks/av/media/img_utils/include"
            "${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiofoundation/include"
            "${ANDROID_ROOT}/frameworks/av/media/liberror/include"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/include"
            "${ANDROID_ROOT}/frameworks/av/media/libmediahelper/include"
            "${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include"
            "${ANDROID_ROOT}/frameworks/av/media/libstagefright/include"
            "${ANDROID_ROOT}/frameworks/av/media/module/foundation/include"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include_platform"
    )

    # add aosp av header file
    file(GLOB_RECURSE HEADERS_AV
            "${ANDROID_ROOT}/frameworks/av/camera/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/camera/include/camera/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/audioaidlconversion/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/img_utils/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libaudiofoundation/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/liberror/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmedia/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmediahelper/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/libstagefright/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/module/foundation/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/av/media/ndk/include_platform/**/*.h"
            )
    # append aosp av header
    list(APPEND HEADERS ${HEADERS_AV})
endif ()

# add aosp out
if (${AOSP_OUT})
    message("enable aosp out")

    # add aosp out header dir
    include_directories(
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
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/core/jni/libandroid_runtime/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/proto/frameworks/base/libs/hwui"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/proto/frameworks/base/libs/hwui"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/camera/common/1.0/android.hardware.camera.common@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/camera/device/3.2/android.hardware.camera.device@3.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/omx/1.0/android.hardware.media.omx@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_shared/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_static/gen/sysprop/include"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
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
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/core/jni/libandroid_runtime/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/proto/frameworks/base/libs/hwui/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_shared/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/proto/frameworks/base/libs/hwui/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/libs/hwui/libhwui/android_arm64_armv8-a-branchprot_static_afdo-libhwui/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/camera/common/1.0/android.hardware.camera.common@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/camera/device/3.2/android.hardware.camera.device@3.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/omx/1.0/android.hardware.media.omx@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_shared/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_static/gen/sysprop/include/**/*.h"
            )
    # append aosp out header
    list(APPEND HEADERS ${HEADERS_OUT})
endif ()

# add aosp art
if (${AOSP_ART})
    message("enable aosp art")

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

# add aosp external
if (${AOSP_BIONIC})
    message("enable aosp bionic")

    # add aosp bionic header dir
    include_directories(
            "${ANDROID_ROOT}/bionic/libc/async_safe/include"
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
    message("enable aosp external")

    # add aosp external header dir
    include_directories(
            "${ANDROID_ROOT}/external/boringssl/src/include"
            "${ANDROID_ROOT}/external/flatbuffers/include"
            "${ANDROID_ROOT}/external/fmtlib/include"
            "${ANDROID_ROOT}/external/giflib"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/harfbuzz_ng/src"
            "${ANDROID_ROOT}/external/icu/libicu/ndk_headers"
            "${ANDROID_ROOT}/external/image_io/includes"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/libjpeg-turbo"
            "${ANDROID_ROOT}/external/libtextclassifier/native"
            "${ANDROID_ROOT}/external/pdfium/public"
            "${ANDROID_ROOT}/external/protobuf/src"
            "${ANDROID_ROOT}/external/scrypt/lib/crypto"
            "${ANDROID_ROOT}/external/selinux/libselinux/include"
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
            "${ANDROID_ROOT}/external/sqlite/android"
            "${ANDROID_ROOT}/external/sqlite/dist"
            "${ANDROID_ROOT}/external/tensorflow"
            "${ANDROID_ROOT}/external/vulkan-headers/include"
            "${ANDROID_ROOT}/external/zlib"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/boringssl/src/include/**/*.h"
            "${ANDROID_ROOT}/external/flatbuffers/include/**/*.h"
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/giflib/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/harfbuzz_ng/src/**/*.h"
            "${ANDROID_ROOT}/external/icu/libicu/ndk_headers/**/*.h"
            "${ANDROID_ROOT}/external/image_io/includes/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/libjpeg-turbo/**/*.h"
            "${ANDROID_ROOT}/external/libtextclassifier/native/**/*.h"
            "${ANDROID_ROOT}/external/pdfium/public/**/*.h"
            "${ANDROID_ROOT}/external/protobuf/src/**/*.h"
            "${ANDROID_ROOT}/external/scrypt/lib/crypto/**/*.h"
            "${ANDROID_ROOT}/external/selinux/libselinux/include/**/*.h"
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
            "${ANDROID_ROOT}/external/sqlite/android/**/*.h"
            "${ANDROID_ROOT}/external/sqlite/dist/**/*.h"
            "${ANDROID_ROOT}/external/tensorflow/**/*.h"
            "${ANDROID_ROOT}/external/vulkan-headers/include/**/*.h"
            "${ANDROID_ROOT}/external/zlib/**/*.h"
            )
    # append aosp external header
    list(APPEND HEADERS ${HEADERS_EXTERNAL})
endif ()

# add aosp packages
if (${AOSP_PACKAGES})
    message("enable aosp packages")

    # add aosp packages header dir
    include_directories(
            "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types"
            "${ANDROID_ROOT}/packages/modules/DnsResolver/include/dnsproxyd_protocol"
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatssocket/include"
    )

    # add aosp packages header file
    file(GLOB_RECURSE HEADERS_PACKAGES
            "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types/**/*.h"
            "${ANDROID_ROOT}/packages/modules/DnsResolver/include/dnsproxyd_protocol/**/*.h"
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include/**/*.h"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatssocket/include/**/*.h"
            )
    # append aosp packages header
    list(APPEND HEADERS ${HEADERS_PACKAGES})
endif ()

# add aosp hardware
if (${AOSP_HARDWARE})
    message("enable aosp hardware")

    # add aosp hardware header dir
    include_directories(
            "${ANDROID_ROOT}/hardware/libhardware/include"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include"
            "${ANDROID_ROOT}/hardware/ril/include"
    )

    # add aosp hardware header file
    file(GLOB_RECURSE HEADERS_HARDWARE
            "${ANDROID_ROOT}/hardware/libhardware/include/**/*.h"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include/**/*.h"
            "${ANDROID_ROOT}/hardware/ril/include/**/*.h"
            "${ANDROID_ROOT}/libnativehelper/header_only_include/**/*.h"
            "${ANDROID_ROOT}/libnativehelper/include/**/*.h"
            "${ANDROID_ROOT}/libnativehelper/include_jni/**/*.h"
            "${ANDROID_ROOT}/libnativehelper/include_platform/**/*.h"
            "${ANDROID_ROOT}/libnativehelper/include_platform_header_only/**/*.h"
            )
    # append aosp hardware header
    list(APPEND HEADERS ${HEADERS_HARDWARE})
endif ()

add_executable(
        android_runtime
        ${SOURCE_FILES}
        ${HEADERS}
)

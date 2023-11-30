cmake_minimum_required(VERSION 3.5)
project(InputFlinger)

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
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -DFLYME_INPUT_EXTEND")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -DFLYME_PRODUCT_2481")
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -DFLYME_INPUT_EXTEND")
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -DFLYME_PRODUCT_2481")

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/benchmarks/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/dispatcher/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/host/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reporter/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/tests/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/tests/fuzzers/*.cpp
        )

include_directories(
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/benchmarks"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/dispatcher"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/dispatcher/include"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/host"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/include"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/include"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reporter"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/tests"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/tests/fuzzers"
)

file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/benchmarks/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/dispatcher/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/dispatcher/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/host/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/reporter/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/tests/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/inputflinger/tests/fuzzers/**/*.h"
        )


# add aosp fwk_native
if (${AOSP_FWK_NATIVE})
    message("${PROJECT_NAME}, enable aosp fwk_native")

    # add aosp fwk_native header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/native/include"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform"
            "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include"
            "${ANDROID_ROOT}/frameworks/native/libs/gui/include"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private"
            "${ANDROID_ROOT}/frameworks/native/libs/runtimelog"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types"
            "${ANDROID_ROOT}/frameworks/native/opengl/include"
            "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include"
            "${ANDROID_ROOT}/frameworks/native/services/batteryservice/include"
    )

    # add aosp fwk_native header file
    file(GLOB_RECURSE HEADERS_FWK_NATIVE
            "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/gui/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/runtimelog/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/services/batteryservice/include/**/*.h"
            )
    # append aosp fwk_native header
    list(APPEND HEADERS ${HEADERS_FWK_NATIVE})
endif ()


# add aosp system_common
if (${AOSP_SYSTEM_COMMON})
    message("${PROJECT_NAME}, enable aosp system_common")

    # add aosp system_common header dir
    include_directories(
            "${ANDROID_ROOT}/system/core/include"
            "${ANDROID_ROOT}/system/core/libcutils/include"
            "${ANDROID_ROOT}/system/core/libpackagelistparser/include"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include"
            "${ANDROID_ROOT}/system/core/libsystem/include"
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

    # add aosp system_common header file
    file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
            "${ANDROID_ROOT}/system/core/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libpackagelistparser/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libsystem/include/**/*.h"
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
    # append aosp system_common header
    list(APPEND HEADERS ${HEADERS_SYSTEM_COMMON})
endif ()


# add aosp fwk_av
if (${AOSP_FWK_AV})
    message("${PROJECT_NAME}, enable aosp fwk_av")

    # add aosp fwk_av header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/av/include"
    )

    # add aosp fwk_av header file
    file(GLOB_RECURSE HEADERS_FWK_AV
            "${ANDROID_ROOT}/frameworks/av/include/**/*.h"
            )
    # append aosp fwk_av header
    list(APPEND HEADERS ${HEADERS_FWK_AV})
endif ()


# add aosp system
if (${AOSP_SYSTEM})
    message("${PROJECT_NAME}, enable aosp system")

    # add aosp system header dir
    include_directories(
            "${ANDROID_ROOT}/system/libvintf"
            "${ANDROID_ROOT}/system/libvintf/include"
            "${ANDROID_ROOT}/system/media/audio/include"
            "${ANDROID_ROOT}/system/server_configurable_flags/libflags/include"
    )

    # add aosp system header file
    file(GLOB_RECURSE HEADERS_SYSTEM
            "${ANDROID_ROOT}/system/libvintf/**/*.h"
            "${ANDROID_ROOT}/system/libvintf/include/**/*.h"
            "${ANDROID_ROOT}/system/media/audio/include/**/*.h"
            "${ANDROID_ROOT}/system/server_configurable_flags/libflags/include/**/*.h"
            )
    # append aosp system header
    list(APPEND HEADERS ${HEADERS_SYSTEM})
endif ()


# add aosp bionic
if (${AOSP_BIONIC})
    message("${PROJECT_NAME}, enable aosp bionic")

    # add aosp bionic header dir
    include_directories(
            "${ANDROID_ROOT}/bionic/libc/async_safe/include"
            "${ANDROID_ROOT}/bionic/libc/include"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/scsi"
            "${ANDROID_ROOT}/bionic/libc/kernel/android/uapi"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi"
            "${ANDROID_ROOT}/bionic/libc/kernel/uapi/asm-arm64"
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
            "${ANDROID_ROOT}/external/fmtlib/include"
            "${ANDROID_ROOT}/external/google-benchmark/include"
            "${ANDROID_ROOT}/external/googletest/googlemock/include"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/jsoncpp/include"
            "${ANDROID_ROOT}/external/libcap/libcap/include"
            "${ANDROID_ROOT}/external/libchrome-gestures"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/libpalmrejection"
            "${ANDROID_ROOT}/external/protobuf/src"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/boringssl/src/include/**/*.h"
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/google-benchmark/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/jsoncpp/include/**/*.h"
            "${ANDROID_ROOT}/external/libcap/libcap/include/**/*.h"
            "${ANDROID_ROOT}/external/libchrome-gestures/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/libpalmrejection/**/*.h"
            "${ANDROID_ROOT}/external/protobuf/src/**/*.h"
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
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libkll/include"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatspull/include"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatssocket/include"
    )

    # add aosp packages header file
    file(GLOB_RECURSE HEADERS_PACKAGES
            "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types/**/*.h"
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include/**/*.h"
            "${ANDROID_ROOT}/packages/modules/StatsD/lib/libkll/include/**/*.h"
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
            "${ANDROID_ROOT}/hardware/libhardware/include"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include"
            "${ANDROID_ROOT}/hardware/ril/include"
    )

    # add aosp hardware header file
    file(GLOB_RECURSE HEADERS_HARDWARE
            "${ANDROID_ROOT}/hardware/libhardware/include/**/*.h"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include/**/*.h"
            "${ANDROID_ROOT}/hardware/ril/include/**/*.h"
            )
    # append aosp hardware header
    list(APPEND HEADERS ${HEADERS_HARDWARE})
endif ()


# add aosp out
if (${AOSP_OUT})
    message("${PROJECT_NAME}, enable aosp out")

    # add aosp out header dir
    include_directories(
            "${ANDROID_ROOT}/out/soong/.intermediates/external/libchrome/libchrome-include/gen/gensrcs/external/libchrome"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/libchrome/libmojo_jni_registration_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_static_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/proto_logging/stats/stats_log_api_gen/statslog.h/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/input/classifier/1.0/android.hardware.input.classifier@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/input/common/1.0/android.hardware.input.common@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/input/common/aidl/android.hardware.input.common-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/input/processor/aidl/android.hardware.input.processor-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/StatsD/lib/libkll/proto/libkll-protos/android_arm64_armv8-a-branchprot_static/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_shared/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libsysprop/srcs/libPlatformProperties/android_arm64_armv8-a-branchprot_static/gen/sysprop/include"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
            "${ANDROID_ROOT}/out/soong/.intermediates/external/libchrome/libchrome-include/gen/gensrcs/external/libchrome/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/libchrome/libmojo_jni_registration_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/input/libinput/android_arm64_armv8-a-branchprot_static_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/proto_logging/stats/stats_log_api_gen/statslog.h/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/input/classifier/1.0/android.hardware.input.classifier@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/input/common/1.0/android.hardware.input.common@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/input/common/aidl/android.hardware.input.common-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/input/processor/aidl/android.hardware.input.processor-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/StatsD/lib/libkll/proto/libkll-protos/android_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
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


add_executable(
        InputFlinger
        ${SOURCE_FILES}
        ${HEADERS}
)

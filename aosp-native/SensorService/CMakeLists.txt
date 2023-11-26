cmake_minimum_required(VERSION 3.5)
project(SensorService)

#set(CMAKE_CXX_STANDARD 20)
set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

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

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/1.0/vts/functional/*.cpp
        ${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/aidl/vts/*.cpp
        ${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/libsensorndkbridge/*.cpp
        ${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/utils/*.cc
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl/fuzzer/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/hidl/*.cpp
        ${ANDROID_ROOT}/frameworks/native/services/sensorservice/tests/*.cpp
        )

include_directories(
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/1.0"
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/1.0/vts/functional"
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/aidl/vts"
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/libsensorndkbridge"
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/utils"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl/fuzzer"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl/include"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl/include/sensorserviceaidl"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/hidl"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/hidl/include"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/hidl/include/sensorservicehidl"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/tests"
)

file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/1.0/**/*.h"
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/1.0/vts/functional/**/*.h"
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/aidl/vts/**/*.h"
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/libsensorndkbridge/**/*.h"
        "${ANDROID_ROOT}/frameworks/hardware/interfaces/sensorservice/utils/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl/fuzzer/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/aidl/include/sensorserviceaidl/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/hidl/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/hidl/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/hidl/include/sensorservicehidl/**/*.h"
        "${ANDROID_ROOT}/frameworks/native/services/sensorservice/tests/**/*.h"
        )


# add aosp fwk_base
if (${AOSP_FWK_BASE})
    message("${PROJECT_NAME}, enable aosp fwk_base")

    # add aosp fwk_base header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/base/libs/protoutil/include"
            "${ANDROID_ROOT}/frameworks/base/native/android/include_platform"
    )

    # add aosp fwk_base header file
    file(GLOB_RECURSE HEADERS_FWK_BASE
            "${ANDROID_ROOT}/frameworks/base/libs/protoutil/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/base/native/android/include_platform/**/*.h"
            )
    # append aosp fwk_base header
    list(APPEND HEADERS ${HEADERS_FWK_BASE})
endif ()


# add aosp fwk_native
if (${AOSP_FWK_NATIVE})
    message("${PROJECT_NAME}, enable aosp fwk_native")

    # add aosp fwk_native header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/native/include"
            "${ANDROID_ROOT}/frameworks/native/include_sensor"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_activitymanager"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_batterystats"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/tests/parcel_fuzzer/include_random_parcel"
            "${ANDROID_ROOT}/frameworks/native/libs/fakeservicemanager/include"
            "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include"
            "${ANDROID_ROOT}/frameworks/native/libs/gui/include"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private"
            "${ANDROID_ROOT}/frameworks/native/libs/permission/include"
            "${ANDROID_ROOT}/frameworks/native/libs/sensor/include"
            "${ANDROID_ROOT}/frameworks/native/libs/sensorprivacy/include"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types"
            "${ANDROID_ROOT}/frameworks/native/opengl/include"
            "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include"
            "${ANDROID_ROOT}/frameworks/native/services/sensorservice"
    )

    # add aosp fwk_native header file
    file(GLOB_RECURSE HEADERS_FWK_NATIVE
            "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/include_sensor/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_activitymanager/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include_batterystats/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/tests/parcel_fuzzer/include_random_parcel/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/fakeservicemanager/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/gralloc/types/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/gui/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/permission/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/sensor/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/sensorprivacy/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/ui/include_types/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/libs/EGL/include/**/*.h"
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
            "${ANDROID_ROOT}/system/core/libcutils/include_outside_system"
            "${ANDROID_ROOT}/system/core/libpackagelistparser/include"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include"
            "${ANDROID_ROOT}/system/core/libsystem/include"
            "${ANDROID_ROOT}/system/core/libutils/include"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
            "${ANDROID_ROOT}/system/libbase/include"
            "${ANDROID_ROOT}/system/libfmq/base"
            "${ANDROID_ROOT}/system/libfmq/include"
            "${ANDROID_ROOT}/system/libhidl/base/include"
            "${ANDROID_ROOT}/system/libhidl/gtest_helper"
            "${ANDROID_ROOT}/system/libhidl/transport/include"
            "${ANDROID_ROOT}/system/libhidl/transport/token/1.0/utils/include"
            "${ANDROID_ROOT}/system/libhwbinder/include"
            "${ANDROID_ROOT}/system/logging/liblog/include"
            "${ANDROID_ROOT}/system/logging/liblog/include_vndk"
            "${ANDROID_ROOT}/system/tools/aidl/metadata/include"
            "${ANDROID_ROOT}/system/tools/hidl/metadata/include"
            "${ANDROID_ROOT}/system/tools/hidl/utils/include"
    )

    # add aosp system_common header file
    file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
            "${ANDROID_ROOT}/system/core/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include_outside_system/**/*.h"
            "${ANDROID_ROOT}/system/core/libpackagelistparser/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libsystem/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
            "${ANDROID_ROOT}/system/libbase/include/**/*.h"
            "${ANDROID_ROOT}/system/libfmq/base/**/*.h"
            "${ANDROID_ROOT}/system/libfmq/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/base/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/gtest_helper/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/token/1.0/utils/include/**/*.h"
            "${ANDROID_ROOT}/system/libhwbinder/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include_vndk/**/*.h"
            "${ANDROID_ROOT}/system/tools/aidl/metadata/include/**/*.h"
            "${ANDROID_ROOT}/system/tools/hidl/metadata/include/**/*.h"
            "${ANDROID_ROOT}/system/tools/hidl/utils/include/**/*.h"
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
            "${ANDROID_ROOT}/system/libvintf/libaidlvintf_test_helper/include"
            "${ANDROID_ROOT}/system/media/audio/include"
    )

    # add aosp system header file
    file(GLOB_RECURSE HEADERS_SYSTEM
            "${ANDROID_ROOT}/system/libvintf/**/*.h"
            "${ANDROID_ROOT}/system/libvintf/include/**/*.h"
            "${ANDROID_ROOT}/system/libvintf/libaidlvintf_test_helper/include/**/*.h"
            "${ANDROID_ROOT}/system/media/audio/include/**/*.h"
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
            "${ANDROID_ROOT}/external/googletest/googlemock/include"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/pcre/include"
            "${ANDROID_ROOT}/external/selinux/libselinux/include"
            "${ANDROID_ROOT}/external/tinyxml2"
            "${ANDROID_ROOT}/external/zlib"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/boringssl/src/include/**/*.h"
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/pcre/include/**/*.h"
            "${ANDROID_ROOT}/external/selinux/libselinux/include/**/*.h"
            "${ANDROID_ROOT}/external/tinyxml2/**/*.h"
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
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include"
    )

    # add aosp packages header file
    file(GLOB_RECURSE HEADERS_PACKAGES
            "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types/**/*.h"
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include/**/*.h"
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
            "${ANDROID_ROOT}/hardware/interfaces/sensors/1.0/default/include"
            "${ANDROID_ROOT}/hardware/interfaces/sensors/aidl/convert/include"
            "${ANDROID_ROOT}/hardware/interfaces/sensors/common/convert/include"
            "${ANDROID_ROOT}/hardware/interfaces/sensors/common/utils"
            "${ANDROID_ROOT}/hardware/libhardware/include"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include"
            "${ANDROID_ROOT}/hardware/ril/include"
    )

    # add aosp hardware header file
    file(GLOB_RECURSE HEADERS_HARDWARE
            "${ANDROID_ROOT}/hardware/interfaces/common/support/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/sensors/1.0/default/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/sensors/aidl/convert/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/sensors/common/convert/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/sensors/common/utils/**/*.h"
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

    # add aosp out src
    file(GLOB SOURCE_FILES_OUT
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/1.0/android.frameworks.sensorservice@1.0_genc++/gen/android/frameworks/sensorservice/1.0/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V1-ndk-source/gen/android/frameworks/sensorservice/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V1-ndk_platform-source/gen/android/frameworks/sensorservice/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V2-ndk-source/gen/android/frameworks/sensorservice/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V2-ndk_platform-source/gen/android/frameworks/sensorservice/*.cpp
            )
    # append aosp out src
    list(APPEND SOURCE_FILES ${SOURCE_FILES_OUT})

    # add aosp out header dir
    include_directories(
            "${ANDROID_ROOT}/out/soong/.intermediates/bionic/libc/libc/android_vendor.34_arm64_armv8-a-branchprot_shared/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/framework-cppstream-protos/gen/gensrcs"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/1.0/android.frameworks.sensorservice@1.0_genc++/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/1.0/android.frameworks.sensorservice@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V1-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V2-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libactivitymanager_aidl/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/sysprop/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/sensorprivacy/libsensorprivacy/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/fmq/aidl/android.hardware.common.fmq-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/fmq/aidl/android.hardware.common.fmq-V1-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/1.0/android.hardware.sensors@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/2.0/android.hardware.sensors@2.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/2.1/android.hardware.sensors@2.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/aidl/android.hardware.sensors-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/aidl/android.hardware.sensors-V2-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/allocator/1.0/android.hidl.allocator@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/vendor/meizu/interfaces/sensor_aux/1.0/vendor.meizu.sensor_aux@1.0_genc++_headers/gen"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
            "${ANDROID_ROOT}/out/soong/.intermediates/bionic/libc/libc/android_vendor.34_arm64_armv8-a-branchprot_shared/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/base/framework-cppstream-protos/gen/gensrcs/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/1.0/android.frameworks.sensorservice@1.0_genc++/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/1.0/android.frameworks.sensorservice@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V1-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/hardware/interfaces/sensorservice/aidl/android.frameworks.sensorservice-V2-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libactivitymanager_aidl/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_aidl_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/libgui_window_info_static/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/gui/sysprop/libLibGuiProperties/android_arm64_armv8-a-branchprot_static_afdo-libgui_lto-thin/gen/sysprop/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/permission/framework-permission-aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/sensorprivacy/libsensorprivacy/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/fmq/aidl/android.hardware.common.fmq-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/fmq/aidl/android.hardware.common.fmq-V1-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/1.0/android.hardware.graphics.bufferqueue@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/bufferqueue/2.0/android.hardware.graphics.bufferqueue@2.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.0/android.hardware.graphics.common@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.1/android.hardware.graphics.common@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/1.2/android.hardware.graphics.common@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/mapper/4.0/android.hardware.graphics.mapper@4.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/media/1.0/android.hardware.media@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/1.0/android.hardware.sensors@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/2.0/android.hardware.sensors@2.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/2.1/android.hardware.sensors@2.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/aidl/android.hardware.sensors-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/sensors/aidl/android.hardware.sensors-V2-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/allocator/1.0/android.hidl.allocator@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/vendor/meizu/interfaces/sensor_aux/1.0/vendor.meizu.sensor_aux@1.0_genc++_headers/gen/**/*.h"
            )
    # append aosp out header
    list(APPEND HEADERS ${HEADERS_OUT})
endif ()


# add aosp other
if (${AOSP_other})
    message("${PROJECT_NAME}, enable aosp other")

    # add aosp other header dir
    include_directories(
            "${ANDROID_ROOT}/test/vts/vts_hal_hidl_target"
    )

    # add aosp other header file
    file(GLOB_RECURSE HEADERS_other
            "${ANDROID_ROOT}/test/vts/vts_hal_hidl_target/**/*.h"
            )
    # append aosp other header
    list(APPEND HEADERS ${HEADERS_other})
endif ()


add_executable(
        SensorService
        ${SOURCE_FILES}
        ${HEADERS}
)

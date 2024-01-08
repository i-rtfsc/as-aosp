cmake_minimum_required(VERSION 3.5)
project(AutoMotiveVehicle)

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

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/src/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/tests/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/userhal/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/tests/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/tests/fuzzer/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/utils/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/utils/tests/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/aidl_test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/JsonConfigLoader/src/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/JsonConfigLoader/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/GeneratorHub/src/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/GeneratorHub/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/hardware/src/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/hardware/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/obd2frame/src/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/obd2frame/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/userhal/src/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/userhal/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/utils/proto_message_converter/src/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/utils/proto_message_converter/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/common/src/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/common/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/src/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/test/*.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/vts/src/*.cpp
        )

include_directories(
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include/vhal_v2_0"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/userhal"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/userhal/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/tests"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/utils"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/aidl_test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/generated_lib/cpp"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/JsonConfigLoader"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/JsonConfigLoader/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/JsonConfigLoader/test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/GeneratorHub"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/GeneratorHub/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/GeneratorHub/test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/hardware"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/hardware/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/hardware/test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/obd2frame"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/obd2frame/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/obd2frame/test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/userhal"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/userhal/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/userhal/test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/utils/proto_message_converter"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/utils/proto_message_converter/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/hardware/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/proto"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/common"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/common/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/common/test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/test/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/include"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/test"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/proto"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/vts"
)

file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/common/include/vhal_v2_0/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/userhal/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/impl/vhal_v2_0/userhal/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/default/tests/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/2.0/utils/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/aidl_test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/generated_lib/cpp/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/JsonConfigLoader/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/JsonConfigLoader/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/JsonConfigLoader/test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/default_config/test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/GeneratorHub/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/GeneratorHub/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/GeneratorHub/test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/hardware/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/hardware/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/hardware/test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/obd2frame/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/obd2frame/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/obd2frame/test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/userhal/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/userhal/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/fake_impl/userhal/test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/utils/proto_message_converter/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/utils/proto_message_converter/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/hardware/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/proto/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/common/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/common/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/common/test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/utils/test/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/include/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/aidl/impl/vhal/test/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/proto/**/*.h"
        "${ANDROID_ROOT}/hardware/interfaces/automotive/vehicle/vts/**/*.h"
        )


# add aosp fwk_native
if (${AOSP_FWK_NATIVE})
    message("${PROJECT_NAME}, enable aosp fwk_native")

    # add aosp fwk_native header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/native/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/tests/parcel_fuzzer/include_random_parcel"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include"
            "${ANDROID_ROOT}/frameworks/native/opengl/include"
    )

    # add aosp fwk_native header file
    file(GLOB_RECURSE HEADERS_FWK_NATIVE
            "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/tests/parcel_fuzzer/include_random_parcel/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
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
            "${ANDROID_ROOT}/system/core/libvndksupport/include"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
            "${ANDROID_ROOT}/system/libbase/include"
            "${ANDROID_ROOT}/system/libfmq/base"
            "${ANDROID_ROOT}/system/libhidl/base/include"
            "${ANDROID_ROOT}/system/libhidl/gtest_helper"
            "${ANDROID_ROOT}/system/libhidl/transport/include"
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
            "${ANDROID_ROOT}/system/core/libvndksupport/include/**/*.h"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
            "${ANDROID_ROOT}/system/libbase/include/**/*.h"
            "${ANDROID_ROOT}/system/libfmq/base/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/base/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/gtest_helper/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/include/**/*.h"
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
            "${ANDROID_ROOT}/external/fmtlib/include"
            "${ANDROID_ROOT}/external/googletest/googlemock/include"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/grpc-grpc/include"
            "${ANDROID_ROOT}/external/jsoncpp/include"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/pcre/include"
            "${ANDROID_ROOT}/external/protobuf/src"
            "${ANDROID_ROOT}/external/selinux/libselinux/include"
            "${ANDROID_ROOT}/external/tinyxml2"
            "${ANDROID_ROOT}/external/zlib"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/grpc-grpc/include/**/*.h"
            "${ANDROID_ROOT}/external/jsoncpp/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/pcre/include/**/*.h"
            "${ANDROID_ROOT}/external/protobuf/src/**/*.h"
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
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include"
            "${ANDROID_ROOT}/packages/services/Car/cpp/car_binder_lib/largeParcelable/include"
            "${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/include"
    )

    # add aosp packages header file
    file(GLOB_RECURSE HEADERS_PACKAGES
            "${ANDROID_ROOT}/packages/modules/Gki/libkver/include/**/*.h"
            "${ANDROID_ROOT}/packages/services/Car/cpp/car_binder_lib/largeParcelable/include/**/*.h"
            "${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/include/**/*.h"
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

    # add aosp out src
    file(GLOB SOURCE_FILES_OUT
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/2.0/android.hardware.automotive.vehicle@2.0_genc++/gen/android/hardware/automotive/vehicle/2.0/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V1-ndk-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V2-ndk-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V3-ndk-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_cc@default-grpc/gen/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/proto/VehicleProtoStub_cc/gen/android/hardware/automotive/vehicle/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/proto/android.hardware.automotive.vehicle@2.0-libproto-native/android_vendor.34_arm64_armv8-a-branchprot_static/gen/proto/hardware/interfaces/automotive/vehicle/proto/*.cc
            )
    # append aosp out src
    list(APPEND SOURCE_FILES ${SOURCE_FILES_OUT})

    # add aosp out header dir
    include_directories(
            "${ANDROID_ROOT}/out/soong/.intermediates/bionic/libc/libc/android_vendor.34_arm64_armv8-a-branchprot_shared/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_vendor.34_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/2.0/android.hardware.automotive.vehicle@2.0_genc++/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/2.0/android.hardware.automotive.vehicle@2.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V1-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V2-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V3-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V3-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_cc@default-grpc/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_h@default-grpc/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/proto/VehicleProtoStub_cc/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/proto/VehicleProtoStub_h/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/proto/android.hardware.automotive.vehicle@2.0-libproto-native/android_vendor.34_arm64_armv8-a-branchprot_static/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/proto/android.hardware.automotive.vehicle@2.0-libproto-native/android_vendor.34_arm64_armv8-a-branchprot_static/gen/proto/hardware/interfaces/automotive/vehicle/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/proto/android.hardware.automotive.vehicle@2.0-libproto-native/android_vendor.34_arm64_armv8-a-branchprot_static_cfi/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/proto/android.hardware.automotive.vehicle@2.0-libproto-native/android_vendor.34_arm64_armv8-a-branchprot_static_cfi/gen/proto/hardware/interfaces/automotive/vehicle/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
            "${ANDROID_ROOT}/out/soong/.intermediates/bionic/libc/libc/android_vendor.34_arm64_armv8-a-branchprot_shared/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_vendor.34_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/2.0/android.hardware.automotive.vehicle@2.0_genc++/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/2.0/android.hardware.automotive.vehicle@2.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V1-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V2-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V3-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/android.hardware.automotive.vehicle-V3-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_cc@default-grpc/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/grpc/VehicleServerProtoStub_h@default-grpc/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/proto/VehicleProtoStub_cc/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl/impl/proto/VehicleProtoStub_h/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/proto/android.hardware.automotive.vehicle@2.0-libproto-native/android_vendor.34_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/proto/android.hardware.automotive.vehicle@2.0-libproto-native/android_vendor.34_arm64_armv8-a-branchprot_static/gen/proto/hardware/interfaces/automotive/vehicle/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/proto/android.hardware.automotive.vehicle@2.0-libproto-native/android_vendor.34_arm64_armv8-a-branchprot_static_cfi/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/automotive/vehicle/proto/android.hardware.automotive.vehicle@2.0-libproto-native/android_vendor.34_arm64_armv8-a-branchprot_static_cfi/gen/proto/hardware/interfaces/automotive/vehicle/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen/**/*.h"
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
        AutoMotiveVehicle
        ${SOURCE_FILES}
        ${HEADERS}
)

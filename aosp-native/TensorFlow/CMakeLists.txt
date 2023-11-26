cmake_minimum_required(VERSION 3.5)
project(TensorFlow)

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
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/eager/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/compiler/jit/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/device/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/example/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/core/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/hash/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/histogram/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/monitoring/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/random/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/strings/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/wav/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/nccl/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/profile_utils/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/profiler/lib/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/tfrt/fallback/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/sparse/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/tensor_bundle/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/c/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/core/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/core/api/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/core/shims/jni/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/flex/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/nnapi/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/utils/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/experimental/resource/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/optimized/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/reference/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/internal/utils/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/nnapi/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/nnapi/sl/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/profiling/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/schema/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/testing/nnapi_tflite_zip_tests/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/tools/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/tools/delegates/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/tools/versioning/*.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/base/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/container/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/debugging/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/flags/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/hash/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/numeric/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/profiling/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/random/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/status/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/strings/internal/str_format/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/synchronization/internal/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/time/internal/cctz/src/*.cc
        ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/absl/types/*.cc
        )

include_directories(
        "${ANDROID_ROOT}/external/tensorflow"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/testing/nnapi_tflite_zip_tests"
        "${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp"
)

file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/external/tensorflow/**/*.h"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/**/*.h"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/**/*.h"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/**/*.h"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/java/src/main/native/**/*.h"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/**/*.h"
        "${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/testing/nnapi_tflite_zip_tests/**/*.h"
        "${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp/**/*.h"
        )


# add aosp fwk_base
if (${AOSP_FWK_BASE})
    message("${PROJECT_NAME}, enable aosp fwk_base")

    # add aosp fwk_base header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/base/native/android/include_platform"
    )

    # add aosp fwk_base header file
    file(GLOB_RECURSE HEADERS_FWK_BASE
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
            "${ANDROID_ROOT}/frameworks/native/cmds/lshal/libprocpartition/include"
            "${ANDROID_ROOT}/frameworks/native/include"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private"
            "${ANDROID_ROOT}/frameworks/native/opengl/include"
    )

    # add aosp fwk_native header file
    file(GLOB_RECURSE HEADERS_FWK_NATIVE
            "${ANDROID_ROOT}/frameworks/native/cmds/lshal/libprocpartition/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include/**/*.h"
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
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
            "${ANDROID_ROOT}/system/libbase/include"
            "${ANDROID_ROOT}/system/logging/liblog/include"
    )

    # add aosp system_common header file
    file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
            "${ANDROID_ROOT}/system/core/include/**/*.h"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
            "${ANDROID_ROOT}/system/libbase/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
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
            "${ANDROID_ROOT}/system/media/audio/include"
    )

    # add aosp system header file
    file(GLOB_RECURSE HEADERS_SYSTEM
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
            "${ANDROID_ROOT}/external/FP16"
            "${ANDROID_ROOT}/external/FP16/include"
            "${ANDROID_ROOT}/external/double-conversion"
            "${ANDROID_ROOT}/external/eigen"
            "${ANDROID_ROOT}/external/fft2d/src/fft2d"
            "${ANDROID_ROOT}/external/flatbuffers/include"
            "${ANDROID_ROOT}/external/fmtlib/include"
            "${ANDROID_ROOT}/external/gemmlowp"
            "${ANDROID_ROOT}/external/googletest/googlemock/include"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/libtextclassifier/native"
            "${ANDROID_ROOT}/external/protobuf/src"
            "${ANDROID_ROOT}/external/ruy"
            "${ANDROID_ROOT}/external/zlib"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/FP16/**/*.h"
            "${ANDROID_ROOT}/external/FP16/include/**/*.h"
            "${ANDROID_ROOT}/external/double-conversion/**/*.h"
            "${ANDROID_ROOT}/external/eigen/**/*.h"
            "${ANDROID_ROOT}/external/fft2d/src/fft2d/**/*.h"
            "${ANDROID_ROOT}/external/flatbuffers/include/**/*.h"
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/gemmlowp/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/libtextclassifier/native/**/*.h"
            "${ANDROID_ROOT}/external/protobuf/src/**/*.h"
            "${ANDROID_ROOT}/external/ruy/**/*.h"
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
            "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/include"
    )

    # add aosp packages header file
    file(GLOB_RECURSE HEADERS_PACKAGES
            "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/include/**/*.h"
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
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/google/protobuf/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/google/protobuf/compiler/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/compiler/xla/pjrt/distributed/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/compiler/xla/stream_executor/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/core/example/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/core/framework/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/core/grappler/costs/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/core/profiler/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/core/profiler/protobuf/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/core/protobuf/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/core/util/*.cc
            ${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tflite_configuration_proto/android_arm64_armv8-a-branchprot_static/gen/proto/tensorflow/lite/experimental/acceleration/configuration/*.cc
            )
    # append aosp out src
    list(APPEND SOURCE_FILES ${SOURCE_FILES_OUT})

    # add aosp out header dir
    include_directories(
            "${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/libtflite_mutable_schema/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tflite_configuration_proto/android_arm64_armv8-a-branchprot_static/gen/proto"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
            "${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/libtflite_mutable_schema/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tensorflow_core_proto_cpp_lite/android_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/tensorflow/tflite_configuration_proto/android_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
            )
    # append aosp out header
    list(APPEND HEADERS ${HEADERS_OUT})
endif ()


add_executable(
        TensorFlow
        ${SOURCE_FILES}
        ${HEADERS}
)

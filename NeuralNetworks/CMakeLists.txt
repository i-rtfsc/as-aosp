cmake_minimum_required(VERSION 3.5)
project(NeuralNetworks)

# set(CMAKE_CXX_STANDARD 20)

#set(ANDROID_ROOT "/home/solo/code/aosp")
set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

set(AOSP_SYSTEM_COMMON true)
set(AOSP_AV false)
set(AOSP_BIONIC false)
set(AOSP_SYSTEM false)
set(AOSP_EXTERNAL false)
set(AOSP_PACKAGES false)
set(AOSP_HARDWARE false)
set(AOSP_OUT false)

file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/cpu_operations/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/random/*.cc
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/src/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/src/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/cache/BlobCache/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/cache/nnCache/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_shim/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/operation_converters/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/packageinfo/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/fibonacci_extension/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/fuzzing/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/fuzzing/operation_signatures/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/generated/spec_AIDL_V2/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/generated/spec_AIDL_V3/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/generated/spec_V1_0/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/generated/spec_V1_1/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/generated/spec_V1_2/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/generated/spec_V1_3/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/generated/spec_V1_3_cts_only/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/generated/spec_experimental/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/ion_watcher/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/nnapi_info/*.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/test_generator/test_harness/*.cpp
        )

include_directories(
        "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/include"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/random"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include/nnapi"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/include"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/cache/BlobCache"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/cache/nnCache"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_shim"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/extensions/example/fibonacci"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/include"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/packageinfo"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/packageinfo/include"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/android_fuzzing"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/fuzzing"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/include"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/public"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/ion_watcher"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/nnapi_info"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/test_generator/test_harness"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/test_generator/test_harness/include"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/test_generator/test_harness/include/nnapi"
        "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatssocket/include"
)

file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/packages/modules/Bluetooth/system/types/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/random/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/include/nnapi/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/common/types/operations/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/cache/BlobCache/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/cache/nnCache/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_shim/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/extensions/example/fibonacci/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/packageinfo/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/packageinfo/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/test/fuzzing/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/public/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/ion_watcher/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/nnapi_info/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/test_generator/test_harness/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/test_generator/test_harness/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/NeuralNetworks/tools/test_generator/test_harness/include/nnapi/**/*.h"
        "${ANDROID_ROOT}/packages/modules/StatsD/lib/libstatssocket/include/**/*.h"
        )


# add aosp framework
if (${AOSP_FKW})
    message("enable aosp framework")

    # add aosp framework header dir
    include_directories(
            "${ANDROID_ROOT}/build/soong/cc/libbuildversion/include"
            "${ANDROID_ROOT}/frameworks/base/native/android/include_platform"
            "${ANDROID_ROOT}/frameworks/native/cmds/lshal/libprocpartition/include"
            "${ANDROID_ROOT}/frameworks/native/include"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/tests/parcel_fuzzer/include_random_parcel"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private"
            "${ANDROID_ROOT}/frameworks/native/opengl/include"
    )

    # add aosp framework header file
    file(GLOB_RECURSE HEADERS_FKW
            "${ANDROID_ROOT}/build/soong/cc/libbuildversion/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/base/native/android/include_platform/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/cmds/lshal/libprocpartition/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/tests/parcel_fuzzer/include_random_parcel/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/math/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
            )
    # append aosp framework header
    list(APPEND HEADERS ${HEADERS_FKW})
endif ()


# add aosp system common
if (${AOSP_SYSTEM_COMMON})
    message("enable aosp system common")

    # add aosp system common header dir
    include_directories(
            "${ANDROID_ROOT}/libnativehelper/include_jni"
            "${ANDROID_ROOT}/system/core/include"
            "${ANDROID_ROOT}/system/core/libcutils/include"
            "${ANDROID_ROOT}/system/core/libcutils/include_outside_system"
            "${ANDROID_ROOT}/system/core/libprocessgroup/cgrouprc/include"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include"
            "${ANDROID_ROOT}/system/core/libsystem/include"
            "${ANDROID_ROOT}/system/core/libutils/include"
            "${ANDROID_ROOT}/system/core/libvndksupport/include"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
            "${ANDROID_ROOT}/system/libbase/include"
            "${ANDROID_ROOT}/system/libfmq/base"
            "${ANDROID_ROOT}/system/libfmq/include"
            "${ANDROID_ROOT}/system/libhidl/base/include"
            "${ANDROID_ROOT}/system/libhidl/libhidlmemory/include"
            "${ANDROID_ROOT}/system/libhidl/transport/include"
            "${ANDROID_ROOT}/system/libhwbinder/include"
            "${ANDROID_ROOT}/system/logging/liblog/include"
            "${ANDROID_ROOT}/system/logging/liblog/include_vndk"
    )

    # add aosp system common header file
    file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
            "${ANDROID_ROOT}/libnativehelper/include_jni/**/*.h"
            "${ANDROID_ROOT}/system/core/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include_outside_system/**/*.h"
            "${ANDROID_ROOT}/system/core/libprocessgroup/cgrouprc/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libsystem/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libvndksupport/include/**/*.h"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
            "${ANDROID_ROOT}/system/libbase/include/**/*.h"
            "${ANDROID_ROOT}/system/libfmq/base/**/*.h"
            "${ANDROID_ROOT}/system/libfmq/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/base/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/libhidlmemory/include/**/*.h"
            "${ANDROID_ROOT}/system/libhidl/transport/include/**/*.h"
            "${ANDROID_ROOT}/system/libhwbinder/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include_vndk/**/*.h"
            )
    # append aosp system common header
    list(APPEND HEADERS ${HEADERS_SYSTEM_COMMON})
endif ()


# add aosp av
if (${AOSP_AV})
    message("enable aosp av")

    # add aosp av header dir
    include_directories(
            "${ANDROID_ROOT}/frameworks/av/include"
    )

    # add aosp av header file
    file(GLOB_RECURSE HEADERS_AV
            "${ANDROID_ROOT}/frameworks/av/include/**/*.h"
            )
    # append aosp av header
    list(APPEND HEADERS ${HEADERS_AV})
endif ()


# add aosp bionic
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


# add aosp system
if (${AOSP_SYSTEM})
    message("enable aosp system")

    # add aosp system header dir
    include_directories(
            "${ANDROID_ROOT}/system/media/audio/include"
            "${ANDROID_ROOT}/system/memory/libmemunreachable/include"
            "${ANDROID_ROOT}/system/server_configurable_flags/libflags/include"
    )

    # add aosp system header file
    file(GLOB_RECURSE HEADERS_SYSTEM
            "${ANDROID_ROOT}/system/media/audio/include/**/*.h"
            "${ANDROID_ROOT}/system/memory/libmemunreachable/include/**/*.h"
            "${ANDROID_ROOT}/system/server_configurable_flags/libflags/include/**/*.h"
            )
    # append aosp system header
    list(APPEND HEADERS ${HEADERS_SYSTEM})
endif ()


# add aosp external
if (${AOSP_EXTERNAL})
    message("enable aosp external")

    # add aosp external header dir
    include_directories(
            "${ANDROID_ROOT}/external/XNNPACK/include"
            "${ANDROID_ROOT}/external/boringssl/src/include"
            "${ANDROID_ROOT}/external/cpuinfo/deps/clog/include"
            "${ANDROID_ROOT}/external/cpuinfo/include"
            "${ANDROID_ROOT}/external/eigen"
            "${ANDROID_ROOT}/external/fft2d/src/fft2d"
            "${ANDROID_ROOT}/external/flatbuffers/include"
            "${ANDROID_ROOT}/external/fmtlib/include"
            "${ANDROID_ROOT}/external/gemmlowp"
            "${ANDROID_ROOT}/external/googletest/googlemock/include"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/jsoncpp/include"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/libprotobuf-mutator"
            "${ANDROID_ROOT}/external/libtextclassifier/native"
            "${ANDROID_ROOT}/external/protobuf/src"
            "${ANDROID_ROOT}/external/pthreadpool/include"
            "${ANDROID_ROOT}/external/ruy"
            "${ANDROID_ROOT}/external/tensorflow"
            "${ANDROID_ROOT}/external/vulkan-headers/include"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/XNNPACK/include/**/*.h"
            "${ANDROID_ROOT}/external/boringssl/src/include/**/*.h"
            "${ANDROID_ROOT}/external/cpuinfo/deps/clog/include/**/*.h"
            "${ANDROID_ROOT}/external/cpuinfo/include/**/*.h"
            "${ANDROID_ROOT}/external/eigen/**/*.h"
            "${ANDROID_ROOT}/external/fft2d/src/fft2d/**/*.h"
            "${ANDROID_ROOT}/external/flatbuffers/include/**/*.h"
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/gemmlowp/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/jsoncpp/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/libprotobuf-mutator/**/*.h"
            "${ANDROID_ROOT}/external/libtextclassifier/native/**/*.h"
            "${ANDROID_ROOT}/external/protobuf/src/**/*.h"
            "${ANDROID_ROOT}/external/pthreadpool/include/**/*.h"
            "${ANDROID_ROOT}/external/ruy/**/*.h"
            "${ANDROID_ROOT}/external/tensorflow/**/*.h"
            "${ANDROID_ROOT}/external/vulkan-headers/include/**/*.h"
            )
    # append aosp external header
    list(APPEND HEADERS ${HEADERS_EXTERNAL})
endif ()


# add aosp hardware
if (${AOSP_HARDWARE})
    message("enable aosp hardware")

    # add aosp hardware header dir
    include_directories(
            "${ANDROID_ROOT}/hardware/interfaces/common/support/include"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/include"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.1/utils/include"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.2/utils/include"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.3/utils/include"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/aidl/utils/include"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/adapter/aidl/include"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/adapter/hidl/include"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/common/include"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/service/include"
            "${ANDROID_ROOT}/hardware/libhardware/include"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include"
            "${ANDROID_ROOT}/hardware/ril/include"
    )

    # add aosp hardware header file
    file(GLOB_RECURSE HEADERS_HARDWARE
            "${ANDROID_ROOT}/hardware/interfaces/common/support/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.1/utils/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.2/utils/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.3/utils/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/aidl/utils/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/adapter/aidl/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/adapter/hidl/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/common/include/**/*.h"
            "${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/service/include/**/*.h"
            "${ANDROID_ROOT}/hardware/libhardware/include/**/*.h"
            "${ANDROID_ROOT}/hardware/libhardware_legacy/include/**/*.h"
            "${ANDROID_ROOT}/hardware/ril/include/**/*.h"
            )
    # append aosp hardware header
    list(APPEND HEADERS ${HEADERS_HARDWARE})
endif ()


# add aosp out
if (${AOSP_OUT})
    message("enable aosp out")

    # add aosp out src
    file(GLOB SOURCE_FILES_OUT
            ${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/libneuralnetworks_fuzzer_proto/android_arm64_armv8-a-branchprot_static/gen/proto/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/*.cc
            )
    # append aosp out src
    list(APPEND SOURCE_FILES ${SOURCE_FILES_OUT})

    # add aosp out header dir
    include_directories(
            "${ANDROID_ROOT}/out/soong/.intermediates/bionic/libc/libc/android_vendor.34_arm64_armv8-a-branchprot_shared/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/1.0/android.hardware.neuralnetworks@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/1.1/android.hardware.neuralnetworks@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/1.2/android.hardware.neuralnetworks@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/1.3/android.hardware.neuralnetworks@1.3_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.h/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/libneuralnetworks_fuzzer_proto/android_arm64_armv8-a-branchprot_static/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/libneuralnetworks_fuzzer_proto/android_arm64_armv8-a-branchprot_static/gen/proto/packages/modules/NeuralNetworks/runtime/test/android_fuzzing"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/libneuralnetworks_fuzzer_proto/android_arm64_armv8-a-branchprot_static_fuzzer/gen/proto"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/libneuralnetworks_fuzzer_proto/android_arm64_armv8-a-branchprot_static_fuzzer/gen/proto/packages/modules/NeuralNetworks/runtime/test/android_fuzzing"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/allocator/1.0/android.hidl.allocator@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/memory/1.0/android.hidl.memory@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/memory/token/1.0/android.hidl.memory.token@1.0_genc++_headers/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/safe_union/1.0/android.hidl.safe_union@1.0_genc++_headers/gen"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
            "${ANDROID_ROOT}/out/soong/.intermediates/bionic/libc/libc/android_vendor.34_arm64_armv8-a-branchprot_shared/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared_fuzzer/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/packagemanager_aidl-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/common/aidl/android.hardware.common-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/1.0/android.hardware.neuralnetworks@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/1.1/android.hardware.neuralnetworks@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/1.2/android.hardware.neuralnetworks@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/1.3/android.hardware.neuralnetworks@1.3_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.h/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/libneuralnetworks_fuzzer_proto/android_arm64_armv8-a-branchprot_static/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/libneuralnetworks_fuzzer_proto/android_arm64_armv8-a-branchprot_static/gen/proto/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/libneuralnetworks_fuzzer_proto/android_arm64_armv8-a-branchprot_static_fuzzer/gen/proto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/libneuralnetworks_fuzzer_proto/android_arm64_armv8-a-branchprot_static_fuzzer/gen/proto/packages/modules/NeuralNetworks/runtime/test/android_fuzzing/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/allocator/1.0/android.hidl.allocator@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.0/android.hidl.manager@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.1/android.hidl.manager@1.1_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/manager/1.2/android.hidl.manager@1.2_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/memory/1.0/android.hidl.memory@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/memory/token/1.0/android.hidl.memory.token@1.0_genc++_headers/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/system/libhidl/transport/safe_union/1.0/android.hidl.safe_union@1.0_genc++_headers/gen/**/*.h"
            )
    # append aosp out header
    list(APPEND HEADERS ${HEADERS_OUT})
endif ()


add_executable(
        NeuralNetworks
        ${SOURCE_FILES}
        ${HEADERS}
)

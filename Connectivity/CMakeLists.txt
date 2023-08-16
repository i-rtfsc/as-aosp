cmake_minimum_required(VERSION 3.5)
project(Connectivity)

# set(CMAKE_CXX_STANDARD 20)

#set(ANDROID_ROOT "/home/solo/code/aosp")
set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})

set(AOSP_SYSTEM_COMMON true)
set(AOSP_FWK true)
set(AOSP_AV false)
set(AOSP_BIONIC false)
set(AOSP_SYSTEM false)
set(AOSP_EXTERNAL false)
set(AOSP_HARDWARE false)
set(AOSP_OUT false)

# add aosp packages src
file(GLOB SOURCE_FILES
        ${ANDROID_ROOT}/packages/modules/Connectivity/Tethering/jni/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/framework-t/jni/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/framework/jni/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/netd/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/service-t/jni/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/service-t/native/libs/libnetworkstats/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/service/jni/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/service/jni/com_android_net_module_util/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/service/libconnectivity/src/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/service/native/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/service/native/libs/libclat/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/jni/*.c
        ${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/jni/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/native/dns/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/native/src/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/tests/mts/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/tests/native/connectivity_native_test/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/tests/native/utilities/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/tests/unit/jni/*.cpp
        ${ANDROID_ROOT}/packages/modules/Connectivity/tests/unit/jni/android_net_frameworktests_util/*.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_headers/*.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/native/bpfmapjni/*.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/native/bpfutiljni/*.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/native/ip_checksum/*.c
        ${ANDROID_ROOT}/frameworks/libs/net/common/native/netjniutils/*.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/native/nettestutils/*.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/native/tcutils/*.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/native/tcutils/tests/*.cpp
        ${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/*.cpp
        )

# add aosp packages header dir
include_directories(
        "${ANDROID_ROOT}/packages/modules/Connectivity/Tethering"
        "${ANDROID_ROOT}/packages/modules/Connectivity/bpf_progs"
        "${ANDROID_ROOT}/packages/modules/Connectivity/framework"
        "${ANDROID_ROOT}/packages/modules/Connectivity/framework-t"
        "${ANDROID_ROOT}/packages/modules/Connectivity/netd"
        "${ANDROID_ROOT}/packages/modules/Connectivity/netd/include"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service-t/native/libs/libnetworkstats"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service-t/native/libs/libnetworkstats/include"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/libconnectivity"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/libconnectivity/include"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/native"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/native/include"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/native/libs/libclat"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/native/libs/libclat/include"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/jni"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/native"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/native/dns"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/mts"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/native/connectivity_native_test"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/native/utilities"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/unit/jni"
        "${ANDROID_ROOT}/packages/modules/DnsResolver/include/dnsproxyd_protocol"
        "${ANDROID_ROOT}/frameworks/libs/modules-utils/build/include"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_headers"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_headers/include"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_syscall_wrappers/include"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpfmapjni"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpfutiljni"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/ip_checksum"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/netjniutils"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/netjniutils/include"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/nettestutils"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/nettestutils/include"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/tcutils"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/tcutils/include"
        "${ANDROID_ROOT}/frameworks/libs/net/common/netd"
        "${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils"
        "${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/include"
)

# add aosp packages header file
file(GLOB_RECURSE HEADERS
        "${ANDROID_ROOT}/packages/modules/Connectivity/Tethering/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/bpf_progs/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/framework-t/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/framework/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/netd/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/netd/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service-t/native/libs/libnetworkstats/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service-t/native/libs/libnetworkstats/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/libconnectivity/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/libconnectivity/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/native/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/native/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/native/libs/libclat/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/service/native/libs/libclat/include/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/jni/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/native/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/cts/net/native/dns/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/mts/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/native/connectivity_native_test/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/native/utilities/**/*.h"
        "${ANDROID_ROOT}/packages/modules/Connectivity/tests/unit/jni/**/*.h"
        "${ANDROID_ROOT}/packages/modules/DnsResolver/include/dnsproxyd_protocol/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_headers/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_headers/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_syscall_wrappers/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpfmapjni/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpfutiljni/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/ip_checksum/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/netjniutils/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/netjniutils/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/nettestutils/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/nettestutils/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/tcutils/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/native/tcutils/include/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/netd/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/**/*.h"
        "${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/include/**/*.h"
        )

# add aosp fwk
if (${AOSP_FWK})
    message("enable aosp fwk")

    include_directories(
            "${ANDROID_ROOT}/frameworks/base/native/android/include_platform"
            "${ANDROID_ROOT}/frameworks/libs/modules-utils/build/include"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_headers/include"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_syscall_wrappers/include"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/ip_checksum"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/netjniutils/include"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/tcutils/include"
            "${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/include"
            "${ANDROID_ROOT}/frameworks/native/include"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private"
            "${ANDROID_ROOT}/frameworks/native/opengl/include"
    )

    file(GLOB_RECURSE HEADERS_FWK
            "${ANDROID_ROOT}/frameworks/base/native/android/include_platform/**/*.h"
            "${ANDROID_ROOT}/frameworks/libs/modules-utils/build/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_headers/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/bpf_syscall_wrappers/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/ip_checksum/**/*.h"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/netjniutils/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/libs/net/common/native/tcutils/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/libs/net/common/netd/libnetdutils/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/arect/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativebase/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include-private/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/libs/nativewindow/include/**/*.h"
            "${ANDROID_ROOT}/frameworks/native/opengl/include/**/*.h"
            )
    # append aosp fwk header
    list(APPEND HEADERS ${HEADERS_FWK})
endif ()

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
            "${ANDROID_ROOT}/system/core/include"
            "${ANDROID_ROOT}/system/core/libcutils/include"
            "${ANDROID_ROOT}/system/core/libnetutils/include"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include"
            "${ANDROID_ROOT}/system/core/libsystem/include"
            "${ANDROID_ROOT}/system/core/libutils/include"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include"
            "${ANDROID_ROOT}/system/libbase/include"
            "${ANDROID_ROOT}/system/logging/liblog/include"
            "${ANDROID_ROOT}/system/tools/aidl/analyzer/include"
    )

    # add aosp system common header file
    file(GLOB_RECURSE HEADERS_SYSTEM_COMMON
            "${ANDROID_ROOT}/libnativehelper/header_only_include/**/*.h"
            "${ANDROID_ROOT}/libnativehelper/include/**/*.h"
            "${ANDROID_ROOT}/libnativehelper/include_jni/**/*.h"
            "${ANDROID_ROOT}/libnativehelper/include_platform/**/*.h"
            "${ANDROID_ROOT}/libnativehelper/include_platform_header_only/**/*.h"
            "${ANDROID_ROOT}/system/core/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libcutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libnetutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libprocessgroup/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libsystem/include/**/*.h"
            "${ANDROID_ROOT}/system/core/libutils/include/**/*.h"
            "${ANDROID_ROOT}/system/core/property_service/libpropertyinfoparser/include/**/*.h"
            "${ANDROID_ROOT}/system/libbase/include/**/*.h"
            "${ANDROID_ROOT}/system/logging/liblog/include/**/*.h"
            "${ANDROID_ROOT}/system/tools/aidl/analyzer/include/**/*.h"
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
            "${ANDROID_ROOT}/system/bpf/libbpf_android/include"
            "${ANDROID_ROOT}/system/media/audio/include"
            "${ANDROID_ROOT}/system/netd/include/mainline"
            "${ANDROID_ROOT}/system/netd/tests"
    )

    # add aosp system header file
    file(GLOB_RECURSE HEADERS_SYSTEM
            "${ANDROID_ROOT}/system/bpf/libbpf_android/include/**/*.h"
            "${ANDROID_ROOT}/system/media/audio/include/**/*.h"
            "${ANDROID_ROOT}/system/netd/include/mainline/**/*.h"
            "${ANDROID_ROOT}/system/netd/tests/**/*.h"
            )
    # append aosp system header
    list(APPEND HEADERS ${HEADERS_SYSTEM})
endif ()


# add aosp external
if (${AOSP_EXTERNAL})
    message("enable aosp external")

    # add aosp external header dir
    include_directories(
            "${ANDROID_ROOT}/external/bcc/src/cc/includes"
            "${ANDROID_ROOT}/external/fmtlib/include"
            "${ANDROID_ROOT}/external/googletest/googlemock/include"
            "${ANDROID_ROOT}/external/googletest/googletest/include"
            "${ANDROID_ROOT}/external/libcxx/include"
            "${ANDROID_ROOT}/external/libcxxabi/include"
            "${ANDROID_ROOT}/external/perfetto/include"
            "${ANDROID_ROOT}/external/perfetto/include/perfetto/base/build_configs/android_tree"
            "${ANDROID_ROOT}/external/protobuf/src"
    )

    # add aosp external header file
    file(GLOB_RECURSE HEADERS_EXTERNAL
            "${ANDROID_ROOT}/external/bcc/src/cc/includes/**/*.h"
            "${ANDROID_ROOT}/external/fmtlib/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googlemock/include/**/*.h"
            "${ANDROID_ROOT}/external/googletest/googletest/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxx/include/**/*.h"
            "${ANDROID_ROOT}/external/libcxxabi/include/**/*.h"
            "${ANDROID_ROOT}/external/perfetto/include/**/*.h"
            "${ANDROID_ROOT}/external/perfetto/include/perfetto/base/build_configs/android_tree/**/*.h"
            "${ANDROID_ROOT}/external/protobuf/src/**/*.h"
            )
    # append aosp external header
    list(APPEND HEADERS ${HEADERS_EXTERNAL})
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
            )
    # append aosp hardware header
    list(APPEND HEADERS ${HEADERS_HARDWARE})
endif ()


# add aosp out
if (${AOSP_OUT})
    message("enable aosp out")

    # add aosp out src
    file(GLOB SOURCE_FILES_OUT
            ${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V1-cpp-source/gen/android/net/connectivity/aidl/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V1-ndk-source/gen/android/net/connectivity/aidl/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V1-ndk_platform-source/gen/android/net/connectivity/aidl/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V2-cpp-source/gen/android/net/connectivity/aidl/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V2-ndk-source/gen/android/net/connectivity/aidl/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V2-ndk_platform-source/gen/android/net/connectivity/aidl/*.cpp
            ${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-cpp-analyzer-source/gen/android/net/connectivity/aidl/*.cpp
            )
    # append aosp out src
    list(APPEND SOURCE_FILES ${SOURCE_FILES_OUT})

    # add aosp out header dir
    include_directories(
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_lite_gen_headers/gen/external/perfetto/protos"
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
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_cpp_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_cpp_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_lite_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_lite_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_zero_gen_headers/gen/external/perfetto"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_zero_gen_headers/gen/external/perfetto/protos"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_src_base_version_gen_h/gen"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V1-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V1-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V1-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V2-cpp-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V2-ndk-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V2-ndk_platform-source/gen/include"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-cpp-analyzer-source/gen/include"
    )

    # add aosp out header file
    file(GLOB_RECURSE HEADERS_OUT
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_common_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_android_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_ftrace_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_gpu_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_inode_file_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_interceptors_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_power_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_process_stats_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_profiling_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_statsd_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_sys_stats_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_system_info_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_config_track_event_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
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
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_android_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_chrome_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_filesystem_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ftrace_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_gpu_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_interned_data_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_minimal_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_non_minimal_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_perfetto_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_power_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_profiling_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_ps_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_statsd_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_sys_stats_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_system_info_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_cpp_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_cpp_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_track_event_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_lite_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_lite_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_zero_gen_headers/gen/external/perfetto/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_protos_perfetto_trace_translation_zero_gen_headers/gen/external/perfetto/protos/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/external/perfetto/perfetto_src_base_version_gen_h/gen/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/frameworks/native/libs/binder/libbinder/android_arm64_armv8-a-branchprot_shared/gen/aidl/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V1-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V1-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V1-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V2-cpp-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V2-ndk-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-V2-ndk_platform-source/gen/include/**/*.h"
            "${ANDROID_ROOT}/out/soong/.intermediates/packages/modules/Connectivity/service/connectivity_native_aidl_interface-cpp-analyzer-source/gen/include/**/*.h"
            )
    # append aosp out header
    list(APPEND HEADERS ${HEADERS_OUT})
endif ()


add_executable(
        Connectivity
        ${SOURCE_FILES}
        ${HEADERS}
)

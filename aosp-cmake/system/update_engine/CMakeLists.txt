# Copyright (c) 2024 anqi.huang@outlook.com
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use self file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

cmake_minimum_required(VERSION 3.5)

project(update_engine)

# Check if the major CMakeLists.txt file has correctly set
# the ANDROID_ROOT directory.
if (IS_DIRECTORY ${ANDROID_ROOT})
    message(STATUS "${PROJECT_NAME}, has been set android root")
else ()
    if (IS_DIRECTORY ${BUILD_NATIVE_ROOT})
        message(WARNING "${PROJECT_NAME}, set ANDROID_ROOT via BUILD_NATIVE_ROOT=${BUILD_NATIVE_ROOT}")
        set(ANDROID_ROOT ${BUILD_NATIVE_ROOT})
        set(OUT_INTERMEDIATES_ROOT ${ANDROID_ROOT}/out/soong/.intermediates)
    else ()
        message(FATAL_ERROR "${PROJECT_NAME}, ANDROID_ROOT was NULL")
    endif ()
endif ()

#################### update_metadata-protos ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto/update_engine/update_metadata.pb.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(update_metadata-protos SHARED ${SOURCE_FILES})

    target_include_directories(update_metadata-protos
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(update_metadata-protos PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(update_metadata-protos
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### update_metadata-protos ####################

#################### update_engine_client ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_engine_client/android_${OUT_ARCH_CPU}/gen/aidl/android/os/IUpdateEngine.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_engine_client/android_${OUT_ARCH_CPU}/gen/aidl/android/os/IUpdateEngineCallback.cpp
        ${ANDROID_ROOT}/system/update_engine/aosp/update_engine_client_android.cc
        ${ANDROID_ROOT}/system/update_engine/common/error_code_utils.cc
        ${ANDROID_ROOT}/system/update_engine/update_status_utils.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_executable(update_engine_client ${SOURCE_FILES})

    target_include_directories(update_engine_client
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_engine_client/android_${OUT_ARCH_CPU}/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(update_engine_client PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(update_engine_client
            libbinder
            libbinderwrapper
            libbrillo
            libbrillo-binder
            libbrillo-stream
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            libm
            libutils
    )
endif ()
#################### update_engine_client ####################

#################### libupdate_engine_android ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/libupdate_engine_android/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IUpdateEngine.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/libupdate_engine_android/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IUpdateEngineCallback.cpp
        ${ANDROID_ROOT}/system/update_engine/common/system_state.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/apex_handler_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/binder_service_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/binder_service_stable_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/daemon_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/daemon_state_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/hardware_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/logging_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/network_selector_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/update_attempter_android.cc
        ${ANDROID_ROOT}/system/update_engine/certificate_checker.cc
        ${ANDROID_ROOT}/system/update_engine/download_action.cc
        ${ANDROID_ROOT}/system/update_engine/libcurl_http_fetcher.cc
        ${ANDROID_ROOT}/system/update_engine/metrics_utils.cc
        ${ANDROID_ROOT}/system/update_engine/update_boot_flags_action.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libupdate_engine_android SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_android
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/libupdate_engine_android/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/statslog_ue.h/gen
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_android PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libupdate_engine_android
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            android.hardware.boot-1.2
            apex_aidl_interface-cpp
            avb_headers
            libPlatformProperties
            libandroid_net
            libavb
            libavb_user
            libbase
            libbinder
            libbinderwrapper
            libboot_control_client
            libbootloader_message
            libbrillo
            libbrillo-binder
            libbrillo-stream
            libbrotli
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcom.android.sysprop.apex
            libcow_operation_convert
            libcrypto
            libcurl
            libcutils
            libdl
            libfec
            libfec_rs
            libfs_mgr_binder
            libgkiprops
            libgsi
            libgtest_prod_headers
            libhidlbase
            liblog
            liblp
            liblz4
            liblz4patch
            libm
            libpayload_consumer
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libpuffpatch
            libsnapshot
            libsnapshot_cow
            libssl
            libstatslog_ue
            libstatssocket
            libupdate_engine_boot_control
            libupdate_engine_stable-V1-cpp
            libutils
            libverity_tree
            libxz
            libz
            libziparchive
            libzstd
            lz4diff-protos
            update_metadata-protos
    )
endif ()
#################### libupdate_engine_android ####################

#################### update_engine_sideload ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/aosp/hardware_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/logging_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/sideload_main.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/update_attempter_android.cc
        ${ANDROID_ROOT}/system/update_engine/common/metrics_reporter_stub.cc
        ${ANDROID_ROOT}/system/update_engine/common/network_selector_stub.cc
        ${ANDROID_ROOT}/system/update_engine/common/system_state.cc
        ${ANDROID_ROOT}/system/update_engine/download_action.cc
        ${ANDROID_ROOT}/system/update_engine/metrics_utils.cc
        ${ANDROID_ROOT}/system/update_engine/update_boot_flags_action.cc
        ${ANDROID_ROOT}/system/update_engine/update_status_utils.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_executable(update_engine_sideload ${SOURCE_FILES})

    target_include_directories(update_engine_sideload
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_recovery_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_recovery_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(update_engine_sideload PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(update_engine_sideload
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            android.hardware.boot-1.2
            avb_headers
            libbase
            libboot_control_client
            libbootloader_message
            libbrillo-stream
            libbrillo
            libbrotli
            libbspatch
            libbz
            libc++
            libc++demangle
            libc++fs
            libc
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcow_operation_convert
            libcrypto
            libcutils
            libdl
            libevent
            libfec
            libfec_rs
            libfs_mgr
            libgkiprops
            libgsi
            libgtest_prod_headers
            libhidlbase
            liblog
            liblp
            liblz4
            liblz4patch
            libm
            libmodpb64
            libpayload_consumer
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libpuffpatch
            libsnapshot_cow
            libsnapshot_nobinder
            libupdate_engine_boot_control
            libutils
            libverity_tree
            libxz
            libz
            libziparchive
            libzstd
            lz4diff-protos
            update_metadata-protos
            libc
            libdl
            libm
    )
endif ()
#################### update_engine_sideload ####################

#################### libpayload_extent_ranges ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/payload_generator/extent_ranges.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libpayload_extent_ranges SHARED ${SOURCE_FILES})

    target_include_directories(libpayload_extent_ranges
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libpayload_extent_ranges PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libpayload_extent_ranges
            libbrillo
            libbrillo-stream
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            libm
            update_metadata-protos
    )
endif ()
#################### libpayload_extent_ranges ####################

#################### libupdate_engine_android_proto-full ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/libupdate_engine_android_proto-full/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IUpdateEngine.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/libupdate_engine_android_proto-full/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IUpdateEngineCallback.cpp
        ${ANDROID_ROOT}/system/update_engine/common/system_state.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/apex_handler_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/binder_service_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/binder_service_stable_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/daemon_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/daemon_state_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/hardware_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/logging_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/network_selector_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/update_attempter_android.cc
        ${ANDROID_ROOT}/system/update_engine/certificate_checker.cc
        ${ANDROID_ROOT}/system/update_engine/download_action.cc
        ${ANDROID_ROOT}/system/update_engine/libcurl_http_fetcher.cc
        ${ANDROID_ROOT}/system/update_engine/metrics_utils.cc
        ${ANDROID_ROOT}/system/update_engine/update_boot_flags_action.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libupdate_engine_android_proto-full SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_android_proto-full
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/libupdate_engine_android_proto-full/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos-full/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/statslog_ue.h/gen
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_android_proto-full PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libupdate_engine_android_proto-full
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            android.hardware.boot-1.2
            apex_aidl_interface-cpp
            avb_headers
            libPlatformProperties
            libandroid_net
            libavb
            libavb_user
            libbase
            libbinder
            libbinderwrapper
            libboot_control_client
            libbootloader_message
            libbrillo
            libbrillo-binder
            libbrillo-stream
            libbrotli
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcom.android.sysprop.apex
            libcow_operation_convert
            libcrypto
            libcurl
            libcutils
            libdl
            libfec
            libfec_rs
            libfs_mgr_binder
            libgkiprops
            libgsi
            libgtest_prod_headers
            libhidlbase
            liblog
            liblp
            liblz4
            liblz4patch
            libm
            libpayload_consumer_proto-full
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libpuffpatch
            libsnapshot
            libsnapshot_cow
            libssl
            libstatslog_ue
            libstatssocket
            libupdate_engine_boot_control_proto-full
            libupdate_engine_stable-V1-cpp
            libutils
            libverity_tree
            libxz
            libz
            libziparchive
            libzstd
            lz4diff-protos
            update_metadata-protos-full
    )
endif ()
#################### libupdate_engine_android_proto-full ####################

#################### update_metadata-protos-full ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos-full/android_${OUT_ARCH_CPU}_static/gen/proto/update_engine/update_metadata.pb.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(update_metadata-protos-full SHARED ${SOURCE_FILES})

    target_include_directories(update_metadata-protos-full
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos-full/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(update_metadata-protos-full PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(update_metadata-protos-full
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### update_metadata-protos-full ####################

#################### libcow_operation_convert ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/common/cow_operation_convert.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libcow_operation_convert SHARED ${SOURCE_FILES})

    target_include_directories(libcow_operation_convert
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libcow_operation_convert PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libcow_operation_convert
            libbrillo
            libbrillo-stream
            libbrotli
            libc
            libc++
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest_prod_headers
            libm
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libsnapshot_cow
            libz
            update_metadata-protos
    )
endif ()
#################### libcow_operation_convert ####################

#################### update_engine ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/main.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/metrics_reporter_android.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_executable(update_engine ${SOURCE_FILES})

    target_include_directories(update_engine
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/statslog_ue.h/gen
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/libupdate_engine_android/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/stable/libupdate_engine_stable-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(update_engine PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(update_engine
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            android.hardware.boot-1.2
            apex_aidl_interface-cpp
            avb_headers
            libPlatformProperties
            libandroid_net
            libavb
            libavb_user
            libbase
            libbinder
            libbinderwrapper
            libboot_control_client
            libbootloader_message
            libbrillo
            libbrillo-binder
            libbrillo-stream
            libbrotli
            libbspatch
            libbz
            libc
            libc++
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcom.android.sysprop.apex
            libcow_operation_convert
            libcrypto
            libcurl
            libcutils
            libdl
            libfec
            libfec_rs
            libfs_mgr_binder
            libgflags
            libgkiprops
            libgsi
            libgtest_prod_headers
            libhidlbase
            liblog
            liblp
            liblz4
            liblz4patch
            libm
            libpayload_consumer
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libpuffpatch
            libsnapshot
            libsnapshot_cow
            libssl
            libstatslog_ue
            libstatssocket
            libupdate_engine_android
            libupdate_engine_boot_control
            libupdate_engine_stable-V1-cpp
            libutils
            libverity_tree
            libxz
            libz
            libziparchive
            libzstd
            lz4diff-protos
            update_metadata-protos
    )
endif ()
#################### update_engine ####################

#################### libpayload_consumer ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/aosp/platform_constants_android.cc
        ${ANDROID_ROOT}/system/update_engine/common/action_processor.cc
        ${ANDROID_ROOT}/system/update_engine/common/boot_control_stub.cc
        ${ANDROID_ROOT}/system/update_engine/common/clock.cc
        ${ANDROID_ROOT}/system/update_engine/common/constants.cc
        ${ANDROID_ROOT}/system/update_engine/common/cpu_limiter.cc
        ${ANDROID_ROOT}/system/update_engine/common/dynamic_partition_control_stub.cc
        ${ANDROID_ROOT}/system/update_engine/common/error_code_utils.cc
        ${ANDROID_ROOT}/system/update_engine/common/file_fetcher.cc
        ${ANDROID_ROOT}/system/update_engine/common/hash_calculator.cc
        ${ANDROID_ROOT}/system/update_engine/common/http_common.cc
        ${ANDROID_ROOT}/system/update_engine/common/http_fetcher.cc
        ${ANDROID_ROOT}/system/update_engine/common/hwid_override.cc
        ${ANDROID_ROOT}/system/update_engine/common/multi_range_http_fetcher.cc
        ${ANDROID_ROOT}/system/update_engine/common/prefs.cc
        ${ANDROID_ROOT}/system/update_engine/common/subprocess.cc
        ${ANDROID_ROOT}/system/update_engine/common/terminator.cc
        ${ANDROID_ROOT}/system/update_engine/common/utils.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/bzip_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/cached_file_descriptor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/certificate_parser_android.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/cow_writer_file_descriptor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/delta_performer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/extent_reader.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/file_descriptor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/file_descriptor_utils.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/file_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/filesystem_verifier_action.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/install_operation_executor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/install_plan.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/mount_history.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/payload_constants.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/payload_metadata.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/payload_verifier.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/partition_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/partition_writer_factory_android.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/vabc_partition_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/xor_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/block_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/snapshot_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/postinstall_runner_action.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/verified_source_fd.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/verity_writer_android.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/xz_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/fec_file_descriptor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/partition_update_generator_android.cc
        ${ANDROID_ROOT}/system/update_engine/update_status_utils.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libpayload_consumer SHARED ${SOURCE_FILES})

    target_include_directories(libpayload_consumer
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libpayload_consumer PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libpayload_consumer
            libbase
            libbrillo
            libbrillo-stream
            libbrotli
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcow_operation_convert
            libcrypto
            libdl
            libfec
            libfec_rs
            libgtest_prod_headers
            liblz4
            liblz4patch
            libm
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libpuffpatch
            libsnapshot_cow
            libverity_tree
            libxz
            libz
            libziparchive
            libzstd
            lz4diff-protos
            update_metadata-protos
    )
endif ()
#################### libpayload_consumer ####################

#################### libstatslog_ue ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/statslog_ue.cpp/gen/statslog_ue.cpp
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libstatslog_ue SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_ue
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/statslog_ue.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/statslog_ue.h/gen
    )

    target_link_libraries(libstatslog_ue
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libstatssocket
    )
endif ()
#################### libstatslog_ue ####################

#################### libpayload_extent_utils ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/payload_generator/extent_utils.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libpayload_extent_utils SHARED ${SOURCE_FILES})

    target_include_directories(libpayload_extent_utils
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libpayload_extent_utils PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libpayload_extent_utils
            libbrillo
            libbrillo-stream
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            libm
            update_metadata-protos
    )
endif ()
#################### libpayload_extent_utils ####################

#################### libupdate_engine_boot_control_proto-full ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/aosp/boot_control_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/cleanup_previous_update_action.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/dynamic_partition_control_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/dynamic_partition_utils.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libupdate_engine_boot_control_proto-full SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_boot_control_proto-full
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/statslog_ue.h/gen
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos-full/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_boot_control_proto-full PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libupdate_engine_boot_control_proto-full
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            android.hardware.boot-1.2
            avb_headers
            libbase
            libboot_control_client
            libbootloader_message
            libbrillo
            libbrillo-stream
            libbrotli
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcow_operation_convert
            libcrypto
            libcutils
            libdl
            libfec
            libfec_rs
            libfs_mgr_binder
            libgsi
            libgtest_prod_headers
            libhidlbase
            liblp
            liblz4
            liblz4patch
            libm
            libpayload_consumer_proto-full
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libpuffpatch
            libsnapshot
            libsnapshot_cow
            libstatslog_ue
            libutils
            libverity_tree
            libxz
            libz
            libziparchive
            libzstd
            lz4diff-protos
            update_metadata-protos-full
    )
endif ()
#################### libupdate_engine_boot_control_proto-full ####################

#################### libpayload_consumer_proto-full ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/aosp/platform_constants_android.cc
        ${ANDROID_ROOT}/system/update_engine/common/action_processor.cc
        ${ANDROID_ROOT}/system/update_engine/common/boot_control_stub.cc
        ${ANDROID_ROOT}/system/update_engine/common/clock.cc
        ${ANDROID_ROOT}/system/update_engine/common/constants.cc
        ${ANDROID_ROOT}/system/update_engine/common/cpu_limiter.cc
        ${ANDROID_ROOT}/system/update_engine/common/dynamic_partition_control_stub.cc
        ${ANDROID_ROOT}/system/update_engine/common/error_code_utils.cc
        ${ANDROID_ROOT}/system/update_engine/common/file_fetcher.cc
        ${ANDROID_ROOT}/system/update_engine/common/hash_calculator.cc
        ${ANDROID_ROOT}/system/update_engine/common/http_common.cc
        ${ANDROID_ROOT}/system/update_engine/common/http_fetcher.cc
        ${ANDROID_ROOT}/system/update_engine/common/hwid_override.cc
        ${ANDROID_ROOT}/system/update_engine/common/multi_range_http_fetcher.cc
        ${ANDROID_ROOT}/system/update_engine/common/prefs.cc
        ${ANDROID_ROOT}/system/update_engine/common/subprocess.cc
        ${ANDROID_ROOT}/system/update_engine/common/terminator.cc
        ${ANDROID_ROOT}/system/update_engine/common/utils.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/bzip_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/cached_file_descriptor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/certificate_parser_android.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/cow_writer_file_descriptor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/delta_performer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/extent_reader.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/file_descriptor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/file_descriptor_utils.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/file_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/filesystem_verifier_action.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/install_operation_executor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/install_plan.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/mount_history.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/payload_constants.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/payload_metadata.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/payload_verifier.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/partition_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/partition_writer_factory_android.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/vabc_partition_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/xor_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/block_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/snapshot_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/postinstall_runner_action.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/verified_source_fd.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/verity_writer_android.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/xz_extent_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/fec_file_descriptor.cc
        ${ANDROID_ROOT}/system/update_engine/payload_consumer/partition_update_generator_android.cc
        ${ANDROID_ROOT}/system/update_engine/update_status_utils.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libpayload_consumer_proto-full SHARED ${SOURCE_FILES})

    target_include_directories(libpayload_consumer_proto-full
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos-full/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libpayload_consumer_proto-full PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libpayload_consumer_proto-full
            libbase
            libbrillo
            libbrillo-stream
            libbrotli
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcow_operation_convert
            libcrypto
            libdl
            libfec
            libfec_rs
            libgtest_prod_headers
            liblz4
            liblz4patch
            libm
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libpuffpatch
            libsnapshot_cow
            libverity_tree
            libxz
            libz
            libziparchive
            libzstd
            lz4diff-protos
            update_metadata-protos-full
    )
endif ()
#################### libpayload_consumer_proto-full ####################

#################### lz4diff-protos ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto/lz4diff/lz4diff.pb.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(lz4diff-protos SHARED ${SOURCE_FILES})

    target_include_directories(lz4diff-protos
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(lz4diff-protos PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(lz4diff-protos
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### lz4diff-protos ####################

#################### liblz4patch ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/lz4diff/lz4patch.cc
        ${ANDROID_ROOT}/system/update_engine/lz4diff/lz4diff_compress.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(liblz4patch SHARED ${SOURCE_FILES})

    target_include_directories(liblz4patch
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(liblz4patch PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(liblz4patch
            libbrillo
            libbrillo-stream
            libbspatch
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest_prod_headers
            liblz4
            libm
            libpuffpatch
            libssl
            lz4diff-protos
            update_metadata-protos
    )
endif ()
#################### liblz4patch ####################

#################### liblz4diff ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/lz4diff/lz4diff.cc
        ${ANDROID_ROOT}/system/update_engine/lz4diff/lz4diff_compress.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(liblz4diff SHARED ${SOURCE_FILES})

    target_include_directories(liblz4diff
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(liblz4diff PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(liblz4diff
            libbrillo
            libbrillo-stream
            libbsdiff
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest_prod_headers
            liblz4
            libm
            libpuffdiff
            libssl
            lz4diff-protos
            update_metadata-protos
    )
endif ()
#################### liblz4diff ####################

#################### libcow_size_estimator ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/payload_generator/cow_size_estimator.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libcow_size_estimator SHARED ${SOURCE_FILES})

    target_include_directories(libcow_size_estimator
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libcow_size_estimator PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libcow_size_estimator
            libbase
            libbrillo
            libbrillo-stream
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcow_operation_convert
            libdl
            libgtest_prod_headers
            libm
            libprotobuf-cpp-lite
            libsnapshot_cow
            update_metadata-protos
    )
endif ()
#################### libcow_size_estimator ####################

#################### libpayload_generator ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/common/system_state.cc
        ${ANDROID_ROOT}/system/update_engine/download_action.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/ab_generator.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/annotated_operation.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/blob_file_writer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/block_mapping.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/boot_img_filesystem.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/bzip.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/deflate_utils.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/delta_diff_generator.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/delta_diff_utils.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/ext2_filesystem.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/erofs_filesystem.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/extent_ranges.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/full_update_generator.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/mapfile_filesystem.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/merge_sequence_generator.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/payload_file.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/payload_generation_config_android.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/payload_generation_config.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/payload_properties.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/payload_signer.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/raw_filesystem.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/squashfs_filesystem.cc
        ${ANDROID_ROOT}/system/update_engine/payload_generator/xz_android.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libpayload_generator SHARED ${SOURCE_FILES})

    target_include_directories(libpayload_generator
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libpayload_generator PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libpayload_generator
            bootimg_headers
            libavb
            libbase
            libbrillo
            libbrillo-stream
            libbrotli
            libbsdiff
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcow_operation_convert
            libcow_size_estimator
            libcrypto
            libdivsufsort
            libdivsufsort64
            libdl
            liberofs
            libext2fs
            libfec
            libfec_rs
            libgtest_prod_headers
            liblz4
            liblz4diff
            liblz4patch
            liblzma
            libm
            libpayload_consumer
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libpuffdiff
            libpuffpatch
            libselinux
            libsnapshot_cow
            libverity_tree
            libxz
            libz
            libziparchive
            libzstd
            libzucchini
            lz4diff-protos
            update_metadata-protos
    )
endif ()
#################### libpayload_generator ####################

#################### libupdate_engine_boot_control ####################
set(FILES
        ${ANDROID_ROOT}/system/update_engine/aosp/boot_control_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/cleanup_previous_update_action.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/dynamic_partition_control_android.cc
        ${ANDROID_ROOT}/system/update_engine/aosp/dynamic_partition_utils.cc
)
set(SOURCE_FILES "")
foreach (file ${FILES})
    if (EXISTS ${file})
        list(APPEND SOURCE_FILES ${file})
    endif ()
endforeach ()

if (NOT SOURCE_FILES)
    message(WARNING "${PROJECT_NAME}, source files not exist")
else ()
    add_library(libupdate_engine_boot_control SHARED ${SOURCE_FILES})

    target_include_directories(libupdate_engine_boot_control
            PUBLIC
            ${ANDROID_ROOT}/system/update_engine/client_library/include
            ${ANDROID_ROOT}/system/update_engine
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/lz4diff-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/update_metadata-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/update_engine/statslog_ue.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libupdate_engine_boot_control PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libupdate_engine_boot_control
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            android.hardware.boot-1.2
            avb_headers
            libbase
            libboot_control_client
            libbootloader_message
            libbrillo
            libbrillo-stream
            libbrotli
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcow_operation_convert
            libcrypto
            libcutils
            libdl
            libfec
            libfec_rs
            libfs_mgr_binder
            libgsi
            libgtest_prod_headers
            libhidlbase
            liblp
            liblz4
            liblz4patch
            libm
            libpayload_consumer
            libpayload_extent_ranges
            libpayload_extent_utils
            libprotobuf-cpp-lite
            libpuffpatch
            libsnapshot
            libsnapshot_cow
            libstatslog_ue
            libutils
            libverity_tree
            libxz
            libz
            libziparchive
            libzstd
            lz4diff-protos
            update_metadata-protos
    )
endif ()
#################### libupdate_engine_boot_control ####################

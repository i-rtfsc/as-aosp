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

project(init_first_stage)

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

#################### init_second_stage ####################
set(FILES
        ${ANDROID_ROOT}/system/core/init/main.cpp
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
    add_executable(init_second_stage ${SOURCE_FILES})

    target_include_directories(init_second_stage
            PUBLIC
            ${ANDROID_ROOT}/system/core/init
            ${OUT_INTERMEDIATES_ROOT}/system/core/init/sysprop/libcom.android.sysprop.init/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
            ${OUT_INTERMEDIATES_ROOT}/system/core/init/libinit/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/init
            ${OUT_INTERMEDIATES_ROOT}/system/core/init/libinit/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(init_second_stage PRIVATE
            -DINSTALL_DEBUG_POLICY_TO_SYSTEM_EXT=0
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(init_second_stage
            bionic_libc_platform_headers
            lib_apex_manifest_proto_lite
            libavb
            libbase
            libbootloader_message
            libc
            libc++
            libc++demangle
            libc++fs
            libcgrouprc_format
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libext4_utils
            libfs_mgr
            libfsverity_init
            libgsi
            libhidl-gen-utils
            libinit
            libkeyutils
            liblmkd_utils
            liblog
            liblogwrap
            liblp
            liblz4
            libm
            libmini_keyctl_static
            libmodprobe
            libprocessgroup
            libprocessgroup_setup
            libprocinfo
            libpropertyinfoparser
            libpropertyinfoserializer
            libprotobuf-cpp-lite
            libselinux
            libsigningutils
            libsnapshot_cow
            libsnapshot_init
            libunwindstack
            libutils
            libxml2
            libziparchive
            update_metadata-protos
    )
endif ()
#################### init_second_stage ####################

#################### init_first_stage ####################
set(FILES
        ${ANDROID_ROOT}/system/core/init/block_dev_initializer.cpp
        ${ANDROID_ROOT}/system/core/init/devices.cpp
        ${ANDROID_ROOT}/system/core/init/first_stage_console.cpp
        ${ANDROID_ROOT}/system/core/init/first_stage_init.cpp
        ${ANDROID_ROOT}/system/core/init/first_stage_main.cpp
        ${ANDROID_ROOT}/system/core/init/first_stage_mount.cpp
        ${ANDROID_ROOT}/system/core/init/reboot_utils.cpp
        ${ANDROID_ROOT}/system/core/init/selabel.cpp
        ${ANDROID_ROOT}/system/core/init/service_utils.cpp
        ${ANDROID_ROOT}/system/core/init/snapuserd_transition.cpp
        ${ANDROID_ROOT}/system/core/init/switch_root.cpp
        ${ANDROID_ROOT}/system/core/init/uevent_listener.cpp
        ${ANDROID_ROOT}/system/core/init/util.cpp
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
    add_executable(init_first_stage ${SOURCE_FILES})

    target_include_directories(init_first_stage
            PUBLIC
            ${ANDROID_ROOT}/system/core/init
    )

    # Add c/c++ flags
    target_compile_options(init_first_stage PRIVATE
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(init_first_stage
            libavb
            libbase
            libc++_static
            libc++demangle
            libc++fs
            libc
            libcap
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libcrypto_utils
            libcutils
            libext2_uuid
            libfec
            libfec_rs
            libfs_avb
            libfs_mgr
            libgsi
            liblog
            liblp
            liblz4
            liblzma
            libm
            libmodprobe
            libprocinfo
            libprotobuf-cpp-lite
            libselinux
            libsnapshot_cow
            libsnapshot_init
            libsquashfs_utils
            libunwind
            libunwindstack_no_dex
            update_metadata-protos
    )
endif ()
#################### init_first_stage ####################

#################### libinit ####################
set(FILES
        ${ANDROID_ROOT}/system/core/init/action.cpp
        ${ANDROID_ROOT}/system/core/init/action_manager.cpp
        ${ANDROID_ROOT}/system/core/init/action_parser.cpp
        ${ANDROID_ROOT}/system/core/init/capabilities.cpp
        ${ANDROID_ROOT}/system/core/init/epoll.cpp
        ${ANDROID_ROOT}/system/core/init/import_parser.cpp
        ${ANDROID_ROOT}/system/core/init/interface_utils.cpp
        ${ANDROID_ROOT}/system/core/init/interprocess_fifo.cpp
        ${ANDROID_ROOT}/system/core/init/keychords.cpp
        ${ANDROID_ROOT}/system/core/init/parser.cpp
        ${ANDROID_ROOT}/system/core/init/property_type.cpp
        ${ANDROID_ROOT}/system/core/init/rlimit_parser.cpp
        ${ANDROID_ROOT}/system/core/init/service.cpp
        ${ANDROID_ROOT}/system/core/init/service_list.cpp
        ${ANDROID_ROOT}/system/core/init/service_parser.cpp
        ${ANDROID_ROOT}/system/core/init/service_utils.cpp
        ${ANDROID_ROOT}/system/core/init/subcontext.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/core/init/libinit/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/init/subcontext.pb.cc
        ${ANDROID_ROOT}/system/core/init/tokenizer.cpp
        ${ANDROID_ROOT}/system/core/init/util.cpp
        ${ANDROID_ROOT}/system/core/init/apex_init_util.cpp
        ${ANDROID_ROOT}/system/core/init/block_dev_initializer.cpp
        ${ANDROID_ROOT}/system/core/init/bootchart.cpp
        ${ANDROID_ROOT}/system/core/init/builtins.cpp
        ${ANDROID_ROOT}/system/core/init/devices.cpp
        ${ANDROID_ROOT}/system/core/init/firmware_handler.cpp
        ${ANDROID_ROOT}/system/core/init/first_stage_console.cpp
        ${ANDROID_ROOT}/system/core/init/first_stage_init.cpp
        ${ANDROID_ROOT}/system/core/init/first_stage_mount.cpp
        ${ANDROID_ROOT}/system/core/init/fscrypt_init_extensions.cpp
        ${ANDROID_ROOT}/system/core/init/init.cpp
        ${ANDROID_ROOT}/system/core/init/lmkd_service.cpp
        ${ANDROID_ROOT}/system/core/init/modalias_handler.cpp
        ${ANDROID_ROOT}/system/core/init/mount_handler.cpp
        ${ANDROID_ROOT}/system/core/init/mount_namespace.cpp
        ${ANDROID_ROOT}/system/core/init/persistent_properties.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/core/init/libinit/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/init/persistent_properties.pb.cc
        ${ANDROID_ROOT}/system/core/init/property_service.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/core/init/libinit/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/init/property_service.pb.cc
        ${ANDROID_ROOT}/system/core/init/reboot.cpp
        ${ANDROID_ROOT}/system/core/init/reboot_utils.cpp
        ${ANDROID_ROOT}/system/core/init/security.cpp
        ${ANDROID_ROOT}/system/core/init/selabel.cpp
        ${ANDROID_ROOT}/system/core/init/selinux.cpp
        ${ANDROID_ROOT}/system/core/init/sigchld_handler.cpp
        ${ANDROID_ROOT}/system/core/init/snapuserd_transition.cpp
        ${ANDROID_ROOT}/system/core/init/switch_root.cpp
        ${ANDROID_ROOT}/system/core/init/uevent_listener.cpp
        ${ANDROID_ROOT}/system/core/init/ueventd.cpp
        ${ANDROID_ROOT}/system/core/init/ueventd_parser.cpp
        ${ANDROID_ROOT}/system/core/init/private_opt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex-info-list/gen/cpp/com_android_apex.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex-info-list/gen/cpp/com_android_apex_enums.cpp
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
    add_library(libinit SHARED ${SOURCE_FILES})

    target_include_directories(libinit
            PUBLIC
            ${ANDROID_ROOT}/system/core/init
            ${OUT_INTERMEDIATES_ROOT}/system/core/init/libinit/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/init
            ${OUT_INTERMEDIATES_ROOT}/system/core/init/libinit/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/core/init/sysprop/libcom.android.sysprop.init/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libinit PRIVATE
            -DINSTALL_DEBUG_POLICY_TO_SYSTEM_EXT=0
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libinit
            bionic_libc_platform_headers
            bootimg_headers
            lib_apex_manifest_proto_lite
            libavb
            libbase
            libbootloader_message
            libc
            libc++_static
            libc++demangle
            libc++fs
            libcap
            libcgrouprc_format
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcom.android.sysprop.apex
            libcom.android.sysprop.init
            libcrypto
            libcutils
            libdl
            libext4_utils
            libfs_mgr
            libfsverity_init
            libgsi
            libhidl-gen-utils
            libkeyutils
            liblmkd_utils
            liblog
            liblogwrap
            liblp
            liblz4
            libm
            libmini_keyctl_static
            libmodprobe
            libprocessgroup
            libprocessgroup_setup
            libprocinfo
            libpropertyinfoparser
            libpropertyinfoserializer
            libprotobuf-cpp-lite
            libselinux
            libsigningutils
            libsnapshot_cow
            libsnapshot_init
            libunwindstack
            libutils
            libxml2
            libziparchive
            update_metadata-protos
    )
endif ()
#################### libinit ####################

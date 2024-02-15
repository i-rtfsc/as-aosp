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

project(libcutils_sockets)

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

#################### libcutils_sockets ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libcutils/sockets.cpp
        ${ANDROID_ROOT}/system/core/libcutils/android_get_control_file.cpp
        ${ANDROID_ROOT}/system/core/libcutils/socket_inaddr_any_server_unix.cpp
        ${ANDROID_ROOT}/system/core/libcutils/socket_local_client_unix.cpp
        ${ANDROID_ROOT}/system/core/libcutils/socket_local_server_unix.cpp
        ${ANDROID_ROOT}/system/core/libcutils/socket_network_client_unix.cpp
        ${ANDROID_ROOT}/system/core/libcutils/sockets_unix.cpp
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
    add_library(libcutils_sockets SHARED ${SOURCE_FILES})

    target_include_directories(libcutils_sockets
            PUBLIC
            ${ANDROID_ROOT}/system/core/libcutils/include
            ${ANDROID_ROOT}/system/core/libcutils
    )

    # Add c/c++ flags
    target_compile_options(libcutils_sockets PRIVATE
            -Werror
    )

    target_link_libraries(libcutils_sockets
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libwinpthread
    )
endif ()
#################### libcutils_sockets ####################

#################### libcutils ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libcutils/config_utils.cpp
        ${ANDROID_ROOT}/system/core/libcutils/iosched_policy.cpp
        ${ANDROID_ROOT}/system/core/libcutils/load_file.cpp
        ${ANDROID_ROOT}/system/core/libcutils/native_handle.cpp
        ${ANDROID_ROOT}/system/core/libcutils/properties.cpp
        ${ANDROID_ROOT}/system/core/libcutils/record_stream.cpp
        ${ANDROID_ROOT}/system/core/libcutils/strlcpy.c
        ${ANDROID_ROOT}/system/core/libcutils/threads.cpp
        ${ANDROID_ROOT}/system/core/libcutils/fs.cpp
        ${ANDROID_ROOT}/system/core/libcutils/hashmap.cpp
        ${ANDROID_ROOT}/system/core/libcutils/multiuser.cpp
        ${ANDROID_ROOT}/system/core/libcutils/str_parms.cpp
        ${ANDROID_ROOT}/system/core/libcutils/android_reboot.cpp
        ${ANDROID_ROOT}/system/core/libcutils/ashmem-dev.cpp
        ${ANDROID_ROOT}/system/core/libcutils/canned_fs_config.cpp
        ${ANDROID_ROOT}/system/core/libcutils/fs_config.cpp
        ${ANDROID_ROOT}/system/core/libcutils/klog.cpp
        ${ANDROID_ROOT}/system/core/libcutils/partition_utils.cpp
        ${ANDROID_ROOT}/system/core/libcutils/qtaguid.cpp
        ${ANDROID_ROOT}/system/core/libcutils/trace-dev.cpp
        ${ANDROID_ROOT}/system/core/libcutils/uevent.cpp
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
    add_library(libcutils SHARED ${SOURCE_FILES})

    target_include_directories(libcutils
            PUBLIC
            ${ANDROID_ROOT}/system/core/libcutils/include
            ${ANDROID_ROOT}/system/core/libcutils
    )

    # Add c/c++ flags
    target_compile_options(libcutils PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(libcutils
            libasync_safe
            libbase
            libbase_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils_headers
            libcutils_sockets
            libdl
            liblog
            libm
            libprocessgroup_headers
            libwinpthread
    )
endif ()
#################### libcutils ####################

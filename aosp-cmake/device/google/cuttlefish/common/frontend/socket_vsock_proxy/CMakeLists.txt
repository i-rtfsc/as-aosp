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

project(socket_vsock_proxy)

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

#################### socket_vsock_proxy ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/common/frontend/socket_vsock_proxy/client.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/frontend/socket_vsock_proxy/server.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/common/frontend/socket_vsock_proxy/socket_vsock_proxy.cpp
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
    add_executable(socket_vsock_proxy ${SOURCE_FILES})

    target_include_directories(socket_vsock_proxy
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/common/frontend/socket_vsock_proxy
    )

    # Add c/c++ flags
    target_compile_options(socket_vsock_proxy PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(socket_vsock_proxy
            cuttlefish_common_headers
            cuttlefish_common_headers
            libbase
            libbase
            libc++
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcuttlefish_fs
            libcuttlefish_fs
            libcuttlefish_host_config
            libcuttlefish_kernel_log_monitor_utils
            libcuttlefish_kernel_log_monitor_utils
            libcuttlefish_utils
            libcuttlefish_utils
            libdl
            libext2_blkid
            libext2_blkid
            libgflags
            libgflags
            libjsoncpp
            libjsoncpp
            liblog
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### socket_vsock_proxy ####################

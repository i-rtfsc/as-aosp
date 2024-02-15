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

project(iperf3)

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

#################### iperf3 ####################
set(FILES
        ${ANDROID_ROOT}/external/iperf3/src/cjson.c
        ${ANDROID_ROOT}/external/iperf3/src/dscp.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_api.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_auth.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_client_api.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_error.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_locale.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_sctp.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_server_api.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_tcp.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_time.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_udp.c
        ${ANDROID_ROOT}/external/iperf3/src/iperf_util.c
        ${ANDROID_ROOT}/external/iperf3/src/main.c
        ${ANDROID_ROOT}/external/iperf3/src/net.c
        ${ANDROID_ROOT}/external/iperf3/src/tcp_info.c
        ${ANDROID_ROOT}/external/iperf3/src/timer.c
        ${ANDROID_ROOT}/external/iperf3/src/units.c
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
    add_executable(iperf3 ${SOURCE_FILES})

    target_include_directories(iperf3
            PUBLIC
            ${ANDROID_ROOT}/external/iperf3
    )

    # Add c/c++ flags
    target_compile_options(iperf3 PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(iperf3
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### iperf3 ####################

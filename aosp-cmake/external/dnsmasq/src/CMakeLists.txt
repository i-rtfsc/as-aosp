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

project(dnsmasq)

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

#################### dnsmasq ####################
set(FILES
        ${ANDROID_ROOT}/external/dnsmasq/src/cache.c
        ${ANDROID_ROOT}/external/dnsmasq/src/dhcp.c
        ${ANDROID_ROOT}/external/dnsmasq/src/dnsmasq.c
        ${ANDROID_ROOT}/external/dnsmasq/src/forward.c
        ${ANDROID_ROOT}/external/dnsmasq/src/helper.c
        ${ANDROID_ROOT}/external/dnsmasq/src/lease.c
        ${ANDROID_ROOT}/external/dnsmasq/src/log.c
        ${ANDROID_ROOT}/external/dnsmasq/src/netlink.c
        ${ANDROID_ROOT}/external/dnsmasq/src/network.c
        ${ANDROID_ROOT}/external/dnsmasq/src/option.c
        ${ANDROID_ROOT}/external/dnsmasq/src/rfc1035.c
        ${ANDROID_ROOT}/external/dnsmasq/src/rfc2131.c
        ${ANDROID_ROOT}/external/dnsmasq/src/util.c
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
    add_executable(dnsmasq ${SOURCE_FILES})

    target_include_directories(dnsmasq
            PUBLIC
            ${ANDROID_ROOT}/external/dnsmasq/src
    )

    # Add c/c++ flags
    target_compile_options(dnsmasq PRIVATE
            -Wall
            -Wno-unused-variable
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(dnsmasq
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            liblog
    )
endif ()
#################### dnsmasq ####################

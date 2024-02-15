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

project(libpcap)

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

#################### libpcap ####################
set(FILES
        ${ANDROID_ROOT}/external/libpcap/bpf_dump.c
        ${ANDROID_ROOT}/external/libpcap/bpf_filter.c
        ${ANDROID_ROOT}/external/libpcap/bpf_image.c
        ${ANDROID_ROOT}/external/libpcap/etherent.c
        ${ANDROID_ROOT}/external/libpcap/fad-getad.c
        ${ANDROID_ROOT}/external/libpcap/fmtutils.c
        ${ANDROID_ROOT}/external/libpcap/gencode.c
        ${ANDROID_ROOT}/external/libpcap/grammar.c
        ${ANDROID_ROOT}/external/libpcap/nametoaddr.c
        ${ANDROID_ROOT}/external/libpcap/optimize.c
        ${ANDROID_ROOT}/external/libpcap/pcap-common.c
        ${ANDROID_ROOT}/external/libpcap/pcap-linux.c
        ${ANDROID_ROOT}/external/libpcap/pcap-netfilter-linux-android.c
        ${ANDROID_ROOT}/external/libpcap/pcap-usb-linux-common.c
        ${ANDROID_ROOT}/external/libpcap/pcap-usb-linux.c
        ${ANDROID_ROOT}/external/libpcap/pcap-util.c
        ${ANDROID_ROOT}/external/libpcap/pcap.c
        ${ANDROID_ROOT}/external/libpcap/savefile.c
        ${ANDROID_ROOT}/external/libpcap/scanner.c
        ${ANDROID_ROOT}/external/libpcap/sf-pcap.c
        ${ANDROID_ROOT}/external/libpcap/sf-pcapng.c
        ${ANDROID_ROOT}/external/libpcap/missing/strlcpy.c
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
    add_library(libpcap SHARED ${SOURCE_FILES})

    target_include_directories(libpcap
            PUBLIC
            ${ANDROID_ROOT}/external/libpcap
    )

    # Add c/c++ flags
    target_compile_options(libpcap PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(libpcap
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libpcap ####################

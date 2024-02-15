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

project(libcups)

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

#################### libcups ####################
set(FILES
        ${ANDROID_ROOT}/external/libcups/cups/array.c
        ${ANDROID_ROOT}/external/libcups/cups/auth.c
        ${ANDROID_ROOT}/external/libcups/cups/backchannel.c
        ${ANDROID_ROOT}/external/libcups/cups/debug.c
        ${ANDROID_ROOT}/external/libcups/cups/dest.c
        ${ANDROID_ROOT}/external/libcups/cups/dest-job.c
        ${ANDROID_ROOT}/external/libcups/cups/dest-localization.c
        ${ANDROID_ROOT}/external/libcups/cups/dest-options.c
        ${ANDROID_ROOT}/external/libcups/cups/dir.c
        ${ANDROID_ROOT}/external/libcups/cups/encode.c
        ${ANDROID_ROOT}/external/libcups/cups/file.c
        ${ANDROID_ROOT}/external/libcups/cups/getdevices.c
        ${ANDROID_ROOT}/external/libcups/cups/getputfile.c
        ${ANDROID_ROOT}/external/libcups/cups/globals.c
        ${ANDROID_ROOT}/external/libcups/cups/hash.c
        ${ANDROID_ROOT}/external/libcups/cups/http.c
        ${ANDROID_ROOT}/external/libcups/cups/http-addr.c
        ${ANDROID_ROOT}/external/libcups/cups/http-addrlist.c
        ${ANDROID_ROOT}/external/libcups/cups/http-support.c
        ${ANDROID_ROOT}/external/libcups/cups/ipp.c
        ${ANDROID_ROOT}/external/libcups/cups/ipp-support.c
        ${ANDROID_ROOT}/external/libcups/cups/langprintf.c
        ${ANDROID_ROOT}/external/libcups/cups/language.c
        ${ANDROID_ROOT}/external/libcups/cups/md5.c
        ${ANDROID_ROOT}/external/libcups/cups/md5passwd.c
        ${ANDROID_ROOT}/external/libcups/cups/notify.c
        ${ANDROID_ROOT}/external/libcups/cups/options.c
        ${ANDROID_ROOT}/external/libcups/cups/pwg-media.c
        ${ANDROID_ROOT}/external/libcups/cups/raster-error.c
        ${ANDROID_ROOT}/external/libcups/cups/raster-stream.c
        ${ANDROID_ROOT}/external/libcups/cups/raster-stubs.c
        ${ANDROID_ROOT}/external/libcups/cups/request.c
        ${ANDROID_ROOT}/external/libcups/cups/sidechannel.c
        ${ANDROID_ROOT}/external/libcups/cups/snmp.c
        ${ANDROID_ROOT}/external/libcups/cups/string.c
        ${ANDROID_ROOT}/external/libcups/cups/tempfile.c
        ${ANDROID_ROOT}/external/libcups/cups/thread.c
        ${ANDROID_ROOT}/external/libcups/cups/tls-boringssl.c
        ${ANDROID_ROOT}/external/libcups/cups/transcode.c
        ${ANDROID_ROOT}/external/libcups/cups/usersys.c
        ${ANDROID_ROOT}/external/libcups/cups/util.c
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
    add_library(libcups SHARED ${SOURCE_FILES})

    target_include_directories(libcups
            PUBLIC
            ${ANDROID_ROOT}/external/libcups/cups
            ${ANDROID_ROOT}/external/libcups/filter
            ${ANDROID_ROOT}/external/libcups
    )

    # Add c/c++ flags
    target_compile_options(libcups PRIVATE
            -Wextra
            -Wall
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-variable
            -Wno-unused-value
            -Werror
    )

    target_link_libraries(libcups
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libssl
            libz
    )
endif ()
#################### libcups ####################

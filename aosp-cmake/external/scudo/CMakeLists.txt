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

project(scudo_get_error_info_fuzzer)

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

#################### scudo_get_error_info_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/scudo/standalone/fuzz/get_error_info_fuzzer.cpp
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
    add_executable(scudo_get_error_info_fuzzer ${SOURCE_FILES})

    target_include_directories(scudo_get_error_info_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/scudo/standalone
            ${ANDROID_ROOT}/external/scudo/standalone/include
            ${ANDROID_ROOT}/external/scudo
    )

    # Add c/c++ flags
    target_compile_options(scudo_get_error_info_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(scudo_get_error_info_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libscudo
    )
endif ()
#################### scudo_get_error_info_fuzzer ####################

#################### libscudo ####################
set(FILES
        ${ANDROID_ROOT}/external/scudo/standalone/checksum.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/common.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/flags.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/flags_parser.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/linux.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/mem_map.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/release.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/report.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/rss_limit_checker.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/string_utils.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/timing.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/wrappers_c_bionic.cpp
        ${ANDROID_ROOT}/external/scudo/standalone/crc32_hw.cpp
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
    add_library(libscudo SHARED ${SOURCE_FILES})

    target_include_directories(libscudo
            PUBLIC
            ${ANDROID_ROOT}/external/scudo/standalone/include
            ${ANDROID_ROOT}/external/scudo
    )

    # Add c/c++ flags
    target_compile_options(libscudo PRIVATE
            -Wall
            -Wextra
            -Wno-unused-result
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=thread-safety
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libscudo
            bionic_libc_platform_headers
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libscudo ####################

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

project(stack_trace_compressor_fuzzer)

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

#################### options_parser_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/gwp_asan/tools/options_parser_fuzzer.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/optional/options_parser.cpp
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
    add_executable(options_parser_fuzzer ${SOURCE_FILES})

    target_include_directories(options_parser_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/gwp_asan
    )

    # Add c/c++ flags
    target_compile_options(options_parser_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(options_parser_fuzzer
            gwp_asan
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
    )
endif ()
#################### options_parser_fuzzer ####################

#################### stack_trace_compressor_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/gwp_asan/tools/stack_trace_compressor_fuzzer.cpp
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
    add_executable(stack_trace_compressor_fuzzer ${SOURCE_FILES})

    target_include_directories(stack_trace_compressor_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/gwp_asan
    )

    # Add c/c++ flags
    target_compile_options(stack_trace_compressor_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(stack_trace_compressor_fuzzer
            gwp_asan
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
    )
endif ()
#################### stack_trace_compressor_fuzzer ####################

#################### gwp_asan ####################
set(FILES
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/common.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/guarded_pool_allocator.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/platform_specific/common_posix.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/platform_specific/guarded_pool_allocator_posix.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/platform_specific/mutex_posix.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/platform_specific/utilities_posix.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/stack_trace_compressor.cpp
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
    add_library(gwp_asan SHARED ${SOURCE_FILES})

    target_include_directories(gwp_asan
            PUBLIC
            ${ANDROID_ROOT}/external/gwp_asan
    )

    # Add c/c++ flags
    target_compile_options(gwp_asan PRIVATE
            -Werror
    )

    target_link_libraries(gwp_asan
            gwp_asan_headers
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### gwp_asan ####################

#################### gwp_asan_crash_handler ####################
set(FILES
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/common.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/crash_handler.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/platform_specific/common_posix.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/platform_specific/utilities_posix.cpp
        ${ANDROID_ROOT}/external/gwp_asan/gwp_asan/stack_trace_compressor.cpp
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
    add_library(gwp_asan_crash_handler SHARED ${SOURCE_FILES})

    target_include_directories(gwp_asan_crash_handler
            PUBLIC
            ${ANDROID_ROOT}/external/gwp_asan
    )

    # Add c/c++ flags
    target_compile_options(gwp_asan_crash_handler PRIVATE
            -Werror
    )

    target_link_libraries(gwp_asan_crash_handler
            gwp_asan_headers
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### gwp_asan_crash_handler ####################

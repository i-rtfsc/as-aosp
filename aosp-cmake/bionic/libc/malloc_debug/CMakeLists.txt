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

project(libc_malloc_debug_backtrace)

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

#################### libc_malloc_debug_backtrace ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/backtrace.cpp
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/MapData.cpp
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
    add_library(libc_malloc_debug_backtrace SHARED ${SOURCE_FILES})

    target_include_directories(libc_malloc_debug_backtrace
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/malloc_debug
    )

    # Add c/c++ flags
    target_compile_options(libc_malloc_debug_backtrace PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libc_malloc_debug_backtrace
            libasync_safe
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libc_malloc_debug_backtrace ####################

#################### libc_malloc_debug ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/Config.cpp
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/DebugData.cpp
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/debug_disable.cpp
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/GuardData.cpp
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/malloc_debug.cpp
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/PointerData.cpp
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/RecordData.cpp
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/Unreachable.cpp
        ${ANDROID_ROOT}/bionic/libc/malloc_debug/UnwindBacktrace.cpp
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
    add_library(libc_malloc_debug SHARED ${SOURCE_FILES})

    target_include_directories(libc_malloc_debug
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/malloc_debug
    )

    # Add c/c++ flags
    target_compile_options(libc_malloc_debug PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libc_malloc_debug
            libasync_safe
            libbase
            libc
            libc++_static
            libc++demangle
            libc_malloc_debug_backtrace
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libmemunreachable
            libunwindstack
    )
endif ()
#################### libc_malloc_debug ####################

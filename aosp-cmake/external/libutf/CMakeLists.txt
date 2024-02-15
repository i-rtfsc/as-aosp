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

project(libutf)

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

#################### libutf ####################
set(FILES
        ${ANDROID_ROOT}/external/libutf/rune.c
        ${ANDROID_ROOT}/external/libutf/runestrcat.c
        ${ANDROID_ROOT}/external/libutf/runestrchr.c
        ${ANDROID_ROOT}/external/libutf/runestrcmp.c
        ${ANDROID_ROOT}/external/libutf/runestrcpy.c
        ${ANDROID_ROOT}/external/libutf/runestrdup.c
        ${ANDROID_ROOT}/external/libutf/runestrlen.c
        ${ANDROID_ROOT}/external/libutf/runestrecpy.c
        ${ANDROID_ROOT}/external/libutf/runestrncat.c
        ${ANDROID_ROOT}/external/libutf/runestrncmp.c
        ${ANDROID_ROOT}/external/libutf/runestrncpy.c
        ${ANDROID_ROOT}/external/libutf/runestrrchr.c
        ${ANDROID_ROOT}/external/libutf/runestrstr.c
        ${ANDROID_ROOT}/external/libutf/runetype.c
        ${ANDROID_ROOT}/external/libutf/utfecpy.c
        ${ANDROID_ROOT}/external/libutf/utflen.c
        ${ANDROID_ROOT}/external/libutf/utfnlen.c
        ${ANDROID_ROOT}/external/libutf/utfrrune.c
        ${ANDROID_ROOT}/external/libutf/utfrune.c
        ${ANDROID_ROOT}/external/libutf/utfutf.c
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
    add_library(libutf SHARED ${SOURCE_FILES})

    target_include_directories(libutf
            PUBLIC
            ${ANDROID_ROOT}/external/libutf
    )

    # Add c/c++ flags
    target_compile_options(libutf PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libutf
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libutf ####################

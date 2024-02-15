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

project(libxkbcommon)

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

#################### libxkbcommon ####################
set(FILES
        ${ANDROID_ROOT}/external/libxkbcommon/config/libxkbcommon.so.0.0.0.p/parser.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/compose/parser.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/compose/paths.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/compose/state.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/compose/table.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/action.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/ast-build.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/compat.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/expr.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/include.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/keycodes.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/keymap.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/keymap-dump.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/keywords.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/rules.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/scanner.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/symbols.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/types.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/vmod.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/xkbcomp/xkbcomp.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/atom.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/context.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/context-priv.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/keymap.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/keymap-priv.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/keysym.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/keysym-utf.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/registry.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/state.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/text.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/utf8.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/util-list.c
        ${ANDROID_ROOT}/external/libxkbcommon/src/utils.c
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
    add_library(libxkbcommon SHARED ${SOURCE_FILES})

    target_include_directories(libxkbcommon
            PUBLIC
            ${ANDROID_ROOT}/external/libxkbcommon/include
            ${ANDROID_ROOT}/external/libxkbcommon/src
            ${ANDROID_ROOT}/external/libxkbcommon/config
            ${ANDROID_ROOT}/external/libxkbcommon/config/libxkbcommon.so.0.0.0.p
            ${ANDROID_ROOT}/external/libxkbcommon
    )

    # Add c/c++ flags
    target_compile_options(libxkbcommon PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libxkbcommon
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libxml2
    )
endif ()
#################### libxkbcommon ####################

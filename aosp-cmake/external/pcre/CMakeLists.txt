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

project(libpcre2)

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

#################### libpcre2 ####################
set(FILES
        ${ANDROID_ROOT}/external/pcre/src/pcre2_auto_possess.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_compile.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_config.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_context.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_convert.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_dfa_match.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_error.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_extuni.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_find_bracket.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_maketables.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_match.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_match_data.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_jit_compile.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_newline.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_ord2utf.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_pattern_info.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_script_run.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_serialize.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_string_utils.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_study.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_substitute.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_substring.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_tables.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_ucd.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_valid_utf.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_xclass.c
        ${ANDROID_ROOT}/external/pcre/src/pcre2_chartables.c
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
    add_library(libpcre2 SHARED ${SOURCE_FILES})

    target_include_directories(libpcre2
            PUBLIC
            ${ANDROID_ROOT}/external/pcre/include
            ${ANDROID_ROOT}/external/pcre
    )

    # Add c/c++ flags
    target_compile_options(libpcre2 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libpcre2
            libc
            libclang_rt.builtins
    )
endif ()
#################### libpcre2 ####################

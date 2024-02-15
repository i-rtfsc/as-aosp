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

project(liblua)

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

#################### liblua ####################
set(FILES
        ${ANDROID_ROOT}/external/lua/src/lapi.c
        ${ANDROID_ROOT}/external/lua/src/lauxlib.c
        ${ANDROID_ROOT}/external/lua/src/lbaselib.c
        ${ANDROID_ROOT}/external/lua/src/lcode.c
        ${ANDROID_ROOT}/external/lua/src/lcorolib.c
        ${ANDROID_ROOT}/external/lua/src/lctype.c
        ${ANDROID_ROOT}/external/lua/src/ldblib.c
        ${ANDROID_ROOT}/external/lua/src/ldebug.c
        ${ANDROID_ROOT}/external/lua/src/ldo.c
        ${ANDROID_ROOT}/external/lua/src/ldump.c
        ${ANDROID_ROOT}/external/lua/src/lfunc.c
        ${ANDROID_ROOT}/external/lua/src/lgc.c
        ${ANDROID_ROOT}/external/lua/src/linit.c
        ${ANDROID_ROOT}/external/lua/src/liolib.c
        ${ANDROID_ROOT}/external/lua/src/llex.c
        ${ANDROID_ROOT}/external/lua/src/lmathlib.c
        ${ANDROID_ROOT}/external/lua/src/lmem.c
        ${ANDROID_ROOT}/external/lua/src/loadlib.c
        ${ANDROID_ROOT}/external/lua/src/lobject.c
        ${ANDROID_ROOT}/external/lua/src/lopcodes.c
        ${ANDROID_ROOT}/external/lua/src/loslib.c
        ${ANDROID_ROOT}/external/lua/src/lparser.c
        ${ANDROID_ROOT}/external/lua/src/lstate.c
        ${ANDROID_ROOT}/external/lua/src/lstring.c
        ${ANDROID_ROOT}/external/lua/src/lstrlib.c
        ${ANDROID_ROOT}/external/lua/src/ltable.c
        ${ANDROID_ROOT}/external/lua/src/ltablib.c
        ${ANDROID_ROOT}/external/lua/src/ltm.c
        ${ANDROID_ROOT}/external/lua/src/luac.c
        ${ANDROID_ROOT}/external/lua/src/lundump.c
        ${ANDROID_ROOT}/external/lua/src/lutf8lib.c
        ${ANDROID_ROOT}/external/lua/src/lvm.c
        ${ANDROID_ROOT}/external/lua/src/lzio.c
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
    add_library(liblua SHARED ${SOURCE_FILES})

    target_include_directories(liblua
            PUBLIC
            ${ANDROID_ROOT}/external/lua/src
            ${ANDROID_ROOT}/external/lua
    )

    # Add c/c++ flags
    target_compile_options(liblua PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(liblua
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### liblua ####################

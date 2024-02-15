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

project(libicutoolutil)

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

#################### libicutoolutil ####################
set(FILES
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/collationinfo.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/dbgutil.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/denseranges.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/filestrm.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/filetools.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/flagparser.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/package.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/pkg_genc.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/pkg_gencmn.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/pkg_icu.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/pkgitems.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/ppucd.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/swapimpl.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/toolutil.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/ucbuf.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/ucln_tu.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/ucm.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/ucmstate.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/udbgutil.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/unewdata.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/uoptions.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/uparse.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/writesrc.cpp
        ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil/xmlparser.cpp
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
    add_library(libicutoolutil SHARED ${SOURCE_FILES})

    target_include_directories(libicutoolutil
            PUBLIC
            ${ANDROID_ROOT}/external/icu/icu4c/source/tools/toolutil
    )

    # Add c/c++ flags
    target_compile_options(libicutoolutil PRIVATE
            -Werror
    )

    target_link_libraries(libicutoolutil
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libicui18n
            libicuuc
            libm
    )
endif ()
#################### libicutoolutil ####################

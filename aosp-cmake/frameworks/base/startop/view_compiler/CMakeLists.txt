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

project(viewcompiler)

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

#################### libviewcompiler ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/startop/view_compiler/apk_layout_compiler.cc
        ${ANDROID_ROOT}/frameworks/base/startop/view_compiler/dex_builder.cc
        ${ANDROID_ROOT}/frameworks/base/startop/view_compiler/dex_layout_compiler.cc
        ${ANDROID_ROOT}/frameworks/base/startop/view_compiler/java_lang_builder.cc
        ${ANDROID_ROOT}/frameworks/base/startop/view_compiler/tinyxml_layout_parser.cc
        ${ANDROID_ROOT}/frameworks/base/startop/view_compiler/util.cc
        ${ANDROID_ROOT}/frameworks/base/startop/view_compiler/layout_validation.cc
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
    add_library(libviewcompiler SHARED ${SOURCE_FILES})

    target_include_directories(libviewcompiler
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/startop/view_compiler
    )

    # Add c/c++ flags
    target_compile_options(libviewcompiler PRIVATE
            -Werror
    )

    target_link_libraries(libviewcompiler
            libandroidfw
            libbase
            libbase_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libtinyxml2
            libutils
            libz
            libziparchive
            slicer
    )
endif ()
#################### libviewcompiler ####################

#################### viewcompiler ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/startop/view_compiler/main.cc
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
    add_executable(viewcompiler ${SOURCE_FILES})

    target_include_directories(viewcompiler
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/startop/view_compiler
    )

    # Add c/c++ flags
    target_compile_options(viewcompiler PRIVATE
            -Werror
    )

    target_link_libraries(viewcompiler
            libandroidfw
            libbase
            libbase_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgflags
            liblog
            libm
            libtinyxml2
            libutils
            libviewcompiler
            libz
            libziparchive
            slicer
    )
endif ()
#################### viewcompiler ####################

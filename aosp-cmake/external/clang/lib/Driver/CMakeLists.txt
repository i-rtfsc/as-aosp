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

project(libclangDriver)

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

#################### libclangDriver ####################
set(FILES
        ${ANDROID_ROOT}/external/clang/lib/Driver/Action.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/Compilation.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/CrossWindowsToolChain.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/Driver.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/DriverOptions.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/Job.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/MSVCToolChain.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/MinGWToolChain.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/Multilib.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/Phases.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/SanitizerArgs.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/Tool.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/ToolChain.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/ToolChains.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/Tools.cpp
        ${ANDROID_ROOT}/external/clang/lib/Driver/Types.cpp
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
    add_library(libclangDriver SHARED ${SOURCE_FILES})

    target_include_directories(libclangDriver
            PUBLIC
            ${ANDROID_ROOT}/external/clang/lib/Driver
    )

    # Add c/c++ flags
    target_compile_options(libclangDriver PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wno-unused-lambda-capture
            -Wno-unused-private-field
    )

    target_link_libraries(libclangDriver
            clang-headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            llvm-headers
            llvm-headers-no-generated-headers
    )
endif ()
#################### libclangDriver ####################

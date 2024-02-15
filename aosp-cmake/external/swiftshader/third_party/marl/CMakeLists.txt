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

project(swiftshader_marl)

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

#################### swiftshader_marl ####################
set(FILES
        ${ANDROID_ROOT}/external/swiftshader/third_party/marl/src/debug.cpp
        ${ANDROID_ROOT}/external/swiftshader/third_party/marl/src/memory.cpp
        ${ANDROID_ROOT}/external/swiftshader/third_party/marl/src/scheduler.cpp
        ${ANDROID_ROOT}/external/swiftshader/third_party/marl/src/thread.cpp
        ${ANDROID_ROOT}/external/swiftshader/third_party/marl/src/trace.cpp
        ${ANDROID_ROOT}/external/swiftshader/third_party/marl/src/osfiber_aarch64.c
        ${ANDROID_ROOT}/external/swiftshader/third_party/marl/src/osfiber_asm_aarch64.S
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
    add_library(swiftshader_marl SHARED ${SOURCE_FILES})

    target_include_directories(swiftshader_marl
            PUBLIC
            ${ANDROID_ROOT}/external/swiftshader/third_party/marl/include
            ${ANDROID_ROOT}/external/swiftshader/third_party/marl/src
            ${ANDROID_ROOT}/external/swiftshader/third_party/marl
    )

    # Add c/c++ flags
    target_compile_options(swiftshader_marl PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(swiftshader_marl
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### swiftshader_marl ####################

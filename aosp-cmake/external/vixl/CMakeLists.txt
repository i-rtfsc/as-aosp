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

project(libvixl)

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

#################### libvixl ####################
set(FILES
        ${ANDROID_ROOT}/external/vixl/src/code-buffer-vixl.cc
        ${ANDROID_ROOT}/external/vixl/src/compiler-intrinsics-vixl.cc
        ${ANDROID_ROOT}/external/vixl/src/cpu-features.cc
        ${ANDROID_ROOT}/external/vixl/src/utils-vixl.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/assembler-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/assembler-sve-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/cpu-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/cpu-features-auditor-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/decoder-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/disasm-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/instructions-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/logic-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/macro-assembler-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/macro-assembler-sve-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/operands-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/pointer-auth-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/registers-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/simulator-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/assembler-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/constants-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/disasm-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/instructions-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/location-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/macro-assembler-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/operands-aarch32.cc
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
    add_library(libvixl SHARED ${SOURCE_FILES})

    target_include_directories(libvixl
            PUBLIC
            ${ANDROID_ROOT}/external/vixl/src
            ${ANDROID_ROOT}/external/vixl
    )

    # Add c/c++ flags
    target_compile_options(libvixl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libvixl
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libvixl ####################

#################### libvixld ####################
set(FILES
        ${ANDROID_ROOT}/external/vixl/src/code-buffer-vixl.cc
        ${ANDROID_ROOT}/external/vixl/src/compiler-intrinsics-vixl.cc
        ${ANDROID_ROOT}/external/vixl/src/cpu-features.cc
        ${ANDROID_ROOT}/external/vixl/src/utils-vixl.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/assembler-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/assembler-sve-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/cpu-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/cpu-features-auditor-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/decoder-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/disasm-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/instructions-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/logic-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/macro-assembler-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/macro-assembler-sve-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/operands-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/pointer-auth-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/registers-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch64/simulator-aarch64.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/assembler-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/constants-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/disasm-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/instructions-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/location-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/macro-assembler-aarch32.cc
        ${ANDROID_ROOT}/external/vixl/src/aarch32/operands-aarch32.cc
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
    add_library(libvixld SHARED ${SOURCE_FILES})

    target_include_directories(libvixld
            PUBLIC
            ${ANDROID_ROOT}/external/vixl/src
            ${ANDROID_ROOT}/external/vixl
    )

    # Add c/c++ flags
    target_compile_options(libvixld PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libvixld
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libvixld ####################

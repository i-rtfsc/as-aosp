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

project(libcapstone)

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

#################### libcapstone ####################
set(FILES
        ${ANDROID_ROOT}/external/capstone/cs.c
        ${ANDROID_ROOT}/external/capstone/utils.c
        ${ANDROID_ROOT}/external/capstone/SStream.c
        ${ANDROID_ROOT}/external/capstone/MCInstrDesc.c
        ${ANDROID_ROOT}/external/capstone/MCRegisterInfo.c
        ${ANDROID_ROOT}/external/capstone/MCInst.c
        ${ANDROID_ROOT}/external/capstone/arch/ARM/ARMDisassembler.c
        ${ANDROID_ROOT}/external/capstone/arch/ARM/ARMInstPrinter.c
        ${ANDROID_ROOT}/external/capstone/arch/ARM/ARMMapping.c
        ${ANDROID_ROOT}/external/capstone/arch/ARM/ARMModule.c
        ${ANDROID_ROOT}/external/capstone/arch/AArch64/AArch64BaseInfo.c
        ${ANDROID_ROOT}/external/capstone/arch/AArch64/AArch64Disassembler.c
        ${ANDROID_ROOT}/external/capstone/arch/AArch64/AArch64InstPrinter.c
        ${ANDROID_ROOT}/external/capstone/arch/AArch64/AArch64Mapping.c
        ${ANDROID_ROOT}/external/capstone/arch/AArch64/AArch64Module.c
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
    add_library(libcapstone SHARED ${SOURCE_FILES})

    target_include_directories(libcapstone
            PUBLIC
            ${ANDROID_ROOT}/external/capstone
            ${ANDROID_ROOT}/external/capstone/include
    )

    # Add c/c++ flags
    target_compile_options(libcapstone PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libcapstone
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libcapstone ####################

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

project(libLLVMMipsCodeGen)

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

#################### libLLVMMipsCodeGen ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/Mips16FrameLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/Mips16HardFloat.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/Mips16HardFloatInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/Mips16ISelDAGToDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/Mips16ISelLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/Mips16InstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/Mips16RegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsAnalyzeImmediate.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsAsmPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsCCState.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsConstantIslandPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsDelaySlotFiller.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsFastISel.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsFrameLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsHazardSchedule.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsISelDAGToDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsISelLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsInstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsLongBranch.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsMCInstLower.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsMachineFunction.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsModuleISelDAGToDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsOptimizePICCall.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsOs16.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsRegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsSEFrameLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsSEISelDAGToDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsSEISelLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsSEInstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsSERegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsSubtarget.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsTargetMachine.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/Mips/MipsTargetObjectFile.cpp
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
    add_library(libLLVMMipsCodeGen SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMMipsCodeGen
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/Target/Mips
            ${OUT_INTERMEDIATES_ROOT}/external/llvm/lib/Target/Mips/llvm-gen-mips/gen
    )

    # Add c/c++ flags
    target_compile_options(libLLVMMipsCodeGen PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMMipsCodeGen
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            llvm-headers
            llvm-headers-no-generated-headers
            llvm-mips-headers
    )
endif ()
#################### libLLVMMipsCodeGen ####################

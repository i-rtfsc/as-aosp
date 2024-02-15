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

project(libLLVMARMCodeGen)

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

#################### libLLVMARMCodeGen ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/A15SDOptimizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMAsmPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMBaseInstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMBaseRegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMConstantIslandPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMConstantPoolValue.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMExpandPseudoInsts.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMFastISel.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMFrameLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMHazardRecognizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMISelDAGToDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMISelLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMInstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMLoadStoreOptimizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMMCInstLower.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMMachineFunctionInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMOptimizeBarriersPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMRegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMSelectionDAGInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMSubtarget.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMTargetMachine.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMTargetObjectFile.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ARMTargetTransformInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/MLxExpansionPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/Thumb1FrameLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/Thumb1InstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/Thumb2ITBlockPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/Thumb2InstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/Thumb2SizeReduction.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/ARM/ThumbRegisterInfo.cpp
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
    add_library(libLLVMARMCodeGen SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMARMCodeGen
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/Target/ARM
            ${OUT_INTERMEDIATES_ROOT}/external/llvm/lib/Target/ARM/llvm-gen-arm/gen
    )

    # Add c/c++ flags
    target_compile_options(libLLVMARMCodeGen PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMARMCodeGen
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            llvm-arm-headers
            llvm-headers
            llvm-headers-no-generated-headers
    )
endif ()
#################### libLLVMARMCodeGen ####################

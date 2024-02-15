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

project(libLLVMCore)

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

#################### libLLVMCore ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/IR/AsmWriter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Attributes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/AutoUpgrade.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/BasicBlock.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Comdat.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/ConstantFold.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/ConstantRange.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Constants.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Core.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/DIBuilder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/DataLayout.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/DebugInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/DebugInfoMetadata.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/DebugLoc.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/DiagnosticInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/DiagnosticPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Dominators.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Function.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/GCOV.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/GVMaterializer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Globals.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/IRBuilder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/IRPrintingPasses.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/InlineAsm.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Instruction.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Instructions.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/IntrinsicInst.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/LLVMContext.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/LLVMContextImpl.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/LegacyPassManager.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/MDBuilder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Mangler.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Metadata.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Module.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/ModuleSummaryIndex.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Operator.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/OptBisect.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Pass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/PassManager.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/PassRegistry.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/ProfileSummary.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Statepoint.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Type.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/TypeFinder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Use.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/User.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Value.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/ValueSymbolTable.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/ValueTypes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/IR/Verifier.cpp
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
    add_library(libLLVMCore SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMCore
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/IR
            ${OUT_INTERMEDIATES_ROOT}/external/llvm/lib/IR/llvm-gen-core/gen
    )

    # Add c/c++ flags
    target_compile_options(libLLVMCore PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMCore
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
#################### libLLVMCore ####################

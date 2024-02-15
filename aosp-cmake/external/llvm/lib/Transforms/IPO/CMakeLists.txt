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

project(libLLVMipo)

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

#################### libLLVMipo ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/ArgumentPromotion.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/BarrierNoopPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/ConstantMerge.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/CrossDSOCFI.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/DeadArgumentElimination.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/ElimAvailExtern.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/ExtractGV.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/ForceFunctionAttrs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/FunctionAttrs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/FunctionImport.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/GlobalDCE.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/GlobalOpt.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/IPConstantPropagation.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/IPO.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/InferFunctionAttrs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/InlineAlways.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/InlineSimple.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/Inliner.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/Internalize.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/LoopExtractor.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/LowerTypeTests.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/MergeFunctions.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/PartialInlining.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/PassManagerBuilder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/PruneEH.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/SampleProfile.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/StripDeadPrototypes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/StripSymbols.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO/WholeProgramDevirt.cpp
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
    add_library(libLLVMipo SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMipo
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/Transforms/IPO
    )

    # Add c/c++ flags
    target_compile_options(libLLVMipo PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMipo
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
#################### libLLVMipo ####################

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

project(libLLVMTransformUtils)

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

#################### libLLVMTransformUtils ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/ASanStackFrameLayout.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/AddDiscriminators.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/BasicBlockUtils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/BreakCriticalEdges.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/BuildLibCalls.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/BypassSlowDivision.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/CloneFunction.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/CloneModule.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/CmpInstAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/CodeExtractor.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/CtorUtils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/DemoteRegToStack.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/Evaluator.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/FlattenCFG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/FunctionImportUtils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/GlobalStatus.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/InlineFunction.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/InstructionNamer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/IntegerDivision.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/LCSSA.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/Local.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/LoopSimplify.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/LoopUnroll.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/LoopUnrollRuntime.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/LoopUtils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/LoopVersioning.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/LowerInvoke.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/LowerSwitch.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/Mem2Reg.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/MemorySSA.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/MetaRenamer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/ModuleUtils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/NameAnonFunctions.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/PromoteMemoryToRegister.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/SSAUpdater.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/SanitizerStats.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/SimplifyCFG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/SimplifyIndVar.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/SimplifyInstructions.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/SimplifyLibCalls.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/SplitModule.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/SymbolRewriter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/UnifyFunctionExitNodes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/Utils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils/ValueMapper.cpp
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
    add_library(libLLVMTransformUtils SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMTransformUtils
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/Transforms/Utils
    )

    # Add c/c++ flags
    target_compile_options(libLLVMTransformUtils PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMTransformUtils
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
#################### libLLVMTransformUtils ####################

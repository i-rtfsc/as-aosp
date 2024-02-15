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

project(libLLVMScalarOpts)

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

#################### libLLVMScalarOpts ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/ADCE.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/AlignmentFromAssumptions.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/BDCE.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/ConstantHoisting.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/ConstantProp.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/CorrelatedValuePropagation.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/DCE.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/DeadStoreElimination.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/EarlyCSE.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/FlattenCFGPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/Float2Int.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/GVN.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/GuardWidening.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/IndVarSimplify.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/InductiveRangeCheckElimination.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/JumpThreading.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LICM.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoadCombine.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopDataPrefetch.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopDeletion.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopDistribute.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopIdiomRecognize.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopInstSimplify.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopInterchange.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopLoadElimination.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopRerollPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopRotation.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopSimplifyCFG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopStrengthReduce.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopUnrollPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopUnswitch.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LoopVersioningLICM.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LowerAtomic.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LowerExpectIntrinsic.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/LowerGuardIntrinsic.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/MemCpyOptimizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/MergedLoadStoreMotion.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/NaryReassociate.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/PartiallyInlineLibCalls.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/PlaceSafepoints.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/Reassociate.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/Reg2Mem.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/RewriteStatepointsForGC.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/SCCP.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/SROA.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/Scalar.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/Scalarizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/SeparateConstOffsetFromGEP.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/SimplifyCFGPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/Sink.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/SpeculativeExecution.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/StraightLineStrengthReduce.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/StructurizeCFG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar/TailRecursionElimination.cpp
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
    add_library(libLLVMScalarOpts SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMScalarOpts
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/Transforms/Scalar
    )

    # Add c/c++ flags
    target_compile_options(libLLVMScalarOpts PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMScalarOpts
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
#################### libLLVMScalarOpts ####################

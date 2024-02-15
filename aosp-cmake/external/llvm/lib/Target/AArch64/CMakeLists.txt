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

project(libLLVMAArch64CodeGen)

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

#################### libLLVMAArch64CodeGen ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64A53Fix835769.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64A57FPLoadBalancing.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64AddressTypePromotion.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64AdvSIMDScalarPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64AsmPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64BranchRelaxation.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64CleanupLocalDynamicTLSPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64CollectLOH.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64ConditionOptimizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64ConditionalCompares.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64DeadRegisterDefinitionsPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64ExpandPseudoInsts.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64FastISel.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64FrameLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64ISelDAGToDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64ISelLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64InstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64LoadStoreOptimizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64MCInstLower.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64PBQPRegAlloc.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64PromoteConstant.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64RedundantCopyElimination.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64RegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64SelectionDAGInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64StorePairSuppress.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64Subtarget.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64TargetMachine.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64TargetObjectFile.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64/AArch64TargetTransformInfo.cpp
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
    add_library(libLLVMAArch64CodeGen SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMAArch64CodeGen
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/Target/AArch64
            ${OUT_INTERMEDIATES_ROOT}/external/llvm/lib/Target/AArch64/llvm-gen-aarch64/gen
    )

    # Add c/c++ flags
    target_compile_options(libLLVMAArch64CodeGen PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMAArch64CodeGen
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            llvm-aarch64-headers
            llvm-headers
            llvm-headers-no-generated-headers
    )
endif ()
#################### libLLVMAArch64CodeGen ####################

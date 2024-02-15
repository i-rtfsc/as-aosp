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

project(libLLVMAnalysis)

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

#################### libLLVMAnalysis ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/AliasAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/AliasAnalysisEvaluator.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/AliasAnalysisSummary.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/AliasSetTracker.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/Analysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/AssumptionCache.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/BasicAliasAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/BlockFrequencyInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/BlockFrequencyInfoImpl.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/BranchProbabilityInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CFG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CFGPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CFLAndersAliasAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CFLSteensAliasAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CGSCCPassManager.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CallGraph.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CallGraphSCCPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CallPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CaptureTracking.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CodeMetrics.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ConstantFolding.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/CostModel.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/Delinearization.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/DemandedBits.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/DependenceAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/DivergenceAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/DomPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/DominanceFrontier.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/EHPersonalities.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/GlobalsModRef.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/IVUsers.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/IndirectCallPromotionAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/InlineCost.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/InstCount.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/InstructionSimplify.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/Interval.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/IntervalPartition.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/IteratedDominanceFrontier.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/LazyBlockFrequencyInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/LazyCallGraph.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/LazyValueInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/Lint.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/Loads.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/LoopAccessAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/LoopInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/LoopPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/LoopPassManager.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/LoopUnrollAnalyzer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/MemDepPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/MemDerefPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/MemoryBuiltins.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/MemoryDependenceAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/MemoryLocation.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ModuleDebugInfoPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ModuleSummaryAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ObjCARCAliasAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ObjCARCAnalysisUtils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ObjCARCInstKind.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/OrderedBasicBlock.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/PHITransAddr.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/PostDominators.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ProfileSummaryInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/PtrUseVisitor.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/RegionInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/RegionPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/RegionPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ScalarEvolution.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ScalarEvolutionAliasAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ScalarEvolutionExpander.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ScalarEvolutionNormalization.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ScopedNoAliasAA.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/SparsePropagation.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/TargetLibraryInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/TargetTransformInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/Trace.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/TypeBasedAliasAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/TypeMetadataUtils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/ValueTracking.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Analysis/VectorUtils.cpp
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
    add_library(libLLVMAnalysis SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMAnalysis
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/Analysis
    )

    # Add c/c++ flags
    target_compile_options(libLLVMAnalysis PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMAnalysis
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
#################### libLLVMAnalysis ####################

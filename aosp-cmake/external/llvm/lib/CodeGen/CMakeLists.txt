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

project(libLLVMCodeGen)

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

#################### libLLVMCodeGen ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/AggressiveAntiDepBreaker.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/AllocationOrder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/Analysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/AtomicExpandPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/BasicTargetTransformInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/BranchFolding.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/BuiltinGCs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/CalcSpillWeights.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/CallingConvLower.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/CodeGen.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/CodeGenPrepare.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/CriticalAntiDepBreaker.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/DFAPacketizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/DeadMachineInstructionElim.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/DetectDeadLanes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/DwarfEHPrepare.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/EarlyIfConversion.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/EdgeBundles.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ExecutionDepsFix.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ExpandISelPseudos.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ExpandPostRAPseudos.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/FaultMaps.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/FuncletLayout.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/GCMetadata.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/GCMetadataPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/GCRootLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/GCStrategy.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/GlobalMerge.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/IfConversion.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ImplicitNullChecks.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/InlineSpiller.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/InterferenceCache.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/InterleavedAccessPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/IntrinsicLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LLVMTargetMachine.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LatencyPriorityQueue.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LexicalScopes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveDebugValues.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveDebugVariables.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveInterval.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveIntervalAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveIntervalUnion.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LivePhysRegs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveRangeCalc.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveRangeEdit.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveRegMatrix.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveStackAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LiveVariables.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LocalStackSlotAllocation.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/LowerEmuTLS.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MIRPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MIRPrintingPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineBasicBlock.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineBlockFrequencyInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineBlockPlacement.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineBranchProbabilityInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineCSE.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineCombiner.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineCopyPropagation.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineDominanceFrontier.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineDominators.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineFunction.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineFunctionAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineFunctionPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineFunctionPrinterPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineInstr.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineInstrBundle.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineLICM.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineLoopInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineModuleInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineModuleInfoImpls.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachinePassRegistry.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachinePipeliner.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachinePostDominators.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineRegionInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineRegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineSSAUpdater.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineScheduler.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineSink.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineTraceMetrics.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/MachineVerifier.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/OptimizePHIs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/PHIElimination.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/PHIEliminationUtils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ParallelCG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/PatchableFunction.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/PeepholeOptimizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/PostRAHazardRecognizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/PostRASchedulerList.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/PreISelIntrinsicLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ProcessImplicitDefs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/PrologEpilogInserter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/PseudoSourceValue.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegAllocBase.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegAllocBasic.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegAllocFast.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegAllocGreedy.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegAllocPBQP.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegUsageInfoCollector.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegUsageInfoPropagate.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegisterClassInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegisterCoalescer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegisterPressure.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegisterScavenging.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RegisterUsageInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/RenameIndependentSubregs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SafeStack.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SafeStackColoring.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SafeStackLayout.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ScheduleDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ScheduleDAGInstrs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ScheduleDAGPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ScoreboardHazardRecognizer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ShadowStackGCLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/ShrinkWrap.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SjLjEHPrepare.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SlotIndexes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SpillPlacement.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SplitKit.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/StackColoring.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/StackMapLivenessAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/StackMaps.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/StackProtector.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/StackSlotColoring.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TailDuplication.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TailDuplicator.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TargetFrameLoweringImpl.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TargetInstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TargetLoweringBase.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TargetLoweringObjectFileImpl.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TargetOptionsImpl.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TargetPassConfig.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TargetRegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TargetSchedule.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/TwoAddressInstructionPass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/UnreachableBlockElim.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/VirtRegMap.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/WinEHPrepare.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/XRayInstrumentation.cpp
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
    add_library(libLLVMCodeGen SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMCodeGen
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/CodeGen
    )

    # Add c/c++ flags
    target_compile_options(libLLVMCodeGen PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMCodeGen
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
#################### libLLVMCodeGen ####################

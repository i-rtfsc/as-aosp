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

project(libLLVMSelectionDAG)

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

#################### libLLVMSelectionDAG ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/DAGCombiner.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/FastISel.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/FunctionLoweringInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/InstrEmitter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/LegalizeDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/LegalizeFloatTypes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/LegalizeIntegerTypes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/LegalizeTypes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/LegalizeTypesGeneric.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/LegalizeVectorOps.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/LegalizeVectorTypes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/ResourcePriorityQueue.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/ScheduleDAGFast.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/ScheduleDAGRRList.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/ScheduleDAGSDNodes.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/ScheduleDAGVLIW.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/SelectionDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/SelectionDAGBuilder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/SelectionDAGDumper.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/SelectionDAGISel.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/SelectionDAGPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/SelectionDAGTargetInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/StatepointLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG/TargetLowering.cpp
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
    add_library(libLLVMSelectionDAG SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMSelectionDAG
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/CodeGen/SelectionDAG
    )

    # Add c/c++ flags
    target_compile_options(libLLVMSelectionDAG PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMSelectionDAG
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
#################### libLLVMSelectionDAG ####################

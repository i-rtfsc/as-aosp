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

project(android.hardware.neuralnetworks-V1-ndk_platform)

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

#################### android.hardware.neuralnetworks-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/android/hardware/neuralnetworks/Timing.cpp
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
    add_library(android.hardware.neuralnetworks-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V1-ndk_platform ####################

#################### android.hardware.neuralnetworks-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/ExecutionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/IExecution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/PrepareModelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/Timing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/android/hardware/neuralnetworks/TokenValuePair.cpp
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
    add_library(android.hardware.neuralnetworks-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V4-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V4-ndk ####################

#################### android.hardware.neuralnetworks-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/android/hardware/neuralnetworks/Timing.cpp
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
    add_library(android.hardware.neuralnetworks-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V1-ndk ####################

#################### android.hardware.neuralnetworks-V5-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/ExecutionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/IExecution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/PrepareModelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/Timing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/android/hardware/neuralnetworks/TokenValuePair.cpp
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
    add_library(android.hardware.neuralnetworks-V5-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V5-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V5-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V5-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V5-ndk ####################

#################### android.hardware.neuralnetworks-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/android/hardware/neuralnetworks/Timing.cpp
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
    add_library(android.hardware.neuralnetworks-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V3-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V3-ndk ####################

#################### android.hardware.neuralnetworks-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/android/hardware/neuralnetworks/Timing.cpp
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
    add_library(android.hardware.neuralnetworks-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V2-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V2-ndk ####################

#################### android.hardware.neuralnetworks-V5-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/IExecution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/PrepareModelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/Timing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/android/hardware/neuralnetworks/TokenValuePair.cpp
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
    add_library(android.hardware.neuralnetworks-V5-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V5-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V5-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V5-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V5-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V5-ndk_platform ####################

#################### android.hardware.neuralnetworks-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/android/hardware/neuralnetworks/Timing.cpp
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
    add_library(android.hardware.neuralnetworks-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V2-ndk_platform ####################

#################### android.hardware.neuralnetworks-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/android/hardware/neuralnetworks/Timing.cpp
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
    add_library(android.hardware.neuralnetworks-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V3-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V3-ndk_platform ####################

#################### android.hardware.neuralnetworks-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/BufferRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/DataLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/DeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/ErrorStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/ExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Extension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionNameAndPrefix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/ExtensionOperandTypeInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/FencedExecutionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/FusedActivationFunc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/IBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/IBurst.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/IDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/IExecution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/IFencedExecutionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/IPreparedModelParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Memory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Model.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/NumberOfCacheFiles.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Operand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandExtraParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandLifeTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandPerformance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/OperandType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Operation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/OperationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/OutputShape.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/PerformanceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/PrepareModelConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Priority.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Request.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestArgument.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/RequestMemoryPool.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Subgraph.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/SymmPerChannelQuantParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/Timing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/android/hardware/neuralnetworks/TokenValuePair.cpp
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
    add_library(android.hardware.neuralnetworks-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/neuralnetworks/aidl/android.hardware.neuralnetworks-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.neuralnetworks-V4-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-V4-ndk_platform ####################

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

project(android.hardware.bluetooth.audio-V3-ndk_platform)

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

#################### android.hardware.bluetooth.audio-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacObjectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveInputMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveLeCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveLeConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveTimeToPlay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxSinkBuffering.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/BluetoothAudioStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/BroadcastCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/ChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProviderFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/LatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/Lc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/Lc3Configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacQualityIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioBroadcastConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioCodecCapabilitiesSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioCodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/OpusCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/OpusConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/PcmCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/PcmConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/PresentationPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcAllocMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/SessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/android/hardware/bluetooth/audio/UnicastCapability.cpp
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
    add_library(android.hardware.bluetooth.audio-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth.audio-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth.audio-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.bluetooth.audio-V3-ndk_platform
            android.hardware.audio.common-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
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
#################### android.hardware.bluetooth.audio-V3-ndk_platform ####################

#################### android.hardware.bluetooth.audio-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacObjectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveInputMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveLeCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveLeConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveTimeToPlay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxSinkBuffering.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/BluetoothAudioStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/BroadcastCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/ChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProviderFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/LatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/Lc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/Lc3Configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacQualityIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioBroadcastConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioCodecCapabilitiesSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioCodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/OpusCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/OpusConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/PcmCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/PcmConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/PresentationPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcAllocMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/SessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/android/hardware/bluetooth/audio/UnicastCapability.cpp
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
    add_library(android.hardware.bluetooth.audio-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth.audio-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth.audio-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.bluetooth.audio-V4-ndk_platform
            android.hardware.audio.common-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
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
#################### android.hardware.bluetooth.audio-V4-ndk_platform ####################

#################### android.hardware.bluetooth.audio-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacObjectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveInputMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveTimeToPlay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxSinkBuffering.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/BluetoothAudioStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/BroadcastCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/ChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProviderFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/LatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/Lc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/Lc3Configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacQualityIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioBroadcastConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioCodecCapabilitiesSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioCodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/OpusCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/OpusConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/PcmCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/PcmConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/PresentationPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcAllocMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/SessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/android/hardware/bluetooth/audio/UnicastCapability.cpp
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
    add_library(android.hardware.bluetooth.audio-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth.audio-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth.audio-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.bluetooth.audio-V2-ndk_platform
            android.hardware.audio.common-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
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
#################### android.hardware.bluetooth.audio-V2-ndk_platform ####################

#################### android.hardware.bluetooth.audio-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AacObjectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveInputMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveTimeToPlay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AptxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AptxConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AptxMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AptxSinkBuffering.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/AudioLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/BluetoothAudioStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/BroadcastCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/ChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/CodecCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/CodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProviderFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/LatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/Lc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/Lc3Configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/LdacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/LdacChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/LdacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/LdacQualityIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioBroadcastConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioCodecCapabilitiesSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioCodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/OpusCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/OpusConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/PcmCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/PcmConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/PresentationPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/SbcAllocMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/SbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/SbcChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/SbcConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/SessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/android/hardware/bluetooth/audio/UnicastCapability.cpp
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
    add_library(android.hardware.bluetooth.audio-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth.audio-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth.audio-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.bluetooth.audio-V2-ndk
            android.hardware.audio.common-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
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
#################### android.hardware.bluetooth.audio-V2-ndk ####################

#################### android.hardware.bluetooth.audio-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AacObjectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveInputMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveLeCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveLeConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveTimeToPlay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AptxSinkBuffering.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/AudioLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/BluetoothAudioStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/BroadcastCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/ChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/CodecCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/CodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProviderFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/LatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/Lc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/Lc3Configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/LdacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/LdacChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/LdacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/LdacQualityIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioBroadcastConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioCodecCapabilitiesSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioCodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/OpusCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/OpusConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/PcmCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/PcmConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/PresentationPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/SbcAllocMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/SbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/SbcChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/SbcConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/SessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/android/hardware/bluetooth/audio/UnicastCapability.cpp
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
    add_library(android.hardware.bluetooth.audio-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth.audio-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth.audio-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.bluetooth.audio-V3-ndk
            android.hardware.audio.common-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
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
#################### android.hardware.bluetooth.audio-V3-ndk ####################

#################### android.hardware.bluetooth.audio-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AacObjectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveInputMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveTimeToPlay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AptxSinkBuffering.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/AudioLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/BluetoothAudioStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/BroadcastCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/ChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProviderFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/LatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/Lc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/Lc3Configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/LdacQualityIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioBroadcastConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioCodecCapabilitiesSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioCodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/LeAudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/PcmCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/PcmConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/PresentationPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcAllocMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/SbcConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/SessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/android/hardware/bluetooth/audio/UnicastCapability.cpp
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
    add_library(android.hardware.bluetooth.audio-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth.audio-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth.audio-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.bluetooth.audio-V1-ndk_platform
            android.hardware.audio.common-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
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
#################### android.hardware.bluetooth.audio-V1-ndk_platform ####################

#################### android.hardware.bluetooth.audio-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AacObjectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveInputMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveLeCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveLeConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveTimeToPlay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AptxSinkBuffering.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/AudioLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/BluetoothAudioStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/BroadcastCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/ChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/CodecCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/CodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProviderFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/LatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/Lc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/Lc3Configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/LdacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/LdacChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/LdacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/LdacQualityIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioBroadcastConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioCodecCapabilitiesSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioCodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/OpusCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/OpusConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/PcmCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/PcmConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/PresentationPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/SbcAllocMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/SbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/SbcChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/SbcConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/SessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/android/hardware/bluetooth/audio/UnicastCapability.cpp
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
    add_library(android.hardware.bluetooth.audio-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth.audio-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth.audio-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.bluetooth.audio-V4-ndk
            android.hardware.audio.common-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
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
#################### android.hardware.bluetooth.audio-V4-ndk ####################

#################### android.hardware.bluetooth.audio-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AacObjectType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveInputMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AptxAdaptiveTimeToPlay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AptxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AptxConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AptxMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AptxSinkBuffering.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/AudioLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/BluetoothAudioStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/BroadcastCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/ChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/CodecCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/CodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/CodecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/IBluetoothAudioProviderFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/LatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/Lc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/Lc3Configuration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/LdacCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/LdacChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/LdacConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/LdacQualityIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioBroadcastConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioCodecCapabilitiesSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioCodecConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/LeAudioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/PcmCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/PcmConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/PresentationPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/SbcAllocMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/SbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/SbcChannelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/SbcConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/SessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/android/hardware/bluetooth/audio/UnicastCapability.cpp
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
    add_library(android.hardware.bluetooth.audio-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth.audio-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/aidl/android.hardware.bluetooth.audio-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth.audio-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.bluetooth.audio-V1-ndk
            android.hardware.audio.common-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
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
#################### android.hardware.bluetooth.audio-V1-ndk ####################

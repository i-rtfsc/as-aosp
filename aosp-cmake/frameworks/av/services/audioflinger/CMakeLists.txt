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

project(libaudioflinger)

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

#################### libaudioflinger ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/AudioFlinger.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/AudioHwDevice.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/AudioStreamOut.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/AudioWatchdog.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/BufLog.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/DeviceEffectManager.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/Effects.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/FastCapture.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/FastCaptureDumpState.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/FastCaptureState.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/FastMixer.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/FastMixerDumpState.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/FastMixerState.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/FastThread.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/FastThreadDumpState.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/FastThreadState.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/MelReporter.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/NBAIO_Tee.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/PatchCommandThread.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/PatchPanel.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/PropertyUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/SpdifStreamOut.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/StateQueue.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/Threads.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/Tracks.cpp
        ${ANDROID_ROOT}/frameworks/av/services/audioflinger/TypedLogger.cpp
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
    add_library(libaudioflinger SHARED ${SOURCE_FILES})

    target_include_directories(libaudioflinger
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/audioflinger
    )

    # Add c/c++ flags
    target_compile_options(libaudioflinger PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
            -Wextra
            -Werror
            -Wall
    )

    target_link_libraries(libaudioflinger
            android.hardware.audio.core.sounddose-V1-ndk
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            audioflinger-aidl-cpp
            av-types-aidl-cpp
            effect-aidl-cpp
            libaaudio_headers
            libactivitymanager_aidl
            libaudioclient
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudioflinger_timing
            libaudiofoundation
            libaudiohal
            libaudiohal_headers
            libaudiomanager
            libaudioprocessing
            libaudiospdif
            libaudioutils
            libaudioutils_headers
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcpustats
            libcutils
            libdl
            liblog
            libm
            libmedia_headers
            libmedia_helper
            libmedialogservice
            libmediametrics
            libmediautils
            libmemunreachable
            libnbaio
            libnblog
            libpermission
            libpowermanager
            libshmemcompat
            libsndfile
            libsounddose
            libutils
            libvibrator
            packagemanager_aidl-cpp
    )
endif ()
#################### libaudioflinger ####################

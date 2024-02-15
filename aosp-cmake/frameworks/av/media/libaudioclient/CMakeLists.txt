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

project(sounddose-aidl-cpp-analyzer)

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

#################### effect-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk_platform-source/gen/android/media/EffectConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk_platform-source/gen/android/media/IEffect.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk_platform-source/gen/android/media/IEffectClient.cpp
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
    add_library(effect-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(effect-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(effect-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(effect-aidl-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            shared-file-region-aidl-ndk_platform
    )
endif ()
#################### effect-aidl-ndk_platform ####################

#################### capture_state_listener-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-ndk_platform-source/gen/android/media/ICaptureStateListener.cpp
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
    add_library(capture_state_listener-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(capture_state_listener-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(capture_state_listener-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(capture_state_listener-aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### capture_state_listener-aidl-ndk_platform ####################

#################### audiopolicy-types-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioAttributesEx.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioMix.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioMixerAttributesInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioMixerBehavior.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioMixCallbackFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioMixMatchCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioMixMatchCriterionValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioMixRouteFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioMixType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioOffloadMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioPolicyDeviceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioPolicyForceUse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioPolicyForcedConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/AudioVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/DeviceRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/SoundTriggerSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/SpatializationLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/SpatializationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/android/media/SpatializerHeadTrackingMode.cpp
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
    add_library(audiopolicy-types-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(audiopolicy-types-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audiopolicy-types-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audiopolicy-types-aidl-cpp
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### audiopolicy-types-aidl-cpp ####################

#################### effect-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/android/media/EffectConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/android/media/IEffect.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/android/media/IEffectClient.cpp
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
    add_library(effect-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(effect-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(effect-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(effect-aidl-cpp
            android.media.audio.common.types-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
            shared-file-region-aidl-cpp
    )
endif ()
#################### effect-aidl-cpp ####################

#################### audiopolicy-types-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioAttributesEx.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioMix.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioMixerAttributesInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioMixerBehavior.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioMixCallbackFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioMixMatchCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioMixMatchCriterionValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioMixRouteFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioMixType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioOffloadMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioPolicyDeviceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioPolicyForceUse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioPolicyForcedConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/AudioVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/DeviceRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/SoundTriggerSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/SpatializationLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/SpatializationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/android/media/SpatializerHeadTrackingMode.cpp
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
    add_library(audiopolicy-types-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(audiopolicy-types-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audiopolicy-types-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audiopolicy-types-aidl-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
            audioclient-types-aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### audiopolicy-types-aidl-ndk_platform ####################

#################### audiopolicy-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk_platform-source/gen/android/media/GetInputForAttrResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk_platform-source/gen/android/media/GetOutputForAttrResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk_platform-source/gen/android/media/GetSpatializerResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk_platform-source/gen/android/media/RecordClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk_platform-source/gen/android/media/IAudioPolicyService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk_platform-source/gen/android/media/IAudioPolicyServiceClient.cpp
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
    add_library(audiopolicy-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(audiopolicy-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audiopolicy-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audiopolicy-aidl-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
            audioclient-types-aidl-ndk_platform
            audiopolicy-types-aidl-ndk_platform
            capture_state_listener-aidl-ndk_platform
            framework-permission-aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            spatializer-aidl-ndk_platform
    )
endif ()
#################### audiopolicy-aidl-ndk_platform ####################

#################### libaudiopolicy ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/VolumeGroupAttributes.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/AudioPolicy.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/AudioProductStrategy.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/AudioVolumeGroup.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/PolicyAidlConversion.cpp
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
    add_library(libaudiopolicy SHARED ${SOURCE_FILES})

    target_include_directories(libaudiopolicy
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libaudiopolicy PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudiopolicy
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            audioflinger-aidl-cpp
            audiopolicy-aidl-cpp
            audiopolicy-types-aidl-cpp
            capture_state_listener-aidl-cpp
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudiofoundation
            libaudioutils
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libaudiopolicy ####################

#################### audiopolicy-types-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioAttributesEx.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioMix.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioMixerAttributesInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioMixerBehavior.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioMixCallbackFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioMixMatchCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioMixMatchCriterionValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioMixRouteFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioMixType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioOffloadMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioPolicyDeviceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioPolicyForceUse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioPolicyForcedConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/AudioVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/DeviceRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/SoundTriggerSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/SpatializationLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/SpatializationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/android/media/SpatializerHeadTrackingMode.cpp
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
    add_library(audiopolicy-types-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(audiopolicy-types-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audiopolicy-types-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audiopolicy-types-aidl-ndk
            android.media.audio.common.types-V2-ndk
            audioclient-types-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### audiopolicy-types-aidl-ndk ####################

#################### libaudioclient ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/libaudioclient/android_${OUT_ARCH_CPU}_static/gen/aidl/android/media/IPlayer.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/libaudioclient/android_${OUT_ARCH_CPU}_static/gen/aidl/android/media/AudioHalVersion.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/AudioEffect.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/AudioRecord.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/AudioSystem.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/AudioTrack.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/AudioTrackShared.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/IAudioFlinger.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/ToneGenerator.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/PlayerBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/RecordingActivityTracker.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/TrackPlayerBase.cpp
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
    add_library(libaudioclient SHARED ${SOURCE_FILES})

    target_include_directories(libaudioclient
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include/media
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/aidl
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/libaudioclient/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libaudioclient PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libaudioclient
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            audioflinger-aidl-cpp
            audiopolicy-aidl-cpp
            audiopolicy-types-aidl-cpp
            av-types-aidl-cpp
            capture_state_listener-aidl-cpp
            effect-aidl-cpp
            framework-permission-aidl-cpp
            libaudio_aidl_conversion_common_cpp
            libaudioclient_aidl_conversion
            libaudioclient_headers
            libaudiofoundation
            libaudiomanager
            libaudiopolicy
            libaudioutils
            libavmediaextentions
            libbase_headers
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libc_malloc_debug_backtrace
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libmedia_headers
            libmedia_helper
            libmediametrics
            libmediautils
            libnblog
            libprocessgroup
            libshmemcompat
            libutils
            packagemanager_aidl-cpp
            shared-file-region-aidl-cpp
            spatializer-aidl-cpp
    )
endif ()
#################### libaudioclient ####################

#################### capture_state_listener-aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-analyzer-source/gen/android/media/ICaptureStateListener.cpp
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
    add_executable(capture_state_listener-aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(capture_state_listener-aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(capture_state_listener-aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(capture_state_listener-aidl-cpp-analyzer
            aidl-analyzer-main
            capture_state_listener-aidl-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### capture_state_listener-aidl-cpp-analyzer ####################

#################### effect-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk-source/gen/android/media/EffectConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk-source/gen/android/media/IEffect.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk-source/gen/android/media/IEffectClient.cpp
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
    add_library(effect-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(effect-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(effect-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(effect-aidl-ndk
            android.media.audio.common.types-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            shared-file-region-aidl-ndk
    )
endif ()
#################### effect-aidl-ndk ####################

#################### sounddose-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-ndk_platform-source/gen/android/media/ISoundDose.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-ndk_platform-source/gen/android/media/ISoundDoseCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-ndk_platform-source/gen/android/media/SoundDoseRecord.cpp
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
    add_library(sounddose-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(sounddose-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(sounddose-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(sounddose-aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### sounddose-aidl-ndk_platform ####################

#################### sounddose-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-ndk-source/gen/android/media/ISoundDose.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-ndk-source/gen/android/media/ISoundDoseCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-ndk-source/gen/android/media/SoundDoseRecord.cpp
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
    add_library(sounddose-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(sounddose-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(sounddose-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(sounddose-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### sounddose-aidl-ndk ####################

#################### capture_state_listener-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/android/media/ICaptureStateListener.cpp
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
    add_library(capture_state_listener-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(capture_state_listener-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(capture_state_listener-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(capture_state_listener-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### capture_state_listener-aidl-cpp ####################

#################### spatializer-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk-source/gen/android/media/INativeSpatializerCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk-source/gen/android/media/ISpatializer.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk-source/gen/android/media/ISpatializerHeadTrackingCallback.cpp
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
    add_library(spatializer-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(spatializer-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(spatializer-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(spatializer-aidl-ndk
            audiopolicy-types-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### spatializer-aidl-ndk ####################

#################### audioflinger-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/CreateEffectRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/CreateEffectResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/CreateRecordRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/CreateRecordResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/CreateTrackRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/CreateTrackResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/OpenInputRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/OpenInputResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/OpenOutputRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/OpenOutputResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/RenderPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/IAudioFlingerService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/IAudioFlingerClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/IAudioRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/IAudioTrack.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/IAudioTrackCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/ISoundDose.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/ISoundDoseCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/android/media/SoundDoseRecord.cpp
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
    add_library(audioflinger-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(audioflinger-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audioflinger-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audioflinger-aidl-ndk
            android.media.audio.common.types-V2-ndk
            audioclient-types-aidl-ndk
            av-types-aidl-ndk
            effect-aidl-ndk
            framework-permission-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            shared-file-region-aidl-ndk
    )
endif ()
#################### audioflinger-aidl-ndk ####################

#################### audiopolicy-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/android/media/GetInputForAttrResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/android/media/GetOutputForAttrResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/android/media/GetSpatializerResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/android/media/RecordClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/android/media/IAudioPolicyService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/android/media/IAudioPolicyServiceClient.cpp
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
    add_library(audiopolicy-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(audiopolicy-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audiopolicy-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audiopolicy-aidl-cpp
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            audiopolicy-types-aidl-cpp
            capture_state_listener-aidl-cpp
            framework-permission-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
            spatializer-aidl-cpp
    )
endif ()
#################### audiopolicy-aidl-cpp ####################

#################### audioclient-types-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioAttributesInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioDirectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioGainSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioHalVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioHwModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioIoConfigEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioIoDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPatchFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPolicyConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPortFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPortSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPortConfigFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPortConfigSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPortDeviceExtSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPortExtSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPortMixExtSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPortRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioPortType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioProfileSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioRoute.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioTimestampInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioUniqueIdUse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/AudioVibratorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/DeviceConnectedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/EffectDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/TrackSecondaryOutputInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/android/media/SurroundSoundConfig.cpp
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
    add_library(audioclient-types-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(audioclient-types-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audioclient-types-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audioclient-types-aidl-cpp
            android.media.audio.common.types-V2-cpp
            framework-permission-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### audioclient-types-aidl-cpp ####################

#################### audioflinger-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/CreateEffectRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/CreateEffectResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/CreateRecordRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/CreateRecordResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/CreateTrackRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/CreateTrackResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/OpenInputRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/OpenInputResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/OpenOutputRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/OpenOutputResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/RenderPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/IAudioFlingerService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/IAudioFlingerClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/IAudioRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/IAudioTrack.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/IAudioTrackCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/ISoundDose.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/ISoundDoseCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/android/media/SoundDoseRecord.cpp
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
    add_library(audioflinger-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(audioflinger-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audioflinger-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audioflinger-aidl-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
            audioclient-types-aidl-ndk_platform
            av-types-aidl-ndk_platform
            effect-aidl-ndk_platform
            framework-permission-aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            shared-file-region-aidl-ndk_platform
    )
endif ()
#################### audioflinger-aidl-ndk_platform ####################

#################### libaudioclient_aidl_conversion ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/AidlConversion.cpp
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
    add_library(libaudioclient_aidl_conversion SHARED ${SOURCE_FILES})

    target_include_directories(libaudioclient_aidl_conversion
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient/include
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libaudioclient_aidl_conversion PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libaudioclient_aidl_conversion
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            framework-permission-aidl-cpp
            libaudio_aidl_conversion_common_cpp
            libaudio_aidl_conversion_common_util_cpp
            libaudio_system_headers
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhardware_headers
            liblog
            libm
            libshmemcompat
            libstagefright_foundation
            libutils
            shared-file-region-aidl-cpp
    )
endif ()
#################### libaudioclient_aidl_conversion ####################

#################### sounddose-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-cpp-source/gen/android/media/ISoundDose.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-cpp-source/gen/android/media/ISoundDoseCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-cpp-source/gen/android/media/SoundDoseRecord.cpp
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
    add_library(sounddose-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(sounddose-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(sounddose-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(sounddose-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### sounddose-aidl-cpp ####################

#################### capture_state_listener-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-ndk-source/gen/android/media/ICaptureStateListener.cpp
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
    add_library(capture_state_listener-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(capture_state_listener-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(capture_state_listener-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(capture_state_listener-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### capture_state_listener-aidl-ndk ####################

#################### sounddose-aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-cpp-analyzer-source/gen/android/media/ISoundDose.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-cpp-analyzer-source/gen/android/media/ISoundDoseCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-cpp-analyzer-source/gen/android/media/SoundDoseRecord.cpp
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
    add_executable(sounddose-aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(sounddose-aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudioclient
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/sounddose-aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(sounddose-aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(sounddose-aidl-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            sounddose-aidl-cpp
    )
endif ()
#################### sounddose-aidl-cpp-analyzer ####################

#################### audioflinger-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/CreateEffectRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/CreateEffectResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/CreateRecordRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/CreateRecordResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/CreateTrackRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/CreateTrackResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/OpenInputRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/OpenInputResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/OpenOutputRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/OpenOutputResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/RenderPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/IAudioFlingerService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/IAudioFlingerClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/IAudioRecord.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/IAudioTrack.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/IAudioTrackCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/ISoundDose.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/ISoundDoseCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/android/media/SoundDoseRecord.cpp
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
    add_library(audioflinger-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(audioflinger-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/effect-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioflinger-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audioflinger-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audioflinger-aidl-cpp
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            effect-aidl-cpp
            framework-permission-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
            shared-file-region-aidl-cpp
    )
endif ()
#################### audioflinger-aidl-cpp ####################

#################### spatializer-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/android/media/INativeSpatializerCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/android/media/ISpatializer.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/android/media/ISpatializerHeadTrackingCallback.cpp
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
    add_library(spatializer-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(spatializer-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(spatializer-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(spatializer-aidl-cpp
            audiopolicy-types-aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### spatializer-aidl-cpp ####################

#################### spatializer-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk_platform-source/gen/android/media/INativeSpatializerCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk_platform-source/gen/android/media/ISpatializer.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk_platform-source/gen/android/media/ISpatializerHeadTrackingCallback.cpp
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
    add_library(spatializer-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(spatializer-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(spatializer-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(spatializer-aidl-ndk_platform
            audiopolicy-types-aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### spatializer-aidl-ndk_platform ####################

#################### audioclient-types-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioAttributesInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioDirectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioGainSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioHalVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioHwModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioIoConfigEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioIoDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPatchFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPolicyConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPortFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPortSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPortConfigFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPortConfigSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPortDeviceExtSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPortExtSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPortMixExtSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPortRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioPortType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioProfileSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioRoute.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioTimestampInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioUniqueIdUse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/AudioVibratorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/DeviceConnectedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/EffectDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/TrackSecondaryOutputInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/android/media/SurroundSoundConfig.cpp
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
    add_library(audioclient-types-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(audioclient-types-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audioclient-types-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audioclient-types-aidl-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
            framework-permission-aidl-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### audioclient-types-aidl-ndk_platform ####################

#################### audiopolicy-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk-source/gen/android/media/GetInputForAttrResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk-source/gen/android/media/GetOutputForAttrResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk-source/gen/android/media/GetSpatializerResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk-source/gen/android/media/RecordClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk-source/gen/android/media/IAudioPolicyService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk-source/gen/android/media/IAudioPolicyServiceClient.cpp
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
    add_library(audiopolicy-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(audiopolicy-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-types-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/capture_state_listener-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/spatializer-aidl-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audiopolicy-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audiopolicy-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audiopolicy-aidl-ndk
            android.media.audio.common.types-V2-ndk
            audioclient-types-aidl-ndk
            audiopolicy-types-aidl-ndk
            capture_state_listener-aidl-ndk
            framework-permission-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            spatializer-aidl-ndk
    )
endif ()
#################### audiopolicy-aidl-ndk ####################

#################### audioclient-types-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioAttributesInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioDirectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioGainSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioHalVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioHwModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioIoConfigEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioIoDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPatchFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPolicyConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPortFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPortSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPortConfigFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPortConfigSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPortDeviceExtSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPortExtSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPortMixExtSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPortRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioPortType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioProfileSys.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioRoute.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioTimestampInternal.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioUniqueIdUse.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/AudioVibratorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/DeviceConnectedState.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/EffectDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/TrackSecondaryOutputInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/android/media/SurroundSoundConfig.cpp
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
    add_library(audioclient-types-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(audioclient-types-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libaudioclient/audioclient-types-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(audioclient-types-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(audioclient-types-aidl-ndk
            android.media.audio.common.types-V2-ndk
            framework-permission-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### audioclient-types-aidl-ndk ####################

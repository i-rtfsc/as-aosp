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

project(android.media.audio.common.types-cpp-analyzer)

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

#################### android.media.audio.common.types-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/android/media/audio/common/PcmType.cpp
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
    add_library(android.media.audio.common.types-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-V1-ndk_platform
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
#################### android.media.audio.common.types-V1-ndk_platform ####################

#################### android.media.soundtrigger.types-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/ConfidenceLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/ModelParameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/ModelParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/Phrase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/PhraseRecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/PhraseRecognitionExtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/PhraseSoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/Properties.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/RecognitionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/RecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/RecognitionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/RecognitionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/SoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/SoundModelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/android/media/soundtrigger/Status.cpp
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
    add_library(android.media.soundtrigger.types-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.media.soundtrigger.types-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.soundtrigger.types-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.soundtrigger.types-V2-ndk
            android.media.audio.common.types-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.media.soundtrigger.types-V2-ndk ####################

#################### android.media.soundtrigger.types-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/ConfidenceLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/ModelParameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/ModelParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/Phrase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/PhraseRecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/PhraseRecognitionExtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/PhraseSoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/Properties.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/RecognitionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/RecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/RecognitionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/RecognitionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/SoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/SoundModelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/android/media/soundtrigger/Status.cpp
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
    add_library(android.media.soundtrigger.types-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.media.soundtrigger.types-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.soundtrigger.types-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.soundtrigger.types-V1-cpp
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
    )
endif ()
#################### android.media.soundtrigger.types-V1-cpp ####################

#################### android.media.soundtrigger.types-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/ConfidenceLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/ModelParameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/ModelParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/Phrase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/PhraseRecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/PhraseRecognitionExtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/PhraseSoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/Properties.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/RecognitionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/RecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/RecognitionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/RecognitionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/SoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/SoundModelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/android/media/soundtrigger/Status.cpp
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
    add_library(android.media.soundtrigger.types-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.media.soundtrigger.types-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.soundtrigger.types-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.soundtrigger.types-V2-cpp
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
    )
endif ()
#################### android.media.soundtrigger.types-V2-cpp ####################

#################### android.media.soundtrigger.types-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/ConfidenceLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/ModelParameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/ModelParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/Phrase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/PhraseRecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/PhraseRecognitionExtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/PhraseSoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/Properties.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/RecognitionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/RecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/RecognitionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/RecognitionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/SoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/SoundModelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/android/media/soundtrigger/Status.cpp
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
    add_library(android.media.soundtrigger.types-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.media.soundtrigger.types-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.soundtrigger.types-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.soundtrigger.types-V2-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.media.soundtrigger.types-V2-ndk_platform ####################

#################### android.media.soundtrigger.types-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/ConfidenceLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/ModelParameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/ModelParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/Phrase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/PhraseRecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/PhraseRecognitionExtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/PhraseSoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/Properties.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/RecognitionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/RecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/RecognitionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/RecognitionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/SoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/SoundModelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/android/media/soundtrigger/Status.cpp
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
    add_library(android.media.soundtrigger.types-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.media.soundtrigger.types-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.soundtrigger.types-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.soundtrigger.types-V1-ndk
            android.media.audio.common.types-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.media.soundtrigger.types-V1-ndk ####################

#################### android.media.audio.common.types-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/android/media/audio/common/PcmType.cpp
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
    add_library(android.media.audio.common.types-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.media.audio.common.types-V1-ndk ####################

#################### android.media.audio.common.types-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioAttributes.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioDualMonoMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioHalAttributesGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioHalCapCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioHalCapCriterionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioHalEngineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioHalProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioHalVolumeCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioHalVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioLatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioPlaybackRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioProductStrategyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/Boolean.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/Byte.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/Float.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/Double.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/Long.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/MicrophoneDynamicInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/MicrophoneInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/PcmType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/android/media/audio/common/Void.cpp
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
    add_library(android.media.audio.common.types-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-V3-cpp
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
#################### android.media.audio.common.types-V3-cpp ####################

#################### android.media.audio.common.types-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioAttributes.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioDualMonoMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioHalAttributesGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioHalCapCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioHalCapCriterionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioHalEngineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioHalProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioHalVolumeCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioHalVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioLatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioPlaybackRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioProductStrategyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/Boolean.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/Byte.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/Double.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/Float.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/Long.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/MicrophoneDynamicInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/MicrophoneInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/PcmType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/android/media/audio/common/Void.cpp
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
    add_library(android.media.audio.common.types-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-V2-ndk_platform
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
#################### android.media.audio.common.types-V2-ndk_platform ####################

#################### android.media.audio.common.types-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioAttributes.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioDualMonoMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioHalAttributesGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioHalCapCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioHalCapCriterionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioHalEngineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioHalProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioHalVolumeCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioHalVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioLatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioPlaybackRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioProductStrategyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/Boolean.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/Byte.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/Float.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/Double.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/Long.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/MicrophoneDynamicInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/MicrophoneInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/PcmType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/android/media/audio/common/Void.cpp
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
    add_executable(android.media.audio.common.types-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/system/hardware/interfaces/media
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-cpp-analyzer
            aidl-analyzer-main
            android.media.audio.common.types-V3-cpp
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
#################### android.media.audio.common.types-cpp-analyzer ####################

#################### android.media.audio.common.types-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioAttributes.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioDualMonoMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioHalAttributesGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioHalCapCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioHalCapCriterionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioHalEngineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioHalProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioHalVolumeCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioHalVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioLatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioPlaybackRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioProductStrategyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/Boolean.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/Byte.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/Double.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/Float.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/Long.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/MicrophoneDynamicInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/MicrophoneInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/PcmType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/android/media/audio/common/Void.cpp
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
    add_library(android.media.audio.common.types-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-V2-cpp
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
#################### android.media.audio.common.types-V2-cpp ####################

#################### android.media.audio.common.types-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioAttributes.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioDualMonoMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioHalAttributesGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioHalCapCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioHalCapCriterionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioHalEngineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioHalProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioHalVolumeCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioHalVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioLatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioPlaybackRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioProductStrategyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/Boolean.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/Byte.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/Float.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/Double.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/Long.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/MicrophoneDynamicInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/MicrophoneInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/PcmType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/android/media/audio/common/Void.cpp
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
    add_library(android.media.audio.common.types-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-V3-ndk_platform
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
#################### android.media.audio.common.types-V3-ndk_platform ####################

#################### android.media.audio.common.types-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioAttributes.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioDualMonoMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioHalAttributesGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioHalCapCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioHalCapCriterionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioHalEngineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioHalProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioHalVolumeCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioHalVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioLatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioPlaybackRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioProductStrategyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/Boolean.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/Byte.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/Float.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/Double.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/Long.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/MicrophoneDynamicInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/MicrophoneInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/PcmType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/android/media/audio/common/Void.cpp
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
    add_library(android.media.audio.common.types-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.media.audio.common.types-V3-ndk ####################

#################### android.media.soundtrigger.types-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/AudioCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/ConfidenceLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/ModelParameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/ModelParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/Phrase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/PhraseRecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/PhraseRecognitionExtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/PhraseSoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/Properties.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/RecognitionConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/RecognitionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/RecognitionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/RecognitionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/SoundModel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/SoundModelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/android/media/soundtrigger/Status.cpp
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
    add_library(android.media.soundtrigger.types-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.media.soundtrigger.types-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.soundtrigger.types-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.soundtrigger.types-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.soundtrigger.types-V1-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.media.soundtrigger.types-V1-ndk_platform ####################

#################### android.media.audio.common.types-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/android/media/audio/common/PcmType.cpp
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
    add_library(android.media.audio.common.types-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-V1-cpp
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
#################### android.media.audio.common.types-V1-cpp ####################

#################### android.media.audio.common.types-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioAttributes.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioChannelLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioConfigBase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioContentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioDeviceAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioDeviceDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioDeviceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioDualMonoMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioEncapsulationMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioEncapsulationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioEncapsulationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioFormatDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioFormatType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioGain.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioGainConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioGainMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioHalAttributesGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioHalCapCriterion.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioHalCapCriterionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioHalEngineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioHalProductStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioHalVolumeCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioHalVolumeGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioInputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioIoFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioLatencyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioMMapPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioMMapPolicyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioMMapPolicyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioOffloadInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioOutputFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioPlaybackRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioPort.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioPortConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioPortDeviceExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioPortExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioPortMixExt.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioPortMixExtUseCase.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioProductStrategyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioProfile.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/AudioUuid.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/Boolean.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/Byte.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/Double.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/ExtraAudioDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/Float.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/Int.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/Long.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/MicrophoneDynamicInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/MicrophoneInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/PcmType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/android/media/audio/common/Void.cpp
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
    add_library(android.media.audio.common.types-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.media.audio.common.types-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/hardware/interfaces/media/android.media.audio.common.types-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.media.audio.common.types-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.media.audio.common.types-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.media.audio.common.types-V2-ndk ####################

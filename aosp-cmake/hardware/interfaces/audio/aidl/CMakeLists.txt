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

project(android.hardware.audio.common-V2-ndk_platform)

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

#################### android.hardware.audio.common-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/android/hardware/audio/common/AudioOffloadMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/android/hardware/audio/common/PlaybackTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/android/hardware/audio/common/RecordTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/android/hardware/audio/common/SinkMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/android/hardware/audio/common/SourceMetadata.cpp
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
    add_library(android.hardware.audio.common-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.common-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.common-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.common-V2-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.hardware.audio.common-V2-ndk_platform ####################

#################### android.hardware.audio.effect-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/AcousticEchoCanceler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/AutomaticGainControlV1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/AutomaticGainControlV2.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/BassBoost.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Capability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/CommandId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/DefaultExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Descriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Downmix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/DynamicsProcessing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/EnvironmentalReverb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Equalizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Flags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/HapticGenerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/IEffect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/IFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/LoudnessEnhancer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/NoiseSuppression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Parameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/PresetReverb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Processing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Range.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/State.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/VendorExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Virtualizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Visualizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/android/hardware/audio/effect/Volume.cpp
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
    add_library(android.hardware.audio.effect-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.effect-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.effect-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.effect-V2-ndk
            android.hardware.audio.common-V2-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-ndk
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
#################### android.hardware.audio.effect-V2-ndk ####################

#################### android.hardware.audio.core-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/AudioPatch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/AudioRoute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IBluetooth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IBluetoothA2dp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IBluetoothLe.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IStreamCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IStreamCommon.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IStreamIn.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IStreamOut.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/IStreamOutEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/ITelephony.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/MmapBufferDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/ModuleDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/StreamDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/SurroundSoundConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/android/hardware/audio/core/VendorParameter.cpp
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
    add_library(android.hardware.audio.core-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.core-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.core-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.core-V1-ndk
            android.hardware.audio.common-V2-ndk
            android.hardware.audio.core.sounddose-V1-ndk
            android.hardware.audio.effect-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-ndk
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
#################### android.hardware.audio.core-V1-ndk ####################

#################### android.hardware.audio.effect-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/AcousticEchoCanceler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/AutomaticGainControlV1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/AutomaticGainControlV2.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/BassBoost.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Capability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/CommandId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/DefaultExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Descriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Downmix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/DynamicsProcessing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/EnvironmentalReverb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Equalizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Flags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/HapticGenerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/IEffect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/IFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/LoudnessEnhancer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/NoiseSuppression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Parameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/PresetReverb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Processing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Range.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/State.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/VendorExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Virtualizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Visualizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/android/hardware/audio/effect/Volume.cpp
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
    add_library(android.hardware.audio.effect-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.effect-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.effect-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.effect-V1-ndk
            android.hardware.audio.common-V2-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-ndk
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
#################### android.hardware.audio.effect-V1-ndk ####################

#################### android.hardware.audio.core-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/AudioPatch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/AudioRoute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IBluetooth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IBluetoothA2dp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IBluetoothLe.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IStreamCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IStreamCommon.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IStreamIn.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IStreamOut.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/IStreamOutEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/ITelephony.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/MmapBufferDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/ModuleDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/StreamDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/SurroundSoundConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/android/hardware/audio/core/VendorParameter.cpp
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
    add_library(android.hardware.audio.core-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.core-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.core-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.core-V1-ndk_platform
            android.hardware.audio.common-V2-ndk_platform
            android.hardware.audio.core.sounddose-V1-ndk_platform
            android.hardware.audio.effect-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.hardware.audio.core-V1-ndk_platform ####################

#################### android.hardware.audio.core-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/AudioPatch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/AudioRoute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IBluetooth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IBluetoothA2dp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IBluetoothLe.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IStreamCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IStreamCommon.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IStreamIn.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IStreamOut.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/IStreamOutEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/ITelephony.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/MmapBufferDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/ModuleDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/StreamDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/SurroundSoundConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/android/hardware/audio/core/VendorParameter.cpp
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
    add_library(android.hardware.audio.core-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.core-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.core-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.core-V2-ndk
            android.hardware.audio.common-V2-ndk
            android.hardware.audio.core.sounddose-V1-ndk
            android.hardware.audio.effect-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-ndk
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
#################### android.hardware.audio.core-V2-ndk ####################

#################### android.hardware.audio.core.sounddose-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V2-ndk_platform-source/gen/android/hardware/audio/core/sounddose/ISoundDose.cpp
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
    add_library(android.hardware.audio.core.sounddose-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.core.sounddose-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.core.sounddose-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.core.sounddose-V2-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.hardware.audio.core.sounddose-V2-ndk_platform ####################

#################### android.hardware.audio.effect-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/AcousticEchoCanceler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/AutomaticGainControlV1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/AutomaticGainControlV2.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/BassBoost.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Capability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/CommandId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/DefaultExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Descriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Downmix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/DynamicsProcessing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/EnvironmentalReverb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Equalizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Flags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/HapticGenerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/IEffect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/IFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/LoudnessEnhancer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/NoiseSuppression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Parameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/PresetReverb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Processing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Range.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/State.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/VendorExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Virtualizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Visualizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/android/hardware/audio/effect/Volume.cpp
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
    add_library(android.hardware.audio.effect-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.effect-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.effect-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.effect-V1-ndk_platform
            android.hardware.audio.common-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.hardware.audio.effect-V1-ndk_platform ####################

#################### android.hardware.audio.common-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk_platform-source/gen/android/hardware/audio/common/AudioOffloadMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk_platform-source/gen/android/hardware/audio/common/PlaybackTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk_platform-source/gen/android/hardware/audio/common/RecordTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk_platform-source/gen/android/hardware/audio/common/SinkMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk_platform-source/gen/android/hardware/audio/common/SourceMetadata.cpp
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
    add_library(android.hardware.audio.common-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.common-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.common-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.common-V3-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.hardware.audio.common-V3-ndk_platform ####################

#################### android.hardware.audio.core.sounddose-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V2-ndk-source/gen/android/hardware/audio/core/sounddose/ISoundDose.cpp
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
    add_library(android.hardware.audio.core.sounddose-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.core.sounddose-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.core.sounddose-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.core.sounddose-V2-ndk
            android.media.audio.common.types-V2-ndk
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
#################### android.hardware.audio.core.sounddose-V2-ndk ####################

#################### android.hardware.audio.common-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-cpp-source/gen/android/hardware/audio/common/AudioOffloadMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-cpp-source/gen/android/hardware/audio/common/PlaybackTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-cpp-source/gen/android/hardware/audio/common/RecordTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-cpp-source/gen/android/hardware/audio/common/SinkMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-cpp-source/gen/android/hardware/audio/common/SourceMetadata.cpp
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
    add_library(android.hardware.audio.common-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.common-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.common-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.common-V2-cpp
            android.media.audio.common.types-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.audio.common-V2-cpp ####################

#################### android.hardware.audio.common-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk-source/gen/android/hardware/audio/common/PlaybackTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk-source/gen/android/hardware/audio/common/RecordTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk-source/gen/android/hardware/audio/common/SinkMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk-source/gen/android/hardware/audio/common/SourceMetadata.cpp
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
    add_library(android.hardware.audio.common-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.common-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.common-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.common-V1-ndk
            android.media.audio.common.types-V2-ndk
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
#################### android.hardware.audio.common-V1-ndk ####################

#################### android.hardware.audio.common-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-cpp-source/gen/android/hardware/audio/common/AudioOffloadMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-cpp-source/gen/android/hardware/audio/common/PlaybackTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-cpp-source/gen/android/hardware/audio/common/RecordTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-cpp-source/gen/android/hardware/audio/common/SinkMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-cpp-source/gen/android/hardware/audio/common/SourceMetadata.cpp
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
    add_library(android.hardware.audio.common-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.common-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.common-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.common-V3-cpp
            android.media.audio.common.types-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.audio.common-V3-cpp ####################

#################### android.hardware.audio.common-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk_platform-source/gen/android/hardware/audio/common/PlaybackTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk_platform-source/gen/android/hardware/audio/common/RecordTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk_platform-source/gen/android/hardware/audio/common/SinkMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk_platform-source/gen/android/hardware/audio/common/SourceMetadata.cpp
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
    add_library(android.hardware.audio.common-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.common-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.common-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.common-V1-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.hardware.audio.common-V1-ndk_platform ####################

#################### android.hardware.audio.effect-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/AcousticEchoCanceler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/AutomaticGainControlV1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/AutomaticGainControlV2.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/BassBoost.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Capability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/CommandId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/DefaultExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Descriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Downmix.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/DynamicsProcessing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/EnvironmentalReverb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Equalizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Flags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/HapticGenerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/IEffect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/IFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/LoudnessEnhancer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/NoiseSuppression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Parameter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/PresetReverb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Processing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Range.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/State.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/VendorExtension.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Virtualizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Visualizer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/android/hardware/audio/effect/Volume.cpp
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
    add_library(android.hardware.audio.effect-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.effect-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.effect-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.effect-V2-ndk_platform
            android.hardware.audio.common-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.hardware.audio.effect-V2-ndk_platform ####################

#################### android.hardware.audio.common-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/android/hardware/audio/common/AudioOffloadMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/android/hardware/audio/common/PlaybackTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/android/hardware/audio/common/RecordTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/android/hardware/audio/common/SinkMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/android/hardware/audio/common/SourceMetadata.cpp
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
    add_library(android.hardware.audio.common-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.common-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.common-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.common-V2-ndk
            android.media.audio.common.types-V2-ndk
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
#################### android.hardware.audio.common-V2-ndk ####################

#################### android.hardware.audio.core.sounddose-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V1-ndk-source/gen/android/hardware/audio/core/sounddose/ISoundDose.cpp
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
    add_library(android.hardware.audio.core.sounddose-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.core.sounddose-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.core.sounddose-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.core.sounddose-V1-ndk
            android.media.audio.common.types-V2-ndk
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
#################### android.hardware.audio.core.sounddose-V1-ndk ####################

#################### android.hardware.audio.common-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk-source/gen/android/hardware/audio/common/AudioOffloadMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk-source/gen/android/hardware/audio/common/PlaybackTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk-source/gen/android/hardware/audio/common/RecordTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk-source/gen/android/hardware/audio/common/SinkMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk-source/gen/android/hardware/audio/common/SourceMetadata.cpp
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
    add_library(android.hardware.audio.common-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.common-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.common-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.common-V3-ndk
            android.media.audio.common.types-V2-ndk
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
#################### android.hardware.audio.common-V3-ndk ####################

#################### android.hardware.audio.core.sounddose-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V1-ndk_platform-source/gen/android/hardware/audio/core/sounddose/ISoundDose.cpp
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
    add_library(android.hardware.audio.core.sounddose-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.core.sounddose-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.core.sounddose-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.core.sounddose-V1-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.hardware.audio.core.sounddose-V1-ndk_platform ####################

#################### android.hardware.audio.core-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/AudioPatch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/AudioRoute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IBluetooth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IBluetoothA2dp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IBluetoothLe.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IModule.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IStreamCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IStreamCommon.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IStreamIn.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IStreamOut.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/IStreamOutEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/ITelephony.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/MmapBufferDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/ModuleDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/StreamDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/SurroundSoundConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/android/hardware/audio/core/VendorParameter.cpp
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
    add_library(android.hardware.audio.core-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.core-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V2-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core.sounddose-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.effect-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.core-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.core-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.core-V2-ndk_platform
            android.hardware.audio.common-V2-ndk_platform
            android.hardware.audio.core.sounddose-V1-ndk_platform
            android.hardware.audio.effect-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### android.hardware.audio.core-V2-ndk_platform ####################

#################### android.hardware.audio.common-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-cpp-source/gen/android/hardware/audio/common/PlaybackTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-cpp-source/gen/android/hardware/audio/common/RecordTrackMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-cpp-source/gen/android/hardware/audio/common/SinkMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-cpp-source/gen/android/hardware/audio/common/SourceMetadata.cpp
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
    add_library(android.hardware.audio.common-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.common-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/audio/aidl/android.hardware.audio.common-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.common-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.audio.common-V1-cpp
            android.media.audio.common.types-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.audio.common-V1-cpp ####################

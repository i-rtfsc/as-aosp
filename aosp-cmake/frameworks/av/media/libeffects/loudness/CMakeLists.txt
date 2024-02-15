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

project(libloudnessenhanceraidl)

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

#################### libloudnessenhanceraidl ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/loudness/aidl/EffectLoudnessEnhancer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/loudness/aidl/LoudnessEnhancerContext.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/loudness/dsp/core/dynamic_range_compression.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/EffectThread.cpp
        ${ANDROID_ROOT}/hardware/interfaces/audio/aidl/default/EffectImpl.cpp
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
    add_library(libloudnessenhanceraidl SHARED ${SOURCE_FILES})

    target_include_directories(libloudnessenhanceraidl
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/loudness
    )

    # Add c/c++ flags
    target_compile_options(libloudnessenhanceraidl PRIVATE
            -Werror
    )

    target_link_libraries(libloudnessenhanceraidl
            android.hardware.audio.effect-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-ndk
            libalsautilsv2
            libaudio_system_headers
            libaudioaidl_headers
            libaudioaidlcommon
            libaudioeffects
            libaudioutils
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libfmq
            libhardware_headers
            liblog
            libm
            libstagefright_foundation
            libtinyalsav2
            libutils
            libxml2
            libxsdc-utils
            libc
            libdl
            libm
    )
endif ()
#################### libloudnessenhanceraidl ####################

#################### libldnhncr ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/loudness/EffectLoudnessEnhancer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/loudness/dsp/core/dynamic_range_compression.cpp
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
    add_library(libldnhncr SHARED ${SOURCE_FILES})

    target_include_directories(libldnhncr
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/loudness
    )

    # Add c/c++ flags
    target_compile_options(libldnhncr PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libldnhncr
            libaudioeffects
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libldnhncr ####################

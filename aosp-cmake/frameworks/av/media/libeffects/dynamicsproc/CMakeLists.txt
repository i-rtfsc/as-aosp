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

project(libdynproc)

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

#################### libdynproc ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/dynamicsproc/dsp/DPBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/dynamicsproc/dsp/DPFrequency.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/dynamicsproc/EffectDynamicsProcessing.cpp
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
    add_library(libdynproc SHARED ${SOURCE_FILES})

    target_include_directories(libdynproc
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/dynamicsproc
    )

    # Add c/c++ flags
    target_compile_options(libdynproc PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libdynproc
            libaudioeffects
            libaudioutils
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libeigen
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libdynproc ####################

#################### libdynamicsprocessingaidl ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/dynamicsproc/dsp/DPBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/dynamicsproc/dsp/DPFrequency.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/dynamicsproc/aidl/DynamicsProcessing.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libeffects/dynamicsproc/aidl/DynamicsProcessingContext.cpp
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
    add_library(libdynamicsprocessingaidl SHARED ${SOURCE_FILES})

    target_include_directories(libdynamicsprocessingaidl
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libeffects/dynamicsproc
    )

    # Add c/c++ flags
    target_compile_options(libdynamicsprocessingaidl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libdynamicsprocessingaidl
            android.hardware.audio.effect-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.media.audio.common.types-V2-ndk
            libalsautilsv2
            libaudio_system_headers
            libaudioaidl_headers
            libaudioaidlcommon
            libaudioaidlranges
            libaudioeffects
            libaudioutils
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libeigen
            libfmq
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
#################### libdynamicsprocessingaidl ####################

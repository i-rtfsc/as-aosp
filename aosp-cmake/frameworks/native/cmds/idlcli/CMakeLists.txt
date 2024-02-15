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

project(idlcli)

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

#################### idlcli ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/main.cpp
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
    add_executable(idlcli ${SOURCE_FILES})

    target_include_directories(idlcli
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/idlcli
    )

    # Add c/c++ flags
    target_compile_options(idlcli PRIVATE
            -Werror
    )

    target_link_libraries(idlcli
            android.hardware.vibrator-V2-ndk
            android.hardware.vibrator-1.0
            android.hardware.vibrator-1.1
            android.hardware.vibrator-1.2
            android.hardware.vibrator-1.3
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libhidlbase
            libidlcli
            liblog
            libm
            libutils
    )
endif ()
#################### idlcli ####################

#################### libidlcli ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/CommandVibrator.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandAlwaysOnDisable.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandAlwaysOnEnable.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandCompose.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandComposePwle.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetBandwidthAmplitudeMap.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetCapabilities.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetCompositionDelayMax.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetCompositionSizeMax.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetFrequencyMinimum.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetFrequencyResolution.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetPrimitiveDuration.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetPwleCompositionSizeMax.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetPwlePrimitiveDurationMax.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetQFactor.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetResonantFrequency.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetSupportedAlwaysOnEffects.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetSupportedBraking.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetSupportedEffects.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandGetSupportedPrimitives.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandOff.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandOn.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandPerform.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandSetAmplitude.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandSetExternalControl.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandSupportsAmplitudeControl.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/idlcli/vibrator/CommandSupportsExternalControl.cpp
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
    add_library(libidlcli SHARED ${SOURCE_FILES})

    target_include_directories(libidlcli
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/idlcli
    )

    # Add c/c++ flags
    target_compile_options(libidlcli PRIVATE
            -Werror
    )

    target_link_libraries(libidlcli
            android.hardware.vibrator-V2-ndk
            android.hardware.vibrator-1.0
            android.hardware.vibrator-1.1
            android.hardware.vibrator-1.2
            android.hardware.vibrator-1.3
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### libidlcli ####################

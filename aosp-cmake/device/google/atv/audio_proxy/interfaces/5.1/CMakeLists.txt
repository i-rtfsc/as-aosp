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

project(device.google.atv.audio_proxy@5.1)

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

#################### device.google.atv.audio_proxy-5.1 ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/device/google/atv/audio_proxy/interfaces/5.1/device.google.atv.audio_proxy@5.1_genc++/gen/device/google/atv/audio_proxy/5.1/AudioProxyDevicesManagerAll.cpp
        ${OUT_INTERMEDIATES_ROOT}/device/google/atv/audio_proxy/interfaces/5.1/device.google.atv.audio_proxy@5.1_genc++/gen/device/google/atv/audio_proxy/5.1/AudioProxyStreamOutAll.cpp
        ${OUT_INTERMEDIATES_ROOT}/device/google/atv/audio_proxy/interfaces/5.1/device.google.atv.audio_proxy@5.1_genc++/gen/device/google/atv/audio_proxy/5.1/BusDeviceAll.cpp
        ${OUT_INTERMEDIATES_ROOT}/device/google/atv/audio_proxy/interfaces/5.1/device.google.atv.audio_proxy@5.1_genc++/gen/device/google/atv/audio_proxy/5.1/StreamEventListenerAll.cpp
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
    add_library(device.google.atv.audio_proxy-5.1 SHARED ${SOURCE_FILES})

    target_include_directories(device.google.atv.audio_proxy-5.1
            PUBLIC
            ${ANDROID_ROOT}/device/google/atv/audio_proxy/interfaces/5.1
            ${OUT_INTERMEDIATES_ROOT}/device/google/atv/audio_proxy/interfaces/5.1/device.google.atv.audio_proxy@5.1_genc++/gen
            ${OUT_INTERMEDIATES_ROOT}/device/google/atv/audio_proxy/interfaces/5.1/device.google.atv.audio_proxy@5.1_genc++_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(device.google.atv.audio_proxy-5.1 PRIVATE
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(device.google.atv.audio_proxy-5.1
            android.hardware.audio.common-5.0
            android.hardware.audio-5.0
            android.hidl.safe_union-1.0
            device.google.atv.audio_proxy-5.0
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### device.google.atv.audio_proxy-5.1 ####################

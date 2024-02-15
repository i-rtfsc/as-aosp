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

project(libbluetooth_audio_session_aidl)

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

#################### libbluetooth_audio_session_aidl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/aidl_session/BluetoothAudioCodecs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/aidl_session/BluetoothAudioSession.cpp
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/aidl_session/HidlToAidlMiddleware.cpp
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/aidl_session/BluetoothLeAudioCodecsProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/utils/le_audio_codec_capabilities/gen/cpp/aidl_android_hardware_bluetooth_audio_setting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/utils/le_audio_codec_capabilities/gen/cpp/aidl_android_hardware_bluetooth_audio_setting_enums.cpp
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
    add_library(libbluetooth_audio_session_aidl SHARED ${SOURCE_FILES})

    target_include_directories(libbluetooth_audio_session_aidl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/aidl_session
            ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/bluetooth/audio/utils/le_audio_codec_capabilities/gen/cpp/include
    )

    # Add c/c++ flags
    target_compile_options(libbluetooth_audio_session_aidl PRIVATE
            -Werror
    )

    target_link_libraries(libbluetooth_audio_session_aidl
            android.hardware.bluetooth.audio-V3-ndk
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware_headers
            libhidlbase
            liblog
            libm
            libxml2
            libxsdc-utils
            libc
            libdl
            libm
    )
endif ()
#################### libbluetooth_audio_session_aidl ####################

#################### libbluetooth_audio_session ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/session/BluetoothAudioSession.cpp
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/session/BluetoothAudioSession_2_1.cpp
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/session/BluetoothAudioSupportedCodecsDB.cpp
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/session/BluetoothAudioSupportedCodecsDB_2_1.cpp
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
    add_library(libbluetooth_audio_session SHARED ${SOURCE_FILES})

    target_include_directories(libbluetooth_audio_session
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/session
            ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils
            ${ANDROID_ROOT}/hardware/interfaces/bluetooth/audio/utils/aidl_session
    )

    # Add c/c++ flags
    target_compile_options(libbluetooth_audio_session PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libbluetooth_audio_session
            android.hardware.audio.common-5.0
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            libbase
            libbluetooth_audio_session_aidl
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware_headers
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libbluetooth_audio_session ####################

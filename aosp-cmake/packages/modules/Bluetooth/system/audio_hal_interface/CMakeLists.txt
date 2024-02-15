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

project(libbt-audio-hal-interface)

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

#################### libbt-audio-hal-interface ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/a2dp_encoding.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/aidl/a2dp_encoding_aidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/aidl/bluetooth_audio_port_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/aidl/client_interface_aidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/aidl/codec_status_aidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/aidl/hearing_aid_software_encoding_aidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/aidl/le_audio_software_aidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/hal_version_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/hearing_aid_software_encoding.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/hidl/a2dp_encoding_hidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/hidl/client_interface_hidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/hidl/codec_status_hidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/hidl/hearing_aid_software_encoding_hidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/hidl/le_audio_software_hidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface/le_audio_software.cc
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
    add_library(libbt-audio-hal-interface SHARED ${SOURCE_FILES})

    target_include_directories(libbt-audio-hal-interface
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/audio_hal_interface
    )

    # Add c/c++ flags
    target_compile_options(libbt-audio-hal-interface PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-audio-hal-interface
            android.hardware.bluetooth.audio-V3-ndk
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            libbinder_ndk
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-common
            libbt-platform-protos-lite
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhidlbase
            liblog
            libm
            libosi
            libstatslog_bt
            libstatssocket
            libutils
    )
endif ()
#################### libbt-audio-hal-interface ####################

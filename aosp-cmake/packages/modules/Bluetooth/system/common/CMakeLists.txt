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

project(libbt-common)

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

#################### libbt-common ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/common/address_obfuscator.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/common/message_loop_thread.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/common/metric_id_allocator.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/common/os_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/common/repeating_timer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/common/stop_watch_legacy.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/common/time_util.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/common/metrics.cc
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
    add_library(libbt-common SHARED ${SOURCE_FILES})

    target_include_directories(libbt-common
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/common
    )

    # Add c/c++ flags
    target_compile_options(libbt-common PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-common
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            libFraunhoferAAC
            libPlatformProperties
            libaaudio
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-platform-protos-lite
            libbt-protos-lite
            libbt-sbc-decoder
            libbt-sbc-encoder
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libflatbuffers-cpp
            libfmq
            libg722codec
            libhidlbase
            liblc3
            liblog
            libm
            libnativehelper
            libopus
            libprotobuf-cpp-lite
            libstatslog_bt
            libstatssocket
            libudrv-uipc
            libutils
    )
endif ()
#################### libbt-common ####################

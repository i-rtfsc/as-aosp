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

project(libbluetooth_jni)

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

#################### libbluetooth_jni ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_BluetoothQualityReport.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_a2dp.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_a2dp_sink.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_avrcp_controller.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_avrcp_target.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_btservice_ActivityAttribution.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_btservice_AdapterService.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_btservice_BluetoothKeystore.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_csip_set_coordinator.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_gatt.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_hap_client.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_hearing_aid.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_hfp.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_hfpclient.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_hid_device.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_hid_host.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_le_audio.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_pan.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_sdp.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app/jni/com_android_bluetooth_vc.cpp
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
    add_library(libbluetooth_jni SHARED ${SOURCE_FILES})

    target_include_directories(libbluetooth_jni
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/android/app
    )

    # Add c/c++ flags
    target_compile_options(libbluetooth_jni PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
    )

    target_link_libraries(libbluetooth_jni
            jni_headers
            libbase
            libbluetooth-types
            libbluetooth_headers
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativehelper
    )
endif ()
#################### libbluetooth_jni ####################

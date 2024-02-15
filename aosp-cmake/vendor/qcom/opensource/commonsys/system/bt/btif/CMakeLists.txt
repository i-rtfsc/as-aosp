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

project(libbtif_qti)

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

#################### libbtif_qti ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/bluetooth.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/bluetooth_data_migration.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_a2dp.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_a2dp_control.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_a2dp_sink.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_a2dp_source.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_a2dp_audio_interface.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_ahim.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_av.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_avrcp_audio_track.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_ble_advertiser.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_ble_scanner.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_bqr.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_config.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_config_cache.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_config_transcode.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_core.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_debug.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_debug_btsnoop.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_debug_conn.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_dm.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_gatt.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_gatt_client.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_gatt_server.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_gatt_test.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_gatt_util.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_hearing_aid.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_has_client.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_hf.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_hf_client.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_hh.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_hd.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_pan.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_profile_queue.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_rc.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sdp.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sdp_server.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sm.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sock.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sock_rfc.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sock_l2cap.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sock_sco.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sock_sdp.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sock_thread.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_sock_util.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_storage.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_uid.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_util.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_keystore.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/btif_atp_locator.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/stack_manager.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/src/stack_interface.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/co/bta_ag_co.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/co/bta_dm_co.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/co/bta_av_co.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/co/bta_hh_co.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/co/bta_pan_co.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/co/bta_gatts_co.cc
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
    add_library(libbtif_qti SHARED ${SOURCE_FILES})

    target_include_directories(libbtif_qti
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/include
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif/co
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/btif
    )

    # Add c/c++ flags
    target_compile_options(libbtif_qti PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbtif_qti
            android.hardware.audio.common-V1-ndk
            android.hardware.bluetooth.audio-V3-ndk
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            com.qualcomm.qti.bluetooth_audio-1.0
            libaaudio
            libaudio-a2dp-hw-utils
            libbluetooth-types
            libbluetooth_headers
            libbt-audio-hal-interface-qti
            libbt-common-qti
            libc
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libfmq
            libhidlbase
            liblog
            libm
            libmedia_headers
            libmediametrics_headers
            libtinyxml2
            libutils
            libz
            vendor.qti.hardware.bluetooth.audio-V1-ndk
            vendor.qti.hardware.bluetooth_audio-2.0
            vendor.qti.hardware.bluetooth_audio-2.1
    )
endif ()
#################### libbtif_qti ####################

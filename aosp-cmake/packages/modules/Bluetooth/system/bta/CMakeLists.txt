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

project(libbt-bta)

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

#################### libbt-bta ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ag/bta_ag_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ag/bta_ag_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ag/bta_ag_at.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ag/bta_ag_cfg.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ag/bta_ag_cmd.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ag/bta_ag_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ag/bta_ag_rfc.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ag/bta_ag_sco.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ag/bta_ag_sdp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/ar/bta_ar.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/av/bta_av_aact.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/av/bta_av_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/av/bta_av_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/av/bta_av_cfg.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/av/bta_av_ci.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/av/bta_av_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/av/bta_av_ssm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/csis/csis_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/groups/groups.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/has/has_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/has/has_ctp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/has/has_journal.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/has/has_preset.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/has/has_types.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hd/bta_hd_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hd/bta_hd_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hd/bta_hd_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hearing_aid/hearing_aid.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hearing_aid/hearing_aid_audio_source.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hf_client/bta_hf_client_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hf_client/bta_hf_client_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hf_client/bta_hf_client_at.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hf_client/bta_hf_client_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hf_client/bta_hf_client_rfc.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hf_client/bta_hf_client_sco.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hf_client/bta_hf_client_sdp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hfp/bta_hfp_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hh/bta_hh_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hh/bta_hh_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hh/bta_hh_cfg.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hh/bta_hh_le.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hh/bta_hh_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hh/bta_hh_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/audio_hal_client/audio_sink_hal_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/audio_hal_client/audio_source_hal_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/broadcaster/broadcaster.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/broadcaster/broadcaster_types.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/broadcaster/state_machine.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/client_parser.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/codec_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/content_control_id_keeper.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/devices.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/hal_verifier.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/le_audio_log_history.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/le_audio_set_configuration_provider.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/le_audio_set_configuration_provider_json.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/le_audio_types.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/le_audio_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/metrics_collector.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/state_machine.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/le_audio/storage_helper.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/pan/bta_pan_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/pan/bta_pan_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/pan/bta_pan_ci.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/pan/bta_pan_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/vc/device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/vc/vc.cc
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
    add_library(libbt-bta SHARED ${SOURCE_FILES})

    target_include_directories(libbt-bta
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/dm
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hd
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hh
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/sys
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/bta/LeAudioSetConfigSchemas_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libbt-bta PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-bta
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            avrcp-target-service
            lib-bt-packets
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
            libbt-bta-core
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
            libcom.android.sysprop.bluetooth
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
#################### libbt-bta ####################

#################### libbt-bta-core ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/dm/bta_dm_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/dm/bta_dm_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/dm/bta_dm_cfg.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/dm/bta_dm_ci.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/dm/bta_dm_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/dm/bta_dm_pm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gattc_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gattc_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gattc_cache.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gattc_db_storage.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gattc_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gattc_queue.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gattc_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gatts_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gatts_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gatts_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gatts_queue.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/bta_gatts_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/database.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/gatt/database_builder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/jv/bta_jv_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/jv/bta_jv_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/jv/bta_jv_cfg.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/sdp/bta_sdp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/sdp/bta_sdp_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/sdp/bta_sdp_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/sdp/bta_sdp_cfg.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/sys/bta_sys_conn.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/sys/bta_sys_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/sys/utl.cc
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
    add_library(libbt-bta-core SHARED ${SOURCE_FILES})

    target_include_directories(libbt-bta-core
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/dm
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hd
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/hh
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta/sys
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/bta
    )

    # Add c/c++ flags
    target_compile_options(libbt-bta-core PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-bta-core
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            lib-bt-packets
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
            libcom.android.sysprop.bluetooth
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
#################### libbt-bta-core ####################

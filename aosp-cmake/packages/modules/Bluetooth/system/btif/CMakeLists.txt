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

project(libbtif-core)

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

#################### libbtif-core ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/co/bta_dm_co.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/co/bta_gatts_co.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_activity_attribution.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_ble_advertiser.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_ble_scanner.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_bqr.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_config.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_config_cache.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_core.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_debug_conn.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_dm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_gatt.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_gatt_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_gatt_server.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_gatt_test.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_gatt_util.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_iot_config.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_keystore.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_metrics_logging.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_profile_queue.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_sdp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_sdp_server.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_sock.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_sock_l2cap.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_sock_rfc.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_sock_sco.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_sock_sdp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_sock_thread.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_sock_util.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_storage.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_uid.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_util.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/stack_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_avrcp_audio_track.cc
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
    add_library(libbtif-core SHARED ${SOURCE_FILES})

    target_include_directories(libbtif-core
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/avrcp
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/co
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/btif/statslog_bt.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libbtif-core PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
    )

    target_link_libraries(libbtif-core
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            lib-bt-packets
            lib-bt-packets-base
            libFraunhoferAAC
            libPlatformProperties
            libaaudio
            libaudio-a2dp-hw-utils
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-audio-hal-interface
            libbt-platform-protos-lite
            libbt-protos-lite
            libbt-sbc-decoder
            libbt-sbc-encoder
            libbt-stack-core
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
#################### libbtif-core ####################

#################### libbtif ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/avrcp/avrcp_service.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/co/bta_av_co.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/co/bta_hh_co.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/co/bta_pan_co.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/bluetooth.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_a2dp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_a2dp_control.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_a2dp_sink.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_a2dp_source.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_av.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_csis_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_has_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_hd.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_hearing_aid.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_hf.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_hf_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_hh.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_iot_config.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_le_audio.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_le_audio_broadcaster.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_pan.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_profile_queue.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_profile_storage.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_rc.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_vc.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/profile_log_levels.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/src/btif_avrcp_audio_track.cc
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
    add_library(libbtif SHARED ${SOURCE_FILES})

    target_include_directories(libbtif
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/avrcp
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif/co
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/btif/statslog_bt.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libbtif PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
    )

    target_link_libraries(libbtif
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            avrcp-target-service
            lib-bt-packets
            lib-bt-packets-avrcp
            lib-bt-packets-base
            libFraunhoferAAC
            libPlatformProperties
            libaaudio
            libaudio-a2dp-hw-utils
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-audio-hal-interface
            libbt-platform-protos-lite
            libbt-protos-lite
            libbt-sbc-decoder
            libbt-sbc-encoder
            libbt-stack
            libbt-stack-core
            libbt_shim_bridge
            libbt_shim_ffi
            libbtif-core
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
#################### libbtif ####################

#################### libstatslog_bt ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/btif/statslog_bt.cpp/gen/statslog_bt.cpp
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
    add_library(libstatslog_bt SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_bt
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btif
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/btif/statslog_bt.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/btif/statslog_bt.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libstatslog_bt PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libstatslog_bt
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libstatssocket
    )
endif ()
#################### libstatslog_bt ####################

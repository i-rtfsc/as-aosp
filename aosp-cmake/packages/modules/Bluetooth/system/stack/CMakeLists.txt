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

project(libbt-stack-core)

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

#################### libbt-stack-core ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/crypto_toolbox/aes.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/crypto_toolbox/aes_cmac.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/crypto_toolbox/crypto_toolbox.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/sdp/sdp_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/sdp/sdp_db.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/sdp/sdp_discovery.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/sdp/sdp_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/sdp/sdp_server.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/sdp/sdp_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/acl/acl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/acl/ble_acl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/acl/btm_acl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/acl/btm_ble_connection_establishment.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/acl/btm_pm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/arbiter/acl_arbiter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/ble_advertiser_hci_interface.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/ble_scanner_hci_interface.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble_addr.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble_adv_filter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble_batchscan.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble_bgconn.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble_cont_energy.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble_gap.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble_multi_adv.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble_privacy.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_ble_scanner.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_client_interface.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_dev.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_devctl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_inq.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_iot_config.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_iso.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_scn.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_sco.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_sco_hci.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_sco_hfp_hal.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/btm_sec.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btu/btu_hcif.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btu/btu_task.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/eatt/eatt.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gap/gap_ble.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gap/gap_conn.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/att_protocol.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/connection_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/gatt_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/gatt_attr.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/gatt_auth.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/gatt_cl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/gatt_db.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/gatt_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/gatt_sr.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/gatt_sr_hash.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt/gatt_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/hcic/hciblecmds.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/hcic/hcicmds.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/l2cap/l2c_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/l2cap/l2c_ble.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/l2cap/l2c_csm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/l2cap/l2c_fcr.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/l2cap/l2c_link.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/l2cap/l2c_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/l2cap/l2c_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/metrics/stack_metrics_logging.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/rfcomm/port_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/rfcomm/port_rfc.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/rfcomm/port_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/rfcomm/rfc_l2cap_if.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/rfcomm/rfc_mx_fsm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/rfcomm/rfc_port_fsm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/rfcomm/rfc_port_if.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/rfcomm/rfc_ts_frames.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/rfcomm/rfc_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/p_256_curvepara.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/p_256_ecc_pp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/p_256_multprecision.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/smp_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/smp_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/smp_br_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/smp_keys.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/smp_l2c.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/smp_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp/smp_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/srvc/srvc_dis.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/srvc/srvc_eng.cc
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
    add_library(libbt-stack-core SHARED ${SOURCE_FILES})

    target_include_directories(libbt-stack-core
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avct
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/bnep
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/eatt
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gap
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/hid
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/l2cap
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/metrics
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/pan
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/sdp
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/srvc
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack
    )

    # Add c/c++ flags
    target_compile_options(libbt-stack-core PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-stack-core
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
            libbluetooth_core_rs
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-hci
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
#################### libbt-stack-core ####################

#################### libbt-stack ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/crypto_toolbox/aes.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/crypto_toolbox/aes_cmac.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/crypto_toolbox/crypto_toolbox.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_aac.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_aac_decoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_aac_encoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_codec_config.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_sbc.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_sbc_decoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_sbc_encoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_sbc_up_sample.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_aptx.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_aptx_encoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_aptx_hd.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_aptx_hd_encoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_ldac.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_ldac_decoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_ldac_encoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_opus.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_opus_decoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/a2dp/a2dp_vendor_opus_encoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avct/avct_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avct/avct_bcb_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avct/avct_ccb.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avct/avct_l2c.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avct/avct_l2c_br.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avct/avct_lcb.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avct/avct_lcb_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt/avdt_ad.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt/avdt_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt/avdt_ccb.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt/avdt_ccb_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt/avdt_l2c.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt/avdt_msg.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt/avdt_scb.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt/avdt_scb_act.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc/avrc_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc/avrc_bld_ct.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc/avrc_bld_tg.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc/avrc_opt.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc/avrc_pars_ct.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc/avrc_pars_tg.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc/avrc_sdp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc/avrc_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/bnep/bnep_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/bnep/bnep_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/bnep/bnep_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/hfp_msbc_decoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm/hfp_msbc_encoder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/hid/hidd_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/hid/hidd_conn.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/hid/hidh_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/hid/hidh_conn.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/pan/pan_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/pan/pan_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/pan/pan_utils.cc
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
    add_library(libbt-stack SHARED ${SOURCE_FILES})

    target_include_directories(libbt-stack
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avct
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avdt
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/avrc
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/bnep
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/btm
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/eatt
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gap
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/gatt
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/hid
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/l2cap
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/metrics
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/pan
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/sdp
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/smp
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack/srvc
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/stack
    )

    # Add c/c++ flags
    target_compile_options(libbt-stack PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-stack
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            libFraunhoferAAC
            libPlatformProperties
            libaaudio
            libaptx_enc
            libaptxhd_enc
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-hci
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
            libcom.android.sysprop.bluetooth
            libcrypto
            libcutils
            libdl
            libflatbuffers-cpp
            libfmq
            libg722codec
            libhidlbase
            liblc3
            libldacBT_abr
            libldacBT_enc
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
#################### libbt-stack ####################

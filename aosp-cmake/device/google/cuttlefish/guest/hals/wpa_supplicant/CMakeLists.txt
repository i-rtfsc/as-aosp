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

project(wpa_supplicant_cf)

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

#################### wpa_supplicant_cf ####################
set(FILES
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/ap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/bss.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/bssid_ignore.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/config.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/config_file.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/ctrl_iface.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/ctrl_iface_unix.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/dpp_supplicant.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/eap_register.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/events.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/gas_query.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/hs20_supplicant.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/interworking.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/main.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/mbo.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/notify.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/offchannel.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/op_classes.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/p2p_supplicant.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/p2p_supplicant_sd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/robust_av.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/rrm.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/scan.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/sme.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/acs.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ap_config.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ap_drv_ops.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ap_list.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/comeback_token.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/pasn/pasn_responder.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ap_mlme.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/authsrv.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/beacon.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/bss_load.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ctrl_iface_ap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/dfs.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/dpp_hostapd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/drv_callbacks.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/eap_user_db.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/fils_hlp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/gas_query_ap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/gas_serv.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/hostapd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/hs20.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/hw_features.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ieee802_11_auth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ieee802_11.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ieee802_11_ht.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ieee802_11_shared.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ieee802_11_vht.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/ieee802_1x.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/mbo_ap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/neighbor_db.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/p2p_hostapd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/pmksa_cache_auth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/rrm.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/sta_info.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/tkip_countermeasures.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/utils.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/wmm.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/wpa_auth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/wpa_auth_glue.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/wpa_auth_ie.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/ap/wps_hostapd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/ctrl_iface_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/dpp_auth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/dpp_backup.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/dpp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/dpp_crypto.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/dpp_pkex.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/dpp_reconfig.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/dpp_tcp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/dragonfly.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/gas.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/gas_server.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/hw_features_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/ieee802_11_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/sae.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/sae_pk.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/common/wpa_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/aes-ctr.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/aes-siv.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/crypto_openssl.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/dh_groups.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/fips_prf_openssl.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/ms_funcs.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/sha1-prf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/sha1-tlsprf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/sha256-kdf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/sha256-prf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/sha256-tlsprf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/sha384-kdf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/sha384-prf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/sha512-kdf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/sha512-prf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/tls_openssl.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/crypto/tls_openssl_ocsp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/driver_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/driver_nl80211_android.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/driver_nl80211.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/driver_nl80211_capa.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/driver_nl80211_event.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/driver_nl80211_monitor.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/driver_nl80211_scan.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/drivers.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/linux_ioctl.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/netlink.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/drivers/rfkill.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_common/chap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_common/eap_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_common/eap_peap_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_common/eap_pwd_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_common/eap_sim_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_common/eap_wsc_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eapol_auth/eapol_auth_sm.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eapol_supp/eapol_supp_sm.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_aka.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_gtc.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_leap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_md5.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_methods.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_mschapv2.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_otp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_peap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_pwd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_sim.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_tls.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_tls_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_ttls.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/eap_wsc.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_peer/mschapv2.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_server/eap_server.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_server/eap_server_identity.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_server/eap_server_methods.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/eap_server/eap_server_wsc.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/l2_packet/l2_packet_linux.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p_build.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p_dev_disc.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p_go_neg.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p_group.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p_invitation.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p_parse.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p_pd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p_sd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/p2p/p2p_utils.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/rsn_supp/pmksa_cache.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/rsn_supp/preauth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/rsn_supp/tdls.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/rsn_supp/wpa.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/rsn_supp/wpa_ft.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/rsn_supp/wpa_ie.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/tls/asn1.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/base64.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/bitfield.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/config.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/crc32.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/eloop.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/ip_addr.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/json.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/os_unix.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/radiotap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/uuid.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/wpabuf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/utils/wpa_debug.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/http_client.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/httpread.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/http_server.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/ndef.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/upnp_xml.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_attr_build.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_attr_parse.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_attr_process.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_dev_attr.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_enrollee.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_er.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_er_ssdp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_registrar.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_upnp_ap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_upnp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_upnp_event.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_upnp_ssdp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/src/wps/wps_upnp_web.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/wifi_display.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/wmm_ac.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/wnm_sta.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/wpas_glue.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/wpa_supplicant.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/wpa_supplicant/wps_supplicant.c
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
    add_executable(wpa_supplicant_cf ${SOURCE_FILES})

    target_include_directories(wpa_supplicant_cf
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/wpa_supplicant
    )

    # Add c/c++ flags
    target_compile_options(wpa_supplicant_cf PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(wpa_supplicant_cf
            android.hardware.wifi.supplicant-V2-ndk
            android.system.keystore2-V1-ndk
            lib_driver_cmd_simulated_cf_bp
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libkeystore-engine-wifi-hidl
            liblog
            libm
            libnl
            libssl
            libutils
            libwpa_aidl_bp
            libwpa_aidl_headers
            vendor.qti.hardware.wifi.supplicant-V1-ndk
            wpa_supplicant_headers
            libc
            libdl
            libm
    )
endif ()
#################### wpa_supplicant_cf ####################

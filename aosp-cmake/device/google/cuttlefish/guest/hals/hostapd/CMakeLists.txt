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

project(hostapd_cf)

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

#################### hostapd_cf ####################
set(FILES
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/main.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/config_file.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/ctrl_iface.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/eap_register.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/hostapd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/wpa_auth_glue.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/drv_callbacks.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ap_drv_ops.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/utils.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/authsrv.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ieee802_1x.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ap_config.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/eap_user_db.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ieee802_11_auth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/sta_info.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/wpa_auth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/tkip_countermeasures.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ap_mlme.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/wpa_auth_ie.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/preauth_auth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/pmksa_cache_auth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ieee802_11_shared.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/beacon.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/bss_load.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/neighbor_db.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/rrm.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/drivers.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/eloop.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/wpa_debug.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/wpabuf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/os_unix.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/ip_addr.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/crc32.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/ieee802_11_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/wpa_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/hw_features_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eapol_auth/eapol_auth_sm.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eapol_auth/eapol_auth_dump.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/vlan_init.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/vlan_ifconfig.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/vlan.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/ctrl_iface_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ctrl_iface_ap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/wpa_auth_ft.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/eth_p_oui.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/sae.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/sae_pk.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/fils_hlp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/wnm_ap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/mbo_ap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/driver_nl80211.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/driver_nl80211_android.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/driver_nl80211_capa.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/driver_nl80211_event.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/driver_nl80211_monitor.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/driver_nl80211_scan.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/netlink.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/rfkill.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/radiotap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/linux_ioctl.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/l2_packet/l2_packet_linux.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_md5.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_tls.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_peap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_common/eap_peap_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_ttls.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_mschapv2.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_gtc.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_sim.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_aka.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_common/eap_sim_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_sim_db.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_pwd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_common/eap_pwd_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/uuid.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/wps_hostapd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_wsc.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_common/eap_wsc_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/wps/wps.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/wps/wps_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/wps/wps_attr_parse.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/wps/wps_attr_build.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/wps/wps_attr_process.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/wps/wps_dev_attr.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/wps/wps_enrollee.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/wps/wps_registrar.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/wps/ndef.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/dpp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/dpp_auth.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/dpp_backup.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/dpp_crypto.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/dpp_pkex.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/dpp_reconfig.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/dpp_tcp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/dpp_hostapd.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/gas_query_ap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_common/eap_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_methods.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_identity.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/dragonfly.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/ms_funcs.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_common/chap.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/eap_server/eap_server_tls_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/tls_openssl.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/tls_openssl_ocsp.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/crypto_openssl.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/fips_prf_openssl.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/aes-siv.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/aes-ctr.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/sha1-prf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/sha1-tlsprf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/sha256-prf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/sha256-tlsprf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/sha256-kdf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/sha384-kdf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/sha512-kdf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/sha384-prf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/sha512-prf.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/tls/asn1.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/crypto/dh_groups.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/base64.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/utils/json.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/wmm.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ap_list.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/comeback_token.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/pasn/pasn_responder.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ieee802_11.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/hw_features.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/dfs.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ieee802_11_ht.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ieee802_11_vht.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/hs20.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/common/gas.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/gas_serv.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/x_snoop.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/dhcp_snoop.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/ndisc_snoop.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/drivers/driver_common.c
        ${ANDROID_ROOT}/external/wpa_supplicant_8/hostapd/src/ap/acs.c
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
    add_executable(hostapd_cf ${SOURCE_FILES})

    target_include_directories(hostapd_cf
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/hostapd
    )

    # Add c/c++ flags
    target_compile_options(hostapd_cf PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(hostapd_cf
            android.hardware.wifi.hostapd-V1-ndk
            hostapd_headers
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
            libhostapd_aidl_bp
            libhostapd_aidl_headers
            liblog
            libm
            libnl
            libssl
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### hostapd_cf ####################

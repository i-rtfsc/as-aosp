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

project(ot-ctl)

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

#################### openthread-radio ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/core/api/diags_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/error_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/instance_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/link_raw_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/logging_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/random_noncrypto_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/tasklet_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/binary_search.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/error.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/instance.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/log.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/random.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/string.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/tasklet.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/timer.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/uptime.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/aes_ccm.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/aes_ecb.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/crypto_platform.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/storage.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/diags/factory_diags.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/link_raw.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/mac_frame.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/mac_types.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/sub_mac.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/sub_mac_callbacks.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/radio/radio.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/radio/radio_callbacks.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/radio/radio_platform.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/link_quality.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/parse_cmdline.cpp
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
    add_library(openthread-radio SHARED ${SOURCE_FILES})

    target_include_directories(openthread-radio
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation
            ${ANDROID_ROOT}/external/openthread/examples/platforms
            ${ANDROID_ROOT}/external/openthread/src/cli
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/ncp
            ${ANDROID_ROOT}/external/openthread/src/lib/hdlc
            ${ANDROID_ROOT}/external/openthread/third_party
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/include
            ${ANDROID_ROOT}/external/openthread
            ${OUT_INTERMEDIATES_ROOT}/external/openthread/ot_version_header/gen
    )

    # Add c/c++ flags
    target_compile_options(openthread-radio PRIVATE
            -Werror
    )

    target_link_libraries(openthread-radio
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### openthread-radio ####################

#################### openthread-spinel-rcp ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/lib/spinel/spinel.c
        ${ANDROID_ROOT}/external/openthread/src/lib/spinel/spinel_buffer.cpp
        ${ANDROID_ROOT}/external/openthread/src/lib/spinel/spinel_decoder.cpp
        ${ANDROID_ROOT}/external/openthread/src/lib/spinel/spinel_encoder.cpp
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
    add_library(openthread-spinel-rcp SHARED ${SOURCE_FILES})

    target_include_directories(openthread-spinel-rcp
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation
            ${ANDROID_ROOT}/external/openthread/examples/platforms
            ${ANDROID_ROOT}/external/openthread/src/cli
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/ncp
            ${ANDROID_ROOT}/external/openthread/src/lib/hdlc
            ${ANDROID_ROOT}/external/openthread/third_party
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/include
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(openthread-spinel-rcp PRIVATE
            -Werror
    )

    target_link_libraries(openthread-spinel-rcp
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### openthread-spinel-rcp ####################

#################### openthread-posix ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/alarm.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/mainloop.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/hdlc_interface.cpp
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
    add_library(openthread-posix SHARED ${SOURCE_FILES})

    target_include_directories(openthread-posix
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/posix/platform
            ${ANDROID_ROOT}/external/openthread/src/posix/platform/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/src/lib/platform
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(openthread-posix PRIVATE
            -Werror
    )

    target_link_libraries(openthread-posix
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### openthread-posix ####################

#################### ot-rcp ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/examples/apps/ncp/main.c
        ${ANDROID_ROOT}/external/openthread/examples/apps/ncp/ncp.c
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
    add_executable(ot-rcp ${SOURCE_FILES})

    target_include_directories(ot-rcp
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation
            ${ANDROID_ROOT}/external/openthread/examples/platforms
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/src/cli
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/ncp
            ${ANDROID_ROOT}/external/openthread/src/lib/hdlc
            ${ANDROID_ROOT}/external/openthread/third_party
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/include
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(ot-rcp PRIVATE
            -Werror
    )

    target_link_libraries(ot-rcp
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
            openthread-hdlc
            openthread-radio
            openthread-rcp
            openthread-simulation
            openthread-spinel-rcp
            libc
            libdl
            libm
    )
endif ()
#################### ot-rcp ####################

#################### ot-core ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/android/thread_network_hal/hal_interface.cpp
        ${ANDROID_ROOT}/external/openthread/src/android/thread_network_hal/vendor_interface.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/backbone_router_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/backbone_router_ftd_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/border_agent_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/border_router_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/border_routing_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/channel_manager_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/channel_monitor_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/child_supervision_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/coap_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/coap_secure_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/commissioner_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/crypto_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/dataset_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/dataset_ftd_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/dataset_updater_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/diags_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/dns_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/dns_server_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/error_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/heap_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/icmp6_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/instance_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/ip6_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/jam_detection_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/joiner_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/link_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/link_metrics_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/link_raw_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/logging_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/message_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/multi_radio_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/netdata_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/netdata_publisher_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/netdiag_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/network_time_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/ping_sender_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/random_crypto_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/random_noncrypto_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/server_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/sntp_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/srp_client_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/srp_client_buffers_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/srp_server_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/tasklet_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/tcp_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/thread_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/thread_ftd_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/udp_api.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/backbone_router/backbone_tmf.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/backbone_router/bbr_leader.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/backbone_router/bbr_local.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/backbone_router/bbr_manager.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/backbone_router/multicast_listeners_table.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/backbone_router/ndproxy_table.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/border_router/infra_if.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/border_router/routing_manager.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/coap/coap.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/coap/coap_message.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/coap/coap_secure.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/binary_search.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/crc16.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/data.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/error.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/frame_builder.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/frame_data.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/heap.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/heap_data.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/heap_string.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/instance.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/log.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/message.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/notifier.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/random.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/settings.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/string.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/tasklet.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/time_ticker.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/timer.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/tlvs.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/common/trickle_timer.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/aes_ccm.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/aes_ecb.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/crypto_platform.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/ecdsa.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/hkdf_sha256.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/hmac_sha256.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/mbedtls.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/pbkdf2_cmac.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/sha256.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/crypto/storage.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/diags/factory_diags.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/channel_mask.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/data_poll_handler.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/data_poll_sender.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/link_raw.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/mac.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/mac_filter.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/mac_frame.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/mac_links.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/mac_types.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/sub_mac.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/mac/sub_mac_callbacks.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/announce_begin_client.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/border_agent.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/commissioner.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/dataset.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/dataset_local.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/dataset_manager.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/dataset_manager_ftd.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/dataset_updater.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/dtls.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/energy_scan_client.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/extended_panid.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/joiner.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/joiner_router.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/meshcop.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/meshcop_leader.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/meshcop_tlvs.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/network_name.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/panid_query_client.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/meshcop/timestamp.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/checksum.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/dhcp6_client.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/dhcp6_server.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/dns_client.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/dns_types.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/dnssd_server.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/icmp6.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/ip4_types.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/ip6.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/ip6_address.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/ip6_filter.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/ip6_headers.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/ip6_mpl.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/nd6.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/nd_agent.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/netif.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/sntp_client.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/socket.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/srp_client.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/srp_server.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/tcp6.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/net/udp6.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/radio/radio.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/radio/radio_callbacks.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/radio/radio_platform.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/radio/trel_interface.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/radio/trel_link.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/radio/trel_packet.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/address_resolver.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/announce_begin_server.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/announce_sender.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/child_table.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/csl_tx_scheduler.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/discover_scanner.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/dua_manager.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/energy_scan_server.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/indirect_sender.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/key_manager.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/link_metrics.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/link_quality.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/lowpan.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/mesh_forwarder.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/mesh_forwarder_ftd.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/mesh_forwarder_mtd.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/mle.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/mle_router.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/mle_types.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/mlr_manager.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/neighbor_table.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_data.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_data_leader.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_data_leader_ftd.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_data_local.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_data_notifier.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_data_publisher.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_data_service.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_data_tlvs.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_data_types.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/network_diagnostic.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/panid_query_server.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/radio_selector.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/router_table.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/src_match_controller.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/thread_netif.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/time_sync_service.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/tmf.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/topology.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/thread/uri_paths.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/channel_manager.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/channel_monitor.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/child_supervision.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/flash.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/heap.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/jam_detector.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/otns.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/parse_cmdline.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/ping_sender.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/slaac_address.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/utils/srp_client_buffers.cpp
        ${ANDROID_ROOT}/external/openthread/src/lib/hdlc/hdlc.cpp
        ${ANDROID_ROOT}/external/openthread/src/lib/platform/exit_code.c
        ${ANDROID_ROOT}/external/openthread/src/lib/spinel/spinel.c
        ${ANDROID_ROOT}/external/openthread/src/lib/spinel/spinel_buffer.cpp
        ${ANDROID_ROOT}/external/openthread/src/lib/spinel/spinel_decoder.cpp
        ${ANDROID_ROOT}/external/openthread/src/lib/spinel/spinel_encoder.cpp
        ${ANDROID_ROOT}/external/openthread/src/lib/url/url.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/alarm.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/backbone.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/daemon.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/entropy.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/hdlc_interface.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/infra_if.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/logging.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/mainloop.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/memory.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/misc.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/multicast_routing.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/netif.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/radio.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/radio_url.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/settings.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/spi_interface.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/system.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/trel.cpp
        ${ANDROID_ROOT}/external/openthread/src/posix/platform/udp.cpp
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/aes.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/asn1parse.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/asn1write.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/base64.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/bignum.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ccm.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/cipher.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/cipher_wrap.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/cmac.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/constant_time.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ctr_drbg.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/debug.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ecdh.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ecdsa.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ecjpake.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ecp.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ecp_curves.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/entropy.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/entropy_poll.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/hmac_drbg.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/md.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/memory_buffer_alloc.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/oid.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/pem.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/pk.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/pk_wrap.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/pkparse.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/pkwrite.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/platform.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/platform_util.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/sha256.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ssl_ciphersuites.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ssl_cli.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ssl_cookie.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ssl_msg.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ssl_srv.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ssl_ticket.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ssl_tls.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/threading.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/x509.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/x509_crt.c
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
    add_library(ot-core SHARED ${SOURCE_FILES})

    target_include_directories(ot-core
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/src/android/thread_network_hal
            ${ANDROID_ROOT}/external/openthread/src/cli
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/ncp
            ${ANDROID_ROOT}/external/openthread/src/posix/platform
            ${ANDROID_ROOT}/external/openthread/src/posix/platform/include
            ${ANDROID_ROOT}/external/openthread/third_party
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/include
            ${ANDROID_ROOT}/external/openthread
            ${OUT_INTERMEDIATES_ROOT}/external/openthread/ot_version_header/gen
    )

    # Add c/c++ flags
    target_compile_options(ot-core PRIVATE
            -Werror
    )

    target_link_libraries(ot-core
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### ot-core ####################

#################### ot-ctl ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/posix/client.cpp
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
    add_executable(ot-ctl ${SOURCE_FILES})

    target_include_directories(ot-ctl
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/posix/platform
            ${ANDROID_ROOT}/external/openthread/src/posix/platform/include
            ${ANDROID_ROOT}/external/openthread
            ${OUT_INTERMEDIATES_ROOT}/external/openthread/ot_version_header/gen
    )

    # Add c/c++ flags
    target_compile_options(ot-ctl PRIVATE
            -Werror
    )

    target_link_libraries(ot-ctl
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### ot-ctl ####################

#################### openthread-rcp ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/ncp/changed_props_set.cpp
        ${ANDROID_ROOT}/external/openthread/src/ncp/ncp_base.cpp
        ${ANDROID_ROOT}/external/openthread/src/ncp/ncp_base_dispatcher.cpp
        ${ANDROID_ROOT}/external/openthread/src/ncp/ncp_base_radio.cpp
        ${ANDROID_ROOT}/external/openthread/src/ncp/ncp_spi.cpp
        ${ANDROID_ROOT}/external/openthread/src/ncp/ncp_hdlc.cpp
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
    add_library(openthread-rcp SHARED ${SOURCE_FILES})

    target_include_directories(openthread-rcp
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation
            ${ANDROID_ROOT}/external/openthread/examples/platforms
            ${ANDROID_ROOT}/external/openthread/src/cli
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/ncp
            ${ANDROID_ROOT}/external/openthread/src/lib/hdlc
            ${ANDROID_ROOT}/external/openthread/third_party
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/include
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(openthread-rcp PRIVATE
            -Werror
    )

    target_link_libraries(openthread-rcp
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### openthread-rcp ####################

#################### openthread-common ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/core/common/error.cpp
        ${ANDROID_ROOT}/external/openthread/src/core/api/error_api.cpp
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
    add_library(openthread-common SHARED ${SOURCE_FILES})

    target_include_directories(openthread-common
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(openthread-common PRIVATE
            -Werror
    )

    target_link_libraries(openthread-common
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### openthread-common ####################

#################### openthread-platform ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/lib/platform/exit_code.c
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
    add_library(openthread-platform SHARED ${SOURCE_FILES})

    target_include_directories(openthread-platform
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/lib/platform
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(openthread-platform PRIVATE
            -Werror
    )

    target_link_libraries(openthread-platform
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### openthread-platform ####################

#################### openthread-tcplp ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/bsdtcp/cc/cc_newreno.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/bsdtcp/tcp_input.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/bsdtcp/tcp_output.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/bsdtcp/tcp_reass.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/bsdtcp/tcp_sack.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/bsdtcp/tcp_subr.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/bsdtcp/tcp_timer.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/bsdtcp/tcp_timewait.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/bsdtcp/tcp_usrreq.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/lib/bitmap.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/lib/cbuf.c
        ${ANDROID_ROOT}/external/openthread/third_party/tcplp/lib/lbuf.c
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
    add_library(openthread-tcplp SHARED ${SOURCE_FILES})

    target_include_directories(openthread-tcplp
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(openthread-tcplp PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(openthread-tcplp
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### openthread-tcplp ####################

#################### openthread-url ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/lib/url/url.cpp
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
    add_library(openthread-url SHARED ${SOURCE_FILES})

    target_include_directories(openthread-url
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/lib/url
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(openthread-url PRIVATE
            -Werror
    )

    target_link_libraries(openthread-url
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### openthread-url ####################

#################### libopenthread-cli ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/cli/cli.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_coap.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_coap_secure.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_commissioner.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_dataset.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_history.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_joiner.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_network_data.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_output.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_srp_client.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_srp_server.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_tcp.cpp
        ${ANDROID_ROOT}/external/openthread/src/cli/cli_udp.cpp
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
    add_library(libopenthread-cli SHARED ${SOURCE_FILES})

    target_include_directories(libopenthread-cli
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/src/cli
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/posix/platform
            ${ANDROID_ROOT}/external/openthread/src/posix/platform/include
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/include
            ${ANDROID_ROOT}/external/openthread
            ${OUT_INTERMEDIATES_ROOT}/external/openthread/ot_version_header/gen
    )

    # Add c/c++ flags
    target_compile_options(libopenthread-cli PRIVATE
            -Werror
    )

    target_link_libraries(libopenthread-cli
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libopenthread-cli ####################

#################### openthread-hdlc ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/src/lib/hdlc/hdlc.cpp
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
    add_library(openthread-hdlc SHARED ${SOURCE_FILES})

    target_include_directories(openthread-hdlc
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation
            ${ANDROID_ROOT}/external/openthread/examples/platforms
            ${ANDROID_ROOT}/external/openthread/src/cli
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/ncp
            ${ANDROID_ROOT}/external/openthread/src/lib/hdlc
            ${ANDROID_ROOT}/external/openthread/third_party
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/include
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(openthread-hdlc PRIVATE
            -Werror
    )

    target_link_libraries(openthread-hdlc
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### openthread-hdlc ####################

#################### openthread-simulation ####################
set(FILES
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/alarm.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/crypto.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/diag.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/entropy.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/flash.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/infra_if.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/logging.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/misc.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/radio.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/spi-stubs.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/system.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/trel.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation/uart.c
        ${ANDROID_ROOT}/external/openthread/examples/platforms/utils/link_metrics.cpp
        ${ANDROID_ROOT}/external/openthread/examples/platforms/utils/mac_frame.cpp
        ${ANDROID_ROOT}/external/openthread/examples/platforms/utils/soft_source_match_table.c
        ${ANDROID_ROOT}/external/openthread/src/lib/platform/exit_code.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/aes.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/asn1parse.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/asn1write.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/base64.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/bignum.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ccm.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/cipher.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/cipher_wrap.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/cmac.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/ctr_drbg.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/debug.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/entropy.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/entropy_poll.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/hmac_drbg.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/md.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/memory_buffer_alloc.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/oid.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/platform.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/platform_util.c
        ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/library/sha256.c
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
    add_library(openthread-simulation SHARED ${SOURCE_FILES})

    target_include_directories(openthread-simulation
            PUBLIC
            ${ANDROID_ROOT}/external/openthread/include
            ${ANDROID_ROOT}/external/openthread/src
            ${ANDROID_ROOT}/external/openthread/examples/platforms/simulation
            ${ANDROID_ROOT}/external/openthread/examples/platforms
            ${ANDROID_ROOT}/external/openthread/src/cli
            ${ANDROID_ROOT}/external/openthread/src/core
            ${ANDROID_ROOT}/external/openthread/src/ncp
            ${ANDROID_ROOT}/external/openthread/src/lib/hdlc
            ${ANDROID_ROOT}/external/openthread/third_party
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls
            ${ANDROID_ROOT}/external/openthread/third_party/mbedtls/repo/include
            ${ANDROID_ROOT}/external/openthread
    )

    # Add c/c++ flags
    target_compile_options(openthread-simulation PRIVATE
            -Werror
    )

    target_link_libraries(openthread-simulation
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### openthread-simulation ####################

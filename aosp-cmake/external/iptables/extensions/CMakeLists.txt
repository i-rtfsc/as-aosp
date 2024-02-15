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

project(libext4)

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

#################### libext4 ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_init/gen/initext.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_CLUSTERIP.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_DNAT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_ECN.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_LOG.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_MASQUERADE.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_NETMAP.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_REDIRECT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_REJECT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_SNAT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_TTL.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_ULOG.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_ah.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_icmp.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_realm.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext4_srcs/gen/gensrcs/external/iptables/extensions/libipt_ttl.c
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
    add_library(libext4 SHARED ${SOURCE_FILES})

    target_include_directories(libext4
            PUBLIC
            ${ANDROID_ROOT}/external/iptables/extensions
    )

    # Add c/c++ flags
    target_compile_options(libext4 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libext4
            iptables_config_header
            iptables_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libext4 ####################

#################### libext6 ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_init/gen/initext.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_DNAT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_DNPT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_HL.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_LOG.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_MASQUERADE.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_NETMAP.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_REDIRECT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_REJECT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_SNAT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_SNPT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_ah.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_dst.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_eui64.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_frag.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_hbh.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_hl.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_icmp6.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_ipv6header.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_mh.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_rt.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext6_srcs/gen/gensrcs/external/iptables/extensions/libip6t_srh.c
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
    add_library(libext6 SHARED ${SOURCE_FILES})

    target_include_directories(libext6
            PUBLIC
            ${ANDROID_ROOT}/external/iptables/extensions
    )

    # Add c/c++ flags
    target_compile_options(libext6 PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libext6
            iptables_config_header
            iptables_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libext6 ####################

#################### libext ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_init/gen/initext.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_AUDIT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_CHECKSUM.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_CLASSIFY.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_CONNMARK.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_CONNSECMARK.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_CT.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_DSCP.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_HARDIDLETIMER.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_HMARK.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_IDLETIMER.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_LED.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_MARK.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_NFLOG.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_NFQUEUE.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_RATEEST.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_SECMARK.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_SET.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_SYNPROXY.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_TCPMSS.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_TEE.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_TOS.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_TPROXY.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_TRACE.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_addrtype.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_bpf.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_cluster.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_comment.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_connbytes.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_connlimit.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_connmark.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_conntrack.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_cpu.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_devgroup.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_dscp.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_ecn.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_esp.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_hashlimit.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_helper.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_ipcomp.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_iprange.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_length.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_limit.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_mac.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_mark.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_multiport.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_nfacct.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_osf.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_owner.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_physdev.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_pkttype.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_policy.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_quota.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_quota2.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_rateest.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_recent.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_rpfilter.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_sctp.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_set.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_socket.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_standard.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_statistic.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_string.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_tcp.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_tcpmss.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_time.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_tos.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_u32.c
        ${OUT_INTERMEDIATES_ROOT}/external/iptables/extensions/libext_srcs/gen/gensrcs/external/iptables/extensions/libxt_udp.c
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
    add_library(libext SHARED ${SOURCE_FILES})

    target_include_directories(libext
            PUBLIC
            ${ANDROID_ROOT}/external/iptables/extensions
    )

    # Add c/c++ flags
    target_compile_options(libext PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libext
            iptables_config_header
            iptables_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libext ####################

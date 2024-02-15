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

project(ip)

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

#################### ip ####################
set(FILES
        ${ANDROID_ROOT}/external/iproute2/ip/ip6tunnel.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipaddress.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipaddrlabel.c
        ${ANDROID_ROOT}/external/iproute2/ip/ip.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipfou.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipila.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipl2tp.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_bond.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_bond_slave.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_bridge.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_bridge_slave.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_can.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_dummy.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_geneve.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_hsr.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_ifb.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_ipoib.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_ipvlan.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_macvlan.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_nlmon.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_team.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_vcan.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_vlan.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_vrf.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_vxcan.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_vxlan.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_xdp.c
        ${ANDROID_ROOT}/external/iproute2/ip/iplink_xstats.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipmacsec.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipmaddr.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipmonitor.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipmroute.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipneigh.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipnetconf.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipnetns.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipntable.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipprefix.c
        ${ANDROID_ROOT}/external/iproute2/ip/iproute.c
        ${ANDROID_ROOT}/external/iproute2/ip/iproute_lwtunnel.c
        ${ANDROID_ROOT}/external/iproute2/ip/iprule.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipseg6.c
        ${ANDROID_ROOT}/external/iproute2/ip/iptoken.c
        ${ANDROID_ROOT}/external/iproute2/ip/iptunnel.c
        ${ANDROID_ROOT}/external/iproute2/ip/iptuntap.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipvrf.c
        ${ANDROID_ROOT}/external/iproute2/ip/ipxfrm.c
        ${ANDROID_ROOT}/external/iproute2/ip/link_gre6.c
        ${ANDROID_ROOT}/external/iproute2/ip/link_gre.c
        ${ANDROID_ROOT}/external/iproute2/ip/link_ip6tnl.c
        ${ANDROID_ROOT}/external/iproute2/ip/link_iptnl.c
        ${ANDROID_ROOT}/external/iproute2/ip/link_veth.c
        ${ANDROID_ROOT}/external/iproute2/ip/link_vti6.c
        ${ANDROID_ROOT}/external/iproute2/ip/link_vti.c
        ${ANDROID_ROOT}/external/iproute2/ip/rtm_map.c
        ${ANDROID_ROOT}/external/iproute2/ip/static-syms.c
        ${ANDROID_ROOT}/external/iproute2/ip/tcp_metrics.c
        ${ANDROID_ROOT}/external/iproute2/ip/tunnel.c
        ${ANDROID_ROOT}/external/iproute2/ip/xfrm_monitor.c
        ${ANDROID_ROOT}/external/iproute2/ip/xfrm_policy.c
        ${ANDROID_ROOT}/external/iproute2/ip/xfrm_state.c
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
    add_executable(ip ${SOURCE_FILES})

    target_include_directories(ip
            PUBLIC
            ${ANDROID_ROOT}/external/iproute2/ip
    )

    # Add c/c++ flags
    target_compile_options(ip PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(ip
            iproute2_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libiprouteutil
            libm
            libnetlink
    )
endif ()
#################### ip ####################

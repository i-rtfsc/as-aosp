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

project(tcpdump)

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

#################### tcpdump_vendor ####################
set(FILES
        ${ANDROID_ROOT}/external/tcpdump/tcpdump.c
        ${ANDROID_ROOT}/external/tcpdump/addrtoname.c
        ${ANDROID_ROOT}/external/tcpdump/addrtostr.c
        ${ANDROID_ROOT}/external/tcpdump/af.c
        ${ANDROID_ROOT}/external/tcpdump/ascii_strcasecmp.c
        ${ANDROID_ROOT}/external/tcpdump/checksum.c
        ${ANDROID_ROOT}/external/tcpdump/cpack.c
        ${ANDROID_ROOT}/external/tcpdump/fptype.c
        ${ANDROID_ROOT}/external/tcpdump/gmpls.c
        ${ANDROID_ROOT}/external/tcpdump/in_cksum.c
        ${ANDROID_ROOT}/external/tcpdump/ipproto.c
        ${ANDROID_ROOT}/external/tcpdump/l2vpn.c
        ${ANDROID_ROOT}/external/tcpdump/machdep.c
        ${ANDROID_ROOT}/external/tcpdump/netdissect-alloc.c
        ${ANDROID_ROOT}/external/tcpdump/nlpid.c
        ${ANDROID_ROOT}/external/tcpdump/ntp.c
        ${ANDROID_ROOT}/external/tcpdump/oui.c
        ${ANDROID_ROOT}/external/tcpdump/parsenfsfh.c
        ${ANDROID_ROOT}/external/tcpdump/print.c
        ${ANDROID_ROOT}/external/tcpdump/print-802_11.c
        ${ANDROID_ROOT}/external/tcpdump/print-802_15_4.c
        ${ANDROID_ROOT}/external/tcpdump/print-ah.c
        ${ANDROID_ROOT}/external/tcpdump/print-ahcp.c
        ${ANDROID_ROOT}/external/tcpdump/print-aodv.c
        ${ANDROID_ROOT}/external/tcpdump/print-aoe.c
        ${ANDROID_ROOT}/external/tcpdump/print-ap1394.c
        ${ANDROID_ROOT}/external/tcpdump/print-arcnet.c
        ${ANDROID_ROOT}/external/tcpdump/print-arista.c
        ${ANDROID_ROOT}/external/tcpdump/print-arp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ascii.c
        ${ANDROID_ROOT}/external/tcpdump/print-atalk.c
        ${ANDROID_ROOT}/external/tcpdump/print-atm.c
        ${ANDROID_ROOT}/external/tcpdump/print-babel.c
        ${ANDROID_ROOT}/external/tcpdump/print-bcm-li.c
        ${ANDROID_ROOT}/external/tcpdump/print-beep.c
        ${ANDROID_ROOT}/external/tcpdump/print-bfd.c
        ${ANDROID_ROOT}/external/tcpdump/print-bgp.c
        ${ANDROID_ROOT}/external/tcpdump/print-bootp.c
        ${ANDROID_ROOT}/external/tcpdump/print-brcmtag.c
        ${ANDROID_ROOT}/external/tcpdump/print-bt.c
        ${ANDROID_ROOT}/external/tcpdump/print-calm-fast.c
        ${ANDROID_ROOT}/external/tcpdump/print-carp.c
        ${ANDROID_ROOT}/external/tcpdump/print-cdp.c
        ${ANDROID_ROOT}/external/tcpdump/print-cfm.c
        ${ANDROID_ROOT}/external/tcpdump/print-chdlc.c
        ${ANDROID_ROOT}/external/tcpdump/print-cip.c
        ${ANDROID_ROOT}/external/tcpdump/print-cnfp.c
        ${ANDROID_ROOT}/external/tcpdump/print-dccp.c
        ${ANDROID_ROOT}/external/tcpdump/print-decnet.c
        ${ANDROID_ROOT}/external/tcpdump/print-dhcp6.c
        ${ANDROID_ROOT}/external/tcpdump/print-domain.c
        ${ANDROID_ROOT}/external/tcpdump/print-dsa.c
        ${ANDROID_ROOT}/external/tcpdump/print-dtp.c
        ${ANDROID_ROOT}/external/tcpdump/print-dvmrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-eap.c
        ${ANDROID_ROOT}/external/tcpdump/print-egp.c
        ${ANDROID_ROOT}/external/tcpdump/print-eigrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-enc.c
        ${ANDROID_ROOT}/external/tcpdump/print-esp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ether.c
        ${ANDROID_ROOT}/external/tcpdump/print-fddi.c
        ${ANDROID_ROOT}/external/tcpdump/print-forces.c
        ${ANDROID_ROOT}/external/tcpdump/print-fr.c
        ${ANDROID_ROOT}/external/tcpdump/print-frag6.c
        ${ANDROID_ROOT}/external/tcpdump/print-ftp.c
        ${ANDROID_ROOT}/external/tcpdump/print-geneve.c
        ${ANDROID_ROOT}/external/tcpdump/print-geonet.c
        ${ANDROID_ROOT}/external/tcpdump/print-gre.c
        ${ANDROID_ROOT}/external/tcpdump/print-hncp.c
        ${ANDROID_ROOT}/external/tcpdump/print-hsrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-http.c
        ${ANDROID_ROOT}/external/tcpdump/print-icmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-icmp6.c
        ${ANDROID_ROOT}/external/tcpdump/print-igmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-igrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ip-demux.c
        ${ANDROID_ROOT}/external/tcpdump/print-ip.c
        ${ANDROID_ROOT}/external/tcpdump/print-ip6.c
        ${ANDROID_ROOT}/external/tcpdump/print-ip6opts.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipcomp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipfc.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipnet.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipoib.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipx.c
        ${ANDROID_ROOT}/external/tcpdump/print-isakmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-isoclns.c
        ${ANDROID_ROOT}/external/tcpdump/print-juniper.c
        ${ANDROID_ROOT}/external/tcpdump/print-krb.c
        ${ANDROID_ROOT}/external/tcpdump/print-l2tp.c
        ${ANDROID_ROOT}/external/tcpdump/print-lane.c
        ${ANDROID_ROOT}/external/tcpdump/print-ldp.c
        ${ANDROID_ROOT}/external/tcpdump/print-lisp.c
        ${ANDROID_ROOT}/external/tcpdump/print-llc.c
        ${ANDROID_ROOT}/external/tcpdump/print-lldp.c
        ${ANDROID_ROOT}/external/tcpdump/print-lmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-loopback.c
        ${ANDROID_ROOT}/external/tcpdump/print-lspping.c
        ${ANDROID_ROOT}/external/tcpdump/print-lwapp.c
        ${ANDROID_ROOT}/external/tcpdump/print-lwres.c
        ${ANDROID_ROOT}/external/tcpdump/print-m3ua.c
        ${ANDROID_ROOT}/external/tcpdump/print-macsec.c
        ${ANDROID_ROOT}/external/tcpdump/print-mobile.c
        ${ANDROID_ROOT}/external/tcpdump/print-mobility.c
        ${ANDROID_ROOT}/external/tcpdump/print-mpcp.c
        ${ANDROID_ROOT}/external/tcpdump/print-mpls.c
        ${ANDROID_ROOT}/external/tcpdump/print-mptcp.c
        ${ANDROID_ROOT}/external/tcpdump/print-msdp.c
        ${ANDROID_ROOT}/external/tcpdump/print-msnlb.c
        ${ANDROID_ROOT}/external/tcpdump/print-nflog.c
        ${ANDROID_ROOT}/external/tcpdump/print-nfs.c
        ${ANDROID_ROOT}/external/tcpdump/print-nsh.c
        ${ANDROID_ROOT}/external/tcpdump/print-ntp.c
        ${ANDROID_ROOT}/external/tcpdump/print-null.c
        ${ANDROID_ROOT}/external/tcpdump/print-olsr.c
        ${ANDROID_ROOT}/external/tcpdump/print-openflow-1.0.c
        ${ANDROID_ROOT}/external/tcpdump/print-openflow-1.3.c
        ${ANDROID_ROOT}/external/tcpdump/print-openflow.c
        ${ANDROID_ROOT}/external/tcpdump/print-ospf.c
        ${ANDROID_ROOT}/external/tcpdump/print-ospf6.c
        ${ANDROID_ROOT}/external/tcpdump/print-otv.c
        ${ANDROID_ROOT}/external/tcpdump/print-pflog.c
        ${ANDROID_ROOT}/external/tcpdump/print-pgm.c
        ${ANDROID_ROOT}/external/tcpdump/print-pim.c
        ${ANDROID_ROOT}/external/tcpdump/print-pktap.c
        ${ANDROID_ROOT}/external/tcpdump/print-ppi.c
        ${ANDROID_ROOT}/external/tcpdump/print-ppp.c
        ${ANDROID_ROOT}/external/tcpdump/print-pppoe.c
        ${ANDROID_ROOT}/external/tcpdump/print-pptp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ptp.c
        ${ANDROID_ROOT}/external/tcpdump/print-radius.c
        ${ANDROID_ROOT}/external/tcpdump/print-raw.c
        ${ANDROID_ROOT}/external/tcpdump/print-realtek.c
        ${ANDROID_ROOT}/external/tcpdump/print-resp.c
        ${ANDROID_ROOT}/external/tcpdump/print-rip.c
        ${ANDROID_ROOT}/external/tcpdump/print-ripng.c
        ${ANDROID_ROOT}/external/tcpdump/print-rpki-rtr.c
        ${ANDROID_ROOT}/external/tcpdump/print-rsvp.c
        ${ANDROID_ROOT}/external/tcpdump/print-rt6.c
        ${ANDROID_ROOT}/external/tcpdump/print-rtsp.c
        ${ANDROID_ROOT}/external/tcpdump/print-rx.c
        ${ANDROID_ROOT}/external/tcpdump/print-sctp.c
        ${ANDROID_ROOT}/external/tcpdump/print-sflow.c
        ${ANDROID_ROOT}/external/tcpdump/print-sip.c
        ${ANDROID_ROOT}/external/tcpdump/print-sl.c
        ${ANDROID_ROOT}/external/tcpdump/print-sll.c
        ${ANDROID_ROOT}/external/tcpdump/print-slow.c
        ${ANDROID_ROOT}/external/tcpdump/print-smb.c
        ${ANDROID_ROOT}/external/tcpdump/print-smtp.c
        ${ANDROID_ROOT}/external/tcpdump/print-snmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-someip.c
        ${ANDROID_ROOT}/external/tcpdump/print-ssh.c
        ${ANDROID_ROOT}/external/tcpdump/print-stp.c
        ${ANDROID_ROOT}/external/tcpdump/print-sunatm.c
        ${ANDROID_ROOT}/external/tcpdump/print-sunrpc.c
        ${ANDROID_ROOT}/external/tcpdump/print-symantec.c
        ${ANDROID_ROOT}/external/tcpdump/print-syslog.c
        ${ANDROID_ROOT}/external/tcpdump/print-tcp.c
        ${ANDROID_ROOT}/external/tcpdump/print-telnet.c
        ${ANDROID_ROOT}/external/tcpdump/print-tftp.c
        ${ANDROID_ROOT}/external/tcpdump/print-timed.c
        ${ANDROID_ROOT}/external/tcpdump/print-tipc.c
        ${ANDROID_ROOT}/external/tcpdump/print-token.c
        ${ANDROID_ROOT}/external/tcpdump/print-udld.c
        ${ANDROID_ROOT}/external/tcpdump/print-udp.c
        ${ANDROID_ROOT}/external/tcpdump/print-unsupported.c
        ${ANDROID_ROOT}/external/tcpdump/print-usb.c
        ${ANDROID_ROOT}/external/tcpdump/print-vjc.c
        ${ANDROID_ROOT}/external/tcpdump/print-vqp.c
        ${ANDROID_ROOT}/external/tcpdump/print-vrrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-vsock.c
        ${ANDROID_ROOT}/external/tcpdump/print-vtp.c
        ${ANDROID_ROOT}/external/tcpdump/print-vxlan-gpe.c
        ${ANDROID_ROOT}/external/tcpdump/print-vxlan.c
        ${ANDROID_ROOT}/external/tcpdump/print-wb.c
        ${ANDROID_ROOT}/external/tcpdump/print-whois.c
        ${ANDROID_ROOT}/external/tcpdump/print-zep.c
        ${ANDROID_ROOT}/external/tcpdump/print-zephyr.c
        ${ANDROID_ROOT}/external/tcpdump/print-zeromq.c
        ${ANDROID_ROOT}/external/tcpdump/netdissect.c
        ${ANDROID_ROOT}/external/tcpdump/signature.c
        ${ANDROID_ROOT}/external/tcpdump/strtoaddr.c
        ${ANDROID_ROOT}/external/tcpdump/util-print.c
        ${ANDROID_ROOT}/external/tcpdump/smbutil.c
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
    add_executable(tcpdump_vendor ${SOURCE_FILES})

    target_include_directories(tcpdump_vendor
            PUBLIC
            ${ANDROID_ROOT}/external/tcpdump
    )

    # Add c/c++ flags
    target_compile_options(tcpdump_vendor PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(tcpdump_vendor
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libpcap
            libssl
            libc
            libdl
            libm
    )
endif ()
#################### tcpdump_vendor ####################

#################### tcpdump ####################
set(FILES
        ${ANDROID_ROOT}/external/tcpdump/tcpdump.c
        ${ANDROID_ROOT}/external/tcpdump/addrtoname.c
        ${ANDROID_ROOT}/external/tcpdump/addrtostr.c
        ${ANDROID_ROOT}/external/tcpdump/af.c
        ${ANDROID_ROOT}/external/tcpdump/ascii_strcasecmp.c
        ${ANDROID_ROOT}/external/tcpdump/checksum.c
        ${ANDROID_ROOT}/external/tcpdump/cpack.c
        ${ANDROID_ROOT}/external/tcpdump/fptype.c
        ${ANDROID_ROOT}/external/tcpdump/gmpls.c
        ${ANDROID_ROOT}/external/tcpdump/in_cksum.c
        ${ANDROID_ROOT}/external/tcpdump/ipproto.c
        ${ANDROID_ROOT}/external/tcpdump/l2vpn.c
        ${ANDROID_ROOT}/external/tcpdump/machdep.c
        ${ANDROID_ROOT}/external/tcpdump/netdissect-alloc.c
        ${ANDROID_ROOT}/external/tcpdump/nlpid.c
        ${ANDROID_ROOT}/external/tcpdump/ntp.c
        ${ANDROID_ROOT}/external/tcpdump/oui.c
        ${ANDROID_ROOT}/external/tcpdump/parsenfsfh.c
        ${ANDROID_ROOT}/external/tcpdump/print.c
        ${ANDROID_ROOT}/external/tcpdump/print-802_11.c
        ${ANDROID_ROOT}/external/tcpdump/print-802_15_4.c
        ${ANDROID_ROOT}/external/tcpdump/print-ah.c
        ${ANDROID_ROOT}/external/tcpdump/print-ahcp.c
        ${ANDROID_ROOT}/external/tcpdump/print-aodv.c
        ${ANDROID_ROOT}/external/tcpdump/print-aoe.c
        ${ANDROID_ROOT}/external/tcpdump/print-ap1394.c
        ${ANDROID_ROOT}/external/tcpdump/print-arcnet.c
        ${ANDROID_ROOT}/external/tcpdump/print-arista.c
        ${ANDROID_ROOT}/external/tcpdump/print-arp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ascii.c
        ${ANDROID_ROOT}/external/tcpdump/print-atalk.c
        ${ANDROID_ROOT}/external/tcpdump/print-atm.c
        ${ANDROID_ROOT}/external/tcpdump/print-babel.c
        ${ANDROID_ROOT}/external/tcpdump/print-bcm-li.c
        ${ANDROID_ROOT}/external/tcpdump/print-beep.c
        ${ANDROID_ROOT}/external/tcpdump/print-bfd.c
        ${ANDROID_ROOT}/external/tcpdump/print-bgp.c
        ${ANDROID_ROOT}/external/tcpdump/print-bootp.c
        ${ANDROID_ROOT}/external/tcpdump/print-brcmtag.c
        ${ANDROID_ROOT}/external/tcpdump/print-bt.c
        ${ANDROID_ROOT}/external/tcpdump/print-calm-fast.c
        ${ANDROID_ROOT}/external/tcpdump/print-carp.c
        ${ANDROID_ROOT}/external/tcpdump/print-cdp.c
        ${ANDROID_ROOT}/external/tcpdump/print-cfm.c
        ${ANDROID_ROOT}/external/tcpdump/print-chdlc.c
        ${ANDROID_ROOT}/external/tcpdump/print-cip.c
        ${ANDROID_ROOT}/external/tcpdump/print-cnfp.c
        ${ANDROID_ROOT}/external/tcpdump/print-dccp.c
        ${ANDROID_ROOT}/external/tcpdump/print-decnet.c
        ${ANDROID_ROOT}/external/tcpdump/print-dhcp6.c
        ${ANDROID_ROOT}/external/tcpdump/print-domain.c
        ${ANDROID_ROOT}/external/tcpdump/print-dsa.c
        ${ANDROID_ROOT}/external/tcpdump/print-dtp.c
        ${ANDROID_ROOT}/external/tcpdump/print-dvmrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-eap.c
        ${ANDROID_ROOT}/external/tcpdump/print-egp.c
        ${ANDROID_ROOT}/external/tcpdump/print-eigrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-enc.c
        ${ANDROID_ROOT}/external/tcpdump/print-esp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ether.c
        ${ANDROID_ROOT}/external/tcpdump/print-fddi.c
        ${ANDROID_ROOT}/external/tcpdump/print-forces.c
        ${ANDROID_ROOT}/external/tcpdump/print-fr.c
        ${ANDROID_ROOT}/external/tcpdump/print-frag6.c
        ${ANDROID_ROOT}/external/tcpdump/print-ftp.c
        ${ANDROID_ROOT}/external/tcpdump/print-geneve.c
        ${ANDROID_ROOT}/external/tcpdump/print-geonet.c
        ${ANDROID_ROOT}/external/tcpdump/print-gre.c
        ${ANDROID_ROOT}/external/tcpdump/print-hncp.c
        ${ANDROID_ROOT}/external/tcpdump/print-hsrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-http.c
        ${ANDROID_ROOT}/external/tcpdump/print-icmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-icmp6.c
        ${ANDROID_ROOT}/external/tcpdump/print-igmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-igrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ip-demux.c
        ${ANDROID_ROOT}/external/tcpdump/print-ip.c
        ${ANDROID_ROOT}/external/tcpdump/print-ip6.c
        ${ANDROID_ROOT}/external/tcpdump/print-ip6opts.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipcomp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipfc.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipnet.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipoib.c
        ${ANDROID_ROOT}/external/tcpdump/print-ipx.c
        ${ANDROID_ROOT}/external/tcpdump/print-isakmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-isoclns.c
        ${ANDROID_ROOT}/external/tcpdump/print-juniper.c
        ${ANDROID_ROOT}/external/tcpdump/print-krb.c
        ${ANDROID_ROOT}/external/tcpdump/print-l2tp.c
        ${ANDROID_ROOT}/external/tcpdump/print-lane.c
        ${ANDROID_ROOT}/external/tcpdump/print-ldp.c
        ${ANDROID_ROOT}/external/tcpdump/print-lisp.c
        ${ANDROID_ROOT}/external/tcpdump/print-llc.c
        ${ANDROID_ROOT}/external/tcpdump/print-lldp.c
        ${ANDROID_ROOT}/external/tcpdump/print-lmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-loopback.c
        ${ANDROID_ROOT}/external/tcpdump/print-lspping.c
        ${ANDROID_ROOT}/external/tcpdump/print-lwapp.c
        ${ANDROID_ROOT}/external/tcpdump/print-lwres.c
        ${ANDROID_ROOT}/external/tcpdump/print-m3ua.c
        ${ANDROID_ROOT}/external/tcpdump/print-macsec.c
        ${ANDROID_ROOT}/external/tcpdump/print-mobile.c
        ${ANDROID_ROOT}/external/tcpdump/print-mobility.c
        ${ANDROID_ROOT}/external/tcpdump/print-mpcp.c
        ${ANDROID_ROOT}/external/tcpdump/print-mpls.c
        ${ANDROID_ROOT}/external/tcpdump/print-mptcp.c
        ${ANDROID_ROOT}/external/tcpdump/print-msdp.c
        ${ANDROID_ROOT}/external/tcpdump/print-msnlb.c
        ${ANDROID_ROOT}/external/tcpdump/print-nflog.c
        ${ANDROID_ROOT}/external/tcpdump/print-nfs.c
        ${ANDROID_ROOT}/external/tcpdump/print-nsh.c
        ${ANDROID_ROOT}/external/tcpdump/print-ntp.c
        ${ANDROID_ROOT}/external/tcpdump/print-null.c
        ${ANDROID_ROOT}/external/tcpdump/print-olsr.c
        ${ANDROID_ROOT}/external/tcpdump/print-openflow-1.0.c
        ${ANDROID_ROOT}/external/tcpdump/print-openflow-1.3.c
        ${ANDROID_ROOT}/external/tcpdump/print-openflow.c
        ${ANDROID_ROOT}/external/tcpdump/print-ospf.c
        ${ANDROID_ROOT}/external/tcpdump/print-ospf6.c
        ${ANDROID_ROOT}/external/tcpdump/print-otv.c
        ${ANDROID_ROOT}/external/tcpdump/print-pflog.c
        ${ANDROID_ROOT}/external/tcpdump/print-pgm.c
        ${ANDROID_ROOT}/external/tcpdump/print-pim.c
        ${ANDROID_ROOT}/external/tcpdump/print-pktap.c
        ${ANDROID_ROOT}/external/tcpdump/print-ppi.c
        ${ANDROID_ROOT}/external/tcpdump/print-ppp.c
        ${ANDROID_ROOT}/external/tcpdump/print-pppoe.c
        ${ANDROID_ROOT}/external/tcpdump/print-pptp.c
        ${ANDROID_ROOT}/external/tcpdump/print-ptp.c
        ${ANDROID_ROOT}/external/tcpdump/print-radius.c
        ${ANDROID_ROOT}/external/tcpdump/print-raw.c
        ${ANDROID_ROOT}/external/tcpdump/print-realtek.c
        ${ANDROID_ROOT}/external/tcpdump/print-resp.c
        ${ANDROID_ROOT}/external/tcpdump/print-rip.c
        ${ANDROID_ROOT}/external/tcpdump/print-ripng.c
        ${ANDROID_ROOT}/external/tcpdump/print-rpki-rtr.c
        ${ANDROID_ROOT}/external/tcpdump/print-rsvp.c
        ${ANDROID_ROOT}/external/tcpdump/print-rt6.c
        ${ANDROID_ROOT}/external/tcpdump/print-rtsp.c
        ${ANDROID_ROOT}/external/tcpdump/print-rx.c
        ${ANDROID_ROOT}/external/tcpdump/print-sctp.c
        ${ANDROID_ROOT}/external/tcpdump/print-sflow.c
        ${ANDROID_ROOT}/external/tcpdump/print-sip.c
        ${ANDROID_ROOT}/external/tcpdump/print-sl.c
        ${ANDROID_ROOT}/external/tcpdump/print-sll.c
        ${ANDROID_ROOT}/external/tcpdump/print-slow.c
        ${ANDROID_ROOT}/external/tcpdump/print-smb.c
        ${ANDROID_ROOT}/external/tcpdump/print-smtp.c
        ${ANDROID_ROOT}/external/tcpdump/print-snmp.c
        ${ANDROID_ROOT}/external/tcpdump/print-someip.c
        ${ANDROID_ROOT}/external/tcpdump/print-ssh.c
        ${ANDROID_ROOT}/external/tcpdump/print-stp.c
        ${ANDROID_ROOT}/external/tcpdump/print-sunatm.c
        ${ANDROID_ROOT}/external/tcpdump/print-sunrpc.c
        ${ANDROID_ROOT}/external/tcpdump/print-symantec.c
        ${ANDROID_ROOT}/external/tcpdump/print-syslog.c
        ${ANDROID_ROOT}/external/tcpdump/print-tcp.c
        ${ANDROID_ROOT}/external/tcpdump/print-telnet.c
        ${ANDROID_ROOT}/external/tcpdump/print-tftp.c
        ${ANDROID_ROOT}/external/tcpdump/print-timed.c
        ${ANDROID_ROOT}/external/tcpdump/print-tipc.c
        ${ANDROID_ROOT}/external/tcpdump/print-token.c
        ${ANDROID_ROOT}/external/tcpdump/print-udld.c
        ${ANDROID_ROOT}/external/tcpdump/print-udp.c
        ${ANDROID_ROOT}/external/tcpdump/print-unsupported.c
        ${ANDROID_ROOT}/external/tcpdump/print-usb.c
        ${ANDROID_ROOT}/external/tcpdump/print-vjc.c
        ${ANDROID_ROOT}/external/tcpdump/print-vqp.c
        ${ANDROID_ROOT}/external/tcpdump/print-vrrp.c
        ${ANDROID_ROOT}/external/tcpdump/print-vsock.c
        ${ANDROID_ROOT}/external/tcpdump/print-vtp.c
        ${ANDROID_ROOT}/external/tcpdump/print-vxlan-gpe.c
        ${ANDROID_ROOT}/external/tcpdump/print-vxlan.c
        ${ANDROID_ROOT}/external/tcpdump/print-wb.c
        ${ANDROID_ROOT}/external/tcpdump/print-whois.c
        ${ANDROID_ROOT}/external/tcpdump/print-zep.c
        ${ANDROID_ROOT}/external/tcpdump/print-zephyr.c
        ${ANDROID_ROOT}/external/tcpdump/print-zeromq.c
        ${ANDROID_ROOT}/external/tcpdump/netdissect.c
        ${ANDROID_ROOT}/external/tcpdump/signature.c
        ${ANDROID_ROOT}/external/tcpdump/strtoaddr.c
        ${ANDROID_ROOT}/external/tcpdump/util-print.c
        ${ANDROID_ROOT}/external/tcpdump/smbutil.c
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
    add_executable(tcpdump ${SOURCE_FILES})

    target_include_directories(tcpdump
            PUBLIC
            ${ANDROID_ROOT}/external/tcpdump
    )

    # Add c/c++ flags
    target_compile_options(tcpdump PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(tcpdump
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libpcap
            libssl
    )
endif ()
#################### tcpdump ####################

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

project(netd)

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

#################### oemnetd_aidl_interface-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-ndk-source/gen/com/android/internal/net/IOemNetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-ndk-source/gen/com/android/internal/net/IOemNetdUnsolicitedEventListener.cpp
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
    add_library(oemnetd_aidl_interface-ndk SHARED ${SOURCE_FILES})

    target_include_directories(oemnetd_aidl_interface-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(oemnetd_aidl_interface-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(oemnetd_aidl_interface-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### oemnetd_aidl_interface-ndk ####################

#################### oemnetd_aidl_interface-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-source/gen/com/android/internal/net/IOemNetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-source/gen/com/android/internal/net/IOemNetdUnsolicitedEventListener.cpp
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
    add_library(oemnetd_aidl_interface-cpp SHARED ${SOURCE_FILES})

    target_include_directories(oemnetd_aidl_interface-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(oemnetd_aidl_interface-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(oemnetd_aidl_interface-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### oemnetd_aidl_interface-cpp ####################

#################### ndc ####################
set(FILES
        ${ANDROID_ROOT}/system/netd/server/ndc.cpp
        ${ANDROID_ROOT}/system/netd/server/UidRanges.cpp
        ${ANDROID_ROOT}/system/netd/server/NdcDispatcher.cpp
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
    add_executable(ndc ${SOURCE_FILES})

    target_include_directories(ndc
            PUBLIC
            ${ANDROID_ROOT}/system/netd/server
    )

    # Add c/c++ flags
    target_compile_options(ndc PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
    )

    target_link_libraries(ndc
            dnsresolver_aidl_interface-V7-cpp
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libnetd_client_headers
            libnetdutils
            libnetutils
            libutils
            netd_aidl_interface-V13-cpp
    )
endif ()
#################### ndc ####################

#################### netd_hw_service_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/netd/server/DummyNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/EventReporter.cpp
        ${ANDROID_ROOT}/system/netd/server/FwmarkServer.cpp
        ${ANDROID_ROOT}/system/netd/server/LocalNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsEventReporter.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsSdListener.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdCommand.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdHwAidlService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdHwService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdNativeService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetlinkHandler.cpp
        ${ANDROID_ROOT}/system/netd/server/Network.cpp
        ${ANDROID_ROOT}/system/netd/server/NetworkController.cpp
        ${ANDROID_ROOT}/system/netd/server/OemNetdListener.cpp
        ${ANDROID_ROOT}/system/netd/server/PhysicalNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/PppController.cpp
        ${ANDROID_ROOT}/system/netd/server/Process.cpp
        ${ANDROID_ROOT}/system/netd/server/UnreachableNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/VirtualNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/oem_iptables_hook.cpp
        ${ANDROID_ROOT}/system/netd/server/Controllers.cpp
        ${ANDROID_ROOT}/system/netd/server/aidl-fuzzers/NetdHwAidlServiceFuzzer.cpp
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
    add_executable(netd_hw_service_fuzzer ${SOURCE_FILES})

    target_include_directories(netd_hw_service_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/netd/server
            ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_hw_service_fuzzer PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
    )

    target_link_libraries(netd_hw_service_fuzzer
            android.system.net.netd-V1-ndk
            android.system.net.netd-1.0
            android.system.net.netd-1.1
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libhidlbase
            libip_checksum
            liblog
            libm
            libmdnssd
            libnetd_resolv
            libnetd_server
            libnetd_updatable
            libnetdutils
            libnetutils
            libpcap
            libprocessgroup
            libselinux
            libsysutils
            libtcutils
            libutils
            mdns_aidl_interface-V1-cpp
            netd_aidl_interface-V13-cpp
            netd_event_listener_interface-V1-cpp
            oemnetd_aidl_interface-cpp
    )
endif ()
#################### netd_hw_service_fuzzer ####################

#################### netd_native_service_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/netd/server/DummyNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/EventReporter.cpp
        ${ANDROID_ROOT}/system/netd/server/FwmarkServer.cpp
        ${ANDROID_ROOT}/system/netd/server/LocalNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsEventReporter.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsSdListener.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdCommand.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdHwAidlService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdHwService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdNativeService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetlinkHandler.cpp
        ${ANDROID_ROOT}/system/netd/server/Network.cpp
        ${ANDROID_ROOT}/system/netd/server/NetworkController.cpp
        ${ANDROID_ROOT}/system/netd/server/OemNetdListener.cpp
        ${ANDROID_ROOT}/system/netd/server/PhysicalNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/PppController.cpp
        ${ANDROID_ROOT}/system/netd/server/Process.cpp
        ${ANDROID_ROOT}/system/netd/server/UnreachableNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/VirtualNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/oem_iptables_hook.cpp
        ${ANDROID_ROOT}/system/netd/server/Controllers.cpp
        ${ANDROID_ROOT}/system/netd/server/aidl-fuzzers/NetdNativeServiceFuzzer.cpp
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
    add_executable(netd_native_service_fuzzer ${SOURCE_FILES})

    target_include_directories(netd_native_service_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/netd/server
            ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_native_service_fuzzer PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
    )

    target_link_libraries(netd_native_service_fuzzer
            android.system.net.netd-V1-ndk
            android.system.net.netd-1.0
            android.system.net.netd-1.1
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libhidlbase
            libip_checksum
            liblog
            libm
            libmdnssd
            libnetd_resolv
            libnetd_server
            libnetd_updatable
            libnetdutils
            libnetutils
            libpcap
            libprocessgroup
            libselinux
            libsysutils
            libtcutils
            libutils
            mdns_aidl_interface-V1-cpp
            netd_aidl_interface-V13-cpp
            netd_event_listener_interface-V1-cpp
            oemnetd_aidl_interface-cpp
    )
endif ()
#################### netd_native_service_fuzzer ####################

#################### oemnetd_aidl_interface-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-ndk_platform-source/gen/com/android/internal/net/IOemNetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-ndk_platform-source/gen/com/android/internal/net/IOemNetdUnsolicitedEventListener.cpp
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
    add_library(oemnetd_aidl_interface-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(oemnetd_aidl_interface-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(oemnetd_aidl_interface-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(oemnetd_aidl_interface-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### oemnetd_aidl_interface-ndk_platform ####################

#################### oemnetd_aidl_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-analyzer-source/gen/com/android/internal/net/IOemNetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-analyzer-source/gen/com/android/internal/net/IOemNetdUnsolicitedEventListener.cpp
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
    add_executable(oemnetd_aidl_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(oemnetd_aidl_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/system/netd/server
            ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(oemnetd_aidl_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(oemnetd_aidl_interface-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            oemnetd_aidl_interface-cpp
    )
endif ()
#################### oemnetd_aidl_interface-cpp-analyzer ####################

#################### mdns_service_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/netd/server/DummyNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/EventReporter.cpp
        ${ANDROID_ROOT}/system/netd/server/FwmarkServer.cpp
        ${ANDROID_ROOT}/system/netd/server/LocalNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsEventReporter.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsSdListener.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdCommand.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdHwAidlService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdHwService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdNativeService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetlinkHandler.cpp
        ${ANDROID_ROOT}/system/netd/server/Network.cpp
        ${ANDROID_ROOT}/system/netd/server/NetworkController.cpp
        ${ANDROID_ROOT}/system/netd/server/OemNetdListener.cpp
        ${ANDROID_ROOT}/system/netd/server/PhysicalNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/PppController.cpp
        ${ANDROID_ROOT}/system/netd/server/Process.cpp
        ${ANDROID_ROOT}/system/netd/server/UnreachableNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/VirtualNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/oem_iptables_hook.cpp
        ${ANDROID_ROOT}/system/netd/server/Controllers.cpp
        ${ANDROID_ROOT}/system/netd/server/aidl-fuzzers/MDnsServiceFuzzer.cpp
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
    add_executable(mdns_service_fuzzer ${SOURCE_FILES})

    target_include_directories(mdns_service_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/netd/server
            ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mdns_service_fuzzer PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
    )

    target_link_libraries(mdns_service_fuzzer
            android.system.net.netd-V1-ndk
            android.system.net.netd-1.0
            android.system.net.netd-1.1
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libhidlbase
            libip_checksum
            liblog
            libm
            libmdnssd
            libnetd_resolv
            libnetd_server
            libnetd_updatable
            libnetdutils
            libnetutils
            libpcap
            libprocessgroup
            libselinux
            libsysutils
            libtcutils
            libutils
            mdns_aidl_interface-V1-cpp
            netd_aidl_interface-V13-cpp
            netd_event_listener_interface-V1-cpp
            oemnetd_aidl_interface-cpp
    )
endif ()
#################### mdns_service_fuzzer ####################

#################### libnetd_server ####################
set(FILES
        ${ANDROID_ROOT}/system/netd/server/BandwidthController.cpp
        ${ANDROID_ROOT}/system/netd/server/Controllers.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdConstants.cpp
        ${ANDROID_ROOT}/system/netd/server/FirewallController.cpp
        ${ANDROID_ROOT}/system/netd/server/IdletimerController.cpp
        ${ANDROID_ROOT}/system/netd/server/InterfaceController.cpp
        ${ANDROID_ROOT}/system/netd/server/IptablesRestoreController.cpp
        ${ANDROID_ROOT}/system/netd/server/NFLogListener.cpp
        ${ANDROID_ROOT}/system/netd/server/NetlinkCommands.cpp
        ${ANDROID_ROOT}/system/netd/server/NetlinkManager.cpp
        ${ANDROID_ROOT}/system/netd/server/RouteController.cpp
        ${ANDROID_ROOT}/system/netd/server/SockDiag.cpp
        ${ANDROID_ROOT}/system/netd/server/StrictController.cpp
        ${ANDROID_ROOT}/system/netd/server/TcpSocketMonitor.cpp
        ${ANDROID_ROOT}/system/netd/server/TetherController.cpp
        ${ANDROID_ROOT}/system/netd/server/UidRanges.cpp
        ${ANDROID_ROOT}/system/netd/server/WakeupController.cpp
        ${ANDROID_ROOT}/system/netd/server/XfrmController.cpp
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
    add_library(libnetd_server SHARED ${SOURCE_FILES})

    target_include_directories(libnetd_server
            PUBLIC
            ${ANDROID_ROOT}/system/netd/server/binder
            ${ANDROID_ROOT}/system/netd/server
    )

    # Add c/c++ flags
    target_compile_options(libnetd_server PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
    )

    target_link_libraries(libnetd_server
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libip_checksum
            libm
            libnetdutils
            libnetutils
            libpcap
            libssl
            libsysutils
            libtcutils
            netd_aidl_interface-V13-cpp
            netd_event_listener_interface-V1-cpp
    )
endif ()
#################### libnetd_server ####################

#################### netd ####################
set(FILES
        ${ANDROID_ROOT}/system/netd/server/DummyNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/EventReporter.cpp
        ${ANDROID_ROOT}/system/netd/server/FwmarkServer.cpp
        ${ANDROID_ROOT}/system/netd/server/LocalNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsEventReporter.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsSdListener.cpp
        ${ANDROID_ROOT}/system/netd/server/MDnsService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdCommand.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdHwAidlService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdHwService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetdNativeService.cpp
        ${ANDROID_ROOT}/system/netd/server/NetlinkHandler.cpp
        ${ANDROID_ROOT}/system/netd/server/Network.cpp
        ${ANDROID_ROOT}/system/netd/server/NetworkController.cpp
        ${ANDROID_ROOT}/system/netd/server/OemNetdListener.cpp
        ${ANDROID_ROOT}/system/netd/server/PhysicalNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/PppController.cpp
        ${ANDROID_ROOT}/system/netd/server/Process.cpp
        ${ANDROID_ROOT}/system/netd/server/UnreachableNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/VirtualNetwork.cpp
        ${ANDROID_ROOT}/system/netd/server/oem_iptables_hook.cpp
        ${ANDROID_ROOT}/system/netd/server/main.cpp
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
    add_executable(netd ${SOURCE_FILES})

    target_include_directories(netd
            PUBLIC
            ${ANDROID_ROOT}/system/netd/server
            ${OUT_INTERMEDIATES_ROOT}/system/netd/server/oemnetd_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd PRIVATE
            -Wall
            -Werror
            -Werror=implicit-fallthrough
            -Werror=sometimes-uninitialized
            -Werror=conditional-uninitialized
    )

    target_link_libraries(netd
            android.system.net.netd-V1-ndk
            android.system.net.netd-1.0
            android.system.net.netd-1.1
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            libip_checksum
            liblog
            libm
            libmdnssd
            libnetd_resolv
            libnetd_server
            libnetd_updatable
            libnetdutils
            libnetutils
            libpcap
            libprocessgroup
            libselinux
            libsysutils
            libtcutils
            libutils
            mdns_aidl_interface-V1-cpp
            netd_aidl_interface-V13-cpp
            netd_event_listener_interface-V1-cpp
            oemnetd_aidl_interface-cpp
    )
endif ()
#################### netd ####################

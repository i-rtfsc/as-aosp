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

project(wificond)

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

#################### libwificond_ipc ####################
set(FILES
        ${ANDROID_ROOT}/system/connectivity/wificond/ipc_constants.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/IApInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/IApInterfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/IClientInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/IInterfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/IPnoScanEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/IScanEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/ISendMgmtFrameEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/IWificond.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/IWificondEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl/android/net/wifi/nl80211/IWifiScannerImpl.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/client/native_wifi_client.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/device_wiphy_capabilities.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/channel_settings.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/hidden_network.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/pno_network.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/pno_settings.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/radio_chain_info.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/scan_result.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/single_scan_settings.cpp
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
    add_library(libwificond_ipc SHARED ${SOURCE_FILES})

    target_include_directories(libwificond_ipc
            PUBLIC
            ${ANDROID_ROOT}/system/connectivity/wificond
            ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    target_link_libraries(libwificond_ipc
            android.hardware.net.nlinterceptor-V1-ndk
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnlinterceptor
    )
endif ()
#################### libwificond_ipc ####################

#################### libwificond_nl ####################
set(FILES
        ${ANDROID_ROOT}/system/connectivity/wificond/net/mlme_event.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/net/netlink_manager.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/net/netlink_utils.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/net/nl80211_attribute.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/net/nl80211_packet.cpp
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
    add_library(libwificond_nl SHARED ${SOURCE_FILES})

    target_include_directories(libwificond_nl
            PUBLIC
            ${ANDROID_ROOT}/system/connectivity/wificond
    )

    target_link_libraries(libwificond_nl
            android.hardware.automotive.can-libnetdevice
            android.hardware.net.nlinterceptor-V1-ndk
            libbase
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnl++
            libnlinterceptor
    )
endif ()
#################### libwificond_nl ####################

#################### wificond ####################
set(FILES
        ${ANDROID_ROOT}/system/connectivity/wificond/main.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/wifi_keystore_hal_connector.cpp
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
    add_executable(wificond ${SOURCE_FILES})

    target_include_directories(wificond
            PUBLIC
            ${ANDROID_ROOT}/system/connectivity/wificond
            ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    target_link_libraries(wificond
            android.hardware.net.nlinterceptor-V1-ndk
            android.security.legacykeystore-ndk
            android.system.keystore2-V3-ndk
            android.system.wifi.keystore-1.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libhidlbase
            libm
            libminijail
            libnlinterceptor
            libprocessgroup
            libssl
            libutils
            libvndksupport
            libwifi-system-iface
            libwificond
            libwifikeystorehal
    )
endif ()
#################### wificond ####################

#################### libwificond ####################
set(FILES
        ${ANDROID_ROOT}/system/connectivity/wificond/ap_interface_binder.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/ap_interface_impl.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/client_interface_binder.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/client_interface_impl.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/device_wiphy_capabilities.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/logging_utils.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/client/native_wifi_client.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/channel_settings.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/hidden_network.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/pno_network.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/pno_settings.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/radio_chain_info.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/scan_result.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/single_scan_settings.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/scan_utils.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/scanning/scanner_impl.cpp
        ${ANDROID_ROOT}/system/connectivity/wificond/server.cpp
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
    add_library(libwificond SHARED ${SOURCE_FILES})

    target_include_directories(libwificond
            PUBLIC
            ${ANDROID_ROOT}/system/connectivity/wificond
            ${OUT_INTERMEDIATES_ROOT}/system/connectivity/wificond/libwificond_ipc/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    target_link_libraries(libwificond
            android.hardware.net.nlinterceptor-V1-ndk
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnlinterceptor
            libutils
            libwifi-system-iface
            libwificond_event_loop
            libwificond_ipc
            libwificond_nl
    )
endif ()
#################### libwificond ####################

#################### libwificond_event_loop ####################
set(FILES
        ${ANDROID_ROOT}/system/connectivity/wificond/looper_backed_event_loop.cpp
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
    add_library(libwificond_event_loop SHARED ${SOURCE_FILES})

    target_include_directories(libwificond_event_loop
            PUBLIC
            ${ANDROID_ROOT}/system/connectivity/wificond
    )

    target_link_libraries(libwificond_event_loop
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
    )
endif ()
#################### libwificond_event_loop ####################

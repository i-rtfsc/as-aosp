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

project(android.hardware.wifi-service)

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

#################### android.hardware.wifi-service-lazy ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/service.cpp
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
    add_executable(android.hardware.wifi-service-lazy ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-service-lazy
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default
    )

    target_link_libraries(android.hardware.wifi-service-lazy
            android.hardware.wifi-V1-ndk
            android.hardware.wifi-service-lib
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libnl
            libutils
            libwifi-hal
            libwifi-system-iface
            libxml2
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.wifi-service-lazy ####################

#################### android.hardware.wifi-service-lib ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/aidl_struct_util.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/aidl_sync_util.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/ringbuffer.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_ap_iface.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_chip.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_feature_flags.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_iface_util.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_legacy_hal.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_legacy_hal_factory.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_legacy_hal_stubs.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_mode_controller.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_nan_iface.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_p2p_iface.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_rtt_controller.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_sta_iface.cpp
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/wifi_status_util.cpp
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
    add_library(android.hardware.wifi-service-lib SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-service-lib
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default
    )

    target_link_libraries(android.hardware.wifi-service-lib
            android.hardware.wifi-V1-ndk
            libbase
            libbinder_ndk
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libnl
            libutils
            libwifi-hal
            libwifi-system-iface
            libxml2
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.wifi-service-lib ####################

#################### android.hardware.wifi-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default/service.cpp
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
    add_executable(android.hardware.wifi-service ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl/default
    )

    target_link_libraries(android.hardware.wifi-service
            android.hardware.wifi-V1-ndk
            android.hardware.wifi-service-lib
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libnl
            libutils
            libwifi-hal
            libwifi-system-iface
            libxml2
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.wifi-service ####################

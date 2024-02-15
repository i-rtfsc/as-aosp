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

project(mdns_aidl_interface-cpp-analyzer)

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

#################### netd_aidl_interface-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V4-ndk_platform
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
#################### netd_aidl_interface-V4-ndk_platform ####################

#################### netd_aidl_interface-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V1-ndk ####################

#################### netd_aidl_interface-V8-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V8-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V8-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V8-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V8-ndk_platform
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
#################### netd_aidl_interface-V8-ndk_platform ####################

#################### netd_event_listener_interface-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V3-ndk_platform-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_library(netd_event_listener_interface-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-V3-ndk_platform
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
#################### netd_event_listener_interface-V3-ndk_platform ####################

#################### netd_event_listener_interface-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V2-ndk-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_library(netd_event_listener_interface-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_event_listener_interface-V2-ndk ####################

#################### mdns_aidl_interface-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-cpp-source/gen/android/net/mdns/aidl/DiscoveryInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-cpp-source/gen/android/net/mdns/aidl/GetAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-cpp-source/gen/android/net/mdns/aidl/IMDns.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-cpp-source/gen/android/net/mdns/aidl/IMDnsEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-cpp-source/gen/android/net/mdns/aidl/RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-cpp-source/gen/android/net/mdns/aidl/ResolutionInfo.cpp
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
    add_library(mdns_aidl_interface-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(mdns_aidl_interface-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mdns_aidl_interface-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mdns_aidl_interface-V2-cpp
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
#################### mdns_aidl_interface-V2-cpp ####################

#################### netd_aidl_interface-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V3-cpp
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
#################### netd_aidl_interface-V3-cpp ####################

#################### netd_aidl_interface-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V4-ndk ####################

#################### netd_aidl_interface-V6-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V6-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V6-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V6-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V6-cpp
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
#################### netd_aidl_interface-V6-cpp ####################

#################### netd_event_listener_interface-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V3-ndk-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_library(netd_event_listener_interface-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_event_listener_interface-V3-ndk ####################

#################### netd_aidl_interface-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-cpp-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V2-cpp
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
#################### netd_aidl_interface-V2-cpp ####################

#################### mdns_aidl_interface-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk-source/gen/android/net/mdns/aidl/DiscoveryInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk-source/gen/android/net/mdns/aidl/GetAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk-source/gen/android/net/mdns/aidl/IMDns.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk-source/gen/android/net/mdns/aidl/IMDnsEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk-source/gen/android/net/mdns/aidl/RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk-source/gen/android/net/mdns/aidl/ResolutionInfo.cpp
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
    add_library(mdns_aidl_interface-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(mdns_aidl_interface-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mdns_aidl_interface-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mdns_aidl_interface-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### mdns_aidl_interface-V1-ndk ####################

#################### mdns_aidl_interface-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk_platform-source/gen/android/net/mdns/aidl/DiscoveryInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk_platform-source/gen/android/net/mdns/aidl/GetAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk_platform-source/gen/android/net/mdns/aidl/IMDns.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk_platform-source/gen/android/net/mdns/aidl/IMDnsEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk_platform-source/gen/android/net/mdns/aidl/RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk_platform-source/gen/android/net/mdns/aidl/ResolutionInfo.cpp
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
    add_library(mdns_aidl_interface-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(mdns_aidl_interface-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mdns_aidl_interface-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mdns_aidl_interface-V1-ndk_platform
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
#################### mdns_aidl_interface-V1-ndk_platform ####################

#################### netd_event_listener_interface-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V1-cpp-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_library(netd_event_listener_interface-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-V1-cpp
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
#################### netd_event_listener_interface-V1-cpp ####################

#################### netd_aidl_interface-V5-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V5-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V5-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V5-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V5-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V5-ndk ####################

#################### netd_aidl_interface-V7-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V7-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V7-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V7-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V7-cpp
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
#################### netd_aidl_interface-V7-cpp ####################

#################### netd_aidl_interface-V11-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V11-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V11-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V11-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V11-ndk_platform
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
#################### netd_aidl_interface-V11-ndk_platform ####################

#################### netd_aidl_interface-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V3-ndk_platform
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
#################### netd_aidl_interface-V3-ndk_platform ####################

#################### netd_aidl_interface-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V3-ndk ####################

#################### netd_aidl_interface-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-cpp-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V1-cpp
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
#################### netd_aidl_interface-V1-cpp ####################

#################### mdns_aidl_interface-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk-source/gen/android/net/mdns/aidl/DiscoveryInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk-source/gen/android/net/mdns/aidl/GetAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk-source/gen/android/net/mdns/aidl/IMDns.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk-source/gen/android/net/mdns/aidl/IMDnsEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk-source/gen/android/net/mdns/aidl/RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk-source/gen/android/net/mdns/aidl/ResolutionInfo.cpp
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
    add_library(mdns_aidl_interface-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(mdns_aidl_interface-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mdns_aidl_interface-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mdns_aidl_interface-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### mdns_aidl_interface-V2-ndk ####################

#################### netd_event_listener_interface-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V2-cpp-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_library(netd_event_listener_interface-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-V2-cpp
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
#################### netd_event_listener_interface-V2-cpp ####################

#################### netd_aidl_interface-V6-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V6-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V6-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V6-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V6-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V6-ndk ####################

#################### netd_aidl_interface-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V2-ndk_platform
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
#################### netd_aidl_interface-V2-ndk_platform ####################

#################### netd_aidl_interface-V4-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V4-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V4-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V4-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V4-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V4-cpp
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
#################### netd_aidl_interface-V4-cpp ####################

#################### netd_aidl_interface-V10-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V10-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V10-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V10-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V10-ndk_platform
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
#################### netd_aidl_interface-V10-ndk_platform ####################

#################### netd_aidl_interface-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V2-ndk ####################

#################### netd_event_listener_interface-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V1-ndk-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_library(netd_event_listener_interface-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_event_listener_interface-V1-ndk ####################

#################### netd_aidl_interface-V5-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V5-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V5-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V5-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V5-ndk_platform
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
#################### netd_aidl_interface-V5-ndk_platform ####################

#################### netd_aidl_interface-V9-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V9-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V9-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V9-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V9-ndk_platform
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
#################### netd_aidl_interface-V9-ndk_platform ####################

#################### mdns_aidl_interface-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-cpp-source/gen/android/net/mdns/aidl/DiscoveryInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-cpp-source/gen/android/net/mdns/aidl/GetAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-cpp-source/gen/android/net/mdns/aidl/IMDns.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-cpp-source/gen/android/net/mdns/aidl/IMDnsEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-cpp-source/gen/android/net/mdns/aidl/RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-cpp-source/gen/android/net/mdns/aidl/ResolutionInfo.cpp
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
    add_library(mdns_aidl_interface-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(mdns_aidl_interface-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mdns_aidl_interface-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mdns_aidl_interface-V1-cpp
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
#################### mdns_aidl_interface-V1-cpp ####################

#################### netd_event_listener_interface-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V2-ndk_platform-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_library(netd_event_listener_interface-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-V2-ndk_platform
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
#################### netd_event_listener_interface-V2-ndk_platform ####################

#################### netd_event_listener_interface-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V3-cpp-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_library(netd_event_listener_interface-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-V3-cpp
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
#################### netd_event_listener_interface-V3-cpp ####################

#################### netd_aidl_interface-V7-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V7-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V7-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V7-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V7-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V7-ndk ####################

#################### netd_aidl_interface-V5-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V5-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V5-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V5-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V5-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V5-cpp
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
#################### netd_aidl_interface-V5-cpp ####################

#################### netd_aidl_interface-V11-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V11-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V11-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V11-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V11-cpp
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
#################### netd_aidl_interface-V11-cpp ####################

#################### mdns_aidl_interface-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk_platform-source/gen/android/net/mdns/aidl/DiscoveryInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk_platform-source/gen/android/net/mdns/aidl/GetAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk_platform-source/gen/android/net/mdns/aidl/IMDns.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk_platform-source/gen/android/net/mdns/aidl/IMDnsEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk_platform-source/gen/android/net/mdns/aidl/RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk_platform-source/gen/android/net/mdns/aidl/ResolutionInfo.cpp
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
    add_library(mdns_aidl_interface-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(mdns_aidl_interface-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mdns_aidl_interface-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mdns_aidl_interface-V2-ndk_platform
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
#################### mdns_aidl_interface-V2-ndk_platform ####################

#################### netd_aidl_interface-V13-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V13-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V13-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V13-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V13-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V13-ndk ####################

#################### netd_aidl_interface-V14-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V14-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V14-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V14-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V14-cpp
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
#################### netd_aidl_interface-V14-cpp ####################

#################### netd_aidl_interface-V12-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V12-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V12-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V12-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V12-ndk_platform
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
#################### netd_aidl_interface-V12-ndk_platform ####################

#################### netd_aidl_interface-V9-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V9-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V9-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V9-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V9-cpp
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
#################### netd_aidl_interface-V9-cpp ####################

#################### netd_aidl_interface-V10-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V10-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V10-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V10-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V10-cpp
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
#################### netd_aidl_interface-V10-cpp ####################

#################### netd_aidl_interface-V12-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V12-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V12-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V12-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V12-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V12-ndk ####################

#################### netd_aidl_interface-V7-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V7-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V7-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V7-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V7-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V7-ndk_platform
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
#################### netd_aidl_interface-V7-ndk_platform ####################

#################### netd_aidl_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_executable(netd_aidl_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/libs/net/common/netd
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            netd_aidl_interface-V14-cpp
    )
endif ()
#################### netd_aidl_interface-cpp-analyzer ####################

#################### netd_aidl_interface-V8-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V8-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V8-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V8-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V8-cpp
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
#################### netd_aidl_interface-V8-cpp ####################

#################### netd_aidl_interface-V13-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V13-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V13-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V13-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V13-cpp
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
#################### netd_aidl_interface-V13-cpp ####################

#################### netd_aidl_interface-V6-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V6-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V6-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V6-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V6-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V6-ndk_platform
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
#################### netd_aidl_interface-V6-ndk_platform ####################

#################### netd_event_listener_interface-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V1-ndk_platform-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_library(netd_event_listener_interface-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-V1-ndk_platform
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
#################### netd_event_listener_interface-V1-ndk_platform ####################

#################### netd_aidl_interface-V11-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V11-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V11-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V11-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V11-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V11-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V11-ndk ####################

#################### netd_aidl_interface-V14-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V14-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V14-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V14-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V14-ndk_platform
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
#################### netd_aidl_interface-V14-ndk_platform ####################

#################### netd_aidl_interface-V9-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V9-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V9-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V9-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V9-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V9-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V9-ndk ####################

#################### netd_aidl_interface-V14-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V14-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V14-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V14-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V14-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V14-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V14-ndk ####################

#################### netd_aidl_interface-V12-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V12-cpp SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V12-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V12-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V12-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V12-cpp
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
#################### netd_aidl_interface-V12-cpp ####################

#################### netd_event_listener_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-cpp-analyzer-source/gen/android/net/metrics/INetdEventListener.cpp
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
    add_executable(netd_event_listener_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(netd_event_listener_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/libs/net/common/netd
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-V3-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_event_listener_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_event_listener_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_event_listener_interface-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            netd_event_listener_interface-V3-cpp
    )
endif ()
#################### netd_event_listener_interface-cpp-analyzer ####################

#################### netd_aidl_interface-V10-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V10-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V10-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V10-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V10-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V10-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V10-ndk ####################

#################### netd_aidl_interface-V13-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/IpSecMigrateInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V13-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V13-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V13-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V13-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V13-ndk_platform
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
#################### netd_aidl_interface-V13-ndk_platform ####################

#################### netd_aidl_interface-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk_platform-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk_platform-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk_platform-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk_platform-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk_platform-source/gen/android/net/UidRangeParcel.cpp
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
    add_library(netd_aidl_interface-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V1-ndk_platform
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
#################### netd_aidl_interface-V1-ndk_platform ####################

#################### netd_aidl_interface-V8-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/INetd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/INetdUnsolicitedEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/InterfaceConfigurationParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/MarkMaskParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/NativeNetworkConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/NativeNetworkType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/NativeVpnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/RouteInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/TetherConfigParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/TetherOffloadRuleParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/TetherStatsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/UidRangeParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/android/net/netd/aidl/NativeUidRangeConfig.cpp
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
    add_library(netd_aidl_interface-V8-ndk SHARED ${SOURCE_FILES})

    target_include_directories(netd_aidl_interface-V8-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/netd_aidl_interface-V8-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(netd_aidl_interface-V8-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(netd_aidl_interface-V8-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### netd_aidl_interface-V8-ndk ####################

#################### mdns_aidl_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-cpp-analyzer-source/gen/android/net/mdns/aidl/DiscoveryInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-cpp-analyzer-source/gen/android/net/mdns/aidl/GetAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-cpp-analyzer-source/gen/android/net/mdns/aidl/IMDns.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-cpp-analyzer-source/gen/android/net/mdns/aidl/IMDnsEventListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-cpp-analyzer-source/gen/android/net/mdns/aidl/RegistrationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-cpp-analyzer-source/gen/android/net/mdns/aidl/ResolutionInfo.cpp
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
    add_executable(mdns_aidl_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(mdns_aidl_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/libs/net/common/netd
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/libs/net/common/netd/mdns_aidl_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mdns_aidl_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mdns_aidl_interface-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            mdns_aidl_interface-V2-cpp
    )
endif ()
#################### mdns_aidl_interface-cpp-analyzer ####################

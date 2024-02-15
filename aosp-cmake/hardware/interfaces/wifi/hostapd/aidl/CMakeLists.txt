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

project(android.hardware.wifi.hostapd-cpp-analyzer)

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

#################### android.hardware.wifi.hostapd-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/ApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/BandMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/ChannelBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/ChannelParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/ClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/EncryptionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/FrequencyRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/Generation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/HostapdStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/HwModeParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/IHostapd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/IHostapdCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/Ieee80211ReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/IfaceParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/NetworkParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/android/hardware/wifi/hostapd/ParamSizeLimits.cpp
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
    add_library(android.hardware.wifi.hostapd-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.hostapd-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.hostapd-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.hostapd-V2-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.wifi.hostapd-V2-cpp ####################

#################### android.hardware.wifi.hostapd-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/ApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/BandMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/ChannelBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/ChannelParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/ClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/EncryptionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/FrequencyRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/Generation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/HostapdStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/HwModeParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/IHostapd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/IHostapdCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/Ieee80211ReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/IfaceParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/NetworkParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/android/hardware/wifi/hostapd/ParamSizeLimits.cpp
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
    add_library(android.hardware.wifi.hostapd-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.hostapd-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.hostapd-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.hostapd-V1-ndk
            libandroid_aidltrace
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
    )
endif ()
#################### android.hardware.wifi.hostapd-V1-ndk ####################

#################### android.hardware.wifi.hostapd-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/ApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/BandMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/ChannelBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/ChannelParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/ClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/EncryptionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/FrequencyRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/Generation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/HostapdStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/HwModeParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/IHostapd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/IHostapdCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/Ieee80211ReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/IfaceParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/NetworkParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/android/hardware/wifi/hostapd/ParamSizeLimits.cpp
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
    add_library(android.hardware.wifi.hostapd-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.hostapd-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.hostapd-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.hostapd-V1-ndk_platform
            libandroid_aidltrace
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
    )
endif ()
#################### android.hardware.wifi.hostapd-V1-ndk_platform ####################

#################### android.hardware.wifi.hostapd-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/ApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/BandMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/ChannelBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/ChannelParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/ClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/EncryptionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/FrequencyRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/Generation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/HostapdStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/HwModeParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/IHostapd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/IHostapdCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/Ieee80211ReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/IfaceParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/NetworkParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/android/hardware/wifi/hostapd/ParamSizeLimits.cpp
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
    add_library(android.hardware.wifi.hostapd-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.hostapd-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.hostapd-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.hostapd-V2-ndk
            libandroid_aidltrace
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
    )
endif ()
#################### android.hardware.wifi.hostapd-V2-ndk ####################

#################### android.hardware.wifi.hostapd-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/ApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/BandMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/ChannelBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/ChannelParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/ClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/EncryptionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/FrequencyRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/Generation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/HostapdStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/HwModeParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/IHostapd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/IHostapdCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/Ieee80211ReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/IfaceParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/NetworkParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/android/hardware/wifi/hostapd/ParamSizeLimits.cpp
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
    add_library(android.hardware.wifi.hostapd-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.hostapd-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.hostapd-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.hostapd-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.wifi.hostapd-V1-cpp ####################

#################### android.hardware.wifi.hostapd-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/ApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/BandMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/ChannelBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/ChannelParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/ClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/EncryptionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/FrequencyRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/Generation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/HostapdStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/HwModeParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/IHostapd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/IHostapdCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/Ieee80211ReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/IfaceParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/NetworkParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/android/hardware/wifi/hostapd/ParamSizeLimits.cpp
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
    add_library(android.hardware.wifi.hostapd-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.hostapd-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.hostapd-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.hostapd-V2-ndk_platform
            libandroid_aidltrace
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
    )
endif ()
#################### android.hardware.wifi.hostapd-V2-ndk_platform ####################

#################### android.hardware.wifi.hostapd-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/ApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/BandMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/ChannelBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/ChannelParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/ClientInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/EncryptionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/FrequencyRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/Generation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/HostapdStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/HwModeParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/IHostapd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/IHostapdCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/Ieee80211ReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/IfaceParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/NetworkParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/android/hardware/wifi/hostapd/ParamSizeLimits.cpp
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
    add_executable(android.hardware.wifi.hostapd-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.hostapd-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/wifi/hostapd/aidl
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/hostapd/aidl/android.hardware.wifi.hostapd-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.hostapd-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.hostapd-cpp-analyzer
            aidl-analyzer-main
            android.hardware.wifi.hostapd-V2-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.wifi.hostapd-cpp-analyzer ####################

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

project(vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform)

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

#################### vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/MsimPreference.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V2-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V4-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V4-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V8-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V8-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V8-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V8-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V8-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V8-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradioconfig-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/CiwlanCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataRecommendation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/MsimPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimTypeInfo.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            vendor.qti.hardware.radio.common-V1-ndk
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V4-ndk ####################

#################### vendor.qti.hardware.radio.qtiradioconfig-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V1-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V10-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V10-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V10-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V10-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V10-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V10-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V9-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V9-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V9-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V9-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V9-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V9-ndk ####################

#################### vendor.qti.hardware.radio.qtiradioconfig-V5-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/CiwlanCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataRecommendation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/MsimPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimTypeInfo.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V5-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V5-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V5-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V5-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            vendor.qti.hardware.radio.common-V1-ndk
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V5-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V3-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/CiwlanCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataRecommendation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/MsimPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimTypeInfo.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            vendor.qti.hardware.radio.common-V1-ndk_platform
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V5-ndk_platform ####################

#################### vendor.qti.hardware.radio.common-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V2-ndk-source/gen/vendor/qti/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V2-ndk-source/gen/vendor/qti/hardware/radio/RadioResponseInfo.cpp
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
    add_library(vendor.qti.hardware.radio.common-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.common-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.common-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.common-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.common-V2-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V11-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/PersoUnlockStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V11-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V11-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V11-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V11-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V11-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V8-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V8-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V8-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V8-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V8-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V8-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V8-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V2-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V12-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/PersoUnlockStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V12-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V12-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V12-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V12-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V12-ndk_platform ####################

#################### vendor.qti.hardware.radio.common-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/RadioResponseInfo.cpp
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
    add_library(vendor.qti.hardware.radio.common-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.common-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.common-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.common-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.common-V1-ndk_platform ####################

#################### vendor.qti.hardware.radio.common-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk-source/gen/vendor/qti/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk-source/gen/vendor/qti/hardware/radio/RadioResponseInfo.cpp
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
    add_library(vendor.qti.hardware.radio.common-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.common-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.common-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.common-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.common-V1-ndk ####################

#################### vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/CiwlanCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataRecommendation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/DualDataSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/MsimPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimTypeInfo.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V1-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            vendor.qti.hardware.radio.common-V1-ndk_platform
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V4-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V12-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/PersoUnlockStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V12-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V12-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V12-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V12-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V12-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V12-ndk ####################

#################### vendor.qti.hardware.radio.qtiradioconfig-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/MsimPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimTypeInfo.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V3-ndk ####################

#################### vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/MsimPreference.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/SimTypeInfo.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V3-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V5-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V5-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V5-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V5-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V5-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V5-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V9-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V9-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V9-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V9-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V9-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V9-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V9-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradioconfig-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradioconfig/MsimPreference.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V2-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V3-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V10-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V10-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V10-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V10-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V10-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V10-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V10-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V6-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V6-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V6-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V6-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V6-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V6-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V7-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V7-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V7-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V7-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V7-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V7-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradioconfig/IQtiRadioConfigResponse.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradioconfig-V1-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V2-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V7-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V7-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V7-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V7-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V7-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V7-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V7-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V6-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V6-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V6-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V6-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V6-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V6-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V6-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V1-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V4-ndk ####################

#################### vendor.qti.hardware.radio.qtiradio-V5-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V5-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V5-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V5-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V5-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V5-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V5-ndk ####################

#################### vendor.qti.hardware.radio.common-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/RadioError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V2-ndk_platform-source/gen/vendor/qti/hardware/radio/RadioResponseInfo.cpp
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
    add_library(vendor.qti.hardware.radio.common-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.common-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.common-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.common-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.common-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.common-V2-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V1-ndk_platform ####################

#################### vendor.qti.hardware.radio.qtiradio-V11-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/AccessNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CagInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CallForwardInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellConnectionStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellIdentityWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoCdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoGsm.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoLte.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoTdscdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CellInfoWcdma.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/CiwlanMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ClosedSubscriberGroupInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EndcStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EutranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/EvdoSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/FacilityLockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GeranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/GsmSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadio.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/IQtiRadioResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/ImeiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/LteSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/McfgRefreshState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NgranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrIconType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconBandwidthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/NrUwbIconRefreshTimerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/OperatorInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/PersoUnlockStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellIdentityNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoNr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiCellInfoRatSpecificInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/QtiNetworkScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessFamily.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioAccessSpecifierBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/RadioTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SearchType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SetNetworkSelectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SignalQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/SnpnInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/TdscdmaSignalStrength.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/UtranBands.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/vendor/qti/hardware/radio/qtiradio/WcdmaSignalStrength.cpp
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
    add_library(vendor.qti.hardware.radio.qtiradio-V11-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.radio.qtiradio-V11-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/commonsys-intf/telephony/interfaces/aidl/qtiradio/vendor.qti.hardware.radio.qtiradio-V11-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.radio.qtiradio-V11-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.hardware.radio.qtiradio-V11-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.radio.qtiradio-V11-ndk_platform ####################

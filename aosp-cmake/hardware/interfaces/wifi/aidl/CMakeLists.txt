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

project(android.hardware.wifi-cpp-analyzer)

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

#################### android.hardware.wifi-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/AfcChannelAllowance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/AvailableAfcChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/AvailableAfcFrequencyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiApIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiChipEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiNanIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiNanIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiRttController.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiRttControllerEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IWifiStaIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IfaceConcurrencyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanBandIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanBandSpecificConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanBootstrappingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanBootstrappingMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanBootstrappingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanBootstrappingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanBootstrappingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanBootstrappingResponseCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanCipherSuiteType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanClusterEventInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanClusterEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanConfigRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanConfigRequestSupplemental.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanDataPathChannelCfg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanDataPathChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanDataPathConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanDataPathRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanDataPathScheduleUpdateInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanDataPathSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanDataPathSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanDebugConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanDiscoveryCommonConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanEnableRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanFollowupReceivedInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanIdentityResolutionAttribute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanInitiateDataPathRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanMatchAlg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanMatchInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPairingAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPairingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPairingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPairingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPairingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPairingRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPairingSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPairingSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPublishRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanPublishType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanRangingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanRespondToDataPathIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanRespondToPairingIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanSrfType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanSubscribeRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanSubscribeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanSuspensionModeChangeInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanTransmitFollowupRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NanTxType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/NpkSecurityAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttBw.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttLciInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttLcrInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttMotionPattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttPeerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttPreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttResponder.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/RttType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/Ssid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaApfPacketFilterCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaBackgroundScanBucketEventReportSchemeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaBackgroundScanBucketParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaBackgroundScanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaBackgroundScanLimits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaBackgroundScanParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaLinkLayerIfaceContentionTimeStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaLinkLayerIfacePacketStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaLinkLayerIfaceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaLinkLayerLinkStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaLinkLayerRadioStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaLinkLayerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaPeerInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaRateStat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaRoamingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaRoamingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaRoamingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaScanDataFlagMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/StaScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiAntennaMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiChannelStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiChannelWidthInMhz.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiChipCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxIcmpPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxMulticastPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugRingBufferFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugRingBufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugRingBufferVerboseLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugRxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugRxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugTxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiDebugTxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiIfaceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiInformationElement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiRadioCombination.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiRadioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiRateInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiRateNss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiRatePreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/android/hardware/wifi/WifiUsableChannel.cpp
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
    add_library(android.hardware.wifi-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi-V2-cpp
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
#################### android.hardware.wifi-V2-cpp ####################

#################### android.hardware.wifi-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/AfcChannelAllowance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/AvailableAfcChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/AvailableAfcFrequencyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiApIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiChipEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiNanIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiNanIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiRttController.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiRttControllerEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IWifiStaIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IfaceConcurrencyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanBandIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanBandSpecificConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanBootstrappingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanBootstrappingMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanBootstrappingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanBootstrappingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanBootstrappingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanBootstrappingResponseCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanCipherSuiteType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanClusterEventInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanClusterEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanConfigRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanConfigRequestSupplemental.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanDataPathChannelCfg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanDataPathChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanDataPathConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanDataPathRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanDataPathScheduleUpdateInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanDataPathSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanDataPathSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanDebugConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanDiscoveryCommonConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanEnableRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanFollowupReceivedInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanIdentityResolutionAttribute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanInitiateDataPathRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanMatchAlg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanMatchInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPairingAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPairingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPairingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPairingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPairingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPairingRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPairingSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPairingSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPublishRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanPublishType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanRangingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanRespondToDataPathIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanRespondToPairingIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanSrfType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanSubscribeRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanSubscribeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanSuspensionModeChangeInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanTransmitFollowupRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NanTxType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/NpkSecurityAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttBw.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttLciInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttLcrInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttMotionPattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttPeerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttPreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttResponder.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/RttType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/Ssid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaApfPacketFilterCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaBackgroundScanBucketEventReportSchemeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaBackgroundScanBucketParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaBackgroundScanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaBackgroundScanLimits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaBackgroundScanParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaLinkLayerIfaceContentionTimeStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaLinkLayerIfacePacketStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaLinkLayerIfaceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaLinkLayerLinkStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaLinkLayerRadioStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaLinkLayerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaPeerInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaRateStat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaRoamingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaRoamingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaRoamingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaScanDataFlagMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/StaScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiAntennaMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiChannelStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiChannelWidthInMhz.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiChipCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxIcmpPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxMulticastPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugRingBufferFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugRingBufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugRingBufferVerboseLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugRxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugRxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugTxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiDebugTxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiIfaceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiInformationElement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiRadioCombination.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiRadioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiRateInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiRateNss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiRatePreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/android/hardware/wifi/WifiUsableChannel.cpp
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
    add_library(android.hardware.wifi-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi-V1-ndk
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
#################### android.hardware.wifi-V1-ndk ####################

#################### android.hardware.wifi-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/AfcChannelAllowance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/AvailableAfcChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/AvailableAfcFrequencyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiApIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiChipEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiNanIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiNanIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiRttController.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiRttControllerEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IWifiStaIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IfaceConcurrencyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanBandIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanBandSpecificConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingResponseCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanCipherSuiteType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanClusterEventInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanClusterEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanConfigRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanConfigRequestSupplemental.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanDataPathChannelCfg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanDataPathChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanDataPathConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanDataPathRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanDataPathScheduleUpdateInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanDataPathSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanDataPathSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanDebugConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanDiscoveryCommonConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanEnableRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanFollowupReceivedInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanIdentityResolutionAttribute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanInitiateDataPathRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanMatchAlg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanMatchInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPairingAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPairingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPairingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPairingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPairingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPairingRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPairingSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPairingSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPublishRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanPublishType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanRangingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanRespondToDataPathIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanRespondToPairingIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanSrfType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanSubscribeRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanSubscribeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanSuspensionModeChangeInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanTransmitFollowupRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NanTxType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/NpkSecurityAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttBw.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttLciInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttLcrInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttMotionPattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttPeerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttPreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttResponder.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/RttType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/Ssid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaApfPacketFilterCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanBucketEventReportSchemeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanBucketParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanLimits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerIfaceContentionTimeStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerIfacePacketStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerIfaceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerLinkStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerRadioStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaPeerInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaRateStat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaRoamingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaRoamingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaRoamingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaScanDataFlagMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/StaScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiAntennaMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiChannelStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiChannelWidthInMhz.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiChipCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxIcmpPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxMulticastPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRingBufferFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRingBufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRingBufferVerboseLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugTxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiDebugTxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiIfaceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiInformationElement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiRadioCombination.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiRadioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiRateInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiRateNss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiRatePreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/android/hardware/wifi/WifiUsableChannel.cpp
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
    add_library(android.hardware.wifi-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi-V2-ndk_platform
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
#################### android.hardware.wifi-V2-ndk_platform ####################

#################### android.hardware.wifi-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/AfcChannelAllowance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/AvailableAfcChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/AvailableAfcFrequencyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiApIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiChipEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiNanIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiNanIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiRttController.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiRttControllerEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IWifiStaIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IfaceConcurrencyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanBandIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanBandSpecificConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanBootstrappingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanBootstrappingMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanBootstrappingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanBootstrappingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanBootstrappingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanBootstrappingResponseCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanCipherSuiteType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanClusterEventInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanClusterEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanConfigRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanConfigRequestSupplemental.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanDataPathChannelCfg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanDataPathChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanDataPathConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanDataPathRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanDataPathScheduleUpdateInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanDataPathSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanDataPathSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanDebugConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanDiscoveryCommonConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanEnableRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanFollowupReceivedInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanIdentityResolutionAttribute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanInitiateDataPathRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanMatchAlg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanMatchInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPairingAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPairingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPairingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPairingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPairingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPairingRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPairingSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPairingSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPublishRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanPublishType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanRangingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanRespondToDataPathIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanRespondToPairingIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanSrfType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanSubscribeRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanSubscribeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanSuspensionModeChangeInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanTransmitFollowupRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NanTxType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/NpkSecurityAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttBw.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttLciInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttLcrInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttMotionPattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttPeerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttPreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttResponder.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/RttType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/Ssid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaApfPacketFilterCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaBackgroundScanBucketEventReportSchemeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaBackgroundScanBucketParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaBackgroundScanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaBackgroundScanLimits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaBackgroundScanParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaLinkLayerIfaceContentionTimeStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaLinkLayerIfacePacketStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaLinkLayerIfaceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaLinkLayerLinkStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaLinkLayerRadioStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaLinkLayerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaPeerInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaRateStat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaRoamingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaRoamingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaRoamingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaScanDataFlagMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/StaScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiAntennaMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiChannelStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiChannelWidthInMhz.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiChipCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxIcmpPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxMulticastPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugRingBufferFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugRingBufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugRingBufferVerboseLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugRxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugRxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugTxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiDebugTxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiIfaceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiInformationElement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiRadioCombination.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiRadioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiRateInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiRateNss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiRatePreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/android/hardware/wifi/WifiUsableChannel.cpp
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
    add_library(android.hardware.wifi-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi-V2-ndk
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
#################### android.hardware.wifi-V2-ndk ####################

#################### android.hardware.wifi-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/AfcChannelAllowance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/AvailableAfcChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/AvailableAfcFrequencyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiApIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiChipEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiNanIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiNanIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiRttController.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiRttControllerEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IWifiStaIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IfaceConcurrencyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanBandIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanBandSpecificConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanBootstrappingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanBootstrappingMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanBootstrappingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanBootstrappingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanBootstrappingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanBootstrappingResponseCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanCipherSuiteType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanClusterEventInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanClusterEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanConfigRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanConfigRequestSupplemental.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanDataPathChannelCfg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanDataPathChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanDataPathConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanDataPathRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanDataPathScheduleUpdateInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanDataPathSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanDataPathSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanDebugConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanDiscoveryCommonConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanEnableRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanFollowupReceivedInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanIdentityResolutionAttribute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanInitiateDataPathRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanMatchAlg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanMatchInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPairingAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPairingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPairingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPairingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPairingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPairingRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPairingSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPairingSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPublishRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanPublishType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanRangingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanRespondToDataPathIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanRespondToPairingIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanSrfType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanSubscribeRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanSubscribeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanSuspensionModeChangeInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanTransmitFollowupRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NanTxType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/NpkSecurityAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttBw.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttLciInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttLcrInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttMotionPattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttPeerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttPreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttResponder.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/RttType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/Ssid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaApfPacketFilterCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaBackgroundScanBucketEventReportSchemeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaBackgroundScanBucketParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaBackgroundScanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaBackgroundScanLimits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaBackgroundScanParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaLinkLayerIfaceContentionTimeStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaLinkLayerIfacePacketStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaLinkLayerIfaceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaLinkLayerLinkStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaLinkLayerRadioStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaLinkLayerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaPeerInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaRateStat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaRoamingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaRoamingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaRoamingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaScanDataFlagMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/StaScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiAntennaMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiChannelStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiChannelWidthInMhz.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiChipCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxIcmpPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxMulticastPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugRingBufferFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugRingBufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugRingBufferVerboseLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugRxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugRxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugTxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiDebugTxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiIfaceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiInformationElement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiRadioCombination.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiRadioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiRateInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiRateNss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiRatePreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/android/hardware/wifi/WifiUsableChannel.cpp
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
    add_library(android.hardware.wifi-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi-V1-cpp
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
#################### android.hardware.wifi-V1-cpp ####################

#################### android.hardware.wifi-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/AfcChannelAllowance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/AvailableAfcChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/AvailableAfcFrequencyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiApIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiChipEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiNanIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiNanIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiRttController.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiRttControllerEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IWifiStaIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IfaceConcurrencyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanBandIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanBandSpecificConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanBootstrappingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanBootstrappingMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanBootstrappingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanBootstrappingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanBootstrappingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanBootstrappingResponseCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanCipherSuiteType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanClusterEventInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanClusterEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanConfigRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanConfigRequestSupplemental.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanDataPathChannelCfg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanDataPathChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanDataPathConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanDataPathRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanDataPathScheduleUpdateInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanDataPathSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanDataPathSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanDebugConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanDiscoveryCommonConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanEnableRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanFollowupReceivedInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanIdentityResolutionAttribute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanInitiateDataPathRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanMatchAlg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanMatchInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPairingAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPairingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPairingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPairingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPairingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPairingRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPairingSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPairingSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPublishRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanPublishType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanRangingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanRespondToDataPathIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanRespondToPairingIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanSrfType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanSubscribeRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanSubscribeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanSuspensionModeChangeInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanTransmitFollowupRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NanTxType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/NpkSecurityAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttBw.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttLciInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttLcrInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttMotionPattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttPeerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttPreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttResponder.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/RttType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/Ssid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaApfPacketFilterCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaBackgroundScanBucketEventReportSchemeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaBackgroundScanBucketParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaBackgroundScanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaBackgroundScanLimits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaBackgroundScanParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaLinkLayerIfaceContentionTimeStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaLinkLayerIfacePacketStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaLinkLayerIfaceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaLinkLayerLinkStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaLinkLayerRadioStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaLinkLayerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaPeerInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaRateStat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaRoamingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaRoamingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaRoamingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaScanDataFlagMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/StaScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiAntennaMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiChannelStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiChannelWidthInMhz.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiChipCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxIcmpPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxMulticastPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugRingBufferFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugRingBufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugRingBufferVerboseLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugRxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugRxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugTxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiDebugTxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiIfaceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiInformationElement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiRadioCombination.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiRadioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiRateInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiRateNss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiRatePreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/android/hardware/wifi/WifiUsableChannel.cpp
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
    add_executable(android.hardware.wifi-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/wifi/aidl
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V2-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi-cpp-analyzer
            aidl-analyzer-main
            android.hardware.wifi-V2-cpp
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
#################### android.hardware.wifi-cpp-analyzer ####################

#################### android.hardware.wifi-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/AfcChannelAllowance.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/AvailableAfcChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/AvailableAfcFrequencyInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifi.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiApIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiChipEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiNanIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiNanIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiRttController.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiRttControllerEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IWifiStaIfaceEventCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IfaceConcurrencyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanBandIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanBandSpecificConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanBootstrappingResponseCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanCipherSuiteType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanClusterEventInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanClusterEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanConfigRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanConfigRequestSupplemental.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanDataPathChannelCfg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanDataPathChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanDataPathConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanDataPathRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanDataPathScheduleUpdateInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanDataPathSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanDataPathSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanDebugConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanDiscoveryCommonConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanEnableRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanFollowupReceivedInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanIdentityResolutionAttribute.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanInitiateDataPathRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanMatchAlg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanMatchInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPairingAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPairingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPairingConfirmInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPairingRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPairingRequestInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPairingRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPairingSecurityConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPairingSecurityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPublishRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanPublishType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanRangingIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanRespondToDataPathIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanRespondToPairingIndicationRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanSrfType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanSubscribeRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanSubscribeType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanSuspensionModeChangeInd.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanTransmitFollowupRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NanTxType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/NpkSecurityAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttBw.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttLciInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttLcrInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttMotionPattern.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttPeerType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttPreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttResponder.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/RttType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/Ssid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaApfPacketFilterCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanBucketEventReportSchemeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanBucketParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanLimits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaBackgroundScanParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerIfaceContentionTimeStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerIfacePacketStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerIfaceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerLinkStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerRadioStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaLinkLayerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaPeerInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaRateStat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaRoamingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaRoamingConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaRoamingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaScanDataFlagMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/StaScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiAntennaMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiBand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiChannelInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiChannelStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiChannelWidthInMhz.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiChipCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxIcmpPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxMulticastPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonRxPacketDetails.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugHostWakeReasonStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugPacketFateFrameType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRingBufferFlags.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRingBufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRingBufferVerboseLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugRxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugTxPacketFate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiDebugTxPacketFateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiIfaceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiInformationElement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiRadioCombination.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiRadioConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiRateInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiRateNss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiRatePreamble.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/android/hardware/wifi/WifiUsableChannel.cpp
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
    add_library(android.hardware.wifi-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/aidl/android.hardware.wifi-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi-V1-ndk_platform
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
#################### android.hardware.wifi-V1-ndk_platform ####################

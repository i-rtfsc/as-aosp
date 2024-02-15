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

project(android.hardware.wifi.supplicant-cpp-analyzer)

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

#################### android.hardware.wifi.supplicant-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppConfigurationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppStatusErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/INonStandardCertCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pClientEapolIpAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pGroupStartedEventParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/PmkSaCacheData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/SignalPollResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/SupplicantStateChangeData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/TlsVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_library(android.hardware.wifi.supplicant-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-V2-ndk_platform
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
#################### android.hardware.wifi.supplicant-V2-ndk_platform ####################

#################### android.hardware.wifi.supplicant-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppConfigurationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/DppStatusErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/INonStandardCertCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/P2pClientEapolIpAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/P2pGroupStartedEventParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/PmkSaCacheData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/SignalPollResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/SupplicantStateChangeData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/TlsVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_library(android.hardware.wifi.supplicant-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-V2-ndk
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
#################### android.hardware.wifi.supplicant-V2-ndk ####################

#################### android.hardware.wifi.supplicant-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppConfigurationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/DppStatusErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/INonStandardCertCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/P2pClientEapolIpAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/P2pGroupStartedEventParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/PmkSaCacheData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/SignalPollResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/SupplicantStateChangeData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/TlsVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_library(android.hardware.wifi.supplicant-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-V3-ndk
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
#################### android.hardware.wifi.supplicant-V3-ndk ####################

#################### android.hardware.wifi.supplicant-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_library(android.hardware.wifi.supplicant-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-V1-cpp
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
#################### android.hardware.wifi.supplicant-V1-cpp ####################

#################### android.hardware.wifi.supplicant-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppConfigurationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/DppStatusErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/INonStandardCertCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/P2pClientEapolIpAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/P2pGroupStartedEventParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/PmkSaCacheData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/SignalPollResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/SupplicantStateChangeData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/TlsVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_library(android.hardware.wifi.supplicant-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-V2-cpp
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
#################### android.hardware.wifi.supplicant-V2-cpp ####################

#################### android.hardware.wifi.supplicant-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppConfigurationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppStatusErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/INonStandardCertCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pClientEapolIpAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pGroupStartedEventParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/PmkSaCacheData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/SignalPollResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/SupplicantStateChangeData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/TlsVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_library(android.hardware.wifi.supplicant-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-V3-ndk_platform
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
#################### android.hardware.wifi.supplicant-V3-ndk_platform ####################

#################### android.hardware.wifi.supplicant-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_library(android.hardware.wifi.supplicant-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-V1-ndk
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
#################### android.hardware.wifi.supplicant-V1-ndk ####################

#################### android.hardware.wifi.supplicant-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppConfigurationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/DppStatusErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/INonStandardCertCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/P2pClientEapolIpAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/P2pGroupStartedEventParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/PmkSaCacheData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/SignalPollResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/SupplicantStateChangeData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/TlsVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_library(android.hardware.wifi.supplicant-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-V3-cpp
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
#################### android.hardware.wifi.supplicant-V3-cpp ####################

#################### android.hardware.wifi.supplicant-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_library(android.hardware.wifi.supplicant-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-V1-ndk_platform
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
#################### android.hardware.wifi.supplicant-V1-ndk_platform ####################

#################### android.hardware.wifi.supplicant-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/AnqpInfoId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/AssociationRejectionData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/AuthAlgMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/AuxiliarySupplicantEventCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/BssTmData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/BssTmDataFlagsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/BssTmStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/BssidChangeReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/BtCoexistenceMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ConnectionCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DebugLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppAkm.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppConfigurationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppConnectionKeys.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppFailureCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppNetRole.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppProgressCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppResponderBootstrapInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/DppStatusErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/EapErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/EapMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/EapPhase2Method.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ExtRadioWorkDefaults.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/FreqRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/GroupCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/GroupMgmtCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/GsmRand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/Hs20AnqpData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/Hs20AnqpSubtypes.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/INonStandardCertCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ISupplicant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ISupplicantCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ISupplicantP2pNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ISupplicantStaIfaceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetwork.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ISupplicantStaNetworkCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/IfaceInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/IfaceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/IpVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/KeyMgmtMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/LegacyMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/MacAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/MboAssocDisallowedReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/MboCellularDataConnectionPrefValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/MboTransitionReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/MiracastMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/MloLink.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/MloLinksInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/NetworkRequestEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimGsmAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/NetworkResponseEapSimUmtsAuthParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/OceRssiBasedAssocRejectAttr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/OcspType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/OsuMethod.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/P2pClientEapolIpAddressInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/P2pFrameTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/P2pGroupCapabilityMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/P2pGroupStartedEventParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/P2pProvDiscStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/P2pStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/PairwiseCipherMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/PmkSaCacheData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/PortRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ProtoMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/ProtocolNextHeader.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyClassifierParamsMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyScsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyScsRequestStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyScsResponseStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/QosPolicyStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/RxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/SaeH2eMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/SignalPollResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/StaIfaceCallbackState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/StaIfaceReasonCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/StaIfaceStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/SupplicantStateChangeData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/SupplicantStatusCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/TlsVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/TransitionDisableIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/WifiTechnology.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/WpaDriverCapabilitiesMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/WpsConfigError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/WpsConfigMethods.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/WpsDevPasswordId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/WpsErrorIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/android/hardware/wifi/supplicant/WpsProvisionMethod.cpp
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
    add_executable(android.hardware.wifi.supplicant-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.hardware.wifi.supplicant-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/wifi/supplicant/aidl
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-V3-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/wifi/supplicant/aidl/android.hardware.wifi.supplicant-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.wifi.supplicant-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.wifi.supplicant-cpp-analyzer
            aidl-analyzer-main
            android.hardware.wifi.supplicant-V3-cpp
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
#################### android.hardware.wifi.supplicant-cpp-analyzer ####################

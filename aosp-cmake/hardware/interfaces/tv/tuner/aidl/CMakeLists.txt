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

project(android.hardware.tv.tuner-V1-ndk_platform)

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

#################### android.hardware.tv.tuner-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/AudioExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/AvStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/Constant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/Constant64Bit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DataFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpLengthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterAvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterIpPayloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMainType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEventExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMmtpRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterPesDataSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterPesEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterRecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterScIndexMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsCondition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsConditionTableInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterTemiEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterTsRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpAddressIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxPid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxQueueNotifyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxRecordScIndexType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScAvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScHevcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTlvFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DvrSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/DvrType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FilterDelayHint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FilterDelayHintType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogAftFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogSifStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Bandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3CodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3DemodOutputFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Fec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3PlpSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3TimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtscCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtscModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtscSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendCableTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcAnnex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcOuterFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsPilot.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsVcmMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtConstellation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtHierarchy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtPlpMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendInnerFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Coderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Rolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsStreamIdType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtLayerSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtPartialReceptionFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendModulationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendRollOff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanMessageStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendSpectralInversion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatusAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatusReadiness.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/IDemux.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/IDescrambler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/IDvr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/IDvrCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/IFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/IFilterCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/IFrontend.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/IFrontendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/ILnb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/ILnbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/ITimeFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/ITuner.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/LnbEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/LnbPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/LnbTone.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/LnbVoltage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/PlaybackSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/PlaybackStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/RecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/RecordStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/Result.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/ScramblingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/android/hardware/tv/tuner/VideoStreamType.cpp
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
    add_library(android.hardware.tv.tuner-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.tv.tuner-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.tv.tuner-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.tv.tuner-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
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
#################### android.hardware.tv.tuner-V1-ndk_platform ####################

#################### android.hardware.tv.tuner-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/AudioExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/AvStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/Constant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/Constant64Bit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DataFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpLengthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterAvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterIpPayloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMainType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEventExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMmtpRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterPesDataSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterPesEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterRecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterScIndexMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsCondition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsConditionTableInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterTemiEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterTsRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxIpAddressIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxIpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxPid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxQueueNotifyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxRecordScIndexType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxScAvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxScHevcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxScIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxTlvFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxTsFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DemuxTsIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DvrSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/DvrType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FilterDelayHint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FilterDelayHintType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogAftFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogSifStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Bandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3CodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3DemodOutputFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Fec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3PlpSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3TimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtscCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtscModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendAtscSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendCableTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcAnnex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcOuterFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsPilot.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsVcmMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtConstellation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtHierarchy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtPlpMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendInnerFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Coderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Rolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsStreamIdType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtLayerSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtPartialReceptionFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendModulationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendRollOff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendScanAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendScanMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendScanMessageStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendScanMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendSpectralInversion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendStatusAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendStatusReadiness.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/FrontendType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/IDemux.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/IDescrambler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/IDvr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/IDvrCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/IFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/IFilterCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/IFrontend.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/IFrontendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/ILnb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/ILnbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/ITimeFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/ITuner.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/LnbEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/LnbPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/LnbTone.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/LnbVoltage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/PlaybackSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/PlaybackStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/RecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/RecordStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/Result.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/ScramblingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/android/hardware/tv/tuner/VideoStreamType.cpp
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
    add_library(android.hardware.tv.tuner-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.tv.tuner-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.tv.tuner-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.tv.tuner-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
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
#################### android.hardware.tv.tuner-V1-ndk ####################

#################### android.hardware.tv.tuner-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/AudioExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/AudioPreselection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/AudioPreselectionLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/AudioPreselectionRenderingIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/AudioPresentation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/AvStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/Constant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/Constant64Bit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DataFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpLengthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterAvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterIpPayloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMainType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEventExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMmtpRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterPesDataSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterPesEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterRecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterScIndexMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsCondition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsConditionTableInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterTemiEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterTsRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpAddressIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxPid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxQueueNotifyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxRecordScIndexType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScAvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScHevcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScVvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTlvFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DvrSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/DvrType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FilterDelayHint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FilterDelayHintType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogAftFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogSifStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Bandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3CodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3DemodOutputFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Fec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3PlpSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3TimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtscCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtscModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtscSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendCableTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcAnnex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcOuterFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsPilot.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsVcmMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtConstellation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtHierarchy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtPlpMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendInnerFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsFecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsIgmp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsProtocol.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Coderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Rolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsStreamIdType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtLayerSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtPartialReceptionFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendModulationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendRollOff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanMessageStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendSpectralInversion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatusAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatusReadiness.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/IDemux.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/IDescrambler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/IDvr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/IDvrCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/IFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/IFilterCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/IFrontend.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/IFrontendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/ILnb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/ILnbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/ITimeFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/ITuner.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/LnbEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/LnbPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/LnbTone.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/LnbVoltage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/PlaybackSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/PlaybackStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/RecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/RecordStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/Result.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/ScramblingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/android/hardware/tv/tuner/VideoStreamType.cpp
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
    add_library(android.hardware.tv.tuner-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.tv.tuner-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.tv.tuner-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.tv.tuner-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
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
#################### android.hardware.tv.tuner-V2-ndk_platform ####################

#################### android.hardware.tv.tuner-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/AudioExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/AudioPreselection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/AudioPreselectionLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/AudioPreselectionRenderingIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/AudioPresentation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/AvStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/Constant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/Constant64Bit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DataFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxAlpLengthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterAvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterIpPayloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMainType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEventExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMmtpRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterPesDataSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterPesEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterRecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterScIndexMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsCondition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsConditionTableInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterTemiEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterTsRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpAddressIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxIpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxPid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxQueueNotifyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxRecordScIndexType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScAvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScHevcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxScVvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTlvFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DemuxTsIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DvrSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/DvrType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FilterDelayHint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FilterDelayHintType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogAftFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogSifStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAnalogType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Bandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3CodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3DemodOutputFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Fec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3PlpSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtsc3TimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtscCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtscModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendAtscSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendCableTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDtmbTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcAnnex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcOuterFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbcSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsPilot.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbsVcmMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtConstellation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtHierarchy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtPlpMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendDvbtTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendInnerFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsFecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsIgmp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsProtocol.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Coderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Rolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbsStreamIdType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtLayerSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtPartialReceptionFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendIsdbtTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendModulationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendRollOff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanMessageStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendSpectralInversion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatusAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatusReadiness.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/FrontendType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/IDemux.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/IDescrambler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/IDvr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/IDvrCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/IFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/IFilterCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/IFrontend.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/IFrontendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/ILnb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/ILnbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/ITimeFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/ITuner.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/LnbEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/LnbPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/LnbTone.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/LnbVoltage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/PlaybackSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/PlaybackStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/RecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/RecordStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/Result.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/ScramblingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/android/hardware/tv/tuner/VideoStreamType.cpp
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
    add_library(android.hardware.tv.tuner-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.tv.tuner-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.tv.tuner-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.tv.tuner-V3-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
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
#################### android.hardware.tv.tuner-V3-ndk_platform ####################

#################### android.hardware.tv.tuner-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/AudioExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/AudioPreselection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/AudioPreselectionLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/AudioPreselectionRenderingIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/AudioPresentation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/AvStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/Constant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/Constant64Bit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DataFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpLengthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterAvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterIpPayloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMainType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEventExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMmtpRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterPesDataSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterPesEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterRecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterScIndexMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsCondition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsConditionTableInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterTemiEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterTsRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxIpAddressIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxIpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxPid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxQueueNotifyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxRecordScIndexType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxScAvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxScHevcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxScIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxScVvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxTlvFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxTsFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DemuxTsIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DvrSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/DvrType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FilterDelayHint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FilterDelayHintType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogAftFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogSifStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Bandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3CodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3DemodOutputFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Fec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3PlpSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3TimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtscCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtscModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendAtscSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendCableTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcAnnex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcOuterFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsPilot.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsVcmMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtConstellation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtHierarchy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtPlpMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendInnerFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsFecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsIgmp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsProtocol.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Coderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Rolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsStreamIdType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtLayerSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtPartialReceptionFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendModulationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendRollOff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendScanAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendScanMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendScanMessageStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendScanMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendSpectralInversion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendStatusAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendStatusReadiness.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/FrontendType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/IDemux.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/IDescrambler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/IDvr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/IDvrCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/IFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/IFilterCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/IFrontend.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/IFrontendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/ILnb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/ILnbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/ITimeFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/ITuner.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/LnbEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/LnbPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/LnbTone.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/LnbVoltage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/PlaybackSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/PlaybackStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/RecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/RecordStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/Result.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/ScramblingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/android/hardware/tv/tuner/VideoStreamType.cpp
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
    add_library(android.hardware.tv.tuner-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.tv.tuner-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.tv.tuner-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.tv.tuner-V3-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
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
#################### android.hardware.tv.tuner-V3-ndk ####################

#################### android.hardware.tv.tuner-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/AudioExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/AudioPreselection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/AudioPreselectionLabel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/AudioPreselectionRenderingIndicationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/AudioPresentation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/AudioStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/AvStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/Constant.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/Constant64Bit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DataFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxAlpLengthType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterAvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterDownloadSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterIpPayloadEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMainType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMediaEventExtraMetaData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMmtpRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterMonitorEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterPesDataSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterPesEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterRecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterScIndexMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsCondition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSectionSettingsConditionTableInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterSubType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterTemiEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterTsRecordEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxIpAddressIpAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxIpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxIpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxMmtpFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxPid.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxQueueNotifyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxRecordScIndexType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxScAvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxScHevcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxScIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxScVvcIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxTlvFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxTlvFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxTsFilterSettingsFilterSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxTsFilterType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DemuxTsIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DvrSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/DvrType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FilterDelayHint.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FilterDelayHintType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogAftFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogSifStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAnalogType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Bandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3CodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3DemodOutputFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Fec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3PlpSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtsc3TimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtscCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtscModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendAtscSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendCableTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDtmbTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcAnnex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcOuterFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbcSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsCodeRate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsPilot.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbsVcmMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtConstellation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtHierarchy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtPlpMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendDvbtTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendInnerFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsFec.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsFecType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsIgmp.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIptvSettingsProtocol.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Capabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Coderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Modulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Rolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbs3Settings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsRolloff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbsStreamIdType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtBandwidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtCoderate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtGuardInterval.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtLayerSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtPartialReceptionFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendIsdbtTimeInterleaveMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendModulation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendModulationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendRollOff.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendScanAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendScanMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendScanMessageStandard.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendScanMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendScanType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendSpectralInversion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendStatusAtsc3PlpInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendStatusReadiness.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendTransmissionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/FrontendType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/IDemux.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/IDescrambler.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/IDvr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/IDvrCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/IFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/IFilterCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/IFrontend.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/IFrontendCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/ILnb.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/ILnbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/ITimeFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/ITuner.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/LnbEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/LnbPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/LnbTone.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/LnbVoltage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/PlaybackSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/PlaybackStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/RecordSettings.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/RecordStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/Result.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/ScramblingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/android/hardware/tv/tuner/VideoStreamType.cpp
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
    add_library(android.hardware.tv.tuner-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.tv.tuner-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/tv/tuner/aidl/android.hardware.tv.tuner-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.tv.tuner-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.tv.tuner-V2-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
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
#################### android.hardware.tv.tuner-V2-ndk ####################

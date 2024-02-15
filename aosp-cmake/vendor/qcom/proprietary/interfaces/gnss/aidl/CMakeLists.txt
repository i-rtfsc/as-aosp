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

project(vendor.qti.gnss-cpp-analyzer)

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

#################### vendor.qti.gnss-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V2-cpp
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
#################### vendor.qti.gnss-V2-cpp ####################

#################### vendor.qti.gnss-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V3-ndk_platform
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
#################### vendor.qti.gnss-V3-ndk_platform ####################

#################### vendor.qti.gnss-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlAltitudeRefType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApRangingData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApRangingScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApServingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlQesdkSessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlRangingBandWidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlWiFiAPPositionQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlWiFiRTTCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_executable(vendor.qti.gnss-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libutils
            vendor.qti.gnss-V6-cpp
    )
endif ()
#################### vendor.qti.gnss-cpp-analyzer ####################

#################### vendor.qti.gnss-V5-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlAltitudeRefType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApRangingData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApRangingScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApServingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlRangingBandWidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlWiFiAPPositionQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlWiFiRTTCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V5-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V5-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V5-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V5-ndk
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
#################### vendor.qti.gnss-V5-ndk ####################

#################### vendor.qti.gnss-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V3-cpp
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
#################### vendor.qti.gnss-V3-cpp ####################

#################### vendor.qti.gnss-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
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
    add_library(vendor.qti.gnss-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V1-ndk
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
#################### vendor.qti.gnss-V1-ndk ####################

#################### vendor.qti.gnss-V6-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlAltitudeRefType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApRangingData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApRangingScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApServingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlRangingBandWidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlWiFiAPPositionQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlWiFiRTTCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V6-cpp SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V6-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V6-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V6-cpp
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
#################### vendor.qti.gnss-V6-cpp ####################

#################### vendor.qti.gnss-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V4-ndk_platform
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
#################### vendor.qti.gnss-V4-ndk_platform ####################

#################### vendor.qti.gnss-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V4-ndk
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
#################### vendor.qti.gnss-V4-ndk ####################

#################### vendor.qti.gnss-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V2-ndk
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
#################### vendor.qti.gnss-V2-ndk ####################

#################### vendor.qti.gnss-V5-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAltitudeRefType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApRangingData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApRangingScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApServingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRangingBandWidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWiFiAPPositionQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWiFiRTTCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V5-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V5-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V5-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V5-ndk_platform
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
#################### vendor.qti.gnss-V5-ndk_platform ####################

#################### vendor.qti.gnss-V5-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlAltitudeRefType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApRangingData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApRangingScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApServingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlRangingBandWidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlWiFiAPPositionQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlWiFiRTTCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V5-cpp SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V5-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V5-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V5-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V5-cpp
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
#################### vendor.qti.gnss-V5-cpp ####################

#################### vendor.qti.gnss-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
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
    add_library(vendor.qti.gnss-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V1-cpp
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
#################### vendor.qti.gnss-V1-cpp ####################

#################### vendor.qti.gnss-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V3-ndk
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
#################### vendor.qti.gnss-V3-ndk ####################

#################### vendor.qti.gnss-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V2-ndk_platform
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
#################### vendor.qti.gnss-V2-ndk_platform ####################

#################### vendor.qti.gnss-V4-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V4-cpp SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V4-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V4-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V4-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V4-cpp
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
#################### vendor.qti.gnss-V4-cpp ####################

#################### vendor.qti.gnss-V6-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlAltitudeRefType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApRangingData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApRangingScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApServingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlQesdkSessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlRangingBandWidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlWiFiAPPositionQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlWiFiRTTCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V6-ndk SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V6-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V6-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V6-ndk
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
#################### vendor.qti.gnss-V6-ndk ####################

#################### vendor.qti.gnss-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
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
    add_library(vendor.qti.gnss-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V1-ndk_platform
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
#################### vendor.qti.gnss-V1-ndk_platform ####################

#################### vendor.qti.gnss-V6-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlDebugReportServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlEsStatusReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlFlpServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoder.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeocoderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGeofenceServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigService.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssConfigServiceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNi.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlGnssNiCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatConfigCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscription.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlIzatSubscriptionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTracking.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlQesdkTrackingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlRilInfoMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWWANDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProvider.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBProviderCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiver.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/ILocAidlWiFiDBReceiverCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV4.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusIpV6.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssStatusValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssSubId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAGnssType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlAltitudeRefType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApRangingData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApRangingScanResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApScanData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApServingStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApTimeStamp.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlApnTypeMask.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBSObsData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchOptions.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatchStatusInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBatteryLevelDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBoolDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsCellInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsLocationData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBsSpecialInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlBtLeDeviceScanDetailsDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellCdmaDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellGwDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellLteDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlCellOooDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlEphemerisType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceBreachType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGeofenceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiEncodingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiNotification.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssNiType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssSvUsedInPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlGnssUserResponseType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatHorizontalAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatLocationStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatOtherAccuracy.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatProviderStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlIzatStreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationError.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationExtended.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationFlagsBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationQualityType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlLocationTechnologyBits.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNetworkPositionSourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlNtripConnectionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlPowerConnectStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionPrecision.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlQesdkSessionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRangingBandWidth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRilServiceInfoDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationConfigValidBit.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlRobustLocationInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlScreenStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlServiceStateDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlStringDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSubscriptionDataItemId.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusBestPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusEphemeris.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusErrRecovery.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusInjectedPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNav.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusNavData.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusPositionFailure.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusReports.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusRfAndParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusSvHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusTimeAndClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXoState.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlSystemStatusXtra.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTime.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlTimeZoneChangeDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlUlpLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWiFiAPPositionQuality.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWiFiRTTCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiDBListStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlWifiSupplicantStatusDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraDataStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/vendor/qti/gnss/LocAidlXtraStatus.cpp
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
    add_library(vendor.qti.gnss-V6-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(vendor.qti.gnss-V6-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/proprietary/interfaces/gnss/aidl/vendor.qti.gnss-V6-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.gnss-V6-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(vendor.qti.gnss-V6-ndk_platform
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
#################### vendor.qti.gnss-V6-ndk_platform ####################

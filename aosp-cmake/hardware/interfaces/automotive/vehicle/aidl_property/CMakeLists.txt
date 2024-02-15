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

project(android.hardware.automotive.vehicle.property-V1-ndk_platform)

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

#################### android.hardware.automotive.vehicle.property-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/CreateUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/CreateUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/CreateUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/CustomInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/DiagnosticFloatSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/DiagnosticIntegerSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvChargeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvConnectorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvRegenerativeBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvsServiceRequestIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvsServiceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvsServiceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponseAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2CommonIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2CompressionIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2FuelSystemStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2IgnitionMonitorKind.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2SecondaryAirStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2SparkIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/PortLocationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/ProcessTerminationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/RemoveUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/RotaryInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/TrailerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationSetValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationGetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/UsersInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateConfigFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReqIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateShutdownParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleArea.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaDoor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaMirror.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaSeat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaWheel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaWindow.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleGear.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHvacFanDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwKeyInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleIgnitionState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleLightState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleLightSwitch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleOilLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleProperty.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehiclePropertyGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehiclePropertyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleSeatOccupancyState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleTurnSignal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleUnit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleVendorPermission.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsAvailabilityStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsBaseMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerAndPublisherIdIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsOfferingMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsPublisherInformationIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsStartSessionMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsSubscriptionsStateIntegerValuesIndex.cpp
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
    add_library(android.hardware.automotive.vehicle.property-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle.property-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle.property-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.vehicle.property-V1-ndk_platform
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
#################### android.hardware.automotive.vehicle.property-V1-ndk_platform ####################

#################### android.hardware.automotive.vehicle.property-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/AutomaticEmergencyBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/BlindSpotWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/CreateUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/CreateUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/CreateUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/CruiseControlCommand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/CruiseControlState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/CruiseControlType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/CustomInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/DiagnosticFloatSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/DiagnosticIntegerSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/EmergencyLaneKeepAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/ErrorState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvChargeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvConnectorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvRegenerativeBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvStoppingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvsServiceRequestIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvsServiceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvsServiceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/ForwardCollisionWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/GsrComplianceRequirementType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/HandsOnDetectionDriverState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/HandsOnDetectionWarning.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponseAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/LaneCenteringAssistCommand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/LaneCenteringAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/LaneDepartureWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/LaneKeepAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/LocationCharacterization.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2CommonIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2CompressionIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2FuelSystemStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2IgnitionMonitorKind.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2SecondaryAirStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2SparkIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/PortLocationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/ProcessTerminationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/RemoveUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/RotaryInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/TrailerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationSetValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationGetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/UsersInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerBootupReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateConfigFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReqIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateShutdownParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleArea.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaDoor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaMirror.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaSeat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaWheel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaWindow.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleGear.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHvacFanDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwKeyInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionButtonStateFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionInputSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionToolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleIgnitionState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleLightState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleLightSwitch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleOilLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleProperty.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehiclePropertyGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehiclePropertyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleSeatOccupancyState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleTurnSignal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleUnit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleVendorPermission.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsAvailabilityStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsBaseMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerAndPublisherIdIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsOfferingMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsPublisherInformationIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsStartSessionMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsSubscriptionsStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/WindshieldWipersState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/android/hardware/automotive/vehicle/WindshieldWipersSwitch.cpp
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
    add_library(android.hardware.automotive.vehicle.property-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle.property-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle.property-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.vehicle.property-V3-ndk_platform
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
#################### android.hardware.automotive.vehicle.property-V3-ndk_platform ####################

#################### android.hardware.automotive.vehicle.property-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/CreateUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/CreateUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/CreateUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/CustomInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/DiagnosticFloatSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/DiagnosticIntegerSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/EvChargeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/EvConnectorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/EvRegenerativeBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/EvsServiceRequestIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/EvsServiceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/EvsServiceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponseAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/Obd2CommonIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/Obd2CompressionIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/Obd2FuelSystemStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/Obd2FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/Obd2IgnitionMonitorKind.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/Obd2SecondaryAirStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/Obd2SparkIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/PortLocationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/ProcessTerminationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/RemoveUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/RotaryInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/TrailerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationSetValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationGetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UserInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/UsersInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateConfigFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReqIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateShutdownParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleArea.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaDoor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaMirror.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaSeat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaWheel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaWindow.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleGear.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHvacFanDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwKeyInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleIgnitionState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleLightState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleLightSwitch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleOilLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleProperty.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehiclePropertyGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehiclePropertyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleSeatOccupancyState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleTurnSignal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleUnit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VehicleVendorPermission.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VmsAvailabilityStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VmsBaseMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VmsMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerAndPublisherIdIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VmsOfferingMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VmsPublisherInformationIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VmsStartSessionMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/android/hardware/automotive/vehicle/VmsSubscriptionsStateIntegerValuesIndex.cpp
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
    add_library(android.hardware.automotive.vehicle.property-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle.property-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle.property-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.vehicle.property-V1-ndk
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
#################### android.hardware.automotive.vehicle.property-V1-ndk ####################

#################### android.hardware.automotive.vehicle.property-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/AutomaticEmergencyBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/BlindSpotWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/CreateUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/CreateUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/CreateUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/CruiseControlCommand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/CruiseControlState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/CruiseControlType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/CustomInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/DiagnosticFloatSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/DiagnosticIntegerSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/EmergencyLaneKeepAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/ErrorState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/EvChargeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/EvConnectorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/EvRegenerativeBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/EvStoppingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/EvsServiceRequestIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/EvsServiceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/EvsServiceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/ForwardCollisionWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/GsrComplianceRequirementType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/HandsOnDetectionDriverState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/HandsOnDetectionWarning.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponseAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/LaneCenteringAssistCommand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/LaneCenteringAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/LaneDepartureWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/LaneKeepAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/LocationCharacterization.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/Obd2CommonIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/Obd2CompressionIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/Obd2FuelSystemStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/Obd2FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/Obd2IgnitionMonitorKind.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/Obd2SecondaryAirStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/Obd2SparkIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/PortLocationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/ProcessTerminationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/RemoveUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/RotaryInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/TrailerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationSetValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationGetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UserInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/UsersInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerBootupReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateConfigFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReqIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateShutdownParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleArea.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaDoor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaMirror.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaSeat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaWheel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaWindow.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleGear.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHvacFanDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwKeyInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionButtonStateFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionInputSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionToolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleIgnitionState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleLightState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleLightSwitch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleOilLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleProperty.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehiclePropertyGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehiclePropertyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleSeatOccupancyState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleTurnSignal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleUnit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VehicleVendorPermission.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VmsAvailabilityStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VmsBaseMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VmsMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerAndPublisherIdIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VmsOfferingMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VmsPublisherInformationIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VmsStartSessionMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/VmsSubscriptionsStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/WindshieldWipersState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/android/hardware/automotive/vehicle/WindshieldWipersSwitch.cpp
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
    add_library(android.hardware.automotive.vehicle.property-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle.property-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle.property-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.vehicle.property-V2-ndk
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
#################### android.hardware.automotive.vehicle.property-V2-ndk ####################

#################### android.hardware.automotive.vehicle.property-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/AutomaticEmergencyBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/BlindSpotWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/CreateUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/CreateUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/CreateUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/CruiseControlCommand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/CruiseControlState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/CruiseControlType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/CustomInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/DiagnosticFloatSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/DiagnosticIntegerSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/EmergencyLaneKeepAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/ErrorState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvChargeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvConnectorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvRegenerativeBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvStoppingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvsServiceRequestIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvsServiceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/EvsServiceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/ForwardCollisionWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/GsrComplianceRequirementType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/HandsOnDetectionDriverState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/HandsOnDetectionWarning.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponseAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/LaneCenteringAssistCommand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/LaneCenteringAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/LaneDepartureWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/LaneKeepAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/LocationCharacterization.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2CommonIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2CompressionIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2FuelSystemStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2IgnitionMonitorKind.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2SecondaryAirStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/Obd2SparkIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/PortLocationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/ProcessTerminationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/RemoveUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/RotaryInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/SwitchUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/TrailerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationSetValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationGetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UserInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/UsersInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerBootupReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateConfigFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReqIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateShutdownParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleArea.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaDoor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaMirror.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaSeat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaWheel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleAreaWindow.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleGear.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHvacFanDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwKeyInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionButtonStateFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionInputSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionToolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleIgnitionState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleLightState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleLightSwitch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleOilLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleProperty.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehiclePropertyGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehiclePropertyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleSeatOccupancyState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleTurnSignal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleUnit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VehicleVendorPermission.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsAvailabilityStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsBaseMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerAndPublisherIdIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsOfferingMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsPublisherInformationIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsStartSessionMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/VmsSubscriptionsStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/WindshieldWipersState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/android/hardware/automotive/vehicle/WindshieldWipersSwitch.cpp
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
    add_library(android.hardware.automotive.vehicle.property-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle.property-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle.property-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.vehicle.property-V2-ndk_platform
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
#################### android.hardware.automotive.vehicle.property-V2-ndk_platform ####################

#################### android.hardware.automotive.vehicle.property-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/AutomaticEmergencyBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/BlindSpotWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/CreateUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/CreateUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/CreateUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/CruiseControlCommand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/CruiseControlState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/CruiseControlType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/CustomInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/DiagnosticFloatSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/DiagnosticIntegerSensorIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/ElectronicTollCollectionCardType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/EmergencyLaneKeepAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/ErrorState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/EvChargeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/EvConnectorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/EvRegenerativeBrakingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/EvStoppingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/EvsServiceRequestIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/EvsServiceState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/EvsServiceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/ForwardCollisionWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/GsrComplianceRequirementType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/HandsOnDetectionDriverState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/HandsOnDetectionWarning.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoRequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/InitialUserInfoResponseAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/LaneCenteringAssistCommand.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/LaneCenteringAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/LaneDepartureWarningState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/LaneKeepAssistState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/LocationCharacterization.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/Obd2CommonIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/Obd2CompressionIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/Obd2FuelSystemStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/Obd2FuelType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/Obd2IgnitionMonitorKind.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/Obd2SecondaryAirStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/Obd2SparkIgnitionMonitors.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/PortLocationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/ProcessTerminationReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/RemoveUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/RotaryInputType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/SwitchUserStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/TrailerState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationSetValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationAssociationValue.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationGetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationResponse.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetAssociation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UserIdentificationSetRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UserInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/UsersInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerBootupReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateConfigFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReq.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateReqIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleApPowerStateShutdownParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleArea.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaDoor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaMirror.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaSeat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaWheel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleAreaWindow.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleGear.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHvacFanDirection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwKeyInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionButtonStateFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionInputAction.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionInputSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleHwMotionToolType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleIgnitionState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleLightState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleLightSwitch.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleOilLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleProperty.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehiclePropertyGroup.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehiclePropertyType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleSeatOccupancyState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleTurnSignal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleUnit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VehicleVendorPermission.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VmsAvailabilityStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VmsBaseMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VmsMessageType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerAndPublisherIdIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VmsMessageWithLayerIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VmsOfferingMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VmsPublisherInformationIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VmsStartSessionMessageIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/VmsSubscriptionsStateIntegerValuesIndex.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/WindshieldWipersState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/android/hardware/automotive/vehicle/WindshieldWipersSwitch.cpp
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
    add_library(android.hardware.automotive.vehicle.property-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.vehicle.property-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/vehicle/aidl_property/android.hardware.automotive.vehicle.property-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.vehicle.property-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.vehicle.property-V3-ndk
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
#################### android.hardware.automotive.vehicle.property-V3-ndk ####################

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

project(android.hardware.gnss-cpp-analyzer)

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

#################### android.hardware.gnss-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_library(android.hardware.gnss-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V3-ndk_platform
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
#################### android.hardware.gnss-V3-ndk_platform ####################

#################### android.hardware.gnss-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_library(android.hardware.gnss-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V4-ndk_platform
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
#################### android.hardware.gnss-V4-ndk_platform ####################

#################### android.hardware.gnss-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_library(android.hardware.gnss-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V2-ndk_platform
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
#################### android.hardware.gnss-V2-ndk_platform ####################

#################### android.hardware.gnss-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_library(android.hardware.gnss-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V2-cpp
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
#################### android.hardware.gnss-V2-cpp ####################

#################### android.hardware.gnss-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
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
    add_library(android.hardware.gnss-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V1-ndk
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
#################### android.hardware.gnss-V1-ndk ####################

#################### android.hardware.gnss-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_library(android.hardware.gnss-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V3-cpp
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
#################### android.hardware.gnss-V3-cpp ####################

#################### android.hardware.gnss-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_library(android.hardware.gnss-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V4-ndk
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
#################### android.hardware.gnss-V4-ndk ####################

#################### android.hardware.gnss-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_library(android.hardware.gnss-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V2-ndk
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
#################### android.hardware.gnss-V2-ndk ####################

#################### android.hardware.gnss-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
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
    add_library(android.hardware.gnss-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V1-ndk_platform
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
#################### android.hardware.gnss-V1-ndk_platform ####################

#################### android.hardware.gnss-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_library(android.hardware.gnss-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V3-ndk
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
#################### android.hardware.gnss-V3-ndk ####################

#################### android.hardware.gnss-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
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
    add_library(android.hardware.gnss-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V1-cpp
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
#################### android.hardware.gnss-V1-cpp ####################

#################### android.hardware.gnss-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_executable(android.hardware.gnss-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/gnss/aidl
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-cpp-analyzer
            aidl-analyzer-main
            android.hardware.gnss-V4-cpp
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
#################### android.hardware.gnss-cpp-analyzer ####################

#################### android.hardware.gnss-V4-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/BlocklistedSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/CorrelationVector.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/ElapsedRealtime.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/GnssClock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/GnssConstellationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/GnssData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/GnssLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/GnssMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/GnssMultipathIndicator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/GnssPowerStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/GnssSignalType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IAGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IAGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IAGnssRil.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IAGnssRilCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnss.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssAntennaInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssAntennaInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssBatching.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssBatchingCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssDebug.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssGeofence.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssGeofenceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssMeasurementCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssMeasurementInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssNavigationMessageCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssNavigationMessageInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssPowerIndication.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssPowerIndicationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssPsds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/IGnssPsdsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/PsdsType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/SatelliteClockInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/SatellitePositionEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/SatellitePvt.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/SatelliteVelocityEcef.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/measurement_corrections/IMeasurementCorrectionsInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/measurement_corrections/MeasurementCorrections.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/measurement_corrections/ReflectingPlane.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/measurement_corrections/SingleSatCorrection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/android/hardware/gnss/visibility_control/IGnssVisibilityControlCallback.cpp
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
    add_library(android.hardware.gnss-V4-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.gnss-V4-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/gnss/aidl/android.hardware.gnss-V4-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.gnss-V4-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.gnss-V4-cpp
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
#################### android.hardware.gnss-V4-cpp ####################

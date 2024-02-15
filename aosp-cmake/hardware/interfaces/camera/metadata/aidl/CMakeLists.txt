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

project(android.hardware.camera.metadata-V2-ndk)

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

#################### android.hardware.camera.metadata-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/AutomotiveLensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/AutomotiveLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/BlackLevelLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/CameraMetadataSection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/CameraMetadataSectionStart.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/CameraMetadataTag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ColorCorrectionAberrationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ColorCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAeAntibandingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAeLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAePrecaptureTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAfMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAfSceneChange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAfState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAfTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAutoframing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAutoframingAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAutoframingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAwbLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAwbLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAwbMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlAwbState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlCaptureIntent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlEffectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlEnableZsl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlExtendedSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlSettingsOverride.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ControlVideoStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/DemosaicMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/DepthDepthIsExclusive.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/DistortionCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/EdgeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/FlashInfoAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/FlashMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/FlashState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/HeicInfoSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/HotPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/InfoSupportedBufferManagementVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/InfoSupportedHardwareLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/JpegrAvailableJpegRStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/JpegrAvailableJpegRStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/LedAvailableLeds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/LedTransmit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/LensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/LensInfoFocusDistanceCalibration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/LensOpticalStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/LensPoseReference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/LensState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/LogicalMultiCameraSensorSyncType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/NoiseReductionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/QuirksPartialResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/RequestAvailableCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/RequestAvailableColorSpaceProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/RequestAvailableDynamicRangeProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/RequestMetadataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/RequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableFormats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableRecommendedStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamUseCases.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ScalerCroppingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ScalerMultiResolutionStreamSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ScalerPhysicalCameraMultiResolutionStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ScalerRotateAndCrop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/SensorInfoColorFilterArrangement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/SensorInfoLensShadingApplied.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/SensorInfoTimestampSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/SensorPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/SensorRawBinningFactorUsed.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/SensorReferenceIlluminant1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/SensorTestPatternMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/ShadingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/StatisticsFaceDetectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/StatisticsHistogramMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/StatisticsHotPixelMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/StatisticsLensShadingMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/StatisticsOisDataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/StatisticsSceneFlicker.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/StatisticsSharpnessMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/SyncFrameNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/SyncMaxLatency.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/TonemapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/android/hardware/camera/metadata/TonemapPresetCurve.cpp
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
    add_library(android.hardware.camera.metadata-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.metadata-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.metadata-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.metadata-V2-ndk
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
#################### android.hardware.camera.metadata-V2-ndk ####################

#################### android.hardware.camera.metadata-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/AutomotiveLensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/AutomotiveLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/BlackLevelLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/CameraMetadataSection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/CameraMetadataSectionStart.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/CameraMetadataTag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ColorCorrectionAberrationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ColorCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAeAntibandingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAeLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAePrecaptureTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAfMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAfSceneChange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAfState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAfTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAutoframing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAutoframingAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAutoframingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAwbLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAwbLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAwbMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlAwbState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlCaptureIntent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlEffectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlEnableZsl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlExtendedSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlSettingsOverride.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ControlVideoStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/DemosaicMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/DepthDepthIsExclusive.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/DistortionCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/EdgeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/FlashInfoAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/FlashMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/FlashState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/HeicInfoSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/HotPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/InfoSupportedBufferManagementVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/InfoSupportedHardwareLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/JpegrAvailableJpegRStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/JpegrAvailableJpegRStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/LedAvailableLeds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/LedTransmit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/LensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/LensInfoFocusDistanceCalibration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/LensOpticalStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/LensPoseReference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/LensState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/LogicalMultiCameraSensorSyncType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/NoiseReductionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/QuirksPartialResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/RequestAvailableCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/RequestAvailableColorSpaceProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/RequestAvailableDynamicRangeProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/RequestMetadataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/RequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableFormats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableRecommendedStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamUseCases.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ScalerCroppingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ScalerMultiResolutionStreamSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ScalerPhysicalCameraMultiResolutionStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ScalerRotateAndCrop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/SensorInfoColorFilterArrangement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/SensorInfoLensShadingApplied.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/SensorInfoTimestampSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/SensorPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/SensorRawBinningFactorUsed.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/SensorReferenceIlluminant1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/SensorTestPatternMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/ShadingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/StatisticsFaceDetectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/StatisticsHistogramMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/StatisticsHotPixelMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/StatisticsLensShadingMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/StatisticsOisDataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/StatisticsSceneFlicker.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/StatisticsSharpnessMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/SyncFrameNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/SyncMaxLatency.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/TonemapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/android/hardware/camera/metadata/TonemapPresetCurve.cpp
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
    add_library(android.hardware.camera.metadata-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.metadata-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.metadata-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.metadata-V3-ndk
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
#################### android.hardware.camera.metadata-V3-ndk ####################

#################### android.hardware.camera.metadata-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/AutomotiveLensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/AutomotiveLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/BlackLevelLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/CameraMetadataSection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/CameraMetadataSectionStart.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/CameraMetadataTag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ColorCorrectionAberrationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ColorCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeAntibandingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAePrecaptureTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfSceneChange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAutoframing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAutoframingAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAutoframingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlCaptureIntent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlEffectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlEnableZsl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlExtendedSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlSettingsOverride.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ControlVideoStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/DemosaicMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/DepthDepthIsExclusive.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/DistortionCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/EdgeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/FlashInfoAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/FlashMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/FlashState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/HeicInfoSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/HotPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/InfoSupportedBufferManagementVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/InfoSupportedHardwareLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/JpegrAvailableJpegRStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/JpegrAvailableJpegRStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/LedAvailableLeds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/LedTransmit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/LensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/LensInfoFocusDistanceCalibration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/LensOpticalStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/LensPoseReference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/LensState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/LogicalMultiCameraSensorSyncType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/NoiseReductionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/QuirksPartialResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/RequestAvailableCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/RequestAvailableColorSpaceProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/RequestAvailableDynamicRangeProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/RequestMetadataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/RequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableFormats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableRecommendedStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamUseCases.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerCroppingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerMultiResolutionStreamSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerPhysicalCameraMultiResolutionStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerRotateAndCrop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/SensorInfoColorFilterArrangement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/SensorInfoLensShadingApplied.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/SensorInfoTimestampSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/SensorPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/SensorRawBinningFactorUsed.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/SensorReferenceIlluminant1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/SensorTestPatternMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/ShadingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsFaceDetectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsHistogramMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsHotPixelMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsLensShadingMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsOisDataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsSceneFlicker.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsSharpnessMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/SyncFrameNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/SyncMaxLatency.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/TonemapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/android/hardware/camera/metadata/TonemapPresetCurve.cpp
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
    add_library(android.hardware.camera.metadata-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.metadata-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.metadata-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.metadata-V3-ndk_platform
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
#################### android.hardware.camera.metadata-V3-ndk_platform ####################

#################### android.hardware.camera.metadata-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/AutomotiveLensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/AutomotiveLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/BlackLevelLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/CameraMetadataSection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/CameraMetadataSectionStart.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/CameraMetadataTag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ColorCorrectionAberrationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ColorCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeAntibandingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAePrecaptureTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfSceneChange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAutoframing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAutoframingAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAutoframingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlCaptureIntent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlEffectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlEnableZsl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlExtendedSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlSettingsOverride.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ControlVideoStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/DemosaicMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/DepthDepthIsExclusive.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/DistortionCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/EdgeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/FlashInfoAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/FlashMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/FlashState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/HeicInfoSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/HotPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/InfoSupportedBufferManagementVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/InfoSupportedHardwareLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/JpegrAvailableJpegRStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/JpegrAvailableJpegRStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/LedAvailableLeds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/LedTransmit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/LensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/LensInfoFocusDistanceCalibration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/LensOpticalStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/LensPoseReference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/LensState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/LogicalMultiCameraSensorSyncType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/NoiseReductionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/QuirksPartialResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/RequestAvailableCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/RequestAvailableColorSpaceProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/RequestAvailableDynamicRangeProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/RequestMetadataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/RequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableFormats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableRecommendedStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamUseCases.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerCroppingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerMultiResolutionStreamSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerPhysicalCameraMultiResolutionStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerRotateAndCrop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/SensorInfoColorFilterArrangement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/SensorInfoLensShadingApplied.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/SensorInfoTimestampSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/SensorPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/SensorRawBinningFactorUsed.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/SensorReferenceIlluminant1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/SensorTestPatternMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/ShadingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsFaceDetectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsHistogramMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsHotPixelMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsLensShadingMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsOisDataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsSceneFlicker.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsSharpnessMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/SyncFrameNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/SyncMaxLatency.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/TonemapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/android/hardware/camera/metadata/TonemapPresetCurve.cpp
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
    add_library(android.hardware.camera.metadata-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.metadata-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.metadata-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.metadata-V2-ndk_platform
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
#################### android.hardware.camera.metadata-V2-ndk_platform ####################

#################### android.hardware.camera.metadata-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/AutomotiveLensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/AutomotiveLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/BlackLevelLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/CameraMetadataSection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/CameraMetadataSectionStart.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/CameraMetadataTag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ColorCorrectionAberrationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ColorCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAeAntibandingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAeLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAePrecaptureTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAfMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAfSceneChange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAfState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAfTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAwbLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAwbLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAwbMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlAwbState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlCaptureIntent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlEffectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlEnableZsl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlExtendedSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ControlVideoStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/DemosaicMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/DepthDepthIsExclusive.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/DistortionCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/EdgeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/FlashInfoAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/FlashMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/FlashState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/HeicInfoSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/HotPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/InfoSupportedBufferManagementVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/InfoSupportedHardwareLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/LedAvailableLeds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/LedTransmit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/LensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/LensInfoFocusDistanceCalibration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/LensOpticalStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/LensPoseReference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/LensState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/LogicalMultiCameraSensorSyncType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/NoiseReductionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/QuirksPartialResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/RequestAvailableCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/RequestAvailableDynamicRangeProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/RequestMetadataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/RequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableFormats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableRecommendedStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamUseCases.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ScalerCroppingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ScalerMultiResolutionStreamSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ScalerPhysicalCameraMultiResolutionStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ScalerRotateAndCrop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/SensorInfoColorFilterArrangement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/SensorInfoLensShadingApplied.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/SensorInfoTimestampSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/SensorPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/SensorRawBinningFactorUsed.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/SensorReferenceIlluminant1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/SensorTestPatternMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/ShadingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/StatisticsFaceDetectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/StatisticsHistogramMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/StatisticsHotPixelMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/StatisticsLensShadingMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/StatisticsOisDataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/StatisticsSceneFlicker.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/StatisticsSharpnessMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/SyncFrameNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/SyncMaxLatency.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/TonemapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/android/hardware/camera/metadata/TonemapPresetCurve.cpp
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
    add_library(android.hardware.camera.metadata-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.metadata-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.metadata-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.metadata-V1-ndk
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
#################### android.hardware.camera.metadata-V1-ndk ####################

#################### android.hardware.camera.metadata-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/AutomotiveLensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/AutomotiveLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/BlackLevelLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/CameraMetadataSection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/CameraMetadataSectionStart.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/CameraMetadataTag.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ColorCorrectionAberrationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ColorCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeAntibandingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAePrecaptureTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAeState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfSceneChange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAfTrigger.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbLock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbLockAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlAwbState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlCaptureIntent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlEffectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlEnableZsl.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlExtendedSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlSceneMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ControlVideoStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/DemosaicMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/DepthAvailableDynamicDepthStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/DepthDepthIsExclusive.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/DistortionCorrectionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/EdgeMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/FlashInfoAvailable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/FlashMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/FlashState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/HeicAvailableHeicStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/HeicInfoSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/HotPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/InfoSupportedBufferManagementVersion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/InfoSupportedHardwareLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/LedAvailableLeds.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/LedTransmit.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/LensFacing.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/LensInfoFocusDistanceCalibration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/LensOpticalStabilizationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/LensPoseReference.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/LensState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/LogicalMultiCameraSensorSyncType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/NoiseReductionMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/QuirksPartialResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/RequestAvailableCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/RequestAvailableDynamicRangeProfilesMap.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/RequestMetadataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/RequestType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableFormats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableRecommendedStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamConfigurationsMaximumResolution.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerAvailableStreamUseCases.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerCroppingType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerMultiResolutionStreamSupported.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerPhysicalCameraMultiResolutionStreamConfigurations.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ScalerRotateAndCrop.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/SensorInfoColorFilterArrangement.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/SensorInfoLensShadingApplied.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/SensorInfoTimestampSource.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/SensorPixelMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/SensorRawBinningFactorUsed.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/SensorReferenceIlluminant1.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/SensorTestPatternMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/ShadingMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsFaceDetectMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsHistogramMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsHotPixelMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsLensShadingMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsOisDataMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsSceneFlicker.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/StatisticsSharpnessMapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/SyncFrameNumber.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/SyncMaxLatency.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/TonemapMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/android/hardware/camera/metadata/TonemapPresetCurve.cpp
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
    add_library(android.hardware.camera.metadata-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.metadata-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/metadata/aidl/android.hardware.camera.metadata-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.metadata-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.metadata-V1-ndk_platform
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
#################### android.hardware.camera.metadata-V1-ndk_platform ####################

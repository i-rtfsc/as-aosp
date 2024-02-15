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

project(libcameraservice_device_independent)

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

#################### libcameraservice_device_independent ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/common/DepthPhotoProcessor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/CoordinateMapper.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/DistortionMapper.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/RotateAndCropMapper.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/ZoomRatioMapper.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/ExifUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/SessionConfigurationUtilsHost.cpp
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
    add_library(libcameraservice_device_independent SHARED ${SOURCE_FILES})

    target_include_directories(libcameraservice_device_independent
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice
    )

    # Add c/c++ flags
    target_compile_options(libcameraservice_device_independent PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libcameraservice_device_independent
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libcamera_metadata
            libclang_rt.builtins
            libdl
            libdynamic_depth
            libdynamic_depth-internal_headers
            libdynamic_depth-public_headers
            libexif
            libjpeg
            liblog
            libm
            libutils
            libxml2
    )
endif ()
#################### libcameraservice_device_independent ####################

#################### libcameraservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/CameraService.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/CameraServiceWatchdog.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/CameraFlashlight.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/common/Camera2ClientBase.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/common/CameraDeviceBase.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/common/CameraOfflineSessionBase.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/common/CameraProviderManager.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/common/FrameProcessorBase.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/common/hidl/HidlProviderInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/common/aidl/AidlProviderInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api1/Camera2Client.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api1/client2/Parameters.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api1/client2/FrameProcessor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api1/client2/StreamingProcessor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api1/client2/JpegProcessor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api1/client2/CallbackProcessor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api1/client2/CaptureSequencer.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api1/client2/ZslProcessor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api2/CameraDeviceClient.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api2/CameraOfflineSessionClient.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api2/CompositeStream.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api2/DepthCompositeStream.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api2/HeicEncoderInfoManager.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api2/HeicCompositeStream.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/api2/JpegRCompositeStream.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/BufferUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3Device.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3OfflineSession.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3Stream.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3IOStreamBase.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3InputStream.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3OutputStream.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3FakeStream.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3SharedOutputStream.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/StatusTracker.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3BufferManager.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3StreamSplitter.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3OutputStreamInterface.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3OutputUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/Camera3DeviceInjectionMethods.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/UHRCropAndMeteringRegionMapper.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/PreviewFrameSpacer.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/hidl/HidlCamera3Device.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/hidl/HidlCamera3OfflineSession.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/hidl/HidlCamera3OutputUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/aidl/AidlCamera3Device.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/aidl/AidlCamera3OutputUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/device3/aidl/AidlCamera3OfflineSession.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/gui/RingBufferConsumer.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/hidl/AidlCameraDeviceCallbacks.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/hidl/AidlCameraServiceListener.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/hidl/HidlCameraDeviceUser.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/hidl/HidlCameraService.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/hidl/Utils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/aidl/AidlCameraDeviceCallbacks.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/aidl/AidlCameraDeviceUser.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/aidl/AidlCameraService.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/aidl/AidlCameraServiceListener.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/aidl/AidlUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/aidl/DeathPipe.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/CameraServiceProxyWrapper.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/CameraThreadState.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/CameraTraces.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/AutoConditionLock.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/SessionConfigurationUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/SessionConfigurationUtilsHidl.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/SessionStatsBuilder.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/TagMonitor.cpp
        ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice/utils/LatencyHistogram.cpp
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
    add_library(libcameraservice SHARED ${SOURCE_FILES})

    target_include_directories(libcameraservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/services/camera/libcameraservice
    )

    # Add c/c++ flags
    target_compile_options(libcameraservice PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libcameraservice
            android.frameworks.cameraservice.common-V1-ndk
            android.frameworks.cameraservice.common-2.0
            android.frameworks.cameraservice.device-V1-ndk
            android.frameworks.cameraservice.device-2.0
            android.frameworks.cameraservice.device-2.1
            android.frameworks.cameraservice.service-V1-ndk
            android.frameworks.cameraservice.service-2.0
            android.frameworks.cameraservice.service-2.1
            android.frameworks.cameraservice.service-2.2
            android.hardware.camera.common-1.0
            android.hardware.camera.device-V2-ndk
            android.hardware.camera.device-3.2
            android.hardware.camera.device-3.3
            android.hardware.camera.device-3.4
            android.hardware.camera.device-3.5
            android.hardware.camera.device-3.6
            android.hardware.camera.device-3.7
            android.hardware.camera.provider-V2-ndk
            android.hardware.camera.provider-2.4
            android.hardware.camera.provider-2.5
            android.hardware.camera.provider-2.6
            android.hardware.camera.provider-2.7
            libactivitymanager_aidl
            libaidlcommonsupport
            libbase
            libbinder
            libbinder_ndk
            libbinderthreadstateutils
            libc
            libc++
            libc++demangle
            libcamera_client
            libcamera_metadata
            libcameraservice_device_independent
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdynamic_depth
            libdynamic_depth-internal_headers
            libdynamic_depth-public_headers
            libexif
            libfmq
            libgui
            libhardware
            libhidlbase
            libimage_io
            libjpeg
            liblog
            libm
            libmedia
            libmedia_codeclist
            libmedia_omx
            libmediadrm_headers
            libmediametrics_headers
            libmediautils
            libmemunreachable
            libpermission
            libprocessinfoservice_aidl
            libsensorprivacy
            libstagefright
            libstagefright_foundation
            libui
            libultrahdr
            libutils
            libutilscallstack
            libxml2
            libyuv
            media_permission-aidl-cpp
    )
endif ()
#################### libcameraservice ####################

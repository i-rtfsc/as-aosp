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

project(android.hardware.camera.device-V2-ndk)

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

#################### android.hardware.camera.device-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/BufferCache.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/BufferRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/BufferRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/CameraBlob.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/CameraBlobId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/CameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/CameraOfflineSessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/CaptureRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/CaptureResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/ErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/ErrorMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/HalStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/ICameraDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/ICameraDeviceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/ICameraDeviceSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/ICameraInjectionSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/ICameraOfflineSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/NotifyMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/OfflineRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/OfflineStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/PhysicalCameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/PhysicalCameraSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/RequestTemplate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/ShutterMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/StreamBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/StreamBufferRequestError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/StreamBufferRet.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/StreamBuffersVal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/StreamConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/StreamConfigurationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/StreamRotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/android/hardware/camera/device/StreamType.cpp
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
    add_library(android.hardware.camera.device-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.device-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.device-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.device-V2-ndk
            android.hardware.camera.common-V1-ndk
            android.hardware.camera.metadata-V2-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.common-V4-ndk
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
#################### android.hardware.camera.device-V2-ndk ####################

#################### android.hardware.camera.device-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/BufferCache.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/BufferRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/BufferRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/CameraBlob.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/CameraBlobId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/CameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/CameraOfflineSessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/CaptureRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/CaptureResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/ErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/ErrorMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/HalStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/ICameraDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/ICameraDeviceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/ICameraDeviceSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/ICameraInjectionSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/ICameraOfflineSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/NotifyMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/OfflineRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/OfflineStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/PhysicalCameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/PhysicalCameraSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/RequestTemplate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/ShutterMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/StreamBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/StreamBufferRequestError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/StreamBufferRet.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/StreamBuffersVal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/StreamConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/StreamConfigurationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/StreamRotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/android/hardware/camera/device/StreamType.cpp
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
    add_library(android.hardware.camera.device-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.device-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.device-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.device-V3-ndk
            android.hardware.camera.common-V1-ndk
            android.hardware.camera.metadata-V2-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.common-V4-ndk
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
#################### android.hardware.camera.device-V3-ndk ####################

#################### android.hardware.camera.device-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/BufferCache.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/BufferRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/BufferRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/CameraBlob.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/CameraBlobId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/CameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/CameraOfflineSessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/CaptureRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/CaptureResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/ErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/ErrorMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/HalStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/ICameraDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/ICameraDeviceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/ICameraDeviceSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/ICameraInjectionSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/ICameraOfflineSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/NotifyMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/OfflineRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/OfflineStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/PhysicalCameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/PhysicalCameraSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/RequestTemplate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/ShutterMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/StreamBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/StreamBufferRequestError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/StreamBufferRet.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/StreamBuffersVal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/StreamConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/StreamConfigurationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/StreamRotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/android/hardware/camera/device/StreamType.cpp
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
    add_library(android.hardware.camera.device-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.device-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.device-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.device-V2-ndk_platform
            android.hardware.camera.common-V1-ndk_platform
            android.hardware.camera.metadata-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### android.hardware.camera.device-V2-ndk_platform ####################

#################### android.hardware.camera.device-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/BufferCache.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/BufferRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/BufferRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/CameraBlob.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/CameraBlobId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/CameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/CameraOfflineSessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/CaptureRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/CaptureResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/ErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/ErrorMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/HalStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/ICameraDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/ICameraDeviceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/ICameraDeviceSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/ICameraInjectionSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/ICameraOfflineSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/NotifyMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/OfflineRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/OfflineStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/PhysicalCameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/PhysicalCameraSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/RequestTemplate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/ShutterMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/StreamBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/StreamBufferRequestError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/StreamBufferRet.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/StreamBuffersVal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/StreamConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/StreamConfigurationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/StreamRotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/android/hardware/camera/device/StreamType.cpp
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
    add_library(android.hardware.camera.device-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.device-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.device-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.device-V3-ndk_platform
            android.hardware.camera.common-V1-ndk_platform
            android.hardware.camera.metadata-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### android.hardware.camera.device-V3-ndk_platform ####################

#################### android.hardware.camera.device-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/BufferCache.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/BufferRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/BufferRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/CameraBlob.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/CameraBlobId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/CameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/CameraOfflineSessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/CaptureRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/CaptureResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/ErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/ErrorMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/HalStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/ICameraDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/ICameraDeviceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/ICameraDeviceSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/ICameraInjectionSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/ICameraOfflineSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/NotifyMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/OfflineRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/OfflineStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/PhysicalCameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/PhysicalCameraSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/RequestTemplate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/ShutterMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/StreamBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/StreamBufferRequestError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/StreamBufferRet.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/StreamBuffersVal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/StreamConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/StreamConfigurationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/StreamRotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/android/hardware/camera/device/StreamType.cpp
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
    add_library(android.hardware.camera.device-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.device-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.device-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.device-V1-ndk
            android.hardware.camera.common-V1-ndk
            android.hardware.camera.metadata-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.common-V4-ndk
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
#################### android.hardware.camera.device-V1-ndk ####################

#################### android.hardware.camera.device-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/BufferCache.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/BufferRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/BufferRequestStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/CameraBlob.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/CameraBlobId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/CameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/CameraOfflineSessionInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/CaptureRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/CaptureResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/ErrorCode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/ErrorMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/HalStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/ICameraDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/ICameraDeviceCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/ICameraDeviceSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/ICameraInjectionSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/ICameraOfflineSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/NotifyMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/OfflineRequest.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/OfflineStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/PhysicalCameraMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/PhysicalCameraSetting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/RequestTemplate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/ShutterMsg.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/StreamBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/StreamBufferRequestError.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/StreamBufferRet.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/StreamBuffersVal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/StreamConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/StreamConfigurationMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/StreamRotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/android/hardware/camera/device/StreamType.cpp
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
    add_library(android.hardware.camera.device-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.device-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/camera/device/aidl/android.hardware.camera.device-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.device-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.camera.device-V1-ndk_platform
            android.hardware.camera.common-V1-ndk_platform
            android.hardware.camera.metadata-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### android.hardware.camera.device-V1-ndk_platform ####################

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

project(libcamera_client)

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

#################### libcamera_client ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl/android/hardware/CameraExtensionSessionStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl/android/hardware/ICameraService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl/android/hardware/ICameraServiceListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl/android/hardware/ICameraServiceProxy.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl/android/hardware/camera2/ICameraDeviceCallbacks.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl/android/hardware/camera2/ICameraDeviceUser.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl/android/hardware/camera2/ICameraInjectionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl/android/hardware/camera2/ICameraInjectionSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl/android/hardware/camera2/ICameraOfflineSession.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/Camera.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/CameraMetadata.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/CameraParameters.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/CaptureResult.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/CameraParameters2.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/CameraSessionStats.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ICamera.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ICameraClient.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ICameraRecordingProxy.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/camera2/CaptureRequest.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/camera2/ConcurrentCamera.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/camera2/OutputConfiguration.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/camera2/SessionConfiguration.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/camera2/SubmitInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/CameraBase.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/CameraUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/VendorTagDescriptor.cpp
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
    add_library(libcamera_client SHARED ${SOURCE_FILES})

    target_include_directories(libcamera_client
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/camera/include
            ${ANDROID_ROOT}/frameworks/av/camera/include/camera
            ${ANDROID_ROOT}/frameworks/av/camera
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/camera/libcamera_client/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libcamera_client PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(libcamera_client
            lib-platform-compat-native-api
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libcamera_metadata
            libclang_rt.builtins
            libcutils
            libdl
            libgui
            liblog
            libm
            libnativewindow
            libutils
    )
endif ()
#################### libcamera_client ####################

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

project(libcamera2ndk)

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

#################### libcamera2ndk ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCameraManager.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCameraMetadata.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCameraDevice.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCaptureRequest.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCameraCaptureSession.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/impl/ACameraManager.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/impl/ACameraMetadata.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/impl/ACameraDevice.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/impl/ACameraCaptureSession.cpp
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
    add_library(libcamera2ndk SHARED ${SOURCE_FILES})

    target_include_directories(libcamera2ndk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/camera/ndk/include
            ${ANDROID_ROOT}/frameworks/av/camera/ndk
    )

    # Add c/c++ flags
    target_compile_options(libcamera2ndk PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libcamera2ndk
            jni_headers
            libandroid_runtime
            libbinder
            libc
            libc++
            libc++demangle
            libcamera_client
            libcamera_metadata
            libclang_rt.builtins
            libcutils
            libdl
            libgui
            liblog
            libm
            libmediandk
            libnativewindow
            libstagefright_foundation
            libutils
    )
endif ()
#################### libcamera2ndk ####################

#################### libcamera2ndk_vendor ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/ndk_vendor/impl/ACameraDevice.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/ndk_vendor/impl/ACameraManager.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/ndk_vendor/impl/utils.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/impl/ACameraMetadata.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/impl/ACameraCaptureSession.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCameraMetadata.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCameraCaptureSession.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCameraManager.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCameraDevice.cpp
        ${ANDROID_ROOT}/frameworks/av/camera/ndk/NdkCaptureRequest.cpp
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
    add_library(libcamera2ndk_vendor SHARED ${SOURCE_FILES})

    target_include_directories(libcamera2ndk_vendor
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/camera/ndk/include
            ${ANDROID_ROOT}/frameworks/av/camera/ndk
            ${ANDROID_ROOT}/frameworks/av/camera/ndk/impl
    )

    # Add c/c++ flags
    target_compile_options(libcamera2ndk_vendor PRIVATE
            -Werror
    )

    target_link_libraries(libcamera2ndk_vendor
            android.frameworks.cameraservice.common-V1-ndk
            android.frameworks.cameraservice.device-V1-ndk
            android.frameworks.cameraservice.service-V1-ndk
            android.hardware.camera.common-1.0-helper
            libaidlcommonsupport
            libarect
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libcamera_metadata
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware
            libhidlbase
            liblog
            libm
            libmediandk
            libnativewindow
            libstagefright_foundation
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libcamera2ndk_vendor ####################

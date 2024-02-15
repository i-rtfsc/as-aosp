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

project(android.hardware.camera.provider@2.7-external-vsock-service)

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

#################### android.hardware.camera.provider-2.7-external-vsock-service ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera/external-service.cpp
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
    add_executable(android.hardware.camera.provider-2.7-external-vsock-service ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.provider-2.7-external-vsock-service
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.provider-2.7-external-vsock-service PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.camera.provider-2.7-external-vsock-service
            android.hardware.camera.provider-2.7
            libbinder
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.camera.provider-2.7-external-vsock-service ####################

#################### android.hardware.camera.provider-2.7-impl-cuttlefish ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera/vsock_camera_provider_2_7.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera/vsock_camera_device_3_4.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera/vsock_camera_device_session_3_4.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera/vsock_camera_metadata.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera/vsock_camera_server.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera/vsock_frame_provider.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera/cached_stream_buffer.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera/stream_buffer_cache.cpp
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
    add_library(android.hardware.camera.provider-2.7-impl-cuttlefish SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.camera.provider-2.7-impl-cuttlefish
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/camera
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.camera.provider-2.7-impl-cuttlefish PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.camera.provider-2.7-impl-cuttlefish
            android.hardware.camera.common-1.0-helper
            android.hardware.camera.common-1.0
            android.hardware.camera.device-1.0
            android.hardware.camera.device-3.2
            android.hardware.camera.device-3.3
            android.hardware.camera.device-3.4
            android.hardware.camera.device-3.5
            android.hardware.camera.provider-2.4-external
            android.hardware.camera.provider-2.4-legacy
            android.hardware.camera.provider-2.4
            android.hardware.camera.provider-2.5
            android.hardware.camera.provider-2.6
            android.hardware.camera.provider-2.7
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            camera.device-1.0-impl
            camera.device-3.2-impl
            camera.device-3.3-impl
            camera.device-3.4-external-impl_headers
            camera.device-3.4-impl
            camera.device-3.4-impl_headers
            camera.device-3.5-external-impl_headers
            camera.device-3.5-impl_headers
            libc++
            libc++demangle
            libc
            libcamera_metadata
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libcuttlefish_fs
            libdl
            libfmq
            libgralloctypes
            libhardware
            libhidlbase
            libjsoncpp
            liblog
            libm
            libsync
            libutils
            libvsock_utils
            libyuv
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.camera.provider-2.7-impl-cuttlefish ####################

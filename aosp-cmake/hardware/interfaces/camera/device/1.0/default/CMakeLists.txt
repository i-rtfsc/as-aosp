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

project(camera.device@1.0-impl)

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

#################### camera.device-1.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/camera/device/1.0/default/CameraDevice.cpp
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
    add_library(camera.device-1.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(camera.device-1.0-impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/camera/device/1.0/default
    )

    # Add c/c++ flags
    target_compile_options(camera.device-1.0-impl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(camera.device-1.0-impl
            android.hardware.camera.common-1.0-helper
            android.hardware.camera.common-1.0
            android.hardware.camera.device-1.0
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.common-1.0
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libc++
            libc++demangle
            libc
            libcamera_metadata
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgralloctypes
            libhardware
            libhidlbase
            libhidlmemory
            liblog
            libm
            libutils
            media_plugin_headers
            vendor.qti.hardware.camera.device-1.0
            libc
            libdl
            libm
    )
endif ()
#################### camera.device-1.0-impl ####################

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

project(android.hardware.automotive.evs@1.1-sample)

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

#################### libevsconfigmanager ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/ConfigManager.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/ConfigManagerUtil.cpp
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
    add_library(libevsconfigmanager SHARED ${SOURCE_FILES})

    target_include_directories(libevsconfigmanager
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl
    )

    # Add c/c++ flags
    target_compile_options(libevsconfigmanager PRIVATE
            -Werror
    )

    target_link_libraries(libevsconfigmanager
            android.hardware.automotive.evs-1.0
            android.hardware.automotive.evs-1.1
            android.hardware.camera.device-3.2
            libbase
            libc++
            libc++_static
            libc++demangle
            libc
            libcamera_metadata
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libm
            libtinyxml2
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libevsconfigmanager ####################

#################### android.hardware.automotive.evs-1.1-sample ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/android.hardware.automotive.evs@1.1-sample/android_vendor.34_${OUT_ARCH_CPU}/gen/aidl/android/gui/FrameEvent.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/service.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/EvsEnumerator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/EvsV4lCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/EvsGlDisplay.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/GlWrapper.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/VideoCapture.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/bufferCopy.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/ConfigManager.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/ConfigManagerUtil.cpp
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
    add_executable(android.hardware.automotive.evs-1.1-sample ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.evs-1.1-sample
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/sampleDriver/hidl/android.hardware.automotive.evs@1.1-sample/android_vendor.34_${OUT_ARCH_CPU}/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.evs-1.1-sample PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.automotive.evs-1.1-sample
            android.frameworks.automotive.display-1.0
            android.hardware.automotive.evs-1.0
            android.hardware.automotive.evs-1.1
            android.hardware.camera.device-3.2
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hidl.token-1.0-utils
            libEGL
            libGLESv2
            libbase
            libbinder
            libbufferqueueconverter
            libc++
            libc++demangle
            libc
            libcamera_metadata
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhardware_legacy
            libhidlbase
            libm
            libtinyxml2
            libui
            libutils
            libyuv
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.automotive.evs-1.1-sample ####################

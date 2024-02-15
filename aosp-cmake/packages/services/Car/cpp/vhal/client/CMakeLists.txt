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

project(libvhalclient)

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

#################### libvhalclient ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/src/AidlHalPropConfig.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/src/AidlHalPropValue.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/src/AidlVhalClient.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/src/HidlHalPropConfig.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/src/HidlHalPropValue.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/src/HidlVhalClient.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/src/IVhalClient.cpp
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
    add_library(libvhalclient SHARED ${SOURCE_FILES})

    target_include_directories(libvhalclient
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/vhal/client
    )

    # Add c/c++ flags
    target_compile_options(libvhalclient PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libvhalclient
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libmath
            libutils
    )
endif ()
#################### libvhalclient ####################

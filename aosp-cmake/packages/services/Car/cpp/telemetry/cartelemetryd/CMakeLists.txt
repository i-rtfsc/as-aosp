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

project(android.automotive.telemetryd@1.0)

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

#################### android.automotive.telemetryd-1.0 ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/src/main.cpp
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
    add_executable(android.automotive.telemetryd-1.0 ${SOURCE_FILES})

    target_include_directories(android.automotive.telemetryd-1.0
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/aidl/android.automotive.telemetry.internal-V2-ndk-source/gen/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/src
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.telemetryd-1.0 PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(android.automotive.telemetryd-1.0
            android.automotive.telemetry.internal-V2-ndk
            android.automotive.telemetryd-1.0-impl
            android.frameworks.automotive.telemetry-V2-ndk
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            liblog
            libm
            libutils
    )
endif ()
#################### android.automotive.telemetryd-1.0 ####################

#################### android.automotive.telemetryd-1.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/src/CarTelemetryImpl.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/src/CarTelemetryInternalImpl.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/src/LooperWrapper.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/src/RingBuffer.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/src/TelemetryServer.cpp
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
    add_library(android.automotive.telemetryd-1.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.telemetryd-1.0-impl
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/src
            ${ANDROID_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/telemetry/cartelemetryd/aidl/android.automotive.telemetry.internal-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.telemetryd-1.0-impl PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(android.automotive.telemetryd-1.0-impl
            android.automotive.telemetry.internal-V2-ndk
            android.frameworks.automotive.telemetry-V2-ndk
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgtest_prod_headers
            liblog
            libm
            libutils
    )
endif ()
#################### android.automotive.telemetryd-1.0-impl ####################

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

project(android.hardware.dumpstate@1.1-service.trout)

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

#################### android.hardware.dumpstate-1.1-service.trout ####################
set(FILES
        ${ANDROID_ROOT}/device/google/trout/hal/dumpstate/1.1/DumpstateDevice.cpp
        ${ANDROID_ROOT}/device/google/trout/hal/dumpstate/1.1/service.cpp
        ${ANDROID_ROOT}/device/google/trout/hal/dumpstate/1.1/WatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/dumpstate/1.1/config/dumpstate_hal_configuration_V1_0/gen/cpp/dumpstate_hal_configuration_V1_0.cpp
        ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/dumpstate/1.1/config/dumpstate_hal_configuration_V1_0/gen/cpp/dumpstate_hal_configuration_V1_0_enums.cpp
        ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/dumpstate/1.1/DumpstateServerProtoStub_cc/gen/DumpstateServer.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/dumpstate/1.1/DumpstateServerProtoStub_cc/gen/DumpstateServer.grpc.pb.cc
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
    add_executable(android.hardware.dumpstate-1.1-service.trout ${SOURCE_FILES})

    target_include_directories(android.hardware.dumpstate-1.1-service.trout
            PUBLIC
            ${ANDROID_ROOT}/device/google/trout/hal/dumpstate/1.1
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/dumpstate/1.1/config/dumpstate_hal_configuration_V1_0/gen/cpp/include
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/dumpstate/1.1/DumpstateServerProtoStub_cc/gen
            ${OUT_INTERMEDIATES_ROOT}/device/google/trout/hal/dumpstate/1.1/DumpstateServerProtoStub_h/gen
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.dumpstate-1.1-service.trout PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(android.hardware.dumpstate-1.1-service.trout
            android.automotive.watchdog-V2-ndk
            android.hardware.automotive.utils.vsockinfo
            android.hardware.automotive.utils.watchdog
            android.hardware.automotive-libc++fs
            android.hardware.dumpstate-1.0
            android.hardware.dumpstate-1.1
            cuttlefish_common_headers
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdumpstateutil
            libgrpc++
            libhidlbase
            liblog
            libm
            libprotobuf-cpp-full
            libutils
            libxml2
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.dumpstate-1.1-service.trout ####################

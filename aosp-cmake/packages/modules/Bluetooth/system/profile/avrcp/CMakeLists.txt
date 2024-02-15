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

project(avrcp_device_fuzz)

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

#################### avrcp-target-service ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/profile/avrcp/connection_handler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/profile/avrcp/device.cc
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
    add_library(avrcp-target-service SHARED ${SOURCE_FILES})

    target_include_directories(avrcp-target-service
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/profile/avrcp
    )

    # Add c/c++ flags
    target_compile_options(avrcp-target-service PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(avrcp-target-service
            lib-bt-packets
            libbase
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-platform-protos-lite
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libevent
            libhidlbase
            liblog
            libm
            libmodpb64
            libosi
            libstatslog_bt
            libstatssocket
            libutils
    )
endif ()
#################### avrcp-target-service ####################

#################### avrcp_device_fuzz ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/profile/avrcp/tests/avrcp_device_fuzz/avrcp_device_fuzz.cc
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
    add_executable(avrcp_device_fuzz ${SOURCE_FILES})

    target_include_directories(avrcp_device_fuzz
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/profile/avrcp
    )

    # Add c/c++ flags
    target_compile_options(avrcp_device_fuzz PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(avrcp_device_fuzz
            avrcp-target-service
            lib-bt-packets
            lib-bt-packets-avrcp
            lib-bt-packets-base
            libbase
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-platform-protos-lite
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libevent
            liblog
            libm
            libmodpb64
            libstatslog_bt
            libutils
    )
endif ()
#################### avrcp_device_fuzz ####################

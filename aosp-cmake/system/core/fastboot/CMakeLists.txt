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

project(fastbootd)

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

#################### fastbootd ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fastboot/device/commands.cpp
        ${ANDROID_ROOT}/system/core/fastboot/device/fastboot_device.cpp
        ${ANDROID_ROOT}/system/core/fastboot/device/flashing.cpp
        ${ANDROID_ROOT}/system/core/fastboot/device/main.cpp
        ${ANDROID_ROOT}/system/core/fastboot/device/usb.cpp
        ${ANDROID_ROOT}/system/core/fastboot/device/usb_iouring.cpp
        ${ANDROID_ROOT}/system/core/fastboot/device/usb_client.cpp
        ${ANDROID_ROOT}/system/core/fastboot/device/tcp_client.cpp
        ${ANDROID_ROOT}/system/core/fastboot/device/utility.cpp
        ${ANDROID_ROOT}/system/core/fastboot/device/variables.cpp
        ${ANDROID_ROOT}/system/core/fastboot/socket.cpp
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
    add_executable(fastbootd ${SOURCE_FILES})

    target_include_directories(fastbootd
            PUBLIC
            ${ANDROID_ROOT}/system/core/fastboot
    )

    # Add c/c++ flags
    target_compile_options(fastbootd PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(fastbootd
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            android.hardware.fastboot-V1-ndk
            android.hardware.fastboot-1.1
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            android.hardware.health-2.0
            avb_headers
            libasyncio
            libbase
            libbinder_ndk
            libboot_control_client
            libbootloader_message
            libc++
            libc++demangle
            libc++fs
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libext2_uuid
            libext4_utils
            libfastbootshim
            libfs_mgr
            libgsi
            libgtest_prod_headers
            libhealthhalutils
            libhealthshim
            libhidlbase
            liblog
            liblp
            liblz4
            libm
            libprotobuf-cpp-lite
            libsnapshot_cow
            libsnapshot_headers
            libsnapshot_nobinder
            libsparse
            libstorage_literals_headers
            liburing
            libutils
            update_metadata-protos
            libc
            libdl
            libm
    )
endif ()
#################### fastbootd ####################

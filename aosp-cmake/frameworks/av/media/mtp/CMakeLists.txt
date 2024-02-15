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

project(libmtp_fuzz)

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

#################### libmtp_fuzz ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDataPacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDebug.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDescriptors.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDevHandle.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDevice.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDeviceInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpEventPacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpFfsCompatHandle.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpFfsHandle.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpObjectInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpPacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpProperty.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpRequestPacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpResponsePacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpServer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpStorage.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpStorageInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpStringBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/PosixAsyncIO.cpp
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
    add_library(libmtp_fuzz SHARED ${SOURCE_FILES})

    target_include_directories(libmtp_fuzz
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/mtp
    )

    # Add c/c++ flags
    target_compile_options(libmtp_fuzz PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libmtp_fuzz
            libasyncio
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils_headers
            libdl
            liblog
            libm
            libusbhost
    )
endif ()
#################### libmtp_fuzz ####################

#################### libmtp ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDataPacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDebug.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDescriptors.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDevHandle.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDevice.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpDeviceInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpEventPacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpFfsCompatHandle.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpFfsHandle.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpObjectInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpPacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpProperty.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpRequestPacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpResponsePacket.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpServer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpStorage.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpStorageInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpStringBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/MtpUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/mtp/PosixAsyncIO.cpp
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
    add_library(libmtp SHARED ${SOURCE_FILES})

    target_include_directories(libmtp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/mtp
    )

    # Add c/c++ flags
    target_compile_options(libmtp PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libmtp
            libasyncio
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils_headers
            libdl
            liblog
            libm
            libusbhost
    )
endif ()
#################### libmtp ####################

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

project(libimsmedia_protocol)

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

#################### libimsmedia_protocol ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/RtpImpl.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/RtpService.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpAppPacket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpByePacket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpChunk.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpConfigInfo.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpFbPacket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpHeader.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpPacket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpReportBlock.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpRrPacket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpSdesPacket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpSrPacket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtcpXrPacket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpHeader.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpPacket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpPayloadInfo.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpReceiverInfo.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpSession.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpSessionManager.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpStack.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpStackProfile.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpStackUtil.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/core/RtpTimerInfo.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/utils/RtpBuffer.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/rtp/utils/RtpOsUtil.cpp
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
    add_library(libimsmedia_protocol SHARED ${SOURCE_FILES})

    target_include_directories(libimsmedia_protocol
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/include/rtp/core
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/include/rtp/utils
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/protocol/interface/rtp
    )

    # Add c/c++ flags
    target_compile_options(libimsmedia_protocol PRIVATE
            -Werror
            -DRTP_DEBUG_ENABLE
            -Wall
    )

    target_link_libraries(libimsmedia_protocol
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libimsmedia_core_interface_headers
            libimsmedia_protocol_headers
            libm
    )
endif ()
#################### libimsmedia_protocol ####################

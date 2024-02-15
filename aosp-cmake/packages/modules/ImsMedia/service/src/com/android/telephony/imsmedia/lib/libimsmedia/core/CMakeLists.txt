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

project(libimsmedia_core)

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

#################### libimsmedia_core ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/BaseJitterBuffer.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/BaseManager.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/BaseSession.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/BaseStreamGraph.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/IRtpSession.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/JitterNetworkAnalyser.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/StreamScheduler.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/AudioJitterBuffer.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/AudioManager.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/AudioSession.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/AudioStreamGraphRtcp.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/AudioStreamGraphRtpRx.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/AudioStreamGraphRtpTx.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/MediaQualityAnalyzer.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/RtcpXrEncoder.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/nodes/AudioRtpPayloadDecoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/nodes/AudioRtpPayloadEncoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/nodes/DtmfEncoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/nodes/IAudioPlayerNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/nodes/IAudioSourceNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/nodes/ImsMediaAudioUtil.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/nodes/BaseNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/nodes/JitterBufferControlNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/nodes/RtcpDecoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/nodes/RtcpEncoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/nodes/RtpDecoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/nodes/RtpEncoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/nodes/SocketReaderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/nodes/SocketWriterNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/TextJitterBuffer.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/TextManager.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/TextSession.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/TextStreamGraphRtcp.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/TextStreamGraphRtpRx.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/TextStreamGraphRtpTx.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/nodes/TextRendererNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/nodes/TextRtpPayloadDecoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/nodes/TextRtpPayloadEncoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/text/nodes/TextSourceNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/IImsMediaThread.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ISocket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaBinaryFormat.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaBitReader.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaBitWriter.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaCondition.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaDataQueue.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaEventHandler.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaImageRotate.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaNetworkUtil.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaSocket.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaTimer.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/utils/ImsMediaTrace.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/ImsMediaVideoUtil.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/VideoJitterBuffer.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/VideoManager.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/VideoSession.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/VideoStreamGraphRtcp.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/VideoStreamGraphRtpRx.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/VideoStreamGraphRtpTx.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/android/ImsMediaCamera.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/android/ImsMediaPauseImageSource.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/android/ImsMediaVideoRenderer.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/android/ImsMediaVideoSource.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/nodes/IVideoRendererNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/nodes/IVideoSourceNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/nodes/VideoRtpPayloadDecoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/video/nodes/VideoRtpPayloadEncoderNode.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/jni/MediaManagerFactory.cpp
        ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/jni/libimsmediajni.cpp
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
    add_library(libimsmedia_core SHARED ${SOURCE_FILES})

    target_include_directories(libimsmedia_core
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/audio/android/include
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include/text
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include/text/nodes
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include/video
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include/video/nodes
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include/video/android
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include/nodes
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include/utils
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include/audio
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/include/audio/nodes
            ${ANDROID_ROOT}/packages/modules/ImsMedia/service/src/com/android/telephony/imsmedia/lib/libimsmedia/core/interface/utils
    )

    # Add c/c++ flags
    target_compile_options(libimsmedia_core PRIVATE
            -Werror
            -DRTP_DEBUG_ENABLE
            -Wall
    )

    target_link_libraries(libimsmedia_core
            framework-permission-aidl-cpp
            libaaudio
            libandroid
            libandroid_runtime
            libbinder
            libc
            libc++_static
            libc++demangle
            libcamera2ndk
            libclang_rt.builtins
            libcutils
            libdl
            libimsmedia_aosp_audio_codec_headers
            libimsmedia_audio_headers
            libimsmedia_config
            libimsmedia_config_headers
            libimsmedia_core_interface_headers
            libimsmedia_headers
            libimsmedia_jni_headers
            libimsmedia_protocol
            libimsmedia_protocol_headers
            libimsmedia_protocol_interface_headers
            libimsmedia_text_headers
            libimsmedia_video_headers
            libjnigraphics
            liblog
            libm
            libmediandk
            libnativewindow
            libutils
    )
endif ()
#################### libimsmedia_core ####################

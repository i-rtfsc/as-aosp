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

project(libmediaplayerservice)

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

#################### libmediaplayerservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/ActivityManager.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/DeathNotifier.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/MediaPlayerFactory.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/MediaPlayerService.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/MediaRecorderClient.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/MetadataRetrieverClient.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/StagefrightMetadataRetriever.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/StagefrightRecorder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/TestPlayerStub.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/mpctl/PerfBoost.cpp
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
    add_library(libmediaplayerservice SHARED ${SOURCE_FILES})

    target_include_directories(libmediaplayerservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/include
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/datasource/include
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/include
    )

    # Add c/c++ flags
    target_compile_options(libmediaplayerservice PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmediaplayerservice
            android.hardware.media.c2-1.0
            android.hardware.media.omx-1.0
            av-types-aidl-cpp
            framework-permission-aidl-cpp
            libactivitymanager_aidl
            libandroid_net
            libaudioclient
            libaudioclient_aidl_conversion
            libavmediaserviceextensions
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libcamera_client
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcodec2_client
            libcrypto
            libcutils
            libdatasource
            libdl
            libdrmframework
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_codeclist
            libmedia_omx
            libmediadrm
            libmediametrics
            libmediautils
            libmediautils_headers
            libmemunreachable
            libnetd_client
            libplayerservice_datasource
            libpowermanager
            libstagefright
            libstagefright_foundation
            libstagefright_httplive
            libstagefright_mpeg2support_headers
            libstagefright_nuplayer
            libstagefright_omx
            libstagefright_rtsp
            libstagefright_rtsp_headers
            libstagefright_timedtext
            libstagefright_webm_headers
            libutils
            media_plugin_headers
            packagemanager_aidl-cpp
    )
endif ()
#################### libmediaplayerservice ####################

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

project(libstagefright_nuplayer)

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

#################### libstagefright_nuplayer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/AWakeLock.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/GenericSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/HTTPLiveSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/NuPlayer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/NuPlayerCCDecoder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/NuPlayerDecoder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/NuPlayerDecoderBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/NuPlayerDecoderPassThrough.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/NuPlayerDriver.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/NuPlayerDrm.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/NuPlayerRenderer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/NuPlayerStreamListener.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/RTSPSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/RTPSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/StreamingSource.cpp
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
    add_library(libstagefright_nuplayer SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_nuplayer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/include
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer/include/nuplayer
            ${ANDROID_ROOT}/frameworks/av/media/libmediaplayerservice/nuplayer
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_nuplayer PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_nuplayer
            android.hardware.drm-V1-ndk
            libbinder
            libc
            libc++_static
            libc++demangle
            libcamera_client
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdatasource
            libdl
            libgui
            libhidlmemory
            libm
            libmedia
            libmediadrm
            libmediadrm_headers
            libmediametrics_headers
            libplayerservice_datasource
            libpowermanager
            libstagefright_esds
            libstagefright_headers
            libstagefright_httplive_headers
            libstagefright_mpeg2support_headers
            libstagefright_rtsp_headers
            libstagefright_timedtext
            libui
            media_plugin_headers
    )
endif ()
#################### libstagefright_nuplayer ####################

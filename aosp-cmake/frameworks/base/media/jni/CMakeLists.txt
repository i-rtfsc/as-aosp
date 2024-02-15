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

project(libmedia_jni_utils)

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

#################### libmedia_jni_utils ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/jni/libmedia_jni_utils/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/gui/FrameEvent.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_Utils.cpp
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
    add_library(libmedia_jni_utils SHARED ${SOURCE_FILES})

    target_include_directories(libmedia_jni_utils
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/media/jni
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/media/jni/libmedia_jni_utils/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libmedia_jni_utils PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libmedia_jni_utils
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libgui_headers
            liblog
            libm
            libui
            libutils
    )
endif ()
#################### libmedia_jni_utils ####################

#################### libmedia_jni ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_ImageWriter.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_ImageReader.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_JetPlayer.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaCrypto.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaCodec.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaCodecList.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaDataSource.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaDescrambler.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaDrm.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaExtractor.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaHTTPConnection.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaMetadataRetriever.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaMuxer.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaPlayer.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaProfiles.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaRecorder.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_MediaSync.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_PublicFormatUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_ResampleInputStream.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_Streams.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_SyncParams.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_mtp_MtpDatabase.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_mtp_MtpDevice.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_mtp_MtpServer.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/JetPlayer.cpp
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
    add_library(libmedia_jni SHARED ${SOURCE_FILES})

    target_include_directories(libmedia_jni
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/media/jni
    )

    # Add c/c++ flags
    target_compile_options(libmedia_jni PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libmedia_jni
            android.hardware.cas.native-1.0
            android.hardware.cas-1.0
            android.hardware.drm-V1-ndk
            android.hardware.drm-1.3
            android.hardware.drm-1.4
            android.hardware.tv.tuner-V2-ndk
            android.hidl.memory-1.0
            android.hidl.token-1.0-utils
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            framework-permission-aidl-cpp
            libandroid_runtime
            libandroidfw
            libaudioclient
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libcamera_client
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcodec2
            libcodec2_internal
            libcodec2_vndk
            libcutils
            libdl
            libgrallocusage
            libgui
            libhardware_headers
            libhidlallocatorutils
            libhidlbase
            libhwui
            liblog
            libm
            libmedia
            libmedia_codeclist
            libmedia_jni_utils
            libmedia_midiiowrapper
            libmedia_omx
            libmediadrm
            libmediadrm_headers
            libmediadrmmetrics_consumer
            libmediametrics
            libmtp
            libnativehelper
            libnativewindow
            libpiex
            libprocessgroup
            libsonivox
            libstagefright
            libstagefright_foundation
            libui
            libutils
    )
endif ()
#################### libmedia_jni ####################

#################### libmedia_tv_tuner ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/media/jni/android_media_tv_Tuner.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/tuner/DemuxClient.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/tuner/DescramblerClient.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/tuner/DvrClient.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/tuner/FilterClient.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/tuner/FrontendClient.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/tuner/LnbClient.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/tuner/TimeFilterClient.cpp
        ${ANDROID_ROOT}/frameworks/base/media/jni/tuner/TunerClient.cpp
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
    add_library(libmedia_tv_tuner SHARED ${SOURCE_FILES})

    target_include_directories(libmedia_tv_tuner
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/media/jni
    )

    # Add c/c++ flags
    target_compile_options(libmedia_tv_tuner PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libmedia_tv_tuner
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.tv.tuner-V2-ndk
            libaidlcommonsupport
            libandroid_runtime
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcodec2
            libcodec2_internal
            libcodec2_vndk
            libcutils
            libdl
            libfmq
            libhidlbase
            liblog
            libm
            libmedia
            libnativehelper
            libstagefright_foundation_headers
            libutils
            tv_tuner_aidl_interface-ndk
    )
endif ()
#################### libmedia_tv_tuner ####################

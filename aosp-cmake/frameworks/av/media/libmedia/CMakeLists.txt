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

project(libmedia_midiiowrapper)

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

#################### libmedia_midiiowrapper ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/MidiIoWrapper.cpp
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
    add_library(libmedia_midiiowrapper SHARED ${SOURCE_FILES})

    target_include_directories(libmedia_midiiowrapper
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmedia
            ${ANDROID_ROOT}/frameworks/av/media/libmedia/include
    )

    # Add c/c++ flags
    target_compile_options(libmedia_midiiowrapper PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmedia_midiiowrapper
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libmedia_datasource_headers
            libsonivoxwithoutjet
            media_ndk_headers
    )
endif ()
#################### libmedia_midiiowrapper ####################

#################### libmedia_codeclist ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaCodecList.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/MediaCodecInfo.cpp
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
    add_library(libmedia_codeclist SHARED ${SOURCE_FILES})

    target_include_directories(libmedia_codeclist
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmedia/include
            ${ANDROID_ROOT}/frameworks/av/media/libmedia
    )

    # Add c/c++ flags
    target_compile_options(libmedia_codeclist PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmedia_codeclist
            android.hardware.media.omx-1.0
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libstagefright_foundation
            libutils
    )
endif ()
#################### libmedia_codeclist ####################

#################### libmedia_omx ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmedia/libmedia_omx/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/IOMXBufferSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IOMX.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/MediaCodecBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/OMXBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/omx/1.0/WOmxBufferSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/omx/1.0/WOmxNode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/omx/1.0/WOmxObserver.cpp
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
    add_library(libmedia_omx SHARED ${SOURCE_FILES})

    target_include_directories(libmedia_omx
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl
            ${ANDROID_ROOT}/frameworks/av/media/libmedia/include
            ${ANDROID_ROOT}/frameworks/av/media/libmedia
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmedia/libmedia_omx/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libmedia_omx PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmedia_omx
            android.hardware.media.omx-1.0
            android.hidl.token-1.0-utils
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libstagefright_foundation
            libstagefright_headers
            libui
            libutils
            media_plugin_headers
    )
endif ()
#################### libmedia_omx ####################

#################### libmedia ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmedia/libmedia/android_${OUT_ARCH_CPU}_static/gen/aidl/android/IMediaExtractorService.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IDataSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/BufferingSettings.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/mediaplayer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaHTTPConnection.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaHTTPService.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaExtractor.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaPlayerService.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaPlayerClient.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaRecorderClient.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaPlayer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaRecorder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IRemoteDisplay.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IRemoteDisplayClient.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IStreamSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/Metadata.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/mediarecorder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaMetadataRetriever.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/mediametadataretriever.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/MediaScanner.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/MediaScannerClient.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/CharacterEncodingDetector.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/IMediaDeathNotifier.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/MediaProfiles.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/MediaResource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/MediaResourcePolicy.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/StringArray.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/NdkMediaFormatPriv.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/NdkMediaErrorPriv.cpp
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
    add_library(libmedia SHARED ${SOURCE_FILES})

    target_include_directories(libmedia
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmedia/include
            ${ANDROID_ROOT}/frameworks/av/media/libmedia
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmedia/libmedia/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${ANDROID_ROOT}/frameworks/av/media/libmediahelper/include
            ${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include
            ${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmedia/libmedia_omx/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libmedia PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmedia
            android.hidl.token-1.0-utils
            android.media.audio.common.types-V2-cpp
            audioclient-types-aidl-cpp
            av-types-aidl-cpp
            framework-permission-aidl-cpp
            jni_headers
            libandroidicu
            libaudioclient
            libavmediaextentions
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libcamera_client
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libexpat
            libgui
            liblog
            libm
            libmedia_codeclist
            libmedia_omx
            libprocessgroup
            libstagefright_foundation
            libstagefright_headers
            libutils
            media_ndk_headers
            resourcemanager_aidl_interface-ndk
    )
endif ()
#################### libmedia ####################

#################### libmedia_omx_client ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/omx/1.0/WOmx.cpp
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
    add_library(libmedia_omx_client SHARED ${SOURCE_FILES})

    target_include_directories(libmedia_omx_client
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmedia/include
            ${ANDROID_ROOT}/frameworks/av/media/libmedia
            ${ANDROID_ROOT}/frameworks/av/media/libmedia/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmedia/libmedia_omx/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libmedia_omx_client PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmedia_omx_client
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libmedia_omx
            libstagefright_foundation
            libstagefright_headers
            libui
            libutils
            media_plugin_headers
    )
endif ()
#################### libmedia_omx_client ####################

#################### libmedia_ndkformatpriv ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/NdkMediaFormatPriv.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmedia/NdkMediaErrorPriv.cpp
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
    add_library(libmedia_ndkformatpriv SHARED ${SOURCE_FILES})

    target_include_directories(libmedia_ndkformatpriv
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmedia/include
            ${ANDROID_ROOT}/frameworks/av/media/libmedia
    )

    # Add c/c++ flags
    target_compile_options(libmedia_ndkformatpriv PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmedia_ndkformatpriv
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libstagefright_foundation_headers
            libstagefright_headers
            media_ndk_headers
    )
endif ()
#################### libmedia_ndkformatpriv ####################

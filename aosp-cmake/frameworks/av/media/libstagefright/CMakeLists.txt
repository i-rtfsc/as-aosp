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

project(libstagefright)

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

#################### libstagefright ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/AACWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/ACodec.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/ACodecBufferChannel.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/AHierarchicalStateMachine.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/AMRWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/AudioSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/BufferImpl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/CallbackDataSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/CallbackMediaSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/CameraSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/CameraSourceTimeLapse.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/CodecErrorLog.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/CryptoAsync.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/FrameDecoder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/HevcUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/InterfaceUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/JPEGSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MPEG2TSWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MPEG4Writer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaAdapter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaAppender.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaClock.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaCodec.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaCodecList.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaCodecListOverrides.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaCodecSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaExtractor.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaExtractorFactory.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaSync.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaTrack.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaMuxer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/NuMediaExtractor.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/OggWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/OMXClient.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/OmxInfoBuilder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/RemoteMediaExtractor.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/RemoteMediaSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/SimpleDecodingSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/StagefrightMediaScanner.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/SurfaceUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/ThrottledSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/Utils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/VideoRenderQualityTracker.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/VideoFrameSchedulerBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/VideoFrameScheduler.cpp
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
    add_library(libstagefright SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/xmlparser/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/webm/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/timedtext/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/include
    )

    # Add c/c++ flags
    target_compile_options(libstagefright PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright
            android.hardware.cas.native-1.0
            android.hardware.drm-1.0
            android.hardware.media.omx-1.0
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            framework-permission-aidl-cpp
            libaudioclient
            libaudioclient_aidl_conversion
            libaudioutils
            libavextensions
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libcamera_client
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcodec2
            libcodec2_vndk
            libcrypto
            libcutils
            libdatasource
            libdl
            libdl_android
            libgui
            libhidlallocatorutils
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia
            libmedia_codeclist
            libmedia_helper
            libmedia_ndkformatpriv
            libmedia_omx
            libmedia_omx_client
            libmediadrm_headers
            libmediaformatshaper_headers
            libmediametrics
            libmediandk_format
            libnativeloader-headers
            libogg
            libsfplugin_ccodec
            libsfplugin_ccodec_utils
            libstagefright_codecbase
            libstagefright_color_conversion
            libstagefright_esds
            libstagefright_foundation
            libstagefright_framecapture_utils
            libstagefright_id3
            libstagefright_mpeg2support_nocrypto
            libstagefright_omx_utils
            libstagefright_timedtext
            libstagefright_webm
            libstagefright_xmlparser_headers
            libui
            libutils
            libyuv_static
            media_ndk_headers
            packagemanager_aidl-cpp
            server_configurable_flags
    )
endif ()
#################### libstagefright ####################

#################### libstagefright_framecapture_utils ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/FrameCaptureLayer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/FrameCaptureProcessor.cpp
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
    add_library(libstagefright_framecapture_utils SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_framecapture_utils
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/include
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_framecapture_utils PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_framecapture_utils
            libEGL
            libGLESv1_CM
            libGLESv2
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgui
            liblog
            libm
            libprocessgroup
            librenderfright
            libstagefright_foundation
            libsync
            libui
            libutils
            libvulkan
    )
endif ()
#################### libstagefright_framecapture_utils ####################

#################### libstagefright_surface_utils ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/SurfaceUtils.cpp
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
    add_library(libstagefright_surface_utils SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_surface_utils
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_surface_utils PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_surface_utils
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgui
            liblog
            libm
            libui
            libutils
    )
endif ()
#################### libstagefright_surface_utils ####################

#################### libstagefright_codecbase ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/CodecBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/DataConverter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/FrameRenderTracker.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaCodecListWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/SkipCutBuffer.cpp
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
    add_library(libstagefright_codecbase SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_codecbase
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_codecbase PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_codecbase
            android.hardware.cas.native-1.0
            android.hardware.drm-1.0
            libaudioutils
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgui
            libhidlallocatorutils
            liblog
            libm
            libmedia_codeclist
            libmedia_omx
            libmediadrm_headers
            libstagefright_foundation
            libui
            libutils
            media_ndk_headers
    )
endif ()
#################### libstagefright_codecbase ####################

#################### libstagefright_mpeg2extractor ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/Utils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/MediaSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/HevcUtils.cpp
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
    add_library(libstagefright_mpeg2extractor SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_mpeg2extractor
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_mpeg2extractor PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_mpeg2extractor
            libaudioclient_headers
            libbase_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libmedia_datasource_headers
            libstagefright_esds
            media_ndk_headers
            media_plugin_headers
    )
endif ()
#################### libstagefright_mpeg2extractor ####################

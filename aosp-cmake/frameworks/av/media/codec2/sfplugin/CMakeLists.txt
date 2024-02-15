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

project(libsfplugin_ccodec)

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

#################### libsfplugin_ccodec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/C2OMXNode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/CCodec.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/CCodecBufferChannel.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/CCodecBuffers.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/CCodecConfig.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/Codec2Buffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/Codec2InfoBuilder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/FrameReassembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/PipelineWatcher.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/ReflectedParamUpdater.cpp
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
    add_library(libsfplugin_ccodec SHARED ${SOURCE_FILES})

    target_include_directories(libsfplugin_ccodec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/include
            ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin
            ${ANDROID_ROOT}/frameworks/av/media/codec2/sfplugin/utils
    )

    # Add c/c++ flags
    target_compile_options(libsfplugin_ccodec PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libsfplugin_ccodec
            android.hardware.cas.native-1.0
            android.hardware.drm-1.0
            android.hardware.graphics.mapper-4.0
            android.hardware.media.c2-1.0
            android.hardware.media.omx-1.0
            libSurfaceFlingerProperties
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcodec2
            libcodec2_client
            libcodec2_internal
            libcodec2_vndk
            libcutils
            libdl
            libgralloctypes
            libgui
            libhidlallocatorutils
            libhidlbase
            liblog
            libm
            libmedia_codeclist
            libmedia_omx
            libmediadrm_headers
            libmediametrics_headers
            libmediautils
            libsfplugin_ccodec_utils
            libstagefright_bufferqueue_helper
            libstagefright_codecbase
            libstagefright_foundation
            libstagefright_omx
            libstagefright_surface_utils
            libstagefright_xmlparser
            libui
            libutils
            media_ndk_headers
            server_configurable_flags
    )
endif ()
#################### libsfplugin_ccodec ####################

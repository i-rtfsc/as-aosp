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

project(libcodec2_vndk)

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

#################### libcodec2_vndk ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/C2AllocatorBlob.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/C2AllocatorIon.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/C2AllocatorGralloc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/C2Buffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/C2Config.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/C2DmaBufAllocator.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/C2Fence.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/C2PlatformStorePluginLoader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/C2Store.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/platform/C2BqBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/platform/C2SurfaceSyncObj.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/types.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/util/C2Debug.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/util/C2InterfaceHelper.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/util/C2InterfaceUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/util/C2ParamUtils.cpp
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
    add_library(libcodec2_vndk SHARED ${SOURCE_FILES})

    target_include_directories(libcodec2_vndk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/include
            ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk/internal
            ${ANDROID_ROOT}/frameworks/av/media/codec2/vndk
    )

    # Add c/c++ flags
    target_compile_options(libcodec2_vndk PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libcodec2_vndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.common-1.2
            android.hardware.media.bufferpool2-V1-ndk
            android.hardware.media.bufferpool-2.0
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcodec2_headers
            libcutils
            libdl
            libdmabufheap
            libfmq
            libgralloctypes
            libhidlbase
            libion
            liblog
            libm
            libnativewindow
            libstagefright_aidl_bufferpool2
            libstagefright_bufferpool-2.0.1
            libstagefright_foundation
            libui
            libutils
            media_plugin_headers
    )
endif ()
#################### libcodec2_vndk ####################

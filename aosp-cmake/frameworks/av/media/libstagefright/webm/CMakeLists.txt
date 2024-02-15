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

project(libstagefright_webm)

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

#################### libstagefright_webm ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/webm/EbmlUtil.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/webm/WebmElement.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/webm/WebmFrame.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/webm/WebmFrameThread.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/webm/WebmWriter.cpp
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
    add_library(libstagefright_webm SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_webm
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/webm/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/webm/include/webm
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/webm
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_webm PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_webm
            android.hardware.media.omx-1.0
            av-headers
            framework-permission-aidl-cpp
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
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_headers
            libmediadrm
            libnativewindow
            libstagefright_foundation
            libstagefright_mpeg2support_nocrypto
            libui
            libutils
    )
endif ()
#################### libstagefright_webm ####################

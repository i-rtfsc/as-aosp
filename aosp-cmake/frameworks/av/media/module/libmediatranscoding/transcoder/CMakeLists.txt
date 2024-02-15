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

project(libmediatranscoder)

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

#################### libmediatranscoder ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/MediaSampleQueue.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/MediaSampleReaderNDK.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/MediaSampleWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/MediaTrackTranscoder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/MediaTranscoder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/NdkCommon.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/PassthroughTrackTranscoder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/VideoTrackTranscoder.cpp
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
    add_library(libmediatranscoder SHARED ${SOURCE_FILES})

    target_include_directories(libmediatranscoder
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder/include
            ${ANDROID_ROOT}/frameworks/av/media/module/libmediatranscoding/transcoder
    )

    # Add c/c++ flags
    target_compile_options(libmediatranscoder PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libmediatranscoder
            libarect
            libbase
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libmediandk
            libnativewindow
            libutils
    )
endif ()
#################### libmediatranscoder ####################

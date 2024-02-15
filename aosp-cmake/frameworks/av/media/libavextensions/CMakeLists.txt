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

project(libavextensions)

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

#################### libavextensions ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libavextensions/stagefright/ExtendedMediaDefs.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libavextensions/stagefright/AVUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libavextensions/stagefright/AVFactory.cpp
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
    add_library(libavextensions SHARED ${SOURCE_FILES})

    target_include_directories(libavextensions
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libavextensions
    )

    # Add c/c++ flags
    target_compile_options(libavextensions PRIVATE
            -Werror
    )

    target_link_libraries(libavextensions
            android.hardware.media.omx-1.0
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libaudioclient
            libbase_headers
            libc
            libc++_static
            libc++demangle
            libcamera_client
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libgui
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_headers
            libmediadrm
            libmediametrics_headers
            libnativewindow
            libstagefright_headers
            libstagefright_mpeg2support_nocrypto
            libui
            libutils
            media_ndk_headers
            media_plugin_headers
    )
endif ()
#################### libavextensions ####################

#################### libavmediaserviceextensions ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libavextensions/mediaplayerservice/AVMediaServiceFactory.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libavextensions/mediaplayerservice/AVMediaServiceUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libavextensions/mediaplayerservice/AVNuFactory.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libavextensions/mediaplayerservice/AVNuUtils.cpp
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
    add_library(libavmediaserviceextensions SHARED ${SOURCE_FILES})

    target_include_directories(libavmediaserviceextensions
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libavextensions
    )

    # Add c/c++ flags
    target_compile_options(libavmediaserviceextensions PRIVATE
            -Werror
    )

    target_link_libraries(libavmediaserviceextensions
            android.hardware.media.omx-1.0
            libbase_headers
            libc
            libc++_static
            libc++demangle
            libcamera_client
            libclang_rt.builtins
            libcutils
            libdatasource
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libmediadrm
            libmediadrm_headers
            libmediametrics_headers
            libmediautils_headers
            libnativewindow
            libstagefright
            libstagefright_headers
            libstagefright_rtsp_headers
            libui
            libutils
            media_plugin_headers
    )
endif ()
#################### libavmediaserviceextensions ####################

#################### libavmediaextentions ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libavextensions/media/AVMediaUtils.cpp
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
    add_library(libavmediaextentions SHARED ${SOURCE_FILES})

    target_include_directories(libavmediaextentions
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libavextensions
    )

    # Add c/c++ flags
    target_compile_options(libavmediaextentions PRIVATE
            -Werror
    )

    target_link_libraries(libavmediaextentions
            libaudioclient_headers
            libbase_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libmedia_headers
            libmediametrics_headers
            libutils
            media_plugin_headers
    )
endif ()
#################### libavmediaextentions ####################

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

project(libmediandk)

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

#################### libmediandk ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkJavaVMHelper.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkMediaCodec.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkMediaCrypto.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkMediaDataSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkMediaExtractor.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkMediaFormat.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkMediaMuxer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkMediaDrm.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkImage.cpp
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkImageReader.cpp
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
    add_library(libmediandk SHARED ${SOURCE_FILES})

    target_include_directories(libmediandk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/include
            ${ANDROID_ROOT}/frameworks/av/media/ndk/include_platform
            ${ANDROID_ROOT}/frameworks/av/media/ndk
    )

    # Add c/c++ flags
    target_compile_options(libmediandk PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmediandk
            android.hardware.drm-V1-ndk
            android.hardware.graphics.bufferqueue-1.0
            android.hidl.token-1.0-utils
            jni_headers
            libandroid_runtime_lazy
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdatasource
            libdl
            libgrallocusage
            libgui
            libhidlbase
            liblog
            libm
            libmedia
            libmedia_jni_utils
            libmedia_omx
            libmediadrm
            libmediadrm_headers
            libmediametrics_headers
            libmediandk_utils
            libnativehelper_lazy
            libnativewindow
            libstagefright
            libstagefright_foundation
            libui
            libutils
    )
endif ()
#################### libmediandk ####################

#################### libmediandk_utils ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkMediaDataSourceCallbacks.cpp
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
    add_library(libmediandk_utils SHARED ${SOURCE_FILES})

    target_include_directories(libmediandk_utils
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/include
            ${ANDROID_ROOT}/frameworks/av/media/ndk
    )

    # Add c/c++ flags
    target_compile_options(libmediandk_utils PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmediandk_utils
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libmedia_headers
            libstagefright_headers
    )
endif ()
#################### libmediandk_utils ####################

#################### libmediandk_format ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/ndk/NdkMediaFormat.cpp
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
    add_library(libmediandk_format SHARED ${SOURCE_FILES})

    target_include_directories(libmediandk_format
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/ndk/include
            ${ANDROID_ROOT}/frameworks/av/media/ndk
    )

    # Add c/c++ flags
    target_compile_options(libmediandk_format PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libmediandk_format
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libstagefright_foundation_headers
    )
endif ()
#################### libmediandk_format ####################

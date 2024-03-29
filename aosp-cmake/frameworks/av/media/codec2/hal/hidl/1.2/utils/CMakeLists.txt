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

project(libcodec2_hidl@1.2)

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

#################### libcodec2_hidl-1.2 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/Component.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/ComponentInterface.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/ComponentStore.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/Configurable.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/InputBufferManager.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/InputSurface.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/InputSurfaceConnection.cpp
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/types.cpp
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
    add_library(libcodec2_hidl-1.2 SHARED ${SOURCE_FILES})

    target_include_directories(libcodec2_hidl-1.2
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/include
            ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils
    )

    # Add c/c++ flags
    target_compile_options(libcodec2_hidl-1.2 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libcodec2_hidl-1.2
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hardware.graphics.common-1.0
            android.hardware.media.bufferpool-2.0
            android.hardware.media.c2-1.0
            android.hardware.media.c2-1.1
            android.hardware.media.c2-1.2
            android.hardware.media.omx-1.0
            android.hardware.media-1.0
            libbase
            libbinder_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcodec2
            libcodec2_hidl-1.0
            libcodec2_hidl-1.1
            libcodec2_hidl_plugin_stub
            libcodec2_internal
            libcodec2_vndk
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libstagefright_bufferpool-2.0.1
            libstagefright_bufferqueue_helper_novndk
            libsystem_headers
            libui
            libutils
    )
endif ()
#################### libcodec2_hidl-1.2 ####################

#################### libcodec2_hidl_client-1.2 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/types.cpp
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
    add_library(libcodec2_hidl_client-1.2 SHARED ${SOURCE_FILES})

    target_include_directories(libcodec2_hidl_client-1.2
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils/include
            ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.2/utils
    )

    # Add c/c++ flags
    target_compile_options(libcodec2_hidl_client-1.2 PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libcodec2_hidl_client-1.2
            android.hardware.media.bufferpool-2.0
            android.hardware.media.c2-1.0
            android.hardware.media.c2-1.1
            android.hardware.media.c2-1.2
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcodec2
            libcodec2_hidl_client-1.0
            libcodec2_hidl_client-1.1
            libcodec2_internal
            libcodec2_vndk
            libcutils
            libdl
            libgui
            libhidlbase
            liblog
            libm
            libstagefright_bufferpool-2.0.1
            libui
            libutils
    )
endif ()
#################### libcodec2_hidl_client-1.2 ####################

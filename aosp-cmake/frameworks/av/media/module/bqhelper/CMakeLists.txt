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

project(libstagefright_bufferqueue_helper_novndk)

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

#################### libstagefright_bufferqueue_helper_novndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/bqhelper/libstagefright_bufferqueue_helper_novndk/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/gui/FrameEvent.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/bqhelper/FrameDropper.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/bqhelper/GraphicBufferSource.cpp
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
    add_library(libstagefright_bufferqueue_helper_novndk SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_bufferqueue_helper_novndk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/bqhelper/include
            ${ANDROID_ROOT}/frameworks/av/media/module/bqhelper
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/bqhelper/libstagefright_bufferqueue_helper_novndk/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_bufferqueue_helper_novndk PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libstagefright_bufferqueue_helper_novndk
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            android.hidl.token-1.0-utils
            libEGL
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgui_bufferqueue_static
            libhidlbase
            libhidlmemory
            liblog
            libm
            libnativewindow
            libstagefright_foundation
            libui
            libutils
            libvndksupport
            media_plugin_headers
    )
endif ()
#################### libstagefright_bufferqueue_helper_novndk ####################

#################### libstagefright_bufferqueue_helper ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/bqhelper/libstagefright_bufferqueue_helper/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/gui/FrameEvent.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/bqhelper/FrameDropper.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/bqhelper/GraphicBufferSource.cpp
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
    add_library(libstagefright_bufferqueue_helper SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_bufferqueue_helper
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/bqhelper/include
            ${ANDROID_ROOT}/frameworks/av/media/module/bqhelper
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/module/bqhelper/libstagefright_bufferqueue_helper/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_bufferqueue_helper PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libstagefright_bufferqueue_helper
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.graphics.bufferqueue-2.0
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgui
            libhidlbase
            libhidlmemory
            liblog
            libm
            libstagefright_foundation
            libui
            libutils
            media_plugin_headers
    )
endif ()
#################### libstagefright_bufferqueue_helper ####################

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

project(libui-types)

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

#################### libui-types ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/ui/ColorSpace.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/Rect.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/Region.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/Transform.cpp
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
    add_library(libui-types SHARED ${SOURCE_FILES})

    target_include_directories(libui-types
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/ui/include
            ${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock
            ${ANDROID_ROOT}/frameworks/native/libs/ui/include_private
            ${ANDROID_ROOT}/frameworks/native/libs/ui/include_types
            ${ANDROID_ROOT}/frameworks/native/libs/ui
    )

    # Add c/c++ flags
    target_compile_options(libui-types PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libui-types
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libmath
            libutils
            libwinpthread
    )
endif ()
#################### libui-types ####################

#################### libui ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/ui/DebugUtils.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/DeviceProductInfo.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/DisplayIdentification.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/DynamicDisplayInfo.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/Fence.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/FenceTime.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/FrameStats.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/Gralloc.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/Gralloc2.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/Gralloc3.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/Gralloc4.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/Gralloc5.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/GraphicBuffer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/GraphicBufferAllocator.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/GraphicBufferMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/PixelFormat.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/PublicFormat.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/ui/StaticAsserts.cpp
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
    add_library(libui SHARED ${SOURCE_FILES})

    target_include_directories(libui
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/ui/include
            ${ANDROID_ROOT}/frameworks/native/libs/ui/include_private
            ${ANDROID_ROOT}/frameworks/native/libs/ui
            ${ANDROID_ROOT}/frameworks/native/libs/ui/include_mock
            ${ANDROID_ROOT}/frameworks/native/libs/ui/include_types
    )

    # Add c/c++ flags
    target_compile_options(libui PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(libui
            android.hardware.graphics.allocator-V2-ndk
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.allocator-4.0
            android.hardware.graphics.common-V4-ndk
            android.hardware.graphics.common-1.2
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-2.1
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
            libaidlcommonsupport
            libarect
            libbase
            libbase_headers
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgralloctypes
            libgrallocusage
            libhardware_headers
            libhidlbase
            libimapper_providerutils
            libimapper_stablec
            liblog
            libm
            libmath
            libnativebase_headers
            libnativewindow_headers
            libsync
            libui-types
            libui_headers
            libutils
            libvndksupport
    )
endif ()
#################### libui ####################

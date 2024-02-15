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

project(android.hardware.graphics.composer@2.2-vts)

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

#################### android.hardware.graphics.composer-2.2-vts ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/graphics/composer/2.2/utils/vts/ComposerVts.cpp
        ${ANDROID_ROOT}/hardware/interfaces/graphics/composer/2.2/utils/vts/ReadbackVts.cpp
        ${ANDROID_ROOT}/hardware/interfaces/graphics/composer/2.2/utils/vts/RenderEngineVts.cpp
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
    add_library(android.hardware.graphics.composer-2.2-vts SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.composer-2.2-vts
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/graphics/composer/2.2/utils/vts/include
            ${ANDROID_ROOT}/hardware/interfaces/graphics/composer/2.2/utils/vts
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.composer-2.2-vts PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.graphics.composer-2.2-vts
            android.hardware.graphics.allocator-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.graphics.composer3-V2-ndk
            android.hardware.graphics.composer-2.1-vts
            android.hardware.graphics.composer-2.2
            android.hardware.graphics.composer-2.2-command-buffer
            android.hardware.graphics.mapper-2.1-vts
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-3.0-vts
            android.hardware.graphics.mapper-4.0
            android.hardware.graphics.mapper-4.0-vts
            libEGL
            libGLESv2
            libarect
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest
            liblog
            libm
            libmath
            libnativewindow
            libperfetto_client_experimental
            libpng
            librenderengine
            libshaders
            libskia_renderengine
            libtiff
            libtonemap
            libui
            libvulkan
            libz
    )
endif ()
#################### android.hardware.graphics.composer-2.2-vts ####################

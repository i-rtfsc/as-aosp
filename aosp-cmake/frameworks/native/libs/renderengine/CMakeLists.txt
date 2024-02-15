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

project(librenderengine_mocks)

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

#################### librenderengine_mocks ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/mock/Framebuffer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/mock/Image.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/mock/RenderEngine.cpp
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
    add_library(librenderengine_mocks SHARED ${SOURCE_FILES})

    target_include_directories(librenderengine_mocks
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/renderengine/include
            ${ANDROID_ROOT}/frameworks/native/libs/renderengine
    )

    # Add c/c++ flags
    target_compile_options(librenderengine_mocks PRIVATE
            -Wall
            -Werror
            -DFLYME_GRAPHICS_EXTEND
    )

    target_link_libraries(librenderengine_mocks
            android.hardware.graphics.composer3-V2-ndk
            libEGL
            libGLESv1_CM
            libGLESv2
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgmock
            libgtest
            libgui
            liblog
            libm
            libnativewindow
            libprocessgroup
            libruntime_log
            libshaders
            libsync
            libtonemap
            libui
            libutils
            libvulkan
    )
endif ()
#################### librenderengine_mocks ####################

#################### librenderengine ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/Description.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/ExternalTexture.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/Mesh.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/RenderEngine.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/Texture.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/GLESRenderEngine.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/GLExtensions.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/GLFramebuffer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/GLImage.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/GLShadowTexture.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/GLShadowVertexGenerator.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/GLSkiaShadowPort.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/GLVertexBuffer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/ImageManager.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/Program.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/ProgramCache.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/filters/BlurFilter.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/gl/filters/GenericProgram.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/threaded/RenderEngineThreaded.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/AutoBackendTexture.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/Cache.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/ColorSpaces.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/SmoothCornerPathGenerator.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/SkiaRenderEngine.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/SkiaGLRenderEngine.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/SkiaVkRenderEngine.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/debug/CaptureTimer.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/debug/CommonPool.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/debug/SkiaCapture.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/debug/SkiaMemoryReporter.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/filters/BlurFilter.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/filters/GaussianBlurFilter.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/filters/KawaseBlurFilter.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/filters/LinearEffect.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/renderengine/skia/filters/StretchShaderFactory.cpp
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
    add_library(librenderengine SHARED ${SOURCE_FILES})

    target_include_directories(librenderengine
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/renderengine/include
            ${ANDROID_ROOT}/frameworks/native/libs/renderengine
    )

    # Add c/c++ flags
    target_compile_options(librenderengine PRIVATE
            -Wall
            -Werror
            -DFLYME_GRAPHICS_EXTEND
            -Werror=format
            -Wno-unused-parameter
    )

    target_link_libraries(librenderengine
            android.hardware.graphics.composer3-V2-ndk
            libEGL
            libGLESv1_CM
            libGLESv2
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgui
            liblog
            libm
            libnativewindow
            libperfetto_client_experimental
            libpng
            libprocessgroup
            libruntime_log
            libshaders
            libskia_renderengine
            libsync
            libtiff
            libtonemap
            libtonemap_headers
            libui
            libutils
            libvulkan
            libz
    )
endif ()
#################### librenderengine ####################

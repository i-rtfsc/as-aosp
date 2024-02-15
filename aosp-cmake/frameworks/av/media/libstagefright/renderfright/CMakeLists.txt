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

project(librenderfright)

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

#################### librenderfright ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/Description.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/Mesh.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/RenderEngine.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/Texture.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/GLESRenderEngine.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/GLExtensions.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/GLFramebuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/GLImage.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/GLShadowTexture.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/GLShadowVertexGenerator.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/GLSkiaShadowPort.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/GLVertexBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/ImageManager.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/Program.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/ProgramCache.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/filters/BlurFilter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/gl/filters/GenericProgram.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/threaded/RenderEngineThreaded.cpp
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
    add_library(librenderfright SHARED ${SOURCE_FILES})

    target_include_directories(librenderfright
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright
    )

    # Add c/c++ flags
    target_compile_options(librenderfright PRIVATE
            -Wall
            -Werror
            -Werror=format
    )

    target_link_libraries(librenderfright
            libEGL
            libGLESv1_CM
            libGLESv2
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsync
            libui
            libutils
            libvulkan
    )
endif ()
#################### librenderfright ####################

#################### librenderfright_mocks ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/mock/Framebuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/mock/Image.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/mock/RenderEngine.cpp
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
    add_library(librenderfright_mocks SHARED ${SOURCE_FILES})

    target_include_directories(librenderfright_mocks
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/renderfright
    )

    # Add c/c++ flags
    target_compile_options(librenderfright_mocks PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(librenderfright_mocks
            libEGL
            libGLESv1_CM
            libGLESv2
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgmock
            libgtest
            liblog
            libm
            libnativewindow
            libprocessgroup
            libsync
            libui
            libutils
            libvulkan
    )
endif ()
#################### librenderfright_mocks ####################

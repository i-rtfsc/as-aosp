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

project(gfxstream_gl_server)

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

#################### gfxstream_gl_server ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/BufferGl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/ColorBufferGl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/CompositorGl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/DebugGl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/DisplayGl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/DisplaySurfaceGl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/EmulatedEglConfig.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/EmulatedEglContext.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/EmulatedEglFenceSync.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/EmulatedEglImage.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/EmulatedEglWindowSurface.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/EmulationGl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/GLESVersionDetector.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/ReadbackWorkerGl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/TextureDraw.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/TextureResize.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/YUVConverter.cpp
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
    add_library(gfxstream_gl_server SHARED ${SOURCE_FILES})

    target_include_directories(gfxstream_gl_server
            PUBLIC
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/include
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/gles2_dec
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glsnapshot
    )

    # Add c/c++ flags
    target_compile_options(gfxstream_gl_server PRIVATE
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-variable
            -Wno-unused-private-field
    )

    target_link_libraries(gfxstream_gl_server
            gfxstream_gles1_dec
            gfxstream_gles2_dec
            gfxstream_glsnapshot
            gfxstream_headers
            gfxstream_x11_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnativewindow
    )
endif ()
#################### gfxstream_gl_server ####################

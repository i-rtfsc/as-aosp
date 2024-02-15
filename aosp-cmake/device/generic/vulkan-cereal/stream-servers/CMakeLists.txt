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

project(libgfxstream_backend)

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

#################### libgfxstream_backend ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/Buffer.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/ChannelStream.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/ColorBuffer.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/DisplaySurface.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/DisplaySurfaceUser.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/Hwc2.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/PostWorker.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/ReadBuffer.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/render_api.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RenderChannelImpl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RenderThread.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RenderThreadInfo.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RenderThreadInfoGl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RenderThreadInfoMagma.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RingStream.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/SyncThread.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RenderControl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RenderWindow.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RenderLibImpl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/RendererImpl.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/FrameBuffer.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/GfxStreamAgents.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/virtio-gpu-gfxstream-renderer.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/VirtioGpuTimelines.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/VsyncThread.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/NativeSubWindow_android.cpp
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
    add_library(libgfxstream_backend SHARED ${SOURCE_FILES})

    target_include_directories(libgfxstream_backend
            PUBLIC
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/include
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/apigen-codec-common
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/gles2_dec
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glsnapshot
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/magma/magma_dec
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/magma
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/cereal
    )

    # Add c/c++ flags
    target_compile_options(libgfxstream_backend PRIVATE
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-variable
            -Wno-unused-private-field
            -Wno-unused-parameter
            -Wno-unused-variable
            -Wno-unused-function
    )

    target_link_libraries(libgfxstream_backend
            gfxstream_apigen_codec_common
            gfxstream_base
            gfxstream_compressedTextures
            gfxstream_dispatch
            gfxstream_emulated_textures
            gfxstream_gl_host_common
            gfxstream_gl_server
            gfxstream_gles1_dec
            gfxstream_gles2_dec
            gfxstream_glm
            gfxstream_glsnapshot
            gfxstream_headers
            gfxstream_host_common
            gfxstream_magma_dec
            gfxstream_magma_server
            gfxstream_renderControl_dec
            gfxstream_translator_egl
            gfxstream_translator_glcommon
            gfxstream_translator_glescm
            gfxstream_translator_glesv2
            gfxstream_vulkan_cereal_host
            gfxstream_vulkan_server
            gfxstream_x11_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libnativewindow
    )
endif ()
#################### libgfxstream_backend ####################

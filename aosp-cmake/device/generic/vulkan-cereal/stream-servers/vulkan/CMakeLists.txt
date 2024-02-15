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

project(gfxstream_vulkan_server)

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

#################### gfxstream_vulkan_server ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/BorrowedImageVk.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/BufferVk.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/ColorBufferVk.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/CompositorVk.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/DebugUtilsHelper.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/DisplayVk.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/DisplaySurfaceVk.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/SwapChainStateVk.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VkAndroidNativeBuffer.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VkCommonOperations.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VkDecoder.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VkDecoderGlobalState.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VkDecoderSnapshot.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VkFormatUtils.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VkReconstruction.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VulkanDispatch.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VulkanHandleMapping.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/VulkanStream.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/vk_util.cpp
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
    add_library(gfxstream_vulkan_server SHARED ${SOURCE_FILES})

    target_include_directories(gfxstream_vulkan_server
            PUBLIC
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/vulkan/cereal
    )

    # Add c/c++ flags
    target_compile_options(gfxstream_vulkan_server PRIVATE
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-variable
            -Wno-unused-private-field
            -Wno-unused-value
            -Wno-unused-parameter
            -Wno-unused-variable
            -Wno-unused-function
    )

    target_link_libraries(gfxstream_vulkan_server
            gfxstream_apigen_codec_common
            gfxstream_base
            gfxstream_compressedTextures
            gfxstream_emulated_textures
            gfxstream_glm
            gfxstream_headers
            gfxstream_host_common
            gfxstream_vulkan_cereal_host
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
#################### gfxstream_vulkan_server ####################

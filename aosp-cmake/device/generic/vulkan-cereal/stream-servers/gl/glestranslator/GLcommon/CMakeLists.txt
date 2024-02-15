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

project(gfxstream_translator_glcommon)

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

#################### gfxstream_translator_glcommon ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/etc.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/rgtc.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/FramebufferData.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/GLBackgroundLoader.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/GLDispatch.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/GLESbuffer.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/GLEScontext.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/GLESpointer.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/GLESvalidate.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/GLutils.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/NamedObject.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/ObjectData.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/ObjectNameSpace.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/PaletteTexture.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/RangeManip.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/SaveableTexture.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/ScopedGLState.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/ShareGroup.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/TextureData.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon/TextureUtils.cpp
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
    add_library(gfxstream_translator_glcommon SHARED ${SOURCE_FILES})

    target_include_directories(gfxstream_translator_glcommon
            PUBLIC
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/glestranslator/GLcommon
    )

    # Add c/c++ flags
    target_compile_options(gfxstream_translator_glcommon PRIVATE
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-variable
            -Wno-unused-private-field
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(gfxstream_translator_glcommon
            gfxstream_base
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
#################### gfxstream_translator_glcommon ####################

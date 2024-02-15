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

project(gfxstream_dispatch)

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

#################### gfxstream_dispatch ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/OpenGLESDispatch/EGLDispatch.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/OpenGLESDispatch/gles1_stubbed_in_translator_namespace.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/OpenGLESDispatch/gles2_stubbed_in_translator_namespace.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/OpenGLESDispatch/GLESv1Dispatch.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/OpenGLESDispatch/GLESv2Dispatch.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/OpenGLESDispatch/OpenGLDispatchLoader.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/OpenGLESDispatch/StaticDispatch.cpp
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
    add_library(gfxstream_dispatch SHARED ${SOURCE_FILES})

    target_include_directories(gfxstream_dispatch
            PUBLIC
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/stream-servers/gl/OpenGLESDispatch
    )

    # Add c/c++ flags
    target_compile_options(gfxstream_dispatch PRIVATE
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-variable
            -Wno-unused-private-field
    )

    target_link_libraries(gfxstream_dispatch
            gfxstream_apigen_codec_common
            gfxstream_base
            gfxstream_gles1_dec
            gfxstream_gles2_dec
            gfxstream_glm
            gfxstream_headers
            gfxstream_translator_egl
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
#################### gfxstream_dispatch ####################

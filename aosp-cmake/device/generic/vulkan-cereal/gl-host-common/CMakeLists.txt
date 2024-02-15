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

project(gfxstream_gl_host_common)

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

#################### gfxstream_gl_host_common ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengles.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengl/EmuglBackendList.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengl/emugl_config.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengl/GLProcessPipe.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengl/GpuFrameBridge.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengl/gpuinfo.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengl/logger.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengl/misc.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengl/OpenglEsPipe.cpp
        ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/opengl/NativeGpuInfo_linux.cpp
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
    add_library(gfxstream_gl_host_common SHARED ${SOURCE_FILES})

    target_include_directories(gfxstream_gl_host_common
            PUBLIC
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/include/host-common
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common
            ${ANDROID_ROOT}/device/generic/vulkan-cereal/gl-host-common/include
    )

    # Add c/c++ flags
    target_compile_options(gfxstream_gl_host_common PRIVATE
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-variable
            -Wno-unused-private-field
    )

    target_link_libraries(gfxstream_gl_host_common
            gfxstream_base
            gfxstream_headers
            gfxstream_host_common
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
#################### gfxstream_gl_host_common ####################

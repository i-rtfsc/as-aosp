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

project(gfxstream_base)

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

#################### gfxstream_base ####################
set(FILES
        ${ANDROID_ROOT}/hardware/google/aemu/base/AlignedBuf.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/CompressingStream.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/CpuTime.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/DecompressingStream.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/FileUtils.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/FunctorThread.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/GLObjectCounter.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/HealthMonitor.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/LayoutResolver.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/MemStream.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/StdioStream.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/MemoryTracker.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/MessageChannel.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/Metrics.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/PathUtils.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/ring_buffer.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/SharedLibrary.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/SharedMemory_posix.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/StringFormat.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/Stream.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/StreamSerializing.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/SubAllocator.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/System.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/Tracing.cpp
        ${ANDROID_ROOT}/hardware/google/aemu/base/Thread_pthread.cpp
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
    add_library(gfxstream_base SHARED ${SOURCE_FILES})

    target_include_directories(gfxstream_base
            PUBLIC
            ${ANDROID_ROOT}/hardware/google/aemu/base
            ${ANDROID_ROOT}/hardware/google/aemu/base/include
    )

    # Add c/c++ flags
    target_compile_options(gfxstream_base PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-variable
            -Wno-unused-private-field
    )

    target_link_libraries(gfxstream_base
            gfxstream_headers
            gfxstream_renderdoc_headers
            gfxstream_x11_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblz4
            libm
            libnativewindow
    )
endif ()
#################### gfxstream_base ####################

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

project(libnbaio_mono)

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

#################### libnbaio_mono ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/MonoPipe.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/MonoPipeReader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/NBAIO.cpp
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
    add_library(libnbaio_mono SHARED ${SOURCE_FILES})

    target_include_directories(libnbaio_mono
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libnbaio/include_mono
            ${ANDROID_ROOT}/frameworks/av/media/libnbaio
    )

    # Add c/c++ flags
    target_compile_options(libnbaio_mono PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libnbaio_mono
            libaudio_system_headers
            libaudioclient_headers
            libaudioutils
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libnbaio_mono ####################

#################### libnbaio ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/MonoPipe.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/MonoPipeReader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/NBAIO.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/AudioBufferProviderSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/AudioStreamInSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/AudioStreamOutSink.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/Pipe.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/PipeReader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libnbaio/SourceAudioBufferProvider.cpp
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
    add_library(libnbaio SHARED ${SOURCE_FILES})

    target_include_directories(libnbaio
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libnbaio/include_mono
            ${ANDROID_ROOT}/frameworks/av/media/libnbaio/include
            ${ANDROID_ROOT}/frameworks/av/media/libnbaio
    )

    # Add c/c++ flags
    target_compile_options(libnbaio PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libnbaio
            libaudio_system_headers
            libaudioclient_headers
            libaudiohal_headers
            libaudioutils
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libmediautils
            libutils
    )
endif ()
#################### libnbaio ####################

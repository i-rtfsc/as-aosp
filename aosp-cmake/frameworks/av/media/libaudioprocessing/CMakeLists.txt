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

project(libaudioprocessing_base)

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

#################### libaudioprocessing_base ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/AudioMixerBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/AudioResampler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/AudioResamplerCubic.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/AudioResamplerSinc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/AudioResamplerDyn.cpp
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
    add_library(libaudioprocessing_base SHARED ${SOURCE_FILES})

    target_include_directories(libaudioprocessing_base
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/include
            ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing
    )

    # Add c/c++ flags
    target_compile_options(libaudioprocessing_base PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudioprocessing_base
            libaudioclient_headers
            libaudioutils
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libaudioprocessing_base ####################

#################### libaudioprocessing ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/AudioMixer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/BufferProviders.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/RecordBufferConverter.cpp
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
    add_library(libaudioprocessing SHARED ${SOURCE_FILES})

    target_include_directories(libaudioprocessing
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing/include
            ${ANDROID_ROOT}/frameworks/av/media/libaudioprocessing
    )

    # Add c/c++ flags
    target_compile_options(libaudioprocessing PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudioprocessing
            libaudioclient_headers
            libaudiohal
            libaudiohal_headers
            libaudioprocessing_base
            libaudioutils
            libbase_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libmedia_headers
            libsonic
            libutils
            libvibrator
    )
endif ()
#################### libaudioprocessing ####################

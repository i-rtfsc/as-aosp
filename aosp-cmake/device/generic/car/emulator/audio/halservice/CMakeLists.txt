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

project(android.hardware.audio.service-caremu)

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

#################### audiocontrol-caremu ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/car/emulator/audio/halservice/AudioControl.cpp
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
    add_library(audiocontrol-caremu SHARED ${SOURCE_FILES})

    target_include_directories(audiocontrol-caremu
            PUBLIC
            ${ANDROID_ROOT}/device/generic/car/emulator/audio/halservice
    )

    target_link_libraries(audiocontrol-caremu
            android.hardware.audio.common-7.0-enums
            android.hardware.automotive.audiocontrol-V1-ndk
            audio.primary.caremu
            libbase
            libbinder_ndk
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### audiocontrol-caremu ####################

#################### android.hardware.audio.service-caremu ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/car/emulator/audio/halservice/service.cpp
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
    add_executable(android.hardware.audio.service-caremu ${SOURCE_FILES})

    target_include_directories(android.hardware.audio.service-caremu
            PUBLIC
            ${ANDROID_ROOT}/device/generic/car/emulator/audio/halservice
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.audio.service-caremu PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(android.hardware.audio.service-caremu
            android.hardware.audio.common-6.0
            android.hardware.audio.effect-6.0
            android.hardware.audio-6.0
            android.hardware.automotive.audiocontrol-V1-ndk
            audiocontrol-caremu
            libbase
            libbinder
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libhardware
            libhardware_headers
            libhidlbase
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.audio.service-caremu ####################

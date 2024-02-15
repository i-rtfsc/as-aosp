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

project(libvibratorservice)

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

#################### libvibratorservice ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/vibratorservice/VibratorCallbackScheduler.cpp
        ${ANDROID_ROOT}/frameworks/native/services/vibratorservice/VibratorHalController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/vibratorservice/VibratorHalWrapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/vibratorservice/VibratorManagerHalController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/vibratorservice/VibratorManagerHalWrapper.cpp
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
    add_library(libvibratorservice SHARED ${SOURCE_FILES})

    target_include_directories(libvibratorservice
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/vibratorservice/include
            ${ANDROID_ROOT}/frameworks/native/services/vibratorservice
    )

    # Add c/c++ flags
    target_compile_options(libvibratorservice PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libvibratorservice
            android.hardware.vibrator-V2-cpp
            android.hardware.vibrator-1.0
            android.hardware.vibrator-1.1
            android.hardware.vibrator-1.2
            android.hardware.vibrator-1.3
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### libvibratorservice ####################

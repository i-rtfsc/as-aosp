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

project(libinputflinger_base)

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

#################### libinputflinger_base ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/InputListener.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/InputReaderBase.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/InputThread.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/NotifyArgs.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/VibrationElement.cpp
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
    add_library(libinputflinger_base SHARED ${SOURCE_FILES})

    target_include_directories(libinputflinger_base
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/include
    )

    # Add c/c++ flags
    target_compile_options(libinputflinger_base PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libinputflinger_base
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libinput
            libinputflinger_headers
            liblog
            libm
            libui-types
            libutils
    )
endif ()
#################### libinputflinger_base ####################

#################### libinputflinger ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/InputCommonConverter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/InputProcessor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/PreferStylusOverTouchBlocker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/UnwantedInteractionBlocker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/InputManager.cpp
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
    add_library(libinputflinger SHARED ${SOURCE_FILES})

    target_include_directories(libinputflinger
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/include
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/dispatcher/include
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/include
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reporter
    )

    # Add c/c++ flags
    target_compile_options(libinputflinger PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libinputflinger
            android.hardware.input.processor-V1-ndk
            libattestation
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libgui
            libhidlbase
            libinput
            libinputdispatcher
            libinputflinger_base
            libinputreader
            libinputreporter
            libkll
            liblog
            libm
            libpalmrejection
            libprotobuf-cpp-lite
            libstatslog
            libstatspull
            libstatssocket
            libui-types
            libutils
            server_configurable_flags
    )
endif ()
#################### libinputflinger ####################

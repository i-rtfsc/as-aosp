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

project(libinputreader)

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

#################### libinputreader ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/EventHub.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/InputDevice.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/InputReader.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/Macros.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/TouchVideoDevice.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller/PeripheralController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/CapturedTouchpadEventConverter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/CursorInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/ExternalStylusInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/InputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/JoystickInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/KeyboardInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/MultiTouchInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/RotaryEncoderInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/SensorInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/SingleTouchInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/SwitchInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/TouchCursorInputMapperCommon.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/TouchInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/TouchpadInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/VibratorInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/CursorButtonAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/CursorScrollAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/HidUsageAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/MultiTouchMotionAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/SingleTouchMotionAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/TouchButtonAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/GestureConverter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/GesturesLogging.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/HardwareStateConverter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/PropertyProvider.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/InputReaderFactory.cpp
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
    add_library(libinputreader SHARED ${SOURCE_FILES})

    target_include_directories(libinputreader
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/include
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures
    )

    # Add c/c++ flags
    target_compile_options(libinputreader PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libinputreader
            libPlatformProperties
            libbase
            libbatteryservice_headers
            libbinder
            libc
            libc++
            libc++demangle
            libc++fs
            libcap
            libchrome-gestures
            libchrome-gestures_headers
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libinput
            libinputflinger_base
            libinputreader_headers
            libjsoncpp
            liblog
            libm
            libstatslog
            libui-types
            libutils
    )
endif ()
#################### libinputreader ####################

#################### libinputreader_static ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/EventHub.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/InputDevice.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/InputReader.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/Macros.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/TouchVideoDevice.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller/PeripheralController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/CapturedTouchpadEventConverter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/CursorInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/ExternalStylusInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/InputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/JoystickInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/KeyboardInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/MultiTouchInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/RotaryEncoderInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/SensorInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/SingleTouchInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/SwitchInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/TouchCursorInputMapperCommon.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/TouchInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/TouchpadInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/VibratorInputMapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/CursorButtonAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/CursorScrollAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/HidUsageAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/MultiTouchMotionAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/SingleTouchMotionAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator/TouchButtonAccumulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/GestureConverter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/GesturesLogging.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/HardwareStateConverter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures/PropertyProvider.cpp
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
    add_library(libinputreader_static SHARED ${SOURCE_FILES})

    target_include_directories(libinputreader_static
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/controller
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/include
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/accumulator
            ${ANDROID_ROOT}/frameworks/native/services/inputflinger/reader/mapper/gestures
    )

    # Add c/c++ flags
    target_compile_options(libinputreader_static PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libinputreader_static
            libPlatformProperties
            libbase
            libbatteryservice_headers
            libc
            libc++_static
            libc++demangle
            libc++fs
            libcap
            libchrome-gestures
            libchrome-gestures_headers
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libinput
            libinputflinger_base
            libinputreader_headers
            libjsoncpp
            liblog
            libm
            libstatslog
            libui-types
            libutils
    )
endif ()
#################### libinputreader_static ####################

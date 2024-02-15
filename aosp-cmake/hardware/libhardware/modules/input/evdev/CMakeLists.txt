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

project(libinput_evdev)

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

#################### libinput_evdev ####################
set(FILES
        ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev/BitUtils.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev/InputHub.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev/InputDevice.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev/InputDeviceManager.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev/InputHost.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev/InputMapper.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev/MouseInputMapper.cpp
        ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev/SwitchInputMapper.cpp
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
    add_library(libinput_evdev SHARED ${SOURCE_FILES})

    target_include_directories(libinput_evdev
            PUBLIC
            ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev
    )

    # Add c/c++ flags
    target_compile_options(libinput_evdev PRIVATE
            -Werror
    )

    target_link_libraries(libinput_evdev
            jni_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libhardware_legacy
            liblog
            libm
            libutils
    )
endif ()
#################### libinput_evdev ####################

#################### input.evdev.default ####################
set(FILES
        ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev/EvdevModule.cpp
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
    add_library(input.evdev.default SHARED ${SOURCE_FILES})

    target_include_directories(input.evdev.default
            PUBLIC
            ${ANDROID_ROOT}/hardware/libhardware/modules/input/evdev
    )

    target_link_libraries(input.evdev.default
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libinput_evdev
            liblog
            libm
    )
endif ()
#################### input.evdev.default ####################

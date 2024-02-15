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

project(virtual_keyboard)

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

#################### virtual_keyboard ####################
set(FILES
        ${ANDROID_ROOT}/vendor/microsoft/proprietary/windowsconnectivity/crossdevice/virtualkeyboard/main.cpp
        ${ANDROID_ROOT}/vendor/microsoft/proprietary/windowsconnectivity/crossdevice/virtualkeyboard/VirtualKeyboardService.cpp
        ${ANDROID_ROOT}/vendor/microsoft/proprietary/windowsconnectivity/crossdevice/virtualkeyboard/EvdevInjector.cpp
        ${ANDROID_ROOT}/vendor/microsoft/proprietary/windowsconnectivity/crossdevice/virtualkeyboard/VirtualKeyboardEvdev.cpp
        ${OUT_INTERMEDIATES_ROOT}/vendor/microsoft/proprietary/windowsconnectivity/crossdevice/virtualkeyboard/virtual_keyboard/android_${OUT_ARCH_CPU}/gen/aidl/com/microsoftsdk/IVirtualKeyboardService.cpp
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
    add_executable(virtual_keyboard ${SOURCE_FILES})

    target_include_directories(virtual_keyboard
            PUBLIC
            ${ANDROID_ROOT}/vendor/microsoft/proprietary/windowsconnectivity/crossdevice/virtualkeyboard/include
            ${ANDROID_ROOT}/vendor/microsoft/proprietary/windowsconnectivity/crossdevice/virtualkeyboard
            ${OUT_INTERMEDIATES_ROOT}/vendor/microsoft/proprietary/windowsconnectivity/crossdevice/virtualkeyboard/virtual_keyboard/android_${OUT_ARCH_CPU}/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(virtual_keyboard PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(virtual_keyboard
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libutils
            packagemanager_aidl-cpp
    )
endif ()
#################### virtual_keyboard ####################

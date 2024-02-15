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

project(libinput_fuzz_propertymap)

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

#################### libinput ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/input/Input.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/InputDevice.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/InputEventLabels.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/InputVerifier.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/Keyboard.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/KeyCharacterMap.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/KeyLayoutMap.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/MotionPredictor.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/PrintTools.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/PropertyMap.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/TfLiteMotionPredictor.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/TouchVideoFrame.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/VelocityControl.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/VelocityTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/VirtualInputDevice.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/VirtualKeyMap.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/InputTransport.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/input/libinput/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IInputFlinger.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/input/libinput/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IInputConstants.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/input/libinput/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/InputEventInjectionResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/input/libinput/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/InputEventInjectionSync.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/input/libinput/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/InputConfig.cpp
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
    add_library(libinput SHARED ${SOURCE_FILES})

    target_include_directories(libinput
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/input
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/libs/input/libinput/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libinput PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libinput
            flatbuffer_headers
            jni_headers
            libPlatformProperties
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgui_window_info_static
            libhostgraphics
            liblog
            libm
            libtflite_static
            libui-types
            libutils
            libvintf
            tensorflow_headers
    )
endif ()
#################### libinput ####################

#################### libinput_fuzz_propertymap ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/input/PropertyMap.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/input/PropertyMap_fuzz.cpp
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
    add_executable(libinput_fuzz_propertymap ${SOURCE_FILES})

    target_include_directories(libinput_fuzz_propertymap
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/input
    )

    # Add c/c++ flags
    target_compile_options(libinput_fuzz_propertymap PRIVATE
            -Werror
    )

    target_link_libraries(libinput_fuzz_propertymap
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libinput_fuzz_propertymap ####################

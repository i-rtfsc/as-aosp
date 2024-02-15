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

project(VtsHalEvsV1_1CameraOpenFuzz)

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

#################### VtsHalEvsV1_1CameraOpenFuzz ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/evs/1.1/vts/fuzzing/VtsHalEvsV1_1CameraOpenFuzz.cpp
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
    add_executable(VtsHalEvsV1_1CameraOpenFuzz ${SOURCE_FILES})

    target_include_directories(VtsHalEvsV1_1CameraOpenFuzz
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/evs/1.1/vts/fuzzing
    )

    # Add c/c++ flags
    target_compile_options(VtsHalEvsV1_1CameraOpenFuzz PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(VtsHalEvsV1_1CameraOpenFuzz
            VtsHalHidlTestUtils
            android.hardware.automotive.evs-1.0
            android.hardware.automotive.evs-1.1
            android.hardware.automotive.evs-common-default-lib
            android.hardware.camera.device-3.2
            android.hardware.graphics.common-1.0
            android.hardware.graphics.common-1.1
            android.hardware.graphics.common-1.2
            libbase
            libc
            libc++
            libc++demangle
            libcamera_metadata
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libhidl-gen-utils
            libhidl_gtest_helper
            libhidlbase
            liblog
            libm
            libui
            libutils
    )
endif ()
#################### VtsHalEvsV1_1CameraOpenFuzz ####################

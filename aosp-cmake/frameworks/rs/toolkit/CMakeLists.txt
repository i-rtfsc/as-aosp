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

project(librenderscripttoolkit)

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

#################### librenderscripttoolkit ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Blend.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Blur.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/ColorMatrix.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Convolve3x3.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Convolve5x5.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Histogram.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Lut.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Lut3d.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/RenderScriptToolkit.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Resize.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/TaskProcessor.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Utils.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/YuvToRgb.cpp
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Blend_advsimd.S
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Blur_advsimd.S
        ${ANDROID_ROOT}/frameworks/rs/toolkit/ColorMatrix_advsimd.S
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Convolve_advsimd.S
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Lut3d_advsimd.S
        ${ANDROID_ROOT}/frameworks/rs/toolkit/Resize_advsimd.S
        ${ANDROID_ROOT}/frameworks/rs/toolkit/YuvToRgb_advsimd.S
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
    add_library(librenderscripttoolkit SHARED ${SOURCE_FILES})

    target_include_directories(librenderscripttoolkit
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs/toolkit
    )

    # Add c/c++ flags
    target_compile_options(librenderscripttoolkit PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(librenderscripttoolkit
            cpufeatures
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libjnigraphics
            liblog
            libm
            libnativehelper
    )
endif ()
#################### librenderscripttoolkit ####################

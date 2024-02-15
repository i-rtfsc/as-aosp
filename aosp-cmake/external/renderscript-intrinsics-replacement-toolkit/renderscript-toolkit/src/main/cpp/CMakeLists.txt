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

project(librenderscripttoolkit_native_static)

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

#################### librenderscripttoolkit_native_static ####################
set(FILES
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Blend.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Blur.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/ColorMatrix.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Convolve3x3.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Convolve5x5.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Histogram.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Lut.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Lut3d.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/RenderScriptToolkit.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Resize.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/TaskProcessor.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Utils.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/YuvToRgb.cpp
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Blend_advsimd.S
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Blur_advsimd.S
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/ColorMatrix_advsimd.S
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Convolve_advsimd.S
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Lut3d_advsimd.S
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/Resize_advsimd.S
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/YuvToRgb_advsimd.S
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
    add_library(librenderscripttoolkit_native_static SHARED ${SOURCE_FILES})

    target_include_directories(librenderscripttoolkit_native_static
            PUBLIC
            ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp
    )

    # Add c/c++ flags
    target_compile_options(librenderscripttoolkit_native_static PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(librenderscripttoolkit_native_static
            cpufeatures
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libjnigraphics
            liblog
            libm
    )
endif ()
#################### librenderscripttoolkit_native_static ####################

#################### librenderscript-toolkit ####################
set(FILES
        ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp/JniEntryPoints.cpp
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
    add_library(librenderscript-toolkit SHARED ${SOURCE_FILES})

    target_include_directories(librenderscript-toolkit
            PUBLIC
            ${ANDROID_ROOT}/external/renderscript-intrinsics-replacement-toolkit/renderscript-toolkit/src/main/cpp
    )

    # Add c/c++ flags
    target_compile_options(librenderscript-toolkit PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(librenderscript-toolkit
            cpufeatures
            jni_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libjnigraphics
            liblog
            libm
            librenderscripttoolkit_native_static
    )
endif ()
#################### librenderscript-toolkit ####################

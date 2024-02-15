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

project(libRSCpuRef)

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

#################### libRSCpuRef ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuCore.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuExecutable.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuScript.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuRuntimeMath.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuScriptGroup.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuScriptGroup2.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsic.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsic3DLUT.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicBLAS.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicBlend.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicBlur.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicColorMatrix.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicConvolve3x3.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicConvolve5x5.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicHistogram.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicResize.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicLUT.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsicYuvToRGB.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_3DLUT.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_Convolve.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_Blur.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_ColorMatrix.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_Resize.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_YuvToRGB.S
        ${ANDROID_ROOT}/frameworks/rs/cpu_ref/rsCpuIntrinsics_advsimd_Blend.S
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
    add_library(libRSCpuRef SHARED ${SOURCE_FILES})

    target_include_directories(libRSCpuRef
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs/cpu_ref
    )

    # Add c/c++ flags
    target_compile_options(libRSCpuRef PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libRSCpuRef
            libRS_internal
            libbcinfo
            libblas
            libbnnmlowp
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libhardware_headers
            liblog
            libm
            libutils_headers
            libz
    )
endif ()
#################### libRSCpuRef ####################

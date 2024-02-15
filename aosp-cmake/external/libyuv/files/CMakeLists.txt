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

project(libyuv)

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

#################### libyuv ####################
set(FILES
        ${ANDROID_ROOT}/external/libyuv/files/source/compare.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/compare_common.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/compare_gcc.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/compare_msa.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/compare_neon.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/compare_neon64.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/convert.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/convert_argb.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/convert_from.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/convert_from_argb.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/convert_jpeg.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/convert_to_argb.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/convert_to_i420.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/cpu_id.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/mjpeg_decoder.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/mjpeg_validate.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/planar_functions.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/rotate.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/rotate_any.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/rotate_argb.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/rotate_common.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/rotate_gcc.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/rotate_msa.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/rotate_neon.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/rotate_neon64.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/row_any.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/row_common.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/row_gcc.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/row_msa.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/row_neon.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/row_neon64.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale_any.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale_argb.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale_common.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale_gcc.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale_msa.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale_neon.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale_neon64.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale_rgb.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/scale_uv.cc
        ${ANDROID_ROOT}/external/libyuv/files/source/video_common.cc
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
    add_library(libyuv SHARED ${SOURCE_FILES})

    target_include_directories(libyuv
            PUBLIC
            ${ANDROID_ROOT}/external/libyuv/files/include
            ${ANDROID_ROOT}/external/libyuv/files
    )

    # Add c/c++ flags
    target_compile_options(libyuv PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libyuv
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libjpeg
            libm
    )
endif ()
#################### libyuv ####################

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

project(libpdfium-font)

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

#################### libpdfium-font ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cfx_cttgsubtable.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cfx_stockfontarray.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_cid2unicodemap.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_cidfont.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_cmap.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_cmapmanager.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_cmapparser.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_font.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_fontencoding.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_fontglobals.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_simplefont.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_tounicodemap.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_truetypefont.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_type1font.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_type3char.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font/cpdf_type3font.cpp
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
    add_library(libpdfium-font SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-font
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/font
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-font PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-font
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpdfium-cmaps
            libpdfium-fxcrt
            libpdfium-fxge
            libpdfium-parser
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-font ####################

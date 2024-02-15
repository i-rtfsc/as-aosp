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

project(libpdfium-render)

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

#################### libpdfium-render ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_charposlist.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_devicebuffer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_docrenderdata.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_imagecacheentry.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_imageloader.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_imagerenderer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_pagerendercache.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_pagerendercontext.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_progressiverenderer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_rendercontext.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_renderoptions.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_rendershading.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_renderstatus.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_scaledrenderbuffer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_textrenderer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_type3cache.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render/cpdf_type3glyphmap.cpp
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
    add_library(libpdfium-render SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-render
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/render
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-render PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-render
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpdfium-constants
            libpdfium-font
            libpdfium-fxcodec
            libpdfium-fxcrt
            libpdfium-fxge
            libpdfium-page
            libpdfium-parser
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-render ####################

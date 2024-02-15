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

project(libpdfium-page)

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

#################### libpdfium-page ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_allstates.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_annotcontext.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_clippath.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_color.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_colorspace.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_colorstate.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_contentmarkitem.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_contentmarks.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_contentparser.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_devicecs.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_dib.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_docpagedata.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_expintfunc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_form.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_formobject.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_function.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_generalstate.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_graphicstates.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_iccprofile.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_image.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_imageobject.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_meshstream.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_occontext.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_page.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_pagemodule.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_pageobject.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_pageobjectholder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_path.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_pathobject.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_pattern.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_patterncs.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_psengine.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_psfunc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_sampledfunc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_shadingobject.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_shadingpattern.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_stitchfunc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_streamcontentparser.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_streamparser.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_textobject.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_textstate.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_tilingpattern.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_transferfunc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_transferfuncdib.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page/cpdf_transparency.cpp
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
    add_library(libpdfium-page SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-page
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/page
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-page PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-page
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
            libpdfium-parser
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-page ####################

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

project(libpdfium-fpdfdoc)

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

#################### libpdfium-fpdfdoc ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cba_fontmap.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cline.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_aaction.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_action.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_annot.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_annotlist.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_apsettings.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_bookmark.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_bookmarktree.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_color_utils.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_defaultappearance.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_dest.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_filespec.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_formcontrol.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_formfield.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_icon.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_iconfit.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_interactiveform.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_link.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_linklist.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_metadata.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_nametree.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_numbertree.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_pagelabel.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_structelement.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_structtree.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_variabletext.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpdf_viewerpreferences.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpvt_fontmap.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpvt_generateap.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/cpvt_wordinfo.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/csection.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc/ctypeset.cpp
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
    add_library(libpdfium-fpdfdoc SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-fpdfdoc
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/core/fpdfdoc
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-fpdfdoc PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-fpdfdoc
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpdfium-constants
            libpdfium-font
            libpdfium-fxcrt
            libpdfium-fxge
            libpdfium-page
            libpdfium-parser
            libpdfium-render
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-fpdfdoc ####################

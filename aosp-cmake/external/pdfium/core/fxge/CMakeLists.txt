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

project(libpdfium-fxge)

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

#################### libpdfium-fxge ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_cliprgn.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_color.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_face.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_folderfontinfo.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_font.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_fontcache.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_fontmapper.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_fontmgr.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_gemodule.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_glyphbitmap.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_glyphcache.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_graphstate.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_graphstatedata.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_pathdata.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_renderdevice.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_substfont.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/cfx_unicodeencoding.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fx_font.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fx_ge_fontmap.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/renderdevicedriver_iface.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/scoped_font_transform.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/text_char_pos.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/text_glyph_pos.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cfx_bitmapcomposer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cfx_bitmapstorer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cfx_cmyk_to_srgb.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cfx_dibbase.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cfx_dibitmap.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cfx_imagerenderer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cfx_imagestretcher.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cfx_imagetransformer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cfx_scanlinecompositor.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/cstretchengine.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/dib/fx_dib_main.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/freetype/fx_freetype.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitDingbats.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitFixed.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitFixedBold.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitFixedBoldItalic.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitFixedItalic.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSans.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSansBold.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSansBoldItalic.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSansItalic.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSansMM.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSerif.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSerifBold.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSerifBoldItalic.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSerifItalic.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSerifMM.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/fontdata/chromefontdata/FoxitSymbol.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/agg/fx_agg_driver.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/android/cfpf_skiadevicemodule.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/android/cfpf_skiafont.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/android/cfpf_skiafontmgr.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/android/cfpf_skiapathfont.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/android/cfx_androidfontinfo.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxge/android/fx_android_imp.cpp
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
    add_library(libpdfium-fxge SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-fxge
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/core/fxge
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-fxge PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-fxge
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpdfium-agg
            libpdfium-fxcrt
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-fxge ####################

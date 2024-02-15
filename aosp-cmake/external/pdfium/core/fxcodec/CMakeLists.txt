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

project(libpdfium-fxcodec)

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

#################### libpdfium-fxcodec ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/cfx_codec_memory.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/fx_codec.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/scanlinedecoder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/basic/basicmodule.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/fax/faxmodule.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/flate/flatemodule.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/icc/iccmodule.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_ArithDecoder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_ArithIntDecoder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_BitStream.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_Context.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_DocumentContext.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_GrdProc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_GrrdProc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_HtrdProc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_HuffmanDecoder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_HuffmanTable.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_Image.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_PatternDict.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_PddProc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_SddProc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_Segment.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_SymbolDict.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/JBig2_TrdProc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jbig2/jbig2module.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jpeg/jpegmodule.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jpx/cjpx_decoder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jpx/jpx_decode_utils.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcodec/jpx/jpxmodule.cpp
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
    add_library(libpdfium-fxcodec SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-fxcodec
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/core/fxcodec
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-fxcodec PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-fxcodec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libjpeg
            libm
            libpdfium-fxcrt
            libpdfium-fxge
            libpdfium-lcms2
            libpdfium-libopenjpeg2
            libz
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-fxcodec ####################

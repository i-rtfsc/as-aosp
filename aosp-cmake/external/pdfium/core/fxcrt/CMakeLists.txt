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

project(libpdfium-fxcrt)

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

#################### libpdfium-fxcrt ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/bytestring.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_binarybuf.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_bitstream.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_datetime.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_fileaccess_posix.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_readonlymemorystream.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_seekablestreamproxy.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_timer.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_utf8decoder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_utf8encoder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/cfx_widetextbuf.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_bidi.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_codepage.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_coordinates.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_extension.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_memory.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_number.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_random.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_stream.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_string.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_system.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/fx_unicode.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/observed_ptr.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/widestring.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/xml/cfx_xmlchardata.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/xml/cfx_xmldocument.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/xml/cfx_xmlelement.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/xml/cfx_xmlinstruction.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/xml/cfx_xmlnode.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/xml/cfx_xmlparser.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fxcrt/xml/cfx_xmltext.cpp
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
    add_library(libpdfium-fxcrt SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-fxcrt
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/core/fxcrt
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-fxcrt PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-fxcrt
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libicu
            libm
            libpdfium-third_party-base
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-fxcrt ####################

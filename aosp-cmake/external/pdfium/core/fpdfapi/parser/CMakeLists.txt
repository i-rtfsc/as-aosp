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

project(libpdfium-parser)

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

#################### libpdfium-parser ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cfdf_document.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_array.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_boolean.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_cross_ref_avail.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_cross_ref_table.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_crypto_handler.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_data_avail.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_dictionary.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_document.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_encryptor.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_flateencoder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_hint_tables.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_indirect_object_holder.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_linearized_header.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_name.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_null.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_number.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_object.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_object_avail.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_object_stream.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_object_walker.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_page_object_avail.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_parser.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_read_validator.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_reference.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_security_handler.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_simple_parser.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_stream.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_stream_acc.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_string.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/cpdf_syntax_parser.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/fpdf_parser_decode.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser/fpdf_parser_utility.cpp
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
    add_library(libpdfium-parser SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-parser
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/parser
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-parser PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-parser
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpdfium-constants
            libpdfium-fdrm
            libpdfium-fxcodec
            libpdfium-fxcrt
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-parser ####################

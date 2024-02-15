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

project(libsfntly)

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

#################### libsfntly ####################
set(FILES
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/font.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/font_factory.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/tag.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/data/byte_array.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/data/font_data.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/data/font_input_stream.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/data/font_output_stream.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/data/growable_memory_byte_array.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/data/memory_byte_array.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/data/readable_font_data.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/data/writable_font_data.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/port/file_input_stream.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/port/lock.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/port/memory_input_stream.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/port/memory_output_stream.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/byte_array_table_builder.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/font_data_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/generic_table_builder.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/header.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/subtable.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/table_based_table_builder.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/big_glyph_metrics.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/bitmap_glyph.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/bitmap_glyph_info.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/bitmap_size_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/composite_bitmap_glyph.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/ebdt_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/eblc_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/ebsc_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/glyph_metrics.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/index_sub_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/index_sub_table_format1.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/index_sub_table_format2.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/index_sub_table_format3.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/index_sub_table_format4.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/index_sub_table_format5.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/simple_bitmap_glyph.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/bitmap/small_glyph_metrics.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/core/cmap_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/core/font_header_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/core/horizontal_device_metrics_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/core/horizontal_header_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/core/horizontal_metrics_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/core/maximum_profile_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/core/name_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/core/os2_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/truetype/glyph_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sfntly/table/truetype/loca_table.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sample/chromium/font_subsetter.cc
        ${ANDROID_ROOT}/external/sfntly/cpp/src/sample/chromium/subsetter_impl.cc
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
    add_library(libsfntly SHARED ${SOURCE_FILES})

    target_include_directories(libsfntly
            PUBLIC
            ${ANDROID_ROOT}/external/sfntly/cpp/src
            ${ANDROID_ROOT}/external/sfntly
    )

    # Add c/c++ flags
    target_compile_options(libsfntly PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
            -DNDEBUG
    )

    target_link_libraries(libsfntly
            libandroidicu
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libicui18n
            libicuuc
            libm
            libwinpthread
    )
endif ()
#################### libsfntly ####################

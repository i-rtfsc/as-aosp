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

project(libimage_io)

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

#################### libimage_io ####################
set(FILES
        ${ANDROID_ROOT}/external/image_io/src/base/byte_buffer.cc
        ${ANDROID_ROOT}/external/image_io/src/base/byte_pointer_data_destination.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_context.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_line_map.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_range_tracking_destination.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_scanner.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_segment.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_segment_data_source.cc
        ${ANDROID_ROOT}/external/image_io/src/base/istream_ref_data_source.cc
        ${ANDROID_ROOT}/external/image_io/src/base/message_handler.cc
        ${ANDROID_ROOT}/external/image_io/src/base/message_writer.cc
        ${ANDROID_ROOT}/external/image_io/src/base/ostream_ref_data_destination.cc
        ${ANDROID_ROOT}/external/image_io/src/base/string_ref_data_source.cc
        ${ANDROID_ROOT}/external/image_io/src/extras/base64_decoder_data_destination.cc
        ${ANDROID_ROOT}/external/image_io/src/gcontainer/gcontainer.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_apple_depth_builder.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_image_extractor.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_info_builder.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_marker.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_scanner.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_segment.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_segment_builder.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_segment_lister.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_xmp_data_extractor.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_xmp_info.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_xmp_info_builder.cc
        ${ANDROID_ROOT}/external/image_io/src/utils/file_utils.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_attribute_rule.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_cdata_and_comment_rules.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_element_rules.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_handler.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_pi_rule.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_reader.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_rule.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_token_context.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_writer.cc
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
    add_library(libimage_io SHARED ${SOURCE_FILES})

    target_include_directories(libimage_io
            PUBLIC
            ${ANDROID_ROOT}/external/image_io/includes
            ${ANDROID_ROOT}/external/image_io
    )

    # Add c/c++ flags
    target_compile_options(libimage_io PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libimage_io
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libimage_io-headers
            libm
            libmodpb64
    )
endif ()
#################### libimage_io ####################

#################### libimage_io_ndk ####################
set(FILES
        ${ANDROID_ROOT}/external/image_io/src/base/byte_buffer.cc
        ${ANDROID_ROOT}/external/image_io/src/base/byte_pointer_data_destination.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_context.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_line_map.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_range_tracking_destination.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_scanner.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_segment.cc
        ${ANDROID_ROOT}/external/image_io/src/base/data_segment_data_source.cc
        ${ANDROID_ROOT}/external/image_io/src/base/istream_ref_data_source.cc
        ${ANDROID_ROOT}/external/image_io/src/base/message_handler.cc
        ${ANDROID_ROOT}/external/image_io/src/base/message_writer.cc
        ${ANDROID_ROOT}/external/image_io/src/base/ostream_ref_data_destination.cc
        ${ANDROID_ROOT}/external/image_io/src/base/string_ref_data_source.cc
        ${ANDROID_ROOT}/external/image_io/src/extras/base64_decoder_data_destination.cc
        ${ANDROID_ROOT}/external/image_io/src/gcontainer/gcontainer.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_apple_depth_builder.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_image_extractor.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_info_builder.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_marker.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_scanner.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_segment.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_segment_builder.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_segment_lister.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_xmp_data_extractor.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_xmp_info.cc
        ${ANDROID_ROOT}/external/image_io/src/jpeg/jpeg_xmp_info_builder.cc
        ${ANDROID_ROOT}/external/image_io/src/utils/file_utils.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_attribute_rule.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_cdata_and_comment_rules.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_element_rules.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_handler.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_pi_rule.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_reader.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_rule.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_token_context.cc
        ${ANDROID_ROOT}/external/image_io/src/xml/xml_writer.cc
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
    add_library(libimage_io_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libimage_io_ndk
            PUBLIC
            ${ANDROID_ROOT}/external/image_io/includes
            ${ANDROID_ROOT}/external/image_io
    )

    # Add c/c++ flags
    target_compile_options(libimage_io_ndk PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libimage_io_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libimage_io_ndk-headers
            libm
            libmodpb64_ndk
    )
endif ()
#################### libimage_io_ndk ####################

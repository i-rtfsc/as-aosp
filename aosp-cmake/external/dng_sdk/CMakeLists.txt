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

project(dng_validate)

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

#################### libdng_sdk_validate ####################
set(FILES
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_1d_function.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_1d_table.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_abort_sniffer.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_area_task.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_bad_pixels.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_bottlenecks.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_camera_profile.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_color_space.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_color_spec.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_date_time.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_exceptions.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_exif.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_file_stream.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_filter_task.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_fingerprint.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_gain_map.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_globals.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_host.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_hue_sat_map.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_ifd.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_image_writer.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_info.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_iptc.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_jpeg_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_jpeg_memory_source.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_lens_correction.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_linearization_info.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_lossless_jpeg.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_matrix.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_memory.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_memory_stream.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_misc_opcodes.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_mosaic_info.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_mutex.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_negative.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_opcode_list.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_opcodes.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_orientation.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_parse_utils.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_pixel_buffer.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_point.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_preview.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_pthread.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_rational.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_read_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_rect.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_ref_counted_block.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_reference.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_render.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_resample.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_safe_arithmetic.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_shared.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_simple_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_spline.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_stream.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_string.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_string_list.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_tag_types.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_temperature.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_tile_iterator.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_tone_curve.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_utils.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_xy_coord.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_xmp.cpp
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
    add_library(libdng_sdk_validate SHARED ${SOURCE_FILES})

    target_include_directories(libdng_sdk_validate
            PUBLIC
            ${ANDROID_ROOT}/external/dng_sdk/source
            ${ANDROID_ROOT}/external/dng_sdk
    )

    # Add c/c++ flags
    target_compile_options(libdng_sdk_validate PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libdng_sdk_validate
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libjpeg_static_ndk
            libm
            libz
    )
endif ()
#################### libdng_sdk_validate ####################

#################### dng_validate ####################
set(FILES
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_1d_function.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_1d_table.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_abort_sniffer.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_area_task.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_bad_pixels.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_bottlenecks.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_camera_profile.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_color_space.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_color_spec.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_date_time.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_exceptions.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_exif.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_file_stream.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_filter_task.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_fingerprint.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_gain_map.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_globals.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_host.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_hue_sat_map.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_ifd.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_image_writer.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_info.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_iptc.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_jpeg_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_jpeg_memory_source.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_lens_correction.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_linearization_info.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_lossless_jpeg.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_matrix.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_memory.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_memory_stream.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_misc_opcodes.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_mosaic_info.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_mutex.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_negative.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_opcode_list.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_opcodes.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_orientation.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_parse_utils.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_pixel_buffer.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_point.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_preview.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_pthread.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_rational.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_read_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_rect.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_ref_counted_block.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_reference.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_render.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_resample.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_safe_arithmetic.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_shared.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_simple_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_spline.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_stream.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_string.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_string_list.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_tag_types.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_temperature.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_tile_iterator.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_tone_curve.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_utils.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_xy_coord.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_xmp.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_validate.cpp
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
    add_executable(dng_validate ${SOURCE_FILES})

    target_include_directories(dng_validate
            PUBLIC
            ${ANDROID_ROOT}/external/dng_sdk
    )

    # Add c/c++ flags
    target_compile_options(dng_validate PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(dng_validate
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libjpeg
            libm
            libz
    )
endif ()
#################### dng_validate ####################

#################### libdng_sdk ####################
set(FILES
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_1d_function.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_1d_table.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_abort_sniffer.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_area_task.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_bad_pixels.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_bottlenecks.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_camera_profile.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_color_space.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_color_spec.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_date_time.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_exceptions.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_exif.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_file_stream.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_filter_task.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_fingerprint.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_gain_map.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_globals.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_host.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_hue_sat_map.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_ifd.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_image_writer.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_info.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_iptc.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_jpeg_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_jpeg_memory_source.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_lens_correction.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_linearization_info.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_lossless_jpeg.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_matrix.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_memory.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_memory_stream.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_misc_opcodes.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_mosaic_info.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_mutex.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_negative.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_opcode_list.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_opcodes.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_orientation.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_parse_utils.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_pixel_buffer.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_point.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_preview.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_pthread.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_rational.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_read_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_rect.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_ref_counted_block.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_reference.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_render.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_resample.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_safe_arithmetic.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_shared.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_simple_image.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_spline.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_stream.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_string.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_string_list.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_tag_types.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_temperature.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_tile_iterator.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_tone_curve.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_utils.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_xy_coord.cpp
        ${ANDROID_ROOT}/external/dng_sdk/source/dng_xmp.cpp
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
    add_library(libdng_sdk SHARED ${SOURCE_FILES})

    target_include_directories(libdng_sdk
            PUBLIC
            ${ANDROID_ROOT}/external/dng_sdk/source
            ${ANDROID_ROOT}/external/dng_sdk
    )

    # Add c/c++ flags
    target_compile_options(libdng_sdk PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libdng_sdk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcompiler_rt-extras
            libdl
            libjpeg
            libm
            libwinpthread
            libz
    )
endif ()
#################### libdng_sdk ####################

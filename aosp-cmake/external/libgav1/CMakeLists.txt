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

project(libgav1_decoder_fuzzer_frame_parallel)

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

#################### libgav1_decoder_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libgav1/examples/file_reader.cc
        ${ANDROID_ROOT}/external/libgav1/examples/file_reader_constants.cc
        ${ANDROID_ROOT}/external/libgav1/examples/file_reader_factory.cc
        ${ANDROID_ROOT}/external/libgav1/examples/ivf_parser.cc
        ${ANDROID_ROOT}/external/libgav1/tests/fuzzer/decoder_fuzzer.cc
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
    add_executable(libgav1_decoder_fuzzer ${SOURCE_FILES})

    target_include_directories(libgav1_decoder_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libgav1
            ${ANDROID_ROOT}/external/libgav1/src
    )

    # Add c/c++ flags
    target_compile_options(libgav1_decoder_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libgav1_decoder_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libgav1
            libm
    )
endif ()
#################### libgav1_decoder_fuzzer ####################

#################### libgav1_obu_parser_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libgav1/examples/file_reader.cc
        ${ANDROID_ROOT}/external/libgav1/examples/file_reader_constants.cc
        ${ANDROID_ROOT}/external/libgav1/examples/file_reader_factory.cc
        ${ANDROID_ROOT}/external/libgav1/examples/ivf_parser.cc
        ${ANDROID_ROOT}/external/libgav1/tests/fuzzer/obu_parser_fuzzer.cc
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
    add_executable(libgav1_obu_parser_fuzzer ${SOURCE_FILES})

    target_include_directories(libgav1_obu_parser_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libgav1
            ${ANDROID_ROOT}/external/libgav1/src
    )

    # Add c/c++ flags
    target_compile_options(libgav1_obu_parser_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libgav1_obu_parser_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libgav1
            libm
    )
endif ()
#################### libgav1_obu_parser_fuzzer ####################

#################### libgav1_decoder_fuzzer_frame_parallel ####################
set(FILES
        ${ANDROID_ROOT}/external/libgav1/examples/file_reader.cc
        ${ANDROID_ROOT}/external/libgav1/examples/file_reader_constants.cc
        ${ANDROID_ROOT}/external/libgav1/examples/file_reader_factory.cc
        ${ANDROID_ROOT}/external/libgav1/examples/ivf_parser.cc
        ${ANDROID_ROOT}/external/libgav1/tests/fuzzer/decoder_fuzzer_frame_parallel.cc
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
    add_executable(libgav1_decoder_fuzzer_frame_parallel ${SOURCE_FILES})

    target_include_directories(libgav1_decoder_fuzzer_frame_parallel
            PUBLIC
            ${ANDROID_ROOT}/external/libgav1
            ${ANDROID_ROOT}/external/libgav1/src
    )

    # Add c/c++ flags
    target_compile_options(libgav1_decoder_fuzzer_frame_parallel PRIVATE
            -Werror
    )

    target_link_libraries(libgav1_decoder_fuzzer_frame_parallel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libgav1
            libm
    )
endif ()
#################### libgav1_decoder_fuzzer_frame_parallel ####################

#################### libgav1 ####################
set(FILES
        ${ANDROID_ROOT}/external/libgav1/src/buffer_pool.cc
        ${ANDROID_ROOT}/external/libgav1/src/decoder.cc
        ${ANDROID_ROOT}/external/libgav1/src/decoder_impl.cc
        ${ANDROID_ROOT}/external/libgav1/src/decoder_settings.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/average_blend_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/cdef_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/convolve_10bit_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/convolve_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/distance_weighted_blend_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/film_grain_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/intra_edge_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/intrapred_cfl_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/intrapred_directional_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/intrapred_filter_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/intrapred_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/intrapred_smooth_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/inverse_transform_10bit_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/inverse_transform_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/loop_filter_10bit_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/loop_filter_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/loop_restoration_10bit_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/loop_restoration_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/mask_blend_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/motion_field_projection_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/motion_vector_search_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/obmc_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/super_res_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/warp_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/arm/weight_mask_neon.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/average_blend.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/cdef.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/constants.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/convolve.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/distance_weighted_blend.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/dsp.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/film_grain.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/intra_edge.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/intrapred.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/intrapred_cfl.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/intrapred_directional.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/intrapred_filter.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/intrapred_smooth.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/inverse_transform.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/loop_filter.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/loop_restoration.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/mask_blend.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/motion_field_projection.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/motion_vector_search.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/obmc.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/super_res.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/warp.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/weight_mask.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/average_blend_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/cdef_avx2.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/cdef_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/convolve_avx2.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/convolve_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/distance_weighted_blend_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/film_grain_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/intra_edge_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/intrapred_cfl_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/intrapred_directional_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/intrapred_filter_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/intrapred_smooth_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/intrapred_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/inverse_transform_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/loop_filter_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/loop_restoration_10bit_avx2.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/loop_restoration_10bit_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/loop_restoration_avx2.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/loop_restoration_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/mask_blend_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/motion_field_projection_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/motion_vector_search_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/obmc_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/super_res_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/warp_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/dsp/x86/weight_mask_sse4.cc
        ${ANDROID_ROOT}/external/libgav1/src/film_grain.cc
        ${ANDROID_ROOT}/external/libgav1/src/frame_buffer.cc
        ${ANDROID_ROOT}/external/libgav1/src/internal_frame_buffer_list.cc
        ${ANDROID_ROOT}/external/libgav1/src/loop_restoration_info.cc
        ${ANDROID_ROOT}/external/libgav1/src/motion_vector.cc
        ${ANDROID_ROOT}/external/libgav1/src/obu_parser.cc
        ${ANDROID_ROOT}/external/libgav1/src/post_filter/cdef.cc
        ${ANDROID_ROOT}/external/libgav1/src/post_filter/deblock.cc
        ${ANDROID_ROOT}/external/libgav1/src/post_filter/loop_restoration.cc
        ${ANDROID_ROOT}/external/libgav1/src/post_filter/post_filter.cc
        ${ANDROID_ROOT}/external/libgav1/src/post_filter/super_res.cc
        ${ANDROID_ROOT}/external/libgav1/src/prediction_mask.cc
        ${ANDROID_ROOT}/external/libgav1/src/quantizer.cc
        ${ANDROID_ROOT}/external/libgav1/src/reconstruction.cc
        ${ANDROID_ROOT}/external/libgav1/src/residual_buffer_pool.cc
        ${ANDROID_ROOT}/external/libgav1/src/status_code.cc
        ${ANDROID_ROOT}/external/libgav1/src/symbol_decoder_context.cc
        ${ANDROID_ROOT}/external/libgav1/src/threading_strategy.cc
        ${ANDROID_ROOT}/external/libgav1/src/tile/bitstream/mode_info.cc
        ${ANDROID_ROOT}/external/libgav1/src/tile/bitstream/palette.cc
        ${ANDROID_ROOT}/external/libgav1/src/tile/bitstream/partition.cc
        ${ANDROID_ROOT}/external/libgav1/src/tile/bitstream/transform_size.cc
        ${ANDROID_ROOT}/external/libgav1/src/tile/prediction.cc
        ${ANDROID_ROOT}/external/libgav1/src/tile/tile.cc
        ${ANDROID_ROOT}/external/libgav1/src/tile_scratch_buffer.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/bit_reader.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/block_parameters_holder.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/constants.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/cpu.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/entropy_decoder.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/executor.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/logging.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/raw_bit_reader.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/segmentation.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/segmentation_map.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/threadpool.cc
        ${ANDROID_ROOT}/external/libgav1/src/version.cc
        ${ANDROID_ROOT}/external/libgav1/src/warp_prediction.cc
        ${ANDROID_ROOT}/external/libgav1/src/yuv_buffer.cc
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
    add_library(libgav1 SHARED ${SOURCE_FILES})

    target_include_directories(libgav1
            PUBLIC
            ${ANDROID_ROOT}/external/libgav1
            ${ANDROID_ROOT}/external/libgav1/src
    )

    # Add c/c++ flags
    target_compile_options(libgav1 PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libgav1
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libgav1 ####################

#################### libgav1_parser ####################
set(FILES
        ${ANDROID_ROOT}/external/libgav1/src/buffer_pool.cc
        ${ANDROID_ROOT}/external/libgav1/src/frame_buffer.cc
        ${ANDROID_ROOT}/external/libgav1/src/obu_parser.cc
        ${ANDROID_ROOT}/external/libgav1/src/internal_frame_buffer_list.cc
        ${ANDROID_ROOT}/external/libgav1/src/status_code.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/bit_reader.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/constants.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/logging.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/raw_bit_reader.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/segmentation.cc
        ${ANDROID_ROOT}/external/libgav1/src/utils/segmentation_map.cc
        ${ANDROID_ROOT}/external/libgav1/src/symbol_decoder_context.cc
        ${ANDROID_ROOT}/external/libgav1/src/quantizer.cc
        ${ANDROID_ROOT}/external/libgav1/src/yuv_buffer.cc
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
    add_library(libgav1_parser SHARED ${SOURCE_FILES})

    target_include_directories(libgav1_parser
            PUBLIC
            ${ANDROID_ROOT}/external/libgav1
            ${ANDROID_ROOT}/external/libgav1/src
    )

    # Add c/c++ flags
    target_compile_options(libgav1_parser PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libgav1_parser
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libgav1_parser ####################

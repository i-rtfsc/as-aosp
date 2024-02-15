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

project(libwebp_animation_api_fuzzer)

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

#################### libwebp-decode ####################
set(FILES
        ${ANDROID_ROOT}/external/webp/src/dec/alpha_dec.c
        ${ANDROID_ROOT}/external/webp/src/dec/buffer_dec.c
        ${ANDROID_ROOT}/external/webp/src/dec/frame_dec.c
        ${ANDROID_ROOT}/external/webp/src/dec/idec_dec.c
        ${ANDROID_ROOT}/external/webp/src/dec/io_dec.c
        ${ANDROID_ROOT}/external/webp/src/dec/quant_dec.c
        ${ANDROID_ROOT}/external/webp/src/dec/tree_dec.c
        ${ANDROID_ROOT}/external/webp/src/dec/vp8_dec.c
        ${ANDROID_ROOT}/external/webp/src/dec/vp8l_dec.c
        ${ANDROID_ROOT}/external/webp/src/dec/webp_dec.c
        ${ANDROID_ROOT}/external/webp/src/demux/anim_decode.c
        ${ANDROID_ROOT}/external/webp/src/demux/demux.c
        ${ANDROID_ROOT}/external/webp/src/dsp/alpha_processing.c
        ${ANDROID_ROOT}/external/webp/src/dsp/alpha_processing_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/alpha_processing_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/alpha_processing_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/alpha_processing_sse41.c
        ${ANDROID_ROOT}/external/webp/src/dsp/cpu.c
        ${ANDROID_ROOT}/external/webp/src/dsp/dec.c
        ${ANDROID_ROOT}/external/webp/src/dsp/dec_clip_tables.c
        ${ANDROID_ROOT}/external/webp/src/dsp/dec_mips32.c
        ${ANDROID_ROOT}/external/webp/src/dsp/dec_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/dec_msa.c
        ${ANDROID_ROOT}/external/webp/src/dsp/dec_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/dec_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/dec_sse41.c
        ${ANDROID_ROOT}/external/webp/src/dsp/filters.c
        ${ANDROID_ROOT}/external/webp/src/dsp/filters_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/filters_msa.c
        ${ANDROID_ROOT}/external/webp/src/dsp/filters_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/filters_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_msa.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_sse41.c
        ${ANDROID_ROOT}/external/webp/src/dsp/rescaler.c
        ${ANDROID_ROOT}/external/webp/src/dsp/rescaler_mips32.c
        ${ANDROID_ROOT}/external/webp/src/dsp/rescaler_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/rescaler_msa.c
        ${ANDROID_ROOT}/external/webp/src/dsp/rescaler_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/rescaler_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/upsampling.c
        ${ANDROID_ROOT}/external/webp/src/dsp/upsampling_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/upsampling_msa.c
        ${ANDROID_ROOT}/external/webp/src/dsp/upsampling_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/upsampling_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/upsampling_sse41.c
        ${ANDROID_ROOT}/external/webp/src/dsp/yuv.c
        ${ANDROID_ROOT}/external/webp/src/dsp/yuv_mips32.c
        ${ANDROID_ROOT}/external/webp/src/dsp/yuv_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/yuv_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/yuv_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/yuv_sse41.c
        ${ANDROID_ROOT}/external/webp/src/utils/bit_reader_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/color_cache_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/filters_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/huffman_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/quant_levels_dec_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/random_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/rescaler_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/thread_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/utils.c
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
    add_library(libwebp-decode SHARED ${SOURCE_FILES})

    target_include_directories(libwebp-decode
            PUBLIC
            ${ANDROID_ROOT}/external/webp
            ${ANDROID_ROOT}/external/webp/src
    )

    # Add c/c++ flags
    target_compile_options(libwebp-decode PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libwebp-decode
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            webp-headers
    )
endif ()
#################### libwebp-decode ####################

#################### libwebp-encode ####################
set(FILES
        ${ANDROID_ROOT}/external/webp/sharpyuv/sharpyuv.c
        ${ANDROID_ROOT}/external/webp/sharpyuv/sharpyuv_cpu.c
        ${ANDROID_ROOT}/external/webp/sharpyuv/sharpyuv_csp.c
        ${ANDROID_ROOT}/external/webp/sharpyuv/sharpyuv_dsp.c
        ${ANDROID_ROOT}/external/webp/sharpyuv/sharpyuv_gamma.c
        ${ANDROID_ROOT}/external/webp/sharpyuv/sharpyuv_neon.c
        ${ANDROID_ROOT}/external/webp/sharpyuv/sharpyuv_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/cost.c
        ${ANDROID_ROOT}/external/webp/src/dsp/cost_mips32.c
        ${ANDROID_ROOT}/external/webp/src/dsp/cost_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/cost_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/cost_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/cpu.c
        ${ANDROID_ROOT}/external/webp/src/dsp/enc.c
        ${ANDROID_ROOT}/external/webp/src/dsp/enc_mips32.c
        ${ANDROID_ROOT}/external/webp/src/dsp/enc_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/enc_msa.c
        ${ANDROID_ROOT}/external/webp/src/dsp/enc_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/enc_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/enc_sse41.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_enc.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_enc_mips32.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_enc_mips_dsp_r2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_enc_msa.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_enc_neon.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_enc_sse2.c
        ${ANDROID_ROOT}/external/webp/src/dsp/lossless_enc_sse41.c
        ${ANDROID_ROOT}/external/webp/src/dsp/ssim.c
        ${ANDROID_ROOT}/external/webp/src/dsp/ssim_sse2.c
        ${ANDROID_ROOT}/external/webp/src/enc/alpha_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/analysis_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/backward_references_cost_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/backward_references_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/config_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/cost_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/filter_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/frame_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/histogram_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/iterator_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/near_lossless_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/picture_csp_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/picture_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/picture_psnr_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/picture_rescale_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/picture_tools_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/predictor_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/quant_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/syntax_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/token_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/tree_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/vp8l_enc.c
        ${ANDROID_ROOT}/external/webp/src/enc/webp_enc.c
        ${ANDROID_ROOT}/external/webp/src/mux/anim_encode.c
        ${ANDROID_ROOT}/external/webp/src/mux/muxedit.c
        ${ANDROID_ROOT}/external/webp/src/mux/muxinternal.c
        ${ANDROID_ROOT}/external/webp/src/mux/muxread.c
        ${ANDROID_ROOT}/external/webp/src/utils/bit_writer_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/huffman_encode_utils.c
        ${ANDROID_ROOT}/external/webp/src/utils/quant_levels_utils.c
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
    add_library(libwebp-encode SHARED ${SOURCE_FILES})

    target_include_directories(libwebp-encode
            PUBLIC
            ${ANDROID_ROOT}/external/webp
            ${ANDROID_ROOT}/external/webp/src
    )

    # Add c/c++ flags
    target_compile_options(libwebp-encode PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libwebp-encode
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            webp-headers
    )
endif ()
#################### libwebp-encode ####################

#################### libwebp_mux_demux_api_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/webp/tests/fuzzer/mux_demux_api_fuzzer.c
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
    add_executable(libwebp_mux_demux_api_fuzzer ${SOURCE_FILES})

    target_include_directories(libwebp_mux_demux_api_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/webp
            ${ANDROID_ROOT}/external/webp/src
    )

    # Add c/c++ flags
    target_compile_options(libwebp_mux_demux_api_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libwebp_mux_demux_api_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libwebp-decode
            libwebp-encode
    )
endif ()
#################### libwebp_mux_demux_api_fuzzer ####################

#################### libwebp_enc_dec_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/webp/tests/fuzzer/enc_dec_fuzzer.cc
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
    add_executable(libwebp_enc_dec_fuzzer ${SOURCE_FILES})

    target_include_directories(libwebp_enc_dec_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/webp
            ${ANDROID_ROOT}/external/webp/src
    )

    # Add c/c++ flags
    target_compile_options(libwebp_enc_dec_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libwebp_enc_dec_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libwebp-decode
            libwebp-encode
    )
endif ()
#################### libwebp_enc_dec_fuzzer ####################

#################### libwebp_advanced_api_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/webp/tests/fuzzer/advanced_api_fuzzer.c
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
    add_executable(libwebp_advanced_api_fuzzer ${SOURCE_FILES})

    target_include_directories(libwebp_advanced_api_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/webp
            ${ANDROID_ROOT}/external/webp/src
    )

    # Add c/c++ flags
    target_compile_options(libwebp_advanced_api_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libwebp_advanced_api_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libwebp-decode
    )
endif ()
#################### libwebp_advanced_api_fuzzer ####################

#################### libwebp_animencoder_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/webp/tests/fuzzer/animencoder_fuzzer.cc
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
    add_executable(libwebp_animencoder_fuzzer ${SOURCE_FILES})

    target_include_directories(libwebp_animencoder_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/webp
            ${ANDROID_ROOT}/external/webp/src
    )

    # Add c/c++ flags
    target_compile_options(libwebp_animencoder_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libwebp_animencoder_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libwebp-decode
            libwebp-encode
    )
endif ()
#################### libwebp_animencoder_fuzzer ####################

#################### libwebp_simple_api_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/webp/tests/fuzzer/simple_api_fuzzer.c
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
    add_executable(libwebp_simple_api_fuzzer ${SOURCE_FILES})

    target_include_directories(libwebp_simple_api_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/webp
            ${ANDROID_ROOT}/external/webp/src
    )

    # Add c/c++ flags
    target_compile_options(libwebp_simple_api_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libwebp_simple_api_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libwebp-decode
    )
endif ()
#################### libwebp_simple_api_fuzzer ####################

#################### libwebp_animation_api_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/webp/tests/fuzzer/animation_api_fuzzer.c
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
    add_executable(libwebp_animation_api_fuzzer ${SOURCE_FILES})

    target_include_directories(libwebp_animation_api_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/webp
            ${ANDROID_ROOT}/external/webp/src
    )

    # Add c/c++ flags
    target_compile_options(libwebp_animation_api_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libwebp_animation_api_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libwebp-decode
    )
endif ()
#################### libwebp_animation_api_fuzzer ####################

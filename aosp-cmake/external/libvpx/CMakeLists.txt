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

project(vp9_dec_fuzzer)

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

#################### libvpx ####################
set(FILES
        ${ANDROID_ROOT}/external/libvpx/config/arm64/vpx_config.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/alloccommon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/loopfilter_arm.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/bilinearpredict_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/copymem_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/dc_only_idct_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/dequant_idct_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/dequantizeb_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/idct_blk_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/iwalsh_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/loopfiltersimplehorizontaledge_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/loopfiltersimpleverticaledge_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/mbloopfilter_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/shortidct4x4llm_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/sixtappredict_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/arm/neon/vp8_loopfilter_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/blockd.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/dequantize.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/entropy.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/entropymode.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/entropymv.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/extend.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/filter.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/findnearmv.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/generic/systemdependent.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/idct_blk.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/idctllm.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/loopfilter_filters.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/mbpitch.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/modecont.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/quant_common.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/reconinter.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/reconintra.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/reconintra4x4.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/rtcd.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/setupintrarecon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/swapyv12buffer.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/treecoder.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/vp8_loopfilter.c
        ${ANDROID_ROOT}/external/libvpx/vp8/common/vp8_skin_detection.c
        ${ANDROID_ROOT}/external/libvpx/vp8/decoder/dboolhuff.c
        ${ANDROID_ROOT}/external/libvpx/vp8/decoder/decodeframe.c
        ${ANDROID_ROOT}/external/libvpx/vp8/decoder/decodemv.c
        ${ANDROID_ROOT}/external/libvpx/vp8/decoder/detokenize.c
        ${ANDROID_ROOT}/external/libvpx/vp8/decoder/onyxd_if.c
        ${ANDROID_ROOT}/external/libvpx/vp8/decoder/threading.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/arm/neon/denoising_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/arm/neon/fastquantizeb_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/arm/neon/shortfdct_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/arm/neon/vp8_shortwalsh4x4_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/bitstream.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/boolhuff.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/copy_c.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/dct.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/denoising.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/encodeframe.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/encodeintra.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/encodemb.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/encodemv.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/ethreading.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/lookahead.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/mcomp.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/modecosts.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/onyx_if.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/pickinter.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/picklpf.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/ratectrl.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/rdopt.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/segmentation.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/tokenize.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/treewriter.c
        ${ANDROID_ROOT}/external/libvpx/vp8/encoder/vp8_quantize.c
        ${ANDROID_ROOT}/external/libvpx/vp8/vp8_cx_iface.c
        ${ANDROID_ROOT}/external/libvpx/vp8/vp8_dx_iface.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/arm/neon/vp9_highbd_iht4x4_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/arm/neon/vp9_highbd_iht8x8_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/arm/neon/vp9_highbd_iht16x16_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/arm/neon/vp9_iht4x4_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/arm/neon/vp9_iht8x8_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/arm/neon/vp9_iht16x16_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_alloccommon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_blockd.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_common_data.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_entropy.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_entropymode.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_entropymv.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_filter.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_frame_buffers.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_idct.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_loopfilter.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_mvref_common.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_pred_common.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_quant_common.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_reconinter.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_reconintra.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_rtcd.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_scale.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_scan.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_seg_common.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_thread_common.c
        ${ANDROID_ROOT}/external/libvpx/vp9/common/vp9_tile_common.c
        ${ANDROID_ROOT}/external/libvpx/vp9/decoder/vp9_decodeframe.c
        ${ANDROID_ROOT}/external/libvpx/vp9/decoder/vp9_decodemv.c
        ${ANDROID_ROOT}/external/libvpx/vp9/decoder/vp9_decoder.c
        ${ANDROID_ROOT}/external/libvpx/vp9/decoder/vp9_detokenize.c
        ${ANDROID_ROOT}/external/libvpx/vp9/decoder/vp9_dsubexp.c
        ${ANDROID_ROOT}/external/libvpx/vp9/decoder/vp9_job_queue.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/arm/neon/vp9_dct_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/arm/neon/vp9_diamond_search_sad_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/arm/neon/vp9_frame_scale_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/arm/neon/vp9_quantize_neon.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_aq_cyclicrefresh.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_bitstream.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_context_tree.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_cost.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_dct.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_encodeframe.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_encodemb.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_encodemv.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_encoder.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_ethread.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_ext_ratectrl.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_extend.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_frame_scale.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_lookahead.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_mcomp.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_multi_thread.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_noise_estimate.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_picklpf.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_pickmode.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_quantize.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_ratectrl.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_rd.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_rdopt.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_resize.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_segmentation.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_skin_detection.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_speed_features.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_subexp.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_svc_layercontext.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_tokenize.c
        ${ANDROID_ROOT}/external/libvpx/vp9/encoder/vp9_treewriter.c
        ${ANDROID_ROOT}/external/libvpx/vp9/vp9_cx_iface.c
        ${ANDROID_ROOT}/external/libvpx/vp9/vp9_dx_iface.c
        ${ANDROID_ROOT}/external/libvpx/vp9/vp9_iface_common.c
        ${ANDROID_ROOT}/external/libvpx/vpx/src/vpx_codec.c
        ${ANDROID_ROOT}/external/libvpx/vpx/src/vpx_decoder.c
        ${ANDROID_ROOT}/external/libvpx/vpx/src/vpx_encoder.c
        ${ANDROID_ROOT}/external/libvpx/vpx/src/vpx_image.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/avg_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/avg_pred_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/fdct4x4_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/fdct8x8_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/fdct16x16_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/fdct32x32_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/fdct_partial_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/hadamard_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_idct4x4_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_idct8x8_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_idct16x16_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_idct32x32_34_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_idct32x32_135_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_idct32x32_1024_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_idct32x32_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_intrapred_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_loopfilter_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_quantize_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_sad_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_variance_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_vpx_convolve8_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_vpx_convolve_avg_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_vpx_convolve_copy_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/highbd_vpx_convolve_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct4x4_1_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct4x4_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct8x8_1_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct8x8_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct16x16_1_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct16x16_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct32x32_1_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct32x32_34_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct32x32_135_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/idct32x32_add_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/intrapred_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/loopfilter_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/quantize_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/sad4d_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/sad_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/subpel_variance_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/subtract_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/sum_squares_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/variance_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/vpx_convolve8_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/vpx_convolve_avg_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/vpx_convolve_copy_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/vpx_convolve_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/arm/vpx_scaled_convolve8_neon.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/avg.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/bitreader.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/bitreader_buffer.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/bitwriter.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/bitwriter_buffer.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/fwd_txfm.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/intrapred.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/inv_txfm.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/loopfilter.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/prob.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/psnr.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/quantize.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/sad.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/skin_detection.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/subtract.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/sum_squares.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/variance.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/vpx_convolve.c
        ${ANDROID_ROOT}/external/libvpx/vpx_dsp/vpx_dsp_rtcd.c
        ${ANDROID_ROOT}/external/libvpx/vpx_mem/vpx_mem.c
        ${ANDROID_ROOT}/external/libvpx/vpx_ports/arm_cpudetect.c
        ${ANDROID_ROOT}/external/libvpx/vpx_scale/generic/gen_scalers.c
        ${ANDROID_ROOT}/external/libvpx/vpx_scale/generic/vpx_scale.c
        ${ANDROID_ROOT}/external/libvpx/vpx_scale/generic/yv12config.c
        ${ANDROID_ROOT}/external/libvpx/vpx_scale/generic/yv12extend.c
        ${ANDROID_ROOT}/external/libvpx/vpx_scale/vpx_scale_rtcd.c
        ${ANDROID_ROOT}/external/libvpx/vpx_util/vpx_thread.c
        ${ANDROID_ROOT}/external/libvpx/vpx_util/vpx_write_yuv_frame.c
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
    add_library(libvpx SHARED ${SOURCE_FILES})

    target_include_directories(libvpx
            PUBLIC
            ${ANDROID_ROOT}/external/libvpx
            ${ANDROID_ROOT}/external/libvpx/config/arm64
    )

    # Add c/c++ flags
    target_compile_options(libvpx PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libvpx
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libvpx ####################

#################### vp8_dec_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libvpx/examples/vpx_dec_fuzzer.cc
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
    add_executable(vp8_dec_fuzzer ${SOURCE_FILES})

    target_include_directories(vp8_dec_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libvpx/config/arm64
            ${ANDROID_ROOT}/external/libvpx
    )

    # Add c/c++ flags
    target_compile_options(vp8_dec_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(vp8_dec_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libvpx
    )
endif ()
#################### vp8_dec_fuzzer ####################

#################### vp9_dec_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libvpx/examples/vpx_dec_fuzzer.cc
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
    add_executable(vp9_dec_fuzzer ${SOURCE_FILES})

    target_include_directories(vp9_dec_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libvpx/config/arm64
            ${ANDROID_ROOT}/external/libvpx
    )

    # Add c/c++ flags
    target_compile_options(vp9_dec_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(vp9_dec_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libvpx
    )
endif ()
#################### vp9_dec_fuzzer ####################

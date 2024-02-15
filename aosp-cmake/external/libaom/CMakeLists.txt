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

project(av1_dec_fuzzer)

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

#################### libaom ####################
set(FILES
        ${ANDROID_ROOT}/external/libaom/av1/arg_defs.c
        ${ANDROID_ROOT}/external/libaom/av1/common/alloccommon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/av1_inv_txfm1d.c
        ${ANDROID_ROOT}/external/libaom/av1/common/av1_inv_txfm2d.c
        ${ANDROID_ROOT}/external/libaom/av1/common/av1_loopfilter.c
        ${ANDROID_ROOT}/external/libaom/av1/common/av1_txfm.c
        ${ANDROID_ROOT}/external/libaom/av1/common/blockd.c
        ${ANDROID_ROOT}/external/libaom/av1/common/cdef.c
        ${ANDROID_ROOT}/external/libaom/av1/common/cdef_block.c
        ${ANDROID_ROOT}/external/libaom/av1/common/cfl.c
        ${ANDROID_ROOT}/external/libaom/av1/common/common_data.c
        ${ANDROID_ROOT}/external/libaom/av1/common/convolve.c
        ${ANDROID_ROOT}/external/libaom/av1/common/debugmodes.c
        ${ANDROID_ROOT}/external/libaom/av1/common/entropy.c
        ${ANDROID_ROOT}/external/libaom/av1/common/entropymode.c
        ${ANDROID_ROOT}/external/libaom/av1/common/entropymv.c
        ${ANDROID_ROOT}/external/libaom/av1/common/frame_buffers.c
        ${ANDROID_ROOT}/external/libaom/av1/common/idct.c
        ${ANDROID_ROOT}/external/libaom/av1/common/mvref_common.c
        ${ANDROID_ROOT}/external/libaom/av1/common/obu_util.c
        ${ANDROID_ROOT}/external/libaom/av1/common/pred_common.c
        ${ANDROID_ROOT}/external/libaom/av1/common/quant_common.c
        ${ANDROID_ROOT}/external/libaom/av1/common/reconinter.c
        ${ANDROID_ROOT}/external/libaom/av1/common/reconintra.c
        ${ANDROID_ROOT}/external/libaom/av1/common/resize.c
        ${ANDROID_ROOT}/external/libaom/av1/common/restoration.c
        ${ANDROID_ROOT}/external/libaom/av1/common/scale.c
        ${ANDROID_ROOT}/external/libaom/av1/common/scan.c
        ${ANDROID_ROOT}/external/libaom/av1/common/seg_common.c
        ${ANDROID_ROOT}/external/libaom/av1/common/thread_common.c
        ${ANDROID_ROOT}/external/libaom/av1/common/tile_common.c
        ${ANDROID_ROOT}/external/libaom/av1/common/timing.c
        ${ANDROID_ROOT}/external/libaom/av1/common/txb_common.c
        ${ANDROID_ROOT}/external/libaom/av1/common/warped_motion.c
        ${ANDROID_ROOT}/external/libaom/common/args_helper.c
        ${ANDROID_ROOT}/external/libaom/common/av1_config.c
        ${ANDROID_ROOT}/external/libaom/av1/av1_dx_iface.c
        ${ANDROID_ROOT}/external/libaom/av1/decoder/decodeframe.c
        ${ANDROID_ROOT}/external/libaom/av1/decoder/decodemv.c
        ${ANDROID_ROOT}/external/libaom/av1/decoder/decoder.c
        ${ANDROID_ROOT}/external/libaom/av1/decoder/decodetxb.c
        ${ANDROID_ROOT}/external/libaom/av1/decoder/detokenize.c
        ${ANDROID_ROOT}/external/libaom/av1/decoder/grain_synthesis.c
        ${ANDROID_ROOT}/external/libaom/av1/decoder/obu.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/aom_convolve.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/bitreader_buffer.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/bitwriter_buffer.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/blend_a64_hmask.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/blend_a64_mask.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/blend_a64_vmask.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/entcode.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/fft.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/intrapred.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/loopfilter.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/odintrin.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/subtract.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/binary_codes_reader.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/bitreader.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/entdec.c
        ${ANDROID_ROOT}/external/libaom/aom_mem/aom_mem.c
        ${ANDROID_ROOT}/external/libaom/aom_ports/arm_cpudetect.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/aom_dsp_rtcd.c
        ${ANDROID_ROOT}/external/libaom/aom_scale/aom_scale_rtcd.c
        ${ANDROID_ROOT}/external/libaom/av1/common/av1_rtcd.c
        ${ANDROID_ROOT}/external/libaom/aom_scale/generic/aom_scale.c
        ${ANDROID_ROOT}/external/libaom/aom_scale/generic/gen_scalers.c
        ${ANDROID_ROOT}/external/libaom/aom_scale/generic/yv12config.c
        ${ANDROID_ROOT}/external/libaom/aom_scale/generic/yv12extend.c
        ${ANDROID_ROOT}/external/libaom/aom/src/aom_codec.c
        ${ANDROID_ROOT}/external/libaom/aom/src/aom_decoder.c
        ${ANDROID_ROOT}/external/libaom/aom/src/aom_encoder.c
        ${ANDROID_ROOT}/external/libaom/aom/src/aom_image.c
        ${ANDROID_ROOT}/external/libaom/aom/src/aom_integer.c
        ${ANDROID_ROOT}/external/libaom/aom_util/aom_thread.c
        ${ANDROID_ROOT}/external/libaom/aom_util/debug_util.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/av1_inv_txfm_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/av1_txfm_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/blend_a64_hmask_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/blend_a64_vmask_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/cdef_block_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/cfl_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/convolve_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/highbd_inv_txfm_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/jnt_convolve_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/reconinter_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/reconintra_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/resize_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/selfguided_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/warp_plane_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/common/arm/wiener_convolve_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/aom_convolve_copy_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/blend_a64_mask_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/fwd_txfm_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/highbd_intrapred_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/highbd_loopfilter_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/intrapred_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/loopfilter_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/subtract_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/av1_error_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/av1_fwd_txfm2d_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/av1_highbd_quantize_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/encodetxb_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/highbd_fwd_txfm_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/hybrid_fwd_txfm_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/ml_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/picksrt_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/quantize_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/rdopt_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/temporal_filter_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/arm/neon/wedge_utils_neon.c
        ${ANDROID_ROOT}/external/libaom/av1/av1_cx_iface.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/allintra_vis.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/aq_complexity.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/aq_cyclicrefresh.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/aq_variance.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/av1_fwd_txfm1d.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/av1_fwd_txfm2d.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/av1_noise_estimate.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/av1_quantize.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/bitstream.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/cnn.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/compound_type.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/context_tree.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/cost.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/dwt.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/encode_strategy.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/encodeframe.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/encodeframe_utils.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/encodemb.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/encodemv.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/encoder.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/encoder_utils.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/encodetxb.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/ethread.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/extend.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/external_partition.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/firstpass.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/global_motion.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/global_motion_facade.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/gop_structure.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/hash.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/hash_motion.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/hybrid_fwd_txfm.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/interp_search.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/intra_mode_search.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/level.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/lookahead.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/mcomp.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/ml.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/motion_search_facade.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/mv_prec.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/nonrd_pickmode.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/palette.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/partition_search.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/partition_strategy.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/pass2_strategy.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/pickcdef.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/picklpf.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/pickrst.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/ratectrl.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/rd.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/rdopt.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/reconinter_enc.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/segmentation.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/speed_features.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/superres_scale.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/svc_layercontext.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/temporal_filter.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/thirdpass.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/tokenize.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/tpl_model.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/tx_search.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/txb_rdopt.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/var_based_part.c
        ${ANDROID_ROOT}/external/libaom/av1/encoder/wedge_utils.c
        ${ANDROID_ROOT}/external/libaom/third_party/fastfeat/fast.c
        ${ANDROID_ROOT}/external/libaom/third_party/fastfeat/fast_9.c
        ${ANDROID_ROOT}/external/libaom/third_party/fastfeat/nonmax.c
        ${ANDROID_ROOT}/external/libaom/third_party/vector/vector.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/avg_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/hadamard_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/highbd_quantize_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/highbd_variance_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/sad4d_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/sad_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/sse_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/subpel_variance_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/sum_squares_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/arm/variance_neon.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/avg.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/binary_codes_writer.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/bitwriter.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/blk_sse_sum.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/entenc.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/flow_estimation/corner_detect.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/flow_estimation/corner_match.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/flow_estimation/disflow.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/flow_estimation/flow_estimation.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/flow_estimation/ransac.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/fwd_txfm.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/grain_table.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/noise_model.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/noise_util.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/psnr.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/quantize.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/sad.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/sad_av1.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/sse.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/ssim.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/sum_squares.c
        ${ANDROID_ROOT}/external/libaom/aom_dsp/variance.c
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
    add_library(libaom SHARED ${SOURCE_FILES})

    target_include_directories(libaom
            PUBLIC
            ${ANDROID_ROOT}/external/libaom
            ${ANDROID_ROOT}/external/libaom/config
            ${ANDROID_ROOT}/external/libaom/config/arm64
    )

    # Add c/c++ flags
    target_compile_options(libaom PRIVATE
            -Werror
    )

    target_link_libraries(libaom
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libaom ####################

#################### av1_dec_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libaom/examples/av1_dec_fuzzer.cc
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
    add_executable(av1_dec_fuzzer ${SOURCE_FILES})

    target_include_directories(av1_dec_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libaom/config
            ${ANDROID_ROOT}/external/libaom/config/arm64
            ${ANDROID_ROOT}/external/libaom
    )

    # Add c/c++ flags
    target_compile_options(av1_dec_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(av1_dec_fuzzer
            libaom
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
    )
endif ()
#################### av1_dec_fuzzer ####################

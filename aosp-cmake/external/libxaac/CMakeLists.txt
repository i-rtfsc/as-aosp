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

project(libxaacdec)

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

#################### libxaacdec ####################
set(FILES
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_aacdecoder.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_aacpluscheck.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_aac_ec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_aac_imdct.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_aac_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_aac_tns.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_acelp_bitparse.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_acelp_decode.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_acelp_mdct.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_acelp_tools.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_adts_crc_check.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_api.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_arith_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_avq_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_avq_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_basic_funcs.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_basic_ops.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_bitbuffer.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_block.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_channel.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_common_initfuncs.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_common_lpfuncs.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_common_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_create.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_decode_main.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_drc_freq_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_dsp_fft32x32s.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_ec_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_env_calc.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_env_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_env_extr.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_esbr_envcal.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_esbr_fft.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_esbr_polyphase.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_esbr_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_ext_ch_ele.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_fft.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_fft_ifft_32x32.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_fft_ifft_32x32_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_freq_sca.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_fwd_alias_cnx.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_hbe_dft_trans.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_hbe_trans.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_headerdecode.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_hufftables.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_huff_code_reorder.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_huff_tools.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_hybrid.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_imdct.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_initfuncs.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_init_config.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_latmdemux.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_ld_mps_config.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_ld_mps_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_longblock.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_lpc.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_lpc_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_lpfuncs.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_lpp_tran.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_lt_predict.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_apply_common.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_apply_m1.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_apply_m2.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_bitdec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_blind.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_calc_m1m2_common.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_calc_m1m2_emm.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_calc_m1m2_tree_515x.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_calc_m1m2_tree_51sx.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_calc_m1m2_tree_52xx.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_calc_m1m2_tree_727x.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_calc_m1m2_tree_757x.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_decorr.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_get_index.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_hybrid_filt.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_initfuncs.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_m1m2_common.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_mdct_2_qmf.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_parse.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_polyphase.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_poly_filt.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_pre_mix.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_process.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_reshape_bb_env.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_res_block.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_res_channel.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_res_channel_info.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_res_longblock.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_res_pns_js_thumb.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_res_pulsedata.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_res_tns.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_smoothing.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_temp_process.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_temp_reshape.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_mps_tonality.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_multichannel.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_peak_limiter.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_pns_js_thumb.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_pred_vec_block.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_process.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_ps_bitdec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_ps_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_ps_dec_flt.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_pvc_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_qmf_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_rev_vlc.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_sbrdecoder.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_sbrdec_initfuncs.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_sbrdec_lpfuncs.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_sbr_crc.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_sbr_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_sbr_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_spectrum_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_stereo.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_tcx_fwd_alcnx.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_tcx_fwd_mdct.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_thumb_ps_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_tns.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_usac_ec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/ixheaacd_Windowing.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_api.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_bitbuffer.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_dynamic_payload.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_eq.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_extr_delta_coded_info.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_filter_bank.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_gain_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_gain_decoder.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_init.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_interface_decoder.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_loudness_control.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_main_td_process.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_multiband.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_parametric_dec.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_peak_limiter.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_process.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_rom.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_selection_process.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_selection_process_drcset_selection.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_selection_process_init.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_shape_filter.c
        ${ANDROID_ROOT}/external/libxaac/decoder/drc_src/impd_drc_static_payload.c
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_function_selector_armv8.c
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_qmf_dec_armv8.c
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_apply_scale_factors.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_calcmaxspectralline.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_cos_sin_mod_loop1.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_cos_sin_mod_loop2.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_fft32x32_ld2_armv8.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_imdct_using_fft.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_inv_dit_fft_8pt.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_no_lap1.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_overlap_add1.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_overlap_add2.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_postradixcompute4.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_post_twiddle.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_post_twiddle_overlap.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_pre_twiddle.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_sbr_imdct_using_fft.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_sbr_qmfsyn64_winadd.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_sbr_qmf_analysis32_neon.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_shiftrountine_with_round.s
        ${ANDROID_ROOT}/external/libxaac/decoder/armv8/ixheaacd_shiftrountine_with_round_eld.s
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
    add_library(libxaacdec SHARED ${SOURCE_FILES})

    target_include_directories(libxaacdec
            PUBLIC
            ${ANDROID_ROOT}/external/libxaac/decoder
            ${ANDROID_ROOT}/external/libxaac/decoder/drc_src
            ${ANDROID_ROOT}/external/libxaac/decoder/armv8
            ${ANDROID_ROOT}/external/libxaac
    )

    # Add c/c++ flags
    target_compile_options(libxaacdec PRIVATE
            -Werror
    )

    target_link_libraries(libxaacdec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libxaacdec ####################

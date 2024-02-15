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

project(libhevcdec)

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

#################### libhevcdec ####################
set(FILES
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_quant_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_inter_pred_filters.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_weighted_pred.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_padding.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_deblk_edge_filter.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_deblk_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_cabac_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_common_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_intra_pred_filters.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_intra_pred_filters.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_mem_fns.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_sao.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_trans_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_itrans.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_itrans_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_iquant_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_iquant_itrans_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_itrans_recon_32x32.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_itrans_recon_16x16.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_itrans_recon_8x8.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_itrans_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_iquant_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_iquant_itrans_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_itrans_recon_16x16.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_itrans_recon_8x8.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_buf_mgr.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_disp_mgr.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_dpb_mgr.c
        ${ANDROID_ROOT}/external/libhevc/common/ithread.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_version.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_api.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_decode.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_nal.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_bitstream.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_parse_headers.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_parse_slice_header.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_parse_slice.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_parse_residual.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_cabac.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_intra_pred_mode_prediction.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_process_slice.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_utils.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_job_queue.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_ref_list.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_get_mv.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_mv_pred.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_mv_merge.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_iquant_itrans_recon_ctb.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_itrans_recon_dc.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_common_tables.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_boundary_strength.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_deblk.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_inter_pred.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_sao.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_ilf_padding.c
        ${ANDROID_ROOT}/external/libhevc/decoder/ihevcd_fmt_conv.c
        ${ANDROID_ROOT}/external/libhevc/decoder/arm/ihevcd_function_selector.c
        ${ANDROID_ROOT}/external/libhevc/decoder/arm/ihevcd_function_selector_noneon.c
        ${ANDROID_ROOT}/external/libhevc/decoder/arm64/ihevcd_function_selector_av8.c
        ${ANDROID_ROOT}/external/libhevc/common/arm/ihevc_intra_pred_filters_neon_intr.c
        ${ANDROID_ROOT}/external/libhevc/common/arm/ihevc_weighted_pred_neon_intr.c
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_mem_fns.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_32x32.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_weighted_pred_bi_default.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_weighted_pred_bi.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_weighted_pred_uni.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_deblk_luma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_deblk_luma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_deblk_chroma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_deblk_chroma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_band_offset_luma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_band_offset_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class0.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class0_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class1.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class1_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class2.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class2_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class3.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class3_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_luma_horz_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_filters_luma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_filters_luma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_horz_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_vert_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_vert_w16inp.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_vert_w16inp_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_filters_luma_vert_w16inp.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_filters_luma_vert_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_luma_vert_w16inp_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_luma_copy_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_luma_copy.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_copy.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_copy_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_4x4_ttype1.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_4x4.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_8x8.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_16x16.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_planar.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_dc.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_ver.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_mode2.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_mode_18_34.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_filters_chroma_mode_11_to_17.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_filters_chroma_mode_19_to_25.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_mode_3_to_9.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_mode_27_to_33.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_planar.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_mode2.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_mode_27_to_33.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_mode_18_34.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_dc.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_filters_luma_mode_11_to_17.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_filters_luma_mode_19_to_25.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_mode_3_to_9.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_padding.s
        ${ANDROID_ROOT}/external/libhevc/decoder/arm64/ihevcd_itrans_recon_dc_luma.s
        ${ANDROID_ROOT}/external/libhevc/decoder/arm64/ihevcd_itrans_recon_dc_chroma.s
        ${ANDROID_ROOT}/external/libhevc/decoder/arm64/ihevcd_fmt_conv_420sp_to_420p.s
        ${ANDROID_ROOT}/external/libhevc/decoder/arm64/ihevcd_fmt_conv_420sp_to_420sp.s
        ${ANDROID_ROOT}/external/libhevc/decoder/arm64/ihevcd_fmt_conv_420sp_to_rgba8888.s
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
    add_library(libhevcdec SHARED ${SOURCE_FILES})

    target_include_directories(libhevcdec
            PUBLIC
            ${ANDROID_ROOT}/external/libhevc/common
            ${ANDROID_ROOT}/external/libhevc/decoder
            ${ANDROID_ROOT}/external/libhevc/decoder/arm
            ${ANDROID_ROOT}/external/libhevc/common/arm
            ${ANDROID_ROOT}/external/libhevc/decoder/arm64
            ${ANDROID_ROOT}/external/libhevc/common/arm64
            ${ANDROID_ROOT}/external/libhevc
    )

    # Add c/c++ flags
    target_compile_options(libhevcdec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libhevcdec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libhevcdec ####################

#################### libhevcenc ####################
set(FILES
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_cabac_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_intra_pred_filters.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_itrans_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_itrans_recon_16x16.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_chroma_itrans_recon_8x8.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_common_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_deblk_edge_filter.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_deblk_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_hbd_deblk_edge_filter.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_inter_pred_filters.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_intra_pred_filters.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_iquant_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_itrans_recon.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_itrans_recon_16x16.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_itrans_recon_32x32.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_itrans_recon_8x8.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_mem_fns.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_padding.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_quant_iquant_ssd.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_quant_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_resi_trans.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_sao.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_trans_tables.c
        ${ANDROID_ROOT}/external/libhevc/common/ihevc_weighted_pred.c
        ${ANDROID_ROOT}/external/libhevc/encoder/bit_allocation.c
        ${ANDROID_ROOT}/external/libhevc/encoder/cbr_buffer_control.c
        ${ANDROID_ROOT}/external/libhevc/encoder/common_rom.c
        ${ANDROID_ROOT}/external/libhevc/encoder/convert_float_to_fix.c
        ${ANDROID_ROOT}/external/libhevc/encoder/est_sad.c
        ${ANDROID_ROOT}/external/libhevc/encoder/fixed_point_error_bits.c
        ${ANDROID_ROOT}/external/libhevc/encoder/frame_info_collector.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_coarse.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_common_utils.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_err_compute.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_fullpel.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_function_selector.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_globals.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_interface.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_refine.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_search_algo.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_subpel.c
        ${ANDROID_ROOT}/external/libhevc/encoder/hme_utils.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_bitstream.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_bs_compute_ctb.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_buffer_que.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_cabac.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_cabac_cu_pu.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_cabac_rdo.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_cabac_tu.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_chroma_had_satd.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_cmn_utils_instr_set_router.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_coarse_me_pass.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_common_utils.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_deblk.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_decomp_pre_intra_pass.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_dep_mngr.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_enc_cu_recursion.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_enc_loop_inter_mode_sifter.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_enc_loop_pass.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_enc_loop_utils.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_enc_sbh_funcs.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_enc_subpel_gen.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_encode_header.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_encode_header_sei_vui.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_entropy_cod.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_entropy_interface.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_error_check.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_frame_process.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_frame_process_utils.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_function_selector.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_global_tables.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_had_satd.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_hle_interface.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_hle_q_func.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_inter_pred.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_ipe_instr_set_router.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_ipe_pass.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_lap_interface.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_me_instr_set_router.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_me_pass.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_memory_init.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_multi_thrd_funcs.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_mv_pred.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_mv_pred_merge.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_nbr_avail.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_plugin.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_profile.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_rc_interface.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_recur_bracketing.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_sao.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_stasino_helpers.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_sub_pic_rc.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_sys_api.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_tile_interface.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_trace.c
        ${ANDROID_ROOT}/external/libhevc/encoder/ihevce_tu_tree_selector.c
        ${ANDROID_ROOT}/external/libhevc/encoder/init_qp.c
        ${ANDROID_ROOT}/external/libhevc/encoder/mb_model_based.c
        ${ANDROID_ROOT}/external/libhevc/encoder/osal.c
        ${ANDROID_ROOT}/external/libhevc/encoder/osal_cond_var.c
        ${ANDROID_ROOT}/external/libhevc/encoder/osal_error.c
        ${ANDROID_ROOT}/external/libhevc/encoder/osal_mutex.c
        ${ANDROID_ROOT}/external/libhevc/encoder/osal_semaphore.c
        ${ANDROID_ROOT}/external/libhevc/encoder/osal_thread.c
        ${ANDROID_ROOT}/external/libhevc/encoder/picture_type.c
        ${ANDROID_ROOT}/external/libhevc/encoder/rate_control_api.c
        ${ANDROID_ROOT}/external/libhevc/encoder/rc_rd_model.c
        ${ANDROID_ROOT}/external/libhevc/encoder/rc_rd_model_fix.c
        ${ANDROID_ROOT}/external/libhevc/encoder/rc_sad_acc.c
        ${ANDROID_ROOT}/external/libhevc/encoder/sqrt_interp.c
        ${ANDROID_ROOT}/external/libhevc/encoder/var_q_operator.c
        ${ANDROID_ROOT}/external/libhevc/encoder/vbr_storage_vbv.c
        ${ANDROID_ROOT}/external/libhevc/encoder/vbr_str_prms.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_coarse_layer_sad_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_common_utils_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_copy_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_had_compute_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_hme_utils_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_itrans_recon_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_me_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_sad_compute_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_scale_by_2_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_scan_coeffs_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_ssd_and_sad_calculator_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_ssd_calculator_neon.c
        ${ANDROID_ROOT}/external/libhevc/encoder/arm/ihevce_subpel_neon.c
        ${ANDROID_ROOT}/external/libhevc/common/arm/ihevc_resi_trans_neon.c
        ${ANDROID_ROOT}/external/libhevc/common/arm/ihevc_resi_trans_neon_32x32.c
        ${ANDROID_ROOT}/external/libhevc/common/arm/ihevc_quant_iquant_ssd_neon_intr.c
        ${ANDROID_ROOT}/external/libhevc/common/arm/ihevc_intra_pred_filters_neon_intr.c
        ${ANDROID_ROOT}/external/libhevc/common/arm/ihevc_weighted_pred_neon_intr.c
        ${ANDROID_ROOT}/external/libhevc/common/arm/ihevc_intra_ref_substitution_a9q.c
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_deblk_chroma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_deblk_chroma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_deblk_luma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_deblk_luma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_copy.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_copy_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_horz_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_vert_w16inp.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_vert_w16inp_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_chroma_vert_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_filters_luma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_filters_luma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_filters_luma_vert_w16inp.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_filters_luma_vert_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_luma_copy.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_luma_copy_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_luma_horz_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_inter_pred_luma_vert_w16inp_w16out.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_dc.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_mode2.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_mode_18_34.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_mode_27_to_33.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_mode_3_to_9.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_planar.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_chroma_ver.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_filters_chroma_mode_11_to_17.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_filters_chroma_mode_19_to_25.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_filters_luma_mode_11_to_17.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_filters_luma_mode_19_to_25.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_dc.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_horz.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_mode2.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_mode_18_34.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_mode_27_to_33.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_mode_3_to_9.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_planar.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_intra_pred_luma_vert.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_16x16.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_32x32.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_4x4.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_4x4_ttype1.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_itrans_recon_8x8.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_mem_fns.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_padding.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_band_offset_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_band_offset_luma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class0.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class0_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class1.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class1_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class2.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class2_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class3.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_sao_edge_offset_class3_chroma.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_weighted_pred_bi.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_weighted_pred_bi_default.s
        ${ANDROID_ROOT}/external/libhevc/common/arm64/ihevc_weighted_pred_uni.s
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
    add_library(libhevcenc SHARED ${SOURCE_FILES})

    target_include_directories(libhevcenc
            PUBLIC
            ${ANDROID_ROOT}/external/libhevc/common
            ${ANDROID_ROOT}/external/libhevc/encoder
            ${ANDROID_ROOT}/external/libhevc/encoder/arm
            ${ANDROID_ROOT}/external/libhevc/common/arm
            ${ANDROID_ROOT}/external/libhevc/common/arm64
            ${ANDROID_ROOT}/external/libhevc
    )

    # Add c/c++ flags
    target_compile_options(libhevcenc PRIVATE
            -Werror
            -Wall
            -Wno-unused-variable
            -Wno-unused-parameter
    )

    target_link_libraries(libhevcenc
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libhevcenc ####################

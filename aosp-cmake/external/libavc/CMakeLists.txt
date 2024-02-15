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

project(libmvcdec)

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

#################### libmvcdec ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/decoder/mvc/imvcd_api.c
        ${ANDROID_ROOT}/external/libavc/decoder/mvc/imvcd_api_utils.c
        ${ANDROID_ROOT}/external/libavc/decoder/mvc/imvcd_dpb_manager.c
        ${ANDROID_ROOT}/external/libavc/decoder/mvc/imvcd_error_handler.c
        ${ANDROID_ROOT}/external/libavc/decoder/mvc/imvcd_nalu_parser.c
        ${ANDROID_ROOT}/external/libavc/decoder/mvc/imvcd_slice_functions.c
        ${ANDROID_ROOT}/external/libavc/decoder/mvc/imvcd_utils.c
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
    add_library(libmvcdec SHARED ${SOURCE_FILES})

    target_include_directories(libmvcdec
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/decoder
            ${ANDROID_ROOT}/external/libavc/decoder/mvc
            ${ANDROID_ROOT}/external/libavc/common
            ${ANDROID_ROOT}/external/libavc/common/mvc
            ${ANDROID_ROOT}/external/libavc/decoder/arm
            ${ANDROID_ROOT}/external/libavc/common/armv8
            ${ANDROID_ROOT}/external/libavc
    )

    # Add c/c++ flags
    target_compile_options(libmvcdec PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libmvcdec
            libavcdec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libmvcdec ####################

#################### libavcdec ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/common/ih264_buf_mgr.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_disp_mgr.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_inter_pred_filters.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_luma_intra_pred_filters.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_chroma_intra_pred_filters.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_padding.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_mem_fns.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_deblk_edge_filters.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_iquant_itrans_recon.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_ihadamard_scaling.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_weighted_pred.c
        ${ANDROID_ROOT}/external/libavc/common/ithread.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_cabac.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_parse_mb_header.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_parse_cabac.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_process_intra_mb.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_inter_pred.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_parse_bslice.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_parse_pslice.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_parse_islice.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_cabac_init_tables.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_bitstrm.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_compute_bs.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_deblocking.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_parse_headers.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_mb_utils.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_mvpred.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_utils.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_process_bslice.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_process_pslice.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_parse_slice.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_quant_scaling.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_parse_cavlc.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_dpb_mgr.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_nal.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_sei.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_tables.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_vui.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_format_conv.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_thread_parse_decode.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_api.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_thread_compute_bs.c
        ${ANDROID_ROOT}/external/libavc/decoder/ih264d_function_selector_generic.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/ih264d_function_selector.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/ih264d_function_selector_av8.c
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_intra_pred_chroma_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_intra_pred_luma_16x16_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_intra_pred_luma_4x4_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_chroma_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_filters_luma_horz_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_filters_luma_vert_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_luma_copy_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_luma_horz_qpel_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_luma_vert_qpel_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_luma_horz_hpel_vert_hpel_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_luma_horz_qpel_vert_qpel_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_luma_horz_qpel_vert_hpel_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_luma_horz_hpel_vert_qpel_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_default_weighted_pred_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_weighted_pred_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_weighted_bi_pred_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_deblk_chroma_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_deblk_luma_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_padding_neon_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_iquant_itrans_recon_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_iquant_itrans_recon_dc_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_ihadamard_scaling_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_intra_pred_luma_8x8_av8.s
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
    add_library(libavcdec SHARED ${SOURCE_FILES})

    target_include_directories(libavcdec
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/common
            ${ANDROID_ROOT}/external/libavc/decoder
            ${ANDROID_ROOT}/external/libavc/decoder/arm
            ${ANDROID_ROOT}/external/libavc/common/armv8
            ${ANDROID_ROOT}/external/libavc
    )

    # Add c/c++ flags
    target_compile_options(libavcdec PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libavcdec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libavcdec ####################

#################### libavcenc ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/common/ih264_resi_trans_quant.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_iquant_itrans_recon.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_ihadamard_scaling.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_inter_pred_filters.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_luma_intra_pred_filters.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_chroma_intra_pred_filters.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_padding.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_mem_fns.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_deblk_edge_filters.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_deblk_tables.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_cavlc_tables.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_cabac_tables.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_common_tables.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_trans_data.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_buf_mgr.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_dpb_mgr.c
        ${ANDROID_ROOT}/external/libavc/common/ih264_list.c
        ${ANDROID_ROOT}/external/libavc/common/ithread.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_globals.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_intra_modes_eval.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_half_pel.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_mc.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_me.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_rc_mem_interface.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_time_stamp.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_modify_frm_rate.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_rate_control.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_core_coding.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_deblk.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_api.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_process.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_encode.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_utils.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_version.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_bitstream.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_cavlc.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_cabac_init.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_cabac.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_cabac_encode.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_encode_header.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_function_selector_generic.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_fmt_conv.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_rate_control_api.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_bit_allocation.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_cbr_buffer_control.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_est_sad.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_fixed_point_error_bits.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_frame_info_collector.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_mb_model_based.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_picture_type.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_rd_model.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_vbr_storage_vbv.c
        ${ANDROID_ROOT}/external/libavc/encoder/irc_vbr_str_prms.c
        ${ANDROID_ROOT}/external/libavc/encoder/ime.c
        ${ANDROID_ROOT}/external/libavc/encoder/ime_distortion_metrics.c
        ${ANDROID_ROOT}/external/libavc/encoder/ih264e_sei.c
        ${ANDROID_ROOT}/external/libavc/encoder/psnr.c
        ${ANDROID_ROOT}/external/libavc/encoder/arm/ih264e_function_selector.c
        ${ANDROID_ROOT}/external/libavc/encoder/arm/ih264e_function_selector_av8.c
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_resi_trans_quant_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_iquant_itrans_recon_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_iquant_itrans_recon_dc_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_ihadamard_scaling_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_intra_pred_chroma_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_intra_pred_luma_16x16_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_intra_pred_luma_4x4_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_intra_pred_luma_8x8_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_luma_copy_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_chroma_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_filters_luma_horz_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_inter_pred_filters_luma_vert_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_padding_neon_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_mem_fns_neon_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_deblk_luma_av8.s
        ${ANDROID_ROOT}/external/libavc/common/armv8/ih264_deblk_chroma_av8.s
        ${ANDROID_ROOT}/external/libavc/encoder/armv8/ih264e_evaluate_intra16x16_modes_av8.s
        ${ANDROID_ROOT}/external/libavc/encoder/armv8/ih264e_evaluate_intra_chroma_modes_av8.s
        ${ANDROID_ROOT}/external/libavc/encoder/armv8/ih264e_half_pel_av8.s
        ${ANDROID_ROOT}/external/libavc/encoder/armv8/ime_distortion_metrics_av8.s
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
    add_library(libavcenc SHARED ${SOURCE_FILES})

    target_include_directories(libavcenc
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/common
            ${ANDROID_ROOT}/external/libavc/encoder
            ${ANDROID_ROOT}/external/libavc/common/arm
            ${ANDROID_ROOT}/external/libavc/common/armv8
            ${ANDROID_ROOT}/external/libavc/encoder/arm
            ${ANDROID_ROOT}/external/libavc/encoder/armv8
            ${ANDROID_ROOT}/external/libavc
    )

    # Add c/c++ flags
    target_compile_options(libavcenc PRIVATE
            -DNDEBUG
            -Wall
            -Werror
    )

    target_link_libraries(libavcenc
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libavcenc ####################

#################### libsvcdec ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_ii_pred.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_intra_resamp.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_iquant_itrans.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_iquant_itrans_residual.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_iquant_itrans_residual_recon.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_mode_mv_resamp.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_pred_residual_recon.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_residual_resamp.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_api.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_cabac.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_cabac_init_tables.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_compute_bs.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_function_selector_generic.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_mb_utils.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_nal.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_nal_parse.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_parse_cavlc.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_parse_ebslice.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_parse_eislice.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_parse_epslice.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_parse_headers.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_parse_slice.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_process_ebslice.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_process_epslice.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_thread_compute_bs.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_thread_parse_decode.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_utils.c
        ${ANDROID_ROOT}/external/libavc/decoder/svc/isvcd_vui.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/svc/isvcd_function_selector.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/svc/isvcd_function_selector_neon.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/svc/isvcd_intra_resamp_neon.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/svc/isvcd_iquant_itrans_neon.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/svc/isvcd_iquant_itrans_residual_neon.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/svc/isvcd_iquant_itrans_residual_recon_neon.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/svc/isvcd_pred_residual_recon_neon.c
        ${ANDROID_ROOT}/external/libavc/decoder/arm/svc/isvcd_residual_resamp_neon.c
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
    add_library(libsvcdec SHARED ${SOURCE_FILES})

    target_include_directories(libsvcdec
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/decoder
            ${ANDROID_ROOT}/external/libavc/decoder/svc
            ${ANDROID_ROOT}/external/libavc/common
            ${ANDROID_ROOT}/external/libavc/common/svc
            ${ANDROID_ROOT}/external/libavc/decoder/arm
            ${ANDROID_ROOT}/external/libavc/common/armv8
            ${ANDROID_ROOT}/external/libavc
    )

    # Add c/c++ flags
    target_compile_options(libsvcdec PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(libsvcdec
            libavcdec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libsvcdec ####################

#################### libsvcenc ####################
set(FILES
        ${ANDROID_ROOT}/external/libavc/common/svc/isvc_cabac_tables.c
        ${ANDROID_ROOT}/external/libavc/common/svc/isvc_common_tables.c
        ${ANDROID_ROOT}/external/libavc/common/svc/isvc_intra_resample.c
        ${ANDROID_ROOT}/external/libavc/common/svc/isvc_iquant_itrans_recon.c
        ${ANDROID_ROOT}/external/libavc/common/svc/isvc_mem_fns.c
        ${ANDROID_ROOT}/external/libavc/common/svc/isvc_resi_trans_quant.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/irc_svc_rate_control_api.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_api.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_cabac.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_cabac_encode.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_cabac_init.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_cavlc.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_core_coding.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_deblk.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_downscaler.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_encode.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_encode_header.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_fmt_conv.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_function_selector_generic.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_globals.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_ibl_eval.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_ilp_mv.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_intra_modes_eval.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_mc.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_me.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_mode_stat_visualiser.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_nalu_stat_aggregator.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_process.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_rate_control.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_rc_mem_interface.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_rc_utils.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_residual_pred.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_sub_pic_rc.c
        ${ANDROID_ROOT}/external/libavc/encoder/svc/isvce_utils.c
        ${ANDROID_ROOT}/external/libavc/encoder/arm/svc/isvce_function_selector.c
        ${ANDROID_ROOT}/external/libavc/encoder/arm/svc/isvce_function_selector_av8.c
        ${ANDROID_ROOT}/external/libavc/common/arm/svc/isvc_intra_sampling_neon.c
        ${ANDROID_ROOT}/external/libavc/common/arm/svc/isvc_iquant_itrans_recon_neon.c
        ${ANDROID_ROOT}/external/libavc/common/arm/svc/isvc_mem_fns_neon.c
        ${ANDROID_ROOT}/external/libavc/common/arm/svc/isvc_resi_trans_quant_neon.c
        ${ANDROID_ROOT}/external/libavc/encoder/arm/svc/isvce_downscaler_neon.c
        ${ANDROID_ROOT}/external/libavc/encoder/arm/svc/isvce_rc_utils_neon.c
        ${ANDROID_ROOT}/external/libavc/encoder/arm/svc/isvce_residual_pred_neon.c
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
    add_library(libsvcenc SHARED ${SOURCE_FILES})

    target_include_directories(libsvcenc
            PUBLIC
            ${ANDROID_ROOT}/external/libavc/common
            ${ANDROID_ROOT}/external/libavc/common/svc
            ${ANDROID_ROOT}/external/libavc/encoder
            ${ANDROID_ROOT}/external/libavc/encoder/svc
            ${ANDROID_ROOT}/external/libavc/common/arm
            ${ANDROID_ROOT}/external/libavc/common/armv8
            ${ANDROID_ROOT}/external/libavc/encoder/arm
            ${ANDROID_ROOT}/external/libavc/encoder/armv8
            ${ANDROID_ROOT}/external/libavc/common/arm/svc
            ${ANDROID_ROOT}/external/libavc/encoder/arm/svc
            ${ANDROID_ROOT}/external/libavc
    )

    # Add c/c++ flags
    target_compile_options(libsvcenc PRIVATE
            -DNDEBUG
            -Wall
            -Werror
    )

    target_link_libraries(libsvcenc
            libavcenc
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libsvcenc ####################

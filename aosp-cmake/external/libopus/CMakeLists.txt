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

project(libopus)

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

#################### libopus ####################
set(FILES
        ${ANDROID_ROOT}/external/libopus/celt/bands.c
        ${ANDROID_ROOT}/external/libopus/celt/celt.c
        ${ANDROID_ROOT}/external/libopus/celt/celt_encoder.c
        ${ANDROID_ROOT}/external/libopus/celt/celt_decoder.c
        ${ANDROID_ROOT}/external/libopus/celt/cwrs.c
        ${ANDROID_ROOT}/external/libopus/celt/entcode.c
        ${ANDROID_ROOT}/external/libopus/celt/entdec.c
        ${ANDROID_ROOT}/external/libopus/celt/entenc.c
        ${ANDROID_ROOT}/external/libopus/celt/kiss_fft.c
        ${ANDROID_ROOT}/external/libopus/celt/laplace.c
        ${ANDROID_ROOT}/external/libopus/celt/mathops.c
        ${ANDROID_ROOT}/external/libopus/celt/mdct.c
        ${ANDROID_ROOT}/external/libopus/celt/modes.c
        ${ANDROID_ROOT}/external/libopus/celt/pitch.c
        ${ANDROID_ROOT}/external/libopus/celt/celt_lpc.c
        ${ANDROID_ROOT}/external/libopus/celt/quant_bands.c
        ${ANDROID_ROOT}/external/libopus/celt/rate.c
        ${ANDROID_ROOT}/external/libopus/celt/vq.c
        ${ANDROID_ROOT}/external/libopus/silk/CNG.c
        ${ANDROID_ROOT}/external/libopus/silk/code_signs.c
        ${ANDROID_ROOT}/external/libopus/silk/init_decoder.c
        ${ANDROID_ROOT}/external/libopus/silk/decode_core.c
        ${ANDROID_ROOT}/external/libopus/silk/decode_frame.c
        ${ANDROID_ROOT}/external/libopus/silk/decode_parameters.c
        ${ANDROID_ROOT}/external/libopus/silk/decode_indices.c
        ${ANDROID_ROOT}/external/libopus/silk/decode_pulses.c
        ${ANDROID_ROOT}/external/libopus/silk/decoder_set_fs.c
        ${ANDROID_ROOT}/external/libopus/silk/dec_API.c
        ${ANDROID_ROOT}/external/libopus/silk/enc_API.c
        ${ANDROID_ROOT}/external/libopus/silk/encode_indices.c
        ${ANDROID_ROOT}/external/libopus/silk/encode_pulses.c
        ${ANDROID_ROOT}/external/libopus/silk/gain_quant.c
        ${ANDROID_ROOT}/external/libopus/silk/interpolate.c
        ${ANDROID_ROOT}/external/libopus/silk/LP_variable_cutoff.c
        ${ANDROID_ROOT}/external/libopus/silk/NLSF_decode.c
        ${ANDROID_ROOT}/external/libopus/silk/NSQ.c
        ${ANDROID_ROOT}/external/libopus/silk/NSQ_del_dec.c
        ${ANDROID_ROOT}/external/libopus/silk/PLC.c
        ${ANDROID_ROOT}/external/libopus/silk/shell_coder.c
        ${ANDROID_ROOT}/external/libopus/silk/tables_gain.c
        ${ANDROID_ROOT}/external/libopus/silk/tables_LTP.c
        ${ANDROID_ROOT}/external/libopus/silk/tables_NLSF_CB_NB_MB.c
        ${ANDROID_ROOT}/external/libopus/silk/tables_NLSF_CB_WB.c
        ${ANDROID_ROOT}/external/libopus/silk/tables_other.c
        ${ANDROID_ROOT}/external/libopus/silk/tables_pitch_lag.c
        ${ANDROID_ROOT}/external/libopus/silk/tables_pulses_per_block.c
        ${ANDROID_ROOT}/external/libopus/silk/VAD.c
        ${ANDROID_ROOT}/external/libopus/silk/control_audio_bandwidth.c
        ${ANDROID_ROOT}/external/libopus/silk/quant_LTP_gains.c
        ${ANDROID_ROOT}/external/libopus/silk/VQ_WMat_EC.c
        ${ANDROID_ROOT}/external/libopus/silk/HP_variable_cutoff.c
        ${ANDROID_ROOT}/external/libopus/silk/NLSF_encode.c
        ${ANDROID_ROOT}/external/libopus/silk/NLSF_VQ.c
        ${ANDROID_ROOT}/external/libopus/silk/NLSF_unpack.c
        ${ANDROID_ROOT}/external/libopus/silk/NLSF_del_dec_quant.c
        ${ANDROID_ROOT}/external/libopus/silk/process_NLSFs.c
        ${ANDROID_ROOT}/external/libopus/silk/stereo_LR_to_MS.c
        ${ANDROID_ROOT}/external/libopus/silk/stereo_MS_to_LR.c
        ${ANDROID_ROOT}/external/libopus/silk/check_control_input.c
        ${ANDROID_ROOT}/external/libopus/silk/control_SNR.c
        ${ANDROID_ROOT}/external/libopus/silk/init_encoder.c
        ${ANDROID_ROOT}/external/libopus/silk/control_codec.c
        ${ANDROID_ROOT}/external/libopus/silk/A2NLSF.c
        ${ANDROID_ROOT}/external/libopus/silk/ana_filt_bank_1.c
        ${ANDROID_ROOT}/external/libopus/silk/biquad_alt.c
        ${ANDROID_ROOT}/external/libopus/silk/bwexpander_32.c
        ${ANDROID_ROOT}/external/libopus/silk/bwexpander.c
        ${ANDROID_ROOT}/external/libopus/silk/debug.c
        ${ANDROID_ROOT}/external/libopus/silk/decode_pitch.c
        ${ANDROID_ROOT}/external/libopus/silk/inner_prod_aligned.c
        ${ANDROID_ROOT}/external/libopus/silk/lin2log.c
        ${ANDROID_ROOT}/external/libopus/silk/log2lin.c
        ${ANDROID_ROOT}/external/libopus/silk/LPC_analysis_filter.c
        ${ANDROID_ROOT}/external/libopus/silk/LPC_fit.c
        ${ANDROID_ROOT}/external/libopus/silk/LPC_inv_pred_gain.c
        ${ANDROID_ROOT}/external/libopus/silk/table_LSF_cos.c
        ${ANDROID_ROOT}/external/libopus/silk/NLSF2A.c
        ${ANDROID_ROOT}/external/libopus/silk/NLSF_stabilize.c
        ${ANDROID_ROOT}/external/libopus/silk/NLSF_VQ_weights_laroia.c
        ${ANDROID_ROOT}/external/libopus/silk/pitch_est_tables.c
        ${ANDROID_ROOT}/external/libopus/silk/resampler.c
        ${ANDROID_ROOT}/external/libopus/silk/resampler_down2_3.c
        ${ANDROID_ROOT}/external/libopus/silk/resampler_down2.c
        ${ANDROID_ROOT}/external/libopus/silk/resampler_private_AR2.c
        ${ANDROID_ROOT}/external/libopus/silk/resampler_private_down_FIR.c
        ${ANDROID_ROOT}/external/libopus/silk/resampler_private_IIR_FIR.c
        ${ANDROID_ROOT}/external/libopus/silk/resampler_private_up2_HQ.c
        ${ANDROID_ROOT}/external/libopus/silk/resampler_rom.c
        ${ANDROID_ROOT}/external/libopus/silk/sigm_Q15.c
        ${ANDROID_ROOT}/external/libopus/silk/sort.c
        ${ANDROID_ROOT}/external/libopus/silk/sum_sqr_shift.c
        ${ANDROID_ROOT}/external/libopus/silk/stereo_decode_pred.c
        ${ANDROID_ROOT}/external/libopus/silk/stereo_encode_pred.c
        ${ANDROID_ROOT}/external/libopus/silk/stereo_find_predictor.c
        ${ANDROID_ROOT}/external/libopus/silk/stereo_quant_pred.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/LTP_analysis_filter_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/LTP_scale_ctrl_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/corrMatrix_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/encode_frame_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/find_LPC_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/find_LTP_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/find_pitch_lags_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/find_pred_coefs_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/noise_shape_analysis_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/process_gains_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/regularize_correlations_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/residual_energy16_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/residual_energy_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/warped_autocorrelation_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/apply_sine_window_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/autocorr_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/burg_modified_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/k2a_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/k2a_Q16_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/pitch_analysis_core_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/vector_ops_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/schur64_FIX.c
        ${ANDROID_ROOT}/external/libopus/silk/fixed/schur_FIX.c
        ${ANDROID_ROOT}/external/libopus/src/mapping_matrix.c
        ${ANDROID_ROOT}/external/libopus/src/opus.c
        ${ANDROID_ROOT}/external/libopus/src/opus_decoder.c
        ${ANDROID_ROOT}/external/libopus/src/opus_encoder.c
        ${ANDROID_ROOT}/external/libopus/src/opus_multistream.c
        ${ANDROID_ROOT}/external/libopus/src/opus_multistream_encoder.c
        ${ANDROID_ROOT}/external/libopus/src/opus_multistream_decoder.c
        ${ANDROID_ROOT}/external/libopus/src/opus_projection_encoder.c
        ${ANDROID_ROOT}/external/libopus/src/opus_projection_decoder.c
        ${ANDROID_ROOT}/external/libopus/src/repacketizer.c
        ${ANDROID_ROOT}/external/libopus/src/analysis.c
        ${ANDROID_ROOT}/external/libopus/src/mlp.c
        ${ANDROID_ROOT}/external/libopus/src/mlp_data.c
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
    add_library(libopus SHARED ${SOURCE_FILES})

    target_include_directories(libopus
            PUBLIC
            ${ANDROID_ROOT}/external/libopus/include
            ${ANDROID_ROOT}/external/libopus/src
            ${ANDROID_ROOT}/external/libopus/silk
            ${ANDROID_ROOT}/external/libopus/celt
            ${ANDROID_ROOT}/external/libopus/silk/fixed
            ${ANDROID_ROOT}/external/libopus
    )

    # Add c/c++ flags
    target_compile_options(libopus PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libopus
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
#################### libopus ####################

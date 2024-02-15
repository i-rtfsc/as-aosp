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

project(libFraunhoferAAC)

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

#################### libFraunhoferAAC ####################
set(FILES
        ${ANDROID_ROOT}/external/aac/libAACdec/src/FDK_delay.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aac_ram.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aac_rom.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aacdec_drc.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aacdec_hcr.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aacdec_hcr_bit.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aacdec_hcrs.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aacdec_pns.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aacdec_tns.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aacdecoder.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/aacdecoder_lib.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/block.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/channel.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/channelinfo.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/conceal.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/ldfiltbank.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/pulsedata.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/rvlc.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/rvlcbit.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/rvlcconceal.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/stereo.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/usacdec_ace_d4t64.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/usacdec_ace_ltp.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/usacdec_acelp.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/usacdec_fac.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/usacdec_lpc.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/usacdec_lpd.cpp
        ${ANDROID_ROOT}/external/aac/libAACdec/src/usacdec_rom.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/aacEnc_ram.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/aacEnc_rom.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/aacenc.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/aacenc_lib.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/aacenc_pns.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/aacenc_tns.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/adj_thr.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/band_nrg.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/bandwidth.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/bit_cnt.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/bitenc.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/block_switch.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/channel_map.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/chaosmeasure.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/dyn_bits.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/grp_data.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/intensity.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/line_pe.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/metadata_compressor.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/metadata_main.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/mps_main.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/ms_stereo.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/noisedet.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/pnsparam.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/pre_echo_control.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/psy_configuration.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/psy_main.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/qc_main.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/quantize.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/sf_estim.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/spreading.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/tonality.cpp
        ${ANDROID_ROOT}/external/aac/libAACenc/src/transform.cpp
        ${ANDROID_ROOT}/external/aac/libPCMutils/src/limiter.cpp
        ${ANDROID_ROOT}/external/aac/libPCMutils/src/pcm_utils.cpp
        ${ANDROID_ROOT}/external/aac/libPCMutils/src/pcmdmx_lib.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_bitbuffer.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_core.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_crc.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_decorrelate.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_hybrid.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_lpc.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_matrixCalloc.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_qmf_domain.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_tools_rom.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/FDK_trigFcts.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/autocorr2nd.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/dct.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/fft.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/fft_rad2.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/fixpoint_math.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/huff_nodes.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/mdct.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/nlc_dec.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/qmf.cpp
        ${ANDROID_ROOT}/external/aac/libFDK/src/scale.cpp
        ${ANDROID_ROOT}/external/aac/libSYS/src/genericStds.cpp
        ${ANDROID_ROOT}/external/aac/libSYS/src/syslib_channelMapDescr.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPDec/src/tpdec_adif.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPDec/src/tpdec_adts.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPDec/src/tpdec_asc.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPDec/src/tpdec_drm.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPDec/src/tpdec_latm.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPDec/src/tpdec_lib.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPEnc/src/tpenc_adif.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPEnc/src/tpenc_adts.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPEnc/src/tpenc_asc.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPEnc/src/tpenc_latm.cpp
        ${ANDROID_ROOT}/external/aac/libMpegTPEnc/src/tpenc_lib.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/HFgen_preFlat.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/env_calc.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/env_dec.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/env_extr.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/hbe.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/huff_dec.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/lpp_tran.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/psbitdec.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/psdec.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/psdec_drm.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/psdecrom_drm.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/pvc_dec.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/sbr_deb.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/sbr_dec.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/sbr_ram.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/sbr_rom.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/sbrdec_drc.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/sbrdec_freq_sca.cpp
        ${ANDROID_ROOT}/external/aac/libSBRdec/src/sbrdecoder.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/bit_sbr.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/code_env.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/env_bit.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/env_est.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/fram_gen.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/invf_est.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/mh_det.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/nf_est.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/ps_bitenc.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/ps_encode.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/ps_main.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/resampler.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/sbr_encoder.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/sbr_misc.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/sbrenc_freq_sca.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/sbrenc_ram.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/sbrenc_rom.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/ton_corr.cpp
        ${ANDROID_ROOT}/external/aac/libSBRenc/src/tran_det.cpp
        ${ANDROID_ROOT}/external/aac/libArithCoding/src/ac_arith_coder.cpp
        ${ANDROID_ROOT}/external/aac/libDRCdec/src/FDK_drcDecLib.cpp
        ${ANDROID_ROOT}/external/aac/libDRCdec/src/drcDec_gainDecoder.cpp
        ${ANDROID_ROOT}/external/aac/libDRCdec/src/drcDec_reader.cpp
        ${ANDROID_ROOT}/external/aac/libDRCdec/src/drcDec_rom.cpp
        ${ANDROID_ROOT}/external/aac/libDRCdec/src/drcDec_selectionProcess.cpp
        ${ANDROID_ROOT}/external/aac/libDRCdec/src/drcDec_tools.cpp
        ${ANDROID_ROOT}/external/aac/libDRCdec/src/drcGainDec_init.cpp
        ${ANDROID_ROOT}/external/aac/libDRCdec/src/drcGainDec_preprocess.cpp
        ${ANDROID_ROOT}/external/aac/libDRCdec/src/drcGainDec_process.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_bitdec.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_calcM1andM2.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_dec.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_dec_conceal.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_dec_lib.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_process.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_qmf.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_reshapeBBEnv.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_rom.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_smoothing.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_stp.cpp
        ${ANDROID_ROOT}/external/aac/libSACdec/src/sac_tsd.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_bitstream.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_delay.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_dmx_tdom_enh.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_filter.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_framewindowing.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_huff_tab.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_lib.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_nlc_enc.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_onsetdetect.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_paramextract.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_staticgain.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_tree.cpp
        ${ANDROID_ROOT}/external/aac/libSACenc/src/sacenc_vectorfunctions.cpp
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
    add_library(libFraunhoferAAC SHARED ${SOURCE_FILES})

    target_include_directories(libFraunhoferAAC
            PUBLIC
            ${ANDROID_ROOT}/external/aac/libAACdec/include
            ${ANDROID_ROOT}/external/aac/libAACenc/include
            ${ANDROID_ROOT}/external/aac/libPCMutils/include
            ${ANDROID_ROOT}/external/aac/libFDK/include
            ${ANDROID_ROOT}/external/aac/libSYS/include
            ${ANDROID_ROOT}/external/aac/libMpegTPDec/include
            ${ANDROID_ROOT}/external/aac/libMpegTPEnc/include
            ${ANDROID_ROOT}/external/aac/libSBRdec/include
            ${ANDROID_ROOT}/external/aac/libSBRenc/include
            ${ANDROID_ROOT}/external/aac/libArithCoding/include
            ${ANDROID_ROOT}/external/aac/libDRCdec/include
            ${ANDROID_ROOT}/external/aac/libSACdec/include
            ${ANDROID_ROOT}/external/aac/libSACenc/include
            ${ANDROID_ROOT}/external/aac
    )

    # Add c/c++ flags
    target_compile_options(libFraunhoferAAC PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libFraunhoferAAC
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
    )
endif ()
#################### libFraunhoferAAC ####################

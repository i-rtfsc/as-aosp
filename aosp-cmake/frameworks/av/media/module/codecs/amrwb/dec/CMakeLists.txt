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

project(libstagefright_amrwbdec)

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

#################### libstagefright_amrwbdec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/agc2_amr_wb.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/band_pass_6k_7k.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/dec_acelp_2p_in_64.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/dec_acelp_4p_in_64.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/dec_alg_codebook.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/dec_gain2_amr_wb.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/deemphasis_32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/dtx_decoder_amr_wb.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/get_amr_wb_bits.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/highpass_400hz_at_12k8.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/highpass_50hz_at_12k8.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/homing_amr_wb_dec.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/interpolate_isp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/isf_extrapolation.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/isp_az.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/isp_isf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/lagconceal.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/low_pass_filt_7k.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/median5.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/mime_io.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/noise_gen_amrwb.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/normalize_amr_wb.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/oversamp_12k8_to_16k.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/phase_dispersion.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/pit_shrp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/pred_lt4.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/preemph_amrwb_dec.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/pvamrwb_math_op.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/pvamrwbdecoder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/q_gain2_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/qisf_ns.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/qisf_ns_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/qpisf_2s.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/qpisf_2s_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/scale_signal.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/synthesis_amr_wb.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/voice_factor.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/wb_syn_filt.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src/weight_amrwb_lpc.cpp
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
    add_library(libstagefright_amrwbdec SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_amrwbdec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/src
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec/include
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/dec
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_amrwbdec PRIVATE
            -Werror
    )

    target_link_libraries(libstagefright_amrwbdec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libstagefright_amrwbdec ####################

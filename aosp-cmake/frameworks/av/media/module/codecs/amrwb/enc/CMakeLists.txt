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

project(libstagefright_amrwbenc)

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

#################### libstagefright_amrwbenc ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/autocorr.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/az_isp.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/bits.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/c2t64fx.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/c4t64fx.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/convolve.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/cor_h_x.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/decim54.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/deemph.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/dtx.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/g_pitch.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/gpclip.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/homing.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/hp400.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/hp50.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/hp6k.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/hp_wsp.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/int_lpc.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/isp_az.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/isp_isf.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/lag_wind.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/levinson.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/log2.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/lp_dec2.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/math_op.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/oper_32b.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/p_med_ol.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/pit_shrp.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/pitch_f4.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/pred_lt4.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/preemph.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/q_gain2.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/q_pulse.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/qisf_ns.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/qpisf_2s.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/random.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/residu.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/scale.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/stream.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/syn_filt.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/updt_tar.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/util.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/voAMRWBEnc.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/voicefac.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/wb_vad.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/weight_a.c
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src/mem_align.c
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
    add_library(libstagefright_amrwbenc SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_amrwbenc
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/inc
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc/src
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrwb/enc
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_amrwbenc PRIVATE
            -Werror
    )

    target_link_libraries(libstagefright_amrwbenc
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libstagefright_enc_common
    )
endif ()
#################### libstagefright_amrwbenc ####################

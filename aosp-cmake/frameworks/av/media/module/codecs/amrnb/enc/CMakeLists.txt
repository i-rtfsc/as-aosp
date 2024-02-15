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

project(libstagefright_amrnbenc)

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

#################### libstagefright_amrnbenc ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/amrencode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/autocorr.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/c1035pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/c2_11pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/c2_9pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/c3_14pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/c4_17pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/c8_31pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/calc_cor.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/calc_en.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/cbsearch.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/cl_ltp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/cod_amr.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/convolve.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/cor_h.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/cor_h_x.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/cor_h_x2.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/corrwght_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/dtx_enc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/enc_lag3.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/enc_lag6.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/enc_output_format_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/ets_to_if2.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/ets_to_wmf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/g_adapt.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/g_code.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/g_pitch.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/gain_q.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/hp_max.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/inter_36.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/inter_36_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/l_comp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/l_extract.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/l_negate.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/lag_wind.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/lag_wind_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/levinson.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/lpc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/ol_ltp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/p_ol_wgh.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/pitch_fr.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/pitch_ol.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/pre_big.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/pre_proc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/prm2bits.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/q_gain_c.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/q_gain_p.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/qgain475.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/qgain795.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/qua_gain.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/s10_8pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/set_sign.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/sid_sync.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/sp_enc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/spreproc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/spstproc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src/ton_stab.cpp
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
    add_library(libstagefright_amrnbenc SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_amrnbenc
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc/src
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/enc
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_amrnbenc PRIVATE
            -Werror
    )

    target_link_libraries(libstagefright_amrnbenc
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libstagefright_amrnb_common
            libstagefright_headers
    )
endif ()
#################### libstagefright_amrnbenc ####################

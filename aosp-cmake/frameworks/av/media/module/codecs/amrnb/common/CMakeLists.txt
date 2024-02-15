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

project(libstagefright_amrnb_common)

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

#################### libstagefright_amrnb_common ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/add.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/az_lsp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/bitno_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/bitreorder_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/bits2prm.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/c2_9pf_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/copy.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/div_32.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/div_s.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/extract_h.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/extract_l.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/gains_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/gc_pred.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/gmed_n.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/grid_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/gray_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/int_lpc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/inv_sqrt.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/inv_sqrt_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/l_abs.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/l_deposit_h.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/l_deposit_l.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/l_shr_r.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/log2.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/log2_norm.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/log2_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/lsfwt.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/lsp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/lsp_az.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/lsp_lsf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/lsp_lsf_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/lsp_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/mult_r.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/norm_l.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/norm_s.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/ph_disp_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/pow2.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/pow2_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/pred_lt.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/q_plsf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/q_plsf_3.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/q_plsf_3_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/q_plsf_5.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/q_plsf_5_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/qua_gain_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/reorder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/residu.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/round.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/set_zero.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/shr.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/shr_r.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/sqrt_l.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/sqrt_l_tbl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/sub.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/syn_filt.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/vad1.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/weight_a.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/src/window_tab.cpp
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
    add_library(libstagefright_amrnb_common SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_amrnb_common
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common/include
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/common
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_amrnb_common PRIVATE
            -Werror
    )

    target_link_libraries(libstagefright_amrnb_common
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libstagefright_amrnb_common ####################

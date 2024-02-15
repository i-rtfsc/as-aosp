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

project(libstagefright_amrnbdec)

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

#################### libstagefright_amrnbdec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/a_refl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/agc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/amrdecode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/b_cn_cod.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/bgnscd.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/c_g_aver.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d1035pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d2_11pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d2_9pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d3_14pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d4_17pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d8_31pf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d_gain_c.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d_gain_p.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d_plsf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d_plsf_3.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/d_plsf_5.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/dec_amr.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/dec_gain.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/dec_input_format_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/dec_lag3.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/dec_lag6.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/dtx_dec.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/ec_gains.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/ex_ctrl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/if2_to_ets.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/int_lsf.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/lsp_avg.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/ph_disp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/post_pro.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/preemph.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/pstfilt.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/qgain475_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/sp_dec.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src/wmf_to_ets.cpp
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
    add_library(libstagefright_amrnbdec SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_amrnbdec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec/src
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/amrnb/dec
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_amrnbdec PRIVATE
            -Werror
    )

    target_link_libraries(libstagefright_amrnbdec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libstagefright_amrnb_common
    )
endif ()
#################### libstagefright_amrnbdec ####################

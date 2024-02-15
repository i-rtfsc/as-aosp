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

project(libstagefright_m4vh263dec)

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

#################### libstagefright_m4vh263dec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/bitstream.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/block_idct.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/cal_dc_scaler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/combined_decode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/conceal.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/datapart_decode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/dcac_prediction.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/dec_pred_intra_dc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/get_pred_adv_b_add.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/get_pred_outside.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/idct.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/idct_vca.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/mb_motion_comp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/mb_utils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/packet_util.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/post_filter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/pvdec_api.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/scaling_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/vlc_decode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/vlc_dequant.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/vlc_tab.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/vop.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src/zigzag_tab.cpp
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
    add_library(libstagefright_m4vh263dec SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_m4vh263dec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/include
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec/src
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/dec
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_m4vh263dec PRIVATE
            -Werror
    )

    target_link_libraries(libstagefright_m4vh263dec
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
#################### libstagefright_m4vh263dec ####################

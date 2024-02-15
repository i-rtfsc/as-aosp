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

project(libmpeg2dec)

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

#################### libmpeg2dec ####################
set(FILES
        ${ANDROID_ROOT}/external/libmpeg2/common/impeg2_buf_mgr.c
        ${ANDROID_ROOT}/external/libmpeg2/common/impeg2_disp_mgr.c
        ${ANDROID_ROOT}/external/libmpeg2/common/impeg2_format_conv.c
        ${ANDROID_ROOT}/external/libmpeg2/common/impeg2_globals.c
        ${ANDROID_ROOT}/external/libmpeg2/common/impeg2_idct.c
        ${ANDROID_ROOT}/external/libmpeg2/common/impeg2_inter_pred.c
        ${ANDROID_ROOT}/external/libmpeg2/common/impeg2_job_queue.c
        ${ANDROID_ROOT}/external/libmpeg2/common/impeg2_mem_func.c
        ${ANDROID_ROOT}/external/libmpeg2/common/ithread.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_api_main.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_bitstream.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_debug.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_dec_hdr.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_decoder.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_d_pic.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_function_selector_generic.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_globals.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_i_pic.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_mc.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_mv_dec.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_pic_proc.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_pnb_pic.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_vld.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_vld_tables.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/impeg2d_deinterlace.c
        ${ANDROID_ROOT}/external/libmpeg2/common/icv_sad.c
        ${ANDROID_ROOT}/external/libmpeg2/common/icv_variance.c
        ${ANDROID_ROOT}/external/libmpeg2/common/ideint.c
        ${ANDROID_ROOT}/external/libmpeg2/common/ideint_cac.c
        ${ANDROID_ROOT}/external/libmpeg2/common/ideint_debug.c
        ${ANDROID_ROOT}/external/libmpeg2/common/ideint_function_selector_generic.c
        ${ANDROID_ROOT}/external/libmpeg2/common/ideint_utils.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/arm/impeg2d_function_selector.c
        ${ANDROID_ROOT}/external/libmpeg2/decoder/arm/impeg2d_function_selector_av8.c
        ${ANDROID_ROOT}/external/libmpeg2/common/arm/ideint_function_selector.c
        ${ANDROID_ROOT}/external/libmpeg2/common/arm/ideint_function_selector_av8.c
        ${ANDROID_ROOT}/external/libmpeg2/common/armv8/icv_sad_av8.s
        ${ANDROID_ROOT}/external/libmpeg2/common/armv8/icv_variance_av8.s
        ${ANDROID_ROOT}/external/libmpeg2/common/armv8/ideint_spatial_filter_av8.s
        ${ANDROID_ROOT}/external/libmpeg2/common/armv8/ideint_cac_av8.s
        ${ANDROID_ROOT}/external/libmpeg2/common/armv8/impeg2_neon_macros.s
        ${ANDROID_ROOT}/external/libmpeg2/common/armv8/impeg2_format_conv.s
        ${ANDROID_ROOT}/external/libmpeg2/common/armv8/impeg2_idct.s
        ${ANDROID_ROOT}/external/libmpeg2/common/armv8/impeg2_inter_pred.s
        ${ANDROID_ROOT}/external/libmpeg2/common/armv8/impeg2_mem_func.s
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
    add_library(libmpeg2dec SHARED ${SOURCE_FILES})

    target_include_directories(libmpeg2dec
            PUBLIC
            ${ANDROID_ROOT}/external/libmpeg2/decoder
            ${ANDROID_ROOT}/external/libmpeg2/common
            ${ANDROID_ROOT}/external/libmpeg2/decoder/arm
            ${ANDROID_ROOT}/external/libmpeg2/common/armv8
            ${ANDROID_ROOT}/external/libmpeg2
    )

    # Add c/c++ flags
    target_compile_options(libmpeg2dec PRIVATE
            -Werror
    )

    target_link_libraries(libmpeg2dec
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
#################### libmpeg2dec ####################

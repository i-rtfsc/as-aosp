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

project(libstagefright_mp3dec)

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

#################### libstagefright_mp3dec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_normalize.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_alias_reduction.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_crc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_decode_header.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_decode_huff_cw.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_getbits.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_dequantize_sample.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_framedecoder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_get_main_data_size.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_get_side_info.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_get_scale_factors.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_mpeg2_get_scale_data.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_mpeg2_get_scale_factors.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_mpeg2_stereo_proc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_huffman_decoding.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_huffman_parsing.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_tables.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_imdct_synth.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_mdct_6.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_dct_6.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_poly_phase_synthesis.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_equalizer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_seek_synch.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_stereo_proc.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_reorder.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_polyphase_filter_window.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_mdct_18.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_dct_9.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src/pvmp3_dct_16.cpp
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
    add_library(libstagefright_mp3dec SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_mp3dec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/include
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/mp3dec/src
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_mp3dec PRIVATE
            -Werror
    )

    target_link_libraries(libstagefright_mp3dec
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libstagefright_mp3dec_headers
    )
endif ()
#################### libstagefright_mp3dec ####################

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

project(libstagefright_m4vh263enc)

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

#################### libstagefright_m4vh263enc ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/bitstream_io.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/combined_encode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/datapart_encode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/dct.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/findhalfpel.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/fastcodemb.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/fastidct.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/fastquant.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/me_utils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/mp4enc_api.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/rate_control.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/motion_est.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/motion_comp.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/sad.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/sad_halfpel.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/vlc_encode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src/vop.cpp
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
    add_library(libstagefright_m4vh263enc SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_m4vh263enc
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/include
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc/src
            ${ANDROID_ROOT}/frameworks/av/media/module/codecs/m4v_h263/enc
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_m4vh263enc PRIVATE
            -Werror
    )

    target_link_libraries(libstagefright_m4vh263enc
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### libstagefright_m4vh263enc ####################

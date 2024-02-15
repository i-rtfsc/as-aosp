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

project(libstagefright_rtsp_player2)

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

#################### libstagefright_rtsp_player2 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AAMRAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AAVCAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AHEVCAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AH263Assembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AMPEG2TSAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AMPEG4AudioAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AMPEG4ElementaryAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/APacketSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARawAudioAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTPAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTPConnection.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTPSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTPWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTSPConnection.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ASessionDescription.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/JitterCalculator.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/SDPLoader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/NetworkUtilsForAppProc.cpp
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
    add_library(libstagefright_rtsp_player2 SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_rtsp_player2
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/include
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_rtsp_player2 PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_rtsp_player2
            libandroid_net
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdatasource
            libdl
            libm
            libmedia
            libstagefright_headers
            libstagefright_rtsp_headers
    )
endif ()
#################### libstagefright_rtsp_player2 ####################

#################### libstagefright_rtsp ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AAMRAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AAVCAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AHEVCAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AH263Assembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AMPEG2TSAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AMPEG4AudioAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/AMPEG4ElementaryAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/APacketSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARawAudioAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTPAssembler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTPConnection.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTPSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTPWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ARTSPConnection.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/ASessionDescription.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/JitterCalculator.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/SDPLoader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/NetworkUtils.cpp
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
    add_library(libstagefright_rtsp SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_rtsp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/rtsp/include
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_rtsp PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_rtsp
            libandroid_net
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdatasource
            libdl
            libm
            libmedia
            libnetd_client_headers
            libstagefright_headers
            libstagefright_rtsp_headers
    )
endif ()
#################### libstagefright_rtsp ####################

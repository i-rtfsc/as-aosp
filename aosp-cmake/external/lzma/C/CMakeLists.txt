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

project(liblzma)

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

#################### liblzma ####################
set(FILES
        ${ANDROID_ROOT}/external/lzma/C/7zAlloc.c
        ${ANDROID_ROOT}/external/lzma/C/7zArcIn.c
        ${ANDROID_ROOT}/external/lzma/C/7zBuf2.c
        ${ANDROID_ROOT}/external/lzma/C/7zBuf.c
        ${ANDROID_ROOT}/external/lzma/C/7zCrc.c
        ${ANDROID_ROOT}/external/lzma/C/7zCrcOpt.c
        ${ANDROID_ROOT}/external/lzma/C/7zDec.c
        ${ANDROID_ROOT}/external/lzma/C/7zFile.c
        ${ANDROID_ROOT}/external/lzma/C/7zStream.c
        ${ANDROID_ROOT}/external/lzma/C/Aes.c
        ${ANDROID_ROOT}/external/lzma/C/AesOpt.c
        ${ANDROID_ROOT}/external/lzma/C/Alloc.c
        ${ANDROID_ROOT}/external/lzma/C/Bcj2.c
        ${ANDROID_ROOT}/external/lzma/C/Bra86.c
        ${ANDROID_ROOT}/external/lzma/C/Bra.c
        ${ANDROID_ROOT}/external/lzma/C/BraIA64.c
        ${ANDROID_ROOT}/external/lzma/C/CpuArch.c
        ${ANDROID_ROOT}/external/lzma/C/Delta.c
        ${ANDROID_ROOT}/external/lzma/C/LzFind.c
        ${ANDROID_ROOT}/external/lzma/C/Lzma2Dec.c
        ${ANDROID_ROOT}/external/lzma/C/Lzma2Enc.c
        ${ANDROID_ROOT}/external/lzma/C/Lzma86Dec.c
        ${ANDROID_ROOT}/external/lzma/C/Lzma86Enc.c
        ${ANDROID_ROOT}/external/lzma/C/LzmaDec.c
        ${ANDROID_ROOT}/external/lzma/C/LzmaEnc.c
        ${ANDROID_ROOT}/external/lzma/C/LzmaLib.c
        ${ANDROID_ROOT}/external/lzma/C/Ppmd7.c
        ${ANDROID_ROOT}/external/lzma/C/Ppmd7Dec.c
        ${ANDROID_ROOT}/external/lzma/C/Ppmd7Enc.c
        ${ANDROID_ROOT}/external/lzma/C/Sha256.c
        ${ANDROID_ROOT}/external/lzma/C/Sort.c
        ${ANDROID_ROOT}/external/lzma/C/Xz.c
        ${ANDROID_ROOT}/external/lzma/C/XzCrc64.c
        ${ANDROID_ROOT}/external/lzma/C/XzCrc64Opt.c
        ${ANDROID_ROOT}/external/lzma/C/XzDec.c
        ${ANDROID_ROOT}/external/lzma/C/XzEnc.c
        ${ANDROID_ROOT}/external/lzma/C/XzIn.c
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
    add_library(liblzma SHARED ${SOURCE_FILES})

    target_include_directories(liblzma
            PUBLIC
            ${ANDROID_ROOT}/external/lzma/C
    )

    # Add c/c++ flags
    target_compile_options(liblzma PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(liblzma
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### liblzma ####################

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

project(tflite_support_libz)

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

#################### tflite_support_libz ####################
set(FILES
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/adler32.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/compress.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/cpu_features.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/crc32.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/deflate.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/gzclose.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/gzlib.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/gzread.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/gzwrite.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/infback.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/inffast.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/inflate.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/inftrees.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/trees.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/uncompr.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/zutil.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/contrib/minizip/ioapi.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/contrib/minizip/unzip.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/adler32_simd.c
        ${ANDROID_ROOT}/external/tflite-support/third_party/zlib/crc32_simd.c
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
    add_library(tflite_support_libz SHARED ${SOURCE_FILES})

    target_include_directories(tflite_support_libz
            PUBLIC
            ${ANDROID_ROOT}/external/tflite-support/third_party/zlib
    )

    # Add c/c++ flags
    target_compile_options(tflite_support_libz PRIVATE
            -Wall
            -Werror
            -Wno-unused
            -Wno-unused-parameter
    )

    target_link_libraries(tflite_support_libz
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### tflite_support_libz ####################

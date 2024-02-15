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

project(libbspatch)

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

#################### libbspatch ####################
set(FILES
        ${ANDROID_ROOT}/external/bsdiff/brotli_decompressor.cc
        ${ANDROID_ROOT}/external/bsdiff/bspatch.cc
        ${ANDROID_ROOT}/external/bsdiff/bz2_decompressor.cc
        ${ANDROID_ROOT}/external/bsdiff/buffer_file.cc
        ${ANDROID_ROOT}/external/bsdiff/decompressor_interface.cc
        ${ANDROID_ROOT}/external/bsdiff/extents.cc
        ${ANDROID_ROOT}/external/bsdiff/extents_file.cc
        ${ANDROID_ROOT}/external/bsdiff/file.cc
        ${ANDROID_ROOT}/external/bsdiff/logging.cc
        ${ANDROID_ROOT}/external/bsdiff/memory_file.cc
        ${ANDROID_ROOT}/external/bsdiff/patch_reader.cc
        ${ANDROID_ROOT}/external/bsdiff/sink_file.cc
        ${ANDROID_ROOT}/external/bsdiff/utils.cc
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
    add_library(libbspatch SHARED ${SOURCE_FILES})

    target_include_directories(libbspatch
            PUBLIC
            ${ANDROID_ROOT}/external/bsdiff/include
            ${ANDROID_ROOT}/external/bsdiff
    )

    # Add c/c++ flags
    target_compile_options(libbspatch PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libbspatch
            libbrotli
            libbz
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libbspatch ####################

#################### libbsdiff ####################
set(FILES
        ${ANDROID_ROOT}/external/bsdiff/brotli_compressor.cc
        ${ANDROID_ROOT}/external/bsdiff/bsdiff.cc
        ${ANDROID_ROOT}/external/bsdiff/bz2_compressor.cc
        ${ANDROID_ROOT}/external/bsdiff/compressor_buffer.cc
        ${ANDROID_ROOT}/external/bsdiff/diff_encoder.cc
        ${ANDROID_ROOT}/external/bsdiff/endsley_patch_writer.cc
        ${ANDROID_ROOT}/external/bsdiff/logging.cc
        ${ANDROID_ROOT}/external/bsdiff/patch_writer.cc
        ${ANDROID_ROOT}/external/bsdiff/patch_writer_factory.cc
        ${ANDROID_ROOT}/external/bsdiff/split_patch_writer.cc
        ${ANDROID_ROOT}/external/bsdiff/suffix_array_index.cc
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
    add_library(libbsdiff SHARED ${SOURCE_FILES})

    target_include_directories(libbsdiff
            PUBLIC
            ${ANDROID_ROOT}/external/bsdiff/include
            ${ANDROID_ROOT}/external/bsdiff
    )

    # Add c/c++ flags
    target_compile_options(libbsdiff PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libbsdiff
            libbrotli
            libbz
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdivsufsort
            libdivsufsort64
            libdl
            libm
    )
endif ()
#################### libbsdiff ####################

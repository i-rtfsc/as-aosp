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

project(puffin)

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

#################### puffin ####################
set(FILES
        ${ANDROID_ROOT}/external/puffin/src/extent_stream.cc
        ${ANDROID_ROOT}/external/puffin/src/main.cc
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
    add_executable(puffin ${SOURCE_FILES})

    target_include_directories(puffin
            PUBLIC
            ${ANDROID_ROOT}/external/puffin
            ${ANDROID_ROOT}/external/puffin/src/include
            ${OUT_INTERMEDIATES_ROOT}/external/puffin/libpuffpatch/android_${OUT_ARCH_CPU}_static/gen/proto/external/puffin
            ${OUT_INTERMEDIATES_ROOT}/external/puffin/libpuffpatch/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(puffin PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(puffin
            libbrillo
            libbrotli
            libbsdiff
            libbspatch
            libbz
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdivsufsort
            libdivsufsort64
            libdl
            libm
            libprotobuf-cpp-lite
            libpuffdiff
            libpuffpatch
            libzucchini
    )
endif ()
#################### puffin ####################

#################### libpuffpatch ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/puffin/libpuffpatch/android_${OUT_ARCH_CPU}_static/gen/proto/external/puffin/puffin/src/puffin.pb.cc
        ${ANDROID_ROOT}/external/puffin/src/bit_reader.cc
        ${ANDROID_ROOT}/external/puffin/src/bit_writer.cc
        ${ANDROID_ROOT}/external/puffin/src/brotli_util.cc
        ${ANDROID_ROOT}/external/puffin/src/huffer.cc
        ${ANDROID_ROOT}/external/puffin/src/huffman_table.cc
        ${ANDROID_ROOT}/external/puffin/src/memory_stream.cc
        ${ANDROID_ROOT}/external/puffin/src/puff_reader.cc
        ${ANDROID_ROOT}/external/puffin/src/puff_writer.cc
        ${ANDROID_ROOT}/external/puffin/src/puffer.cc
        ${ANDROID_ROOT}/external/puffin/src/puffin_stream.cc
        ${ANDROID_ROOT}/external/puffin/src/puffpatch.cc
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
    add_library(libpuffpatch SHARED ${SOURCE_FILES})

    target_include_directories(libpuffpatch
            PUBLIC
            ${ANDROID_ROOT}/external/puffin/src/include
            ${ANDROID_ROOT}/external/puffin
            ${OUT_INTERMEDIATES_ROOT}/external/puffin/libpuffpatch/android_${OUT_ARCH_CPU}_static/gen/proto/external/puffin
            ${OUT_INTERMEDIATES_ROOT}/external/puffin/libpuffpatch/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libpuffpatch PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libpuffpatch
            libbrotli
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libprotobuf-cpp-lite
            libzucchini
    )
endif ()
#################### libpuffpatch ####################

#################### libpuffdiff ####################
set(FILES
        ${ANDROID_ROOT}/external/puffin/src/file_stream.cc
        ${ANDROID_ROOT}/external/puffin/src/puffdiff.cc
        ${ANDROID_ROOT}/external/puffin/src/utils.cc
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
    add_library(libpuffdiff SHARED ${SOURCE_FILES})

    target_include_directories(libpuffdiff
            PUBLIC
            ${ANDROID_ROOT}/external/puffin/src/include
            ${ANDROID_ROOT}/external/puffin
            ${OUT_INTERMEDIATES_ROOT}/external/puffin/libpuffpatch/android_${OUT_ARCH_CPU}_static/gen/proto/external/puffin
            ${OUT_INTERMEDIATES_ROOT}/external/puffin/libpuffpatch/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libpuffdiff PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libpuffdiff
            libbrotli
            libbsdiff
            libbz
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libprotobuf-cpp-lite
            libpuffpatch
            libzucchini
    )
endif ()
#################### libpuffdiff ####################

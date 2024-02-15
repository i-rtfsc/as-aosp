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

project(zlib_bench)

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

#################### zlib_uncompress_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zlib/contrib/tests/fuzzers/uncompress_fuzzer.cc
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
    add_executable(zlib_uncompress_fuzzer ${SOURCE_FILES})

    target_include_directories(zlib_uncompress_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zlib
    )

    # Add c/c++ flags
    target_compile_options(zlib_uncompress_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(zlib_uncompress_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libz
    )
endif ()
#################### zlib_uncompress_fuzzer ####################

#################### zlib_streaming_inflate_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zlib/contrib/tests/fuzzers/streaming_inflate_fuzzer.cc
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
    add_executable(zlib_streaming_inflate_fuzzer ${SOURCE_FILES})

    target_include_directories(zlib_streaming_inflate_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zlib
    )

    # Add c/c++ flags
    target_compile_options(zlib_streaming_inflate_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(zlib_streaming_inflate_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libz
    )
endif ()
#################### zlib_streaming_inflate_fuzzer ####################

#################### zlib_google_compression_utils_portable ####################
set(FILES
        ${ANDROID_ROOT}/external/zlib/google/compression_utils_portable.cc
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
    add_library(zlib_google_compression_utils_portable SHARED ${SOURCE_FILES})

    target_include_directories(zlib_google_compression_utils_portable
            PUBLIC
            ${ANDROID_ROOT}/external/zlib
            ${ANDROID_ROOT}/external/zlib/google
    )

    # Add c/c++ flags
    target_compile_options(zlib_google_compression_utils_portable PRIVATE
            -Wall
            -Werror
            -Wno-unused
            -Wno-unused-parameter
    )

    target_link_libraries(zlib_google_compression_utils_portable
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            libz
    )
endif ()
#################### zlib_google_compression_utils_portable ####################

#################### zlib_inflate_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zlib/contrib/tests/fuzzers/inflate_fuzzer.cc
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
    add_executable(zlib_inflate_fuzzer ${SOURCE_FILES})

    target_include_directories(zlib_inflate_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zlib
    )

    # Add c/c++ flags
    target_compile_options(zlib_inflate_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(zlib_inflate_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libz
    )
endif ()
#################### zlib_inflate_fuzzer ####################

#################### zlib_deflate_set_dictionary_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zlib/contrib/tests/fuzzers/deflate_set_dictionary_fuzzer.cc
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
    add_executable(zlib_deflate_set_dictionary_fuzzer ${SOURCE_FILES})

    target_include_directories(zlib_deflate_set_dictionary_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zlib
    )

    # Add c/c++ flags
    target_compile_options(zlib_deflate_set_dictionary_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(zlib_deflate_set_dictionary_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libz
    )
endif ()
#################### zlib_deflate_set_dictionary_fuzzer ####################

#################### libz_stable ####################
set(FILES
        ${ANDROID_ROOT}/external/zlib/adler32.c
        ${ANDROID_ROOT}/external/zlib/adler32_simd.c
        ${ANDROID_ROOT}/external/zlib/compress.c
        ${ANDROID_ROOT}/external/zlib/cpu_features.c
        ${ANDROID_ROOT}/external/zlib/crc32.c
        ${ANDROID_ROOT}/external/zlib/crc32_simd.c
        ${ANDROID_ROOT}/external/zlib/crc_folding.c
        ${ANDROID_ROOT}/external/zlib/deflate.c
        ${ANDROID_ROOT}/external/zlib/gzclose.c
        ${ANDROID_ROOT}/external/zlib/gzlib.c
        ${ANDROID_ROOT}/external/zlib/gzread.c
        ${ANDROID_ROOT}/external/zlib/gzwrite.c
        ${ANDROID_ROOT}/external/zlib/infback.c
        ${ANDROID_ROOT}/external/zlib/inffast.c
        ${ANDROID_ROOT}/external/zlib/inflate.c
        ${ANDROID_ROOT}/external/zlib/inftrees.c
        ${ANDROID_ROOT}/external/zlib/trees.c
        ${ANDROID_ROOT}/external/zlib/uncompr.c
        ${ANDROID_ROOT}/external/zlib/zutil.c
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
    add_library(libz_stable SHARED ${SOURCE_FILES})

    target_include_directories(libz_stable
            PUBLIC
            ${ANDROID_ROOT}/external/zlib
    )

    # Add c/c++ flags
    target_compile_options(libz_stable PRIVATE
            -Wall
            -Werror
            -Wno-unused
            -Wno-unused-parameter
    )

    target_link_libraries(libz_stable
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libz_stable ####################

#################### zlib_deflate_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zlib/contrib/tests/fuzzers/deflate_fuzzer.cc
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
    add_executable(zlib_deflate_fuzzer ${SOURCE_FILES})

    target_include_directories(zlib_deflate_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zlib
    )

    # Add c/c++ flags
    target_compile_options(zlib_deflate_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(zlib_deflate_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libz
    )
endif ()
#################### zlib_deflate_fuzzer ####################

#################### libz_static ####################
set(FILES
        ${ANDROID_ROOT}/external/zlib/adler32.c
        ${ANDROID_ROOT}/external/zlib/adler32_simd.c
        ${ANDROID_ROOT}/external/zlib/compress.c
        ${ANDROID_ROOT}/external/zlib/cpu_features.c
        ${ANDROID_ROOT}/external/zlib/crc32.c
        ${ANDROID_ROOT}/external/zlib/crc32_simd.c
        ${ANDROID_ROOT}/external/zlib/crc_folding.c
        ${ANDROID_ROOT}/external/zlib/deflate.c
        ${ANDROID_ROOT}/external/zlib/gzclose.c
        ${ANDROID_ROOT}/external/zlib/gzlib.c
        ${ANDROID_ROOT}/external/zlib/gzread.c
        ${ANDROID_ROOT}/external/zlib/gzwrite.c
        ${ANDROID_ROOT}/external/zlib/infback.c
        ${ANDROID_ROOT}/external/zlib/inffast.c
        ${ANDROID_ROOT}/external/zlib/inflate.c
        ${ANDROID_ROOT}/external/zlib/inftrees.c
        ${ANDROID_ROOT}/external/zlib/trees.c
        ${ANDROID_ROOT}/external/zlib/uncompr.c
        ${ANDROID_ROOT}/external/zlib/zutil.c
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
    add_library(libz_static SHARED ${SOURCE_FILES})

    target_include_directories(libz_static
            PUBLIC
            ${ANDROID_ROOT}/external/zlib
    )

    # Add c/c++ flags
    target_compile_options(libz_static PRIVATE
            -Wall
            -Werror
            -Wno-unused
            -Wno-unused-parameter
    )

    target_link_libraries(libz_static
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libz_static ####################

#################### zlib_bench ####################
set(FILES
        ${ANDROID_ROOT}/external/zlib/contrib/bench/zlib_bench.cc
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
    add_executable(zlib_bench ${SOURCE_FILES})

    target_include_directories(zlib_bench
            PUBLIC
            ${ANDROID_ROOT}/external/zlib
    )

    # Add c/c++ flags
    target_compile_options(zlib_bench PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(zlib_bench
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
    )
endif ()
#################### zlib_bench ####################

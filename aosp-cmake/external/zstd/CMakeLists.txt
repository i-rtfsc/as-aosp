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

project(zstd_dict_decompress_fuzzer)

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

#################### zstd_frame_info_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_frame_info.c
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
    add_executable(zstd_frame_info_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_frame_info_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_frame_info_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_frame_info_fuzzer ####################

#################### zstd_dict_stream_round_trip_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/dictionary_stream_round_trip.c
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
    add_executable(zstd_dict_stream_round_trip_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_dict_stream_round_trip_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_dict_stream_round_trip_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_dict_stream_round_trip_fuzzer ####################

#################### zstd_stream_decompress_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/stream_decompress.c
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
    add_executable(zstd_stream_decompress_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_stream_decompress_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_stream_decompress_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_stream_decompress_fuzzer ####################

#################### zstd_stream_round_trip_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/stream_round_trip.c
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
    add_executable(zstd_stream_round_trip_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_stream_round_trip_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_stream_round_trip_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_stream_round_trip_fuzzer ####################

#################### libzstd ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/lib/common/debug.c
        ${ANDROID_ROOT}/external/zstd/lib/common/entropy_common.c
        ${ANDROID_ROOT}/external/zstd/lib/common/error_private.c
        ${ANDROID_ROOT}/external/zstd/lib/common/fse_decompress.c
        ${ANDROID_ROOT}/external/zstd/lib/common/pool.c
        ${ANDROID_ROOT}/external/zstd/lib/common/threading.c
        ${ANDROID_ROOT}/external/zstd/lib/common/xxhash.c
        ${ANDROID_ROOT}/external/zstd/lib/common/zstd_common.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/fse_compress.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/hist.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/huf_compress.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstd_compress.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstd_compress_literals.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstd_compress_sequences.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstd_compress_superblock.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstd_double_fast.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstd_fast.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstd_lazy.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstd_ldm.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstd_opt.c
        ${ANDROID_ROOT}/external/zstd/lib/compress/zstdmt_compress.c
        ${ANDROID_ROOT}/external/zstd/lib/decompress/huf_decompress.c
        ${ANDROID_ROOT}/external/zstd/lib/decompress/zstd_ddict.c
        ${ANDROID_ROOT}/external/zstd/lib/decompress/zstd_decompress.c
        ${ANDROID_ROOT}/external/zstd/lib/decompress/zstd_decompress_block.c
        ${ANDROID_ROOT}/external/zstd/lib/deprecated/zbuff_common.c
        ${ANDROID_ROOT}/external/zstd/lib/deprecated/zbuff_compress.c
        ${ANDROID_ROOT}/external/zstd/lib/deprecated/zbuff_decompress.c
        ${ANDROID_ROOT}/external/zstd/lib/dictBuilder/cover.c
        ${ANDROID_ROOT}/external/zstd/lib/dictBuilder/divsufsort.c
        ${ANDROID_ROOT}/external/zstd/lib/dictBuilder/fastcover.c
        ${ANDROID_ROOT}/external/zstd/lib/dictBuilder/zdict.c
        ${ANDROID_ROOT}/external/zstd/lib/legacy/zstd_v01.c
        ${ANDROID_ROOT}/external/zstd/lib/legacy/zstd_v02.c
        ${ANDROID_ROOT}/external/zstd/lib/legacy/zstd_v03.c
        ${ANDROID_ROOT}/external/zstd/lib/legacy/zstd_v04.c
        ${ANDROID_ROOT}/external/zstd/lib/legacy/zstd_v05.c
        ${ANDROID_ROOT}/external/zstd/lib/legacy/zstd_v06.c
        ${ANDROID_ROOT}/external/zstd/lib/legacy/zstd_v07.c
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
    add_library(libzstd SHARED ${SOURCE_FILES})

    target_include_directories(libzstd
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd
    )

    # Add c/c++ flags
    target_compile_options(libzstd PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libzstd
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libzstd ####################

#################### zstd_raw_dict_round_trip_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/raw_dictionary_round_trip.c
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
    add_executable(zstd_raw_dict_round_trip_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_raw_dict_round_trip_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_raw_dict_round_trip_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_raw_dict_round_trip_fuzzer ####################

#################### zstd_simple_decompress_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/simple_decompress.c
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
    add_executable(zstd_simple_decompress_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_simple_decompress_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_simple_decompress_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_simple_decompress_fuzzer ####################

#################### zstd_simple_round_trip_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/simple_round_trip.c
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
    add_executable(zstd_simple_round_trip_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_simple_round_trip_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_simple_round_trip_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_simple_round_trip_fuzzer ####################

#################### zstd_block_round_trip_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/block_round_trip.c
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
    add_executable(zstd_block_round_trip_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_block_round_trip_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_block_round_trip_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_block_round_trip_fuzzer ####################

#################### zstd_simple_compress_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/simple_compress.c
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
    add_executable(zstd_simple_compress_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_simple_compress_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_simple_compress_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_simple_compress_fuzzer ####################

#################### zstd_dict_round_trip_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/dictionary_round_trip.c
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
    add_executable(zstd_dict_round_trip_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_dict_round_trip_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_dict_round_trip_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_dict_round_trip_fuzzer ####################

#################### zstd_block_decompress_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/block_decompress.c
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
    add_executable(zstd_block_decompress_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_block_decompress_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_block_decompress_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_block_decompress_fuzzer ####################

#################### zstd_dict_loader_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/dictionary_loader.c
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
    add_executable(zstd_dict_loader_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_dict_loader_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_dict_loader_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_dict_loader_fuzzer ####################

#################### zstd_dict_decompress_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_data_producer.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/fuzz_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/zstd_helpers.c
        ${ANDROID_ROOT}/external/zstd/tests/fuzz/dictionary_decompress.c
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
    add_executable(zstd_dict_decompress_fuzzer ${SOURCE_FILES})

    target_include_directories(zstd_dict_decompress_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/zstd/lib/common
            ${ANDROID_ROOT}/external/zstd/lib/dictBuilder
            ${ANDROID_ROOT}/external/zstd/tests/fuzz
            ${ANDROID_ROOT}/external/zstd
            ${ANDROID_ROOT}/external/zstd/lib
    )

    target_link_libraries(zstd_dict_decompress_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libzstd
    )
endif ()
#################### zstd_dict_decompress_fuzzer ####################

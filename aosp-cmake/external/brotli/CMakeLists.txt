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

project(brotli)

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

#################### libbrotli_encoder_jni ####################
set(FILES
        ${ANDROID_ROOT}/external/brotli/java/org/brotli/wrapper/enc/encoder_jni.cc
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
    add_library(libbrotli_encoder_jni SHARED ${SOURCE_FILES})

    target_include_directories(libbrotli_encoder_jni
            PUBLIC
            ${ANDROID_ROOT}/external/brotli
            ${ANDROID_ROOT}/external/brotli/c/include
    )

    # Add c/c++ flags
    target_compile_options(libbrotli_encoder_jni PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libbrotli_encoder_jni
            libbrotli
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnativehelper_lazy
    )
endif ()
#################### libbrotli_encoder_jni ####################

#################### brotli-fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/brotli/c/fuzz/decode_fuzzer.c
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
    add_executable(brotli-fuzzer ${SOURCE_FILES})

    target_include_directories(brotli-fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/brotli
            ${ANDROID_ROOT}/external/brotli/c/include
    )

    # Add c/c++ flags
    target_compile_options(brotli-fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(brotli-fuzzer
            libbrotli
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
    )
endif ()
#################### brotli-fuzzer ####################

#################### brotli ####################
set(FILES
        ${ANDROID_ROOT}/external/brotli/c/tools/brotli.c
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
    add_executable(brotli ${SOURCE_FILES})

    target_include_directories(brotli
            PUBLIC
            ${ANDROID_ROOT}/external/brotli
            ${ANDROID_ROOT}/external/brotli/c/include
    )

    # Add c/c++ flags
    target_compile_options(brotli PRIVATE
            -Werror
    )

    target_link_libraries(brotli
            libbrotli
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### brotli ####################

#################### libbrotli ####################
set(FILES
        ${ANDROID_ROOT}/external/brotli/c/common/constants.c
        ${ANDROID_ROOT}/external/brotli/c/common/context.c
        ${ANDROID_ROOT}/external/brotli/c/common/dictionary.c
        ${ANDROID_ROOT}/external/brotli/c/common/platform.c
        ${ANDROID_ROOT}/external/brotli/c/common/transform.c
        ${ANDROID_ROOT}/external/brotli/c/dec/bit_reader.c
        ${ANDROID_ROOT}/external/brotli/c/dec/decode.c
        ${ANDROID_ROOT}/external/brotli/c/dec/huffman.c
        ${ANDROID_ROOT}/external/brotli/c/dec/state.c
        ${ANDROID_ROOT}/external/brotli/c/enc/backward_references.c
        ${ANDROID_ROOT}/external/brotli/c/enc/backward_references_hq.c
        ${ANDROID_ROOT}/external/brotli/c/enc/bit_cost.c
        ${ANDROID_ROOT}/external/brotli/c/enc/block_splitter.c
        ${ANDROID_ROOT}/external/brotli/c/enc/brotli_bit_stream.c
        ${ANDROID_ROOT}/external/brotli/c/enc/cluster.c
        ${ANDROID_ROOT}/external/brotli/c/enc/command.c
        ${ANDROID_ROOT}/external/brotli/c/enc/compress_fragment.c
        ${ANDROID_ROOT}/external/brotli/c/enc/compress_fragment_two_pass.c
        ${ANDROID_ROOT}/external/brotli/c/enc/dictionary_hash.c
        ${ANDROID_ROOT}/external/brotli/c/enc/encode.c
        ${ANDROID_ROOT}/external/brotli/c/enc/encoder_dict.c
        ${ANDROID_ROOT}/external/brotli/c/enc/entropy_encode.c
        ${ANDROID_ROOT}/external/brotli/c/enc/fast_log.c
        ${ANDROID_ROOT}/external/brotli/c/enc/histogram.c
        ${ANDROID_ROOT}/external/brotli/c/enc/literal_cost.c
        ${ANDROID_ROOT}/external/brotli/c/enc/memory.c
        ${ANDROID_ROOT}/external/brotli/c/enc/metablock.c
        ${ANDROID_ROOT}/external/brotli/c/enc/static_dict.c
        ${ANDROID_ROOT}/external/brotli/c/enc/utf8_util.c
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
    add_library(libbrotli SHARED ${SOURCE_FILES})

    target_include_directories(libbrotli
            PUBLIC
            ${ANDROID_ROOT}/external/brotli/c/include
            ${ANDROID_ROOT}/external/brotli
    )

    # Add c/c++ flags
    target_compile_options(libbrotli PRIVATE
            -Werror
    )

    target_link_libraries(libbrotli
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libbrotli ####################

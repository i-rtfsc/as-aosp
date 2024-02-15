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

project(sbcdecoder_fuzzer)

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

#################### libbt-sbc-decoder ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/alloc.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/bitalloc-sbc.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/bitalloc.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/bitstream-decode.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/decoder-oina.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/decoder-private.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/decoder-sbc.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/dequant.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/framing-sbc.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/framing.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/oi_codec_version.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/synthesis-8-generated.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/synthesis-dct8.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce/synthesis-sbc.c
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
    add_library(libbt-sbc-decoder SHARED ${SOURCE_FILES})

    target_include_directories(libbt-sbc-decoder
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/srce
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder
    )

    # Add c/c++ flags
    target_compile_options(libbt-sbc-decoder PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
    )

    target_link_libraries(libbt-sbc-decoder
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-platform-protos-lite
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libstatslog_bt
            libstatssocket
            libutils
    )
endif ()
#################### libbt-sbc-decoder ####################

#################### sbcdecoder_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/sbcdecoder_fuzzer.c
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
    add_executable(sbcdecoder_fuzzer ${SOURCE_FILES})

    target_include_directories(sbcdecoder_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/sbc/decoder
    )

    # Add c/c++ flags
    target_compile_options(sbcdecoder_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(sbcdecoder_fuzzer
            libbt-sbc-decoder
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
#################### sbcdecoder_fuzzer ####################

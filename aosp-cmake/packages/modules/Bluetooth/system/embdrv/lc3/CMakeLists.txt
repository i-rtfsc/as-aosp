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

project(liblc3_fuzzer)

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

#################### liblc3 ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/attdet.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/bits.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/bwdet.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/energy.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/lc3.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/ltpf.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/mdct.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/plc.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/sns.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/spec.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/tables.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/src/tns.c
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
    add_library(liblc3 SHARED ${SOURCE_FILES})

    target_include_directories(liblc3
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3
    )

    # Add c/c++ flags
    target_compile_options(liblc3 PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(liblc3
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
#################### liblc3 ####################

#################### lc3_decoder ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/tools/dlc3.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/tools/lc3bin.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/tools/wave.c
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
    add_executable(lc3_decoder ${SOURCE_FILES})

    target_include_directories(lc3_decoder
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/include
    )

    # Add c/c++ flags
    target_compile_options(lc3_decoder PRIVATE
            -Werror
    )

    target_link_libraries(lc3_decoder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblc3
            libm
    )
endif ()
#################### lc3_decoder ####################

#################### lc3_encoder ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/tools/elc3.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/tools/lc3bin.c
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/tools/wave.c
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
    add_executable(lc3_encoder ${SOURCE_FILES})

    target_include_directories(lc3_encoder
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/include
    )

    # Add c/c++ flags
    target_compile_options(lc3_encoder PRIVATE
            -Werror
    )

    target_link_libraries(lc3_encoder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblc3
            libm
    )
endif ()
#################### lc3_encoder ####################

#################### liblc3_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/fuzzer/liblc3_fuzzer.cpp
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
    add_executable(liblc3_fuzzer ${SOURCE_FILES})

    target_include_directories(liblc3_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/embdrv/lc3/include
    )

    # Add c/c++ flags
    target_compile_options(liblc3_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(liblc3_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblc3
            libm
    )
endif ()
#################### liblc3_fuzzer ####################

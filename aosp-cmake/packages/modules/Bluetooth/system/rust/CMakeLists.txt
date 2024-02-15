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

project(libbluetooth_core_rs_bridge)

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

#################### libbluetooth_core_rs_bridge ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/rust/src/connection/ffi/connection_shim.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/rust/src/core/ffi/module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/rust/src/gatt/ffi/gatt_shim.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/rust/libbluetooth_core_rs_bridge_codegen/gen/gensrcs/packages/modules/Bluetooth/system/rust/src/connection/ffi.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/rust/libbluetooth_core_rs_bridge_codegen/gen/gensrcs/packages/modules/Bluetooth/system/rust/src/core/ffi.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/rust/libbluetooth_core_rs_bridge_codegen/gen/gensrcs/packages/modules/Bluetooth/system/rust/src/gatt/ffi.cc
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
    add_library(libbluetooth_core_rs_bridge SHARED ${SOURCE_FILES})

    target_include_directories(libbluetooth_core_rs_bridge
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/rust
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/rust/libbluetooth_core_rs_bridge_codegen/gen/gensrcs
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/rust/libbluetooth_core_rs_bridge_codegen_header/gen/gensrcs/packages/modules/Bluetooth/system/rust
    )

    # Add c/c++ flags
    target_compile_options(libbluetooth_core_rs_bridge PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libbluetooth_core_rs_bridge
            libbase
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libcutils
            libdl
            libevent
            libflatbuffers-cpp
            liblog
            libm
            libmodpb64
    )
endif ()
#################### libbluetooth_core_rs_bridge ####################

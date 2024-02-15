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

project(libring-core)

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

#################### libring-core ####################
set(FILES
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/cpu-intel.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/fipsmodule/aes/aes_nohw.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/crypto.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/poly1305/poly1305.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/poly1305/poly1305_vec.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/curve25519/curve25519.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/fipsmodule/ec/ecp_nistz.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/fipsmodule/ec/gfp_p256.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/fipsmodule/ec/gfp_p384.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/fipsmodule/ec/p256.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/fipsmodule/rsa/padding.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/limbs/limbs.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/mem.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/fipsmodule/bn/montgomery.c
        ${ANDROID_ROOT}/external/rust/crates/ring/crypto/fipsmodule/bn/montgomery_inv.c
        ${ANDROID_ROOT}/external/rust/crates/ring/pregenerated/aesv8-armx-linux64.S
        ${ANDROID_ROOT}/external/rust/crates/ring/pregenerated/armv8-mont-linux64.S
        ${ANDROID_ROOT}/external/rust/crates/ring/pregenerated/chacha-armv8-linux64.S
        ${ANDROID_ROOT}/external/rust/crates/ring/pregenerated/ghash-neon-armv8-linux64.S
        ${ANDROID_ROOT}/external/rust/crates/ring/pregenerated/ghashv8-armx-linux64.S
        ${ANDROID_ROOT}/external/rust/crates/ring/pregenerated/sha256-armv8-linux64.S
        ${ANDROID_ROOT}/external/rust/crates/ring/pregenerated/sha512-armv8-linux64.S
        ${ANDROID_ROOT}/external/rust/crates/ring/pregenerated/vpaes-armv8-linux64.S
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
    add_library(libring-core SHARED ${SOURCE_FILES})

    target_include_directories(libring-core
            PUBLIC
            ${ANDROID_ROOT}/external/rust/crates/ring/include
            ${ANDROID_ROOT}/external/rust/crates/ring
    )

    # Add c/c++ flags
    target_compile_options(libring-core PRIVATE
            -DNDEBUG
    )

    target_link_libraries(libring-core
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libring-core ####################

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

project(libsrtp2-fuzzer)

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

#################### libsrtp2-fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libsrtp2/fuzzer/fuzzer.c
        ${ANDROID_ROOT}/external/libsrtp2/fuzzer/mt19937.cpp
        ${ANDROID_ROOT}/external/libsrtp2/fuzzer/testmem.c
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
    add_executable(libsrtp2-fuzzer ${SOURCE_FILES})

    target_include_directories(libsrtp2-fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libsrtp2
            ${ANDROID_ROOT}/external/libsrtp2/crypto/include
            ${ANDROID_ROOT}/external/libsrtp2/include
            ${ANDROID_ROOT}/external/libsrtp2/exported_include
    )

    # Add c/c++ flags
    target_compile_options(libsrtp2-fuzzer PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libsrtp2-fuzzer
            libc++
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libsrtp2
            libc
            libdl
            libm
    )
endif ()
#################### libsrtp2-fuzzer ####################

#################### libsrtp2 ####################
set(FILES
        ${ANDROID_ROOT}/external/libsrtp2/crypto/cipher/aes.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/cipher/aes_icm.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/cipher/cipher.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/cipher/null_cipher.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/hash/auth.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/hash/hmac.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/hash/null_auth.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/hash/sha1.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/kernel/alloc.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/kernel/crypto_kernel.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/kernel/err.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/kernel/key.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/math/datatypes.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/math/stat.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/replay/rdb.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/replay/rdbx.c
        ${ANDROID_ROOT}/external/libsrtp2/crypto/replay/ut_sim.c
        ${ANDROID_ROOT}/external/libsrtp2/srtp/ekt.c
        ${ANDROID_ROOT}/external/libsrtp2/srtp/srtp.c
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
    add_library(libsrtp2 SHARED ${SOURCE_FILES})

    target_include_directories(libsrtp2
            PUBLIC
            ${ANDROID_ROOT}/external/libsrtp2/crypto/include
            ${ANDROID_ROOT}/external/libsrtp2/include
            ${ANDROID_ROOT}/external/libsrtp2/exported_include
            ${ANDROID_ROOT}/external/libsrtp2
    )

    # Add c/c++ flags
    target_compile_options(libsrtp2 PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libsrtp2
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libsrtp2 ####################

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

project(libconscrypt_jni)

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

#################### libconscrypt_jni ####################
set(FILES
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/compatibility_close_monitor.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/jniload.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/jniutil.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/native_crypto.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/netutil.cc
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
    add_library(libconscrypt_jni SHARED ${SOURCE_FILES})

    target_include_directories(libconscrypt_jni
            PUBLIC
            ${ANDROID_ROOT}/external/conscrypt/common/src/jni/unbundled/include
            ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/include
            ${ANDROID_ROOT}/external/conscrypt
    )

    # Add c/c++ flags
    target_compile_options(libconscrypt_jni PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libconscrypt_jni
            jni_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto_static
            libdl
            liblog
            libm
            libssl
    )
endif ()
#################### libconscrypt_jni ####################

#################### libconscrypt_static ####################
set(FILES
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/compatibility_close_monitor.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/jniload.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/jniutil.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/native_crypto.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/netutil.cc
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
    add_library(libconscrypt_static SHARED ${SOURCE_FILES})

    target_include_directories(libconscrypt_static
            PUBLIC
            ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/include
            ${ANDROID_ROOT}/external/conscrypt/common/src/jni/unbundled/include
            ${ANDROID_ROOT}/external/conscrypt
    )

    # Add c/c++ flags
    target_compile_options(libconscrypt_static PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libconscrypt_static
            jni_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto_static
            libdl
            libm
            libssl
    )
endif ()
#################### libconscrypt_static ####################

#################### libjavacrypto ####################
set(FILES
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/compatibility_close_monitor.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/jniload.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/jniutil.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/native_crypto.cc
        ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/cpp/conscrypt/netutil.cc
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
    add_library(libjavacrypto SHARED ${SOURCE_FILES})

    target_include_directories(libjavacrypto
            PUBLIC
            ${ANDROID_ROOT}/external/conscrypt/common/src/jni/main/include
            ${ANDROID_ROOT}/external/conscrypt
    )

    # Add c/c++ flags
    target_compile_options(libjavacrypto PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libjavacrypto
            jni_headers
            libc++
            libclang_rt.builtins
            libcrypto
            liblog
            libnativehelper_header_only
            libssl
    )
endif ()
#################### libjavacrypto ####################

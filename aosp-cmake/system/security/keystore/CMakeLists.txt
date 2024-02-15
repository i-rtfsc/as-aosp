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

project(keystore_cli_v2)

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

#################### libkeystore-wifi-hidl ####################
set(FILES
        ${ANDROID_ROOT}/system/security/keystore/keystore_get_wifi_hidl.cpp
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
    add_library(libkeystore-wifi-hidl SHARED ${SOURCE_FILES})

    target_include_directories(libkeystore-wifi-hidl
            PUBLIC
            ${ANDROID_ROOT}/system/security/keystore/include
            ${ANDROID_ROOT}/system/security/keystore
    )

    # Add c/c++ flags
    target_compile_options(libkeystore-wifi-hidl PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libkeystore-wifi-hidl
            android.system.wifi.keystore-1.0
            libbase
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libkeystore-wifi-hidl ####################

#################### keystore_cli_v2 ####################
set(FILES
        ${ANDROID_ROOT}/system/security/keystore/keystore_cli_v2.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore/keystore_cli_v2/android_${OUT_ARCH_CPU}_cfi/gen/proto/system/security/keystore/keystore_client.pb.cc
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
    add_executable(keystore_cli_v2 ${SOURCE_FILES})

    target_include_directories(keystore_cli_v2
            PUBLIC
            ${ANDROID_ROOT}/system/security/keystore/include
            ${ANDROID_ROOT}/system/security/keystore
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore/keystore_cli_v2/android_${OUT_ARCH_CPU}_cfi/gen/proto/system/security/keystore
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore/keystore_cli_v2/android_${OUT_ARCH_CPU}_cfi/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(keystore_cli_v2 PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(keystore_cli_v2
            android.security.apc-ndk
            android.system.keystore2-V3-ndk
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            libkeymint_support
            libm
            libprotobuf-cpp-lite
            libutils
    )
endif ()
#################### keystore_cli_v2 ####################

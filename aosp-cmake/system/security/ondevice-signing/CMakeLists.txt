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

project(odsign)

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

#################### libsigningutils ####################
set(FILES
        ${ANDROID_ROOT}/system/security/ondevice-signing/CertUtils.cpp
        ${ANDROID_ROOT}/system/security/ondevice-signing/VerityUtils.cpp
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
    add_library(libsigningutils SHARED ${SOURCE_FILES})

    target_include_directories(libsigningutils
            PUBLIC
            ${ANDROID_ROOT}/system/security/ondevice-signing/include
            ${ANDROID_ROOT}/system/security/ondevice-signing
    )

    # Add c/c++ flags
    target_compile_options(libsigningutils PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libsigningutils
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libcrypto
            libcrypto_utils
            libdl
            libfsverity
            libm
            libprotobuf-cpp-lite
            libutils
    )
endif ()
#################### libsigningutils ####################

#################### odsign ####################
set(FILES
        ${ANDROID_ROOT}/system/security/ondevice-signing/KeystoreKey.cpp
        ${ANDROID_ROOT}/system/security/ondevice-signing/KeystoreHmacKey.cpp
        ${ANDROID_ROOT}/system/security/ondevice-signing/odsign_main.cpp
        ${ANDROID_ROOT}/system/security/ondevice-signing/StatsReporter.cpp
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
    add_executable(odsign ${SOURCE_FILES})

    target_include_directories(odsign
            PUBLIC
            ${ANDROID_ROOT}/system/security/ondevice-signing
            ${ANDROID_ROOT}/system/security/ondevice-signing/include
            ${OUT_INTERMEDIATES_ROOT}/system/security/ondevice-signing/proto/lib_odsign_proto/android_${OUT_ARCH_CPU}_static_cfi/gen/proto/system/security/ondevice-signing/proto
            ${OUT_INTERMEDIATES_ROOT}/system/security/ondevice-signing/proto/lib_odsign_proto/android_${OUT_ARCH_CPU}_static_cfi/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/security/ondevice-signing/statslog_odsign.h/gen
    )

    # Add c/c++ flags
    target_compile_options(odsign PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(odsign
            android.hardware.security.keymint-V1-cpp
            android.system.keystore2-V1-cpp
            lib_odsign_proto
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libcrypto
            libcrypto_utils
            libdl
            libfsverity
            liblogwrap
            libm
            libprotobuf-cpp-lite
            libsigningutils
            libutils
            odrefresh_headers
    )
endif ()
#################### odsign ####################

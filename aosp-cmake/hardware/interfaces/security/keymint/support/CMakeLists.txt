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

project(libkeymint_support)

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

#################### libkeymint_support ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/security/keymint/support/attestation_record.cpp
        ${ANDROID_ROOT}/hardware/interfaces/security/keymint/support/authorization_set.cpp
        ${ANDROID_ROOT}/hardware/interfaces/security/keymint/support/keymint_utils.cpp
        ${ANDROID_ROOT}/hardware/interfaces/security/keymint/support/key_param_output.cpp
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
    add_library(libkeymint_support SHARED ${SOURCE_FILES})

    target_include_directories(libkeymint_support
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/security/keymint/support/include
            ${ANDROID_ROOT}/hardware/interfaces/security/keymint/support
    )

    # Add c/c++ flags
    target_compile_options(libkeymint_support PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libkeymint_support
            android.hardware.security.keymint-V3-ndk
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            libhardware
            libm
            libutils
    )
endif ()
#################### libkeymint_support ####################

#################### libkeymint_remote_prov_support ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/security/keymint/support/remote_prov_utils.cpp
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
    add_library(libkeymint_remote_prov_support SHARED ${SOURCE_FILES})

    target_include_directories(libkeymint_remote_prov_support
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/security/keymint/support/include
            ${ANDROID_ROOT}/hardware/interfaces/security/keymint/support
    )

    # Add c/c++ flags
    target_compile_options(libkeymint_remote_prov_support PRIVATE
            -Werror
    )

    target_link_libraries(libkeymint_remote_prov_support
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.rkp-V3-ndk
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcppbor_external
            libcppcose_rkp
            libcrypto
            libdl
            libhwtrust_cxx
            libjsoncpp
            libkeymaster_portable
            libm
    )
endif ()
#################### libkeymint_remote_prov_support ####################

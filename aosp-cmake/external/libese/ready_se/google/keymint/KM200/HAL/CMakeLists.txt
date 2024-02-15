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

project(android.hardware.security.keymint-service.strongbox)

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

#################### libjc_keymint ####################
set(FILES
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/CborConverter.cpp
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/JavacardKeyMintDevice.cpp
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/JavacardKeyMintOperation.cpp
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/JavacardRemotelyProvisionedComponentDevice.cpp
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/JavacardSecureElement.cpp
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/JavacardSharedSecret.cpp
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/keymint_utils.cpp
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
    add_library(libjc_keymint SHARED ${SOURCE_FILES})

    target_include_directories(libjc_keymint
            PUBLIC
            ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL
    )

    # Add c/c++ flags
    target_compile_options(libjc_keymint PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libjc_keymint
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.rkp-V3-ndk
            android.hardware.security.secureclock-V1-ndk
            android.hardware.security.sharedsecret-V1-ndk
            lib_android_keymaster_keymint_utils
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcppbor_external
            libcrypto
            libcutils
            libdl
            libjc_keymint_transport
            libkeymaster_messages
            libkeymaster_portable
            liblog
            libm
            libsoft_attestation_cert
    )
endif ()
#################### libjc_keymint ####################

#################### libjc_keymint_transport ####################
set(FILES
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/SocketTransport.cpp
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/OmapiTransport.cpp
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
    add_library(libjc_keymint_transport SHARED ${SOURCE_FILES})

    target_include_directories(libjc_keymint_transport
            PUBLIC
            ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL
    )

    # Add c/c++ flags
    target_compile_options(libjc_keymint_transport PRIVATE
            -Werror
    )

    target_link_libraries(libjc_keymint_transport
            android.hardware.security.keymint-V3-ndk
            android.se.omapi-V1-ndk
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libhardware
            liblog
            libm
    )
endif ()
#################### libjc_keymint_transport ####################

#################### android.hardware.security.keymint-service.strongbox ####################
set(FILES
        ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL/service.cpp
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
    add_executable(android.hardware.security.keymint-service.strongbox ${SOURCE_FILES})

    target_include_directories(android.hardware.security.keymint-service.strongbox
            PUBLIC
            ${ANDROID_ROOT}/external/libese/ready_se/google/keymint/KM200/HAL
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.security.keymint-service.strongbox PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(android.hardware.security.keymint-service.strongbox
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.rkp-V3-ndk
            android.hardware.security.sharedsecret-V1-ndk
            android.se.omapi-V1-ndk
            lib_android_keymaster_keymint_utils
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcppbor_external
            libcrypto
            libdl
            libjc_keymint
            libjc_keymint_transport
            libkeymaster_portable
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.security.keymint-service.strongbox ####################

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

project(credstore)

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

#################### credstore ####################
set(FILES
        ${ANDROID_ROOT}/system/security/identity/Credential.cpp
        ${ANDROID_ROOT}/system/security/identity/CredentialData.cpp
        ${ANDROID_ROOT}/system/security/identity/CredentialStore.cpp
        ${ANDROID_ROOT}/system/security/identity/CredentialStoreFactory.cpp
        ${ANDROID_ROOT}/system/security/identity/Session.cpp
        ${ANDROID_ROOT}/system/security/identity/Util.cpp
        ${ANDROID_ROOT}/system/security/identity/WritableCredential.cpp
        ${ANDROID_ROOT}/system/security/identity/main.cpp
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
    add_executable(credstore ${SOURCE_FILES})

    target_include_directories(credstore
            PUBLIC
            ${ANDROID_ROOT}/system/security/identity
            ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(credstore PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(credstore
            android.hardware.identity-V5-cpp
            android.hardware.identity-support-lib
            android.hardware.keymaster-V3-cpp
            android.hardware.keymaster-4.0
            android.hardware.security.keymint-V3-cpp
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.rkp-V3-cpp
            android.security.authorization-ndk
            android.security.rkp_aidl-cpp
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcppbor_external
            libcredstore_aidl
            libcrypto
            libdl
            libhidlbase
            libkeymaster4support
            libkeystore-attestation-application-id
            libm
            librkp_support
            libutils
            libutilscallstack
            libvintf
    )
endif ()
#################### credstore ####################

#################### libcredstore_aidl ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/ICredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/IWritableCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/ICredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/AccessControlProfileParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/EntryNamespaceParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/EntryParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/RequestNamespaceParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/RequestEntryParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/ResultNamespaceParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/ResultEntryParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/GetEntriesResultParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/AuthKeyParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/SecurityHardwareInfoParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/ICredentialStoreFactory.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl/android/security/identity/ISession.cpp
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
    add_library(libcredstore_aidl SHARED ${SOURCE_FILES})

    target_include_directories(libcredstore_aidl
            PUBLIC
            ${ANDROID_ROOT}/system/security/identity
            ${OUT_INTERMEDIATES_ROOT}/system/security/identity/libcredstore_aidl/android_${OUT_ARCH_CPU}_shared_cfi/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libcredstore_aidl PRIVATE
            -Werror
    )

    target_link_libraries(libcredstore_aidl
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libkeymaster4support
            libm
            libutils
    )
endif ()
#################### libcredstore_aidl ####################

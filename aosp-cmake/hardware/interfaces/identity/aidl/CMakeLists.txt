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

project(android.hardware.identity-V1-ndk_platform)

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

#################### android.hardware.identity-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk_platform-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk_platform-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk_platform-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk_platform-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk_platform-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk_platform-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk_platform-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V1-ndk_platform
            android.hardware.keymaster-V3-ndk_platform
            android.hardware.security.rkp-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V1-ndk_platform ####################

#################### android.hardware.identity-V6-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/IPresentationSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V6-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V6-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V6-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V6-ndk_platform
            android.hardware.keymaster-V3-ndk_platform
            android.hardware.security.rkp-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V6-ndk_platform ####################

#################### android.hardware.identity-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V3-ndk_platform
            android.hardware.keymaster-V3-ndk_platform
            android.hardware.security.rkp-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V3-ndk_platform ####################

#################### android.hardware.identity-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V3-ndk
            android.hardware.keymaster-V3-ndk
            android.hardware.security.rkp-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V3-ndk ####################

#################### android.hardware.identity-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-cpp-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-cpp-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-cpp-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-cpp-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-cpp-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-cpp-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-cpp-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V1-cpp
            android.hardware.keymaster-V3-cpp
            android.hardware.security.rkp-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.identity-V1-cpp ####################

#################### android.hardware.identity-V6-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/IPresentationSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V6-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V6-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V6-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V6-ndk
            android.hardware.keymaster-V3-ndk
            android.hardware.security.rkp-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V6-ndk ####################

#################### android.hardware.identity-V4-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/IPresentationSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V4-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V4-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V4-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V4-cpp
            android.hardware.keymaster-V3-cpp
            android.hardware.security.rkp-V3-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.identity-V4-cpp ####################

#################### android.hardware.identity-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V2-ndk
            android.hardware.keymaster-V3-ndk
            android.hardware.security.rkp-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V2-ndk ####################

#################### android.hardware.identity-V5-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/IPresentationSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V5-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V5-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V5-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V5-cpp
            android.hardware.keymaster-V3-cpp
            android.hardware.security.rkp-V3-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.identity-V5-cpp ####################

#################### android.hardware.identity-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/IPresentationSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V4-ndk_platform
            android.hardware.keymaster-V3-ndk_platform
            android.hardware.security.rkp-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V4-ndk_platform ####################

#################### android.hardware.identity-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V1-ndk
            android.hardware.keymaster-V3-ndk
            android.hardware.security.rkp-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V1-ndk ####################

#################### android.hardware.identity-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V3-cpp
            android.hardware.keymaster-V3-cpp
            android.hardware.security.rkp-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.identity-V3-cpp ####################

#################### android.hardware.identity-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/IPresentationSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V4-ndk
            android.hardware.keymaster-V3-ndk
            android.hardware.security.rkp-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V4-ndk ####################

#################### android.hardware.identity-V5-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/IPresentationSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V5-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V5-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V5-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V5-ndk_platform
            android.hardware.keymaster-V3-ndk_platform
            android.hardware.security.rkp-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V5-ndk_platform ####################

#################### android.hardware.identity-V6-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/IPresentationSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V6-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V6-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V6-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V6-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V6-cpp
            android.hardware.keymaster-V3-cpp
            android.hardware.security.rkp-V3-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.identity-V6-cpp ####################

#################### android.hardware.identity-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V2-ndk_platform
            android.hardware.keymaster-V3-ndk_platform
            android.hardware.security.rkp-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V2-ndk_platform ####################

#################### android.hardware.identity-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V2-cpp
            android.hardware.keymaster-V3-cpp
            android.hardware.security.rkp-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.identity-V2-cpp ####################

#################### android.hardware.identity-V5-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/Certificate.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/CipherSuite.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/HardwareInformation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/IIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/IIdentityCredentialStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/IPresentationSession.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/IWritableIdentityCredential.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/RequestDataItem.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/RequestNamespace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/android/hardware/identity/SecureAccessControlProfile.cpp
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
    add_library(android.hardware.identity-V5-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.identity-V5-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/identity/aidl/android.hardware.identity-V5-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.identity-V5-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.identity-V5-ndk
            android.hardware.keymaster-V3-ndk
            android.hardware.security.rkp-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.identity-V5-ndk ####################

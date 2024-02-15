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

project(android.security.legacykeystore-cpp-analyzer)

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

#################### android.security.apc-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-cpp-analyzer-source/gen/android/security/apc/IConfirmationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-cpp-analyzer-source/gen/android/security/apc/IProtectedConfirmation.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-cpp-analyzer-source/gen/android/security/apc/ResponseCode.cpp
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
    add_executable(android.security.apc-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.security.apc-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/system/security/keystore2/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.apc-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.apc-cpp-analyzer
            aidl-analyzer-main
            android.security.apc-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### android.security.apc-cpp-analyzer ####################

#################### android.security.metrics-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/Algorithm.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/AtomID.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/CrashStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/EcCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/HardwareAuthenticatorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/IKeystoreMetrics.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/KeyCreationWithAuthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/KeyCreationWithGeneralInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/KeyCreationWithPurposeAndModesInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/KeyOperationWithGeneralInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/KeyOperationWithPurposeAndModesInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/KeyOrigin.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/Keystore2AtomWithOverflow.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/KeystoreAtom.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/KeystoreAtomPayload.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/Outcome.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/Purpose.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/RkpError.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/RkpErrorStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/SecurityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/Storage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/android/security/metrics/StorageStats.cpp
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
    add_library(android.security.metrics-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.security.metrics-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.metrics-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.metrics-ndk_platform
            android.system.keystore2-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.metrics-ndk_platform ####################

#################### android.security.legacykeystore-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.legacykeystore-ndk_platform-source/gen/android/security/legacykeystore/ILegacyKeystore.cpp
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
    add_library(android.security.legacykeystore-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.security.legacykeystore-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.legacykeystore-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.legacykeystore-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.legacykeystore-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.legacykeystore-ndk_platform ####################

#################### android.security.apc-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-ndk-source/gen/android/security/apc/IConfirmationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-ndk-source/gen/android/security/apc/IProtectedConfirmation.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-ndk-source/gen/android/security/apc/ResponseCode.cpp
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
    add_library(android.security.apc-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.security.apc-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.apc-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.apc-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.apc-ndk ####################

#################### android.security.compat-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.compat-ndk-source/gen/android/security/compat/IKeystoreCompatService.cpp
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
    add_library(android.security.compat-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.security.compat-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.compat-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.compat-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.compat-ndk
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.secureclock-V1-ndk
            android.hardware.security.sharedsecret-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.compat-ndk ####################

#################### android.security.apc-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-cpp-source/gen/android/security/apc/IConfirmationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-cpp-source/gen/android/security/apc/IProtectedConfirmation.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-cpp-source/gen/android/security/apc/ResponseCode.cpp
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
    add_library(android.security.apc-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.security.apc-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.apc-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.apc-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.security.apc-cpp ####################

#################### android.security.apc-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-ndk_platform-source/gen/android/security/apc/IConfirmationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-ndk_platform-source/gen/android/security/apc/IProtectedConfirmation.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-ndk_platform-source/gen/android/security/apc/ResponseCode.cpp
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
    add_library(android.security.apc-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.security.apc-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.apc-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.apc-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.apc-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.apc-ndk_platform ####################

#################### android.security.legacykeystore-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.legacykeystore-cpp-analyzer-source/gen/android/security/legacykeystore/ILegacyKeystore.cpp
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
    add_executable(android.security.legacykeystore-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.security.legacykeystore-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/system/security/keystore2/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.legacykeystore-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.legacykeystore-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.legacykeystore-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.legacykeystore-cpp-analyzer
            aidl-analyzer-main
            android.security.legacykeystore-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### android.security.legacykeystore-cpp-analyzer ####################

#################### android.security.compat-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.compat-cpp-source/gen/android/security/compat/IKeystoreCompatService.cpp
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
    add_library(android.security.compat-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.security.compat-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.compat-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.compat-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.compat-cpp
            android.hardware.security.keymint-V3-cpp
            android.hardware.security.secureclock-V1-cpp
            android.hardware.security.sharedsecret-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.security.compat-cpp ####################

#################### android.security.authorization-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk_platform-source/gen/android/security/authorization/AuthorizationTokens.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk_platform-source/gen/android/security/authorization/IKeystoreAuthorization.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk_platform-source/gen/android/security/authorization/LockScreenEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk_platform-source/gen/android/security/authorization/ResponseCode.cpp
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
    add_library(android.security.authorization-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.security.authorization-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.authorization-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.authorization-ndk_platform
            android.hardware.security.keymint-V3-ndk_platform
            android.hardware.security.secureclock-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.authorization-ndk_platform ####################

#################### android.security.metrics-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/Algorithm.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/AtomID.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/CrashStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/EcCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/HardwareAuthenticatorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/IKeystoreMetrics.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/KeyCreationWithAuthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/KeyCreationWithGeneralInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/KeyCreationWithPurposeAndModesInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/KeyOperationWithGeneralInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/KeyOperationWithPurposeAndModesInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/KeyOrigin.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/Keystore2AtomWithOverflow.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/KeystoreAtom.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/KeystoreAtomPayload.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/Outcome.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/Purpose.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/RkpError.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/RkpErrorStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/SecurityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/Storage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/android/security/metrics/StorageStats.cpp
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
    add_library(android.security.metrics-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.security.metrics-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.metrics-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.metrics-cpp
            android.system.keystore2-V3-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.security.metrics-cpp ####################

#################### android.security.authorization-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-cpp-source/gen/android/security/authorization/AuthorizationTokens.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-cpp-source/gen/android/security/authorization/IKeystoreAuthorization.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-cpp-source/gen/android/security/authorization/LockScreenEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-cpp-source/gen/android/security/authorization/ResponseCode.cpp
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
    add_library(android.security.authorization-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.security.authorization-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.authorization-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.authorization-cpp
            android.hardware.security.keymint-V3-cpp
            android.hardware.security.secureclock-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.security.authorization-cpp ####################

#################### android.security.maintenance-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.maintenance-ndk-source/gen/android/security/maintenance/IKeystoreMaintenance.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.maintenance-ndk-source/gen/android/security/maintenance/UserState.cpp
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
    add_library(android.security.maintenance-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.security.maintenance-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.maintenance-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.maintenance-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.maintenance-ndk
            android.system.keystore2-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.maintenance-ndk ####################

#################### android.security.legacykeystore-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.legacykeystore-cpp-source/gen/android/security/legacykeystore/ILegacyKeystore.cpp
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
    add_library(android.security.legacykeystore-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.security.legacykeystore-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.legacykeystore-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.legacykeystore-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.legacykeystore-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.security.legacykeystore-cpp ####################

#################### android.security.compat-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.compat-ndk_platform-source/gen/android/security/compat/IKeystoreCompatService.cpp
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
    add_library(android.security.compat-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.security.compat-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.compat-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.compat-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.compat-ndk_platform
            android.hardware.security.keymint-V3-ndk_platform
            android.hardware.security.secureclock-V1-ndk_platform
            android.hardware.security.sharedsecret-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.compat-ndk_platform ####################

#################### android.security.metrics-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/Algorithm.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/AtomID.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/CrashStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/EcCurve.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/HardwareAuthenticatorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/IKeystoreMetrics.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/KeyCreationWithAuthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/KeyCreationWithGeneralInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/KeyCreationWithPurposeAndModesInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/KeyOperationWithGeneralInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/KeyOperationWithPurposeAndModesInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/KeyOrigin.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/Keystore2AtomWithOverflow.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/KeystoreAtom.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/KeystoreAtomPayload.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/Outcome.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/Purpose.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/RkpError.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/RkpErrorStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/SecurityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/Storage.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/android/security/metrics/StorageStats.cpp
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
    add_library(android.security.metrics-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.security.metrics-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.metrics-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.metrics-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.metrics-ndk
            android.system.keystore2-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.metrics-ndk ####################

#################### android.security.maintenance-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.maintenance-ndk_platform-source/gen/android/security/maintenance/IKeystoreMaintenance.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.maintenance-ndk_platform-source/gen/android/security/maintenance/UserState.cpp
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
    add_library(android.security.maintenance-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.security.maintenance-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.maintenance-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.maintenance-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.maintenance-ndk_platform
            android.system.keystore2-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.maintenance-ndk_platform ####################

#################### android.security.maintenance-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.maintenance-cpp-source/gen/android/security/maintenance/IKeystoreMaintenance.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.maintenance-cpp-source/gen/android/security/maintenance/UserState.cpp
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
    add_library(android.security.maintenance-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.security.maintenance-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.maintenance-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.maintenance-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.maintenance-cpp
            android.system.keystore2-V3-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.security.maintenance-cpp ####################

#################### android.security.legacykeystore-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.legacykeystore-ndk-source/gen/android/security/legacykeystore/ILegacyKeystore.cpp
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
    add_library(android.security.legacykeystore-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.security.legacykeystore-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.legacykeystore-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.legacykeystore-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.legacykeystore-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.legacykeystore-ndk ####################

#################### android.security.authorization-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk-source/gen/android/security/authorization/AuthorizationTokens.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk-source/gen/android/security/authorization/IKeystoreAuthorization.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk-source/gen/android/security/authorization/LockScreenEvent.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk-source/gen/android/security/authorization/ResponseCode.cpp
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
    add_library(android.security.authorization-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.security.authorization-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/system/security/keystore2/aidl/android.security.authorization-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.authorization-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.authorization-ndk
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.secureclock-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.authorization-ndk ####################

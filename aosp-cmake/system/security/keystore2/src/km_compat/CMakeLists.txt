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

project(libkm_compat_service)

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

#################### libkm_compat_service ####################
set(FILES
        ${ANDROID_ROOT}/system/security/keystore2/src/km_compat/km_compat_service.cpp
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
    add_library(libkm_compat_service SHARED ${SOURCE_FILES})

    target_include_directories(libkm_compat_service
            PUBLIC
            ${ANDROID_ROOT}/system/security/keystore2/src/km_compat
    )

    # Add c/c++ flags
    target_compile_options(libkm_compat_service PRIVATE
            -Werror
    )

    target_link_libraries(libkm_compat_service
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.secureclock-V1-ndk
            android.hardware.security.sharedsecret-V1-ndk
            android.security.compat-ndk
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libkeymaster4_1support
            libkeystore2_crypto
            libkm_compat
            libm
    )
endif ()
#################### libkm_compat_service ####################

#################### libkm_compat ####################
set(FILES
        ${ANDROID_ROOT}/system/security/keystore2/src/km_compat/km_compat.cpp
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
    add_library(libkm_compat SHARED ${SOURCE_FILES})

    target_include_directories(libkm_compat
            PUBLIC
            ${ANDROID_ROOT}/system/security/keystore2/src/km_compat
    )

    # Add c/c++ flags
    target_compile_options(libkm_compat PRIVATE
            -Werror
    )

    target_link_libraries(libkm_compat
            android.hardware.keymaster-3.0
            android.hardware.keymaster-4.0
            android.hardware.keymaster-4.1
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.secureclock-V1-ndk
            android.hardware.security.sharedsecret-V1-ndk
            android.security.compat-ndk
            android.system.keystore2-V3-ndk
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libhidlbase
            libkeymaster4_1support
            libkeymint
            libkeymint_support
            libkeystore2_crypto
            libm
            libutils
    )
endif ()
#################### libkm_compat ####################

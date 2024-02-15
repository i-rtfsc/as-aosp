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

project(libkeymaster4support)

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

#################### libkeymaster4support ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/keymaster/4.0/support/attestation_record.cpp
        ${ANDROID_ROOT}/hardware/interfaces/keymaster/4.0/support/authorization_set.cpp
        ${ANDROID_ROOT}/hardware/interfaces/keymaster/4.0/support/key_param_output.cpp
        ${ANDROID_ROOT}/hardware/interfaces/keymaster/4.0/support/keymaster_utils.cpp
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
    add_library(libkeymaster4support SHARED ${SOURCE_FILES})

    target_include_directories(libkeymaster4support
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/keymaster/4.0/support/include
            ${ANDROID_ROOT}/hardware/interfaces/keymaster/4.0/support
    )

    # Add c/c++ flags
    target_compile_options(libkeymaster4support PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libkeymaster4support
            android.hardware.keymaster-3.0
            android.hardware.keymaster-4.0
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
            libhidlbase
            libm
    )
endif ()
#################### libkeymaster4support ####################

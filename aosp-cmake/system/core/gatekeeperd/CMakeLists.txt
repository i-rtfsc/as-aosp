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

project(gatekeeperd)

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

#################### libgatekeeper_aidl ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/core/gatekeeperd/libgatekeeper_aidl/android_${OUT_ARCH_CPU}_shared/gen/aidl/android/service/gatekeeper/IGateKeeperService.cpp
        ${ANDROID_ROOT}/system/core/gatekeeperd/GateKeeperResponse.cpp
        ${ANDROID_ROOT}/system/core/gatekeeperd/GateKeeperPerfLock.cpp
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
    add_library(libgatekeeper_aidl SHARED ${SOURCE_FILES})

    target_include_directories(libgatekeeper_aidl
            PUBLIC
            ${ANDROID_ROOT}/system/core/gatekeeperd/include
            ${ANDROID_ROOT}/system/core/gatekeeperd
            ${OUT_INTERMEDIATES_ROOT}/system/core/gatekeeperd/libgatekeeper_aidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libgatekeeper_aidl PRIVATE
            -Werror
    )

    target_link_libraries(libgatekeeper_aidl
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libgatekeeper_aidl ####################

#################### gatekeeperd ####################
set(FILES
        ${ANDROID_ROOT}/system/core/gatekeeperd/gatekeeperd.cpp
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
    add_executable(gatekeeperd ${SOURCE_FILES})

    target_include_directories(gatekeeperd
            PUBLIC
            ${ANDROID_ROOT}/system/core/gatekeeperd
            ${ANDROID_ROOT}/system/core/gatekeeperd/include
            ${OUT_INTERMEDIATES_ROOT}/system/core/gatekeeperd/libgatekeeper_aidl/android_${OUT_ARCH_CPU}_shared/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(gatekeeperd PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(gatekeeperd
            android.hardware.gatekeeper-V1-ndk
            android.hardware.gatekeeper-1.0
            android.hardware.security.keymint-V3-ndk
            android.security.authorization-ndk
            lib_android_keymaster_keymint_utils
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libgatekeeper
            libgatekeeper_aidl
            libgsi
            libhardware
            libhidlbase
            liblog
            libm
            libscrypt_static
            libutils
    )
endif ()
#################### gatekeeperd ####################

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

project(android.hardware.biometrics.fingerprint-service.example)

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

#################### libandroid.hardware.biometrics.fingerprint.VirtualProps ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/libandroid.hardware.biometrics.fingerprint.VirtualProps/android_vendor.34_${OUT_ARCH_CPU}_static/gen/sysprop/fingerprint.sysprop.cpp
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
    add_library(libandroid.hardware.biometrics.fingerprint.VirtualProps SHARED ${SOURCE_FILES})

    target_include_directories(libandroid.hardware.biometrics.fingerprint.VirtualProps
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/libandroid.hardware.biometrics.fingerprint.VirtualProps/android_vendor.34_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libandroid.hardware.biometrics.fingerprint.VirtualProps PRIVATE
            -Werror
    )

    target_link_libraries(libandroid.hardware.biometrics.fingerprint.VirtualProps
            libbase_headers
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libandroid.hardware.biometrics.fingerprint.VirtualProps ####################

#################### android.hardware.biometrics.fingerprint-service.example ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/FakeLockoutTracker.cpp
        ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/FakeFingerprintEngine.cpp
        ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/FakeFingerprintEngineRear.cpp
        ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/FakeFingerprintEngineUdfps.cpp
        ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/FakeFingerprintEngineSide.cpp
        ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/Fingerprint.cpp
        ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/Session.cpp
        ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/main.cpp
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
    add_executable(android.hardware.biometrics.fingerprint-service.example ${SOURCE_FILES})

    target_include_directories(android.hardware.biometrics.fingerprint-service.example
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/include
            ${ANDROID_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/biometrics/fingerprint/aidl/default/libandroid.hardware.biometrics.fingerprint.VirtualProps/android_vendor.34_${OUT_ARCH_CPU}_static_cfi/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.biometrics.fingerprint-service.example PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.biometrics.fingerprint-service.example
            android.hardware.biometrics.common-V3-ndk
            android.hardware.biometrics.common.thread
            android.hardware.biometrics.common.util
            android.hardware.biometrics.fingerprint-V3-ndk
            libandroid.hardware.biometrics.fingerprint.VirtualProps
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.biometrics.fingerprint-service.example ####################

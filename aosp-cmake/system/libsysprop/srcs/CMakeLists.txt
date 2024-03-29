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

project(libPlatformProperties)

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

#################### libPlatformProperties ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/AdbProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/ApkVerityProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/BluetoothProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/CarProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/ContactsProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/CryptoProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/DeviceProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/DisplayProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/HdmiProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/HypervisorProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/InputProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/MediaProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/MemoryProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/NetworkProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/NfcProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/OtaProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/PowerProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/SetupWizardProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/SocProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/TelephonyProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/TraceProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/VndkProperties.sysprop.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/android/sysprop/WifiProperties.sysprop.cpp
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
    add_library(libPlatformProperties SHARED ${SOURCE_FILES})

    target_include_directories(libPlatformProperties
            PUBLIC
            ${ANDROID_ROOT}/system/libsysprop/srcs
            ${OUT_INTERMEDIATES_ROOT}/system/libsysprop/srcs/libPlatformProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libPlatformProperties PRIVATE
            -Werror
    )

    target_link_libraries(libPlatformProperties
            libbase
            libbase_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libPlatformProperties ####################

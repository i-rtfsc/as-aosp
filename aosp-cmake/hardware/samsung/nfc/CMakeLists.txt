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

project(nfc_nci_samsung)

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

#################### nfc_nci_samsung ####################
set(FILES
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/osi/osi_main.cc
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/osi/osi_memory.cc
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/osi/osi_task.cc
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/osi/osi_timer.cc
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/src/device.cc
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/src/hal.cc
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/src/hal_nci.cc
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/src/hal_task.cc
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/src/util.cc
        ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/src/config.cpp
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
    add_library(nfc_nci_samsung SHARED ${SOURCE_FILES})

    target_include_directories(nfc_nci_samsung
            PUBLIC
            ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/include
            ${ANDROID_ROOT}/hardware/samsung/nfc/halimpl/osi
            ${ANDROID_ROOT}/hardware/samsung/nfc
    )

    # Add c/c++ flags
    target_compile_options(nfc_nci_samsung PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nfc_nci_samsung
            android.hardware.nfc-1.0
            android.hardware.nfc-1.1
            android.hardware.nfc-1.2
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libhardware
            libhardware_legacy
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### nfc_nci_samsung ####################

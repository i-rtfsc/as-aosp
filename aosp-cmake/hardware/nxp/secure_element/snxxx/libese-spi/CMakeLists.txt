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

project(ese_spi_nxp_snxxx)

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

#################### ese_spi_nxp_snxxx ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/lib/phNxpEseDataMgr.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/lib/phNxpEseProto7816_3.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/lib/phNxpEse_Apdu_Api.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/lib/phNxpEse_Api.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/pal/phNxpEsePal.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/pal/EseTransportFactory.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/pal/spi/EseSpiTransport.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/pal/NxpTimer.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/spm/phNxpEse_Spm.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/utils/ese_config.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/utils/config.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/utils/ringbuffer.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/src/adaptation/NfcAdaptation.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/src/adaptation/CondVar.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/src/adaptation/Mutex.cpp
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
    add_library(ese_spi_nxp_snxxx SHARED ${SOURCE_FILES})

    target_include_directories(ese_spi_nxp_snxxx
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/common/include
            ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/common
            ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/inc
            ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/pal
            ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/src/include
            ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/lib
            ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/pal/spi
            ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi/p73/utils
            ${ANDROID_ROOT}/hardware/nxp/secure_element/snxxx/libese-spi
    )

    # Add c/c++ flags
    target_compile_options(ese_spi_nxp_snxxx PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(ese_spi_nxp_snxxx
            android.hardware.nfc-1.0
            android.hardware.nfc-1.1
            android.hardware.secure_element-1.0
            libbase
            libbinder
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libchrome
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            vendor.nxp.nxpese-1.0
            vendor.nxp.nxpnfc-2.0
            vendor.nxp.nxpnfc_aidl-V1-ndk
            libc
            libdl
            libm
    )
endif ()
#################### ese_spi_nxp_snxxx ####################

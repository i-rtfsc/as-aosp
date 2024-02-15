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

project(android.hardware.nfc@1.1-service)

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

#################### android.hardware.nfc-1.2-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/extns/impl/NxpNfc.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/1.2/NxpNfcService.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/1.2/Nfc.cpp
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
    add_executable(android.hardware.nfc-1.2-service ${SOURCE_FILES})

    target_include_directories(android.hardware.nfc-1.2-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/extns/impl
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/common
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/inc
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.nfc-1.2-service PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.nfc-1.2-service
            android.hardware.nfc-1.0
            android.hardware.nfc-1.1
            android.hardware.nfc-1.2
            ese_spi_nxp
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            nfc_nci_nxp
            vendor.nxp.nxpese-1.0
            vendor.nxp.nxpnfc-1.0
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.nfc-1.2-service ####################

#################### android.hardware.nfc-1.1-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/extns/impl/NxpNfc.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/1.1/NxpNfcService.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/1.1/Nfc.cpp
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
    add_executable(android.hardware.nfc-1.1-service ${SOURCE_FILES})

    target_include_directories(android.hardware.nfc-1.1-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/extns/impl
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/common
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/inc
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.nfc-1.1-service PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.nfc-1.1-service
            android.hardware.nfc-1.0
            android.hardware.nfc-1.1
            android.hardware.nfc-1.2
            ese_spi_nxp
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            nfc_nci_nxp
            vendor.nxp.nxpese-1.0
            vendor.nxp.nxpnfc-1.0
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.nfc-1.1-service ####################

#################### nfc_nci_nxp ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/dnld/phDnldNfc.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/dnld/phDnldNfc_Internal.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/dnld/phDnldNfc_Utils.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/dnld/phNxpNciHal_Dnld.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/hal/phNxpNciHal.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/hal/phNxpNciHal_NfcDepSWPrio.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/hal/phNxpNciHal_dta.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/hal/phNxpNciHal_ext.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/log/phNxpLog.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/self-test/phNxpNciHal_SelfTest.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/src/adaptation/EseAdaptation.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/tml/phDal4Nfc_messageQueueLib.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/tml/phOsalNfc_Timer.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/tml/phTmlNfc.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/tml/phTmlNfc_i2c.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/tml/spi_spm.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/utils/NxpNfcCapability.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/utils/phNxpConfig.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/utils/phNxpNciHal_utils.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/utils/sparse_crc32.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/utils/NfccPowerTracker.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/mifare/NxpMfcReader.cc
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
    add_library(nfc_nci_nxp SHARED ${SOURCE_FILES})

    target_include_directories(nfc_nci_nxp
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/extns/impl
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/common
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/inc
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/dnld
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/hal
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/log
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/self-test
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/src/include
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/tml
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/utils
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x/halimpl/mifare
            ${ANDROID_ROOT}/hardware/nxp/nfc/pn8x
    )

    # Add c/c++ flags
    target_compile_options(nfc_nci_nxp PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(nfc_nci_nxp
            android.hardware.nfc-1.0
            android.hardware.nfc-1.1
            android.hardware.nfc-1.2
            android.hardware.secure_element-1.0
            ese_spi_nxp
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhardware_legacy
            libhidlbase
            liblog
            libm
            libutils
            vendor.nxp.nxpese-1.0
            libc
            libdl
            libm
    )
endif ()
#################### nfc_nci_nxp ####################

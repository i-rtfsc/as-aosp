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

project(android.hardware.nfc-service.nxp)

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

#################### nfc_nci_nxp_snxxx ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/dnld/phDnldNfc.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/dnld/phDnldNfc_Internal.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/dnld/phDnldNfc_Utils.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/dnld/phNxpNciHal_Dnld.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal/phNxpNciHal.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal/phNxpNciHal_NfcDepSWPrio.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal/phNxpNciHal_dta.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal/phNxpNciHal_ext.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal/phNxpNciHal_nciParser.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/log/phNxpLog.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/self-test/phNxpNciHal_SelfTest.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/src/adaptation/EseAdaptation.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/tml/phDal4Nfc_messageQueueLib.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/tml/phOsalNfc_Timer.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/tml/phTmlNfc.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/tml/NfccTransportFactory.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/tml/transport/NfccI2cTransport.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/tml/transport/NfccTransport.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/utils/NxpNfcCapability.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/utils/NxpNfcThreadMutex.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/utils/phNxpConfig.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/utils/phNxpNciHal_utils.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/utils/sparse_crc32.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/utils/IntervalTimer.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/eseclients_extns/src/eSEClientExtns.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal/phNxpNciHal_IoctlOperations.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal/phNxpNciHal_extOperations.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal/phNxpNciHal_ULPDet.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/mifare/NxpMfcReader.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/recovery/phNxpNciHal_Recovery.cc
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal/phNxpNciHal_PowerTrackerIface.cc
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
    add_library(nfc_nci_nxp_snxxx SHARED ${SOURCE_FILES})

    target_include_directories(nfc_nci_nxp_snxxx
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/common
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/dnld
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/inc
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/log
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/self-test
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/tml/transport
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/tml
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/utils
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/eseclients_extns/inc
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/src/include
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/mifare
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/recovery
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/power-tracker/include
    )

    # Add c/c++ flags
    target_compile_options(nfc_nci_nxp_snxxx PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(nfc_nci_nxp_snxxx
            libbase
            libbinder_ndk
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
            power_tracker_headers
            vendor.nxp.nxpese-1.0
            libc
            libdl
            libm
    )
endif ()
#################### nfc_nci_nxp_snxxx ####################

#################### android.hardware.nfc-service.nxp ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/aidl/main.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/aidl/Nfc.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/aidl/NfcExtns.cpp
        ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/extns/impl/nxpnfc/aidl/NxpNfc.cpp
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
    add_executable(android.hardware.nfc-service.nxp ${SOURCE_FILES})

    target_include_directories(android.hardware.nfc-service.nxp
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/common
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/dnld
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/hal
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/inc
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/log
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/utils
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/mifare
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/recovery
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/halimpl/eseclients_extns/inc
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx/extns/impl/nxpnfc/aidl
            ${ANDROID_ROOT}/hardware/nxp/nfc/snxxx
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.nfc-service.nxp PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(android.hardware.nfc-service.nxp
            android.hardware.nfc-V1-ndk
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libhardware
            liblog
            libm
            libutils
            nfc_nci_nxp_snxxx
            vendor.nxp.nxpnfc_aidl-V1-ndk
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.nfc-service.nxp ####################

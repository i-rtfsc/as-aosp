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

project(libsn100nfc_nci_jni)

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

#################### libsn100nfc_nci_jni ####################
set(FILES
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/CondVar.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/DataQueue.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/DwpChannel.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/HciEventManager.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/IntervalTimer.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/MposManager.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/Mutex.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeExtFieldDetect.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeExtFieldDetectManager.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeJniExtns.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeLlcpConnectionlessSocket.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeLlcpServiceSocket.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeLlcpSocket.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeMposManager.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeNfcManager.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeNfcTag.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeNxpFeature.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeP2pDevice.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeSecureElement.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeT4tNfcee.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NativeT4tNfceeManager.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NfcDta.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NfcJniUtil.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NfcSelfTest.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NfcStatsUtil.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NfcTag.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/NfcTagExtns.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/PeerToPeer.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/PowerSwitch.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/RouteDataSet.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/RoutingManager.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/SecureDigitization.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/SecureElement.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/extns/pn54x/src/phNxpExtns.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/extns/pn54x/src/mifare/phFriNfc_MifStdFormat.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/extns/pn54x/src/mifare/phFriNfc_MifareStdMap.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/extns/pn54x/src/mifare/phFriNfc_SmtCrdFmt.cpp
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/extns/pn54x/src/mifare/phNxpExtns_MifareStd.cpp
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
    add_library(libsn100nfc_nci_jni SHARED ${SOURCE_FILES})

    target_include_directories(libsn100nfc_nci_jni
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/extns/pn54x/inc
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/extns/pn54x/src/common
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni/extns/pn54x/src/mifare
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/packages/apps/Nfc/nci/SN100x/jni
    )

    # Add c/c++ flags
    target_compile_options(libsn100nfc_nci_jni PRIVATE
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libsn100nfc_nci_jni
            libandroidicu
            libbase
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libnativehelper
            libsn100nfc-nci
            libstatslog_nfc
            libutils
            libxml2
            vendor.nxp.hardware.nfc-2.0
    )
endif ()
#################### libsn100nfc_nci_jni ####################

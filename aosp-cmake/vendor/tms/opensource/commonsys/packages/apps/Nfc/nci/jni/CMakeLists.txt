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

project(libtmsnfc_nci_jni)

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

#################### libtmsnfc_nci_jni ####################
set(FILES
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/CondVar.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/DataQueue.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/HciEventManager.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/IntervalTimer.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/Mutex.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeLlcpConnectionlessSocket.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeLlcpServiceSocket.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeLlcpSocket.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeNfcManager.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeNfcTag.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeP2pDevice.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeT4tNfcee.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeT4tNfceeManager.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NfcDta.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NfcJniUtil.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NfcStatsUtil.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/NfcTag.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/PeerToPeer.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/PowerSwitch.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/RouteDataSet.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/RoutingManager.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/TmsNativeNfcConfig.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/TmsNativeNfcManager.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/TmsNativeWiredSe.cpp
        ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni/TmsSecureElement.cpp
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
    add_library(libtmsnfc_nci_jni SHARED ${SOURCE_FILES})

    target_include_directories(libtmsnfc_nci_jni
            PUBLIC
            ${ANDROID_ROOT}/vendor/tms/opensource/commonsys/packages/apps/Nfc/nci/jni
    )

    # Add c/c++ flags
    target_compile_options(libtmsnfc_nci_jni PRIVATE
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libtmsnfc_nci_jni
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
            libstatslog_nfc_tms
            libtmsnfc-nci
            libutils
            libxml2
    )
endif ()
#################### libtmsnfc_nci_jni ####################

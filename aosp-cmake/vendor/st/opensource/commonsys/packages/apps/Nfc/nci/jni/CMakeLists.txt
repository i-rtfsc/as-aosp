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

project(libstnfc_jni)

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

#################### libstnfc_jni ####################
set(FILES
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/CondVar.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/DataQueue.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/IntervalTimer.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeLlcpConnectionlessSocket.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeLlcpServiceSocket.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeLlcpSocket.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeP2pDevice.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/PeerToPeer.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/RouteDataSet.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/PowerSwitch.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NfcDta.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NfcStatsUtil.cpp
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
    add_library(libstnfc_jni SHARED ${SOURCE_FILES})

    target_include_directories(libstnfc_jni
            PUBLIC
            ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni
    )

    # Add c/c++ flags
    target_compile_options(libstnfc_jni PRIVATE
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libstnfc_jni
            libbase
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libnativehelper
            libstatslog_nfc_st
            libstnfc-nci
            libutils
            libxml2
    )
endif ()
#################### libstnfc_jni ####################

#################### nfc_jni_files_for_factory ####################
set(FILES
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/CondVar.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/DataQueue.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/IntervalTimer.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/Mutex.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeLlcpConnectionlessSocket.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeLlcpServiceSocket.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeLlcpSocket.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/NativeP2pDevice.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/PeerToPeer.cpp
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni/RouteDataSet.cpp
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
    add_library(nfc_jni_files_for_factory SHARED ${SOURCE_FILES})

    target_include_directories(nfc_jni_files_for_factory
            PUBLIC
            ${ANDROID_ROOT}/vendor/st/opensource/commonsys/packages/apps/Nfc/nci/jni
    )

    # Add c/c++ flags
    target_compile_options(nfc_jni_files_for_factory PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(nfc_jni_files_for_factory
            jni_headers
            libbase
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhardware_headers
            liblog
            libm
            libnativehelper
            libstnfc-nci
            libstnfcutils
            libutils
            libxml2
    )
endif ()
#################### nfc_jni_files_for_factory ####################

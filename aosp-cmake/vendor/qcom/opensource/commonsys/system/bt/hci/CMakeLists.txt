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

project(libbt-hci_qti)

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

#################### libbt-hci_qti ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/btsnoop.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/btsnoop_mem.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/btsnoop_net.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/buffer_allocator.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/hci_inject.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/hci_layer.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/hci_layer_android.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/hci_packet_factory.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/hci_packet_parser.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/src/packet_fragmenter.cc
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
    add_library(libbt-hci_qti SHARED ${SOURCE_FILES})

    target_include_directories(libbt-hci_qti
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci/include
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/hci
    )

    # Add c/c++ flags
    target_compile_options(libbt-hci_qti PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-hci_qti
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            libbluetooth-types
            libbluetooth_headers
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libm
    )
endif ()
#################### libbt-hci_qti ####################

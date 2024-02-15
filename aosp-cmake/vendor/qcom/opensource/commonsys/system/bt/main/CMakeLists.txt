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

project(libbluetooth_qti)

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

#################### libbluetooth_qti ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/main/bte_conf.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/main/bte_init.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/main/bte_init_cpp_logging.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/main/bte_logmsg.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/main/bte_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/main/stack_config.cc
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
    add_library(libbluetooth_qti SHARED ${SOURCE_FILES})

    target_include_directories(libbluetooth_qti
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/main
    )

    # Add c/c++ flags
    target_compile_options(libbluetooth_qti PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbluetooth_qti
            android.hardware.audio.common-V1-ndk
            android.hardware.bluetooth.audio-V3-ndk
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            com.qualcomm.qti.bluetooth_audio-1.0
            libFraunhoferAAC
            libaaudio
            libbinder_ndk
            libbluetooth-types
            libbluetooth_headers
            libbt-bta-ext
            libbt-bta_qti
            libbt-bta_qti_adva
            libbt-btif_qti_adva
            libbt-hci_qti
            libbt-protos_qti
            libbt-sbc-decoder
            libbt-sbc-encoder
            libbt-stack_ext
            libbt-stack_qti
            libbt-utils_qti
            libbtconfigstore
            libbtcore_qti
            libbtdevice_ext
            libbtdevice_qti
            libbtif_ext
            libbtif_qti
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libfmq
            libg722codec_qti
            libhidlbase
            liblog
            libm
            libosi_ext
            libosi_qti
            libprocessgroup
            libprotobuf-cpp-lite
            libtinyxml2
            libudrv-uipc_qti
            libutils
            libxml2
            libz
            vendor.qti.hardware.bluetooth.audio-V1-ndk
            vendor.qti.hardware.bluetooth_audio-2.0
            vendor.qti.hardware.bluetooth_audio-2.1
    )
endif ()
#################### libbluetooth_qti ####################

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

project(bluetooth_gd_l2cap_fuzzer)

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

#################### bluetooth_gd_l2cap_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/fuzz/helpers.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/fuzz/fuzz_hci_layer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/fuzz/hci_layer_fuzz_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/fuzz/status_vs_complete_commands.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/fuzz/channel_fuzz_controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/fuzz/fuzz_l2cap.cc
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
    add_executable(bluetooth_gd_l2cap_fuzzer ${SOURCE_FILES})

    target_include_directories(bluetooth_gd_l2cap_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/fuzz
    )

    # Add c/c++ flags
    target_compile_options(bluetooth_gd_l2cap_fuzzer PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -Wno-unused-result
    )

    target_link_libraries(bluetooth_gd_l2cap_fuzzer
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            android.system.suspend.control-V1-ndk
            jni_headers
            libbinder_ndk
            libbluetooth-protos
            libbluetooth_gd_fuzzing
            libbluetooth_rust_interop
            libbt_shim_bridge
            libbt_shim_ffi
            libbte
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libcutils
            libdl
            libflatbuffers-cpp
            libgmock
            libgrpc++
            libgrpc_wrap
            libgtest
            libhidlbase
            liblog
            libm
            libstatslog_bt
            libutils
    )
endif ()
#################### bluetooth_gd_l2cap_fuzzer ####################

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

project(btcore_device_class_fuzzer)

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

#################### btcore_property_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btcore/fuzzer/btcore_property_fuzzer.cpp
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
    add_executable(btcore_property_fuzzer ${SOURCE_FILES})

    target_include_directories(btcore_property_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btcore/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(btcore_property_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(btcore_property_fuzzer
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            libbase
            libbluetooth-protos
            libbluetooth-types
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-common
            libbt_shim_bridge
            libbt_shim_ffi
            libbtcore
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libdl
            libevent
            liblog
            libm
            libmodpb64
            libosi
            libprotobuf-cpp-lite
            libstatslog_bt
            libvndksupport
    )
endif ()
#################### btcore_property_fuzzer ####################

#################### btcore_device_class_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btcore/fuzzer/btcore_device_class_fuzzer.cpp
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
    add_executable(btcore_device_class_fuzzer ${SOURCE_FILES})

    target_include_directories(btcore_device_class_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/btcore/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(btcore_device_class_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(btcore_device_class_fuzzer
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            libbase
            libbluetooth-protos
            libbluetooth-types
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-common
            libbt_shim_bridge
            libbt_shim_ffi
            libbtcore
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcrypto
            libdl
            libevent
            liblog
            libm
            libmodpb64
            libosi
            libprotobuf-cpp-lite
            libstatslog_bt
            libvndksupport
    )
endif ()
#################### btcore_device_class_fuzzer ####################

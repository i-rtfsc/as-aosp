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

project(android.hardware.bluetooth-service.default)

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

#################### android.hardware.bluetooth-service.default_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/aidl/default/test/fuzzer.cpp
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
    add_executable(android.hardware.bluetooth-service.default_fuzzer ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth-service.default_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/bluetooth/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth-service.default_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.bluetooth-service.default_fuzzer
            android.hardware.bluetooth-V1-ndk
            android.hardware.bluetooth.async
            android.hardware.bluetooth.hci
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libbluetoothhcihalimpl
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### android.hardware.bluetooth-service.default_fuzzer ####################

#################### android.hardware.bluetooth-service.default ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/aidl/default/service.cpp
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
    add_executable(android.hardware.bluetooth-service.default ${SOURCE_FILES})

    target_include_directories(android.hardware.bluetooth-service.default
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/bluetooth/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.bluetooth-service.default PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(android.hardware.bluetooth-service.default
            android.hardware.bluetooth-V1-ndk
            android.hardware.bluetooth.async
            android.hardware.bluetooth.hci
            libbase
            libbinder_ndk
            libbluetoothhcihalimpl
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.bluetooth-service.default ####################

#################### libbluetoothhcihalimpl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/aidl/default/BluetoothHci.cpp
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/bit_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/byte_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/byte_observer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/fragmenting_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/iterator.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/packet_view.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/raw_builder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/view.cc
        ${ANDROID_ROOT}/hardware/interfaces/bluetooth/aidl/default/net_bluetooth_mgmt.cpp
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
    add_library(libbluetoothhcihalimpl SHARED ${SOURCE_FILES})

    target_include_directories(libbluetoothhcihalimpl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/bluetooth/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(libbluetoothhcihalimpl PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(libbluetoothhcihalimpl
            android.hardware.bluetooth-V1-ndk
            android.hardware.bluetooth.async
            android.hardware.bluetooth.hci
            libbase
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### libbluetoothhcihalimpl ####################

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

project(libosi-AllocationTestHarness)

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

#################### libosi-AllocationTestHarness ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/test/AllocationTestHarness.cc
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
    add_library(libosi-AllocationTestHarness SHARED ${SOURCE_FILES})

    target_include_directories(libosi-AllocationTestHarness
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi
    )

    # Add c/c++ flags
    target_compile_options(libosi-AllocationTestHarness PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libosi-AllocationTestHarness
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-platform-protos-lite
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest
            libgtest_main
            libhidlbase
            liblog
            libm
            libstatslog_bt
            libstatssocket
            libutils
    )
endif ()
#################### libosi-AllocationTestHarness ####################

#################### libosi-AlarmTestHarness ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/test/AlarmTestHarness.cc
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
    add_library(libosi-AlarmTestHarness SHARED ${SOURCE_FILES})

    target_include_directories(libosi-AlarmTestHarness
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi
    )

    # Add c/c++ flags
    target_compile_options(libosi-AlarmTestHarness PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libosi-AlarmTestHarness
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-platform-protos-lite
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgtest
            libgtest_main
            libhidlbase
            liblog
            libm
            libstatslog_bt
            libstatssocket
            libutils
    )
endif ()
#################### libosi-AlarmTestHarness ####################

#################### libosi ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/compat.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/alarm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/allocation_tracker.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/allocator.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/array.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/buffer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/config.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/fixed_queue.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/future.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/hash_map_utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/list.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/mutex.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/osi.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/properties.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/reactor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/ringbuffer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/socket.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/socket_utils/socket_local_client.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/socket_utils/socket_local_server.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/thread.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/thread_scheduler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/wakelock.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/src/internal/semaphore.cc
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
    add_library(libosi SHARED ${SOURCE_FILES})

    target_include_directories(libosi
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi/include_internal
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/osi
    )

    # Add c/c++ flags
    target_compile_options(libosi PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libosi
            android.hardware.bluetooth.a2dp-1.0
            android.hardware.bluetooth.audio-2.0
            android.hardware.bluetooth.audio-2.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            libFraunhoferAAC
            libPlatformProperties
            libaaudio
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-types
            libbluetooth_gd
            libbluetooth_headers
            libbluetooth_rust_interop
            libbt-platform-protos-lite
            libbt-protos-lite
            libbt-sbc-decoder
            libbt-sbc-encoder
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libflatbuffers-cpp
            libfmq
            libg722codec
            libhidlbase
            liblc3
            liblog
            libm
            libnativehelper
            libopus
            libprotobuf-cpp-lite
            libstatslog_bt
            libstatssocket
            libudrv-uipc
            libutils
    )
endif ()
#################### libosi ####################

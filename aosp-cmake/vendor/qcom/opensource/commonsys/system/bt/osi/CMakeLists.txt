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

project(libosi_qti)

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

#################### libosi_qti ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/alarm.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/allocation_tracker.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/allocator.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/array.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/buffer.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/compat.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/config.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/config_legacy.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/fixed_queue.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/future.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/hash_map_utils.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/list.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/metrics.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/mutex.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/osi.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/properties.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/reactor.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/ringbuffer.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/semaphore.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/socket.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/socket_utils/socket_local_client.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/socket_utils/socket_local_server.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/thread.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/time.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/wakelock.cc
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
    add_library(libosi_qti SHARED ${SOURCE_FILES})

    target_include_directories(libosi_qti
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/protos
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/libbt-protos_qti/android_${OUT_ARCH_CPU}_static/gen/proto/vendor/qcom/opensource/commonsys/system/bt/osi
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/libbt-protos_qti/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libosi_qti PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DALOCATION_TRACKER_DEBUG
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libosi_qti
            libbluetooth-types
            libbluetooth_headers
            libbt-protos_qti
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
    )
endif ()
#################### libosi_qti ####################

#################### libosi-AllocationTestHarness_qti ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/test/AllocationTestHarness.cc
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
    add_library(libosi-AllocationTestHarness_qti SHARED ${SOURCE_FILES})

    target_include_directories(libosi-AllocationTestHarness_qti
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/protos
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi
    )

    # Add c/c++ flags
    target_compile_options(libosi-AllocationTestHarness_qti PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libosi-AllocationTestHarness_qti
            libbluetooth-types
            libbluetooth_headers
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest
            libgtest_main
            libm
    )
endif ()
#################### libosi-AllocationTestHarness_qti ####################

#################### libbt-protos_qti ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/libbt-protos_qti/android_${OUT_ARCH_CPU}_static/gen/proto/vendor/qcom/opensource/commonsys/system/bt/osi/src/protos/bluetooth.pb.cc
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
    add_library(libbt-protos_qti SHARED ${SOURCE_FILES})

    target_include_directories(libbt-protos_qti
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/libbt-protos_qti/android_${OUT_ARCH_CPU}_static/gen/proto/vendor/qcom/opensource/commonsys/system/bt/osi
            ${OUT_INTERMEDIATES_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/libbt-protos_qti/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libbt-protos_qti PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-protos_qti
            libbluetooth-types
            libbluetooth_headers
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### libbt-protos_qti ####################

#################### libosi-AlarmTestHarness_qti ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/test/AlarmTestHarness.cc
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
    add_library(libosi-AlarmTestHarness_qti SHARED ${SOURCE_FILES})

    target_include_directories(libosi-AlarmTestHarness_qti
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi/src/protos
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/osi
    )

    # Add c/c++ flags
    target_compile_options(libosi-AlarmTestHarness_qti PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libosi-AlarmTestHarness_qti
            libbluetooth-types
            libbluetooth_headers
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest
            libgtest_main
            libm
    )
endif ()
#################### libosi-AlarmTestHarness_qti ####################

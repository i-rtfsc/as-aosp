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

project(abb)

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

#################### adbd ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/adb/daemon/main.cpp
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
    add_executable(adbd ${SOURCE_FILES})

    target_include_directories(adbd
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb
            ${ANDROID_ROOT}/packages/modules/adb/crypto/include
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/include
            ${ANDROID_ROOT}/packages/modules/adb/tls/include
            ${ANDROID_ROOT}/packages/modules/adb/libs/adbconnection/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${ANDROID_ROOT}/packages/modules/adb/libs/libadbd_fs/include
    )

    # Add c/c++ flags
    target_compile_options(adbd PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(adbd
            libadb_crypto
            libadb_pairing_connection
            libadb_protos
            libadb_sysdeps
            libadb_tls_connection
            libadbconnection_server
            libadbd
            libadbd_auth
            libadbd_core
            libadbd_fs
            libadbd_services
            libapp_processes_protos_lite
            libasyncio
            libbase
            libbrotli
            libbuildversion
            libc++
            libcap
            libclang_rt.builtins
            libcrypto_static
            libcrypto_utils
            libcutils
            libcutils_sockets
            libdiagnose_usb
            liblog
            liblz4
            libmdnssd
            libminijail
            libprotobuf-cpp-lite
            libselinux
            libssl
            libzstd
    )
endif ()
#################### adbd ####################

#################### libadbd_core ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/adb/adb.cpp
        ${ANDROID_ROOT}/packages/modules/adb/adb_io.cpp
        ${ANDROID_ROOT}/packages/modules/adb/adb_listeners.cpp
        ${ANDROID_ROOT}/packages/modules/adb/adb_mdns.cpp
        ${ANDROID_ROOT}/packages/modules/adb/adb_trace.cpp
        ${ANDROID_ROOT}/packages/modules/adb/adb_unique_fd.cpp
        ${ANDROID_ROOT}/packages/modules/adb/adb_utils.cpp
        ${ANDROID_ROOT}/packages/modules/adb/fdevent/fdevent.cpp
        ${ANDROID_ROOT}/packages/modules/adb/services.cpp
        ${ANDROID_ROOT}/packages/modules/adb/sockets.cpp
        ${ANDROID_ROOT}/packages/modules/adb/socket_spec.cpp
        ${ANDROID_ROOT}/packages/modules/adb/sysdeps/env.cpp
        ${ANDROID_ROOT}/packages/modules/adb/sysdeps/errno.cpp
        ${ANDROID_ROOT}/packages/modules/adb/transport.cpp
        ${ANDROID_ROOT}/packages/modules/adb/transport_fd.cpp
        ${ANDROID_ROOT}/packages/modules/adb/types.cpp
        ${ANDROID_ROOT}/packages/modules/adb/fdevent/fdevent_epoll.cpp
        ${ANDROID_ROOT}/packages/modules/adb/sysdeps_unix.cpp
        ${ANDROID_ROOT}/packages/modules/adb/sysdeps/posix/network.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/adb_wifi.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/auth.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/jdwp_service.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/logging.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/transport_local.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/property_monitor.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/usb.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/usb_ffs.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/watchdog.cpp
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
    add_library(libadbd_core SHARED ${SOURCE_FILES})

    target_include_directories(libadbd_core
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb
            ${ANDROID_ROOT}/packages/modules/adb/crypto/include
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_shared/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_shared/gen/proto
            ${ANDROID_ROOT}/packages/modules/adb/tls/include
            ${ANDROID_ROOT}/packages/modules/adb/libs/adbconnection/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_shared/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_shared/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libadbd_core PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libadbd_core
            libadb_crypto
            libadb_pairing_connection
            libadb_protos
            libadb_tls_connection
            libadbconnection_server
            libadbd_auth
            libapp_processes_protos_lite
            libasyncio
            libbase
            libbuildversion
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils_sockets
            libdiagnose_usb
            libdl
            liblog
            libm
    )
endif ()
#################### libadbd_core ####################

#################### abb ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/adb/daemon/abb.cpp
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
    add_executable(abb ${SOURCE_FILES})

    target_include_directories(abb
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb
    )

    # Add c/c++ flags
    target_compile_options(abb PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(abb
            libadbd_core
            libadbd_services
            libbase
            libbinder
            libbuildversion
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcmd
            libdl
            liblog
            libm
            libselinux
            libutils
    )
endif ()
#################### abb ####################

#################### libadb_sysdeps ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/adb/sysdeps/env.cpp
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
    add_library(libadb_sysdeps SHARED ${SOURCE_FILES})

    target_include_directories(libadb_sysdeps
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb
    )

    # Add c/c++ flags
    target_compile_options(libadb_sysdeps PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libadb_sysdeps
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libwinpthread
    )
endif ()
#################### libadb_sysdeps ####################

#################### libadbd_services ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/adb/daemon/file_sync_service.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/services.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/shell_service.cpp
        ${ANDROID_ROOT}/packages/modules/adb/shell_service_protocol.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/abb_service.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/framebuffer_service.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/mdns.cpp
        ${ANDROID_ROOT}/packages/modules/adb/daemon/restart_service.cpp
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
    add_library(libadbd_services SHARED ${SOURCE_FILES})

    target_include_directories(libadbd_services
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb
            ${ANDROID_ROOT}/packages/modules/adb/libs/adbconnection/include
            ${ANDROID_ROOT}/packages/modules/adb/crypto/include
            ${ANDROID_ROOT}/packages/modules/adb/pairing_connection/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_shared/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_shared/gen/proto
            ${ANDROID_ROOT}/packages/modules/adb/tls/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_shared/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_shared/gen/proto
            ${ANDROID_ROOT}/packages/modules/adb/libs/libadbd_fs/include
    )

    # Add c/c++ flags
    target_compile_options(libadbd_services PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libadbd_services
            libadb_crypto
            libadb_pairing_connection
            libadb_protos
            libadb_tls_connection
            libadbconnection_server
            libadbd_auth
            libadbd_core
            libadbd_fs
            libapp_processes_protos_lite
            libasyncio
            libbase
            libbrotli
            libbuildversion
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils_sockets
            libdiagnose_usb
            libdl
            liblog
            liblz4
            libm
            libmdnssd
            libprotobuf-cpp-lite
            libselinux
            libzstd
    )
endif ()
#################### libadbd_services ####################

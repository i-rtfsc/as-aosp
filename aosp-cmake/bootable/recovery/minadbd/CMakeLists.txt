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

project(minadbd)

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

#################### libminadbd_services ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/minadbd/fuse_adb_provider.cpp
        ${ANDROID_ROOT}/bootable/recovery/minadbd/minadbd_services.cpp
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
    add_library(libminadbd_services SHARED ${SOURCE_FILES})

    target_include_directories(libminadbd_services
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery/minadbd
            ${ANDROID_ROOT}/bootable/recovery/minadbd/include
    )

    # Add c/c++ flags
    target_compile_options(libminadbd_services PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(libminadbd_services
            android.hardware.health-translate-ndk
            android.hardware.health-2.0
            libadbd
            libbase
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            libext4_utils
            libfs_mgr
            libfstab
            libfusesideload
            libhealthhalutils
            libhealthshim
            libhidlbase
            libm
            libminadbd_headers
            libotautil
            librecovery_utils
            libselinux
            libutils
    )
endif ()
#################### libminadbd_services ####################

#################### minadbd ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/minadbd/minadbd.cpp
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
    add_executable(minadbd ${SOURCE_FILES})

    target_include_directories(minadbd
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery/minadbd
            ${ANDROID_ROOT}/bootable/recovery/minadbd/include
    )

    # Add c/c++ flags
    target_compile_options(minadbd PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(minadbd
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            android.hardware.health-2.0
            libadb_crypto
            libadb_protos
            libadb_sysdeps
            libadb_tls_connection
            libadbconnection_server
            libadbd
            libadbd_auth
            libadbd_core
            libadbd_fs
            libapp_processes_protos_lite
            libasyncio
            libbase
            libbinder_ndk
            libbrotli
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils_sockets
            libdiagnose_usb
            libdl
            libext4_utils
            libfs_mgr
            libfstab
            libfusesideload
            libhealthhalutils
            libhealthshim
            libhidlbase
            liblog
            libm
            libmdnssd
            libminadbd_headers
            libminadbd_services
            libotautil
            libprotobuf-cpp-lite
            librecovery_utils
            libselinux
            libutils
            libzstd
            libc
            libdl
            libm
    )
endif ()
#################### minadbd ####################

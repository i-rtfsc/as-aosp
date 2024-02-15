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

project(libupdater_device)

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

#################### libupdater_device ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/updater/dynamic_partitions.cpp
        ${ANDROID_ROOT}/bootable/recovery/updater/updater_runtime.cpp
        ${ANDROID_ROOT}/bootable/recovery/updater/updater_runtime_dynamic_partitions.cpp
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
    add_library(libupdater_device SHARED ${SOURCE_FILES})

    target_include_directories(libupdater_device
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery/updater/include
            ${ANDROID_ROOT}/bootable/recovery/updater
    )

    # Add c/c++ flags
    target_compile_options(libupdater_device PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libupdater_device
            libapplypatch
            libavb
            libbase
            libbootloader_message
            libbrotli
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils
            libdl
            libdm
            libedify
            libext2_blkid
            libext2_com_err
            libext2_e2p
            libext2_quota
            libext2_uuid
            libext2fs
            libext4_utils
            libfec
            libfec_rs
            libfs_mgr
            libgtest_prod_headers
            liblog
            liblp
            libm
            libotautil
            libselinux
            libsparse
            libsquashfs_utils
            libtune2fs
            libupdater_core
            libutils
            libverity_tree
            libz
            libziparchive
    )
endif ()
#################### libupdater_device ####################

#################### libupdater_core ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/updater/blockimg.cpp
        ${ANDROID_ROOT}/bootable/recovery/updater/commands.cpp
        ${ANDROID_ROOT}/bootable/recovery/updater/install.cpp
        ${ANDROID_ROOT}/bootable/recovery/updater/mounts.cpp
        ${ANDROID_ROOT}/bootable/recovery/updater/updater.cpp
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
    add_library(libupdater_core SHARED ${SOURCE_FILES})

    target_include_directories(libupdater_core
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery/updater/include
            ${ANDROID_ROOT}/bootable/recovery/updater
    )

    # Add c/c++ flags
    target_compile_options(libupdater_core PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libupdater_core
            libapplypatch
            libavb
            libbase
            libbootloader_message
            libbrotli
            libbspatch
            libbz
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils
            libdl
            libdm
            libedify
            libext4_utils
            libfec
            libfec_rs
            libgtest_prod_headers
            liblog
            liblp
            libm
            libotautil
            libselinux
            libsparse
            libsquashfs_utils
            libutils
            libverity_tree
            libz
            libziparchive
    )
endif ()
#################### libupdater_core ####################

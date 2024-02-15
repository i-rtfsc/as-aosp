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

project(libspl_check)

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

#################### libspl_check ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/install/spl_check.cpp
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
    add_library(libspl_check SHARED ${SOURCE_FILES})

    target_include_directories(libspl_check
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery/install/include
            ${ANDROID_ROOT}/bootable/recovery/install
    )

    # Add c/c++ flags
    target_compile_options(libspl_check PRIVATE
            -Werror
    )

    target_link_libraries(libspl_check
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libprotobuf-cpp-lite
            libz
            libziparchive
            ota_metadata_proto_cc
    )
endif ()
#################### libspl_check ####################

#################### libinstall ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/install/snapshot_utils.cpp
        ${ANDROID_ROOT}/bootable/recovery/install/wipe_data.cpp
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
    add_library(libinstall SHARED ${SOURCE_FILES})

    target_include_directories(libinstall
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery/install/include
            ${ANDROID_ROOT}/bootable/recovery/install
    )

    # Add c/c++ flags
    target_compile_options(libinstall PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libinstall
            libbase
            libbootloader_message
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libext4_utils
            libfs_mgr
            libfusesideload
            libhidl-gen-utils
            libhidlbase
            liblog
            libm
            libminadbd_headers
            libminigui
            libminui
            libotautil
            libprotobuf-cpp-lite
            librecovery_ui
            librecovery_utils
            libselinux
            libsnapshot_nobinder
            libtinyxml2
            libutils
            libvintf
            libz
            libziparchive
            ota_metadata_proto_cc
    )
endif ()
#################### libinstall ####################

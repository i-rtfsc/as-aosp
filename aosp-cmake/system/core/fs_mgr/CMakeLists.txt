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

project(remount)

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

#################### libfs_mgr ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/blockdev.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/file_wait.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_format.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_dm_linear.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_overlayfs.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_roots.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_vendor_overlay.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/fiemap_writer.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/fiemap_status.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/image_manager.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/metadata.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/split_fiemap_writer.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/utility.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/passthrough.cpp
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
    add_library(libfs_mgr SHARED ${SOURCE_FILES})

    target_include_directories(libfs_mgr
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/include
            ${ANDROID_ROOT}/system/core/fs_mgr
            ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libstorage_literals
            ${ANDROID_ROOT}/system/core/fs_mgr/libdm/include
            ${ANDROID_ROOT}/system/core/fs_mgr/include_fstab
            ${ANDROID_ROOT}/system/core/fs_mgr/libfs_avb/include
            ${ANDROID_ROOT}/system/core/fs_mgr/liblp/include
    )

    # Add c/c++ flags
    target_compile_options(libfs_mgr PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libfs_mgr
            libavb
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils
            libdl
            libdm
            libext2_uuid
            libext4_utils
            libfec
            libfiemap_headers
            libfs_avb
            libfscrypt
            libfstab
            libgsi
            liblog
            liblogwrap
            liblp
            libm
            libselinux
            libstorage_literals_headers
    )
endif ()
#################### libfs_mgr ####################

#################### libfstab ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_fstab.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_boot_config.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_slotselect.cpp
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
    add_library(libfstab SHARED ${SOURCE_FILES})

    target_include_directories(libfstab
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/include_fstab
            ${ANDROID_ROOT}/system/core/fs_mgr/include
            ${ANDROID_ROOT}/system/core/fs_mgr
    )

    # Add c/c++ flags
    target_compile_options(libfstab PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libfstab
            libbase_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgsi_headers
            libm
    )
endif ()
#################### libfstab ####################

#################### libfs_mgr_binder ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/binder.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/blockdev.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/file_wait.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_format.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_dm_linear.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_overlayfs.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_roots.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_vendor_overlay.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/fiemap_writer.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/fiemap_status.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/image_manager.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/metadata.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/split_fiemap_writer.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/utility.cpp
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
    add_library(libfs_mgr_binder SHARED ${SOURCE_FILES})

    target_include_directories(libfs_mgr_binder
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/include
            ${ANDROID_ROOT}/system/core/fs_mgr
            ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libstorage_literals
            ${ANDROID_ROOT}/system/core/fs_mgr/libdm/include
            ${ANDROID_ROOT}/system/core/fs_mgr/include_fstab
            ${ANDROID_ROOT}/system/core/fs_mgr/libfs_avb/include
            ${ANDROID_ROOT}/system/core/fs_mgr/liblp/include
    )

    # Add c/c++ flags
    target_compile_options(libfs_mgr_binder PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libfs_mgr_binder
            gsi_aidl_interface-cpp
            libavb
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils
            libdl
            libdm
            libext2_uuid
            libext4_utils
            libfec
            libfiemap_headers
            libfs_avb
            libfscrypt
            libfstab
            libgsi
            libgsid
            liblog
            liblogwrap
            liblp
            libm
            libselinux
            libstorage_literals_headers
            libutils
    )
endif ()
#################### libfs_mgr_binder ####################

#################### remount ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/fs_mgr_remount.cpp
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
    add_executable(remount ${SOURCE_FILES})

    target_include_directories(remount
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/include
            ${ANDROID_ROOT}/system/core/fs_mgr
            ${ANDROID_ROOT}/system/core/fs_mgr/libfiemap/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libdm/include
            ${ANDROID_ROOT}/system/core/fs_mgr/include_fstab
            ${ANDROID_ROOT}/system/core/fs_mgr/libfs_avb/include
            ${ANDROID_ROOT}/system/core/fs_mgr/liblp/include
    )

    # Add c/c++ flags
    target_compile_options(remount PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(remount
            libavb_user
            libbase
            libbinder
            libbootloader_message
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libcutils_headers
            libdl
            libext4_utils
            libfs_mgr_binder
            libgsid
            liblog
            liblp
            libm
            libselinux
            libutils
            libvold_binder
    )
endif ()
#################### remount ####################

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

project(vdc)

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

#################### secdiscard ####################
set(FILES
        ${ANDROID_ROOT}/system/vold/FileDeviceUtils.cpp
        ${ANDROID_ROOT}/system/vold/secdiscard.cpp
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
    add_executable(secdiscard ${SOURCE_FILES})

    target_include_directories(secdiscard
            PUBLIC
            ${ANDROID_ROOT}/system/vold
    )

    # Add c/c++ flags
    target_compile_options(secdiscard PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(secdiscard
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### secdiscard ####################

#################### vold ####################
set(FILES
        ${ANDROID_ROOT}/system/vold/main.cpp
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
    add_executable(vold ${SOURCE_FILES})

    target_include_directories(vold
            PUBLIC
            ${ANDROID_ROOT}/system/vold
            ${OUT_INTERMEDIATES_ROOT}/system/vold/libvold_binder/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(vold PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(vold
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.health.storage-V1-ndk
            android.hardware.health.storage-1.0
            android.security.maintenance-ndk
            android.system.keystore2-V3-ndk
            libasync_safe
            libavb
            libbase
            libbinder
            libbinder_ndk
            libboot_control_client
            libbootloader_message
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto_utils
            libcutils
            libdiskconfig
            libdl
            libdm
            libext2_uuid
            libext4_utils
            libf2fs_sparseblock
            libfec
            libfec_rs
            libfs_avb
            libfs_mgr
            libgsi
            libhardware
            libhardware_legacy
            libhidlbase
            libincfs
            libkeymint_support
            libkeyutils
            liblog
            liblogwrap
            libm
            libselinux
            libsquashfs_utils
            libsysutils
            libutils
            libvold
            libvold_binder
    )
endif ()
#################### vold ####################

#################### vold_prepare_subdirs ####################
set(FILES
        ${ANDROID_ROOT}/system/vold/vold_prepare_subdirs.cpp
        ${ANDROID_ROOT}/system/vold/Utils.cpp
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
    add_executable(vold_prepare_subdirs ${SOURCE_FILES})

    target_include_directories(vold_prepare_subdirs
            PUBLIC
            ${ANDROID_ROOT}/system/vold
            ${OUT_INTERMEDIATES_ROOT}/system/vold/libvold_binder/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(vold_prepare_subdirs PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(vold_prepare_subdirs
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblogwrap
            libm
            libselinux
            libutils
            libvold_binder
    )
endif ()
#################### vold_prepare_subdirs ####################

#################### libvold ####################
set(FILES
        ${ANDROID_ROOT}/system/vold/AppFuseUtil.cpp
        ${ANDROID_ROOT}/system/vold/Benchmark.cpp
        ${ANDROID_ROOT}/system/vold/Checkpoint.cpp
        ${ANDROID_ROOT}/system/vold/CryptoType.cpp
        ${ANDROID_ROOT}/system/vold/EncryptInplace.cpp
        ${ANDROID_ROOT}/system/vold/FileDeviceUtils.cpp
        ${ANDROID_ROOT}/system/vold/FsCrypt.cpp
        ${ANDROID_ROOT}/system/vold/IdleMaint.cpp
        ${ANDROID_ROOT}/system/vold/KeyBuffer.cpp
        ${ANDROID_ROOT}/system/vold/KeyStorage.cpp
        ${ANDROID_ROOT}/system/vold/KeyUtil.cpp
        ${ANDROID_ROOT}/system/vold/Keystore.cpp
        ${ANDROID_ROOT}/system/vold/Loop.cpp
        ${ANDROID_ROOT}/system/vold/MetadataCrypt.cpp
        ${ANDROID_ROOT}/system/vold/MoveStorage.cpp
        ${ANDROID_ROOT}/system/vold/NetlinkHandler.cpp
        ${ANDROID_ROOT}/system/vold/NetlinkManager.cpp
        ${ANDROID_ROOT}/system/vold/Process.cpp
        ${ANDROID_ROOT}/system/vold/Utils.cpp
        ${ANDROID_ROOT}/system/vold/VoldNativeService.cpp
        ${ANDROID_ROOT}/system/vold/VoldNativeServiceValidation.cpp
        ${ANDROID_ROOT}/system/vold/VoldUtil.cpp
        ${ANDROID_ROOT}/system/vold/VolumeManager.cpp
        ${ANDROID_ROOT}/system/vold/cryptfs.cpp
        ${ANDROID_ROOT}/system/vold/fs/Exfat.cpp
        ${ANDROID_ROOT}/system/vold/fs/Ext4.cpp
        ${ANDROID_ROOT}/system/vold/fs/F2fs.cpp
        ${ANDROID_ROOT}/system/vold/fs/Vfat.cpp
        ${ANDROID_ROOT}/system/vold/model/Disk.cpp
        ${ANDROID_ROOT}/system/vold/model/EmulatedVolume.cpp
        ${ANDROID_ROOT}/system/vold/model/ObbVolume.cpp
        ${ANDROID_ROOT}/system/vold/model/PrivateVolume.cpp
        ${ANDROID_ROOT}/system/vold/model/PublicVolume.cpp
        ${ANDROID_ROOT}/system/vold/model/StubVolume.cpp
        ${ANDROID_ROOT}/system/vold/model/VolumeBase.cpp
        ${ANDROID_ROOT}/system/vold/model/VolumeEncryption.cpp
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
    add_library(libvold SHARED ${SOURCE_FILES})

    target_include_directories(libvold
            PUBLIC
            ${ANDROID_ROOT}/system/vold
            ${OUT_INTERMEDIATES_ROOT}/system/vold/libvold_binder/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libvold PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libvold
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.health.storage-V1-ndk
            android.hardware.health.storage-1.0
            android.security.maintenance-ndk
            android.system.keystore2-V3-ndk
            libasync_safe
            libavb
            libbase
            libbinder
            libbinder_ndk
            libboot_control_client
            libbootloader_message
            libc
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcom.android.sysprop.apex
            libcrypto
            libcrypto_utils
            libcutils
            libdiskconfig
            libdl
            libdm
            libext2_uuid
            libext4_utils
            libf2fs_sparseblock
            libfec
            libfec_rs
            libfs_avb
            libfs_mgr
            libgsi
            libhardware
            libhardware_legacy
            libhidlbase
            libincfs
            libkeymint_support
            libkeyutils
            liblog
            liblogwrap
            libm
            libselinux
            libsquashfs_utils
            libsysutils
            libutils
            libvold_binder
    )
endif ()
#################### libvold ####################

#################### vdc ####################
set(FILES
        ${ANDROID_ROOT}/system/vold/vdc.cpp
        ${ANDROID_ROOT}/system/vold/Utils.cpp
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
    add_executable(vdc ${SOURCE_FILES})

    target_include_directories(vdc
            PUBLIC
            ${ANDROID_ROOT}/system/vold
            ${OUT_INTERMEDIATES_ROOT}/system/vold/libvold_binder/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(vdc PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(vdc
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblogwrap
            libm
            libselinux
            libutils
            libvold_binder
    )
endif ()
#################### vdc ####################

#################### libvold_binder ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/vold/libvold_binder/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IVold.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/vold/libvold_binder/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IVoldListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/vold/libvold_binder/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IVoldMountCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/vold/libvold_binder/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IVoldTaskListener.cpp
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
    add_library(libvold_binder SHARED ${SOURCE_FILES})

    target_include_directories(libvold_binder
            PUBLIC
            ${ANDROID_ROOT}/system/vold
            ${OUT_INTERMEDIATES_ROOT}/system/vold/libvold_binder/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libvold_binder PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libvold_binder
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libincremental_aidl-cpp
            libm
            libutils
    )
endif ()
#################### libvold_binder ####################

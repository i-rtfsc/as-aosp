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

project(recovery-refresh)

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

#################### recovery-persist ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/recovery-persist.cpp
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
    add_executable(recovery-persist ${SOURCE_FILES})

    target_include_directories(recovery-persist
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery
            ${ANDROID_ROOT}/bootable/recovery/recovery_utils/include
    )

    # Add c/c++ flags
    target_compile_options(recovery-persist PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(recovery-persist
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            librecovery_utils
    )
endif ()
#################### recovery-persist ####################

#################### recovery ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/recovery_main.cpp
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
    add_executable(recovery ${SOURCE_FILES})

    target_include_directories(recovery
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery
            ${ANDROID_ROOT}/bootable/recovery/install/include
            ${ANDROID_ROOT}/bootable/recovery/recovery_ui/include
            ${ANDROID_ROOT}/bootable/recovery/recovery_utils/include
            ${ANDROID_ROOT}/bootable/recovery/otautil/include
            ${ANDROID_ROOT}/bootable/recovery/minui/include
            ${ANDROID_ROOT}/bootable/recovery/bootloader_message/include
            ${ANDROID_ROOT}/bootable/recovery/fuse_sideload/include
    )

    # Add c/c++ flags
    target_compile_options(recovery PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(recovery
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            android.hardware.boot-1.2
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            android.hardware.health-2.0
            libbase
            libbinder_ndk
            libboot_control_client
            libbootloader_message
            libc++
            libc++demangle
            libc++fs
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libext4_utils
            libfs_mgr
            libfstab
            libfusesideload
            libhealthhalutils
            libhealthshim
            libhidl-gen-utils
            libhidlbase
            libinstall
            liblog
            liblp
            liblz4
            libm
            libminigui
            libminui
            libotautil
            libprotobuf-cpp-lite
            librecovery
            librecovery_fastboot
            librecovery_ui
            librecovery_ui_default
            librecovery_utils
            libselinux
            libsnapshot_cow
            libsnapshot_nobinder
            libtinyxml2
            libutils
            libvintf
            libz
            libziparchive
            libzstd
            ota_metadata_proto_cc
            update_metadata-protos
            libc
            libdl
            libm
    )
endif ()
#################### recovery ####################

#################### recovery-refresh ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/recovery-refresh.cpp
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
    add_executable(recovery-refresh ${SOURCE_FILES})

    target_include_directories(recovery-refresh
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery
            ${ANDROID_ROOT}/bootable/recovery/recovery_utils/include
    )

    # Add c/c++ flags
    target_compile_options(recovery-refresh PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(recovery-refresh
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            librecovery_utils
    )
endif ()
#################### recovery-refresh ####################

#################### librecovery_fastboot ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/fastboot/fastboot.cpp
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
    add_library(librecovery_fastboot SHARED ${SOURCE_FILES})

    target_include_directories(librecovery_fastboot
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery
            ${ANDROID_ROOT}/bootable/recovery/recovery_ui/include
            ${ANDROID_ROOT}/bootable/recovery/minui/include
            ${ANDROID_ROOT}/bootable/recovery/bootloader_message/include
    )

    # Add c/c++ flags
    target_compile_options(librecovery_fastboot PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(librecovery_fastboot
            libbase
            libbootloader_message
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblock_passwd
            liblog
            libm
            libminigui
            libminui
            librecovery_ui
            librecovery_ui_default
    )
endif ()
#################### librecovery_fastboot ####################

#################### librecovery ####################
set(FILES
        ${ANDROID_ROOT}/bootable/recovery/recovery.cpp
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
    add_library(librecovery SHARED ${SOURCE_FILES})

    target_include_directories(librecovery
            PUBLIC
            ${ANDROID_ROOT}/bootable/recovery
            ${ANDROID_ROOT}/bootable/recovery/install/include
            ${ANDROID_ROOT}/bootable/recovery/recovery_ui/include
            ${ANDROID_ROOT}/bootable/recovery/minui/include
            ${ANDROID_ROOT}/bootable/recovery/recovery_utils/include
            ${ANDROID_ROOT}/bootable/recovery/otautil/include
            ${ANDROID_ROOT}/bootable/recovery/bootloader_message/include
    )

    # Add c/c++ flags
    target_compile_options(librecovery PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(librecovery
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            android.hardware.boot-1.2
            libbase
            libboot_control_client
            libbootloader_message
            libc
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libfs_mgr
            libinstall
            liblock_passwd
            liblog
            liblp
            liblz4
            libm
            libminigui
            libminui
            libotautil
            libprotobuf-cpp-lite
            librecovery_fastboot
            librecovery_ui
            librecovery_utils
            libsnapshot_cow
            libsnapshot_nobinder
            libziparchive
            libzstd
            update_metadata-protos
    )
endif ()
#################### librecovery ####################

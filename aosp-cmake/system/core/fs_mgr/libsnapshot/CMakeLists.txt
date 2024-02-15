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

project(inspect_cow)

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

#################### libsnapshot_nobinder ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_nobinder/android_${OUT_ARCH_CPU}_static/gen/proto/android/snapshot/snapshot.pb.cc
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/device_info.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_reader.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_stats.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_stub.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_metadata_updater.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_writer.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/partition_cow_creator.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/return.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/utility.cpp
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
    add_library(libsnapshot_nobinder SHARED ${SOURCE_FILES})

    target_include_directories(libsnapshot_nobinder
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot
            ${OUT_INTERMEDIATES_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_nobinder/android_${OUT_ARCH_CPU}_static/gen/proto
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/include
    )

    # Add c/c++ flags
    target_compile_options(libsnapshot_nobinder PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(libsnapshot_nobinder
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            libbase
            libboot_control_client
            libbrotli
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdm
            libext2_uuid
            libext4_utils
            libfiemap_headers
            libfs_mgr
            libfstab
            liblog
            libm
            libsnapshot_snapuserd
            libstorage_literals_headers
            libupdate_engine_headers
            libz
            update_metadata-protos
    )
endif ()
#################### libsnapshot_nobinder ####################

#################### snapshotctl ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshotctl.cpp
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
    add_executable(snapshotctl ${SOURCE_FILES})

    target_include_directories(snapshotctl
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/include
            ${OUT_INTERMEDIATES_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(snapshotctl PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(snapshotctl
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            libbase
            libboot_control_client
            libbrotli
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libext2_uuid
            libext4_utils
            libfs_mgr_binder
            libfstab
            libhidlbase
            liblog
            liblp
            liblz4
            libm
            libprotobuf-cpp-lite
            libsnapshot
            libsnapshot_cow
            libstatslog
            libstorage_literals_headers
            libutils
            libz
            update_metadata-protos
    )
endif ()
#################### snapshotctl ####################

#################### libsnapshot_cow ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_cow/cow_decompress.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_cow/cow_reader.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_cow/cow_writer.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_cow/cow_format.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_cow/cow_compress.cpp
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
    add_library(libsnapshot_cow SHARED ${SOURCE_FILES})

    target_include_directories(libsnapshot_cow
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot
    )

    # Add c/c++ flags
    target_compile_options(libsnapshot_cow PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(libsnapshot_cow
            libbase
            libbrotli
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            liblz4
            libm
            libz
    )
endif ()
#################### libsnapshot_cow ####################

#################### libsnapshot ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot/android_${OUT_ARCH_CPU}_static/gen/proto/android/snapshot/snapshot.pb.cc
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/device_info.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_reader.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_stats.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_stub.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_metadata_updater.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_writer.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/partition_cow_creator.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/return.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/utility.cpp
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
    add_library(libsnapshot SHARED ${SOURCE_FILES})

    target_include_directories(libsnapshot
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot
            ${OUT_INTERMEDIATES_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot/android_${OUT_ARCH_CPU}_static/gen/proto
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/include
    )

    # Add c/c++ flags
    target_compile_options(libsnapshot PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(libsnapshot
            android.hardware.boot-V1-ndk
            android.hardware.boot-1.0
            android.hardware.boot-1.1
            libbase
            libboot_control_client
            libbrotli
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdm
            libext2_uuid
            libext4_utils
            libfiemap_headers
            libfs_mgr_binder
            libfstab
            liblog
            libm
            libsnapshot_snapuserd
            libstorage_literals_headers
            libupdate_engine_headers
            libz
            update_metadata-protos
    )
endif ()
#################### libsnapshot ####################

#################### inspect_cow ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_cow/inspect_cow.cpp
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
    add_executable(inspect_cow ${SOURCE_FILES})

    target_include_directories(inspect_cow
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot
    )

    # Add c/c++ flags
    target_compile_options(inspect_cow PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(inspect_cow
            libbase
            libbrotli
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libdl
            liblog
            liblz4
            libm
            libsnapshot_cow
            libz
    )
endif ()
#################### inspect_cow ####################

#################### libsnapshot_init ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_init/android_${OUT_ARCH_CPU}_static/gen/proto/android/snapshot/snapshot.pb.cc
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/device_info.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_reader.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_stats.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_stub.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_metadata_updater.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapshot_writer.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/partition_cow_creator.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/return.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/utility.cpp
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
    add_library(libsnapshot_init SHARED ${SOURCE_FILES})

    target_include_directories(libsnapshot_init
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot
            ${OUT_INTERMEDIATES_ROOT}/system/core/fs_mgr/libsnapshot/libsnapshot_init/android_${OUT_ARCH_CPU}_static/gen/proto
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/include
    )

    # Add c/c++ flags
    target_compile_options(libsnapshot_init PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(libsnapshot_init
            libbase
            libbrotli
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libdm
            libext2_uuid
            libext4_utils
            libfiemap_headers
            libfs_mgr
            libfstab
            liblog
            libm
            libsnapshot_snapuserd
            libstorage_literals_headers
            libupdate_engine_headers
            libz
            update_metadata-protos
    )
endif ()
#################### libsnapshot_init ####################

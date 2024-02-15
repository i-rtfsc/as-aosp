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

project(snapuserd)

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

#################### snapuserd_ramdisk ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/dm-snapshot-merge/snapuserd_server.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/snapuserd_daemon.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/user-space-merge/snapuserd_server.cpp
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
    add_executable(snapuserd_ramdisk ${SOURCE_FILES})

    target_include_directories(snapuserd_ramdisk
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd
    )

    # Add c/c++ flags
    target_compile_options(snapuserd_ramdisk PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(snapuserd_ramdisk
            libbase
            libbrotli
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils_sockets
            libdm
            libext4_utils
            libfs_mgr
            libgflags
            liblog
            liblz4
            libm
            libsnapshot_cow
            libsnapshot_snapuserd
            libsnapuserd
            libstorage_literals_headers
            libunwind
            liburing
            libz
    )
endif ()
#################### snapuserd_ramdisk ####################

#################### libsnapshot_snapuserd ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/snapuserd_client.cpp
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
    add_library(libsnapshot_snapuserd SHARED ${SOURCE_FILES})

    target_include_directories(libsnapshot_snapuserd
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd
    )

    # Add c/c++ flags
    target_compile_options(libsnapshot_snapuserd PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libsnapshot_snapuserd
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils_sockets
            libdl
            libfs_mgr
            liblog
            libm
    )
endif ()
#################### libsnapshot_snapuserd ####################

#################### snapuserd ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/dm-snapshot-merge/snapuserd_server.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/snapuserd_daemon.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/user-space-merge/snapuserd_server.cpp
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
    add_executable(snapuserd ${SOURCE_FILES})

    target_include_directories(snapuserd
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd
    )

    # Add c/c++ flags
    target_compile_options(snapuserd PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(snapuserd
            libbase
            libbrotli
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils_sockets
            libdm
            libext4_utils
            libfs_mgr
            libgflags
            liblog
            liblz4
            libm
            libsnapshot_cow
            libsnapshot_snapuserd
            libsnapuserd
            libstorage_literals_headers
            libunwind
            liburing
            libz
    )
endif ()
#################### snapuserd ####################

#################### libsnapuserd ####################
set(FILES
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/dm-snapshot-merge/snapuserd.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/dm-snapshot-merge/snapuserd_worker.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/dm-snapshot-merge/snapuserd_readahead.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/snapuserd_buffer.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/user-space-merge/snapuserd_core.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/user-space-merge/snapuserd_dm_user.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/user-space-merge/snapuserd_merge.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/user-space-merge/snapuserd_readahead.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/user-space-merge/snapuserd_transitions.cpp
        ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/user-space-merge/snapuserd_verify.cpp
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
    add_library(libsnapuserd SHARED ${SOURCE_FILES})

    target_include_directories(libsnapuserd
            PUBLIC
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd/include
            ${ANDROID_ROOT}/system/core/fs_mgr/libsnapshot/snapuserd
    )

    # Add c/c++ flags
    target_compile_options(libsnapuserd PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libsnapuserd
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libdm
            libext4_utils
            libm
            libsnapshot_cow
            libstorage_literals_headers
            liburing
    )
endif ()
#################### libsnapuserd ####################

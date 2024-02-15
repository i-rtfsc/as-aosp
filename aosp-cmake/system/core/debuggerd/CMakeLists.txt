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

project(pbtombstone)

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

#################### libdebuggerd ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/backtrace.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/gwp_asan.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/open_files_list.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/tombstone.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/tombstone_proto.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/tombstone_proto_to_text.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/utility.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/scudo.cpp
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
    add_library(libdebuggerd SHARED ${SOURCE_FILES})

    target_include_directories(libdebuggerd
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/debuggerd/proto
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libdebuggerd PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(libdebuggerd
            bionic_libc_platform_headers
            gwp_asan_crash_handler
            gwp_asan_headers
            libasync_safe
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdexfile_support
            libdl
            liblog_headers
            liblzma
            libm
            libprocinfo
            libprotobuf-cpp-lite
            libscudo
            libtombstone_proto
            libunwindstack
            scudo_headers
    )
endif ()
#################### libdebuggerd ####################

#################### debuggerd ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/debuggerd.cpp
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
    add_executable(debuggerd ${SOURCE_FILES})

    target_include_directories(debuggerd
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/common/include
    )

    # Add c/c++ flags
    target_compile_options(debuggerd PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(debuggerd
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdebuggerd_client
            libdl
            liblog
            libm
            libprocinfo
    )
endif ()
#################### debuggerd ####################

#################### libdebuggerd_handler ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/handler/debuggerd_fallback_nop.cpp
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
    add_library(libdebuggerd_handler SHARED ${SOURCE_FILES})

    target_include_directories(libdebuggerd_handler
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/common/include
            ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/include
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/debuggerd/proto
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libdebuggerd_handler PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(libdebuggerd_handler
            bionic_libc_platform_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdebuggerd_handler_core
            libdl
            libm
    )
endif ()
#################### libdebuggerd_handler ####################

#################### libdebuggerd_client ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/client/debuggerd_client.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/util.cpp
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
    add_library(libdebuggerd_client SHARED ${SOURCE_FILES})

    target_include_directories(libdebuggerd_client
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/common/include
    )

    # Add c/c++ flags
    target_compile_options(libdebuggerd_client PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(libdebuggerd_client
            bionic_libc_platform_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdebuggerd_common_headers
            libdl
            libm
            libprocinfo
    )
endif ()
#################### libdebuggerd_client ####################

#################### tombstoned ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/util.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/tombstoned/intercept_manager.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/tombstoned/tombstoned.cpp
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
    add_executable(tombstoned ${SOURCE_FILES})

    target_include_directories(tombstoned
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/common/include
    )

    # Add c/c++ flags
    target_compile_options(tombstoned PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(tombstoned
            bionic_libc_platform_headers
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdebuggerd_common_headers
            libdl
            libevent
            liblog
            libm
    )
endif ()
#################### tombstoned ####################

#################### tombstoned.microdroid ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/util.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/tombstoned/intercept_manager.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/tombstoned/tombstoned.cpp
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
    add_executable(tombstoned.microdroid ${SOURCE_FILES})

    target_include_directories(tombstoned.microdroid
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/common/include
    )

    # Add c/c++ flags
    target_compile_options(tombstoned.microdroid PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(tombstoned.microdroid
            bionic_libc_platform_headers
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdebuggerd_common_headers
            libdl
            libevent
            liblog
            libm
    )
endif ()
#################### tombstoned.microdroid ####################

#################### libtombstoned_client ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/tombstoned/tombstoned_client.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/util.cpp
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
    add_library(libtombstoned_client SHARED ${SOURCE_FILES})

    target_include_directories(libtombstoned_client
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/tombstoned/include
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/common/include
    )

    # Add c/c++ flags
    target_compile_options(libtombstoned_client PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(libtombstoned_client
            libasync_safe
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdebuggerd_common_headers
            libdl
            libm
    )
endif ()
#################### libtombstoned_client ####################

#################### libdebuggerd_handler_fallback ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/handler/debuggerd_fallback.cpp
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
    add_library(libdebuggerd_handler_fallback SHARED ${SOURCE_FILES})

    target_include_directories(libdebuggerd_handler_fallback
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/common/include
            ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/include
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/debuggerd/proto
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto
            ${ANDROID_ROOT}/system/core/debuggerd/tombstoned/include
    )

    # Add c/c++ flags
    target_compile_options(libdebuggerd_handler_fallback PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(libdebuggerd_handler_fallback
            bionic_libc_platform_headers
            libasync_safe
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdebuggerd
            libdebuggerd_handler_core
            libdl
            liblzma
            libm
            libtombstoned_client_static
            libunwindstack_no_dex
    )
endif ()
#################### libdebuggerd_handler_fallback ####################

#################### libdebuggerd_handler_core ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/handler/debuggerd_handler.cpp
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
    add_library(libdebuggerd_handler_core SHARED ${SOURCE_FILES})

    target_include_directories(libdebuggerd_handler_core
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/common/include
            ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/include
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/debuggerd/proto
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libdebuggerd_handler_core PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(libdebuggerd_handler_core
            bionic_libc_platform_headers
            gwp_asan_headers
            libasync_safe
            libbase_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdebuggerd
            libdebuggerd_common_headers
            libdl
            libm
    )
endif ()
#################### libdebuggerd_handler_core ####################

#################### pbtombstone ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/pbtombstone.cpp
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
    add_executable(pbtombstone ${SOURCE_FILES})

    target_include_directories(pbtombstone
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/libdebuggerd/include
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/debuggerd/proto
            ${OUT_INTERMEDIATES_ROOT}/system/core/debuggerd/proto/libtombstone_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(pbtombstone PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(pbtombstone
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdebuggerd
            libdl
            liblog
            libm
            libprotobuf-cpp-lite
            libtombstone_proto
            libunwindstack
    )
endif ()
#################### pbtombstone ####################

#################### libtombstoned_client_static ####################
set(FILES
        ${ANDROID_ROOT}/system/core/debuggerd/tombstoned/tombstoned_client.cpp
        ${ANDROID_ROOT}/system/core/debuggerd/util.cpp
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
    add_library(libtombstoned_client_static SHARED ${SOURCE_FILES})

    target_include_directories(libtombstoned_client_static
            PUBLIC
            ${ANDROID_ROOT}/system/core/debuggerd/tombstoned/include
            ${ANDROID_ROOT}/system/core/debuggerd/include
            ${ANDROID_ROOT}/system/core/debuggerd
            ${ANDROID_ROOT}/system/core/debuggerd/common/include
    )

    # Add c/c++ flags
    target_compile_options(libtombstoned_client_static PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-argument
            -Wno-unused-function
            -DANDROID_DEBUGGABLE=0
            -UANDROID_DEBUGGABLE
            -DANDROID_DEBUGGABLE=1
    )

    target_link_libraries(libtombstoned_client_static
            libasync_safe
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdebuggerd_common_headers
            libdl
            libm
    )
endif ()
#################### libtombstoned_client_static ####################

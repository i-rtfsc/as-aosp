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

project(libartbased)

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

#################### libartbased ####################
set(FILES
        ${ANDROID_ROOT}/art/libartbase/arch/instruction_set.cc
        ${ANDROID_ROOT}/art/libartbase/base/allocator.cc
        ${ANDROID_ROOT}/art/libartbase/base/arena_allocator.cc
        ${ANDROID_ROOT}/art/libartbase/base/arena_bit_vector.cc
        ${ANDROID_ROOT}/art/libartbase/base/bit_vector.cc
        ${ANDROID_ROOT}/art/libartbase/base/compiler_filter.cc
        ${ANDROID_ROOT}/art/libartbase/base/enums.cc
        ${ANDROID_ROOT}/art/libartbase/base/file_magic.cc
        ${ANDROID_ROOT}/art/libartbase/base/file_utils.cc
        ${ANDROID_ROOT}/art/libartbase/base/flags.cc
        ${ANDROID_ROOT}/art/libartbase/base/hex_dump.cc
        ${ANDROID_ROOT}/art/libartbase/base/logging.cc
        ${ANDROID_ROOT}/art/libartbase/base/malloc_arena_pool.cc
        ${ANDROID_ROOT}/art/libartbase/base/membarrier.cc
        ${ANDROID_ROOT}/art/libartbase/base/memfd.cc
        ${ANDROID_ROOT}/art/libartbase/base/memory_region.cc
        ${ANDROID_ROOT}/art/libartbase/base/mem_map.cc
        ${ANDROID_ROOT}/art/libartbase/base/metrics/metrics_common.cc
        ${ANDROID_ROOT}/art/libartbase/base/os_linux.cc
        ${ANDROID_ROOT}/art/libartbase/base/runtime_debug.cc
        ${ANDROID_ROOT}/art/libartbase/base/safe_copy.cc
        ${ANDROID_ROOT}/art/libartbase/base/scoped_arena_allocator.cc
        ${ANDROID_ROOT}/art/libartbase/base/scoped_flock.cc
        ${ANDROID_ROOT}/art/libartbase/base/socket_peer_is_trusted.cc
        ${ANDROID_ROOT}/art/libartbase/base/time_utils.cc
        ${ANDROID_ROOT}/art/libartbase/base/unix_file/fd_file.cc
        ${ANDROID_ROOT}/art/libartbase/base/unix_file/random_access_file_utils.cc
        ${ANDROID_ROOT}/art/libartbase/base/utils.cc
        ${ANDROID_ROOT}/art/libartbase/base/zip_archive.cc
        ${ANDROID_ROOT}/art/libartbase/base/globals_unix.cc
        ${ANDROID_ROOT}/art/libartbase/base/mem_map_unix.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libartbase/art_libartbase_operator_srcs/gen/gensrcs/art/libartbase/arch/instruction_set.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libartbase/art_libartbase_operator_srcs/gen/gensrcs/art/libartbase/base/allocator.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libartbase/art_libartbase_operator_srcs/gen/gensrcs/art/libartbase/base/unix_file/fd_file.operator_out.cc
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
    add_library(libartbased SHARED ${SOURCE_FILES})

    target_include_directories(libartbased
            PUBLIC
            ${ANDROID_ROOT}/art/libartbase
            ${OUT_INTERMEDIATES_ROOT}/art/libartbase/art_libartbase_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libartbased PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(libartbased
            libartpalette
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libmodules-utils-build
            libtinyxml2
            libwinpthread
            libz
            libziparchive
    )
endif ()
#################### libartbased ####################

#################### libartbase ####################
set(FILES
        ${ANDROID_ROOT}/art/libartbase/arch/instruction_set.cc
        ${ANDROID_ROOT}/art/libartbase/base/allocator.cc
        ${ANDROID_ROOT}/art/libartbase/base/arena_allocator.cc
        ${ANDROID_ROOT}/art/libartbase/base/arena_bit_vector.cc
        ${ANDROID_ROOT}/art/libartbase/base/bit_vector.cc
        ${ANDROID_ROOT}/art/libartbase/base/compiler_filter.cc
        ${ANDROID_ROOT}/art/libartbase/base/enums.cc
        ${ANDROID_ROOT}/art/libartbase/base/file_magic.cc
        ${ANDROID_ROOT}/art/libartbase/base/file_utils.cc
        ${ANDROID_ROOT}/art/libartbase/base/flags.cc
        ${ANDROID_ROOT}/art/libartbase/base/hex_dump.cc
        ${ANDROID_ROOT}/art/libartbase/base/logging.cc
        ${ANDROID_ROOT}/art/libartbase/base/malloc_arena_pool.cc
        ${ANDROID_ROOT}/art/libartbase/base/membarrier.cc
        ${ANDROID_ROOT}/art/libartbase/base/memfd.cc
        ${ANDROID_ROOT}/art/libartbase/base/memory_region.cc
        ${ANDROID_ROOT}/art/libartbase/base/mem_map.cc
        ${ANDROID_ROOT}/art/libartbase/base/metrics/metrics_common.cc
        ${ANDROID_ROOT}/art/libartbase/base/os_linux.cc
        ${ANDROID_ROOT}/art/libartbase/base/runtime_debug.cc
        ${ANDROID_ROOT}/art/libartbase/base/safe_copy.cc
        ${ANDROID_ROOT}/art/libartbase/base/scoped_arena_allocator.cc
        ${ANDROID_ROOT}/art/libartbase/base/scoped_flock.cc
        ${ANDROID_ROOT}/art/libartbase/base/socket_peer_is_trusted.cc
        ${ANDROID_ROOT}/art/libartbase/base/time_utils.cc
        ${ANDROID_ROOT}/art/libartbase/base/unix_file/fd_file.cc
        ${ANDROID_ROOT}/art/libartbase/base/unix_file/random_access_file_utils.cc
        ${ANDROID_ROOT}/art/libartbase/base/utils.cc
        ${ANDROID_ROOT}/art/libartbase/base/zip_archive.cc
        ${ANDROID_ROOT}/art/libartbase/base/globals_unix.cc
        ${ANDROID_ROOT}/art/libartbase/base/mem_map_unix.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libartbase/art_libartbase_operator_srcs/gen/gensrcs/art/libartbase/arch/instruction_set.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libartbase/art_libartbase_operator_srcs/gen/gensrcs/art/libartbase/base/allocator.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libartbase/art_libartbase_operator_srcs/gen/gensrcs/art/libartbase/base/unix_file/fd_file.operator_out.cc
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
    add_library(libartbase SHARED ${SOURCE_FILES})

    target_include_directories(libartbase
            PUBLIC
            ${ANDROID_ROOT}/art/libartbase
            ${OUT_INTERMEDIATES_ROOT}/art/libartbase/art_libartbase_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libartbase PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libartbase
            libartpalette
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libmodules-utils-build
            libtinyxml2
            libwinpthread
            libz
            libziparchive
    )
endif ()
#################### libartbase ####################

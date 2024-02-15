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

project(leveldb)

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

#################### leveldb ####################
set(FILES
        ${ANDROID_ROOT}/external/leveldb/db/builder.cc
        ${ANDROID_ROOT}/external/leveldb/db/c.cc
        ${ANDROID_ROOT}/external/leveldb/db/db_impl.cc
        ${ANDROID_ROOT}/external/leveldb/db/db_iter.cc
        ${ANDROID_ROOT}/external/leveldb/db/dbformat.cc
        ${ANDROID_ROOT}/external/leveldb/db/dumpfile.cc
        ${ANDROID_ROOT}/external/leveldb/db/filename.cc
        ${ANDROID_ROOT}/external/leveldb/db/log_reader.cc
        ${ANDROID_ROOT}/external/leveldb/db/log_writer.cc
        ${ANDROID_ROOT}/external/leveldb/db/memtable.cc
        ${ANDROID_ROOT}/external/leveldb/db/repair.cc
        ${ANDROID_ROOT}/external/leveldb/db/table_cache.cc
        ${ANDROID_ROOT}/external/leveldb/db/version_edit.cc
        ${ANDROID_ROOT}/external/leveldb/db/version_set.cc
        ${ANDROID_ROOT}/external/leveldb/db/write_batch.cc
        ${ANDROID_ROOT}/external/leveldb/table/block_builder.cc
        ${ANDROID_ROOT}/external/leveldb/table/block.cc
        ${ANDROID_ROOT}/external/leveldb/table/filter_block.cc
        ${ANDROID_ROOT}/external/leveldb/table/format.cc
        ${ANDROID_ROOT}/external/leveldb/table/iterator.cc
        ${ANDROID_ROOT}/external/leveldb/table/merger.cc
        ${ANDROID_ROOT}/external/leveldb/table/table_builder.cc
        ${ANDROID_ROOT}/external/leveldb/table/table.cc
        ${ANDROID_ROOT}/external/leveldb/table/two_level_iterator.cc
        ${ANDROID_ROOT}/external/leveldb/util/arena.cc
        ${ANDROID_ROOT}/external/leveldb/util/bloom.cc
        ${ANDROID_ROOT}/external/leveldb/util/cache.cc
        ${ANDROID_ROOT}/external/leveldb/util/coding.cc
        ${ANDROID_ROOT}/external/leveldb/util/comparator.cc
        ${ANDROID_ROOT}/external/leveldb/util/crc32c.cc
        ${ANDROID_ROOT}/external/leveldb/util/env.cc
        ${ANDROID_ROOT}/external/leveldb/util/env_posix.cc
        ${ANDROID_ROOT}/external/leveldb/util/filter_policy.cc
        ${ANDROID_ROOT}/external/leveldb/util/hash.cc
        ${ANDROID_ROOT}/external/leveldb/util/logging.cc
        ${ANDROID_ROOT}/external/leveldb/util/options.cc
        ${ANDROID_ROOT}/external/leveldb/util/status.cc
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
    add_library(leveldb SHARED ${SOURCE_FILES})

    target_include_directories(leveldb
            PUBLIC
            ${ANDROID_ROOT}/external/leveldb/include
            ${ANDROID_ROOT}/external/leveldb/db
            ${ANDROID_ROOT}/external/leveldb/table
            ${ANDROID_ROOT}/external/leveldb/port
            ${ANDROID_ROOT}/external/leveldb
    )

    # Add c/c++ flags
    target_compile_options(leveldb PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(leveldb
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### leveldb ####################

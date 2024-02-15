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

project(libext2fs)

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

#################### libext2fs ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/ext2_err.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/alloc.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/alloc_sb.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/alloc_stats.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/alloc_tables.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/atexit.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/badblocks.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/bb_inode.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/bitmaps.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/bitops.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/blkmap64_ba.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/blkmap64_rb.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/blknum.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/block.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/bmap.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/check_desc.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/crc16.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/crc32c.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/csum.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/closefs.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/dblist.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/dblist_dir.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/digest_encode.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/dirblock.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/dirhash.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/dir_iterate.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/dupfs.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/expanddir.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/ext_attr.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/extent.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/fallocate.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/fileio.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/finddev.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/flushb.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/freefs.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/gen_bitmap.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/gen_bitmap64.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/get_num_dirs.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/get_pathname.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/getsize.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/getsectsize.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/hashmap.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/i_block.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/icount.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/imager.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/ind_block.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/initialize.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/inline.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/inline_data.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/inode.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/io_manager.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/ismounted.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/link.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/llseek.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/lookup.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/mmp.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/mkdir.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/mkjournal.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/namei.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/native.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/newdir.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/nls_utf8.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/openfs.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/progress.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/punch.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/qcow2.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/rbtree.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/read_bb.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/read_bb_file.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/res_gdt.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/rw_bitmaps.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/sha256.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/sha512.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/swapfs.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/symlink.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/undo_io.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/unix_io.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/sparse_io.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/unlink.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/valid_blk.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/version.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs/test_io.c
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
    add_library(libext2fs SHARED ${SOURCE_FILES})

    target_include_directories(libext2fs
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/lib/ext2fs
    )

    # Add c/c++ flags
    target_compile_options(libext2fs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libext2fs
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2-headers
            libext2_com_err
            libext2_uuid
            libsparse
            libwinpthread
            libz
    )
endif ()
#################### libext2fs ####################

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

project(debugfs_static)

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

#################### debugfs ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/debug_cmds.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/debugfs.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/util.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/ncheck.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/icheck.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/ls.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/lsdel.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/dump.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/set_fields.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/logdump.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/htree.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/unused.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/e2freefrag.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/filefrag.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/extent_cmds.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/extent_inode.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/zap.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/quota.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/xattrs.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/journal.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/revoke.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/recovery.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/do_journal.c
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
    add_executable(debugfs ${SOURCE_FILES})

    target_include_directories(debugfs
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/debugfs
    )

    # Add c/c++ flags
    target_compile_options(debugfs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -DDEBUGFS
    )

    target_link_libraries(debugfs
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_blkid
            libext2_com_err
            libext2_e2p
            libext2_misc
            libext2_quota
            libext2_ss
            libext2_support
            libext2_uuid
            libext2fs
    )
endif ()
#################### debugfs ####################

#################### debugfs_static ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/debug_cmds.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/debugfs.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/util.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/ncheck.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/icheck.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/ls.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/lsdel.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/dump.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/set_fields.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/logdump.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/htree.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/unused.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/e2freefrag.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/filefrag.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/extent_cmds.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/extent_inode.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/zap.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/quota.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/xattrs.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/journal.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/revoke.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/recovery.c
        ${ANDROID_ROOT}/external/e2fsprogs/debugfs/do_journal.c
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
    add_executable(debugfs_static ${SOURCE_FILES})

    target_include_directories(debugfs_static
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/debugfs
    )

    # Add c/c++ flags
    target_compile_options(debugfs_static PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -DDEBUGFS
    )

    target_link_libraries(debugfs_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_blkid
            libext2_com_err
            libext2_e2p
            libext2_misc
            libext2_quota
            libext2_ss
            libext2_support
            libext2_uuid
            libext2fs
            libm
            libunwind
    )
endif ()
#################### debugfs_static ####################

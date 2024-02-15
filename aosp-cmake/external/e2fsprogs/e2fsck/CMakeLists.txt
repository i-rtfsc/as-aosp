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

project(e2fsck_static)

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

#################### e2fsck ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/e2fsck.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/super.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass1.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass1b.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass2.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass3.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass4.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass5.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/logfile.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/journal.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/recovery.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/revoke.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/badblocks.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/util.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/unix.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/dirinfo.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/dx_dirinfo.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/ehandler.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/problem.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/message.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/ea_refcount.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/quota.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/rehash.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/region.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/sigcatcher.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/readahead.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/extents.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/encrypted_files.c
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
    add_executable(e2fsck ${SOURCE_FILES})

    target_include_directories(e2fsck
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/e2fsck
    )

    # Add c/c++ flags
    target_compile_options(e2fsck PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(e2fsck
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_blkid
            libext2_com_err
            libext2_e2p
            libext2_quota
            libext2_uuid
            libext2fs
    )
endif ()
#################### e2fsck ####################

#################### e2fsck_ramdisk ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/e2fsck.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/super.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass1.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass1b.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass2.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass3.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass4.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass5.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/logfile.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/journal.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/recovery.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/revoke.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/badblocks.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/util.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/unix.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/dirinfo.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/dx_dirinfo.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/ehandler.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/problem.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/message.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/ea_refcount.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/quota.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/rehash.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/region.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/sigcatcher.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/readahead.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/extents.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/encrypted_files.c
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
    add_executable(e2fsck_ramdisk ${SOURCE_FILES})

    target_include_directories(e2fsck_ramdisk
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/e2fsck
    )

    # Add c/c++ flags
    target_compile_options(e2fsck_ramdisk PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(e2fsck_ramdisk
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libext2_blkid
            libext2_com_err
            libext2_e2p
            libext2_quota
            libext2_uuid
            libext2fs
            libm
            libunwind
            libc
            libdl
    )
endif ()
#################### e2fsck_ramdisk ####################

#################### e2fsck_static ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/e2fsck.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/super.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass1.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass1b.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass2.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass3.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass4.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/pass5.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/logfile.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/journal.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/recovery.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/revoke.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/badblocks.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/util.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/unix.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/dirinfo.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/dx_dirinfo.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/ehandler.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/problem.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/message.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/ea_refcount.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/quota.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/rehash.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/region.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/sigcatcher.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/readahead.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/extents.c
        ${ANDROID_ROOT}/external/e2fsprogs/e2fsck/encrypted_files.c
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
    add_executable(e2fsck_static ${SOURCE_FILES})

    target_include_directories(e2fsck_static
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/e2fsck
    )

    # Add c/c++ flags
    target_compile_options(e2fsck_static PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(e2fsck_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_blkid
            libext2_com_err
            libext2_e2p
            libext2_quota
            libext2_uuid
            libext2fs
            libm
            libunwind
    )
endif ()
#################### e2fsck_static ####################

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

project(f2fs.fibmap)

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

#################### make_f2fs ####################
set(FILES
        ${ANDROID_ROOT}/external/f2fs-tools/lib/libf2fs_io.c
        ${ANDROID_ROOT}/external/f2fs-tools/mkfs/f2fs_format_main.c
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
    add_executable(make_f2fs ${SOURCE_FILES})

    target_include_directories(make_f2fs
            PUBLIC
            ${ANDROID_ROOT}/external/f2fs-tools/include
            ${ANDROID_ROOT}/external/f2fs-tools/mkfs
            ${ANDROID_ROOT}/external/f2fs-tools/fsck
            ${ANDROID_ROOT}/external/f2fs-tools
    )

    # Add c/c++ flags
    target_compile_options(make_f2fs PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(make_f2fs
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_uuid
            libf2fs_fmt
            libf2fs_fmt_host
            libm
            libsparse
            libwinpthread
            libz
    )
endif ()
#################### make_f2fs ####################

#################### fsck.f2fs ####################
set(FILES
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/dir.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/dict.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/mkquota.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/quotaio.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/quotaio_tree.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/quotaio_v2.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/node.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/segment.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/xattr.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/main.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/mount.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/libf2fs.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/libf2fs_io.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/libf2fs_zoned.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/nls_utf8.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/dump.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/fsck.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/resize.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/defrag.c
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
    add_executable(fsck.f2fs ${SOURCE_FILES})

    target_include_directories(fsck.f2fs
            PUBLIC
            ${ANDROID_ROOT}/external/f2fs-tools/include
            ${ANDROID_ROOT}/external/f2fs-tools/mkfs
            ${ANDROID_ROOT}/external/f2fs-tools/fsck
            ${ANDROID_ROOT}/external/f2fs-tools
    )

    # Add c/c++ flags
    target_compile_options(fsck.f2fs PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(fsck.f2fs
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_uuid
            libm
            libsparse
    )
endif ()
#################### fsck.f2fs ####################

#################### check_f2fs ####################
set(FILES
        ${ANDROID_ROOT}/external/f2fs-tools/tools/check_f2fs.c
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
    add_executable(check_f2fs ${SOURCE_FILES})

    target_include_directories(check_f2fs
            PUBLIC
            ${ANDROID_ROOT}/external/f2fs-tools
    )

    # Add c/c++ flags
    target_compile_options(check_f2fs PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(check_f2fs
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### check_f2fs ####################

#################### sload_f2fs ####################
set(FILES
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/dir.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/dict.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/mkquota.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/quotaio.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/quotaio_tree.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/quotaio_v2.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/node.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/segment.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/xattr.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/main.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/mount.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/libf2fs.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/libf2fs_io.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/libf2fs_zoned.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/nls_utf8.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/dump.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/fsck.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/sload.c
        ${ANDROID_ROOT}/external/f2fs-tools/fsck/compress.c
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
    add_executable(sload_f2fs ${SOURCE_FILES})

    target_include_directories(sload_f2fs
            PUBLIC
            ${ANDROID_ROOT}/external/f2fs-tools/include
            ${ANDROID_ROOT}/external/f2fs-tools/mkfs
            ${ANDROID_ROOT}/external/f2fs-tools/fsck
            ${ANDROID_ROOT}/external/f2fs-tools
    )

    # Add c/c++ flags
    target_compile_options(sload_f2fs PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(sload_f2fs
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libext2_uuid
            liblog
            liblz4
            libm
            libselinux
            libsparse
            libz
    )
endif ()
#################### sload_f2fs ####################

#################### f2fscrypt ####################
set(FILES
        ${ANDROID_ROOT}/external/f2fs-tools/tools/f2fscrypt.c
        ${ANDROID_ROOT}/external/f2fs-tools/tools/sha512.c
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
    add_executable(f2fscrypt ${SOURCE_FILES})

    target_include_directories(f2fscrypt
            PUBLIC
            ${ANDROID_ROOT}/external/f2fs-tools/include
            ${ANDROID_ROOT}/external/f2fs-tools
    )

    # Add c/c++ flags
    target_compile_options(f2fscrypt PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(f2fscrypt
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libext2_uuid
            libm
            libc
            libdl
            libm
    )
endif ()
#################### f2fscrypt ####################

#################### libf2fs_fmt ####################
set(FILES
        ${ANDROID_ROOT}/external/f2fs-tools/lib/libf2fs.c
        ${ANDROID_ROOT}/external/f2fs-tools/mkfs/f2fs_format.c
        ${ANDROID_ROOT}/external/f2fs-tools/mkfs/f2fs_format_utils.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/libf2fs_zoned.c
        ${ANDROID_ROOT}/external/f2fs-tools/lib/nls_utf8.c
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
    add_library(libf2fs_fmt SHARED ${SOURCE_FILES})

    target_include_directories(libf2fs_fmt
            PUBLIC
            ${ANDROID_ROOT}/external/f2fs-tools/include
            ${ANDROID_ROOT}/external/f2fs-tools/mkfs
            ${ANDROID_ROOT}/external/f2fs-tools/fsck
            ${ANDROID_ROOT}/external/f2fs-tools
    )

    # Add c/c++ flags
    target_compile_options(libf2fs_fmt PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libf2fs_fmt
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_uuid
            libm
    )
endif ()
#################### libf2fs_fmt ####################

#################### f2fs.fibmap ####################
set(FILES
        ${ANDROID_ROOT}/external/f2fs-tools/tools/fibmap.c
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
    add_executable(f2fs.fibmap ${SOURCE_FILES})

    target_include_directories(f2fs.fibmap
            PUBLIC
            ${ANDROID_ROOT}/external/f2fs-tools/include
            ${ANDROID_ROOT}/external/f2fs-tools
    )

    # Add c/c++ flags
    target_compile_options(f2fs.fibmap PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(f2fs.fibmap
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### f2fs.fibmap ####################

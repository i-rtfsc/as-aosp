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

project(mkfs.erofs)

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

#################### fsck.erofs ####################
set(FILES
        ${ANDROID_ROOT}/external/erofs-utils/lib/blobchunk.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/block_list.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/cache.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compress.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compress_hints.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_liblzma.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_lz4.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_lz4hc.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/config.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/data.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/decompress.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/dir.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/exclude.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/hashmap.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/inode.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/io.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/namei.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/sha256.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/super.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/xattr.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/zmap.c
        ${ANDROID_ROOT}/external/erofs-utils/fsck/main.c
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
    add_executable(fsck.erofs ${SOURCE_FILES})

    target_include_directories(fsck.erofs
            PUBLIC
            ${ANDROID_ROOT}/external/erofs-utils/include
            ${ANDROID_ROOT}/external/erofs-utils
            ${OUT_INTERMEDIATES_ROOT}/external/erofs-utils/erofs-utils-version.h/gen
    )

    # Add c/c++ flags
    target_compile_options(fsck.erofs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(fsck.erofs
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liberofs
            libext2_uuid
            liblog
            liblz4
            libm
            libselinux
    )
endif ()
#################### fsck.erofs ####################

#################### dump.erofs ####################
set(FILES
        ${ANDROID_ROOT}/external/erofs-utils/lib/blobchunk.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/block_list.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/cache.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compress.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compress_hints.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_liblzma.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_lz4.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_lz4hc.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/config.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/data.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/decompress.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/dir.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/exclude.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/hashmap.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/inode.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/io.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/namei.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/sha256.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/super.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/xattr.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/zmap.c
        ${ANDROID_ROOT}/external/erofs-utils/dump/main.c
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
    add_executable(dump.erofs ${SOURCE_FILES})

    target_include_directories(dump.erofs
            PUBLIC
            ${ANDROID_ROOT}/external/erofs-utils/include
            ${ANDROID_ROOT}/external/erofs-utils
            ${OUT_INTERMEDIATES_ROOT}/external/erofs-utils/erofs-utils-version.h/gen
    )

    # Add c/c++ flags
    target_compile_options(dump.erofs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(dump.erofs
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liberofs
            libext2_uuid
            liblog
            liblz4
            libm
            libselinux
    )
endif ()
#################### dump.erofs ####################

#################### liberofs ####################
set(FILES
        ${ANDROID_ROOT}/external/erofs-utils/lib/blobchunk.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/block_list.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/cache.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compress.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compress_hints.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_liblzma.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_lz4.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_lz4hc.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/config.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/data.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/decompress.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/dir.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/exclude.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/hashmap.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/inode.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/io.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/namei.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/sha256.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/super.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/xattr.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/zmap.c
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
    add_library(liberofs SHARED ${SOURCE_FILES})

    target_include_directories(liberofs
            PUBLIC
            ${ANDROID_ROOT}/external/erofs-utils/include
            ${ANDROID_ROOT}/external/erofs-utils
            ${OUT_INTERMEDIATES_ROOT}/external/erofs-utils/erofs-utils-version.h/gen
    )

    # Add c/c++ flags
    target_compile_options(liberofs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(liberofs
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libext2_uuid
            liblog
            liblz4
            libm
            libselinux
    )
endif ()
#################### liberofs ####################

#################### mkfs.erofs ####################
set(FILES
        ${ANDROID_ROOT}/external/erofs-utils/lib/blobchunk.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/block_list.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/cache.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compress.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compress_hints.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_liblzma.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_lz4.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/compressor_lz4hc.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/config.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/data.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/decompress.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/dir.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/exclude.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/hashmap.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/inode.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/io.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/namei.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/sha256.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/super.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/xattr.c
        ${ANDROID_ROOT}/external/erofs-utils/lib/zmap.c
        ${ANDROID_ROOT}/external/erofs-utils/mkfs/main.c
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
    add_executable(mkfs.erofs ${SOURCE_FILES})

    target_include_directories(mkfs.erofs
            PUBLIC
            ${ANDROID_ROOT}/external/erofs-utils/include
            ${ANDROID_ROOT}/external/erofs-utils
            ${OUT_INTERMEDIATES_ROOT}/external/erofs-utils/erofs-utils-version.h/gen
    )

    # Add c/c++ flags
    target_compile_options(mkfs.erofs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(mkfs.erofs
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liberofs
            libext2_uuid
            liblog
            liblz4
            libm
            libselinux
    )
endif ()
#################### mkfs.erofs ####################

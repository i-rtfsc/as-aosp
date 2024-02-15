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

project(mke2fs)

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

#################### libtune2fs ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/tune2fs.c
        ${ANDROID_ROOT}/external/e2fsprogs/misc/util.c
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
    add_library(libtune2fs SHARED ${SOURCE_FILES})

    target_include_directories(libtune2fs
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(libtune2fs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libtune2fs
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
    )
endif ()
#################### libtune2fs ####################

#################### e4crypt ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/e4crypt.c
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
    add_executable(e4crypt ${SOURCE_FILES})

    target_include_directories(e4crypt
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(e4crypt PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(e4crypt
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_uuid
            libext2fs
    )
endif ()
#################### e4crypt ####################

#################### e2freefrag ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/e2freefrag.c
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
    add_executable(e2freefrag ${SOURCE_FILES})

    target_include_directories(e2freefrag
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(e2freefrag PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(e2freefrag
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2-headers
            libext2_com_err
            libext2fs
    )
endif ()
#################### e2freefrag ####################

#################### blkid ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/blkid.c
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
    add_executable(blkid ${SOURCE_FILES})

    target_include_directories(blkid
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(blkid PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(blkid
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_blkid
            libext2_com_err
            libext2_e2p
            libext2fs
    )
endif ()
#################### blkid ####################

#################### blkid_static ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/blkid.c
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
    add_executable(blkid_static ${SOURCE_FILES})

    target_include_directories(blkid_static
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(blkid_static PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(blkid_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_blkid
            libext2_com_err
            libext2_e2p
            libext2_uuid
            libext2fs
            libm
            libunwind
    )
endif ()
#################### blkid_static ####################

#################### badblocks ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/badblocks.c
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
    add_executable(badblocks ${SOURCE_FILES})

    target_include_directories(badblocks
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(badblocks PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(badblocks
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_blkid
            libext2_com_err
            libext2_e2p
            libext2_uuid
            libext2fs
    )
endif ()
#################### badblocks ####################

#################### lsattr-e2fsprogs ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/lsattr.c
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
    add_executable(lsattr-e2fsprogs ${SOURCE_FILES})

    target_include_directories(lsattr-e2fsprogs
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(lsattr-e2fsprogs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(lsattr-e2fsprogs
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_com_err
            libext2_e2p
    )
endif ()
#################### lsattr-e2fsprogs ####################

#################### filefrag ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/filefrag.c
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
    add_executable(filefrag ${SOURCE_FILES})

    target_include_directories(filefrag
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(filefrag PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(filefrag
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2fs
    )
endif ()
#################### filefrag ####################

#################### e2image ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/e2image.c
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
    add_executable(e2image ${SOURCE_FILES})

    target_include_directories(e2image
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(e2image PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(e2image
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_blkid
            libext2_com_err
            libext2_quota
            libext2fs
    )
endif ()
#################### e2image ####################

#################### libext2_misc ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/create_inode.c
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
    add_library(libext2_misc SHARED ${SOURCE_FILES})

    target_include_directories(libext2_misc
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(libext2_misc PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libext2_misc
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_com_err
            libext2_quota
            libext2fs
            libwinpthread
    )
endif ()
#################### libext2_misc ####################

#################### tune2fs_static ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/tune2fs.c
        ${ANDROID_ROOT}/external/e2fsprogs/misc/util.c
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
    add_executable(tune2fs_static ${SOURCE_FILES})

    target_include_directories(tune2fs_static
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(tune2fs_static PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(tune2fs_static
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
#################### tune2fs_static ####################

#################### lsattr_static ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/lsattr.c
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
    add_executable(lsattr_static ${SOURCE_FILES})

    target_include_directories(lsattr_static
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(lsattr_static PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(lsattr_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_com_err
            libext2_e2p
            libm
            libunwind
    )
endif ()
#################### lsattr_static ####################

#################### chattr-e2fsprogs ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/chattr.c
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
    add_executable(chattr-e2fsprogs ${SOURCE_FILES})

    target_include_directories(chattr-e2fsprogs
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(chattr-e2fsprogs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(chattr-e2fsprogs
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_com_err
            libext2_e2p
    )
endif ()
#################### chattr-e2fsprogs ####################

#################### tune2fs_ramdisk ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/tune2fs.c
        ${ANDROID_ROOT}/external/e2fsprogs/misc/util.c
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
    add_executable(tune2fs_ramdisk ${SOURCE_FILES})

    target_include_directories(tune2fs_ramdisk
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(tune2fs_ramdisk PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(tune2fs_ramdisk
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
#################### tune2fs_ramdisk ####################

#################### tune2fs ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/tune2fs.c
        ${ANDROID_ROOT}/external/e2fsprogs/misc/util.c
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
    add_executable(tune2fs ${SOURCE_FILES})

    target_include_directories(tune2fs
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(tune2fs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(tune2fs
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
#################### tune2fs ####################

#################### mke2fs ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/misc/mke2fs.c
        ${ANDROID_ROOT}/external/e2fsprogs/misc/util.c
        ${ANDROID_ROOT}/external/e2fsprogs/misc/mk_hugefiles.c
        ${ANDROID_ROOT}/external/e2fsprogs/misc/default_profile.c
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
    add_executable(mke2fs ${SOURCE_FILES})

    target_include_directories(mke2fs
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/misc
    )

    # Add c/c++ flags
    target_compile_options(mke2fs PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(mke2fs
            libbase
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
            libext2_uuid
            libext2fs
            libsparse
            libwinpthread
            libz
    )
endif ()
#################### mke2fs ####################

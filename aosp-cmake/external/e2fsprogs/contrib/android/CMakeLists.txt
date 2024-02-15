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

project(e2fsdroid)

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

#################### ext2simg ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/contrib/android/ext2simg.c
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
    add_executable(ext2simg ${SOURCE_FILES})

    target_include_directories(ext2simg
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/contrib/android
    )

    # Add c/c++ flags
    target_compile_options(ext2simg PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(ext2simg
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_com_err
            libext2fs
            libsparse
            libz
    )
endif ()
#################### ext2simg ####################

#################### e2fsdroid ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/contrib/android/e2fsdroid.c
        ${ANDROID_ROOT}/external/e2fsprogs/contrib/android/block_range.c
        ${ANDROID_ROOT}/external/e2fsprogs/contrib/android/fsmap.c
        ${ANDROID_ROOT}/external/e2fsprogs/contrib/android/block_list.c
        ${ANDROID_ROOT}/external/e2fsprogs/contrib/android/base_fs.c
        ${ANDROID_ROOT}/external/e2fsprogs/contrib/android/perms.c
        ${ANDROID_ROOT}/external/e2fsprogs/contrib/android/basefs_allocator.c
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
    add_executable(e2fsdroid ${SOURCE_FILES})

    target_include_directories(e2fsdroid
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/contrib/android
    )

    # Add c/c++ flags
    target_compile_options(e2fsdroid PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(e2fsdroid
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libext2_com_err
            libext2_misc
            libext2fs
            liblog
            libselinux
            libsparse
            libz
    )
endif ()
#################### e2fsdroid ####################

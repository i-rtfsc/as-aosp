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

project(libext2_support)

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

#################### libext2_support ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/cstring.c
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
    add_library(libext2_support SHARED ${SOURCE_FILES})

    target_include_directories(libext2_support
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/lib/support
    )

    # Add c/c++ flags
    target_compile_options(libext2_support PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libext2_support
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2-headers
    )
endif ()
#################### libext2_support ####################

#################### libext2_quota ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/devname.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/dict.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/mkquota.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/parse_qtype.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/plausible.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/profile.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/profile_helpers.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/prof_err.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/quotaio.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/quotaio_tree.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/quotaio_v2.c
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
    add_library(libext2_quota SHARED ${SOURCE_FILES})

    target_include_directories(libext2_quota
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/lib/support
    )

    # Add c/c++ flags
    target_compile_options(libext2_quota PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libext2_quota
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2-headers
            libext2_blkid
            libext2_com_err
            libext2fs
            libwinpthread
    )
endif ()
#################### libext2_quota ####################

#################### libext2_profile ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/prof_err.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/support/profile.c
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
    add_library(libext2_profile SHARED ${SOURCE_FILES})

    target_include_directories(libext2_profile
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/lib/support
    )

    # Add c/c++ flags
    target_compile_options(libext2_profile PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libext2_profile
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2-headers
            libext2_com_err
    )
endif ()
#################### libext2_profile ####################

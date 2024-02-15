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

project(fsverity)

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

#################### fsverity_multilib ####################
set(FILES
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_digest.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_dump_metadata.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_enable.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_measure.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_sign.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/fsverity.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/utils.c
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
    add_executable(fsverity_multilib ${SOURCE_FILES})

    target_include_directories(fsverity_multilib
            PUBLIC
            ${ANDROID_ROOT}/external/fsverity-utils
            ${ANDROID_ROOT}/external/fsverity-utils/include
    )

    # Add c/c++ flags
    target_compile_options(fsverity_multilib PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(fsverity_multilib
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libfsverity
            libm
    )
endif ()
#################### fsverity_multilib ####################

#################### libfsverity ####################
set(FILES
        ${ANDROID_ROOT}/external/fsverity-utils/lib/compute_digest.c
        ${ANDROID_ROOT}/external/fsverity-utils/lib/enable.c
        ${ANDROID_ROOT}/external/fsverity-utils/lib/hash_algs.c
        ${ANDROID_ROOT}/external/fsverity-utils/lib/sign_digest.c
        ${ANDROID_ROOT}/external/fsverity-utils/lib/utils.c
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
    add_library(libfsverity SHARED ${SOURCE_FILES})

    target_include_directories(libfsverity
            PUBLIC
            ${ANDROID_ROOT}/external/fsverity-utils/include
            ${ANDROID_ROOT}/external/fsverity-utils
    )

    # Add c/c++ flags
    target_compile_options(libfsverity PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libfsverity
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libm
    )
endif ()
#################### libfsverity ####################

#################### fsverity ####################
set(FILES
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_digest.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_dump_metadata.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_enable.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_measure.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/cmd_sign.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/fsverity.c
        ${ANDROID_ROOT}/external/fsverity-utils/programs/utils.c
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
    add_executable(fsverity ${SOURCE_FILES})

    target_include_directories(fsverity
            PUBLIC
            ${ANDROID_ROOT}/external/fsverity-utils
            ${ANDROID_ROOT}/external/fsverity-utils/include
    )

    # Add c/c++ flags
    target_compile_options(fsverity PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(fsverity
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libfsverity
            libm
    )
endif ()
#################### fsverity ####################

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

project(sg_write_buffer)

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

#################### sg_read_buffer ####################
set(FILES
        ${ANDROID_ROOT}/external/sg3_utils/src/sg_read_buffer.c
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
    add_executable(sg_read_buffer ${SOURCE_FILES})

    target_include_directories(sg_read_buffer
            PUBLIC
            ${ANDROID_ROOT}/external/sg3_utils/include
            ${ANDROID_ROOT}/external/sg3_utils
    )

    # Add c/c++ flags
    target_compile_options(sg_read_buffer PRIVATE
            -Wextra
            -Wall
            -Werror
    )

    target_link_libraries(sg_read_buffer
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libsgutils2
            libc
            libdl
            libm
    )
endif ()
#################### sg_read_buffer ####################

#################### libsgutils2 ####################
set(FILES
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_lib.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_pr2serr.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_lib_data.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_lib_names.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_cmds_basic.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_cmds_basic2.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_cmds_extra.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_cmds_mmc.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_pt_common.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_json_builder.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_pt_linux.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_io_linux.c
        ${ANDROID_ROOT}/external/sg3_utils/lib/sg_pt_linux_nvme.c
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
    add_library(libsgutils2 SHARED ${SOURCE_FILES})

    target_include_directories(libsgutils2
            PUBLIC
            ${ANDROID_ROOT}/external/sg3_utils/include
            ${ANDROID_ROOT}/external/sg3_utils
    )

    # Add c/c++ flags
    target_compile_options(libsgutils2 PRIVATE
            -Wextra
            -Wall
            -Werror
    )

    target_link_libraries(libsgutils2
            libc++
            libc++_static
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
#################### libsgutils2 ####################

#################### sg_write_buffer ####################
set(FILES
        ${ANDROID_ROOT}/external/sg3_utils/src/sg_write_buffer.c
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
    add_executable(sg_write_buffer ${SOURCE_FILES})

    target_include_directories(sg_write_buffer
            PUBLIC
            ${ANDROID_ROOT}/external/sg3_utils/include
            ${ANDROID_ROOT}/external/sg3_utils
    )

    # Add c/c++ flags
    target_compile_options(sg_write_buffer PRIVATE
            -Wextra
            -Wall
            -Werror
    )

    target_link_libraries(sg_write_buffer
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libsgutils2
            libc
            libdl
            libm
    )
endif ()
#################### sg_write_buffer ####################

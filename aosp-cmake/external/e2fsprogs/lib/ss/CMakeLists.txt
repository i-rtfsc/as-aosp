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

project(libext2_ss)

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

#################### libext2_ss ####################
set(FILES
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/ss_err.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/std_rqs.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/invocation.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/help.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/execute_cmd.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/listen.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/parse.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/error.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/prompt.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/request_tbl.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/list_rqs.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/pager.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/requests.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/data.c
        ${ANDROID_ROOT}/external/e2fsprogs/lib/ss/get_readline.c
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
    add_library(libext2_ss SHARED ${SOURCE_FILES})

    target_include_directories(libext2_ss
            PUBLIC
            ${ANDROID_ROOT}/external/e2fsprogs/lib/ss
    )

    # Add c/c++ flags
    target_compile_options(libext2_ss PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libext2_ss
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2-headers
            libext2_com_err
    )
endif ()
#################### libext2_ss ####################

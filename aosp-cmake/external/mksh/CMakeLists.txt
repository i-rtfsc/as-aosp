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

project(sh_vendor)

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

#################### sh ####################
set(FILES
        ${ANDROID_ROOT}/external/mksh/src/lalloc.c
        ${ANDROID_ROOT}/external/mksh/src/edit.c
        ${ANDROID_ROOT}/external/mksh/src/eval.c
        ${ANDROID_ROOT}/external/mksh/src/exec.c
        ${ANDROID_ROOT}/external/mksh/src/expr.c
        ${ANDROID_ROOT}/external/mksh/src/funcs.c
        ${ANDROID_ROOT}/external/mksh/src/histrap.c
        ${ANDROID_ROOT}/external/mksh/src/jobs.c
        ${ANDROID_ROOT}/external/mksh/src/lex.c
        ${ANDROID_ROOT}/external/mksh/src/main.c
        ${ANDROID_ROOT}/external/mksh/src/misc.c
        ${ANDROID_ROOT}/external/mksh/src/shf.c
        ${ANDROID_ROOT}/external/mksh/src/syn.c
        ${ANDROID_ROOT}/external/mksh/src/tree.c
        ${ANDROID_ROOT}/external/mksh/src/ulimit.c
        ${ANDROID_ROOT}/external/mksh/src/var.c
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
    add_executable(sh ${SOURCE_FILES})

    target_include_directories(sh
            PUBLIC
            ${ANDROID_ROOT}/external/mksh/src
            ${ANDROID_ROOT}/external/mksh
    )

    # Add c/c++ flags
    target_compile_options(sh PRIVATE
            -Wall
            -Werror
            -DDEBUG_LEAKS
    )

    target_link_libraries(sh
            libc
            libclang_rt.builtins
    )
endif ()
#################### sh ####################

#################### sh_vendor ####################
set(FILES
        ${ANDROID_ROOT}/external/mksh/src/lalloc.c
        ${ANDROID_ROOT}/external/mksh/src/edit.c
        ${ANDROID_ROOT}/external/mksh/src/eval.c
        ${ANDROID_ROOT}/external/mksh/src/exec.c
        ${ANDROID_ROOT}/external/mksh/src/expr.c
        ${ANDROID_ROOT}/external/mksh/src/funcs.c
        ${ANDROID_ROOT}/external/mksh/src/histrap.c
        ${ANDROID_ROOT}/external/mksh/src/jobs.c
        ${ANDROID_ROOT}/external/mksh/src/lex.c
        ${ANDROID_ROOT}/external/mksh/src/main.c
        ${ANDROID_ROOT}/external/mksh/src/misc.c
        ${ANDROID_ROOT}/external/mksh/src/shf.c
        ${ANDROID_ROOT}/external/mksh/src/syn.c
        ${ANDROID_ROOT}/external/mksh/src/tree.c
        ${ANDROID_ROOT}/external/mksh/src/ulimit.c
        ${ANDROID_ROOT}/external/mksh/src/var.c
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
    add_executable(sh_vendor ${SOURCE_FILES})

    target_include_directories(sh_vendor
            PUBLIC
            ${ANDROID_ROOT}/external/mksh/src
            ${ANDROID_ROOT}/external/mksh
    )

    # Add c/c++ flags
    target_compile_options(sh_vendor PRIVATE
            -Wall
            -Werror
            -DDEBUG_LEAKS
    )

    target_link_libraries(sh_vendor
            libc
            libclang_rt.builtins
            libc
    )
endif ()
#################### sh_vendor ####################

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

project(awk)

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

#################### awk_vendor ####################
set(FILES
        ${ANDROID_ROOT}/external/one-true-awk/b.c
        ${ANDROID_ROOT}/external/one-true-awk/lex.c
        ${ANDROID_ROOT}/external/one-true-awk/lib.c
        ${ANDROID_ROOT}/external/one-true-awk/main.c
        ${ANDROID_ROOT}/external/one-true-awk/parse.c
        ${OUT_INTERMEDIATES_ROOT}/external/one-true-awk/proctab.c/gen/proctab.c
        ${ANDROID_ROOT}/external/one-true-awk/run.c
        ${ANDROID_ROOT}/external/one-true-awk/tran.c
        ${OUT_INTERMEDIATES_ROOT}/external/one-true-awk/awkgram.tab.c/gen/awkgram.tab.c
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
    add_executable(awk_vendor ${SOURCE_FILES})

    target_include_directories(awk_vendor
            PUBLIC
            ${ANDROID_ROOT}/external/one-true-awk
            ${OUT_INTERMEDIATES_ROOT}/external/one-true-awk/awkgram.tab.h/gen
    )

    # Add c/c++ flags
    target_compile_options(awk_vendor PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(awk_vendor
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### awk_vendor ####################

#################### awk ####################
set(FILES
        ${ANDROID_ROOT}/external/one-true-awk/b.c
        ${ANDROID_ROOT}/external/one-true-awk/lex.c
        ${ANDROID_ROOT}/external/one-true-awk/lib.c
        ${ANDROID_ROOT}/external/one-true-awk/main.c
        ${ANDROID_ROOT}/external/one-true-awk/parse.c
        ${OUT_INTERMEDIATES_ROOT}/external/one-true-awk/proctab.c/gen/proctab.c
        ${ANDROID_ROOT}/external/one-true-awk/run.c
        ${ANDROID_ROOT}/external/one-true-awk/tran.c
        ${OUT_INTERMEDIATES_ROOT}/external/one-true-awk/awkgram.tab.c/gen/awkgram.tab.c
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
    add_executable(awk ${SOURCE_FILES})

    target_include_directories(awk
            PUBLIC
            ${ANDROID_ROOT}/external/one-true-awk
            ${OUT_INTERMEDIATES_ROOT}/external/one-true-awk/awkgram.tab.h/gen
    )

    # Add c/c++ flags
    target_compile_options(awk PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-result
    )

    target_link_libraries(awk
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### awk ####################

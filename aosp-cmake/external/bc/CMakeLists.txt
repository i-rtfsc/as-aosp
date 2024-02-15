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

project(bc)

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

#################### bc_vendor ####################
set(FILES
        ${ANDROID_ROOT}/external/bc/src/args.c
        ${ANDROID_ROOT}/external/bc/src/bc.c
        ${ANDROID_ROOT}/external/bc/src/bc_lex.c
        ${ANDROID_ROOT}/external/bc/src/bc_parse.c
        ${ANDROID_ROOT}/external/bc/src/data.c
        ${ANDROID_ROOT}/external/bc/src/dc.c
        ${ANDROID_ROOT}/external/bc/src/dc_lex.c
        ${ANDROID_ROOT}/external/bc/src/dc_parse.c
        ${ANDROID_ROOT}/external/bc/src/file.c
        ${ANDROID_ROOT}/external/bc/src/history.c
        ${ANDROID_ROOT}/external/bc/src/lang.c
        ${ANDROID_ROOT}/external/bc/src/lex.c
        ${ANDROID_ROOT}/external/bc/src/library.c
        ${ANDROID_ROOT}/external/bc/src/main.c
        ${ANDROID_ROOT}/external/bc/src/num.c
        ${ANDROID_ROOT}/external/bc/src/opt.c
        ${ANDROID_ROOT}/external/bc/src/parse.c
        ${ANDROID_ROOT}/external/bc/src/program.c
        ${ANDROID_ROOT}/external/bc/src/rand.c
        ${ANDROID_ROOT}/external/bc/src/read.c
        ${ANDROID_ROOT}/external/bc/src/vector.c
        ${ANDROID_ROOT}/external/bc/src/vm.c
        ${OUT_INTERMEDIATES_ROOT}/external/bc/bc-bc_help.c/gen/gen/bc_help.c
        ${OUT_INTERMEDIATES_ROOT}/external/bc/bc-lib.c/gen/gen/lib.c
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
    add_executable(bc_vendor ${SOURCE_FILES})

    target_include_directories(bc_vendor
            PUBLIC
            ${ANDROID_ROOT}/external/bc/include
            ${ANDROID_ROOT}/external/bc
            ${OUT_INTERMEDIATES_ROOT}/external/bc/bc-version.h/gen
    )

    # Add c/c++ flags
    target_compile_options(bc_vendor PRIVATE
            -Werror
    )

    target_link_libraries(bc_vendor
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### bc_vendor ####################

#################### bc ####################
set(FILES
        ${ANDROID_ROOT}/external/bc/src/args.c
        ${ANDROID_ROOT}/external/bc/src/bc.c
        ${ANDROID_ROOT}/external/bc/src/bc_lex.c
        ${ANDROID_ROOT}/external/bc/src/bc_parse.c
        ${ANDROID_ROOT}/external/bc/src/data.c
        ${ANDROID_ROOT}/external/bc/src/dc.c
        ${ANDROID_ROOT}/external/bc/src/dc_lex.c
        ${ANDROID_ROOT}/external/bc/src/dc_parse.c
        ${ANDROID_ROOT}/external/bc/src/file.c
        ${ANDROID_ROOT}/external/bc/src/history.c
        ${ANDROID_ROOT}/external/bc/src/lang.c
        ${ANDROID_ROOT}/external/bc/src/lex.c
        ${ANDROID_ROOT}/external/bc/src/library.c
        ${ANDROID_ROOT}/external/bc/src/main.c
        ${ANDROID_ROOT}/external/bc/src/num.c
        ${ANDROID_ROOT}/external/bc/src/opt.c
        ${ANDROID_ROOT}/external/bc/src/parse.c
        ${ANDROID_ROOT}/external/bc/src/program.c
        ${ANDROID_ROOT}/external/bc/src/rand.c
        ${ANDROID_ROOT}/external/bc/src/read.c
        ${ANDROID_ROOT}/external/bc/src/vector.c
        ${ANDROID_ROOT}/external/bc/src/vm.c
        ${OUT_INTERMEDIATES_ROOT}/external/bc/bc-bc_help.c/gen/gen/bc_help.c
        ${OUT_INTERMEDIATES_ROOT}/external/bc/bc-lib.c/gen/gen/lib.c
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
    add_executable(bc ${SOURCE_FILES})

    target_include_directories(bc
            PUBLIC
            ${ANDROID_ROOT}/external/bc/include
            ${ANDROID_ROOT}/external/bc
            ${OUT_INTERMEDIATES_ROOT}/external/bc/bc-version.h/gen
    )

    # Add c/c++ flags
    target_compile_options(bc PRIVATE
            -Werror
    )

    target_link_libraries(bc
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bc ####################

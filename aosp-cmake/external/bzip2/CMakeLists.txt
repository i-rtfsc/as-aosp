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

project(bzip2)

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

#################### libbz ####################
set(FILES
        ${ANDROID_ROOT}/external/bzip2/blocksort.c
        ${ANDROID_ROOT}/external/bzip2/bzlib.c
        ${ANDROID_ROOT}/external/bzip2/compress.c
        ${ANDROID_ROOT}/external/bzip2/crctable.c
        ${ANDROID_ROOT}/external/bzip2/decompress.c
        ${ANDROID_ROOT}/external/bzip2/huffman.c
        ${ANDROID_ROOT}/external/bzip2/randtable.c
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
    add_library(libbz SHARED ${SOURCE_FILES})

    target_include_directories(libbz
            PUBLIC
            ${ANDROID_ROOT}/external/bzip2
    )

    # Add c/c++ flags
    target_compile_options(libbz PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libbz
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libbz ####################

#################### bzip2 ####################
set(FILES
        ${ANDROID_ROOT}/external/bzip2/bzip2.c
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
    add_executable(bzip2 ${SOURCE_FILES})

    target_include_directories(bzip2
            PUBLIC
            ${ANDROID_ROOT}/external/bzip2
    )

    # Add c/c++ flags
    target_compile_options(bzip2 PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(bzip2
            libbz
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bzip2 ####################

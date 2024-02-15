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

project(libjemalloc5_jet)

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

#################### libjemalloc5_jet ####################
set(FILES
        ${ANDROID_ROOT}/external/jemalloc_new/src/jemalloc.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/arena.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/background_thread.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/base.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/bin.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/bitmap.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/ckh.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/ctl.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/div.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/extent.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/extent_dss.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/extent_mmap.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/hash.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/hooks.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/large.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/log.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/malloc_io.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/mutex.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/mutex_pool.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/nstime.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/pages.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/prng.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/prof.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/rtree.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/stats.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/sz.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/tcache.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/ticker.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/tsd.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/witness.c
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
    add_library(libjemalloc5_jet SHARED ${SOURCE_FILES})

    target_include_directories(libjemalloc5_jet
            PUBLIC
            ${ANDROID_ROOT}/external/jemalloc_new/src
            ${ANDROID_ROOT}/external/jemalloc_new/include
            ${ANDROID_ROOT}/external/jemalloc_new
    )

    # Add c/c++ flags
    target_compile_options(libjemalloc5_jet PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(libjemalloc5_jet
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libjemalloc5_jet ####################

#################### libjemalloc5 ####################
set(FILES
        ${ANDROID_ROOT}/external/jemalloc_new/src/jemalloc.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/arena.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/background_thread.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/base.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/bin.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/bitmap.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/ckh.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/ctl.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/div.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/extent.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/extent_dss.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/extent_mmap.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/hash.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/hooks.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/large.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/log.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/malloc_io.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/mutex.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/mutex_pool.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/nstime.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/pages.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/prng.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/prof.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/rtree.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/stats.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/sz.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/tcache.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/ticker.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/tsd.c
        ${ANDROID_ROOT}/external/jemalloc_new/src/witness.c
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
    add_library(libjemalloc5 SHARED ${SOURCE_FILES})

    target_include_directories(libjemalloc5
            PUBLIC
            ${ANDROID_ROOT}/external/jemalloc_new/include
            ${ANDROID_ROOT}/external/jemalloc_new/src
            ${ANDROID_ROOT}/external/jemalloc_new
    )

    # Add c/c++ flags
    target_compile_options(libjemalloc5 PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(libjemalloc5
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libjemalloc5 ####################

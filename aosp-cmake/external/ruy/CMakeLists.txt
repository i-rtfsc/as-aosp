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

project(libruy)

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

#################### libruy ####################
set(FILES
        ${ANDROID_ROOT}/external/ruy/ruy/allocator.cc
        ${ANDROID_ROOT}/external/ruy/ruy/apply_multiplier.cc
        ${ANDROID_ROOT}/external/ruy/ruy/block_map.cc
        ${ANDROID_ROOT}/external/ruy/ruy/blocking_counter.cc
        ${ANDROID_ROOT}/external/ruy/ruy/context.cc
        ${ANDROID_ROOT}/external/ruy/ruy/context_get_ctx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/cpuinfo.cc
        ${ANDROID_ROOT}/external/ruy/ruy/ctx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/denormal.cc
        ${ANDROID_ROOT}/external/ruy/ruy/frontend.cc
        ${ANDROID_ROOT}/external/ruy/ruy/have_built_path_for_avx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/have_built_path_for_avx2_fma.cc
        ${ANDROID_ROOT}/external/ruy/ruy/have_built_path_for_avx512.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_arm32.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_arm64.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_avx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_avx2_fma.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_avx512.cc
        ${ANDROID_ROOT}/external/ruy/ruy/pack_arm.cc
        ${ANDROID_ROOT}/external/ruy/ruy/pack_avx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/pack_avx2_fma.cc
        ${ANDROID_ROOT}/external/ruy/ruy/pack_avx512.cc
        ${ANDROID_ROOT}/external/ruy/ruy/prepacked_cache.cc
        ${ANDROID_ROOT}/external/ruy/ruy/prepare_packed_matrices.cc
        ${ANDROID_ROOT}/external/ruy/ruy/system_aligned_alloc.cc
        ${ANDROID_ROOT}/external/ruy/ruy/thread_pool.cc
        ${ANDROID_ROOT}/external/ruy/ruy/trmul.cc
        ${ANDROID_ROOT}/external/ruy/ruy/tune.cc
        ${ANDROID_ROOT}/external/ruy/ruy/wait.cc
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
    add_library(libruy SHARED ${SOURCE_FILES})

    target_include_directories(libruy
            PUBLIC
            ${ANDROID_ROOT}/external/ruy
    )

    # Add c/c++ flags
    target_compile_options(libruy PRIVATE
            -Werror
    )

    target_link_libraries(libruy
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libruy ####################

#################### libruy_static ####################
set(FILES
        ${ANDROID_ROOT}/external/ruy/ruy/allocator.cc
        ${ANDROID_ROOT}/external/ruy/ruy/apply_multiplier.cc
        ${ANDROID_ROOT}/external/ruy/ruy/block_map.cc
        ${ANDROID_ROOT}/external/ruy/ruy/blocking_counter.cc
        ${ANDROID_ROOT}/external/ruy/ruy/context.cc
        ${ANDROID_ROOT}/external/ruy/ruy/context_get_ctx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/cpuinfo.cc
        ${ANDROID_ROOT}/external/ruy/ruy/ctx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/denormal.cc
        ${ANDROID_ROOT}/external/ruy/ruy/frontend.cc
        ${ANDROID_ROOT}/external/ruy/ruy/have_built_path_for_avx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/have_built_path_for_avx2_fma.cc
        ${ANDROID_ROOT}/external/ruy/ruy/have_built_path_for_avx512.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_arm32.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_arm64.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_avx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_avx2_fma.cc
        ${ANDROID_ROOT}/external/ruy/ruy/kernel_avx512.cc
        ${ANDROID_ROOT}/external/ruy/ruy/pack_arm.cc
        ${ANDROID_ROOT}/external/ruy/ruy/pack_avx.cc
        ${ANDROID_ROOT}/external/ruy/ruy/pack_avx2_fma.cc
        ${ANDROID_ROOT}/external/ruy/ruy/pack_avx512.cc
        ${ANDROID_ROOT}/external/ruy/ruy/prepacked_cache.cc
        ${ANDROID_ROOT}/external/ruy/ruy/prepare_packed_matrices.cc
        ${ANDROID_ROOT}/external/ruy/ruy/system_aligned_alloc.cc
        ${ANDROID_ROOT}/external/ruy/ruy/thread_pool.cc
        ${ANDROID_ROOT}/external/ruy/ruy/trmul.cc
        ${ANDROID_ROOT}/external/ruy/ruy/tune.cc
        ${ANDROID_ROOT}/external/ruy/ruy/wait.cc
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
    add_library(libruy_static SHARED ${SOURCE_FILES})

    target_include_directories(libruy_static
            PUBLIC
            ${ANDROID_ROOT}/external/ruy
    )

    # Add c/c++ flags
    target_compile_options(libruy_static PRIVATE
            -Werror
    )

    target_link_libraries(libruy_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libruy_static ####################

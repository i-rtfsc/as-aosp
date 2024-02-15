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

project(icing-c-proto)

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

#################### icing-c-proto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/index/numeric/wildcard-property-storage.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/legacy/index/proto/icing-dynamic-trie-header.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/debug.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/document.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/document_wrapper.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/initialize.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/logging.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/optimize.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/persist.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/reset.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/schema.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/scoring.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/search.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/status.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/storage.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/term.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/usage.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto/icing/proto/internal/optimize.pb.cc
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
    add_library(icing-c-proto SHARED ${SOURCE_FILES})

    target_include_directories(icing-c-proto
            PUBLIC
            ${ANDROID_ROOT}/external/icing/proto
            ${OUT_INTERMEDIATES_ROOT}/external/icing/proto/icing-c-proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(icing-c-proto PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(icing-c-proto
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### icing-c-proto ####################

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

project(service.incremental)

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

#################### service.incremental ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/services/incremental/incremental_service.c
        ${ANDROID_ROOT}/frameworks/base/services/incremental/IncrementalService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/incremental/IncrementalServiceValidation.cpp
        ${ANDROID_ROOT}/frameworks/base/services/incremental/BinderIncrementalService.cpp
        ${ANDROID_ROOT}/frameworks/base/services/incremental/path.cpp
        ${ANDROID_ROOT}/frameworks/base/services/incremental/ServiceWrappers.cpp
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
    add_library(service.incremental SHARED ${SOURCE_FILES})

    target_include_directories(service.incremental
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/services/incremental/include
            ${ANDROID_ROOT}/frameworks/base/services/incremental
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/services/incremental/service.incremental.proto/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/services/incremental
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/services/incremental/service.incremental.proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(service.incremental PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(service.incremental
            jni_headers
            libandroidfw
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdataloader_aidl-cpp
            libdl
            libext2_uuid
            libincfs
            libincremental_aidl-cpp
            libincremental_manager_aidl-cpp
            liblog
            libm
            libpermission
            libprotobuf-cpp-lite
            libutils
            libvold_binder
            libz
            libziparchive_for_incfs
            service.incremental.proto
    )
endif ()
#################### service.incremental ####################

#################### service.incremental.proto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/services/incremental/service.incremental.proto/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/services/incremental/Metadata.pb.cc
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
    add_library(service.incremental.proto SHARED ${SOURCE_FILES})

    target_include_directories(service.incremental.proto
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/services/incremental
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/services/incremental/service.incremental.proto/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/base/services/incremental
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/services/incremental/service.incremental.proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(service.incremental.proto PRIVATE
            -Werror
    )

    target_link_libraries(service.incremental.proto
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### service.incremental.proto ####################

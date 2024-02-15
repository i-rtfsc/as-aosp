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

project(libnosprotos)

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

#################### libnosprotos ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/avb/avb.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/identity/identity.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/identity/identity_defs.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/identity/identity_types.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/keymaster/keymaster.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/keymaster/keymaster_defs.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/keymaster/keymaster_types.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/protoapi/control.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/protoapi/diagnostics_api.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/protoapi/gchips_types.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/protoapi/header.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/protoapi/testing_api.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/app/weaver/weaver.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/nugget/protobuf/options.pb.cc
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
    add_library(libnosprotos SHARED ${SOURCE_FILES})

    target_include_directories(libnosprotos
            PUBLIC
            ${ANDROID_ROOT}/external/nos/host/generic/nugget/proto
            ${OUT_INTERMEDIATES_ROOT}/external/nos/host/generic/nugget/proto/libnosprotos/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libnosprotos PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libnosprotos
            libc++
            libc++
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-full
            libprotobuf-cpp-full
            libc
            libdl
            libm
    )
endif ()
#################### libnosprotos ####################

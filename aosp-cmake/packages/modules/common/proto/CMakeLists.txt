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

project(libclasspaths_proto)

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

#################### libclasspaths_proto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/common/proto/libclasspaths_proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/common/proto/classpaths.pb.cc
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
    add_library(libclasspaths_proto SHARED ${SOURCE_FILES})

    target_include_directories(libclasspaths_proto
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/common/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/common/proto/libclasspaths_proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/common/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/common/proto/libclasspaths_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libclasspaths_proto PRIVATE
            -Werror
    )

    target_link_libraries(libclasspaths_proto
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### libclasspaths_proto ####################

#################### libsdk_proto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/common/proto/libsdk_proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/common/proto/sdk.pb.cc
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
    add_library(libsdk_proto SHARED ${SOURCE_FILES})

    target_include_directories(libsdk_proto
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/common/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/common/proto/libsdk_proto/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/common/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/common/proto/libsdk_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libsdk_proto PRIVATE
            -Werror
    )

    target_link_libraries(libsdk_proto
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### libsdk_proto ####################

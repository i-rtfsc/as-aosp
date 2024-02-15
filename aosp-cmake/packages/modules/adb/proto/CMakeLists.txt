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

project(libapp_processes_protos_lite)

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

#################### libapp_processes_protos_lite ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto/app_processes.pb.cc
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
    add_library(libapp_processes_protos_lite SHARED ${SOURCE_FILES})

    target_include_directories(libapp_processes_protos_lite
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libapp_processes_protos_lite/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libapp_processes_protos_lite PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libapp_processes_protos_lite
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
            libwinpthread
    )
endif ()
#################### libapp_processes_protos_lite ####################

#################### libadb_protos_static ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos_static/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto/adb_known_hosts.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos_static/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto/key_type.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos_static/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto/pairing.pb.cc
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
    add_library(libadb_protos_static SHARED ${SOURCE_FILES})

    target_include_directories(libadb_protos_static
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos_static/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos_static/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libadb_protos_static PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libadb_protos_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
            libwinpthread
    )
endif ()
#################### libadb_protos_static ####################

#################### libadb_protos ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto/adb_known_hosts.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto/key_type.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto/pairing.pb.cc
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
    add_library(libadb_protos SHARED ${SOURCE_FILES})

    target_include_directories(libadb_protos
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/adb/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/adb/proto/libadb_protos/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libadb_protos PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libadb_protos
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
            libwinpthread
    )
endif ()
#################### libadb_protos ####################

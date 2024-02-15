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

project(libnetlink)

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

#################### libnetlink ####################
set(FILES
        ${ANDROID_ROOT}/external/iproute2/lib/ll_map.c
        ${ANDROID_ROOT}/external/iproute2/lib/libnetlink.c
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
    add_library(libnetlink SHARED ${SOURCE_FILES})

    target_include_directories(libnetlink
            PUBLIC
            ${ANDROID_ROOT}/external/iproute2/lib
    )

    # Add c/c++ flags
    target_compile_options(libnetlink PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libnetlink
            iproute2_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
    )
endif ()
#################### libnetlink ####################

#################### libiprouteutil ####################
set(FILES
        ${ANDROID_ROOT}/external/iproute2/lib/color.c
        ${ANDROID_ROOT}/external/iproute2/lib/utils.c
        ${ANDROID_ROOT}/external/iproute2/lib/rt_names.c
        ${ANDROID_ROOT}/external/iproute2/lib/ll_types.c
        ${ANDROID_ROOT}/external/iproute2/lib/ll_proto.c
        ${ANDROID_ROOT}/external/iproute2/lib/ll_addr.c
        ${ANDROID_ROOT}/external/iproute2/lib/inet_proto.c
        ${ANDROID_ROOT}/external/iproute2/lib/mpls_pton.c
        ${ANDROID_ROOT}/external/iproute2/lib/namespace.c
        ${ANDROID_ROOT}/external/iproute2/lib/names.c
        ${ANDROID_ROOT}/external/iproute2/lib/libgenl.c
        ${ANDROID_ROOT}/external/iproute2/lib/libnetlink.c
        ${ANDROID_ROOT}/external/iproute2/lib/bpf.c
        ${ANDROID_ROOT}/external/iproute2/lib/exec.c
        ${ANDROID_ROOT}/external/iproute2/lib/fs.c
        ${ANDROID_ROOT}/external/iproute2/lib/json_print.c
        ${ANDROID_ROOT}/external/iproute2/lib/json_writer.c
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
    add_library(libiprouteutil SHARED ${SOURCE_FILES})

    target_include_directories(libiprouteutil
            PUBLIC
            ${ANDROID_ROOT}/external/iproute2/lib
    )

    # Add c/c++ flags
    target_compile_options(libiprouteutil PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libiprouteutil
            iproute2_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
    )
endif ()
#################### libiprouteutil ####################

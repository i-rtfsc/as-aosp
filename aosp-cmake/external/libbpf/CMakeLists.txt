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

project(btfloader)

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

#################### libbpf_minimal ####################
set(FILES
        ${ANDROID_ROOT}/external/libbpf/android/bpf_stub.c
        ${ANDROID_ROOT}/external/libbpf/src/bpf.c
        ${ANDROID_ROOT}/external/libbpf/src/libbpf_errno.c
        ${ANDROID_ROOT}/external/libbpf/src/netlink.c
        ${ANDROID_ROOT}/external/libbpf/src/nlattr.c
        ${ANDROID_ROOT}/external/libbpf/src/ringbuf.c
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
    add_library(libbpf_minimal SHARED ${SOURCE_FILES})

    target_include_directories(libbpf_minimal
            PUBLIC
            ${ANDROID_ROOT}/external/libbpf/include/uapi
            ${ANDROID_ROOT}/external/libbpf/android
            ${ANDROID_ROOT}/external/libbpf/include
            ${ANDROID_ROOT}/external/libbpf
            ${OUT_INTERMEDIATES_ROOT}/external/libbpf/libbpf_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(libbpf_minimal PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libbpf_minimal
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libbpf_minimal ####################

#################### btfloader ####################
set(FILES
        ${ANDROID_ROOT}/external/libbpf/android/btfloader.cpp
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
    add_executable(btfloader ${SOURCE_FILES})

    target_include_directories(btfloader
            PUBLIC
            ${ANDROID_ROOT}/external/libbpf
            ${ANDROID_ROOT}/external/libbpf/include/uapi
            ${OUT_INTERMEDIATES_ROOT}/external/libbpf/libbpf_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(btfloader PRIVATE
            -Werror
    )

    target_link_libraries(btfloader
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### btfloader ####################

#################### libbpf ####################
set(FILES
        ${ANDROID_ROOT}/external/libbpf/src/bpf.c
        ${ANDROID_ROOT}/external/libbpf/src/bpf_prog_linfo.c
        ${ANDROID_ROOT}/external/libbpf/src/btf.c
        ${ANDROID_ROOT}/external/libbpf/src/btf_dump.c
        ${ANDROID_ROOT}/external/libbpf/src/gen_loader.c
        ${ANDROID_ROOT}/external/libbpf/src/hashmap.c
        ${ANDROID_ROOT}/external/libbpf/src/libbpf.c
        ${ANDROID_ROOT}/external/libbpf/src/libbpf_errno.c
        ${ANDROID_ROOT}/external/libbpf/src/libbpf_probes.c
        ${ANDROID_ROOT}/external/libbpf/src/linker.c
        ${ANDROID_ROOT}/external/libbpf/src/netlink.c
        ${ANDROID_ROOT}/external/libbpf/src/nlattr.c
        ${ANDROID_ROOT}/external/libbpf/src/relo_core.c
        ${ANDROID_ROOT}/external/libbpf/src/ringbuf.c
        ${ANDROID_ROOT}/external/libbpf/src/str_error.c
        ${ANDROID_ROOT}/external/libbpf/src/strset.c
        ${ANDROID_ROOT}/external/libbpf/src/xsk.c
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
    add_library(libbpf SHARED ${SOURCE_FILES})

    target_include_directories(libbpf
            PUBLIC
            ${ANDROID_ROOT}/external/libbpf/include/uapi
            ${ANDROID_ROOT}/external/libbpf/android
            ${ANDROID_ROOT}/external/libbpf/include
            ${ANDROID_ROOT}/external/libbpf
            ${OUT_INTERMEDIATES_ROOT}/external/libbpf/libbpf_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(libbpf PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libbpf
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### libbpf ####################

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

project(linker)

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

#################### linker ####################
set(FILES
        ${ANDROID_ROOT}/bionic/linker/dlfcn.cpp
        ${ANDROID_ROOT}/bionic/linker/linker.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_block_allocator.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_dlwarning.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_cfi.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_config.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_debug.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_gdb_support.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_globals.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_libc_support.c
        ${ANDROID_ROOT}/bionic/linker/linker_libcxx_support.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_namespaces.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_logger.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_mapped_file_fragment.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_note_gnu_property.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_phdr.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_relocate.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_sdk_versions.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_soinfo.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_transparent_hugepage_support.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_tls.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_utils.cpp
        ${ANDROID_ROOT}/bionic/linker/rt.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_translate_path.cpp
        ${ANDROID_ROOT}/bionic/linker/linker_debuggerd_android.cpp
        ${ANDROID_ROOT}/bionic/linker/arch/arm64/begin.S
        ${ANDROID_ROOT}/bionic/linker/arch/arm64/tlsdesc_resolver.S
        ${ANDROID_ROOT}/bionic/linker/arch/arm_neon/linker_gnu_hash_neon.cpp
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
    add_executable(linker ${SOURCE_FILES})

    target_include_directories(linker
            PUBLIC
            ${ANDROID_ROOT}/bionic/linker
    )

    # Add c/c++ flags
    target_compile_options(linker PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(linker
            libasync_safe
            libbase
            libc++_static
            libc++demangle
            libc_dynamic_dispatch
            libc_nomalloc
            libc_unwind_static
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdebuggerd_handler_fallback
            liblinker_main
            liblinker_malloc
            liblog_for_runtime_apex
            libm
            libunwind
            libz
            libziparchive
    )
endif ()
#################### linker ####################

#################### liblinker_malloc ####################
set(FILES
        ${ANDROID_ROOT}/bionic/linker/linker_memory.cpp
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
    add_library(liblinker_malloc SHARED ${SOURCE_FILES})

    target_include_directories(liblinker_malloc
            PUBLIC
            ${ANDROID_ROOT}/bionic/linker
    )

    # Add c/c++ flags
    target_compile_options(liblinker_malloc PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(liblinker_malloc
            libasync_safe
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog_for_runtime_apex
            libm
            libz
            libziparchive
    )
endif ()
#################### liblinker_malloc ####################

#################### liblinker_debuggerd_stub ####################
set(FILES
        ${ANDROID_ROOT}/bionic/linker/linker_debuggerd_stub.cpp
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
    add_library(liblinker_debuggerd_stub SHARED ${SOURCE_FILES})

    target_include_directories(liblinker_debuggerd_stub
            PUBLIC
            ${ANDROID_ROOT}/bionic/linker
    )

    # Add c/c++ flags
    target_compile_options(liblinker_debuggerd_stub PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(liblinker_debuggerd_stub
            libasync_safe
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog_for_runtime_apex
            libm
            libz
            libziparchive
    )
endif ()
#################### liblinker_debuggerd_stub ####################

#################### liblinker_main ####################
set(FILES
        ${ANDROID_ROOT}/bionic/linker/linker_main.cpp
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
    add_library(liblinker_main SHARED ${SOURCE_FILES})

    target_include_directories(liblinker_main
            PUBLIC
            ${ANDROID_ROOT}/bionic/linker
    )

    # Add c/c++ flags
    target_compile_options(liblinker_main PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(liblinker_main
            libasync_safe
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog_for_runtime_apex
            libm
            libz
            libziparchive
    )
endif ()
#################### liblinker_main ####################

#################### ld-android ####################
set(FILES
        ${ANDROID_ROOT}/bionic/linker/ld_android.cpp
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
    add_library(ld-android SHARED ${SOURCE_FILES})

    target_include_directories(ld-android
            PUBLIC
            ${ANDROID_ROOT}/bionic/linker
    )

    # Add c/c++ flags
    target_compile_options(ld-android PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(ld-android
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### ld-android ####################

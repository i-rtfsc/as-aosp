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

project(libziparchive_fuzzer)

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

#################### libziparchive_for_incfs_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/libziparchive/libziparchive_fuzzer.cpp
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
    add_executable(libziparchive_for_incfs_fuzzer ${SOURCE_FILES})

    target_include_directories(libziparchive_for_incfs_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/libziparchive
            ${ANDROID_ROOT}/system/libziparchive/incfs_support/include
            ${ANDROID_ROOT}/system/libziparchive/include
    )

    # Add c/c++ flags
    target_compile_options(libziparchive_for_incfs_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libziparchive_for_incfs_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libz
            libziparchive_for_incfs
    )
endif ()
#################### libziparchive_for_incfs_fuzzer ####################

#################### libziparchive_for_incfs ####################
set(FILES
        ${ANDROID_ROOT}/system/libziparchive/zip_archive.cc
        ${ANDROID_ROOT}/system/libziparchive/zip_archive_stream_entry.cc
        ${ANDROID_ROOT}/system/libziparchive/zip_cd_entry_map.cc
        ${ANDROID_ROOT}/system/libziparchive/zip_error.cpp
        ${ANDROID_ROOT}/system/libziparchive/zip_writer.cc
        ${ANDROID_ROOT}/system/libziparchive/incfs_support/signal_handling.cpp
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
    add_library(libziparchive_for_incfs SHARED ${SOURCE_FILES})

    target_include_directories(libziparchive_for_incfs
            PUBLIC
            ${ANDROID_ROOT}/system/libziparchive/incfs_support/include
            ${ANDROID_ROOT}/system/libziparchive/include
            ${ANDROID_ROOT}/system/libziparchive
    )

    # Add c/c++ flags
    target_compile_options(libziparchive_for_incfs PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libziparchive_for_incfs
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest_prod_headers
            liblog
            libm
            libwinpthread
            libz
    )
endif ()
#################### libziparchive_for_incfs ####################

#################### libziparchive_writer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/libziparchive/libziparchive_writer_fuzzer.cpp
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
    add_executable(libziparchive_writer_fuzzer ${SOURCE_FILES})

    target_include_directories(libziparchive_writer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/libziparchive
            ${ANDROID_ROOT}/system/libziparchive/include
    )

    # Add c/c++ flags
    target_compile_options(libziparchive_writer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libziparchive_writer_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libz
            libziparchive
    )
endif ()
#################### libziparchive_writer_fuzzer ####################

#################### ziptool ####################
set(FILES
        ${ANDROID_ROOT}/system/libziparchive/ziptool.cpp
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
    add_executable(ziptool ${SOURCE_FILES})

    target_include_directories(ziptool
            PUBLIC
            ${ANDROID_ROOT}/system/libziparchive
            ${ANDROID_ROOT}/system/libziparchive/include
    )

    # Add c/c++ flags
    target_compile_options(ziptool PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(ziptool
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libz
            libziparchive
    )
endif ()
#################### ziptool ####################

#################### libziparchive ####################
set(FILES
        ${ANDROID_ROOT}/system/libziparchive/zip_archive.cc
        ${ANDROID_ROOT}/system/libziparchive/zip_archive_stream_entry.cc
        ${ANDROID_ROOT}/system/libziparchive/zip_cd_entry_map.cc
        ${ANDROID_ROOT}/system/libziparchive/zip_error.cpp
        ${ANDROID_ROOT}/system/libziparchive/zip_writer.cc
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
    add_library(libziparchive SHARED ${SOURCE_FILES})

    target_include_directories(libziparchive
            PUBLIC
            ${ANDROID_ROOT}/system/libziparchive/include
            ${ANDROID_ROOT}/system/libziparchive/incfs_support/include
            ${ANDROID_ROOT}/system/libziparchive
    )

    # Add c/c++ flags
    target_compile_options(libziparchive PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libziparchive
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest_prod_headers
            liblog
            libm
            libwinpthread
            libz
    )
endif ()
#################### libziparchive ####################

#################### libziparchive_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/libziparchive/libziparchive_fuzzer.cpp
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
    add_executable(libziparchive_fuzzer ${SOURCE_FILES})

    target_include_directories(libziparchive_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/libziparchive
            ${ANDROID_ROOT}/system/libziparchive/include
    )

    # Add c/c++ flags
    target_compile_options(libziparchive_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libziparchive_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libz
            libziparchive
    )
endif ()
#################### libziparchive_fuzzer ####################

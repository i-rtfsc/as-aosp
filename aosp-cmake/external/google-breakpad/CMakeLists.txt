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

project(core2md)

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

#################### core2md ####################
set(FILES
        ${ANDROID_ROOT}/external/google-breakpad/src/common/convert_UTF.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/path_helper.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/tools/linux/core2md/core2md.cc
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
    add_executable(core2md ${SOURCE_FILES})

    target_include_directories(core2md
            PUBLIC
            ${ANDROID_ROOT}/external/google-breakpad
            ${ANDROID_ROOT}/external/google-breakpad/src
    )

    # Add c/c++ flags
    target_compile_options(core2md PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(core2md
            breakpad_client
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### core2md ####################

#################### breakpad_client ####################
set(FILES
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/crash_generation/crash_generation_client.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/dump_writer_common/thread_info.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/dump_writer_common/ucontext_reader.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/handler/exception_handler.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/handler/minidump_descriptor.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/log/log.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/microdump_writer/microdump_writer.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/minidump_writer/linux_core_dumper.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/minidump_writer/linux_dumper.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/minidump_writer/linux_ptrace_dumper.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/linux/minidump_writer/minidump_writer.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/client/minidump_file_writer.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/convert_UTF.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/linux/breakpad_getcontext.S
        ${ANDROID_ROOT}/external/google-breakpad/src/common/linux/elf_core_dump.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/linux/elfutils.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/linux/file_id.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/linux/guid_creator.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/linux/linux_libc_support.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/linux/memory_mapped_file.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/linux/safe_readlink.cc
        ${ANDROID_ROOT}/external/google-breakpad/src/common/string_conversion.cc
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
    add_library(breakpad_client SHARED ${SOURCE_FILES})

    target_include_directories(breakpad_client
            PUBLIC
            ${ANDROID_ROOT}/external/google-breakpad/src
            ${ANDROID_ROOT}/external/google-breakpad/src/common/android/include
            ${ANDROID_ROOT}/external/google-breakpad
    )

    # Add c/c++ flags
    target_compile_options(breakpad_client PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(breakpad_client
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### breakpad_client ####################

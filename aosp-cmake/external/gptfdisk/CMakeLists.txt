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

project(sgdisk)

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

#################### sgdisk ####################
set(FILES
        ${ANDROID_ROOT}/external/gptfdisk/sgdisk.cc
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
    add_executable(sgdisk ${SOURCE_FILES})

    target_include_directories(sgdisk
            PUBLIC
            ${ANDROID_ROOT}/external/gptfdisk
    )

    # Add c/c++ flags
    target_compile_options(sgdisk PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(sgdisk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_uuid
            libgptf
            libm
    )
endif ()
#################### sgdisk ####################

#################### libgptf_fuzzer_lib ####################
set(FILES
        ${ANDROID_ROOT}/external/gptfdisk/gptcl.cc
        ${ANDROID_ROOT}/external/gptfdisk/crc32.cc
        ${ANDROID_ROOT}/external/gptfdisk/support.cc
        ${ANDROID_ROOT}/external/gptfdisk/guid.cc
        ${ANDROID_ROOT}/external/gptfdisk/gptpart.cc
        ${ANDROID_ROOT}/external/gptfdisk/mbrpart.cc
        ${ANDROID_ROOT}/external/gptfdisk/basicmbr.cc
        ${ANDROID_ROOT}/external/gptfdisk/mbr.cc
        ${ANDROID_ROOT}/external/gptfdisk/gpt.cc
        ${ANDROID_ROOT}/external/gptfdisk/bsd.cc
        ${ANDROID_ROOT}/external/gptfdisk/parttypes.cc
        ${ANDROID_ROOT}/external/gptfdisk/attributes.cc
        ${ANDROID_ROOT}/external/gptfdisk/diskio.cc
        ${ANDROID_ROOT}/external/gptfdisk/android_popt.cc
        ${ANDROID_ROOT}/external/gptfdisk/diskio-heap.cc
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
    add_library(libgptf_fuzzer_lib SHARED ${SOURCE_FILES})

    target_include_directories(libgptf_fuzzer_lib
            PUBLIC
            ${ANDROID_ROOT}/external/gptfdisk
    )

    # Add c/c++ flags
    target_compile_options(libgptf_fuzzer_lib PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libgptf_fuzzer_lib
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_uuid
            libm
    )
endif ()
#################### libgptf_fuzzer_lib ####################

#################### libgptf ####################
set(FILES
        ${ANDROID_ROOT}/external/gptfdisk/gptcl.cc
        ${ANDROID_ROOT}/external/gptfdisk/crc32.cc
        ${ANDROID_ROOT}/external/gptfdisk/support.cc
        ${ANDROID_ROOT}/external/gptfdisk/guid.cc
        ${ANDROID_ROOT}/external/gptfdisk/gptpart.cc
        ${ANDROID_ROOT}/external/gptfdisk/mbrpart.cc
        ${ANDROID_ROOT}/external/gptfdisk/basicmbr.cc
        ${ANDROID_ROOT}/external/gptfdisk/mbr.cc
        ${ANDROID_ROOT}/external/gptfdisk/gpt.cc
        ${ANDROID_ROOT}/external/gptfdisk/bsd.cc
        ${ANDROID_ROOT}/external/gptfdisk/parttypes.cc
        ${ANDROID_ROOT}/external/gptfdisk/attributes.cc
        ${ANDROID_ROOT}/external/gptfdisk/diskio.cc
        ${ANDROID_ROOT}/external/gptfdisk/android_popt.cc
        ${ANDROID_ROOT}/external/gptfdisk/diskio-unix.cc
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
    add_library(libgptf SHARED ${SOURCE_FILES})

    target_include_directories(libgptf
            PUBLIC
            ${ANDROID_ROOT}/external/gptfdisk
    )

    # Add c/c++ flags
    target_compile_options(libgptf PRIVATE
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libgptf
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libext2_uuid
            libm
    )
endif ()
#################### libgptf ####################

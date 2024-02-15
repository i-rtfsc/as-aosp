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

project(ogg_writer_fuzzer)

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

#################### webm_writer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/webm_writer_fuzzer.cpp
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
    add_executable(webm_writer_fuzzer ${SOURCE_FILES})

    target_include_directories(webm_writer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers
    )

    # Add c/c++ flags
    target_compile_options(webm_writer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(webm_writer_fuzzer
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_minimal
            libclang_rt.ubsan_standalone
            libcutils
            libdatasource
            libdl
            liblog
            libm
            libstagefright
            libstagefright_foundation
            libstagefright_headers
            libstagefright_webm
            libutils
            libwriterfuzzerbase
    )
endif ()
#################### webm_writer_fuzzer ####################

#################### mpeg4_writer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/mpeg4_writer_fuzzer.cpp
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
    add_executable(mpeg4_writer_fuzzer ${SOURCE_FILES})

    target_include_directories(mpeg4_writer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers
    )

    # Add c/c++ flags
    target_compile_options(mpeg4_writer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(mpeg4_writer_fuzzer
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libstagefright
            libstagefright_esds
            libstagefright_foundation
            libutils
            libwriterfuzzerbase
    )
endif ()
#################### mpeg4_writer_fuzzer ####################

#################### amrwb_writer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/amr_writer_fuzzer.cpp
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
    add_executable(amrwb_writer_fuzzer ${SOURCE_FILES})

    target_include_directories(amrwb_writer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers
    )

    # Add c/c++ flags
    target_compile_options(amrwb_writer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(amrwb_writer_fuzzer
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libstagefright
            libstagefright_foundation
            libutils
            libwriterfuzzerbase
    )
endif ()
#################### amrwb_writer_fuzzer ####################

#################### libwriterfuzzerbase ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/WriterFuzzerBase.cpp
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
    add_library(libwriterfuzzerbase SHARED ${SOURCE_FILES})

    target_include_directories(libwriterfuzzerbase
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers
    )

    # Add c/c++ flags
    target_compile_options(libwriterfuzzerbase PRIVATE
            -Werror
    )

    target_link_libraries(libwriterfuzzerbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libstagefright
            libstagefright_foundation
            libutils
    )
endif ()
#################### libwriterfuzzerbase ####################

#################### amrnb_writer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/amr_writer_fuzzer.cpp
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
    add_executable(amrnb_writer_fuzzer ${SOURCE_FILES})

    target_include_directories(amrnb_writer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers
    )

    # Add c/c++ flags
    target_compile_options(amrnb_writer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(amrnb_writer_fuzzer
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libstagefright
            libstagefright_foundation
            libutils
            libwriterfuzzerbase
    )
endif ()
#################### amrnb_writer_fuzzer ####################

#################### ogg_writer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/ogg_writer_fuzzer.cpp
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
    add_executable(ogg_writer_fuzzer ${SOURCE_FILES})

    target_include_directories(ogg_writer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/writer_fuzzers
    )

    # Add c/c++ flags
    target_compile_options(ogg_writer_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(ogg_writer_fuzzer
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libogg
            libstagefright
            libstagefright_foundation
            libutils
            libwriterfuzzerbase
    )
endif ()
#################### ogg_writer_fuzzer ####################

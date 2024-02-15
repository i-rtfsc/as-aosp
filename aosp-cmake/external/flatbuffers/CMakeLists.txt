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

project(flatbuffers_parser_fuzzer)

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

#################### flatbuffers_verifier_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/flatbuffers/tests/test_assert.cpp
        ${ANDROID_ROOT}/external/flatbuffers/tests/fuzzer/flatbuffers_verifier_fuzzer.cc
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
    add_executable(flatbuffers_verifier_fuzzer ${SOURCE_FILES})

    target_include_directories(flatbuffers_verifier_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/flatbuffers/tests
            ${ANDROID_ROOT}/external/flatbuffers/tests/fuzzer
            ${ANDROID_ROOT}/external/flatbuffers
            ${ANDROID_ROOT}/external/flatbuffers/include
    )

    # Add c/c++ flags
    target_compile_options(flatbuffers_verifier_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(flatbuffers_verifier_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libflatbuffers-cpp
            liblog
            libm
    )
endif ()
#################### flatbuffers_verifier_fuzzer ####################

#################### libflatbuffers-cpp ####################
set(FILES
        ${ANDROID_ROOT}/external/flatbuffers/src/code_generators.cpp
        ${ANDROID_ROOT}/external/flatbuffers/src/idl_gen_fbs.cpp
        ${ANDROID_ROOT}/external/flatbuffers/src/idl_gen_text.cpp
        ${ANDROID_ROOT}/external/flatbuffers/src/idl_parser.cpp
        ${ANDROID_ROOT}/external/flatbuffers/src/reflection.cpp
        ${ANDROID_ROOT}/external/flatbuffers/src/util.cpp
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
    add_library(libflatbuffers-cpp SHARED ${SOURCE_FILES})

    target_include_directories(libflatbuffers-cpp
            PUBLIC
            ${ANDROID_ROOT}/external/flatbuffers/include
            ${ANDROID_ROOT}/external/flatbuffers
    )

    # Add c/c++ flags
    target_compile_options(libflatbuffers-cpp PRIVATE
            -Werror
    )

    target_link_libraries(libflatbuffers-cpp
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libflatbuffers-cpp ####################

#################### flatbuffers_scalar_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/flatbuffers/tests/test_assert.cpp
        ${ANDROID_ROOT}/external/flatbuffers/tests/fuzzer/flatbuffers_scalar_fuzzer.cc
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
    add_executable(flatbuffers_scalar_fuzzer ${SOURCE_FILES})

    target_include_directories(flatbuffers_scalar_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/flatbuffers/tests
            ${ANDROID_ROOT}/external/flatbuffers/tests/fuzzer
            ${ANDROID_ROOT}/external/flatbuffers
            ${ANDROID_ROOT}/external/flatbuffers/include
    )

    # Add c/c++ flags
    target_compile_options(flatbuffers_scalar_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(flatbuffers_scalar_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libflatbuffers-cpp
            liblog
            libm
    )
endif ()
#################### flatbuffers_scalar_fuzzer ####################

#################### flatbuffers_parser_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/flatbuffers/tests/test_assert.cpp
        ${ANDROID_ROOT}/external/flatbuffers/tests/fuzzer/flatbuffers_parser_fuzzer.cc
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
    add_executable(flatbuffers_parser_fuzzer ${SOURCE_FILES})

    target_include_directories(flatbuffers_parser_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/flatbuffers/tests
            ${ANDROID_ROOT}/external/flatbuffers/tests/fuzzer
            ${ANDROID_ROOT}/external/flatbuffers
            ${ANDROID_ROOT}/external/flatbuffers/include
    )

    # Add c/c++ flags
    target_compile_options(flatbuffers_parser_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(flatbuffers_parser_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libflatbuffers-cpp
            liblog
            libm
    )
endif ()
#################### flatbuffers_parser_fuzzer ####################

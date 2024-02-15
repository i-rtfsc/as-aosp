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

project(aidl_parser_fuzzer)

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

#################### libaidl-common ####################
set(FILES
        ${ANDROID_ROOT}/system/tools/aidl/aidl_checkapi.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl_const_expressions.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl_dumpapi.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/tools/aidl/libaidl-common/android_${OUT_ARCH_CPU}_static/gen/lex/system/tools/aidl/aidl_language_l.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/tools/aidl/libaidl-common/android_${OUT_ARCH_CPU}_static/gen/yacc/system/tools/aidl/aidl_language_y.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl_language.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl_to_cpp_common.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl_to_cpp.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl_to_java.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl_to_ndk.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl_to_rust.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl_typenames.cpp
        ${ANDROID_ROOT}/system/tools/aidl/aidl.cpp
        ${ANDROID_ROOT}/system/tools/aidl/ast_java.cpp
        ${ANDROID_ROOT}/system/tools/aidl/check_valid.cpp
        ${ANDROID_ROOT}/system/tools/aidl/code_writer.cpp
        ${ANDROID_ROOT}/system/tools/aidl/comments.cpp
        ${ANDROID_ROOT}/system/tools/aidl/diagnostics.cpp
        ${ANDROID_ROOT}/system/tools/aidl/generate_aidl_mappings.cpp
        ${ANDROID_ROOT}/system/tools/aidl/generate_cpp.cpp
        ${ANDROID_ROOT}/system/tools/aidl/generate_cpp_analyzer.cpp
        ${ANDROID_ROOT}/system/tools/aidl/generate_java_binder.cpp
        ${ANDROID_ROOT}/system/tools/aidl/generate_java.cpp
        ${ANDROID_ROOT}/system/tools/aidl/generate_ndk.cpp
        ${ANDROID_ROOT}/system/tools/aidl/generate_rust.cpp
        ${ANDROID_ROOT}/system/tools/aidl/import_resolver.cpp
        ${ANDROID_ROOT}/system/tools/aidl/io_delegate.cpp
        ${ANDROID_ROOT}/system/tools/aidl/location.cpp
        ${ANDROID_ROOT}/system/tools/aidl/logging.cpp
        ${ANDROID_ROOT}/system/tools/aidl/options.cpp
        ${ANDROID_ROOT}/system/tools/aidl/parser.cpp
        ${ANDROID_ROOT}/system/tools/aidl/permission.cpp
        ${ANDROID_ROOT}/system/tools/aidl/preprocess.cpp
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
    add_library(libaidl-common SHARED ${SOURCE_FILES})

    target_include_directories(libaidl-common
            PUBLIC
            ${ANDROID_ROOT}/system/tools/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/tools/aidl/libaidl-common/android_${OUT_ARCH_CPU}_static/gen/yacc/system/tools/aidl
    )

    # Add c/c++ flags
    target_compile_options(libaidl-common PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libaidl-common
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest
            libgtest_prod_headers
            libm
            libwinpthread
    )
endif ()
#################### libaidl-common ####################

#################### aidl_parser_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/tools/aidl/tests/aidl_parser_fuzzer.cpp
        ${ANDROID_ROOT}/system/tools/aidl/tests/fake_io_delegate.cpp
        ${ANDROID_ROOT}/system/tools/aidl/tests/test_util.cpp
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
    add_executable(aidl_parser_fuzzer ${SOURCE_FILES})

    target_include_directories(aidl_parser_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/tools/aidl
    )

    # Add c/c++ flags
    target_compile_options(aidl_parser_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(aidl_parser_fuzzer
            libaidl-common
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libgtest
            liblog
            libm
    )
endif ()
#################### aidl_parser_fuzzer ####################

#################### aidl-analyzer-main ####################
set(FILES
        ${ANDROID_ROOT}/system/tools/aidl/analyzer/analyzerMain.cpp
        ${ANDROID_ROOT}/system/tools/aidl/analyzer/Analyzer.cpp
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
    add_library(aidl-analyzer-main SHARED ${SOURCE_FILES})

    target_include_directories(aidl-analyzer-main
            PUBLIC
            ${ANDROID_ROOT}/system/tools/aidl/analyzer/include
            ${ANDROID_ROOT}/system/tools/aidl
    )

    # Add c/c++ flags
    target_compile_options(aidl-analyzer-main PRIVATE
            -Werror
    )

    target_link_libraries(aidl-analyzer-main
            libbase
            libbinder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### aidl-analyzer-main ####################

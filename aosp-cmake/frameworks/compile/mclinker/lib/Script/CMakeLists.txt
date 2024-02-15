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

project(libmcldScript)

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

#################### libmcldScript ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/AssertCmd.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/Assignment.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/BinaryOp.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/EntryCmd.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/FileToken.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/GroupCmd.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/InputCmd.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/InputSectDesc.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/InputToken.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/NameSpec.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/NullaryOp.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/Operand.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/Operator.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/OutputArchCmd.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/OutputCmd.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/OutputFormatCmd.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/OutputSectDesc.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/RpnEvaluator.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/RpnExpr.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/ScriptCommand.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/ScriptFile.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/ScriptReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/SearchDirCmd.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/SectionsCmd.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/compile/mclinker/lib/Script/libmcldScript/android_${OUT_ARCH_CPU}_static/gen/lex/frameworks/compile/mclinker/lib/Script/ScriptScanner.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/compile/mclinker/lib/Script/libmcldScript/android_${OUT_ARCH_CPU}_static/gen/yacc/frameworks/compile/mclinker/lib/Script/ScriptParser.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/StrToken.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/StringList.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/TernaryOp.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/UnaryOp.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script/WildcardPattern.cpp
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
    add_library(libmcldScript SHARED ${SOURCE_FILES})

    target_include_directories(libmcldScript
            PUBLIC
            ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/Script
            ${OUT_INTERMEDIATES_ROOT}/frameworks/compile/mclinker/lib/Script/libmcldScript/android_${OUT_ARCH_CPU}_static/gen/yacc/frameworks/compile/mclinker/lib/Script
    )

    # Add c/c++ flags
    target_compile_options(libmcldScript PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libmcldScript
            FlexLexer
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            llvm-headers
            llvm-headers-no-generated-headers
    )
endif ()
#################### libmcldScript ####################

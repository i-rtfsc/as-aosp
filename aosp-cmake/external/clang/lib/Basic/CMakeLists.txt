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

project(libclangBasic)

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

#################### libclangBasic ####################
set(FILES
        ${ANDROID_ROOT}/external/clang/lib/Basic/Attributes.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/Builtins.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/CharInfo.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/Cuda.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/Diagnostic.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/DiagnosticIDs.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/DiagnosticOptions.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/FileManager.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/FileSystemStatCache.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/IdentifierTable.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/LangOptions.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/Module.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/ObjCRuntime.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/OpenMPKinds.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/OperatorPrecedence.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/SanitizerBlacklist.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/Sanitizers.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/SourceLocation.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/SourceManager.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/TargetInfo.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/Targets.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/TokenKinds.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/Version.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/VersionTuple.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/VirtualFileSystem.cpp
        ${ANDROID_ROOT}/external/clang/lib/Basic/Warnings.cpp
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
    add_library(libclangBasic SHARED ${SOURCE_FILES})

    target_include_directories(libclangBasic
            PUBLIC
            ${ANDROID_ROOT}/external/clang/lib/Basic
    )

    # Add c/c++ flags
    target_compile_options(libclangBasic PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wno-unused-lambda-capture
            -Wno-unused-private-field
    )

    target_link_libraries(libclangBasic
            clang-headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            llvm-headers
            llvm-headers-no-generated-headers
    )
endif ()
#################### libclangBasic ####################

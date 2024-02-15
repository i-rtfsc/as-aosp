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

project(libclangFrontend)

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

#################### libclangFrontend ####################
set(FILES
        ${ANDROID_ROOT}/external/clang/lib/Frontend/ASTConsumers.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/ASTMerge.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/ASTUnit.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/CacheTokens.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/ChainedDiagnosticConsumer.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/ChainedIncludesSource.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/CodeGenOptions.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/CompilerInstance.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/CompilerInvocation.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/CreateInvocationFromCommandLine.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/DependencyFile.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/DependencyGraph.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/DiagnosticRenderer.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/FrontendAction.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/FrontendActions.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/FrontendOptions.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/HeaderIncludeGen.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/InitHeaderSearch.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/InitPreprocessor.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/LangStandards.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/LayoutOverrideSource.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/LogDiagnosticPrinter.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/ModuleDependencyCollector.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/MultiplexConsumer.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/PCHContainerOperations.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/PrintPreprocessedOutput.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/SerializedDiagnosticPrinter.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/SerializedDiagnosticReader.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/TestModuleFileExtension.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/TextDiagnostic.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/TextDiagnosticBuffer.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/TextDiagnosticPrinter.cpp
        ${ANDROID_ROOT}/external/clang/lib/Frontend/VerifyDiagnosticConsumer.cpp
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
    add_library(libclangFrontend SHARED ${SOURCE_FILES})

    target_include_directories(libclangFrontend
            PUBLIC
            ${ANDROID_ROOT}/external/clang/lib/Frontend
    )

    # Add c/c++ flags
    target_compile_options(libclangFrontend PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wno-unused-lambda-capture
            -Wno-unused-private-field
    )

    target_link_libraries(libclangFrontend
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
#################### libclangFrontend ####################

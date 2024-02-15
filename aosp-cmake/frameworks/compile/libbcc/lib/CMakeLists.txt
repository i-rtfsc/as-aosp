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

project(libbcc)

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

#################### libbcc ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/BCCContext.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/BCCContextImpl.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/Compiler.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/CompilerConfig.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/FileBase.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/Initialization.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSAddDebugInfoPass.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSCompilerDriver.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSEmbedInfo.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSGlobalInfoPass.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSInvariant.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSInvokeHelperPass.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSIsThreadablePass.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSKernelExpand.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSScreenFunctionsPass.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSScriptGroupFusion.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSFunctionsList.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSX86CallConvPass.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/RSX86TranslateGEPPass.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/Script.cpp
        ${ANDROID_ROOT}/frameworks/compile/libbcc/lib/Source.cpp
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
    add_library(libbcc SHARED ${SOURCE_FILES})

    target_include_directories(libbcc
            PUBLIC
            ${ANDROID_ROOT}/frameworks/compile/libbcc/lib
    )

    # Add c/c++ flags
    target_compile_options(libbcc PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libbcc
            libLLVMLinker
            libLLVM_android
            libbcinfo
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils
            libwinpthread
            llvm-headers
            llvm-headers-no-generated-headers
            slang_headers
    )
endif ()
#################### libbcc ####################

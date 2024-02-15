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

project(libclangCodeGen)

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

#################### libclangCodeGen ####################
set(FILES
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/BackendUtil.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGAtomic.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGBlocks.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGBuiltin.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGCUDABuiltin.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGCUDANV.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGCUDARuntime.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGCXX.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGCXXABI.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGCall.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGClass.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGCleanup.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGDebugInfo.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGDecl.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGDeclCXX.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGException.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGExpr.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGExprAgg.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGExprCXX.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGExprComplex.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGExprConstant.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGExprScalar.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGLoopInfo.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGObjC.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGObjCGNU.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGObjCMac.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGObjCRuntime.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGOpenCLRuntime.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGOpenMPRuntime.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGOpenMPRuntimeNVPTX.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGRecordLayoutBuilder.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGStmt.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGStmtOpenMP.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGVTT.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CGVTables.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CodeGenABITypes.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CodeGenAction.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CodeGenFunction.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CodeGenModule.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CodeGenPGO.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CodeGenTBAA.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CodeGenTypes.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/CoverageMappingGen.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/ItaniumCXXABI.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/MicrosoftCXXABI.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/ModuleBuilder.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/ObjectFilePCHContainerOperations.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/SanitizerMetadata.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/SwiftCallingConv.cpp
        ${ANDROID_ROOT}/external/clang/lib/CodeGen/TargetInfo.cpp
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
    add_library(libclangCodeGen SHARED ${SOURCE_FILES})

    target_include_directories(libclangCodeGen
            PUBLIC
            ${ANDROID_ROOT}/external/clang/lib/CodeGen
    )

    # Add c/c++ flags
    target_compile_options(libclangCodeGen PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wno-unused-lambda-capture
            -Wno-unused-private-field
    )

    target_link_libraries(libclangCodeGen
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
#################### libclangCodeGen ####################

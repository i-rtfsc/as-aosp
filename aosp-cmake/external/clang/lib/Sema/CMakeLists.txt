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

project(libclangSema)

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

#################### libclangSema ####################
set(FILES
        ${ANDROID_ROOT}/external/clang/lib/Sema/AnalysisBasedWarnings.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/AttributeList.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/CodeCompleteConsumer.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/DeclSpec.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/DelayedDiagnostic.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/IdentifierResolver.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/JumpDiagnostics.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/MultiplexExternalSemaSource.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/Scope.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/ScopeInfo.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/Sema.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaAccess.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaAttr.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaCUDA.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaCXXScopeSpec.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaCast.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaChecking.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaCodeComplete.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaConsumer.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaCoroutine.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaDecl.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaDeclAttr.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaDeclCXX.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaDeclObjC.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaExceptionSpec.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaExpr.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaExprCXX.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaExprMember.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaExprObjC.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaFixItUtils.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaInit.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaLambda.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaLookup.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaObjCProperty.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaOpenMP.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaOverload.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaPseudoObject.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaStmt.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaStmtAsm.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaStmtAttr.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaTemplate.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaTemplateDeduction.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaTemplateInstantiate.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaTemplateInstantiateDecl.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaTemplateVariadic.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/SemaType.cpp
        ${ANDROID_ROOT}/external/clang/lib/Sema/TypeLocBuilder.cpp
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
    add_library(libclangSema SHARED ${SOURCE_FILES})

    target_include_directories(libclangSema
            PUBLIC
            ${ANDROID_ROOT}/external/clang/lib/Sema
    )

    # Add c/c++ flags
    target_compile_options(libclangSema PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wno-unused-lambda-capture
            -Wno-unused-private-field
    )

    target_link_libraries(libclangSema
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
#################### libclangSema ####################

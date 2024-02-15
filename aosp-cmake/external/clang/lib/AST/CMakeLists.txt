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

project(libclangAST)

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

#################### libclangAST ####################
set(FILES
        ${ANDROID_ROOT}/external/clang/lib/AST/APValue.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ASTConsumer.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ASTContext.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ASTDiagnostic.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ASTDumper.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ASTImporter.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ASTTypeTraits.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/AttrImpl.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/CXXInheritance.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/Comment.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/CommentBriefParser.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/CommentCommandTraits.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/CommentLexer.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/CommentParser.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/CommentSema.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/Decl.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/DeclBase.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/DeclCXX.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/DeclFriend.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/DeclGroup.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/DeclObjC.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/DeclOpenMP.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/DeclPrinter.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/DeclTemplate.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/DeclarationName.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/Expr.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ExprCXX.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ExprClassification.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ExprConstant.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ExprObjC.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ExternalASTSource.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/InheritViz.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ItaniumCXXABI.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ItaniumMangle.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/Mangle.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/MicrosoftCXXABI.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/MicrosoftMangle.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/NSAPI.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/NestedNameSpecifier.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/OpenMPClause.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/ParentMap.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/RawCommentList.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/RecordLayout.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/RecordLayoutBuilder.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/SelectorLocationsKind.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/Stmt.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/StmtCXX.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/StmtIterator.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/StmtObjC.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/StmtOpenMP.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/StmtPrinter.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/StmtProfile.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/StmtViz.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/TemplateBase.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/TemplateName.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/Type.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/TypeLoc.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/TypePrinter.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/VTTBuilder.cpp
        ${ANDROID_ROOT}/external/clang/lib/AST/VTableBuilder.cpp
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
    add_library(libclangAST SHARED ${SOURCE_FILES})

    target_include_directories(libclangAST
            PUBLIC
            ${ANDROID_ROOT}/external/clang/lib/AST
    )

    # Add c/c++ flags
    target_compile_options(libclangAST PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wno-unused-lambda-capture
            -Wno-unused-private-field
    )

    target_link_libraries(libclangAST
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
#################### libclangAST ####################

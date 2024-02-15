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

project(libclangAnalysis)

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

#################### libclangAnalysis ####################
set(FILES
        ${ANDROID_ROOT}/external/clang/lib/Analysis/AnalysisDeclContext.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/BodyFarm.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/CFG.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/CFGReachabilityAnalysis.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/CFGStmtMap.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/CallGraph.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/CocoaConventions.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/CodeInjector.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/Consumed.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/Dominators.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/FormatString.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/LiveVariables.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/ObjCNoReturn.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/PostOrderCFGView.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/PrintfFormatString.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/ProgramPoint.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/PseudoConstantAnalysis.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/ReachableCode.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/ScanfFormatString.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/ThreadSafety.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/ThreadSafetyCommon.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/ThreadSafetyLogical.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/ThreadSafetyTIL.cpp
        ${ANDROID_ROOT}/external/clang/lib/Analysis/UninitializedValues.cpp
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
    add_library(libclangAnalysis SHARED ${SOURCE_FILES})

    target_include_directories(libclangAnalysis
            PUBLIC
            ${ANDROID_ROOT}/external/clang/lib/Analysis
    )

    # Add c/c++ flags
    target_compile_options(libclangAnalysis PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wno-unused-lambda-capture
            -Wno-unused-private-field
    )

    target_link_libraries(libclangAnalysis
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
#################### libclangAnalysis ####################

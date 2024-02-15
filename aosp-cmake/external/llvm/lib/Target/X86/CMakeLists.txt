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

project(libLLVMX86CodeGen)

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

#################### libLLVMX86CodeGen ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86AsmPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86CallFrameOptimization.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86ExpandPseudo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86FastISel.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86FixupBWInsts.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86FixupLEAs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86FixupSetCC.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86FloatingPoint.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86FrameLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86ISelDAGToDAG.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86ISelLowering.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86InstrInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86MCInstLower.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86MachineFunctionInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86OptimizeLEAs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86PadShortFunction.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86RegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86SelectionDAGInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86ShuffleDecodeConstantPool.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86Subtarget.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86TargetMachine.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86TargetObjectFile.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86TargetTransformInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86VZeroUpper.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86WinAllocaExpander.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Target/X86/X86WinEHState.cpp
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
    add_library(libLLVMX86CodeGen SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMX86CodeGen
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/Target/X86
            ${OUT_INTERMEDIATES_ROOT}/external/llvm/lib/Target/X86/llvm-gen-x86/gen
    )

    # Add c/c++ flags
    target_compile_options(libLLVMX86CodeGen PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMX86CodeGen
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            llvm-headers
            llvm-headers-no-generated-headers
            llvm-x86-headers
    )
endif ()
#################### libLLVMX86CodeGen ####################

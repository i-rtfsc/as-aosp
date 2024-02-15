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

project(libLLVMMC)

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

#################### libLLVMMC ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/MC/ConstantPools.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/ELFObjectWriter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCAsmBackend.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCAsmInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCAsmInfoCOFF.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCAsmInfoDarwin.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCAsmInfoELF.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCAsmStreamer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCAssembler.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCCodeEmitter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCCodeView.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCContext.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCDwarf.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCELFObjectTargetWriter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCELFStreamer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCExpr.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCFragment.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCInst.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCInstPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCInstrAnalysis.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCInstrDesc.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCLabel.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCLinkerOptimizationHint.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCMachOStreamer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCMachObjectTargetWriter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCNullStreamer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCObjectFileInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCObjectStreamer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCObjectWriter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCRegisterInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCSchedule.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCSection.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCSectionCOFF.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCSectionELF.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCSectionMachO.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCStreamer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCSubtargetInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCSymbol.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCSymbolELF.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCTargetOptions.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCValue.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCWin64EH.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MCWinEH.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/MachObjectWriter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/StringTableBuilder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/SubtargetFeature.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/WinCOFFObjectWriter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/MC/WinCOFFStreamer.cpp
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
    add_library(libLLVMMC SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMMC
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/MC
    )

    # Add c/c++ flags
    target_compile_options(libLLVMMC PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMMC
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
#################### libLLVMMC ####################

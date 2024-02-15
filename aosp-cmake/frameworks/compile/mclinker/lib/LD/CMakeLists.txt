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

project(libmcldLD)

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

#################### libmcldLD ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/Archive.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ArchiveReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/BranchIsland.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/BranchIslandFactory.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/BinaryReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/DWARFLineInfo.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/Diagnostic.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/DiagnosticEngine.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/DiagnosticInfos.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/DiagnosticLineInfo.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/DiagnosticPrinter.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/DebugString.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/DynObjReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFBinaryReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFSegment.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFSegmentFactory.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/EhFrame.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/EhFrameHdr.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/EhFrameReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/GarbageCollection.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/GroupReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/IdenticalCodeFolding.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/LDContext.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/LDFileFormat.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/LDReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/LDSection.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/LDSymbol.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/MergedStringTable.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/MsgHandler.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/NamePool.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ObjectWriter.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/RelocData.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/RelocationFactory.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/Relocator.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ResolveInfo.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/Resolver.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/SectionData.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/SectionSymbolSet.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/StaticResolver.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/StubFactory.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/TextDiagnosticPrinter.cpp
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
    add_library(libmcldLD SHARED ${SOURCE_FILES})

    target_include_directories(libmcldLD
            PUBLIC
            ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD
    )

    # Add c/c++ flags
    target_compile_options(libmcldLD PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libmcldLD
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
            llvm-headers
            llvm-headers-no-generated-headers
    )
endif ()
#################### libmcldLD ####################

#################### libmcldLDVariant ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/BSDArchiveReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/GNUArchiveReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFDynObjFileFormat.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFDynObjReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFExecFileFormat.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFFileFormat.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFObjectReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFObjectWriter.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFReader.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD/ELFReaderIf.cpp
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
    add_library(libmcldLDVariant SHARED ${SOURCE_FILES})

    target_include_directories(libmcldLDVariant
            PUBLIC
            ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/LD
    )

    # Add c/c++ flags
    target_compile_options(libmcldLDVariant PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libmcldLDVariant
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
#################### libmcldLDVariant ####################

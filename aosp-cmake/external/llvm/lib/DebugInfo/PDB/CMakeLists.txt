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

project(libLLVMDebugInfoPDB)

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

#################### libLLVMDebugInfoPDB ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/GenericError.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/IPDBSourceFile.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDB.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBContext.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBExtras.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBInterfaceAnchors.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymDumper.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbol.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolAnnotation.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolBlock.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolCompiland.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolCompilandDetails.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolCompilandEnv.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolCustom.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolData.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolExe.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolFunc.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolFuncDebugEnd.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolFuncDebugStart.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolLabel.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolPublicSymbol.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolThunk.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeArray.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeBaseClass.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeBuiltin.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeCustom.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeDimension.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeEnum.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeFriend.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeFunctionArg.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeFunctionSig.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeManaged.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypePointer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeTypedef.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeUDT.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeVTable.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolTypeVTableShape.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolUnknown.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/PDBSymbolUsingNamespace.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/DbiStream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/DbiStreamBuilder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/EnumTables.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/Hash.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/IndexedStreamData.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/InfoStream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/InfoStreamBuilder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/MappedBlockStream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/ModInfo.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/ModStream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/NameHashTable.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/NameMap.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/PDBFile.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/PDBFileBuilder.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/PublicsStream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/RawError.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/RawSession.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/SymbolStream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB/Raw/TpiStream.cpp
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
    add_library(libLLVMDebugInfoPDB SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMDebugInfoPDB
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/DebugInfo/PDB
    )

    # Add c/c++ flags
    target_compile_options(libLLVMDebugInfoPDB PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
            -Wall
            -Wno-unused-lambda-capture
            -Wno-unused-parameter
    )

    target_link_libraries(libLLVMDebugInfoPDB
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
#################### libLLVMDebugInfoPDB ####################

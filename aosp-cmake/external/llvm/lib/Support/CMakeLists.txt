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

project(libLLVMSupport)

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

#################### libLLVMSupport ####################
set(FILES
        ${ANDROID_ROOT}/external/llvm/lib/Support/ConvertUTF.c
        ${ANDROID_ROOT}/external/llvm/lib/Support/regcomp.c
        ${ANDROID_ROOT}/external/llvm/lib/Support/regerror.c
        ${ANDROID_ROOT}/external/llvm/lib/Support/regexec.c
        ${ANDROID_ROOT}/external/llvm/lib/Support/regfree.c
        ${ANDROID_ROOT}/external/llvm/lib/Support/regstrlcpy.c
        ${ANDROID_ROOT}/external/llvm/lib/Support/APFloat.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/APInt.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/APSInt.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ARMBuildAttrs.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ARMWinEH.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Allocator.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Atomic.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/BlockFrequency.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/BranchProbability.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/COM.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/CachePruning.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/CommandLine.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Compression.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ConvertUTFWrapper.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/CrashRecoveryContext.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/DAGDeltaAlgorithm.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/DataExtractor.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/DataStream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Debug.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/DeltaAlgorithm.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Dwarf.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/DynamicLibrary.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Errno.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Error.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ErrorHandling.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/FileOutputBuffer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/FileUtilities.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/FoldingSet.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/FormattedStream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/GraphWriter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Hashing.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Host.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/IntEqClasses.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/IntervalMap.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/IntrusiveRefCntPtr.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/JamCRC.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/LEB128.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/LineIterator.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Locale.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/LockFileManager.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/MD5.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ManagedStatic.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/MathExtras.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Memory.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/MemoryBuffer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/MemoryObject.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Mutex.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Options.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Path.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/PluginLoader.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/PrettyStackTrace.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Process.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Program.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/RWMutex.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/RandomNumberGenerator.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Regex.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/SHA1.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ScaledNumber.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ScopedPrinter.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/SearchForAddressOfSpecialSymbol.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Signals.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/SmallPtrSet.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/SmallVector.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/SourceMgr.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/SpecialCaseList.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Statistic.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/StreamingMemoryObject.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/StringExtras.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/StringMap.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/StringPool.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/StringRef.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/StringSaver.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/SystemUtils.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/TargetParser.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/TargetRegistry.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ThreadLocal.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ThreadPool.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Threading.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/TimeValue.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Timer.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/ToolOutputFile.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Triple.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Twine.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Unicode.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Valgrind.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/Watchdog.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/YAMLParser.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/YAMLTraits.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/circular_raw_ostream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/raw_os_ostream.cpp
        ${ANDROID_ROOT}/external/llvm/lib/Support/raw_ostream.cpp
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
    add_library(libLLVMSupport SHARED ${SOURCE_FILES})

    target_include_directories(libLLVMSupport
            PUBLIC
            ${ANDROID_ROOT}/external/llvm/lib/Support
    )

    # Add c/c++ flags
    target_compile_options(libLLVMSupport PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libLLVMSupport
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            llvm-headers-no-generated-headers
    )
endif ()
#################### libLLVMSupport ####################

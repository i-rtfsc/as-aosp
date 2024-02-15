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

project(unwind_symbols)

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

#################### libunwindstack_stdout_log ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/AndroidUnwinder.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ArmExidx.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Demangle.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DexFiles.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfCfa.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfEhFrameWithHdr.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfMemory.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfOp.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfSection.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Elf.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ElfInterface.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ElfInterfaceArm.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Global.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/JitDebug.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/MapInfo.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Maps.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Memory.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/MemoryMte.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/MemoryXz.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Regs.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsArm.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsArm64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsX86.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsX86_64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsRiscv64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsMips.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsMips64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Symbols.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ThreadEntry.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ThreadUnwinder.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Unwinder.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/LogStdout.cpp
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
    add_library(libunwindstack_stdout_log SHARED ${SOURCE_FILES})

    target_include_directories(libunwindstack_stdout_log
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
    )

    # Add c/c++ flags
    target_compile_options(libunwindstack_stdout_log PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libunwindstack_stdout_log
            bionic_libc_platform_headers
            libasync_safe
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            liblzma
            libm
            libprocinfo
            librustc_demangle_static
    )
endif ()
#################### libunwindstack_stdout_log ####################

#################### libunwindstack_local ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/tests/TestLocal.cpp
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
    add_library(libunwindstack_local SHARED ${SOURCE_FILES})

    target_include_directories(libunwindstack_local
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
    )

    # Add c/c++ flags
    target_compile_options(libunwindstack_local PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libunwindstack_local
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libunwindstack
    )
endif ()
#################### libunwindstack_local ####################

#################### unwind ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/tools/unwind.cpp
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
    add_executable(unwind ${SOURCE_FILES})

    target_include_directories(unwind
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
    )

    # Add c/c++ flags
    target_compile_options(unwind PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(unwind
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblzma
            libm
            libunwindstack_stdout_log
    )
endif ()
#################### unwind ####################

#################### libunwindstack_fuzz_unwinder ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/tests/ElfFake.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/tests/fuzz/UnwinderComponentCreator.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/tests/fuzz/UnwinderFuzz.cpp
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
    add_executable(libunwindstack_fuzz_unwinder ${SOURCE_FILES})

    target_include_directories(libunwindstack_fuzz_unwinder
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/utils
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
    )

    # Add c/c++ flags
    target_compile_options(libunwindstack_fuzz_unwinder PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libunwindstack_fuzz_unwinder
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdexfile_support
            libdl
            liblog
            liblzma
            libm
            libunwindstack
            libunwindstack_utils
    )
endif ()
#################### libunwindstack_fuzz_unwinder ####################

#################### unwind_reg_info ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/tools/unwind_reg_info.cpp
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
    add_executable(unwind_reg_info ${SOURCE_FILES})

    target_include_directories(unwind_reg_info
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
    )

    # Add c/c++ flags
    target_compile_options(unwind_reg_info PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(unwind_reg_info
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblzma
            libm
            libunwindstack_stdout_log
    )
endif ()
#################### unwind_reg_info ####################

#################### libunwindstack ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/AndroidUnwinder.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ArmExidx.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Demangle.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DexFiles.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfCfa.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfEhFrameWithHdr.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfMemory.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfOp.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfSection.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Elf.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ElfInterface.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ElfInterfaceArm.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Global.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/JitDebug.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/MapInfo.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Maps.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Memory.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/MemoryMte.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/MemoryXz.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Regs.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsArm.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsArm64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsX86.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsX86_64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsRiscv64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsMips.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsMips64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Symbols.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ThreadEntry.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ThreadUnwinder.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Unwinder.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DexFile.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/LogAndroid.cpp
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
    add_library(libunwindstack SHARED ${SOURCE_FILES})

    target_include_directories(libunwindstack
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
    )

    # Add c/c++ flags
    target_compile_options(libunwindstack PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libunwindstack
            bionic_libc_platform_headers
            libasync_safe
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfile_support
            libdl
            liblog
            liblzma
            libm
            libprocinfo
            librustc_demangle_static
    )
endif ()
#################### libunwindstack ####################

#################### unwind_info ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/tools/unwind_info.cpp
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
    add_executable(unwind_info ${SOURCE_FILES})

    target_include_directories(unwind_info
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
    )

    # Add c/c++ flags
    target_compile_options(unwind_info PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(unwind_info
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblzma
            libm
            libunwindstack_stdout_log
    )
endif ()
#################### unwind_info ####################

#################### libunwindstack_utils ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/utils/MemoryFake.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/utils/OfflineUnwindUtils.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/utils/PidUtils.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/utils/ProcessTracer.cpp
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
    add_library(libunwindstack_utils SHARED ${SOURCE_FILES})

    target_include_directories(libunwindstack_utils
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/utils
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
    )

    # Add c/c++ flags
    target_compile_options(libunwindstack_utils PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libunwindstack_utils
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libdl
            libm
            libprocinfo
            libunwindstack
            libz
    )
endif ()
#################### libunwindstack_utils ####################

#################### libunwindstack_no_dex ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/AndroidUnwinder.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ArmExidx.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Demangle.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DexFiles.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfCfa.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfEhFrameWithHdr.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfMemory.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfOp.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/DwarfSection.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Elf.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ElfInterface.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ElfInterfaceArm.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Global.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/JitDebug.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/MapInfo.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Maps.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Memory.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/MemoryMte.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/MemoryXz.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Regs.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsArm.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsArm64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsX86.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsX86_64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsRiscv64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsMips.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/RegsMips64.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Symbols.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ThreadEntry.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/ThreadUnwinder.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/Unwinder.cpp
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/LogAndroid.cpp
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
    add_library(libunwindstack_no_dex SHARED ${SOURCE_FILES})

    target_include_directories(libunwindstack_no_dex
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
    )

    # Add c/c++ flags
    target_compile_options(libunwindstack_no_dex PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libunwindstack_no_dex
            bionic_libc_platform_headers
            libasync_safe
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            liblzma
            libm
            libprocinfo
            librustc_demangle_static
    )
endif ()
#################### libunwindstack_no_dex ####################

#################### unwind_for_offline ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/tools/unwind_for_offline.cpp
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
    add_executable(unwind_for_offline ${SOURCE_FILES})

    target_include_directories(unwind_for_offline
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/utils
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
    )

    # Add c/c++ flags
    target_compile_options(unwind_for_offline PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(unwind_for_offline
            libbase
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libdl
            liblzma
            libm
            libunwindstack_stdout_log
            libunwindstack_utils
    )
endif ()
#################### unwind_for_offline ####################

#################### unwind_symbols ####################
set(FILES
        ${ANDROID_ROOT}/system/unwinding/libunwindstack/tools/unwind_symbols.cpp
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
    add_executable(unwind_symbols ${SOURCE_FILES})

    target_include_directories(unwind_symbols
            PUBLIC
            ${ANDROID_ROOT}/system/unwinding/libunwindstack
            ${ANDROID_ROOT}/system/unwinding/libunwindstack/include
    )

    # Add c/c++ flags
    target_compile_options(unwind_symbols PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(unwind_symbols
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblzma
            libm
            libunwindstack_stdout_log
    )
endif ()
#################### unwind_symbols ####################

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

project(libutils_fuzz_string8)

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

#################### libutils_fuzz_refbase ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/RefBase_fuzz.cpp
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
    add_executable(libutils_fuzz_refbase ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_refbase
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_refbase PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_refbase
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libutils_fuzz_refbase ####################

#################### libutilscallstack ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/CallStack.cpp
        ${ANDROID_ROOT}/system/core/libutils/ProcessCallStack.cpp
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
    add_library(libutilscallstack SHARED ${SOURCE_FILES})

    target_include_directories(libutilscallstack
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutilscallstack PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libutilscallstack
            libbase_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libunwindstack
            libutils
            libutils_headers
            libvndksupport
    )
endif ()
#################### libutilscallstack ####################

#################### libutils_fuzz_printer ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/Printer_fuzz.cpp
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
    add_executable(libutils_fuzz_printer ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_printer
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_printer PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_printer
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libutils_fuzz_printer ####################

#################### libutils ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/Errors.cpp
        ${ANDROID_ROOT}/system/core/libutils/FileMap.cpp
        ${ANDROID_ROOT}/system/core/libutils/JenkinsHash.cpp
        ${ANDROID_ROOT}/system/core/libutils/LightRefBase.cpp
        ${ANDROID_ROOT}/system/core/libutils/NativeHandle.cpp
        ${ANDROID_ROOT}/system/core/libutils/Printer.cpp
        ${ANDROID_ROOT}/system/core/libutils/RefBase.cpp
        ${ANDROID_ROOT}/system/core/libutils/SharedBuffer.cpp
        ${ANDROID_ROOT}/system/core/libutils/StopWatch.cpp
        ${ANDROID_ROOT}/system/core/libutils/String8.cpp
        ${ANDROID_ROOT}/system/core/libutils/String16.cpp
        ${ANDROID_ROOT}/system/core/libutils/StrongPointer.cpp
        ${ANDROID_ROOT}/system/core/libutils/SystemClock.cpp
        ${ANDROID_ROOT}/system/core/libutils/Threads.cpp
        ${ANDROID_ROOT}/system/core/libutils/Timers.cpp
        ${ANDROID_ROOT}/system/core/libutils/Tokenizer.cpp
        ${ANDROID_ROOT}/system/core/libutils/Unicode.cpp
        ${ANDROID_ROOT}/system/core/libutils/VectorImpl.cpp
        ${ANDROID_ROOT}/system/core/libutils/misc.cpp
        ${ANDROID_ROOT}/system/core/libutils/Looper.cpp
        ${ANDROID_ROOT}/system/core/libutils/Trace.cpp
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
    add_library(libutils SHARED ${SOURCE_FILES})

    target_include_directories(libutils
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libutils
            libbase_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils_headers
            libvndksupport
            libwinpthread
    )
endif ()
#################### libutils ####################

#################### libutils_fuzz_string16 ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/String16_fuzz.cpp
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
    add_executable(libutils_fuzz_string16 ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_string16
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_string16 PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_string16
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libutils_fuzz_string16 ####################

#################### libutils_fuzz_vector ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/Vector_fuzz.cpp
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
    add_executable(libutils_fuzz_vector ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_vector
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_vector PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_vector
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libutils_fuzz_vector ####################

#################### libutils_fuzz_lrucache ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/LruCache_fuzz.cpp
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
    add_executable(libutils_fuzz_lrucache ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_lrucache
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_lrucache PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_lrucache
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libutils_fuzz_lrucache ####################

#################### libutils_fuzz_process_callstack ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/ProcessCallStack_fuzz.cpp
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
    add_executable(libutils_fuzz_process_callstack ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_process_callstack
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_process_callstack PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_process_callstack
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
            libutilscallstack
    )
endif ()
#################### libutils_fuzz_process_callstack ####################

#################### libutils_fuzz_callstack ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/CallStack_fuzz.cpp
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
    add_executable(libutils_fuzz_callstack ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_callstack
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_callstack PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_callstack
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
            libutilscallstack
    )
endif ()
#################### libutils_fuzz_callstack ####################

#################### libutils_fuzz_filemap ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/FileMap_fuzz.cpp
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
    add_executable(libutils_fuzz_filemap ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_filemap
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_filemap PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_filemap
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libutils_fuzz_filemap ####################

#################### libutils_fuzz_bitset ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/BitSet_fuzz.cpp
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
    add_executable(libutils_fuzz_bitset ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_bitset
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_bitset PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_bitset
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libutils_fuzz_bitset ####################

#################### libutils_fuzz_looper ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/Looper_fuzz.cpp
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
    add_executable(libutils_fuzz_looper ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_looper
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_looper PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_looper
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libutils_fuzz_looper ####################

#################### libutils_fuzz_string8 ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libutils/String8_fuzz.cpp
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
    add_executable(libutils_fuzz_string8 ${SOURCE_FILES})

    target_include_directories(libutils_fuzz_string8
            PUBLIC
            ${ANDROID_ROOT}/system/core/libutils
            ${ANDROID_ROOT}/system/core/libutils/include
    )

    # Add c/c++ flags
    target_compile_options(libutils_fuzz_string8 PRIVATE
            -Werror
    )

    target_link_libraries(libutils_fuzz_string8
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libutils_fuzz_string8 ####################

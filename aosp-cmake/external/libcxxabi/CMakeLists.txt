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

project(cxa_demangle_fuzzer)

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

#################### libc++demangle ####################
set(FILES
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_demangle.cpp
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
    add_library(libc++demangle SHARED ${SOURCE_FILES})

    target_include_directories(libc++demangle
            PUBLIC
            ${ANDROID_ROOT}/external/libcxxabi/include
            ${ANDROID_ROOT}/external/libcxxabi
    )

    # Add c/c++ flags
    target_compile_options(libc++demangle PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libc++demangle ####################

#################### cxa_demangle_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libcxxabi/fuzz/cxa_demangle_fuzzer.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_demangle.cpp
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
    add_executable(cxa_demangle_fuzzer ${SOURCE_FILES})

    target_include_directories(cxa_demangle_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libcxxabi/include
            ${ANDROID_ROOT}/external/libcxxabi
    )

    # Add c/c++ flags
    target_compile_options(cxa_demangle_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(cxa_demangle_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
    )
endif ()
#################### cxa_demangle_fuzzer ####################

#################### libc++abi ####################
set(FILES
        ${ANDROID_ROOT}/external/libcxxabi/src/abort_message.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_aux_runtime.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_default_handlers.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_exception.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_exception_storage.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_guard.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_handlers.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_personality.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_thread_atexit.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_unexpected.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_vector.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/cxa_virtual.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/fallback_malloc.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/private_typeinfo.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/stdlib_exception.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/stdlib_new_delete.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/stdlib_stdexcept.cpp
        ${ANDROID_ROOT}/external/libcxxabi/src/stdlib_typeinfo.cpp
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
    add_library(libc++abi SHARED ${SOURCE_FILES})

    target_include_directories(libc++abi
            PUBLIC
            ${ANDROID_ROOT}/external/libcxxabi/include
            ${ANDROID_ROOT}/external/libcxxabi
    )

    # Add c/c++ flags
    target_compile_options(libc++abi PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libc++abi
            libc
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libc++abi ####################

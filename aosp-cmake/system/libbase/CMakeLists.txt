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

project(libbase_parsenetaddress_fuzzer)

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

#################### libbase_ndk ####################
set(FILES
        ${ANDROID_ROOT}/system/libbase/abi_compatibility.cpp
        ${ANDROID_ROOT}/system/libbase/chrono_utils.cpp
        ${ANDROID_ROOT}/system/libbase/cmsg.cpp
        ${ANDROID_ROOT}/system/libbase/file.cpp
        ${ANDROID_ROOT}/system/libbase/hex.cpp
        ${ANDROID_ROOT}/system/libbase/logging.cpp
        ${ANDROID_ROOT}/system/libbase/mapped_file.cpp
        ${ANDROID_ROOT}/system/libbase/parsebool.cpp
        ${ANDROID_ROOT}/system/libbase/parsenetaddress.cpp
        ${ANDROID_ROOT}/system/libbase/posix_strerror_r.cpp
        ${ANDROID_ROOT}/system/libbase/process.cpp
        ${ANDROID_ROOT}/system/libbase/properties.cpp
        ${ANDROID_ROOT}/system/libbase/stringprintf.cpp
        ${ANDROID_ROOT}/system/libbase/strings.cpp
        ${ANDROID_ROOT}/system/libbase/threads.cpp
        ${ANDROID_ROOT}/system/libbase/test_utils.cpp
        ${ANDROID_ROOT}/system/libbase/errors_unix.cpp
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
    add_library(libbase_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libbase_ndk
            PUBLIC
            ${ANDROID_ROOT}/system/libbase/include
            ${ANDROID_ROOT}/system/libbase
    )

    # Add c/c++ flags
    target_compile_options(libbase_ndk PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libbase_ndk
            fmtlib_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
    )
endif ()
#################### libbase_ndk ####################

#################### libbase ####################
set(FILES
        ${ANDROID_ROOT}/system/libbase/abi_compatibility.cpp
        ${ANDROID_ROOT}/system/libbase/chrono_utils.cpp
        ${ANDROID_ROOT}/system/libbase/cmsg.cpp
        ${ANDROID_ROOT}/system/libbase/file.cpp
        ${ANDROID_ROOT}/system/libbase/hex.cpp
        ${ANDROID_ROOT}/system/libbase/logging.cpp
        ${ANDROID_ROOT}/system/libbase/mapped_file.cpp
        ${ANDROID_ROOT}/system/libbase/parsebool.cpp
        ${ANDROID_ROOT}/system/libbase/parsenetaddress.cpp
        ${ANDROID_ROOT}/system/libbase/posix_strerror_r.cpp
        ${ANDROID_ROOT}/system/libbase/process.cpp
        ${ANDROID_ROOT}/system/libbase/properties.cpp
        ${ANDROID_ROOT}/system/libbase/stringprintf.cpp
        ${ANDROID_ROOT}/system/libbase/strings.cpp
        ${ANDROID_ROOT}/system/libbase/threads.cpp
        ${ANDROID_ROOT}/system/libbase/test_utils.cpp
        ${ANDROID_ROOT}/system/libbase/errors_unix.cpp
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
    add_library(libbase SHARED ${SOURCE_FILES})

    target_include_directories(libbase
            PUBLIC
            ${ANDROID_ROOT}/system/libbase
            ${ANDROID_ROOT}/system/libbase/include
    )

    # Add c/c++ flags
    target_compile_options(libbase PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libbase
            fmtlib
            libbase_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libwinpthread
    )
endif ()
#################### libbase ####################

#################### libbase_parsenetaddress_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/libbase/parsenetaddress_fuzzer.cpp
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
    add_executable(libbase_parsenetaddress_fuzzer ${SOURCE_FILES})

    target_include_directories(libbase_parsenetaddress_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/libbase
            ${ANDROID_ROOT}/system/libbase/include
    )

    # Add c/c++ flags
    target_compile_options(libbase_parsenetaddress_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libbase_parsenetaddress_fuzzer
            libbase
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
#################### libbase_parsenetaddress_fuzzer ####################

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

project(libpdx_service_dispatcher_fuzzer)

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

#################### libpdx_message_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/vr/libpdx/fuzz/message_fuzzer.cpp
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
    add_executable(libpdx_message_fuzzer ${SOURCE_FILES})

    target_include_directories(libpdx_message_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/vr/libpdx/fuzz
    )

    # Add c/c++ flags
    target_compile_options(libpdx_message_fuzzer PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libpdx_message_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libpdx
            libutils
    )
endif ()
#################### libpdx_message_fuzzer ####################

#################### libpdx_serialization_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/vr/libpdx/fuzz/serialization_fuzzer.cpp
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
    add_executable(libpdx_serialization_fuzzer ${SOURCE_FILES})

    target_include_directories(libpdx_serialization_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/vr/libpdx/fuzz
    )

    # Add c/c++ flags
    target_compile_options(libpdx_serialization_fuzzer PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libpdx_serialization_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libpdx
            libutils
    )
endif ()
#################### libpdx_serialization_fuzzer ####################

#################### libpdx_service_dispatcher_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/vr/libpdx/fuzz/service_dispatcher_fuzzer.cpp
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
    add_executable(libpdx_service_dispatcher_fuzzer ${SOURCE_FILES})

    target_include_directories(libpdx_service_dispatcher_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/vr/libpdx/fuzz
    )

    # Add c/c++ flags
    target_compile_options(libpdx_service_dispatcher_fuzzer PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libpdx_service_dispatcher_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libpdx
            libutils
    )
endif ()
#################### libpdx_service_dispatcher_fuzzer ####################

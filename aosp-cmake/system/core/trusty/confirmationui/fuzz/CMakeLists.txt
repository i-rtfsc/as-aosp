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

project(trusty_confirmationui_msg_fuzzer)

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

#################### trusty_confirmationui_tipc_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/fuzz/tipc_fuzzer.cpp
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
    add_executable(trusty_confirmationui_tipc_fuzzer ${SOURCE_FILES})

    target_include_directories(trusty_confirmationui_tipc_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/confirmationui/fuzz
    )

    # Add c/c++ flags
    target_compile_options(trusty_confirmationui_tipc_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(trusty_confirmationui_tipc_fuzzer
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
            libtrusty
            libtrusty_coverage
            libtrusty_fuzz_utils
    )
endif ()
#################### trusty_confirmationui_tipc_fuzzer ####################

#################### trusty_confirmationui_msg_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/confirmationui/fuzz/msg_fuzzer.cpp
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
    add_executable(trusty_confirmationui_msg_fuzzer ${SOURCE_FILES})

    target_include_directories(trusty_confirmationui_msg_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/confirmationui/fuzz
    )

    # Add c/c++ flags
    target_compile_options(trusty_confirmationui_msg_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(trusty_confirmationui_msg_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libdmabufheap
            liblog
            libm
            libtrusty
            libtrusty_coverage
            libtrusty_fuzz_utils
    )
endif ()
#################### trusty_confirmationui_msg_fuzzer ####################

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

project(libchainagents)

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

#################### libchainagents ####################
set(FILES
        ${ANDROID_ROOT}/art/tools/jvmti-agents/chain-agents/chainagents.cc
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
    add_library(libchainagents SHARED ${SOURCE_FILES})

    target_include_directories(libchainagents
            PUBLIC
            ${ANDROID_ROOT}/art/tools/jvmti-agents/chain-agents
    )

    # Add c/c++ flags
    target_compile_options(libchainagents PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libchainagents
            jni_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnativehelper_header_only
            libopenjdkjvmti_headers
    )
endif ()
#################### libchainagents ####################

#################### libchainagentsd ####################
set(FILES
        ${ANDROID_ROOT}/art/tools/jvmti-agents/chain-agents/chainagents.cc
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
    add_library(libchainagentsd SHARED ${SOURCE_FILES})

    target_include_directories(libchainagentsd
            PUBLIC
            ${ANDROID_ROOT}/art/tools/jvmti-agents/chain-agents
    )

    # Add c/c++ flags
    target_compile_options(libchainagentsd PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(libchainagentsd
            jni_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnativehelper_header_only
            libopenjdkjvmti_headers
    )
endif ()
#################### libchainagentsd ####################

#################### libchainagentss ####################
set(FILES
        ${ANDROID_ROOT}/art/tools/jvmti-agents/chain-agents/chainagents.cc
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
    add_library(libchainagentss SHARED ${SOURCE_FILES})

    target_include_directories(libchainagentss
            PUBLIC
            ${ANDROID_ROOT}/art/tools/jvmti-agents/chain-agents
    )

    # Add c/c++ flags
    target_compile_options(libchainagentss PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libchainagentss
            jni_headers
            libbase_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libnativehelper_header_only
            libopenjdkjvmti_headers
    )
endif ()
#################### libchainagentss ####################

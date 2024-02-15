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

project(libsystemproperties)

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

#################### libsystemproperties ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libc/system_properties/context_node.cpp
        ${ANDROID_ROOT}/bionic/libc/system_properties/contexts_split.cpp
        ${ANDROID_ROOT}/bionic/libc/system_properties/contexts_serialized.cpp
        ${ANDROID_ROOT}/bionic/libc/system_properties/prop_area.cpp
        ${ANDROID_ROOT}/bionic/libc/system_properties/prop_info.cpp
        ${ANDROID_ROOT}/bionic/libc/system_properties/system_properties.cpp
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
    add_library(libsystemproperties SHARED ${SOURCE_FILES})

    target_include_directories(libsystemproperties
            PUBLIC
            ${ANDROID_ROOT}/bionic/libc/system_properties/include
            ${ANDROID_ROOT}/bionic/libc/system_properties
    )

    # Add c/c++ flags
    target_compile_options(libsystemproperties PRIVATE
            -Wall
            -Wextra
            -Werror=pointer-to-int-cast
            -Werror=int-to-pointer-cast
            -Werror=type-limits
            -Werror
    )

    target_link_libraries(libsystemproperties
            libasync_safe_headers
            libc_headers
            libclang_rt.builtins
            liblog_headers
            libpropertyinfoparser
    )
endif ()
#################### libsystemproperties ####################

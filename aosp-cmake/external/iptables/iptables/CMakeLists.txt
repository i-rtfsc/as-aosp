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

project(iptables)

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

#################### iptables ####################
set(FILES
        ${ANDROID_ROOT}/external/iptables/iptables/xtables-legacy-multi.c
        ${ANDROID_ROOT}/external/iptables/iptables/iptables-xml.c
        ${ANDROID_ROOT}/external/iptables/iptables/xshared.c
        ${ANDROID_ROOT}/external/iptables/iptables/iptables-save.c
        ${ANDROID_ROOT}/external/iptables/iptables/iptables-restore.c
        ${ANDROID_ROOT}/external/iptables/iptables/iptables-standalone.c
        ${ANDROID_ROOT}/external/iptables/iptables/iptables.c
        ${ANDROID_ROOT}/external/iptables/iptables/ip6tables-standalone.c
        ${ANDROID_ROOT}/external/iptables/iptables/ip6tables.c
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
    add_executable(iptables ${SOURCE_FILES})

    target_include_directories(iptables
            PUBLIC
            ${ANDROID_ROOT}/external/iptables/iptables
    )

    # Add c/c++ flags
    target_compile_options(iptables PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(iptables
            iptables_config_header
            iptables_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libext
            libext4
            libext6
            libip4tc
            libip6tc
            libm
            libxtables
    )
endif ()
#################### iptables ####################

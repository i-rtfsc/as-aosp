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

project(libmcldMC)

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

#################### libmcldMC ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/Attribute.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/AttributeSet.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/CommandAction.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/ContextFactory.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/FileAction.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/InputAction.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/InputBuilder.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/InputFactory.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/Input.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/MCLDDirectory.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/SearchDirs.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/SymbolCategory.cpp
        ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC/ZOption.cpp
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
    add_library(libmcldMC SHARED ${SOURCE_FILES})

    target_include_directories(libmcldMC
            PUBLIC
            ${ANDROID_ROOT}/frameworks/compile/mclinker/lib/MC
    )

    # Add c/c++ flags
    target_compile_options(libmcldMC PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(libmcldMC
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            llvm-headers
            llvm-headers-no-generated-headers
    )
endif ()
#################### libmcldMC ####################

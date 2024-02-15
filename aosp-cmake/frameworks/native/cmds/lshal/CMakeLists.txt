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

project(lshal)

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

#################### liblshal ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/DebugCommand.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/HelpCommand.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/Lshal.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/ListCommand.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/PipeRelay.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/TableEntry.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/TextTable.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/utils.cpp
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/WaitCommand.cpp
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
    add_library(liblshal SHARED ${SOURCE_FILES})

    target_include_directories(liblshal
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/lshal
            ${ANDROID_ROOT}/frameworks/native/cmds/lshal/libprocpartition/include
    )

    # Add c/c++ flags
    target_compile_options(liblshal PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(liblshal
            libbase
            libbinderdebug
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidl-gen-hash
            libhidl-gen-utils
            libhidlbase
            libm
            libprocpartition
            libutils
            libvintf
    )
endif ()
#################### liblshal ####################

#################### lshal ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/cmds/lshal/main.cpp
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
    add_executable(lshal ${SOURCE_FILES})

    target_include_directories(lshal
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/cmds/lshal
            ${ANDROID_ROOT}/frameworks/native/cmds/lshal/libprocpartition/include
    )

    # Add c/c++ flags
    target_compile_options(lshal PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(lshal
            libbase
            libbinderdebug
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidl-gen-hash
            libhidl-gen-utils
            libhidlbase
            liblshal
            libm
            libprocpartition
            libutils
            libvintf
    )
endif ()
#################### lshal ####################

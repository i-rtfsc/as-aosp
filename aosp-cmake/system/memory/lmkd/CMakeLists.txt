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

project(lmkd)

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

#################### libstatslogc ####################
set(FILES
        ${ANDROID_ROOT}/system/memory/lmkd/statslog.cpp
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
    add_library(libstatslogc SHARED ${SOURCE_FILES})

    target_include_directories(libstatslogc
            PUBLIC
            ${ANDROID_ROOT}/system/memory/lmkd/include
            ${ANDROID_ROOT}/system/memory/lmkd
    )

    # Add c/c++ flags
    target_compile_options(libstatslogc PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libstatslogc
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libprocessgroup
    )
endif ()
#################### libstatslogc ####################

#################### liblmkd_utils ####################
set(FILES
        ${ANDROID_ROOT}/system/memory/lmkd/liblmkd_utils.cpp
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
    add_library(liblmkd_utils SHARED ${SOURCE_FILES})

    target_include_directories(liblmkd_utils
            PUBLIC
            ${ANDROID_ROOT}/system/memory/lmkd/include
            ${ANDROID_ROOT}/system/memory/lmkd
    )

    target_link_libraries(liblmkd_utils
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libprocessgroup
    )
endif ()
#################### liblmkd_utils ####################

#################### lmkd ####################
set(FILES
        ${ANDROID_ROOT}/system/memory/lmkd/lmkd.cpp
        ${ANDROID_ROOT}/system/memory/lmkd/reaper.cpp
        ${ANDROID_ROOT}/system/memory/lmkd/watchdog.cpp
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
    add_executable(lmkd ${SOURCE_FILES})

    target_include_directories(lmkd
            PUBLIC
            ${ANDROID_ROOT}/system/memory/lmkd/include
            ${ANDROID_ROOT}/system/memory/lmkd
            ${ANDROID_ROOT}/system/memory/lmkd/libpsi/include
    )

    # Add c/c++ flags
    target_compile_options(lmkd PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(lmkd
            bpf_syscall_wrappers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblmkd_utils
            liblog
            libm
            libprocessgroup
            libpsi
            libstatslogc
    )
endif ()
#################### lmkd ####################

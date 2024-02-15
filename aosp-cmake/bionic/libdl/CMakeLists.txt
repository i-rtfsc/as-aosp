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

project(libdl_android)

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

#################### libdl_android ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libdl/libdl_android.cpp
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
    add_library(libdl_android SHARED ${SOURCE_FILES})

    target_include_directories(libdl_android
            PUBLIC
            ${ANDROID_ROOT}/bionic/libdl
    )

    # Add c/c++ flags
    target_compile_options(libdl_android PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libdl_android
            ld-android
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libdl_android ####################

#################### libdl_static ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libdl/libdl.cpp
        ${ANDROID_ROOT}/bionic/libdl/libdl_cfi.cpp
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
    add_library(libdl_static SHARED ${SOURCE_FILES})

    target_include_directories(libdl_static
            PUBLIC
            ${ANDROID_ROOT}/bionic/libdl
    )

    # Add c/c++ flags
    target_compile_options(libdl_static PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libdl_static
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libdl_static ####################

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

project(libcuttlefish-ril-2)

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

#################### libcuttlefish-ril-2 ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-ril/atchannel.c
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-ril/at_tok.c
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-ril/base64util.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-ril/misc.c
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-ril/reference-ril.c
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
    add_library(libcuttlefish-ril-2 SHARED ${SOURCE_FILES})

    target_include_directories(libcuttlefish-ril-2
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/ril/reference-ril
    )

    # Add c/c++ flags
    target_compile_options(libcuttlefish-ril-2 PRIVATE
            -Wall
            -Wextra
            -Wno-unused-variable
            -Wno-unused-function
            -Werror
    )

    target_link_libraries(libcuttlefish-ril-2
            libbase
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libcuttlefish_fs
            libcuttlefish_utils
            libdl
            liblog
            libm
            libril-modem-lib
            librilutils
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libcuttlefish-ril-2 ####################

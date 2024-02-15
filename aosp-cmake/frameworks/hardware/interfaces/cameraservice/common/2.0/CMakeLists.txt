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

project(android.frameworks.cameraservice.common@2.0)

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

#################### android.frameworks.cameraservice.common-2.0 ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/hardware/interfaces/cameraservice/common/2.0/android.frameworks.cameraservice.common@2.0_genc++/gen/android/frameworks/cameraservice/common/2.0/types.cpp
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
    add_library(android.frameworks.cameraservice.common-2.0 SHARED ${SOURCE_FILES})

    target_include_directories(android.frameworks.cameraservice.common-2.0
            PUBLIC
            ${ANDROID_ROOT}/frameworks/hardware/interfaces/cameraservice/common/2.0
            ${OUT_INTERMEDIATES_ROOT}/frameworks/hardware/interfaces/cameraservice/common/2.0/android.frameworks.cameraservice.common@2.0_genc++/gen
            ${OUT_INTERMEDIATES_ROOT}/frameworks/hardware/interfaces/cameraservice/common/2.0/android.frameworks.cameraservice.common@2.0_genc++_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(android.frameworks.cameraservice.common-2.0 PRIVATE
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.frameworks.cameraservice.common-2.0
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### android.frameworks.cameraservice.common-2.0 ####################

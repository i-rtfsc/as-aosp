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

project(android.hidl.base@1.0)

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

#################### android.hidl.base-1.0 ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++/gen/android/hidl/base/1.0/BaseAll.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++/gen/android/hidl/base/1.0/types.cpp
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
    add_library(android.hidl.base-1.0 SHARED ${SOURCE_FILES})

    target_include_directories(android.hidl.base-1.0
            PUBLIC
            ${ANDROID_ROOT}/system/libhidl/transport/base/1.0
            ${OUT_INTERMEDIATES_ROOT}/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++_headers/gen
            ${OUT_INTERMEDIATES_ROOT}/system/libhidl/transport/base/1.0/android.hidl.base@1.0_genc++/gen
    )

    target_link_libraries(android.hidl.base-1.0
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhidlbase
            libhidltransport
            libhwbinder
            liblog
            libm
            libutils
    )
endif ()
#################### android.hidl.base-1.0 ####################

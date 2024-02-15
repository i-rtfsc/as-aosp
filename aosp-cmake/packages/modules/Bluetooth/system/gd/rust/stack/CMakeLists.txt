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

project(libbt_hidl_hal_cxx)

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

#################### libbt_hidl_hal_cxx ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/stack/src/hal/ffi/hidl.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/stack/libbt_hidl_hal_bridge_code/gen/hidl_hal_generated.cc
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
    add_library(libbt_hidl_hal_cxx SHARED ${SOURCE_FILES})

    target_include_directories(libbt_hidl_hal_cxx
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/stack/src/hal/ffi
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/stack
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/stack/libbt_hidl_hal_bridge_code/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/stack/libbt_hidl_hal_bridge_header/gen
    )

    # Add c/c++ flags
    target_compile_options(libbt_hidl_hal_cxx PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libbt_hidl_hal_cxx
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libhidlbase
            libm
            libutils
    )
endif ()
#################### libbt_hidl_hal_cxx ####################

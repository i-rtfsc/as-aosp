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

project(libvtsclearkey)

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

#################### libvtsclearkey ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.2/vts/functional/drm_hal_clearkey_module.cpp
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
    add_library(libvtsclearkey SHARED ${SOURCE_FILES})

    target_include_directories(libvtsclearkey
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.2/vts/functional
    )

    # Add c/c++ flags
    target_compile_options(libvtsclearkey PRIVATE
            -Werror
    )

    target_link_libraries(libvtsclearkey
            drm_hal_vendor_module_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgtest
            libm
    )
endif ()
#################### libvtsclearkey ####################

#################### android.hardware.drm-1.2-vts ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.2/vts/functional/drm_hal_clearkey_module.cpp
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.2/vts/functional/drm_hal_common.cpp
        ${ANDROID_ROOT}/hardware/interfaces/drm/1.2/vts/functional/drm_hal_test.cpp
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
    add_library(android.hardware.drm-1.2-vts SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-1.2-vts
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.2/vts/functional/include
            ${ANDROID_ROOT}/hardware/interfaces/drm/1.2/vts/functional
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-1.2-vts PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.drm-1.2-vts
            VtsHalHidlTestUtils
            android.hardware.drm-1.0
            android.hardware.drm-1.0-helper
            android.hardware.drm-1.1
            android.hardware.drm-1.2
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libdrmvtshelper
            libhidl-gen-utils
            libhidl_gtest_helper
            libhidlbase
            libhidlmemory
            liblog
            libm
            libnativehelper
            libutils
    )
endif ()
#################### android.hardware.drm-1.2-vts ####################

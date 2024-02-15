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

project(libbt_keystore_cc)

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

#################### libbt_keystore_cc ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/common/keystore/fake_bt_keystore.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/common/libbt_common_bridge_code/gen/bridge.cc
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
    add_library(libbt_keystore_cc SHARED ${SOURCE_FILES})

    target_include_directories(libbt_keystore_cc
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/common
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/common/libbt_common_bridge_code/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/common/libbt_common_bridge_header/gen
    )

    # Add c/c++ flags
    target_compile_options(libbt_keystore_cc PRIVATE
            -Werror
    )

    target_link_libraries(libbt_keystore_cc
            libbt_keystore_cc_headers
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libbt_keystore_cc ####################

#################### libbt_common_bridge ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/common/libbt_common_bridge_code/gen/bridge.cc
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
    add_library(libbt_common_bridge SHARED ${SOURCE_FILES})

    target_include_directories(libbt_common_bridge
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/common
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/common/libbt_common_bridge_code/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/common/libbt_common_bridge_header/gen
    )

    # Add c/c++ flags
    target_compile_options(libbt_common_bridge PRIVATE
            -Wno-unused-const-variable
    )

    target_link_libraries(libbt_common_bridge
            libbt_keystore_cc_headers
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libbt_common_bridge ####################

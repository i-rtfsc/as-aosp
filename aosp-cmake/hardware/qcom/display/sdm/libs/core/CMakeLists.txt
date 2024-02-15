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

project(libsdmcore)

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

#################### libsdmcore ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/core_interface.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/core_impl.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/display_base.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/display_builtin.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/display_pluggable.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/display_virtual.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/display_null.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/noise_plugin_intf_impl.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/comp_manager.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/strategy.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/resource_default.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/color_manager.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core/hw_info_default.cpp
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
    add_library(libsdmcore SHARED ${SOURCE_FILES})

    target_include_directories(libsdmcore
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/core
    )

    # Add c/c++ flags
    target_compile_options(libsdmcore PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libsdmcore
            device_kernel_headers
            display_headers
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdisplaydebug
            libdl
            libdrm
            libdrmutils
            liblog
            libm
            libsdedrm
            libsdmdal
            libsdmutils
            libsynclight
            libutils
            libxml2
            qti_display_kernel_headers
            qti_kernel_headers
            libc
            libdl
            libm
    )
endif ()
#################### libsdmcore ####################

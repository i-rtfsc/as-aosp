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

project(libsdmdal)

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

#################### libsdmdal ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_events_interface.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_info_interface.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_interface.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_info_drm.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_device_drm.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_peripheral_drm.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_tv_drm.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_events_drm.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_scale_drm.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_virtual_drm.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal/hw_color_manager_drm.cpp
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
    add_library(libsdmdal SHARED ${SOURCE_FILES})

    target_include_directories(libsdmdal
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/sdm/libs/dal
    )

    # Add c/c++ flags
    target_compile_options(libsdmdal PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libsdmdal
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
            libsdmutils
            libutils
            qti_display_kernel_headers
            qti_kernel_headers
            libc
            libdl
            libm
    )
endif ()
#################### libsdmdal ####################

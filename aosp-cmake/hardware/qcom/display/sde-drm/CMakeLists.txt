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

project(libsdedrm)

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

#################### libsdedrm ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_manager.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_connector.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_encoder.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_crtc.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_plane.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_atomic_req.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_utils.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_pp_manager.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_property.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_dpps_mgr_imp.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/sde-drm/drm_panel_feature_mgr.cpp
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
    add_library(libsdedrm SHARED ${SOURCE_FILES})

    target_include_directories(libsdedrm
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/sde-drm
    )

    # Add c/c++ flags
    target_compile_options(libsdedrm PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libsdedrm
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
            libutils
            qti_display_kernel_headers
            qti_kernel_headers
            libc
            libdl
            libm
    )
endif ()
#################### libsdedrm ####################

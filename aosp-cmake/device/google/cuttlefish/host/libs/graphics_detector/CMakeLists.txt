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

project(libcuttlefish_graphics_configuration)

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

#################### libcuttlefish_graphics_configuration ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/graphics_configuration.cpp
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
    add_library(libcuttlefish_graphics_configuration SHARED ${SOURCE_FILES})

    target_include_directories(libcuttlefish_graphics_configuration
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector
    )

    # Add c/c++ flags
    target_compile_options(libcuttlefish_graphics_configuration PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libcuttlefish_graphics_configuration
            cuttlefish_common_headers
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcuttlefish_host_config
            libdl
            libjsoncpp
            liblog
            libm
    )
endif ()
#################### libcuttlefish_graphics_configuration ####################

#################### libcuttlefish_graphics_detector ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/egl.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/gles.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/img.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/lib.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/graphics_detector.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/graphics_detector_gl.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/graphics_detector_vk.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/graphics_detector_vk_precision_qualifiers_on_yuv_samplers.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/subprocess.cpp
        ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/vk.cpp
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
    add_library(libcuttlefish_graphics_detector SHARED ${SOURCE_FILES})

    target_include_directories(libcuttlefish_graphics_detector
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector/include
            ${ANDROID_ROOT}/device/google/cuttlefish/host/libs/graphics_detector
    )

    # Add c/c++ flags
    target_compile_options(libcuttlefish_graphics_detector PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libcuttlefish_graphics_detector
            cuttlefish_common_headers
            egl_headers
            gl_headers
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libcuttlefish_graphics_detector ####################

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

project(libgnssauto_power)

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

#################### libgnssauto_power ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/gps/android/auto_power/GnssAutoPowerHandler.cpp
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
    add_library(libgnssauto_power SHARED ${SOURCE_FILES})

    target_include_directories(libgnssauto_power
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/gps/android/auto_power
    )

    # Add c/c++ flags
    target_compile_options(libgnssauto_power PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libgnssauto_power
            android.hardware.automotive.vehicle-2.0
            android.hardware.gnss-1.0
            android.hardware.health-1.0
            android.hardware.health-2.0
            android.hardware.power-1.2
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgps.utils
            libgps.utils_headers
            libhidlbase
            libloc_core_headers
            libloc_pla_headers
            liblocation_api
            liblocation_api_headers
            liblog
            libm
            libutils
            vhal_v2_0_common_headers
            libc
            libdl
            libm
    )
endif ()
#################### libgnssauto_power ####################

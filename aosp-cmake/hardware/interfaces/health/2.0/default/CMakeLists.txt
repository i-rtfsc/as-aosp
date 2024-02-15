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

project(android.hardware.health@2.0-impl)

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

#################### android.hardware.health-2.0-impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/health/2.0/default/Health.cpp
        ${ANDROID_ROOT}/hardware/interfaces/health/2.0/default/healthd_common_adapter.cpp
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
    add_library(android.hardware.health-2.0-impl SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-2.0-impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/health/2.0/default/include
            ${ANDROID_ROOT}/hardware/interfaces/health/2.0/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-2.0-impl PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.health-2.0-impl
            android.hardware.health-1.0-convert
            android.hardware.health-2.0
            libbase
            libbatterymonitor
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealthloop
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### android.hardware.health-2.0-impl ####################

#################### android.hardware.health-2.0-impl-default ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/health/2.0/default/HealthImplDefault.cpp
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
    add_library(android.hardware.health-2.0-impl-default SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-2.0-impl-default
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/health/2.0/default
            ${ANDROID_ROOT}/hardware/interfaces/health/2.0/default/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-2.0-impl-default PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.health-2.0-impl-default
            android.hardware.health-1.0-convert
            android.hardware.health-2.0
            android.hardware.health-2.0-impl
            libbase
            libbatterymonitor
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealthstoragedefault
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### android.hardware.health-2.0-impl-default ####################

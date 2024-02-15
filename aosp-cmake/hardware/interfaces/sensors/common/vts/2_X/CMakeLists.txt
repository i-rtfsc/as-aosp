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

project(VtsHalSensorsV2_1TargetTest-lib)

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

#################### VtsHalSensorsV2_1TargetTest-lib ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/sensors/common/vts/2_X/SensorsHidlEnvironmentV2_X.cpp
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
    add_library(VtsHalSensorsV2_1TargetTest-lib SHARED ${SOURCE_FILES})

    target_include_directories(VtsHalSensorsV2_1TargetTest-lib
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/sensors/common/vts/2_X
    )

    # Add c/c++ flags
    target_compile_options(VtsHalSensorsV2_1TargetTest-lib PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(VtsHalSensorsV2_1TargetTest-lib
            VtsHalHidlTestUtils
            VtsHalSensorsTargetTestUtils
            android.hardware.sensors-1.0
            android.hardware.sensors-1.0-convert
            android.hardware.sensors-2.0
            android.hardware.sensors-2.1
            android.hardware.sensors-2.X-shared-utils
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libgtest
            libgtest_main
            libhidl-gen-utils
            libhidl_gtest_helper
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### VtsHalSensorsV2_1TargetTest-lib ####################

#################### VtsHalSensorsV2_0TargetTest-lib ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/sensors/common/vts/2_X/SensorsHidlEnvironmentV2_X.cpp
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
    add_library(VtsHalSensorsV2_0TargetTest-lib SHARED ${SOURCE_FILES})

    target_include_directories(VtsHalSensorsV2_0TargetTest-lib
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/sensors/common/vts/2_X
    )

    # Add c/c++ flags
    target_compile_options(VtsHalSensorsV2_0TargetTest-lib PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(VtsHalSensorsV2_0TargetTest-lib
            VtsHalHidlTestUtils
            VtsHalSensorsTargetTestUtils
            android.hardware.sensors-1.0
            android.hardware.sensors-1.0-convert
            android.hardware.sensors-2.0
            android.hardware.sensors-2.1
            android.hardware.sensors-2.X-shared-utils
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libgtest
            libgtest_main
            libhidl-gen-utils
            libhidl_gtest_helper
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### VtsHalSensorsV2_0TargetTest-lib ####################

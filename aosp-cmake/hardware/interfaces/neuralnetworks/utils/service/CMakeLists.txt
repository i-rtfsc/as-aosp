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

project(neuralnetworks_utils_hal_service)

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

#################### neuralnetworks_utils_hal_service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/service/src/Service.cpp
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
    add_library(neuralnetworks_utils_hal_service SHARED ${SOURCE_FILES})

    target_include_directories(neuralnetworks_utils_hal_service
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/service/include
            ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/service/include/nnapi/hal
            ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/utils/service
    )

    # Add c/c++ flags
    target_compile_options(neuralnetworks_utils_hal_service PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(neuralnetworks_utils_hal_service
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            neuralnetworks_types
            neuralnetworks_utils_hal_1_0
            neuralnetworks_utils_hal_1_1
            neuralnetworks_utils_hal_1_2
            neuralnetworks_utils_hal_1_3
            neuralnetworks_utils_hal_aidl
            neuralnetworks_utils_hal_common
    )
endif ()
#################### neuralnetworks_utils_hal_service ####################

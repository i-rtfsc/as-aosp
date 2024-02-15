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

project(neuralnetworks_utils_hal_1_0)

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

#################### neuralnetworks_utils_hal_1_0 ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/src/Assertions.cpp
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/src/Burst.cpp
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/src/Callbacks.cpp
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/src/Conversions.cpp
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/src/Device.cpp
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/src/Execution.cpp
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/src/PreparedModel.cpp
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/src/ProtectCallback.cpp
        ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/src/Service.cpp
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
    add_library(neuralnetworks_utils_hal_1_0 SHARED ${SOURCE_FILES})

    target_include_directories(neuralnetworks_utils_hal_1_0
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/include
            ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils/include/nnapi/hal/1.0
            ${ANDROID_ROOT}/hardware/interfaces/neuralnetworks/1.0/utils
    )

    # Add c/c++ flags
    target_compile_options(neuralnetworks_utils_hal_1_0 PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(neuralnetworks_utils_hal_1_0
            android.hardware.neuralnetworks-1.0
            libarect
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnativewindow
            neuralnetworks_types
            neuralnetworks_utils_hal_common
    )
endif ()
#################### neuralnetworks_utils_hal_1_0 ####################

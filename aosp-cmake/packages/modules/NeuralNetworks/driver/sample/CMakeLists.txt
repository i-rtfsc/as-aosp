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

project(android.hardware.neuralnetworks-service.example_fuzzer)

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

#################### neuralnetworks_sample_sl_driver ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/SampleDriverSL.cpp
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
    add_library(neuralnetworks_sample_sl_driver SHARED ${SOURCE_FILES})

    target_include_directories(neuralnetworks_sample_sl_driver
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample
    )

    # Add c/c++ flags
    target_compile_options(neuralnetworks_sample_sl_driver PRIVATE
            -Werror
    )

    target_link_libraries(neuralnetworks_sample_sl_driver
            libbase_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libdl
            liblog
            libm
            libnativewindow
            libneuralnetworks_cl
            libneuralnetworks_common_cl_cpu
            libtflite_static
            neuralnetworks_canonical_sample_driver_cl
            neuralnetworks_types_cl
    )
endif ()
#################### neuralnetworks_sample_sl_driver ####################

#################### neuralnetworks_canonical_sample_driver ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/CanonicalBuffer.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/CanonicalBurst.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/CanonicalDevice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/CanonicalPreparedModel.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/LimitedSupportDevice.cpp
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
    add_library(neuralnetworks_canonical_sample_driver SHARED ${SOURCE_FILES})

    target_include_directories(neuralnetworks_canonical_sample_driver
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample
    )

    # Add c/c++ flags
    target_compile_options(neuralnetworks_canonical_sample_driver PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(neuralnetworks_canonical_sample_driver
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnativewindow
            libneuralnetworks_common
            libneuralnetworks_headers
    )
endif ()
#################### neuralnetworks_canonical_sample_driver ####################

#################### neuralnetworks_canonical_sample_driver_cl ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/CanonicalBuffer.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/CanonicalBurst.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/CanonicalDevice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/CanonicalPreparedModel.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/LimitedSupportDevice.cpp
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
    add_library(neuralnetworks_canonical_sample_driver_cl SHARED ${SOURCE_FILES})

    target_include_directories(neuralnetworks_canonical_sample_driver_cl
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample
    )

    # Add c/c++ flags
    target_compile_options(neuralnetworks_canonical_sample_driver_cl PRIVATE
            -Werror
    )

    target_link_libraries(neuralnetworks_canonical_sample_driver_cl
            libbase_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnativewindow
            libneuralnetworks_common_cl
            libneuralnetworks_headers_ndk
            neuralnetworks_types_cl
    )
endif ()
#################### neuralnetworks_canonical_sample_driver_cl ####################

#################### android.hardware.neuralnetworks-service.example_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample/Fuzzer.cpp
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
    add_executable(android.hardware.neuralnetworks-service.example_fuzzer ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-service.example_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-service.example_fuzzer PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(android.hardware.neuralnetworks-service.example_fuzzer
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            libaidlcommonsupport
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libnativewindow
            libneuralnetworks_common
            libneuralnetworks_headers
            libtextclassifier_hash
            libutils
            neuralnetworks_canonical_sample_driver
            neuralnetworks_utils_hal_adapter_aidl
            neuralnetworks_utils_hal_aidl
    )
endif ()
#################### android.hardware.neuralnetworks-service.example_fuzzer ####################

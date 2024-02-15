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

project(android.hardware.neuralnetworks-service-sample-all)

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

#################### android.hardware.neuralnetworks-service-sample-limited ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidlLimited.cpp
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
    add_executable(android.hardware.neuralnetworks-service-sample-limited ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-service-sample-limited
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-service-sample-limited PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(android.hardware.neuralnetworks-service-sample-limited
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            libaidlcommonsupport
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
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
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-service-sample-limited ####################

#################### libSampleDriverAidl ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidl.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidlFull.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidlPartial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidlUtils.cpp
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
    add_library(libSampleDriverAidl SHARED ${SOURCE_FILES})

    target_include_directories(libSampleDriverAidl
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl
    )

    # Add c/c++ flags
    target_compile_options(libSampleDriverAidl PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(libSampleDriverAidl
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            libaidlcommonsupport
            libbase
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware
            libhidlbase
            libhidlmemory
            liblog
            libm
            libnativewindow
            libneuralnetworks_common
            libneuralnetworks_headers
            libtextclassifier_hash
            libutils
            neuralnetworks_utils_hal_aidl
            neuralnetworks_utils_hal_common
    )
endif ()
#################### libSampleDriverAidl ####################

#################### android.hardware.neuralnetworks-service-sample-all ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidl.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidlFull.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidlPartial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidlUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl/SampleDriverAidlAll.cpp
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
    add_executable(android.hardware.neuralnetworks-service-sample-all ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-service-sample-all
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_aidl
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-service-sample-all PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(android.hardware.neuralnetworks-service-sample-all
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            libaidlcommonsupport
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libfmq
            libhardware
            libhidlbase
            libhidlmemory
            liblog
            libm
            libnativewindow
            libneuralnetworks_common
            libneuralnetworks_headers
            libtextclassifier_hash
            libutils
            neuralnetworks_utils_hal_aidl
            neuralnetworks_utils_hal_common
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-service-sample-all ####################

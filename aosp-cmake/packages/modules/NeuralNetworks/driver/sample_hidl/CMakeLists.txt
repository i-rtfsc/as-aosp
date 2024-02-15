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

project(android.hardware.neuralnetworks@1.3-service-sample-all)

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

#################### android.hardware.neuralnetworks-1.3-service-sample-limited ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverLimited.cpp
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
    add_executable(android.hardware.neuralnetworks-1.3-service-sample-limited ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-1.3-service-sample-limited
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-1.3-service-sample-limited PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(android.hardware.neuralnetworks-1.3-service-sample-limited
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            libbase
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
            neuralnetworks_canonical_sample_driver
            neuralnetworks_utils_hal_adapter
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-1.3-service-sample-limited ####################

#################### android.hardware.neuralnetworks-1.3-service-sample-float-xnnpack ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriver.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverFull.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverPartial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverFloatXNNPACK.cpp
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
    add_executable(android.hardware.neuralnetworks-1.3-service-sample-float-xnnpack ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-1.3-service-sample-float-xnnpack
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-1.3-service-sample-float-xnnpack PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DNN_DEBUGGABLE
    )

    target_link_libraries(android.hardware.neuralnetworks-1.3-service-sample-float-xnnpack
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            libXNNPACK
            libbase
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
            libpthreadpool
            libtextclassifier_hash
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-1.3-service-sample-float-xnnpack ####################

#################### android.hardware.neuralnetworks-1.3-service-sample-all ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriver.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverFull.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverPartial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverUtils.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverAll.cpp
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
    add_executable(android.hardware.neuralnetworks-1.3-service-sample-all ${SOURCE_FILES})

    target_include_directories(android.hardware.neuralnetworks-1.3-service-sample-all
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.neuralnetworks-1.3-service-sample-all PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(android.hardware.neuralnetworks-1.3-service-sample-all
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            libbase
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
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.neuralnetworks-1.3-service-sample-all ####################

#################### libSampleDriver ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriver.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverFull.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverPartial.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl/SampleDriverUtils.cpp
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
    add_library(libSampleDriver SHARED ${SOURCE_FILES})

    target_include_directories(libSampleDriver
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/driver/sample_hidl
    )

    # Add c/c++ flags
    target_compile_options(libSampleDriver PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(libSampleDriver
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            libbase
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
    )
endif ()
#################### libSampleDriver ####################

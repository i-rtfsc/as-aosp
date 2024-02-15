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

project(libneuralnetworks_cl)

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

#################### libneuralnetworks_cl ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/BurstBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/CompilationBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionCallback.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionPlan.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Manager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Memory.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArchHasher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArgumentInfo.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/NeuralNetworks.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ServerFlag.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/SupportLibraryDiagnostic.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Telemetry.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/TypeManager.cpp
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
    add_library(libneuralnetworks_cl SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks_cl
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/include
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks_cl PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(libneuralnetworks_cl
            libbase_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libdl
            libm
            libnativewindow
            libneuralnetworks_common_cl
            libneuralnetworks_headers_ndk
            neuralnetworks_supportlibrary_types_ndk
            neuralnetworks_types_cl
    )
endif ()
#################### libneuralnetworks_cl ####################

#################### libneuralnetworks_v2_static_experimental ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/AppInfoFetcher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/BurstBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/CompilationBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionCallback.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionPlan.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Manager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Memory.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArchHasher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArgumentInfo.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ServerFlag.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Telemetry.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/TypeManager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/FlatbufferModelBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/NeuralNetworksV2.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/operation_converters/AddOperationConverter.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/operation_converters/ArithmeticOperationConverter.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/operation_converters/Conv2DOperationConverter.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/operation_converters/DepthwiseConv2DOperationConverter.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/operation_converters/LogisticOperationConverter.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/operation_converters/OperationConverterResolver.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/operation_converters/SubGraphContext.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/TelemetryStatsd.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen/statslog_neuralnetworks.cpp
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
    add_library(libneuralnetworks_v2_static_experimental SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks_v2_static_experimental
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/packageinfo/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks_v2_static_experimental PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(libneuralnetworks_v2_static_experimental
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libaidlcommonsupport
            libbase
            libbinder_ndk
            libbuildversion
            libc
            libc++_static
            libc++demangle
            libcgrouprc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libcutils
            libdl
            libfmq
            libhidlbase
            libhidlmemory
            libjsoncpp
            liblog
            libm
            libmath
            libnativewindow
            libneuralnetworks_common_experimental
            libneuralnetworks_headers
            libneuralnetworks_packageinfo
            libprocessgroup
            libprocpartition
            libstatssocket
            libtextclassifier_hash_static
            libtflite_static
            libutils
            libvndksupport
            neuralnetworks_types_experimental
            neuralnetworks_utils_hal_aidl
            neuralnetworks_utils_hal_service
    )
endif ()
#################### libneuralnetworks_v2_static_experimental ####################

#################### libneuralnetworks_static ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/AppInfoFetcher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/BurstBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/CompilationBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionCallback.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionPlan.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Manager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Memory.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArchHasher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArgumentInfo.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/NeuralNetworks.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ServerFlag.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Telemetry.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/TypeManager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/TelemetryStatsd.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen/statslog_neuralnetworks.cpp
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
    add_library(libneuralnetworks_static SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks_static
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/packageinfo/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks_static PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(libneuralnetworks_static
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libaidlcommonsupport
            libbase
            libbinder_ndk
            libbuildversion
            libc
            libc++_static
            libc++demangle
            libcgrouprc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libcutils
            libdl
            libfmq
            libhidlbase
            libhidlmemory
            libjsoncpp
            liblog
            libm
            libmath
            libnativewindow
            libneuralnetworks_common
            libneuralnetworks_headers
            libneuralnetworks_packageinfo
            libprocessgroup
            libprocpartition
            libstatssocket
            libtextclassifier_hash_static
            libutils
            libvndksupport
            neuralnetworks_types
            neuralnetworks_utils_hal_aidl
            neuralnetworks_utils_hal_service
            server_configurable_flags
    )
endif ()
#################### libneuralnetworks_static ####################

#################### libneuralnetworks ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/AppInfoFetcher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/BurstBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/CompilationBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionCallback.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionPlan.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Manager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Memory.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArchHasher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArgumentInfo.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/NeuralNetworks.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ServerFlag.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Telemetry.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/TypeManager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/TelemetryStatsd.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen/statslog_neuralnetworks.cpp
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
    add_library(libneuralnetworks SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/packageinfo/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(libneuralnetworks
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libaidlcommonsupport
            libbase
            libbinder_ndk
            libbuildversion
            libc
            libc++_static
            libcgrouprc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libcutils
            libdl
            libfmq
            libhidlbase
            libhidlmemory
            libjsoncpp
            liblog
            libm
            libmath
            libnativewindow
            libneuralnetworks_common
            libneuralnetworks_headers
            libneuralnetworks_packageinfo
            libprocessgroup
            libprocpartition
            libstatssocket
            libtextclassifier_hash_static
            libutils
            libvndksupport
            neuralnetworks_types
            neuralnetworks_utils_hal_aidl
            neuralnetworks_utils_hal_service
            server_configurable_flags
    )
endif ()
#################### libneuralnetworks ####################

#################### libneuralnetworks_static_experimental ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/AppInfoFetcher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/BurstBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/CompilationBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionCallback.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ExecutionPlan.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Manager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Memory.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArchHasher.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelArgumentInfo.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ModelBuilder.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/NeuralNetworks.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/ServerFlag.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/Telemetry.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/TypeManager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/TelemetryStatsd.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen/statslog_neuralnetworks.cpp
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
    add_library(libneuralnetworks_static_experimental SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks_static_experimental
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/runtime/packageinfo/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/NeuralNetworks/runtime/statslog_neuralnetworks.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks_static_experimental PRIVATE
            -Wall
            -Werror
            -Wextra
            -DNN_DEBUGGABLE
    )

    target_link_libraries(libneuralnetworks_static_experimental
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            android.hardware.neuralnetworks-1.0
            android.hardware.neuralnetworks-1.1
            android.hardware.neuralnetworks-1.2
            android.hardware.neuralnetworks-1.3
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libaidlcommonsupport
            libbase
            libbinder_ndk
            libbuildversion
            libc
            libc++_static
            libc++demangle
            libcgrouprc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libcutils
            libdl
            libfmq
            libhidlbase
            libhidlmemory
            libjsoncpp
            liblog
            libm
            libmath
            libnativewindow
            libneuralnetworks_common_experimental
            libneuralnetworks_headers
            libneuralnetworks_packageinfo
            libprocessgroup
            libprocpartition
            libstatssocket
            libtextclassifier_hash_static
            libutils
            libvndksupport
            neuralnetworks_types_experimental
            neuralnetworks_utils_hal_aidl
            neuralnetworks_utils_hal_service
    )
endif ()
#################### libneuralnetworks_static_experimental ####################

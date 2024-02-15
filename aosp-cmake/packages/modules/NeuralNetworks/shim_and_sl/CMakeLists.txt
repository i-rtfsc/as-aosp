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

project(libneuralnetworks_shim_static)

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

#################### libneuralnetworks_shim_static ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/NeuralNetworksShim.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/ShimBufferTracker.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/ShimConverter.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/ShimDevice.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/ShimDeviceManager.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/ShimPreparedModel.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/ShimUtils.cpp
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
    add_library(libneuralnetworks_shim_static SHARED ${SOURCE_FILES})

    target_include_directories(libneuralnetworks_shim_static
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/public
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl
    )

    # Add c/c++ flags
    target_compile_options(libneuralnetworks_shim_static PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libneuralnetworks_shim_static
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            android.hardware.neuralnetworks-V4-ndk
            libaidlcommonsupport
            libarect
            libbase
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            libhidlmemory
            liblog
            libm
            libnativewindow
            libneuralnetworks_common
            libneuralnetworks_headers
            neuralnetworks_supportlibrary_loader
            neuralnetworks_utils_hal_aidl
            neuralnetworks_utils_hal_common
    )
endif ()
#################### libneuralnetworks_shim_static ####################

#################### neuralnetworks_supportlibrary_loader ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/SupportLibrary.cpp
        ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/SupportLibraryWrapper.cpp
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
    add_library(neuralnetworks_supportlibrary_loader SHARED ${SOURCE_FILES})

    target_include_directories(neuralnetworks_supportlibrary_loader
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/include
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl/public
            ${ANDROID_ROOT}/packages/modules/NeuralNetworks/shim_and_sl
    )

    # Add c/c++ flags
    target_compile_options(neuralnetworks_supportlibrary_loader PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(neuralnetworks_supportlibrary_loader
            libarect
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnativewindow
            libneuralnetworks_headers
    )
endif ()
#################### neuralnetworks_supportlibrary_loader ####################

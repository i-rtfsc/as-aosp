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

project(androidx.core.uwb.backend.aidl_interface-cpp-analyzer)

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

#################### androidx.core.uwb.backend.aidl_interface-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/IRangingSessionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/IUwb.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/IUwbClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/RangingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/RangingMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/RangingParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/RangingPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/UwbAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/UwbComplexChannel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/androidx/core/uwb/backend/UwbDevice.cpp
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
    add_library(androidx.core.uwb.backend.aidl_interface-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(androidx.core.uwb.backend.aidl_interface-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(androidx.core.uwb.backend.aidl_interface-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(androidx.core.uwb.backend.aidl_interface-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### androidx.core.uwb.backend.aidl_interface-V1-ndk_platform ####################

#################### androidx.core.uwb.backend.aidl_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/IRangingSessionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/IUwb.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/IUwbClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/RangingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/RangingMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/RangingParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/RangingPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/UwbAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/UwbComplexChannel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/androidx/core/uwb/backend/UwbDevice.cpp
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
    add_executable(androidx.core.uwb.backend.aidl_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(androidx.core.uwb.backend.aidl_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Uwb/androidx_backend
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(androidx.core.uwb.backend.aidl_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(androidx.core.uwb.backend.aidl_interface-cpp-analyzer
            aidl-analyzer-main
            androidx.core.uwb.backend.aidl_interface-V1-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### androidx.core.uwb.backend.aidl_interface-cpp-analyzer ####################

#################### androidx.core.uwb.backend.aidl_interface-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/IRangingSessionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/IUwb.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/IUwbClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/RangingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/RangingMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/RangingParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/RangingPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/UwbAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/UwbComplexChannel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/androidx/core/uwb/backend/UwbDevice.cpp
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
    add_library(androidx.core.uwb.backend.aidl_interface-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(androidx.core.uwb.backend.aidl_interface-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(androidx.core.uwb.backend.aidl_interface-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(androidx.core.uwb.backend.aidl_interface-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### androidx.core.uwb.backend.aidl_interface-V1-ndk ####################

#################### androidx.core.uwb.backend.aidl_interface-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/IRangingSessionCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/IUwb.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/IUwbClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/RangingCapabilities.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/RangingMeasurement.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/RangingParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/RangingPosition.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/UwbAddress.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/UwbComplexChannel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/androidx/core/uwb/backend/UwbDevice.cpp
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
    add_library(androidx.core.uwb.backend.aidl_interface-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(androidx.core.uwb.backend.aidl_interface-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Uwb/androidx_backend/androidx.core.uwb.backend.aidl_interface-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(androidx.core.uwb.backend.aidl_interface-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(androidx.core.uwb.backend.aidl_interface-V1-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### androidx.core.uwb.backend.aidl_interface-V1-cpp ####################

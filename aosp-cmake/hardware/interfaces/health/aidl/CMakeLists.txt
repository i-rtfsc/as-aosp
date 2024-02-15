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

project(android.hardware.health-V1-ndk_platform)

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

#################### android.hardware.health-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk_platform-source/gen/android/hardware/health/BatteryCapacityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk_platform-source/gen/android/hardware/health/BatteryHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk_platform-source/gen/android/hardware/health/BatteryStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk_platform-source/gen/android/hardware/health/DiskStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk_platform-source/gen/android/hardware/health/HealthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk_platform-source/gen/android/hardware/health/IHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk_platform-source/gen/android/hardware/health/IHealthInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk_platform-source/gen/android/hardware/health/StorageInfo.cpp
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
    add_library(android.hardware.health-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.health-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.health-V1-ndk_platform ####################

#################### android.hardware.health-translate-V1-ndk ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/android/hardware/health/translate-ndk.cpp
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
    add_library(android.hardware.health-translate-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-translate-V1-ndk
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/include
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-translate-V1-ndk PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.health-translate-V1-ndk
            android.hardware.health-V1-ndk
            android.hardware.health-2.0
            android.hardware.health-2.1
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libm
    )
endif ()
#################### android.hardware.health-translate-V1-ndk ####################

#################### android.hardware.health-translate-ndk ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/android/hardware/health/translate-ndk.cpp
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
    add_library(android.hardware.health-translate-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-translate-ndk
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/include
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-translate-ndk PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.health-translate-ndk
            android.hardware.health-V2-ndk
            android.hardware.health-2.0
            android.hardware.health-2.1
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlbase
            libm
    )
endif ()
#################### android.hardware.health-translate-ndk ####################

#################### android.hardware.health-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/BatteryCapacityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/BatteryChargingPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/BatteryChargingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/BatteryHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/BatteryHealthData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/BatteryStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/DiskStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/HealthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/IHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/IHealthInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/android/hardware/health/StorageInfo.cpp
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
    add_library(android.hardware.health-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.health-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.health-V3-ndk ####################

#################### android.hardware.health-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/BatteryCapacityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/BatteryChargingPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/BatteryChargingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/BatteryHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/BatteryHealthData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/BatteryStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/DiskStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/HealthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/IHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/IHealthInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/android/hardware/health/StorageInfo.cpp
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
    add_library(android.hardware.health-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.health-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.health-V2-ndk ####################

#################### android.hardware.health-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/BatteryCapacityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/BatteryChargingPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/BatteryChargingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/BatteryHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/BatteryHealthData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/BatteryStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/DiskStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/HealthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/IHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/IHealthInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/android/hardware/health/StorageInfo.cpp
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
    add_library(android.hardware.health-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.health-V2-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.health-V2-ndk_platform ####################

#################### android.hardware.health-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/android/hardware/health/BatteryCapacityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/android/hardware/health/BatteryHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/android/hardware/health/BatteryStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/android/hardware/health/DiskStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/android/hardware/health/HealthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/android/hardware/health/IHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/android/hardware/health/IHealthInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/android/hardware/health/StorageInfo.cpp
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
    add_library(android.hardware.health-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.health-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.health-V1-ndk ####################

#################### android.hardware.health-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/BatteryCapacityLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/BatteryChargingPolicy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/BatteryChargingState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/BatteryHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/BatteryHealthData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/BatteryStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/DiskStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/HealthInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/IHealth.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/IHealthInfoCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/android/hardware/health/StorageInfo.cpp
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
    add_library(android.hardware.health-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/health/aidl/android.hardware.health-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.health-V3-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
    )
endif ()
#################### android.hardware.health-V3-ndk_platform ####################

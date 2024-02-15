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

project(android.automotive.watchdog-cpp-analyzer)

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

#################### android.automotive.watchdog-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_library(android.automotive.watchdog-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.automotive.watchdog-V2-ndk ####################

#################### android.automotive.watchdog-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/IResourceOveruseListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/IoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/PerStateBytes.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/ResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_library(android.automotive.watchdog-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.automotive.watchdog-V3-ndk ####################

#################### android.automotive.watchdog-V4-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/IResourceOveruseListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/IoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/PerStateBytes.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/ResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_library(android.automotive.watchdog-V4-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-V4-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-V4-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-V4-cpp
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
#################### android.automotive.watchdog-V4-cpp ####################

#################### android.automotive.watchdog.internal-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ApplicationCategoryType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/CpuUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/GarageMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ICarWatchdogServiceForSystem.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/IoOveruseAlertThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/IoOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/IoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageIoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/PerStateIoOveruseThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ProcessCpuUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ProcessIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ResourceOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ResourceSpecificConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ResourceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ResourceUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/SystemSummaryUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/ThreadPolicyWithPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/UidIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/UidResourceUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/UidType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/UserPackageIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/android/automotive/watchdog/internal/UserState.cpp
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
    add_library(android.automotive.watchdog.internal-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog.internal-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog.internal-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog.internal-V3-ndk_platform
            android.automotive.watchdog-V3-ndk_platform
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
#################### android.automotive.watchdog.internal-V3-ndk_platform ####################

#################### android.automotive.watchdog-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-cpp-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-cpp-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-cpp-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-cpp-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-cpp-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-cpp-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-cpp-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-cpp-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_library(android.automotive.watchdog-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-V2-cpp
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
#################### android.automotive.watchdog-V2-cpp ####################

#################### android.automotive.watchdog-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/IResourceOveruseListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/IoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/PerStateBytes.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/ResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_executable(android.automotive.watchdog-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/watchdog/aidl
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-cpp-analyzer
            aidl-analyzer-main
            android.automotive.watchdog-V4-cpp
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
#################### android.automotive.watchdog-cpp-analyzer ####################

#################### android.automotive.watchdog.internal-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/ApplicationCategoryType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/ComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/GarageMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/ICarWatchdogServiceForSystem.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/IoOveruseAlertThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/IoOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/IoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageIoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/PerStateIoOveruseThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/ProcessIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/ResourceOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/ResourceSpecificConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/ThreadPolicyWithPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/UidType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/UserPackageIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/android/automotive/watchdog/internal/UserState.cpp
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
    add_library(android.automotive.watchdog.internal-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog.internal-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog.internal-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog.internal-V2-ndk_platform
            android.automotive.watchdog-V3-ndk_platform
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
#################### android.automotive.watchdog.internal-V2-ndk_platform ####################

#################### android.automotive.watchdog-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/IResourceOveruseListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/IoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/PerStateBytes.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/ResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_library(android.automotive.watchdog-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-V3-cpp
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
#################### android.automotive.watchdog-V3-cpp ####################

#################### android.automotive.watchdog-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/IResourceOveruseListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/IoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/PerStateBytes.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/ResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_library(android.automotive.watchdog-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-V4-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.automotive.watchdog-V4-ndk ####################

#################### android.automotive.watchdog.internal-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/ApplicationCategoryType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/ComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/GarageMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/ICarWatchdogServiceForSystem.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/IoOveruseAlertThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/IoOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/IoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/PackageIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/PackageInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/PackageIoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/PackageMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/PerStateIoOveruseThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/ProcessIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/ResourceOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/ResourceSpecificConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/UidType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/UserPackageIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/android/automotive/watchdog/internal/UserState.cpp
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
    add_library(android.automotive.watchdog.internal-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog.internal-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog.internal-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog.internal-V1-ndk
            android.automotive.watchdog-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.automotive.watchdog.internal-V1-ndk ####################

#################### android.automotive.watchdog.internal-V3-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ApplicationCategoryType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/CpuUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/GarageMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ICarWatchdogServiceForSystem.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/IoOveruseAlertThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/IoOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/IoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/PackageIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/PackageInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/PackageIoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/PackageMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/PerStateIoOveruseThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ProcessCpuUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ProcessIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ResourceOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ResourceSpecificConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ResourceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ResourceUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/SystemSummaryUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/ThreadPolicyWithPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/UidIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/UidResourceUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/UidType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/UserPackageIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/android/automotive/watchdog/internal/UserState.cpp
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
    add_library(android.automotive.watchdog.internal-V3-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog.internal-V3-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog.internal-V3-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog.internal-V3-cpp
            android.automotive.watchdog-V3-cpp
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
#################### android.automotive.watchdog.internal-V3-cpp ####################

#################### android.automotive.watchdog.internal-V2-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/ApplicationCategoryType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/ComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/GarageMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/ICarWatchdogServiceForSystem.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/IoOveruseAlertThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/IoOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/IoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/PackageIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/PackageInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/PackageIoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/PackageMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/PerStateIoOveruseThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/ProcessIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/ResourceOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/ResourceSpecificConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/ThreadPolicyWithPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/UidType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/UserPackageIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/android/automotive/watchdog/internal/UserState.cpp
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
    add_library(android.automotive.watchdog.internal-V2-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog.internal-V2-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog.internal-V2-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog.internal-V2-cpp
            android.automotive.watchdog-V3-cpp
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
#################### android.automotive.watchdog.internal-V2-cpp ####################

#################### android.automotive.watchdog-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk_platform-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk_platform-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk_platform-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk_platform-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk_platform-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk_platform-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk_platform-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk_platform-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_library(android.automotive.watchdog-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-V2-ndk_platform
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
#################### android.automotive.watchdog-V2-ndk_platform ####################

#################### android.automotive.watchdog.internal-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ApplicationCategoryType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/CpuUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/GarageMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ICarWatchdogServiceForSystem.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/IoOveruseAlertThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/IoOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/IoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/PackageIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/PackageInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/PackageIoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/PackageMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/PerStateIoOveruseThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ProcessCpuUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ProcessIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ResourceOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ResourceSpecificConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ResourceStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ResourceUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/SystemSummaryUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/ThreadPolicyWithPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/UidIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/UidResourceUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/UidType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/UserPackageIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/android/automotive/watchdog/internal/UserState.cpp
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
    add_library(android.automotive.watchdog.internal-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog.internal-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog.internal-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog.internal-V3-ndk
            android.automotive.watchdog-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.automotive.watchdog.internal-V3-ndk ####################

#################### android.automotive.watchdog.internal-V1-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/ApplicationCategoryType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/ComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/GarageMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/ICarWatchdogServiceForSystem.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/IoOveruseAlertThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/IoOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/IoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/PackageIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/PackageInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/PackageIoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/PackageMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/PerStateIoOveruseThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/ProcessIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/ResourceOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/ResourceSpecificConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/UidType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/UserPackageIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/android/automotive/watchdog/internal/UserState.cpp
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
    add_library(android.automotive.watchdog.internal-V1-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog.internal-V1-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog.internal-V1-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog.internal-V1-cpp
            android.automotive.watchdog-V3-cpp
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
#################### android.automotive.watchdog.internal-V1-cpp ####################

#################### android.automotive.watchdog-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/IResourceOveruseListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/IoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/PerStateBytes.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/ResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_library(android.automotive.watchdog-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-V3-ndk_platform
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
#################### android.automotive.watchdog-V3-ndk_platform ####################

#################### android.automotive.watchdog-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/ICarWatchdogClient.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/IResourceOveruseListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/IoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/PerStateBytes.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/ResourceOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/ResourceType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/android/automotive/watchdog/UserState.cpp
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
    add_library(android.automotive.watchdog-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog-V4-ndk_platform
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
#################### android.automotive.watchdog-V4-ndk_platform ####################

#################### android.automotive.watchdog.internal-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/ApplicationCategoryType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/ComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/GarageMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/ICarWatchdogServiceForSystem.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/IoOveruseAlertThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/IoOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/IoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/PackageIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/PackageInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/PackageIoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/PackageMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/PerStateIoOveruseThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/ProcessIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/ResourceOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/ResourceSpecificConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/ThreadPolicyWithPriority.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/UidType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/UserPackageIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/android/automotive/watchdog/internal/UserState.cpp
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
    add_library(android.automotive.watchdog.internal-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog.internal-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog.internal-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog.internal-V2-ndk
            android.automotive.watchdog-V3-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.automotive.watchdog.internal-V2-ndk ####################

#################### android.automotive.watchdog.internal-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/ApplicationCategoryType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/BootPhase.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/ComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/GarageMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/ICarWatchdog.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/ICarWatchdogMonitor.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/ICarWatchdogServiceForSystem.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/IoOveruseAlertThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/IoOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/IoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageIoOveruseStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/PackageMetadata.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/PerStateIoOveruseThreshold.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/PowerCycle.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/ProcessIdentifier.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/ResourceOveruseConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/ResourceSpecificConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/StateType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/TimeoutLength.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/UidType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/UserPackageIoUsageStats.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/android/automotive/watchdog/internal/UserState.cpp
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
    add_library(android.automotive.watchdog.internal-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.automotive.watchdog.internal-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog-V3-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/watchdog/aidl/android.automotive.watchdog.internal-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.watchdog.internal-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.automotive.watchdog.internal-V1-ndk_platform
            android.automotive.watchdog-V3-ndk_platform
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
#################### android.automotive.watchdog.internal-V1-ndk_platform ####################

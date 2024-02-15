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

project(android.system.virtualizationcommon-cpp-analyzer)

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

#################### android.system.virtualizationcommon-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk_platform-source/gen/android/system/virtualizationcommon/DeathReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk_platform-source/gen/android/system/virtualizationcommon/ErrorCode.cpp
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
    add_library(android.system.virtualizationcommon-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationcommon-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationcommon-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationcommon-ndk_platform
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
#################### android.system.virtualizationcommon-ndk_platform ####################

#################### android.system.virtualizationcommon-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk-source/gen/android/system/virtualizationcommon/DeathReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk-source/gen/android/system/virtualizationcommon/ErrorCode.cpp
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
    add_library(android.system.virtualizationcommon-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationcommon-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationcommon-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationcommon-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.virtualizationcommon-ndk ####################

#################### android.system.virtualizationservice_internal-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk-source/gen/android/system/virtualizationservice_internal/AtomVmBooted.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk-source/gen/android/system/virtualizationservice_internal/AtomVmCreationRequested.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk-source/gen/android/system/virtualizationservice_internal/AtomVmExited.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk-source/gen/android/system/virtualizationservice_internal/IGlobalVmContext.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk-source/gen/android/system/virtualizationservice_internal/IVirtualizationServiceInternal.cpp
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
    add_library(android.system.virtualizationservice_internal-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationservice_internal-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationservice_internal-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationservice_internal-ndk
            android.system.virtualizationcommon-ndk
            android.system.virtualizationservice-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.virtualizationservice_internal-ndk ####################

#################### android.system.virtualizationcommon-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-analyzer-source/gen/android/system/virtualizationcommon/DeathReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-analyzer-source/gen/android/system/virtualizationcommon/ErrorCode.cpp
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
    add_executable(android.system.virtualizationcommon-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationcommon-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationcommon-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationcommon-cpp-analyzer
            aidl-analyzer-main
            android.system.virtualizationcommon-cpp
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
#################### android.system.virtualizationcommon-cpp-analyzer ####################

#################### android.system.virtualmachineservice-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualmachineservice-ndk_platform-source/gen/android/system/virtualmachineservice/IVirtualMachineService.cpp
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
    add_library(android.system.virtualmachineservice-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualmachineservice-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualmachineservice-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualmachineservice-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualmachineservice-ndk_platform
            android.system.virtualizationcommon-ndk_platform
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
#################### android.system.virtualmachineservice-ndk_platform ####################

#################### android.system.virtualizationservice_internal-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-cpp-source/gen/android/system/virtualizationservice_internal/AtomVmBooted.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-cpp-source/gen/android/system/virtualizationservice_internal/AtomVmCreationRequested.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-cpp-source/gen/android/system/virtualizationservice_internal/AtomVmExited.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-cpp-source/gen/android/system/virtualizationservice_internal/IGlobalVmContext.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-cpp-source/gen/android/system/virtualizationservice_internal/IVirtualizationServiceInternal.cpp
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
    add_library(android.system.virtualizationservice_internal-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationservice_internal-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationservice_internal-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationservice_internal-cpp
            android.system.virtualizationcommon-cpp
            android.system.virtualizationservice-cpp
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
#################### android.system.virtualizationservice_internal-cpp ####################

#################### android.system.virtualizationcommon-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-source/gen/android/system/virtualizationcommon/DeathReason.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-source/gen/android/system/virtualizationcommon/ErrorCode.cpp
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
    add_library(android.system.virtualizationcommon-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationcommon-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationcommon-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationcommon-cpp
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
#################### android.system.virtualizationcommon-cpp ####################

#################### android.system.virtualizationservice-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/CpuTopology.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/DiskImage.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/IVirtualMachine.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/IVirtualMachineCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/IVirtualizationService.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/MemoryTrimLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/Partition.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/PartitionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/VirtualMachineAppConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/VirtualMachineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/VirtualMachineDebugInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/VirtualMachinePayloadConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/VirtualMachineRawConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/android/system/virtualizationservice/VirtualMachineState.cpp
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
    add_library(android.system.virtualizationservice-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationservice-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationservice-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationservice-cpp
            android.system.virtualizationcommon-cpp
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
#################### android.system.virtualizationservice-cpp ####################

#################### android.system.virtualizationservice_internal-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk_platform-source/gen/android/system/virtualizationservice_internal/AtomVmBooted.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk_platform-source/gen/android/system/virtualizationservice_internal/AtomVmCreationRequested.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk_platform-source/gen/android/system/virtualizationservice_internal/AtomVmExited.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk_platform-source/gen/android/system/virtualizationservice_internal/IGlobalVmContext.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk_platform-source/gen/android/system/virtualizationservice_internal/IVirtualizationServiceInternal.cpp
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
    add_library(android.system.virtualizationservice_internal-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationservice_internal-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice_internal-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationservice_internal-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationservice_internal-ndk_platform
            android.system.virtualizationcommon-ndk_platform
            android.system.virtualizationservice-ndk_platform
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
#################### android.system.virtualizationservice_internal-ndk_platform ####################

#################### android.system.virtualmachineservice-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualmachineservice-ndk-source/gen/android/system/virtualmachineservice/IVirtualMachineService.cpp
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
    add_library(android.system.virtualmachineservice-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualmachineservice-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualmachineservice-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualmachineservice-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualmachineservice-ndk
            android.system.virtualizationcommon-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.virtualmachineservice-ndk ####################

#################### android.system.virtualizationservice-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/CpuTopology.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/DiskImage.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/IVirtualMachine.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/IVirtualMachineCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/IVirtualizationService.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/MemoryTrimLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/Partition.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/PartitionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/VirtualMachineAppConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/VirtualMachineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/VirtualMachineDebugInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/VirtualMachinePayloadConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/VirtualMachineRawConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/android/system/virtualizationservice/VirtualMachineState.cpp
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
    add_library(android.system.virtualizationservice-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationservice-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk_platform-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationservice-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationservice-ndk_platform
            android.system.virtualizationcommon-ndk_platform
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
#################### android.system.virtualizationservice-ndk_platform ####################

#################### android.system.virtualizationservice-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/CpuTopology.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/DiskImage.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/IVirtualMachine.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/IVirtualMachineCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/IVirtualizationService.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/MemoryTrimLevel.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/Partition.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/PartitionType.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/VirtualMachineAppConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/VirtualMachineConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/VirtualMachineDebugInfo.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/VirtualMachinePayloadConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/VirtualMachineRawConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/android/system/virtualizationservice/VirtualMachineState.cpp
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
    add_library(android.system.virtualizationservice-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualizationservice-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-ndk-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationservice-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualizationservice-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualizationservice-ndk
            android.system.virtualizationcommon-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.virtualizationservice-ndk ####################

#################### android.system.virtualmachineservice-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualmachineservice-cpp-source/gen/android/system/virtualmachineservice/IVirtualMachineService.cpp
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
    add_library(android.system.virtualmachineservice-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualmachineservice-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualizationcommon-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/virtualizationservice/aidl/android.system.virtualmachineservice-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualmachineservice-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualmachineservice-cpp
            android.system.virtualizationcommon-cpp
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
#################### android.system.virtualmachineservice-cpp ####################

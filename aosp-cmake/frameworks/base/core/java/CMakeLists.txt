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

project(android.os.statsbootstrap_aidl-cpp-analyzer)

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

#################### libactivity_manager_procstate_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libactivity_manager_procstate_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/app/ProcessStateEnum.cpp
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
    add_library(libactivity_manager_procstate_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(libactivity_manager_procstate_aidl-cpp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/core/java
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libactivity_manager_procstate_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libactivity_manager_procstate_aidl-cpp PRIVATE
            -Werror
    )

    target_link_libraries(libactivity_manager_procstate_aidl-cpp
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libactivity_manager_procstate_aidl-cpp ####################

#################### android.os.statsbootstrap_aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-ndk-source/gen/android/os/IStatsBootstrapAtomService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-ndk-source/gen/android/os/StatsBootstrapAtom.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-ndk-source/gen/android/os/StatsBootstrapAtomValue.cpp
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
    add_library(android.os.statsbootstrap_aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.os.statsbootstrap_aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.os.statsbootstrap_aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.os.statsbootstrap_aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.os.statsbootstrap_aidl-ndk ####################

#################### android.debug_aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-analyzer-source/gen/android/debug/AdbTransportType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-analyzer-source/gen/android/debug/FingerprintAndPairDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-analyzer-source/gen/android/debug/IAdbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-analyzer-source/gen/android/debug/IAdbManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-analyzer-source/gen/android/debug/PairDevice.cpp
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
    add_executable(android.debug_aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.debug_aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/core/java
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.debug_aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.debug_aidl-cpp-analyzer
            aidl-analyzer-main
            android.debug_aidl-cpp
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
#################### android.debug_aidl-cpp-analyzer ####################

#################### android.os.permissions_aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.permissions_aidl-cpp-analyzer-source/gen/android/os/IPermissionController.cpp
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
    add_executable(android.os.permissions_aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.os.permissions_aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/core/java
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.permissions_aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.permissions_aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.os.permissions_aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.os.permissions_aidl-cpp-analyzer
            aidl-analyzer-main
            android.os.permissions_aidl-cpp
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
#################### android.os.permissions_aidl-cpp-analyzer ####################

#################### android.os.permissions_aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.permissions_aidl-ndk-source/gen/android/os/IPermissionController.cpp
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
    add_library(android.os.permissions_aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.os.permissions_aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.permissions_aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.os.permissions_aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.os.permissions_aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.os.permissions_aidl-ndk ####################

#################### android.debug_aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk_platform-source/gen/android/debug/AdbTransportType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk_platform-source/gen/android/debug/FingerprintAndPairDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk_platform-source/gen/android/debug/IAdbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk_platform-source/gen/android/debug/IAdbManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk_platform-source/gen/android/debug/PairDevice.cpp
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
    add_library(android.debug_aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.debug_aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.debug_aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.debug_aidl-ndk_platform
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
#################### android.debug_aidl-ndk_platform ####################

#################### android.os.permissions_aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.permissions_aidl-ndk_platform-source/gen/android/os/IPermissionController.cpp
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
    add_library(android.os.permissions_aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.os.permissions_aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.permissions_aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.os.permissions_aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.os.permissions_aidl-ndk_platform
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
#################### android.os.permissions_aidl-ndk_platform ####################

#################### android.os.statsbootstrap_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-cpp-source/gen/android/os/IStatsBootstrapAtomService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-cpp-source/gen/android/os/StatsBootstrapAtom.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-cpp-source/gen/android/os/StatsBootstrapAtomValue.cpp
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
    add_library(android.os.statsbootstrap_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.os.statsbootstrap_aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.os.statsbootstrap_aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.os.statsbootstrap_aidl-cpp
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
#################### android.os.statsbootstrap_aidl-cpp ####################

#################### android.os.permissions_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.permissions_aidl-cpp-source/gen/android/os/IPermissionController.cpp
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
    add_library(android.os.permissions_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.os.permissions_aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.permissions_aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.os.permissions_aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.os.permissions_aidl-cpp
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
#################### android.os.permissions_aidl-cpp ####################

#################### libincremental_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/incremental/IIncrementalServiceConnector.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/incremental/IncrementalFileSystemControlParcel.cpp
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
    add_library(libincremental_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(libincremental_aidl-cpp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/core/java
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libincremental_aidl-cpp PRIVATE
            -Wall
            -Wextra
            -Wextra-semi
            -Werror
    )

    target_link_libraries(libincremental_aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
    )
endif ()
#################### libincremental_aidl-cpp ####################

#################### android.debug_aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk-source/gen/android/debug/AdbTransportType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk-source/gen/android/debug/FingerprintAndPairDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk-source/gen/android/debug/IAdbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk-source/gen/android/debug/IAdbManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk-source/gen/android/debug/PairDevice.cpp
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
    add_library(android.debug_aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.debug_aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.debug_aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.debug_aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.debug_aidl-ndk ####################

#################### android.security.rkp_aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-analyzer-source/gen/android/security/rkp/IGetKeyCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-analyzer-source/gen/android/security/rkp/IGetRegistrationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-analyzer-source/gen/android/security/rkp/IRegistration.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-analyzer-source/gen/android/security/rkp/IRemoteProvisioning.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-analyzer-source/gen/android/security/rkp/IStoreUpgradedKeyCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-analyzer-source/gen/android/security/rkp/RemotelyProvisionedKey.cpp
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
    add_executable(android.security.rkp_aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.security.rkp_aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/core/java
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.rkp_aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.rkp_aidl-cpp-analyzer
            aidl-analyzer-main
            android.security.rkp_aidl-cpp
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
#################### android.security.rkp_aidl-cpp-analyzer ####################

#################### android.security.rkp_aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk-source/gen/android/security/rkp/IGetKeyCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk-source/gen/android/security/rkp/IGetRegistrationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk-source/gen/android/security/rkp/IRegistration.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk-source/gen/android/security/rkp/IRemoteProvisioning.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk-source/gen/android/security/rkp/IStoreUpgradedKeyCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk-source/gen/android/security/rkp/RemotelyProvisionedKey.cpp
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
    add_library(android.security.rkp_aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.security.rkp_aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.rkp_aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.rkp_aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.security.rkp_aidl-ndk ####################

#################### libdataloader_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/content/pm/DataLoaderParamsParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/content/pm/DataLoaderType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/content/pm/FileSystemControlParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/content/pm/IDataLoader.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/content/pm/IDataLoaderManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/content/pm/InstallationFileParcel.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/content/pm/InstallationFileLocation.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/content/pm/IDataLoaderStatusListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/content/pm/IPackageInstallerSessionFileSystemConnector.cpp
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
    add_library(libdataloader_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(libdataloader_aidl-cpp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/core/java
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libdataloader_aidl-cpp PRIVATE
            -Wall
            -Wextra
            -Wextra-semi
            -Werror
    )

    target_link_libraries(libdataloader_aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libincremental_aidl-cpp
            libm
            libutils
    )
endif ()
#################### libdataloader_aidl-cpp ####################

#################### android.os.statsbootstrap_aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-ndk_platform-source/gen/android/os/IStatsBootstrapAtomService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-ndk_platform-source/gen/android/os/StatsBootstrapAtom.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-ndk_platform-source/gen/android/os/StatsBootstrapAtomValue.cpp
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
    add_library(android.os.statsbootstrap_aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.os.statsbootstrap_aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.os.statsbootstrap_aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.os.statsbootstrap_aidl-ndk_platform
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
#################### android.os.statsbootstrap_aidl-ndk_platform ####################

#################### android.security.rkp_aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk_platform-source/gen/android/security/rkp/IGetKeyCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk_platform-source/gen/android/security/rkp/IGetRegistrationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk_platform-source/gen/android/security/rkp/IRegistration.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk_platform-source/gen/android/security/rkp/IRemoteProvisioning.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk_platform-source/gen/android/security/rkp/IStoreUpgradedKeyCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk_platform-source/gen/android/security/rkp/RemotelyProvisionedKey.cpp
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
    add_library(android.security.rkp_aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.security.rkp_aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.rkp_aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.rkp_aidl-ndk_platform
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
#################### android.security.rkp_aidl-ndk_platform ####################

#################### android.debug_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-source/gen/android/debug/AdbTransportType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-source/gen/android/debug/FingerprintAndPairDevice.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-source/gen/android/debug/IAdbCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-source/gen/android/debug/IAdbManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-source/gen/android/debug/PairDevice.cpp
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
    add_library(android.debug_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.debug_aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.debug_aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.debug_aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.debug_aidl-cpp
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
#################### android.debug_aidl-cpp ####################

#################### android.os.statsbootstrap_aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-cpp-analyzer-source/gen/android/os/IStatsBootstrapAtomService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-cpp-analyzer-source/gen/android/os/StatsBootstrapAtom.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-cpp-analyzer-source/gen/android/os/StatsBootstrapAtomValue.cpp
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
    add_executable(android.os.statsbootstrap_aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.os.statsbootstrap_aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/core/java
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.os.statsbootstrap_aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.os.statsbootstrap_aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.os.statsbootstrap_aidl-cpp-analyzer
            aidl-analyzer-main
            android.os.statsbootstrap_aidl-cpp
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
#################### android.os.statsbootstrap_aidl-cpp-analyzer ####################

#################### android.security.rkp_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-source/gen/android/security/rkp/IGetKeyCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-source/gen/android/security/rkp/IGetRegistrationCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-source/gen/android/security/rkp/IRegistration.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-source/gen/android/security/rkp/IRemoteProvisioning.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-source/gen/android/security/rkp/IStoreUpgradedKeyCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-source/gen/android/security/rkp/RemotelyProvisionedKey.cpp
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
    add_library(android.security.rkp_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.security.rkp_aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/android.security.rkp_aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.security.rkp_aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.security.rkp_aidl-cpp
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
#################### android.security.rkp_aidl-cpp ####################

#################### libincremental_manager_aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_manager_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/incremental/IIncrementalService.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_manager_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/incremental/IStorageLoadingProgressListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_manager_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/incremental/IncrementalNewFileParams.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_manager_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/incremental/IStorageHealthListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_manager_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/incremental/PerUidReadTimeouts.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_manager_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/incremental/StorageHealthCheckParams.cpp
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
    add_library(libincremental_manager_aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(libincremental_manager_aidl-cpp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/core/java
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_manager_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libincremental_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/core/java/libdataloader_aidl-cpp/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libincremental_manager_aidl-cpp PRIVATE
            -Wall
            -Wextra
            -Wextra-semi
            -Werror
    )

    target_link_libraries(libincremental_manager_aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdataloader_aidl-cpp
            libdl
            libincremental_aidl-cpp
            libm
            libutils
    )
endif ()
#################### libincremental_manager_aidl-cpp ####################

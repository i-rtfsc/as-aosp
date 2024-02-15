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

project(authfs_aidl_interface-cpp-analyzer)

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

#################### authfs_aidl_interface-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk-source/gen/com/android/virt/fs/AuthFsConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk-source/gen/com/android/virt/fs/IAuthFs.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk-source/gen/com/android/virt/fs/IAuthFsService.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk-source/gen/com/android/virt/fs/IVirtFdService.cpp
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
    add_library(authfs_aidl_interface-ndk SHARED ${SOURCE_FILES})

    target_include_directories(authfs_aidl_interface-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(authfs_aidl_interface-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(authfs_aidl_interface-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### authfs_aidl_interface-ndk ####################

#################### authfs_aidl_interface-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-source/gen/com/android/virt/fs/AuthFsConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-source/gen/com/android/virt/fs/IAuthFs.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-source/gen/com/android/virt/fs/IAuthFsService.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-source/gen/com/android/virt/fs/IVirtFdService.cpp
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
    add_library(authfs_aidl_interface-cpp SHARED ${SOURCE_FILES})

    target_include_directories(authfs_aidl_interface-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(authfs_aidl_interface-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(authfs_aidl_interface-cpp
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
#################### authfs_aidl_interface-cpp ####################

#################### authfs_aidl_interface-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-analyzer-source/gen/com/android/virt/fs/AuthFsConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-analyzer-source/gen/com/android/virt/fs/IAuthFs.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-analyzer-source/gen/com/android/virt/fs/IAuthFsService.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-analyzer-source/gen/com/android/virt/fs/IVirtFdService.cpp
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
    add_executable(authfs_aidl_interface-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(authfs_aidl_interface-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Virtualization/authfs/aidl
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(authfs_aidl_interface-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(authfs_aidl_interface-cpp-analyzer
            aidl-analyzer-main
            authfs_aidl_interface-cpp
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
#################### authfs_aidl_interface-cpp-analyzer ####################

#################### authfs_aidl_interface-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk_platform-source/gen/com/android/virt/fs/AuthFsConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk_platform-source/gen/com/android/virt/fs/IAuthFs.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk_platform-source/gen/com/android/virt/fs/IAuthFsService.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk_platform-source/gen/com/android/virt/fs/IVirtFdService.cpp
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
    add_library(authfs_aidl_interface-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(authfs_aidl_interface-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/authfs/aidl/authfs_aidl_interface-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(authfs_aidl_interface-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(authfs_aidl_interface-ndk_platform
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
#################### authfs_aidl_interface-ndk_platform ####################

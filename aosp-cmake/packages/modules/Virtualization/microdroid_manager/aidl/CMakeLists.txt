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

project(android.system.virtualization.payload-cpp-analyzer)

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

#################### android.system.virtualization.payload-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl/android.system.virtualization.payload-cpp-analyzer-source/gen/android/system/virtualization/payload/IVmPayloadService.cpp
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
    add_executable(android.system.virtualization.payload-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.system.virtualization.payload-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl/android.system.virtualization.payload-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl/android.system.virtualization.payload-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualization.payload-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualization.payload-cpp-analyzer
            aidl-analyzer-main
            android.system.virtualization.payload-cpp
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
#################### android.system.virtualization.payload-cpp-analyzer ####################

#################### android.system.virtualization.payload-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl/android.system.virtualization.payload-ndk_platform-source/gen/android/system/virtualization/payload/IVmPayloadService.cpp
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
    add_library(android.system.virtualization.payload-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualization.payload-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl/android.system.virtualization.payload-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualization.payload-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualization.payload-ndk_platform
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
#################### android.system.virtualization.payload-ndk_platform ####################

#################### android.system.virtualization.payload-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl/android.system.virtualization.payload-cpp-source/gen/android/system/virtualization/payload/IVmPayloadService.cpp
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
    add_library(android.system.virtualization.payload-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualization.payload-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl/android.system.virtualization.payload-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualization.payload-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualization.payload-cpp
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
#################### android.system.virtualization.payload-cpp ####################

#################### android.system.virtualization.payload-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl/android.system.virtualization.payload-ndk-source/gen/android/system/virtualization/payload/IVmPayloadService.cpp
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
    add_library(android.system.virtualization.payload-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.virtualization.payload-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/microdroid_manager/aidl/android.system.virtualization.payload-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.virtualization.payload-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.virtualization.payload-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.virtualization.payload-ndk ####################

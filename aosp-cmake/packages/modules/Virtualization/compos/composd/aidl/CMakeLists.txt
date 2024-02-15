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

project(android.system.composd-cpp-analyzer)

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

#################### android.system.composd-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-ndk-source/gen/android/system/composd/ICompilationTask.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-ndk-source/gen/android/system/composd/ICompilationTaskCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-ndk-source/gen/android/system/composd/IIsolatedCompilationService.cpp
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
    add_library(android.system.composd-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.system.composd-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.composd-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.composd-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### android.system.composd-ndk ####################

#################### android.system.composd-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-cpp-source/gen/android/system/composd/ICompilationTask.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-cpp-source/gen/android/system/composd/ICompilationTaskCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-cpp-source/gen/android/system/composd/IIsolatedCompilationService.cpp
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
    add_library(android.system.composd-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.system.composd-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.composd-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.composd-cpp
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
#################### android.system.composd-cpp ####################

#################### android.system.composd-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-cpp-analyzer-source/gen/android/system/composd/ICompilationTask.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-cpp-analyzer-source/gen/android/system/composd/ICompilationTaskCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-cpp-analyzer-source/gen/android/system/composd/IIsolatedCompilationService.cpp
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
    add_executable(android.system.composd-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.system.composd-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Virtualization/compos/composd/aidl
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.composd-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.composd-cpp-analyzer
            aidl-analyzer-main
            android.system.composd-cpp
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
#################### android.system.composd-cpp-analyzer ####################

#################### android.system.composd-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-ndk_platform-source/gen/android/system/composd/ICompilationTask.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-ndk_platform-source/gen/android/system/composd/ICompilationTaskCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-ndk_platform-source/gen/android/system/composd/IIsolatedCompilationService.cpp
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
    add_library(android.system.composd-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.system.composd-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Virtualization/compos/composd/aidl/android.system.composd-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.system.composd-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.system.composd-ndk_platform
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
#################### android.system.composd-ndk_platform ####################

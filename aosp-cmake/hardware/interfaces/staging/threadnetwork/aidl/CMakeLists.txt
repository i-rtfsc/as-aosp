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

project(android.hardware.threadnetwork-cpp-analyzer)

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

#################### android.hardware.threadnetwork-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-ndk-source/gen/android/hardware/threadnetwork/IThreadChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-ndk-source/gen/android/hardware/threadnetwork/IThreadChipCallback.cpp
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
    add_library(android.hardware.threadnetwork-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.threadnetwork-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.threadnetwork-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.threadnetwork-ndk
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
#################### android.hardware.threadnetwork-ndk ####################

#################### android.hardware.threadnetwork-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-ndk_platform-source/gen/android/hardware/threadnetwork/IThreadChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-ndk_platform-source/gen/android/hardware/threadnetwork/IThreadChipCallback.cpp
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
    add_library(android.hardware.threadnetwork-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.threadnetwork-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.threadnetwork-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.threadnetwork-ndk_platform
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
#################### android.hardware.threadnetwork-ndk_platform ####################

#################### android.hardware.threadnetwork-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-cpp-analyzer-source/gen/android/hardware/threadnetwork/IThreadChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-cpp-analyzer-source/gen/android/hardware/threadnetwork/IThreadChipCallback.cpp
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
    add_executable(android.hardware.threadnetwork-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(android.hardware.threadnetwork-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/staging/threadnetwork/aidl
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.threadnetwork-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.threadnetwork-cpp-analyzer
            aidl-analyzer-main
            android.hardware.threadnetwork-cpp
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.threadnetwork-cpp-analyzer ####################

#################### android.hardware.threadnetwork-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-cpp-source/gen/android/hardware/threadnetwork/IThreadChip.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-cpp-source/gen/android/hardware/threadnetwork/IThreadChipCallback.cpp
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
    add_library(android.hardware.threadnetwork-cpp SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.threadnetwork-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/threadnetwork/aidl/android.hardware.threadnetwork-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.threadnetwork-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.threadnetwork-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libutils
    )
endif ()
#################### android.hardware.threadnetwork-cpp ####################

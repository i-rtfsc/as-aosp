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

project(android.hardware.media.bufferpool2-V2-ndk_platform)

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

#################### android.hardware.media.bufferpool2-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/android/hardware/media/bufferpool2/Buffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/android/hardware/media/bufferpool2/BufferInvalidationMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/android/hardware/media/bufferpool2/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/android/hardware/media/bufferpool2/BufferStatusMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/android/hardware/media/bufferpool2/IAccessor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/android/hardware/media/bufferpool2/IClientManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/android/hardware/media/bufferpool2/IConnection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/android/hardware/media/bufferpool2/IObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/android/hardware/media/bufferpool2/ResultStatus.cpp
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
    add_library(android.hardware.media.bufferpool2-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.media.bufferpool2-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.media.bufferpool2-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.media.bufferpool2-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
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
#################### android.hardware.media.bufferpool2-V2-ndk_platform ####################

#################### android.hardware.media.bufferpool2-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/android/hardware/media/bufferpool2/Buffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/android/hardware/media/bufferpool2/BufferInvalidationMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/android/hardware/media/bufferpool2/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/android/hardware/media/bufferpool2/BufferStatusMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/android/hardware/media/bufferpool2/IAccessor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/android/hardware/media/bufferpool2/IClientManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/android/hardware/media/bufferpool2/IConnection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/android/hardware/media/bufferpool2/IObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/android/hardware/media/bufferpool2/ResultStatus.cpp
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
    add_library(android.hardware.media.bufferpool2-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.media.bufferpool2-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.media.bufferpool2-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.media.bufferpool2-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
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
#################### android.hardware.media.bufferpool2-V1-ndk ####################

#################### android.hardware.media.bufferpool2-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/android/hardware/media/bufferpool2/Buffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/android/hardware/media/bufferpool2/BufferInvalidationMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/android/hardware/media/bufferpool2/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/android/hardware/media/bufferpool2/BufferStatusMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/android/hardware/media/bufferpool2/IAccessor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/android/hardware/media/bufferpool2/IClientManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/android/hardware/media/bufferpool2/IConnection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/android/hardware/media/bufferpool2/IObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/android/hardware/media/bufferpool2/ResultStatus.cpp
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
    add_library(android.hardware.media.bufferpool2-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.media.bufferpool2-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.media.bufferpool2-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.media.bufferpool2-V2-ndk
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
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
#################### android.hardware.media.bufferpool2-V2-ndk ####################

#################### android.hardware.media.bufferpool2-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/android/hardware/media/bufferpool2/Buffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/android/hardware/media/bufferpool2/BufferInvalidationMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/android/hardware/media/bufferpool2/BufferStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/android/hardware/media/bufferpool2/BufferStatusMessage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/android/hardware/media/bufferpool2/IAccessor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/android/hardware/media/bufferpool2/IClientManager.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/android/hardware/media/bufferpool2/IConnection.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/android/hardware/media/bufferpool2/IObserver.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/android/hardware/media/bufferpool2/ResultStatus.cpp
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
    add_library(android.hardware.media.bufferpool2-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.media.bufferpool2-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/media/bufferpool/aidl/android.hardware.media.bufferpool2-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.media.bufferpool2-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.media.bufferpool2-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.common.fmq-V1-ndk_platform
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
#################### android.hardware.media.bufferpool2-V1-ndk_platform ####################

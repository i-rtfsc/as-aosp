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

project(android.hardware.media.c2-ndk_platform)

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

#################### android.hardware.media.c2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/BaseBlock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/Block.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/Buffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/FieldDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/FieldId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/FieldSupportedValues.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/FieldSupportedValuesQuery.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/FieldSupportedValuesQueryResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/FrameData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/IComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/IComponentInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/IComponentListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/IComponentStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/IConfigurable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/InfoBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/ParamDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/ParamField.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/ParamFieldValues.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/Params.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/SettingResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/StructDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/SurfaceSyncObj.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/ValueRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/Work.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/WorkBundle.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/WorkOrdinal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/android/hardware/media/c2/Worklet.cpp
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
    add_library(android.hardware.media.c2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.media.c2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.media.c2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.media.c2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.media.bufferpool2-V1-ndk_platform
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnativewindow
    )
endif ()
#################### android.hardware.media.c2-ndk_platform ####################

#################### android.hardware.media.c2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/BaseBlock.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/Block.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/Buffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/FieldDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/FieldId.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/FieldSupportedValues.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/FieldSupportedValuesQuery.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/FieldSupportedValuesQueryResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/FrameData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/IComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/IComponentInterface.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/IComponentListener.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/IComponentStore.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/IConfigurable.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/InfoBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/ParamDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/ParamField.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/ParamFieldValues.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/Params.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/SettingResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/Status.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/StructDescriptor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/SurfaceSyncObj.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/ValueRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/Work.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/WorkBundle.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/WorkOrdinal.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/android/hardware/media/c2/Worklet.cpp
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
    add_library(android.hardware.media.c2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.media.c2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/staging/c2/aidl/android.hardware.media.c2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.media.c2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.media.c2-ndk
            android.hardware.common-V2-ndk
            android.hardware.media.bufferpool2-V1-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libnativewindow
    )
endif ()
#################### android.hardware.media.c2-ndk ####################

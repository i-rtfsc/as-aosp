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

project(android.hardware.graphics.common-V5-ndk_platform)

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

#################### android.hardware.graphics.common-V5-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/AlphaInterpretation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/ColorTransform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/DisplayDecorationSupport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/FRect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/Hdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/HdrConversionCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/HdrConversionStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/Point.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/Transform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V5-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V5-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V5-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V5-ndk_platform
            android.hardware.common-V2-ndk_platform
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
#################### android.hardware.graphics.common-V5-ndk_platform ####################

#################### android.hardware.graphics.common-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
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
#################### android.hardware.graphics.common-V2-ndk_platform ####################

#################### android.hardware.graphics.common-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/AlphaInterpretation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/ColorTransform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/DisplayDecorationSupport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/FRect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/Hdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/Point.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/Transform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V3-ndk_platform
            android.hardware.common-V2-ndk_platform
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
#################### android.hardware.graphics.common-V3-ndk_platform ####################

#################### android.hardware.graphics.common-V4-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/AlphaInterpretation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/ColorTransform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/DisplayDecorationSupport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/FRect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/Hdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/HdrConversionCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/HdrConversionStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/Point.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/Transform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V4-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V4-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V4-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V4-ndk_platform
            android.hardware.common-V2-ndk_platform
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
#################### android.hardware.graphics.common-V4-ndk_platform ####################

#################### android.hardware.graphics.common-V5-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/AlphaInterpretation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/ColorTransform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/DisplayDecorationSupport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/FRect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/Hdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/HdrConversionCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/HdrConversionStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/Point.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/Transform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V5-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V5-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V5-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V5-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V5-ndk
            android.hardware.common-V2-ndk
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
#################### android.hardware.graphics.common-V5-ndk ####################

#################### android.hardware.graphics.common-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
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
#################### android.hardware.graphics.common-V1-ndk_platform ####################

#################### android.hardware.graphics.common-V4-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/AlphaInterpretation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/ColorTransform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/DisplayDecorationSupport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/FRect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/Hdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/HdrConversionCapability.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/HdrConversionStrategy.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/Point.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/Transform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V4-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V4-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V4-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V4-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V4-ndk
            android.hardware.common-V2-ndk
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
#################### android.hardware.graphics.common-V4-ndk ####################

#################### android.hardware.graphics.common-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V1-ndk
            android.hardware.common-V2-ndk
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
#################### android.hardware.graphics.common-V1-ndk ####################

#################### android.hardware.graphics.common-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V2-ndk
            android.hardware.common-V2-ndk
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
#################### android.hardware.graphics.common-V2-ndk ####################

#################### android.hardware.graphics.common-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/AlphaInterpretation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/BlendMode.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/BufferUsage.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/ChromaSiting.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/ColorTransform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/Compression.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/Cta861_3.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/Dataspace.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/DisplayDecorationSupport.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/ExtendableType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/FRect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/HardwareBuffer.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/HardwareBufferDescription.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/Hdr.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/Interlaced.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/PixelFormat.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/PlaneLayout.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponent.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/PlaneLayoutComponentType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/Point.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/Rect.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/Smpte2086.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/StandardMetadataType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/Transform.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/android/hardware/graphics/common/XyColor.cpp
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
    add_library(android.hardware.graphics.common-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.graphics.common-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/graphics/common/aidl/android.hardware.graphics.common-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.graphics.common-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.graphics.common-V3-ndk
            android.hardware.common-V2-ndk
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
#################### android.hardware.graphics.common-V3-ndk ####################

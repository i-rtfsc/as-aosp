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

project(android.hardware.automotive.evs-V3-ndk_platform)

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

#################### android.hardware.automotive.evs-V3-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/CameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/CameraParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/ColorChannel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/DeviceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/DeviceStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/DisplayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/DisplayState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/EmbeddedData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/EvsEventDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/EvsEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/EvsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/ExposureParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/GridStatisticDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/GridStatisticType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/GridStatistics.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/Histogram.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsCamera.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsCameraStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsEnumerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsEnumeratorStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArray.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArrayStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/ParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/Rotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/RotationQuaternion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/SensorPose.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/Size.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/StreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/Translation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/UltrasonicSensor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/UltrasonicsArrayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/android/hardware/automotive/evs/UltrasonicsDataFrameDesc.cpp
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
    add_library(android.hardware.automotive.evs-V3-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.evs-V3-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.evs-V3-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.evs-V3-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### android.hardware.automotive.evs-V3-ndk_platform ####################

#################### android.hardware.automotive.evs-V2-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/CameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/CameraParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/ColorChannel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/DeviceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/DeviceStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/DisplayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/DisplayState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/EmbeddedData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/EvsEventDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/EvsEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/EvsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/ExposureParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/GridStatisticDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/GridStatisticType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/GridStatistics.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/Histogram.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsCamera.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsCameraStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsEnumerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsEnumeratorStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArray.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArrayStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/ParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/Rotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/RotationQuaternion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/SensorPose.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/Size.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/StreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/Translation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/UltrasonicSensor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/UltrasonicsArrayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/android/hardware/automotive/evs/UltrasonicsDataFrameDesc.cpp
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
    add_library(android.hardware.automotive.evs-V2-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.evs-V2-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.evs-V2-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.evs-V2-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### android.hardware.automotive.evs-V2-ndk_platform ####################

#################### android.hardware.automotive.evs-V1-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/CameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/CameraParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/DeviceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/DeviceStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/DisplayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/DisplayState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/EvsEventDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/EvsEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/EvsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/IEvsCamera.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/IEvsCameraStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/IEvsDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/IEvsEnumerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/IEvsEnumeratorStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArray.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArrayStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/ParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/Rotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/RotationQuaternion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/SensorPose.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/StreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/Translation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/UltrasonicSensor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/UltrasonicsArrayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/android/hardware/automotive/evs/UltrasonicsDataFrameDesc.cpp
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
    add_library(android.hardware.automotive.evs-V1-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.evs-V1-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.evs-V1-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.evs-V1-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
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
#################### android.hardware.automotive.evs-V1-ndk ####################

#################### android.hardware.automotive.evs-V3-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/CameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/CameraParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/ColorChannel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/DeviceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/DeviceStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/DisplayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/DisplayState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/EmbeddedData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/EvsEventDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/EvsEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/EvsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/ExposureParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/GridStatisticDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/GridStatisticType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/GridStatistics.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/Histogram.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/IEvsCamera.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/IEvsCameraStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/IEvsDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/IEvsEnumerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/IEvsEnumeratorStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArray.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArrayStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/ParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/Rotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/RotationQuaternion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/SensorPose.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/Size.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/StreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/Translation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/UltrasonicSensor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/UltrasonicsArrayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/android/hardware/automotive/evs/UltrasonicsDataFrameDesc.cpp
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
    add_library(android.hardware.automotive.evs-V3-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.evs-V3-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V3-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.evs-V3-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.evs-V3-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
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
#################### android.hardware.automotive.evs-V3-ndk ####################

#################### android.hardware.automotive.evs-V2-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/CameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/CameraParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/ColorChannel.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/DeviceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/DeviceStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/DisplayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/DisplayState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/EmbeddedData.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/EvsEventDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/EvsEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/EvsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/ExposureParameters.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/GridStatisticDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/GridStatisticType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/GridStatistics.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/Histogram.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/IEvsCamera.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/IEvsCameraStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/IEvsDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/IEvsEnumerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/IEvsEnumeratorStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArray.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArrayStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/ParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/Rotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/RotationQuaternion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/SensorPose.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/Size.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/StreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/Translation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/UltrasonicSensor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/UltrasonicsArrayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/android/hardware/automotive/evs/UltrasonicsDataFrameDesc.cpp
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
    add_library(android.hardware.automotive.evs-V2-ndk SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.evs-V2-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V2-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.evs-V2-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.evs-V2-ndk
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
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
#################### android.hardware.automotive.evs-V2-ndk ####################

#################### android.hardware.automotive.evs-V1-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/BufferDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/CameraDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/CameraParam.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/DeviceStatus.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/DeviceStatusType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/DisplayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/DisplayState.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/EvsEventDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/EvsEventType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/EvsResult.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsCamera.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsCameraStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsDisplay.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsEnumerator.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsEnumeratorStatusCallback.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArray.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/IEvsUltrasonicsArrayStream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/ParameterRange.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/Rotation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/RotationQuaternion.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/SensorPose.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/Stream.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/StreamType.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/Translation.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/UltrasonicSensor.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/UltrasonicsArrayDesc.cpp
        ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/android/hardware/automotive/evs/UltrasonicsDataFrameDesc.cpp
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
    add_library(android.hardware.automotive.evs-V1-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.automotive.evs-V1-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/evs/aidl/android.hardware.automotive.evs-V1-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.automotive.evs-V1-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(android.hardware.automotive.evs-V1-ndk_platform
            android.hardware.common-V2-ndk_platform
            android.hardware.graphics.common-V4-ndk_platform
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
#################### android.hardware.automotive.evs-V1-ndk_platform ####################

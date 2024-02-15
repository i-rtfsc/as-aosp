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

project(SensorPoseProvider-example)

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

#################### libheadtracking-binding ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/SensorPoseProvider.cpp
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
    add_library(libheadtracking-binding SHARED ${SOURCE_FILES})

    target_include_directories(libheadtracking-binding
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libheadtracking
            ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/include
    )

    # Add c/c++ flags
    target_compile_options(libheadtracking-binding PRIVATE
            -Werror
    )

    target_link_libraries(libheadtracking-binding
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libheadtracking
            liblog
            libm
            libsensor
            libutils
    )
endif ()
#################### libheadtracking-binding ####################

#################### libheadtracking ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/HeadTrackingProcessor.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/ModeSelector.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/Pose.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/PoseBias.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/PoseDriftCompensator.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/PosePredictor.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/PoseRateLimiter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/QuaternionUtil.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/ScreenHeadFusion.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/StillnessDetector.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/Twist.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/VectorRecorder.cpp
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
    add_library(libheadtracking SHARED ${SOURCE_FILES})

    target_include_directories(libheadtracking
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/include
            ${ANDROID_ROOT}/frameworks/av/media/libheadtracking
    )

    # Add c/c++ flags
    target_compile_options(libheadtracking PRIVATE
            -Werror
    )

    target_link_libraries(libheadtracking
            libaudioutils
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libeigen
            libm
    )
endif ()
#################### libheadtracking ####################

#################### SensorPoseProvider-example ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/SensorPoseProvider-example.cpp
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
    add_executable(SensorPoseProvider-example ${SOURCE_FILES})

    target_include_directories(SensorPoseProvider-example
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libheadtracking
            ${ANDROID_ROOT}/frameworks/av/media/libheadtracking/include
    )

    # Add c/c++ flags
    target_compile_options(SensorPoseProvider-example PRIVATE
            -Werror
    )

    target_link_libraries(SensorPoseProvider-example
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libheadtracking
            libheadtracking-binding
            libm
            libsensor
            libutils
    )
endif ()
#################### SensorPoseProvider-example ####################

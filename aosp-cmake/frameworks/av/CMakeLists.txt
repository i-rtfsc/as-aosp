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

project(av-types-aidl-cpp)

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

#################### av-types-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/android/media/InterpolatorConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/android/media/InterpolatorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/android/media/MicrophoneInfoFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/android/media/VolumeShaperConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/android/media/VolumeShaperConfigurationOptionFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/android/media/VolumeShaperConfigurationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/android/media/VolumeShaperOperation.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/android/media/VolumeShaperOperationFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/android/media/VolumeShaperState.cpp
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
    add_library(av-types-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(av-types-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(av-types-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(av-types-aidl-cpp
            android.media.audio.common.types-V2-cpp
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
#################### av-types-aidl-cpp ####################

#################### av-types-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/android/media/InterpolatorConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/android/media/InterpolatorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/android/media/MicrophoneInfoFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/android/media/VolumeShaperConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/android/media/VolumeShaperConfigurationOptionFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/android/media/VolumeShaperConfigurationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/android/media/VolumeShaperOperation.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/android/media/VolumeShaperOperationFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/android/media/VolumeShaperState.cpp
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
    add_library(av-types-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(av-types-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(av-types-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(av-types-aidl-ndk
            android.media.audio.common.types-V2-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### av-types-aidl-ndk ####################

#################### av-types-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/android/media/InterpolatorConfig.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/android/media/InterpolatorType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/android/media/MicrophoneInfoFw.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/android/media/VolumeShaperConfiguration.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/android/media/VolumeShaperConfigurationOptionFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/android/media/VolumeShaperConfigurationType.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/android/media/VolumeShaperOperation.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/android/media/VolumeShaperOperationFlag.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/android/media/VolumeShaperState.cpp
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
    add_library(av-types-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(av-types-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/av-types-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(av-types-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(av-types-aidl-ndk_platform
            android.media.audio.common.types-V2-ndk_platform
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
#################### av-types-aidl-ndk_platform ####################

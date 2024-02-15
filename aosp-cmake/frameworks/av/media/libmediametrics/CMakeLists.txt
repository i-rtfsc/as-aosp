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

project(mediametricsservice-aidl-cpp-analyzer)

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

#################### mediametricsservice-aidl-cpp ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-cpp-source/gen/android/media/IMediaMetricsService.cpp
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
    add_library(mediametricsservice-aidl-cpp SHARED ${SOURCE_FILES})

    target_include_directories(mediametricsservice-aidl-cpp
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mediametricsservice-aidl-cpp PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mediametricsservice-aidl-cpp
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
#################### mediametricsservice-aidl-cpp ####################

#################### mediametricsservice-aidl-ndk ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-ndk-source/gen/android/media/IMediaMetricsService.cpp
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
    add_library(mediametricsservice-aidl-ndk SHARED ${SOURCE_FILES})

    target_include_directories(mediametricsservice-aidl-ndk
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-ndk-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mediametricsservice-aidl-ndk PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mediametricsservice-aidl-ndk
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### mediametricsservice-aidl-ndk ####################

#################### libmediametrics ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libmediametrics/MediaMetricsItem.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libmediametrics/MediaMetrics.cpp
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
    add_library(libmediametrics SHARED ${SOURCE_FILES})

    target_include_directories(libmediametrics
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmediametrics/include
            ${ANDROID_ROOT}/frameworks/av/media/libmediametrics
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-cpp-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(libmediametrics PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libmediametrics
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils
            mediametricsservice-aidl-cpp
    )
endif ()
#################### libmediametrics ####################

#################### mediametricsservice-aidl-ndk_platform ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-ndk_platform-source/gen/android/media/IMediaMetricsService.cpp
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
    add_library(mediametricsservice-aidl-ndk_platform SHARED ${SOURCE_FILES})

    target_include_directories(mediametricsservice-aidl-ndk_platform
            PUBLIC
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-ndk_platform-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mediametricsservice-aidl-ndk_platform PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mediametricsservice-aidl-ndk_platform
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
#################### mediametricsservice-aidl-ndk_platform ####################

#################### mediametricsservice-aidl-cpp-analyzer ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-cpp-analyzer-source/gen/android/media/IMediaMetricsService.cpp
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
    add_executable(mediametricsservice-aidl-cpp-analyzer ${SOURCE_FILES})

    target_include_directories(mediametricsservice-aidl-cpp-analyzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libmediametrics
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-cpp-source/gen/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/media/libmediametrics/mediametricsservice-aidl-cpp-analyzer-source/gen/include
    )

    # Add c/c++ flags
    target_compile_options(mediametricsservice-aidl-cpp-analyzer PRIVATE
            -Wextra
            -Wall
            -Werror
            -Wextra-semi
    )

    target_link_libraries(mediametricsservice-aidl-cpp-analyzer
            aidl-analyzer-main
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libutils
            mediametricsservice-aidl-cpp
    )
endif ()
#################### mediametricsservice-aidl-cpp-analyzer ####################

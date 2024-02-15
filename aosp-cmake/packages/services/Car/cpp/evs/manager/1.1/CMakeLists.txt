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

project(android.automotive.evs.manager@1.1)

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

#################### libevsmanager_static ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/Enumerator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/EnumeratorProxy.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/HalCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/HalDisplay.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/VirtualCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/emul/EvsEmulatedCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/emul/VideoCapture.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/stats/CameraUsageStats.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/stats/LooperWrapper.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/stats/StatsCollector.cpp
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
    add_library(libevsmanager_static SHARED ${SOURCE_FILES})

    target_include_directories(libevsmanager_static
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/stats
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/1.1/statslog_evs.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libevsmanager_static PRIVATE
            -Wall
            -Werror
            -DEVS_DEBUG
    )

    target_link_libraries(libevsmanager_static
            android.hardware.automotive.evs-1.0
            android.hardware.automotive.evs-1.1
            libbase
            libc
            libc++_static
            libc++demangle
            libc++fs
            libcamera_metadata
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            libm
            libpng
            libprocessgroup
            libstatslog_evs
            libui
            libutils
    )
endif ()
#################### libevsmanager_static ####################

#################### android.automotive.evs.manager-1.1 ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/Enumerator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/EnumeratorProxy.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/HalCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/HalDisplay.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/VirtualCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/emul/EvsEmulatedCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/emul/VideoCapture.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/stats/CameraUsageStats.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/stats/LooperWrapper.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/stats/StatsCollector.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1/service.cpp
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
    add_executable(android.automotive.evs.manager-1.1 ${SOURCE_FILES})

    target_include_directories(android.automotive.evs.manager-1.1
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/1.1/statslog_evs.h/gen
    )

    # Add c/c++ flags
    target_compile_options(android.automotive.evs.manager-1.1 PRIVATE
            -Wall
            -Werror
            -DEVS_DEBUG
    )

    target_link_libraries(android.automotive.evs.manager-1.1
            android.hardware.automotive.evs-1.0
            android.hardware.automotive.evs-1.1
            libbase
            libc
            libc++
            libc++demangle
            libc++fs
            libcamera_metadata
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            libm
            libpng
            libprocessgroup
            libstatslog_evs
            libui
            libutils
    )
endif ()
#################### android.automotive.evs.manager-1.1 ####################

#################### libstatslog_evs ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/1.1/statslog_evs.cpp/gen/statslog_evs.cpp
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
    add_library(libstatslog_evs SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_evs
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/1.1
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/1.1/statslog_evs.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/1.1/statslog_evs.h/gen
    )

    target_link_libraries(libstatslog_evs
            android.os.statsbootstrap_aidl-cpp
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libstatsbootstrap
            libutils
    )
endif ()
#################### libstatslog_evs ####################

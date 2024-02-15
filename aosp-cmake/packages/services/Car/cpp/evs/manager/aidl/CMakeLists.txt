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

project(evs_enumerator_fuzzer)

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

#################### evsmanagerd ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/service.cpp
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
    add_executable(evsmanagerd ${SOURCE_FILES})

    target_include_directories(evsmanagerd
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/stats/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/aidl/statslog_evsmanagerd.h/gen
    )

    # Add c/c++ flags
    target_compile_options(evsmanagerd PRIVATE
            -Wall
            -Werror
            -DEVS_DEBUG
    )

    target_link_libraries(evsmanagerd
            android.hardware.automotive.evs-V2-ndk
            android.hardware.automotive.evs-1.0
            android.hardware.automotive.evs-1.1
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libaidlcommonsupport
            libarect_headers
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libc++fs
            libcamera_metadata
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libevsmanagerd
            libhardware
            libhidlbase
            liblog
            libm
            libprocessgroup
            libstatslog_evsmanagerd
            libui_headers
            libutils
    )
endif ()
#################### evsmanagerd ####################

#################### libevsmanagerd ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/HidlCameraStream.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/AidlDisplay.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/AidlCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/HidlEnumerator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/AidlCameraStream.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/AidlEnumerator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/HidlDisplay.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/HidlCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/stats/src/CameraUsageStats.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/stats/src/StatsCollector.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/stats/src/LooperWrapper.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/utils/src/Utils.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/HalDisplay.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/VirtualCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/Enumerator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/HalCamera.cpp
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
    add_library(libevsmanagerd SHARED ${SOURCE_FILES})

    target_include_directories(libevsmanagerd
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/stats/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/include
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/aidl/statslog_evsmanagerd.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libevsmanagerd PRIVATE
            -Wall
            -Werror
            -DEVS_DEBUG
    )

    target_link_libraries(libevsmanagerd
            android.hardware.automotive.evs-V2-ndk
            android.hardware.automotive.evs-1.0
            android.hardware.automotive.evs-1.1
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libaidlcommonsupport
            libarect_headers
            libbase
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libc++fs
            libcamera_metadata
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libprocessgroup
            libstatslog_evsmanagerd
            libui_headers
            libutils
    )
endif ()
#################### libevsmanagerd ####################

#################### evs_enumerator_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/HidlCameraStream.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/AidlDisplay.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/AidlCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/HidlEnumerator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/AidlCameraStream.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/AidlEnumerator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/HidlDisplay.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/src/HidlCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/stats/src/CameraUsageStats.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/stats/src/StatsCollector.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/stats/src/LooperWrapper.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/utils/src/Utils.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/HalDisplay.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/VirtualCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/Enumerator.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/HalCamera.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/src/EnumeratorFuzzer.cpp
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
    add_executable(evs_enumerator_fuzzer ${SOURCE_FILES})

    target_include_directories(evs_enumerator_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/stats/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/wrappers/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/aidl/statslog_evsmanagerd.h/gen
    )

    # Add c/c++ flags
    target_compile_options(evs_enumerator_fuzzer PRIVATE
            -Wall
            -Werror
            -DEVS_DEBUG
    )

    target_link_libraries(evs_enumerator_fuzzer
            android.hardware.automotive.evs-V2-ndk
            android.hardware.automotive.evs-1.0
            android.hardware.automotive.evs-1.1
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libaidlcommonsupport
            libarect_headers
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libc++fs
            libcamera_metadata
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libprocessgroup
            libstatslog_evsmanagerd
            libui_headers
            libutils
    )
endif ()
#################### evs_enumerator_fuzzer ####################

#################### libstatslog_evsmanagerd ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/aidl/statslog_evsmanagerd.cpp/gen/statslog_evsmanagerd.cpp
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
    add_library(libstatslog_evsmanagerd SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_evsmanagerd
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/aidl/statslog_evsmanagerd.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/services/Car/cpp/evs/manager/aidl/statslog_evsmanagerd.h/gen
    )

    target_link_libraries(libstatslog_evsmanagerd
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
#################### libstatslog_evsmanagerd ####################

#################### libmockevshal ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/tests/src/MockEvsHal.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/tests/src/MockHidlEvsHal.cpp
        ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/tests/src/MockHidlEvsHal_1_0.cpp
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
    add_library(libmockevshal SHARED ${SOURCE_FILES})

    target_include_directories(libmockevshal
            PUBLIC
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/tests/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl/include
            ${ANDROID_ROOT}/packages/services/Car/cpp/evs/manager/aidl
    )

    # Add c/c++ flags
    target_compile_options(libmockevshal PRIVATE
            -Werror
    )

    target_link_libraries(libmockevshal
            android.hardware.automotive.evs-V2-ndk
            android.hardware.automotive.evs-1.0
            android.hardware.automotive.evs-1.1
            android.hardware.common-V2-ndk
            android.hardware.graphics.common-V4-ndk
            libaidlcommonsupport
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libcamera_client
            libcamera_metadata
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgmock
            libgtest
            libgtest_main
            libhidlbase
            libm
            libnativewindow
            libui
            libutils
    )
endif ()
#################### libmockevshal ####################

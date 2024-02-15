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

project(storaged)

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

#################### libstoraged ####################
set(FILES
        ${ANDROID_ROOT}/system/core/storaged/storaged.cpp
        ${ANDROID_ROOT}/system/core/storaged/storaged_diskstats.cpp
        ${ANDROID_ROOT}/system/core/storaged/storaged_info.cpp
        ${ANDROID_ROOT}/system/core/storaged/storaged_service.cpp
        ${ANDROID_ROOT}/system/core/storaged/storaged_utils.cpp
        ${ANDROID_ROOT}/system/core/storaged/storaged_uid_monitor.cpp
        ${ANDROID_ROOT}/system/core/storaged/uid_info.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/core/storaged/libstoraged/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/storaged/storaged.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/system/core/storaged/libstoraged/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/IStoraged.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/core/storaged/libstoraged/android_${OUT_ARCH_CPU}_static/gen/aidl/android/os/storaged/IStoragedPrivate.cpp
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
    add_library(libstoraged SHARED ${SOURCE_FILES})

    target_include_directories(libstoraged
            PUBLIC
            ${ANDROID_ROOT}/system/core/storaged/include
            ${ANDROID_ROOT}/system/core/storaged
            ${OUT_INTERMEDIATES_ROOT}/system/core/storaged/libstoraged/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/storaged
            ${OUT_INTERMEDIATES_ROOT}/system/core/storaged/libstoraged/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/core/storaged/libstoraged/android_${OUT_ARCH_CPU}_static/gen/aidl
    )

    # Add c/c++ flags
    target_compile_options(libstoraged PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libstoraged
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            android.hardware.health-1.0
            android.hardware.health-2.0
            libbase
            libbatteryservice_headers
            libbinder
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealthhalutils
            libhealthshim
            libhidlbase
            liblog
            libm
            libprotobuf-cpp-lite
            libutils
            libz
            packagemanager_aidl-cpp
    )
endif ()
#################### libstoraged ####################

#################### storaged ####################
set(FILES
        ${ANDROID_ROOT}/system/core/storaged/main.cpp
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
    add_executable(storaged ${SOURCE_FILES})

    target_include_directories(storaged
            PUBLIC
            ${ANDROID_ROOT}/system/core/storaged
            ${ANDROID_ROOT}/system/core/storaged/include
            ${OUT_INTERMEDIATES_ROOT}/system/core/storaged/libstoraged/android_${OUT_ARCH_CPU}_static/gen/aidl
            ${OUT_INTERMEDIATES_ROOT}/system/core/storaged/libstoraged/android_${OUT_ARCH_CPU}_static/gen/proto/system/core/storaged
            ${OUT_INTERMEDIATES_ROOT}/system/core/storaged/libstoraged/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(storaged PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(storaged
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            android.hardware.health-1.0
            android.hardware.health-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealthhalutils
            libhealthshim
            libhidlbase
            liblog
            libm
            libprotobuf-cpp-lite
            libstoraged
            libutils
            libz
            packagemanager_aidl-cpp
    )
endif ()
#################### storaged ####################

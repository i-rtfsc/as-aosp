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

project(canhalconfigurator-aidl)

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

#################### canhalconfigurator-aidl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/aidl/default/tools/configurator/canhalconfigurator.cpp
        ${ANDROID_ROOT}/hardware/interfaces/automotive/can/aidl/default/tools/configurator/canprototools.cpp
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
    add_executable(canhalconfigurator-aidl ${SOURCE_FILES})

    target_include_directories(canhalconfigurator-aidl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/automotive/can/aidl/default/tools/configurator
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/can/aidl/default/tools/configurator/proto/android.hardware.automotive.can-aidl-config-format/android_${OUT_ARCH_CPU}_static_cfi/gen/proto/hardware/interfaces/automotive/can/aidl/default/tools/configurator/proto
            ${OUT_INTERMEDIATES_ROOT}/hardware/interfaces/automotive/can/aidl/default/tools/configurator/proto/android.hardware.automotive.can-aidl-config-format/android_${OUT_ARCH_CPU}_static_cfi/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(canhalconfigurator-aidl PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(canhalconfigurator-aidl
            android.hardware.automotive.can-V1-ndk
            android.hardware.automotive.can-aidl-config-format
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libprotobuf-cpp-full
            libutils
    )
endif ()
#################### canhalconfigurator-aidl ####################

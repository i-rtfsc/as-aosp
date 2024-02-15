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

project(VtsHalWifiV1_6TargetTestUtil)

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

#################### VtsHalWifiV1_6TargetTestUtil ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/wifi/1.6/vts/functional/wifi_hidl_test_utils_1_6.cpp
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
    add_library(VtsHalWifiV1_6TargetTestUtil SHARED ${SOURCE_FILES})

    target_include_directories(VtsHalWifiV1_6TargetTestUtil
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/wifi/1.6/vts/functional
    )

    # Add c/c++ flags
    target_compile_options(VtsHalWifiV1_6TargetTestUtil PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(VtsHalWifiV1_6TargetTestUtil
            VtsHalHidlTestUtils
            VtsHalWifiV1_0TargetTestUtil
            android.hardware.wifi-1.0
            android.hardware.wifi-1.3
            android.hardware.wifi-1.5
            android.hardware.wifi-1.6
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidl-gen-utils
            libhidl_gtest_helper
            libhidlbase
            liblog
            libm
            libnativehelper
            libutils
            libwifi-system-iface
    )
endif ()
#################### VtsHalWifiV1_6TargetTestUtil ####################

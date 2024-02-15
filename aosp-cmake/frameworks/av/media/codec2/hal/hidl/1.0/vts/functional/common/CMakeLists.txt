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

project(VtsHalMediaC2V1_0CommonUtil)

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

#################### VtsHalMediaC2V1_0CommonUtil ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.0/vts/functional/common/media_c2_hidl_test_common.cpp
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
    add_library(VtsHalMediaC2V1_0CommonUtil SHARED ${SOURCE_FILES})

    target_include_directories(VtsHalMediaC2V1_0CommonUtil
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/hal/hidl/1.0/vts/functional/common
    )

    # Add c/c++ flags
    target_compile_options(VtsHalMediaC2V1_0CommonUtil PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(VtsHalMediaC2V1_0CommonUtil
            VtsHalHidlTestUtils
            android.hardware.media.c2-1.0
            android.hardware.media.c2-1.1
            android.hardware.media.c2-1.2
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcodec2
            libcodec2_client_headers
            libcodec2_hidl_client-1.0
            libcodec2_hidl_client-1.1
            libcodec2_hidl_client-1.2
            libcodec2_vndk
            libcutils
            libdl
            libhidl-gen-utils
            libhidl_gtest_helper
            libhidlbase
            liblog
            libm
            libutils
    )
endif ()
#################### VtsHalMediaC2V1_0CommonUtil ####################

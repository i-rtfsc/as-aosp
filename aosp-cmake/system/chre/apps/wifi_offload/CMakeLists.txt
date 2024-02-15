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

project(wifi_offload_types)

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

#################### wifi_offload_types ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/channel_histogram.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/flatbuffers_serialization.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/preferred_network.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/rpc_log_record.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/scan_config.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/scan_filter.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/scan_params.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/scan_record.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/scan_result.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/scan_result_message.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/scan_stats.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/ssid.cc
        ${ANDROID_ROOT}/system/chre/apps/wifi_offload/utility.cc
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
    add_library(wifi_offload_types SHARED ${SOURCE_FILES})

    target_include_directories(wifi_offload_types
            PUBLIC
            ${ANDROID_ROOT}/system/chre/apps/wifi_offload/include
            ${ANDROID_ROOT}/system/chre/apps/wifi_offload
    )

    # Add c/c++ flags
    target_compile_options(wifi_offload_types PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(wifi_offload_types
            chre_api
            chre_flatbuffers
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### wifi_offload_types ####################

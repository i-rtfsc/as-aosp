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

project(libbt-bta-ext)

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

#################### libbt-bta-ext ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/bluetooth_ext/system_bt_ext/bta/tws_plus/bta_tws_plus_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/bluetooth_ext/system_bt_ext/bta/tws_plus/bta_tws_plus_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/bluetooth_ext/system_bt_ext/bta/tws_plus/bta_tws_plus_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/bluetooth_ext/system_bt_ext/bta/ba/bta_ba.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/bluetooth_ext/system_bt_ext/bta/tws_plus/ag/bta_ag_twsp_dev.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/bluetooth_ext/system_bt_ext/bta/tws_plus/ag/bta_ag_twsp_sco.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/bluetooth_ext/system_bt_ext/bta/swb/bta_ag_swb.cc
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
    add_library(libbt-bta-ext SHARED ${SOURCE_FILES})

    target_include_directories(libbt-bta-ext
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/bluetooth_ext/system_bt_ext/bta/include
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/bluetooth_ext/system_bt_ext/bta
    )

    # Add c/c++ flags
    target_compile_options(libbt-bta-ext PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-bta-ext
            libbluetooth-types
            libbluetooth_headers
            libbt-utils
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhardware
            libm
            libosi
            libutils
    )
endif ()
#################### libbt-bta-ext ####################

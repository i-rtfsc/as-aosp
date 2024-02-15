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

project(libbt-bta_qti_adva)

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

#################### libbt-bta_qti_adva ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/csip/bta_csip_act.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/csip/bta_csip_api.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/csip/bta_csip_main.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/csip/bta_csip_utils.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap/ascs_client.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap/pacs_client.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap/gattc_ops_queue.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap/gatts_ops_queue.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap/uclient_main.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap/uclient_strm_mgr.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap/uclient_strm_tracker.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap/connected_iso.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap/uclient_alarm.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/vcp/bta_vcp_controller.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/dm/bta_dm_adv_audio.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/mcp/bta_mcp_main.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/tmap/bta_tmap_main.cc
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/cc/bta_cc_main.cc
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
    add_library(libbt-bta_qti_adva SHARED ${SOURCE_FILES})

    target_include_directories(libbt-bta_qti_adva
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/include
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta/bap
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/bt/bt_adv_audio/system/bt/bta
    )

    # Add c/c++ flags
    target_compile_options(libbt-bta_qti_adva PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-bta_qti_adva
            libbluetooth-types
            libbluetooth_headers
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
    )
endif ()
#################### libbt-bta_qti_adva ####################

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

project(libbt-bta_qti)

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

#################### libbt-bta_qti ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_at.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_cfg.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_ci.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_cmd.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_rfc.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_sco.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag/bta_ag_sdp.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ar/bta_ar.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/av/bta_av_aact.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/av/bta_av_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/av/bta_av_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/av/bta_av_cfg.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/av/bta_av_ci.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/av/bta_av_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/av/bta_av_ssm.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/dm/bta_dm_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/dm/bta_dm_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/dm/bta_dm_cfg.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/dm/bta_dm_ci.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/dm/bta_dm_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/dm/bta_dm_pm.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/dm/bta_dm_sco.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gattc_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gattc_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gattc_cache.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gattc_db_storage.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gattc_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gattc_queue.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gattc_utils.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gatts_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gatts_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gatts_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/bta_gatts_utils.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/database.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/gatt/database_builder.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hearing_aid/hearing_aid.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hearing_aid/hearing_aid_audio_source.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/has/has_client.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/has/has_ctp.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/has/has_journal.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/has/has_preset.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/has/has_types.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hf_client/bta_hf_client_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hf_client/bta_hf_client_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hf_client/bta_hf_client_at.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hf_client/bta_hf_client_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hf_client/bta_hf_client_rfc.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hf_client/bta_hf_client_sco.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hf_client/bta_hf_client_sdp.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hh/bta_hh_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hh/bta_hh_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hh/bta_hh_cfg.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hh/bta_hh_le.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hh/bta_hh_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hh/bta_hh_utils.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hd/bta_hd_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hd/bta_hd_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hd/bta_hd_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/jv/bta_jv_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/jv/bta_jv_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/jv/bta_jv_cfg.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/jv/bta_jv_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/mce/bta_mce_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/mce/bta_mce_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/mce/bta_mce_cfg.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/mce/bta_mce_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/pan/bta_pan_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/pan/bta_pan_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/pan/bta_pan_ci.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/pan/bta_pan_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/sdp/bta_sdp.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/sdp/bta_sdp_act.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/sdp/bta_sdp_api.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/sdp/bta_sdp_cfg.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/sys/bta_sys_conn.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/sys/bta_sys_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/sys/utl.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/direction_finder/bta_atp_locator.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/direction_finder/gattc_ops_queue.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/direction_finder/bta_ctes_main.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/direction_finder/gatts_ops_queue.cc
        ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/direction_finder/bta_atp_result_collector.cc
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
    add_library(libbt-bta_qti SHARED ${SOURCE_FILES})

    target_include_directories(libbt-bta_qti
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/include
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/sys
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/dm
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hd
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/hh
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/direction_finder
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta/ag
            ${ANDROID_ROOT}/vendor/qcom/opensource/commonsys/system/bt/bta
    )

    # Add c/c++ flags
    target_compile_options(libbt-bta_qti PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -DBLUEDROID_DEBUG
    )

    target_link_libraries(libbt-bta_qti
            libaudio-hearing-aid-hw-utils_qti
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
#################### libbt-bta_qti ####################

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

project(libstnfc-nci)

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

#################### libstnfc-nci ####################
set(FILES
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/ce/nfa_ce_act.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/ce/nfa_ce_api.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/ce/nfa_ce_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/dm/nfa_dm_act.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/dm/nfa_dm_api.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/dm/nfa_dm_cfg.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/dm/nfa_dm_discover.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/dm/nfa_dm_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/dm/nfa_dm_ndef.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/ee/nfa_ee_act.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/ee/nfa_ee_api.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/ee/nfa_ee_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/hci/nfa_hci_act.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/hci/nfa_hci_api.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/hci/nfa_hci_ci.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/hci/nfa_hci_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/hci/nfa_hci_utils.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/p2p/nfa_p2p_act.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/p2p/nfa_p2p_api.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/p2p/nfa_p2p_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/rw/nfa_rw_act.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/rw/nfa_rw_api.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/rw/nfa_rw_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/sys/nfa_sys_cback.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/sys/nfa_sys_cfg.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/sys/nfa_sys_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/sys/nfa_sys_ptim.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/llcp/llcp_api.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/llcp/llcp_dlc.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/llcp/llcp_link.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/llcp/llcp_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/llcp/llcp_sdp.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/llcp/llcp_util.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/nci/nci_hmsgs.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/nci/nci_hrcv.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/ndef/ndef_utils.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/nfc/nfc_ee.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/nfc/nfc_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/nfc/nfc_ncif.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/nfc/nfc_task.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/nfc/nfc_test.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/nfc/nfc_utils.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/nfc/nfc_vs.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/ce_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/ce_t3t.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/ce_t4t.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_ci.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_i93.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_main.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_mfc.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_t1t.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_t1t_ndef.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_t2t.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_t2t_ndef.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_t3t.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_t4t.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/rw_t5t.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/tags/tags_int.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/adaptation/CrcChecksum.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/adaptation/NfcAdaptation.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/adaptation/debug_lmrt.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/adaptation/debug_nfcsnoop.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/adaptation/libmain.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/adaptation/nfc_config.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/gki/ulinux/gki_ulinux.cc
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
    add_library(libstnfc-nci SHARED ${SOURCE_FILES})

    target_include_directories(libstnfc-nci
            PUBLIC
            ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/include
            ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/gki/ulinux
            ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/gki/common
            ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfc/include
            ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/nfa/include
            ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src
            ${OUT_INTERMEDIATES_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/statslog_nfc_st.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libstnfc-nci PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libstnfc-nci
            android.hardware.nfc-V1-ndk
            android.hardware.nfc-1.0
            android.hardware.nfc-1.1
            android.hardware.nfc-1.2
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libstatslog_nfc_st
            libstnfcutils
            libutils
            libutils_headers
            libz
    )
endif ()
#################### libstnfc-nci ####################

#################### libstatslog_nfc_st ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/statslog_nfc_st.cpp/gen/statslog_nfc_st.cpp
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
    add_library(libstatslog_nfc_st SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_nfc_st
            PUBLIC
            ${ANDROID_ROOT}/vendor/st/opensource/commonsys/system/nfc/src
            ${OUT_INTERMEDIATES_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/statslog_nfc_st.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/vendor/st/opensource/commonsys/system/nfc/src/statslog_nfc_st.h/gen
    )

    target_link_libraries(libstatslog_nfc_st
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libm
            libstatssocket
    )
endif ()
#################### libstatslog_nfc_st ####################

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

project(nqnfc_rw_fuzzer)

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

#################### nqnfc_ce_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/ce_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/ce_t3t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/ce_t4t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/tags_int.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/ce/main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/ce/stubs.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/ce/t3t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/ce/t4t.cc
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
    add_executable(nqnfc_ce_fuzzer ${SOURCE_FILES})

    target_include_directories(nqnfc_ce_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/inc
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src
    )

    # Add c/c++ flags
    target_compile_options(nqnfc_ce_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nqnfc_ce_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libnfcutils
    )
endif ()
#################### nqnfc_ce_fuzzer ####################

#################### nqnfc_ndef_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/dm/nfa_dm_ndef.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/ndef/ndef_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/ndef/main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/ndef/stubs.cc
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
    add_executable(nqnfc_ndef_fuzzer ${SOURCE_FILES})

    target_include_directories(nqnfc_ndef_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/inc
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src
    )

    # Add c/c++ flags
    target_compile_options(nqnfc_ndef_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nqnfc_ndef_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libnfcutils
    )
endif ()
#################### nqnfc_ndef_fuzzer ####################

#################### nqnfc_llcp_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_dlc.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_link.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_sdp.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_util.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/llcp/llcp.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/llcp/stubs.cc
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
    add_executable(nqnfc_llcp_fuzzer ${SOURCE_FILES})

    target_include_directories(nqnfc_llcp_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/inc
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src
    )

    # Add c/c++ flags
    target_compile_options(nqnfc_llcp_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nqnfc_llcp_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libnfcutils
    )
endif ()
#################### nqnfc_llcp_fuzzer ####################

#################### libsn100nfc-nci ####################
set(FILES
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/ce/nfa_ce_act.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/ce/nfa_ce_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/ce/nfa_ce_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/dm/nfa_dm_act.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/dm/nfa_dm_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/dm/nfa_dm_cfg.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/dm/nfa_dm_discover.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/dm/nfa_dm_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/dm/nfa_dm_ndef.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/ee/nfa_ee_act.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/ee/nfa_ee_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/ee/nfa_ee_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/hci/nfa_hci_act.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/hci/nfa_hci_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/hci/nfa_hci_ci.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/hci/nfa_hci_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/hci/nfa_hci_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/p2p/nfa_p2p_act.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/p2p/nfa_p2p_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/p2p/nfa_p2p_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/rw/nfa_rw_act.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/rw/nfa_rw_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/rw/nfa_rw_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/scr/nfa_scr_act.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/scr/nfa_scr_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/scr/nfa_scr_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/sys/nfa_sys_cback.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/sys/nfa_sys_cfg.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/sys/nfa_sys_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/sys/nfa_sys_ptim.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/t4tnfcee/nfa_nfcee_act.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/t4tnfcee/nfa_nfcee_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/t4tnfcee/nfa_nfcee_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/wlc/nfa_wlc_act.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/wlc/nfa_wlc_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/wlc/nfa_wlc_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/srd/nfa_srd_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_api.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_dlc.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_link.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_sdp.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/llcp/llcp_util.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nci/nci_hmsgs.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nci/nci_hrcv.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/ndef/ndef_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_ee.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_ncif.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_task.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_test.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_vs.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/ce_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/ce_t3t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/ce_t4t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_i93.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_mfc.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t1t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t1t_ndef.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t2t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t2t_ndef.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t3t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t4t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t5t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/tags_int.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/adaptation/CrcChecksum.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/adaptation/NfcAdaptation.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/adaptation/debug_lmrt.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/adaptation/debug_nfcsnoop.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/adaptation/libmain.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/adaptation/nfc_config.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux/gki_ulinux.cc
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
    add_library(libsn100nfc-nci SHARED ${SOURCE_FILES})

    target_include_directories(libsn100nfc-nci
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src
    )

    # Add c/c++ flags
    target_compile_options(libsn100nfc-nci PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libsn100nfc-nci
            android.hardware.nfc-V1-ndk
            android.hardware.nfc-1.0
            android.hardware.nfc-1.1
            android.hardware.nfc-1.2
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libmemunreachable
            libsn100nfcutils
            libstatslog_nfc
            libutils
            libz
            vendor.nxp.hardware.nfc-1.0
            vendor.nxp.hardware.nfc-1.1
            vendor.nxp.hardware.nfc-2.0
            vendor.nxp.nxpnfc_aidl-V1-ndk
    )
endif ()
#################### libsn100nfc-nci ####################

#################### nqnfc_nci_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nci/nci_hmsgs.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nci/nci_hrcv.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_ee.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_ncif.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_task.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_test.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/nfc/nfc_vs.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/adaptation/debug_lmrt.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/nci/hal.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/nci/nci.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/nci/stubs.cc
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
    add_executable(nqnfc_nci_fuzzer ${SOURCE_FILES})

    target_include_directories(nqnfc_nci_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/inc
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src
    )

    # Add c/c++ flags
    target_compile_options(nqnfc_nci_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nqnfc_nci_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libnfcutils
            libstatslog_nfc
    )
endif ()
#################### nqnfc_nci_fuzzer ####################

#################### nqlibstatslog_nfc ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nqstatslog_nfc.cpp/gen/nqstatslog_nfc.cpp
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
    add_library(nqlibstatslog_nfc SHARED ${SOURCE_FILES})

    target_include_directories(nqlibstatslog_nfc
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src
            ${OUT_INTERMEDIATES_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nqstatslog_nfc.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nqstatslog_nfc.h/gen
    )

    target_link_libraries(nqlibstatslog_nfc
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
#################### nqlibstatslog_nfc ####################

#################### nqnfc_rw_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_i93.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_mfc.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t1t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t1t_ndef.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t2t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t2t_ndef.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t3t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t4t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/rw_t5t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/tags/tags_int.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/rw/main.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/rw/mfc.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/rw/stubs.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/rw/t1t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/rw/t2t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/rw/t3t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/rw/t4t.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/rw/t5t.cc
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
    add_executable(nqnfc_rw_fuzzer ${SOURCE_FILES})

    target_include_directories(nqnfc_rw_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/ulinux
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/gki/common
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfc/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/nfa/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src/fuzzers/inc
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/src
    )

    # Add c/c++ flags
    target_compile_options(nqnfc_rw_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nqnfc_rw_fuzzer
            libbase
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libnfcutils
    )
endif ()
#################### nqnfc_rw_fuzzer ####################

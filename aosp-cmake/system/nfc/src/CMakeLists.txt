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

project(nfc_ndef_fuzzer)

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

#################### libstatslog_nfc ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/statslog_nfc.cpp/gen/statslog_nfc.cpp
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
    add_library(libstatslog_nfc SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_nfc
            PUBLIC
            ${ANDROID_ROOT}/system/nfc/src
            ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/statslog_nfc.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/statslog_nfc.h/gen
    )

    target_link_libraries(libstatslog_nfc
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
#################### libstatslog_nfc ####################

#################### nfc_nci_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/nfc/src/nfc/nci/nci_hmsgs.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nci/nci_hrcv.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_ee.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_ncif.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_task.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_test.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_vs.cc
        ${ANDROID_ROOT}/system/nfc/src/adaptation/debug_lmrt.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/nci/hal.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/nci/nci.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/nci/stubs.cc
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
    add_executable(nfc_nci_fuzzer ${SOURCE_FILES})

    target_include_directories(nfc_nci_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/nfc/src/include
            ${ANDROID_ROOT}/system/nfc/src/gki/ulinux
            ${ANDROID_ROOT}/system/nfc/src/gki/common
            ${ANDROID_ROOT}/system/nfc/src/nfc/include
            ${ANDROID_ROOT}/system/nfc/src/nfa/include
            ${ANDROID_ROOT}/system/nfc/src/fuzzers/inc
            ${ANDROID_ROOT}/system/nfc/src
            ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/statslog_nfc.h/gen
    )

    # Add c/c++ flags
    target_compile_options(nfc_nci_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nfc_nci_fuzzer
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
#################### nfc_nci_fuzzer ####################

#################### nfc_ce_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/ce_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/ce_t3t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/ce_t4t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/tags_int.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/ce/main.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/ce/stubs.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/ce/t3t.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/ce/t4t.cc
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
    add_executable(nfc_ce_fuzzer ${SOURCE_FILES})

    target_include_directories(nfc_ce_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/nfc/src/include
            ${ANDROID_ROOT}/system/nfc/src/gki/ulinux
            ${ANDROID_ROOT}/system/nfc/src/gki/common
            ${ANDROID_ROOT}/system/nfc/src/nfc/include
            ${ANDROID_ROOT}/system/nfc/src/nfa/include
            ${ANDROID_ROOT}/system/nfc/src/fuzzers/inc
            ${ANDROID_ROOT}/system/nfc/src
    )

    # Add c/c++ flags
    target_compile_options(nfc_ce_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nfc_ce_fuzzer
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
#################### nfc_ce_fuzzer ####################

#################### nfc_llcp_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_dlc.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_link.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_sdp.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_util.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/llcp/llcp.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/llcp/stubs.cc
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
    add_executable(nfc_llcp_fuzzer ${SOURCE_FILES})

    target_include_directories(nfc_llcp_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/nfc/src/include
            ${ANDROID_ROOT}/system/nfc/src/gki/ulinux
            ${ANDROID_ROOT}/system/nfc/src/gki/common
            ${ANDROID_ROOT}/system/nfc/src/nfc/include
            ${ANDROID_ROOT}/system/nfc/src/nfa/include
            ${ANDROID_ROOT}/system/nfc/src/fuzzers/inc
            ${ANDROID_ROOT}/system/nfc/src
    )

    # Add c/c++ flags
    target_compile_options(nfc_llcp_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nfc_llcp_fuzzer
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
#################### nfc_llcp_fuzzer ####################

#################### nfc_integration_fuzzer_proto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/nfc_integration_fuzzer_proto/android_${OUT_ARCH_CPU}_static/gen/proto/system/nfc/src/fuzzers/integration/nfc_integration_fuzzer.pb.cc
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
    add_library(nfc_integration_fuzzer_proto SHARED ${SOURCE_FILES})

    target_include_directories(nfc_integration_fuzzer_proto
            PUBLIC
            ${ANDROID_ROOT}/system/nfc/src
            ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/nfc_integration_fuzzer_proto/android_${OUT_ARCH_CPU}_static/gen/proto/system/nfc/src
            ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/nfc_integration_fuzzer_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(nfc_integration_fuzzer_proto PRIVATE
            -Werror
    )

    target_link_libraries(nfc_integration_fuzzer_proto
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-full
    )
endif ()
#################### nfc_integration_fuzzer_proto ####################

#################### nfc_rw_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_i93.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_mfc.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t1t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t1t_ndef.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t2t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t2t_ndef.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t3t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t4t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t5t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/tags_int.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/rw/main.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/rw/mfc.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/rw/stubs.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/rw/t1t.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/rw/t2t.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/rw/t3t.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/rw/t4t.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/rw/t5t.cc
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
    add_executable(nfc_rw_fuzzer ${SOURCE_FILES})

    target_include_directories(nfc_rw_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/nfc/src/include
            ${ANDROID_ROOT}/system/nfc/src/gki/ulinux
            ${ANDROID_ROOT}/system/nfc/src/gki/common
            ${ANDROID_ROOT}/system/nfc/src/nfc/include
            ${ANDROID_ROOT}/system/nfc/src/nfa/include
            ${ANDROID_ROOT}/system/nfc/src/fuzzers/inc
            ${ANDROID_ROOT}/system/nfc/src
    )

    # Add c/c++ flags
    target_compile_options(nfc_rw_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nfc_rw_fuzzer
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
#################### nfc_rw_fuzzer ####################

#################### libnfc-nci ####################
set(FILES
        ${ANDROID_ROOT}/system/nfc/src/nfa/ce/nfa_ce_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ce/nfa_ce_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ce/nfa_ce_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_cfg.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_discover.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_ndef.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ee/nfa_ee_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ee/nfa_ee_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ee/nfa_ee_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_ci.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/p2p/nfa_p2p_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/p2p/nfa_p2p_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/p2p/nfa_p2p_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/rw/nfa_rw_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/rw/nfa_rw_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/rw/nfa_rw_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/sys/nfa_sys_cback.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/sys/nfa_sys_cfg.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/sys/nfa_sys_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/sys/nfa_sys_ptim.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_dlc.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_link.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_sdp.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_util.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nci/nci_hmsgs.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nci/nci_hrcv.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/ndef/ndef_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_ee.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_ncif.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_task.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_test.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_vs.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/ce_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/ce_t3t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/ce_t4t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_i93.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_mfc.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t1t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t1t_ndef.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t2t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t2t_ndef.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t3t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t4t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t5t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/tags_int.cc
        ${ANDROID_ROOT}/system/nfc/src/adaptation/CrcChecksum.cc
        ${ANDROID_ROOT}/system/nfc/src/adaptation/NfcAdaptation.cc
        ${ANDROID_ROOT}/system/nfc/src/adaptation/debug_lmrt.cc
        ${ANDROID_ROOT}/system/nfc/src/adaptation/debug_nfcsnoop.cc
        ${ANDROID_ROOT}/system/nfc/src/adaptation/libmain.cc
        ${ANDROID_ROOT}/system/nfc/src/adaptation/nfc_config.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/ulinux/gki_ulinux.cc
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
    add_library(libnfc-nci SHARED ${SOURCE_FILES})

    target_include_directories(libnfc-nci
            PUBLIC
            ${ANDROID_ROOT}/system/nfc/src/include
            ${ANDROID_ROOT}/system/nfc/src/gki/ulinux
            ${ANDROID_ROOT}/system/nfc/src/gki/common
            ${ANDROID_ROOT}/system/nfc/src/nfa/include
            ${ANDROID_ROOT}/system/nfc/src/nfc/include
            ${ANDROID_ROOT}/system/nfc/src
            ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/statslog_nfc.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libnfc-nci PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libnfc-nci
            android.hardware.nfc-V1-ndk
            android.hardware.nfc-1.0
            android.hardware.nfc-1.1
            android.hardware.nfc-1.2
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libnfcutils
            libstatslog_nfc
            libutils
            libz
    )
endif ()
#################### libnfc-nci ####################

#################### nfc_integration_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/nfc/src/adaptation/debug_nfcsnoop.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/nfc_integration_fuzzer.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/nfc_integration_fuzzer_impl.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/nfc_task_helpers.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/fakes/base64_fakes.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/fakes/gki_time_fakes.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/fakes/gki_ulinux_fakes.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/fakes/hal_fakes.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/fakes/nfc_adaptation_fakes.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/fakes/nfc_adaptation_libmain_fakes.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/fakes/nfc_config_fakes.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/fakes/statslog_fakes.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/system/nfc/src/adaptation/debug_lmrt.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ce/nfa_ce_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ce/nfa_ce_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ce/nfa_ce_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_cfg.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_discover.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_ndef.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ee/nfa_ee_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ee/nfa_ee_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/ee/nfa_ee_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_ci.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/hci/nfa_hci_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/p2p/nfa_p2p_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/p2p/nfa_p2p_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/p2p/nfa_p2p_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/rw/nfa_rw_act.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/rw/nfa_rw_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/rw/nfa_rw_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/sys/nfa_sys_cback.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/sys/nfa_sys_cfg.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/sys/nfa_sys_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfa/sys/nfa_sys_ptim.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_api.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_dlc.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_link.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_sdp.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/llcp/llcp_util.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nci/nci_hmsgs.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nci/nci_hrcv.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/ndef/ndef_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_ee.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_ncif.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_task.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_test.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/nfc/nfc_vs.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/ce_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/ce_t3t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/ce_t4t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_i93.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_main.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_mfc.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t1t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t1t_ndef.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t2t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t2t_ndef.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t3t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t4t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/rw_t5t.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/tags/tags_int.cc
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
    add_executable(nfc_integration_fuzzer ${SOURCE_FILES})

    target_include_directories(nfc_integration_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/nfc/src/include
            ${ANDROID_ROOT}/system/nfc/src/gki/ulinux
            ${ANDROID_ROOT}/system/nfc/src/gki/common
            ${ANDROID_ROOT}/system/nfc/src/nfc/include
            ${ANDROID_ROOT}/system/nfc/src/nfa/include
            ${ANDROID_ROOT}/system/nfc/src/fuzzers/inc
            ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration
            ${ANDROID_ROOT}/system/nfc/src/fuzzers/integration/fakes
            ${ANDROID_ROOT}/system/nfc/src
            ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/nfc_integration_fuzzer_proto/android_${OUT_ARCH_CPU}_static_fuzzer/gen/proto/system/nfc/src
            ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/nfc_integration_fuzzer_proto/android_${OUT_ARCH_CPU}_static_fuzzer/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/system/nfc/src/statslog_nfc.h/gen
    )

    # Add c/c++ flags
    target_compile_options(nfc_integration_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nfc_integration_fuzzer
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
            libprotobuf-cpp-full
            libprotobuf-mutator
            libstatslog_nfc
            libutils
            libz
            nfc_integration_fuzzer_proto
    )
endif ()
#################### nfc_integration_fuzzer ####################

#################### nfc_ndef_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/nfc/src/nfa/dm/nfa_dm_ndef.cc
        ${ANDROID_ROOT}/system/nfc/src/nfc/ndef/ndef_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_buffer.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/common/gki_time.cc
        ${ANDROID_ROOT}/system/nfc/src/gki/ulinux/gki_ulinux.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_cmn.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/fuzz_utils.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/ndef/main.cc
        ${ANDROID_ROOT}/system/nfc/src/fuzzers/ndef/stubs.cc
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
    add_executable(nfc_ndef_fuzzer ${SOURCE_FILES})

    target_include_directories(nfc_ndef_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/nfc/src/include
            ${ANDROID_ROOT}/system/nfc/src/gki/ulinux
            ${ANDROID_ROOT}/system/nfc/src/gki/common
            ${ANDROID_ROOT}/system/nfc/src/nfc/include
            ${ANDROID_ROOT}/system/nfc/src/nfa/include
            ${ANDROID_ROOT}/system/nfc/src/fuzzers/inc
            ${ANDROID_ROOT}/system/nfc/src
    )

    # Add c/c++ flags
    target_compile_options(nfc_ndef_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(nfc_ndef_fuzzer
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
#################### nfc_ndef_fuzzer ####################

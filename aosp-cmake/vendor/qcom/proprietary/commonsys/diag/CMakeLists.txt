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

project(libdiag_system)

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

#################### libdiag_system ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_dci.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/ts_linux.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_event.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_log.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/msg_arrays_i.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_msg.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_pkt.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diagsvc_malloc.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_logger.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_qshrink4_db_parser.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_qdss.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_adpl.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_diag_id.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diagcallback.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_hidl_client.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_key_info.c
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
    add_library(libdiag_system SHARED ${SOURCE_FILES})

    target_include_directories(libdiag_system
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag
    )

    target_link_libraries(libdiag_system
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libcutils_headers
            libdiag_headers_system
            libdl
            libhidlbase
            libhidlmemory
            liblog
            libm
            libutils
            vendor.qti.diaghal-1.0
            vendor_common_inc
    )
endif ()
#################### libdiag_system ####################

#################### libdiag_system.qti ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_dci.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/ts_linux.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_event.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_log.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/msg_arrays_i.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_msg.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_pkt.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diagsvc_malloc.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_logger.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_qshrink4_db_parser.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_qdss.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_adpl.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_diag_id.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diagcallback.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_hidl_client.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src/diag_lsm_key_info.c
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
    add_library(libdiag_system.qti SHARED ${SOURCE_FILES})

    target_include_directories(libdiag_system.qti
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag/src
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/diag
    )

    target_link_libraries(libdiag_system.qti
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libcutils_headers
            libdiag_headers_system
            libdl
            libhidlbase
            libhidlmemory
            liblog
            libm
            libutils
            vendor.qti.diaghal-1.0
            vendor_common_inc
    )
endif ()
#################### libdiag_system.qti ####################

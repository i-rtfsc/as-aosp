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

project(libqmi_csi_system)

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

#################### libqmi_csi_system ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/qcsi/src/qmi_csi_common.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/qcsi/src/qmi_csi_xport_ipc_router.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/qcsi/src/qmi_csi_target.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/qcsi/src/qmi_csi_xport_qrtr.c
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
    add_library(libqmi_csi_system SHARED ${SOURCE_FILES})

    target_include_directories(libqmi_csi_system
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework
    )

    # Add c/c++ flags
    target_compile_options(libqmi_csi_system PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libqmi_csi_system
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libqmi_common_headers_sys
            libqmi_common_system
            libqmi_csi_headers_sys
            libqmi_encdec_system
            libutils
    )
endif ()
#################### libqmi_csi_system ####################

#################### libqmi_encdec_system ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/encdec/qmi_idl_accessor.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/encdec/qmi_idl_lib.c
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
    add_library(libqmi_encdec_system SHARED ${SOURCE_FILES})

    target_include_directories(libqmi_encdec_system
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework
    )

    # Add c/c++ flags
    target_compile_options(libqmi_encdec_system PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libqmi_encdec_system
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libqmi_common_headers_sys
            libutils
    )
endif ()
#################### libqmi_encdec_system ####################

#################### libqmi_cci_system ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/qcci/src/qmi_cci_common.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/qcci/src/qmi_cci_target.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/qcci/src/qmi_cci_xport_ipc_router.c
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/qcci/src/qmi_cci_xport_qrtr.c
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
    add_library(libqmi_cci_system SHARED ${SOURCE_FILES})

    target_include_directories(libqmi_cci_system
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework
    )

    # Add c/c++ flags
    target_compile_options(libqmi_cci_system PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libqmi_cci_system
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libcutils_headers
            libdl
            liblog
            libm
            libqmi_cci_headers_sys
            libqmi_common_headers_sys
            libqmi_common_system
            libqmi_encdec_system
            libutils
            vendor_common_inc
    )
endif ()
#################### libqmi_cci_system ####################

#################### libqmi_common_system ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework/common/src/common_v01.c
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
    add_library(libqmi_common_system SHARED ${SOURCE_FILES})

    target_include_directories(libqmi_common_system
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/qmi-framework
    )

    # Add c/c++ flags
    target_compile_options(libqmi_common_system PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libqmi_common_system
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libm
            libqmi_common_headers_sys
            libutils
    )
endif ()
#################### libqmi_common_system ####################

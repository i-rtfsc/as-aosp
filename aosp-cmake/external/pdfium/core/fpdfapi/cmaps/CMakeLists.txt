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

project(libpdfium-cmaps)

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

#################### libpdfium-cmaps ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/fpdf_cmaps.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/Adobe-CNS1-UCS2_5.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/B5pc-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/B5pc-V_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/CNS-EUC-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/CNS-EUC-V_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/ETen-B5-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/ETen-B5-V_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/ETenms-B5-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/ETenms-B5-V_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/HKscs-B5-H_5.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/HKscs-B5-V_5.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/UniCNS-UCS2-H_3.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/UniCNS-UCS2-V_3.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/UniCNS-UTF16-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/CNS1/cmaps_cns1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/Adobe-GB1-UCS2_5.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GB-EUC-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GB-EUC-V_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GBK-EUC-H_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GBK-EUC-V_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GBK2K-H_5.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GBK2K-V_5.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GBKp-EUC-H_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GBKp-EUC-V_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GBpc-EUC-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/GBpc-EUC-V_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/UniGB-UCS2-H_4.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/UniGB-UCS2-V_4.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/GB1/cmaps_gb1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/83pv-RKSJ-H_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/90ms-RKSJ-H_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/90ms-RKSJ-V_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/90msp-RKSJ-H_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/90msp-RKSJ-V_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/90pv-RKSJ-H_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/Add-RKSJ-H_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/Add-RKSJ-V_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/Adobe-Japan1-UCS2_4.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/EUC-H_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/EUC-V_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/Ext-RKSJ-H_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/Ext-RKSJ-V_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/H_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/UniJIS-UCS2-HW-H_4.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/UniJIS-UCS2-HW-V_4.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/UniJIS-UCS2-H_4.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/UniJIS-UCS2-V_4.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/V_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Japan1/cmaps_japan1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/Adobe-Korea1-UCS2_2.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/KSC-EUC-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/KSC-EUC-V_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/KSCms-UHC-HW-H_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/KSCms-UHC-HW-V_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/KSCms-UHC-H_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/KSCms-UHC-V_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/KSCpc-EUC-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/UniKS-UCS2-H_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/UniKS-UCS2-V_1.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/UniKS-UTF16-H_0.cpp
        ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps/Korea1/cmaps_korea1.cpp
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
    add_library(libpdfium-cmaps SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-cmaps
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/core/fpdfapi/cmaps
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-cmaps PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-cmaps
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpdfium-fxcrt
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-cmaps ####################

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

project(libwfds)

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

#################### libwfds ####################
set(FILES
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/lib/lib_wprint.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/lib/plugin_db.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/lib/printable_area.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/lib/printer.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/lib/wprint_msgq.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/lib/wprintJNI.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/ipphelper/ipp_print.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/ipphelper/ipphelper.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/ipphelper/ippstatus_capabilities.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/ipphelper/ippstatus_monitor.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/lib_pclm.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/lib_pwg.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/genPCLm/src/genPCLm.cpp
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/genPCLm/src/genJPEGStrips.cpp
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/pdf_render.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/plugin_pcl.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/plugin_pdf.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/pclm_wrapper_api.cpp
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/wprint_image.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/wprint_image_platform.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/wprint_mupdf.c
        ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/wprint_scaler.c
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
    add_library(libwfds SHARED ${SOURCE_FILES})

    target_include_directories(libwfds
            PUBLIC
            ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/include
            ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/plugins/genPCLm/inc
            ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni/ipphelper
            ${ANDROID_ROOT}/packages/services/BuiltInPrintService/jni
    )

    # Add c/c++ flags
    target_compile_options(libwfds PRIVATE
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libwfds
            jni_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcups
            libdl
            libjpeg_static_ndk
            liblog
            libm
            libz
    )
endif ()
#################### libwfds ####################

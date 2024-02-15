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

project(libmmi)

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

#################### libmmi ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/util_comm.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/util_string.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/util_system.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/util_ui.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/textview.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/button.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/listview.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/layout.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/nv.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/events.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/xmlparse.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/lang.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/board.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/mmi_log.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi/interface.cpp
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
    add_library(libmmi SHARED ${SOURCE_FILES})

    target_include_directories(libmmi
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/libmmi
    )

    # Add c/c++ flags
    target_compile_options(libmmi PRIVATE
            -Wall
    )

    target_link_libraries(libmmi
            libEGL
            libGLESv2
            libarect
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdiag_headers_system
            libdiag_system
            libdl
            libdng_sdk
            libexpat
            libft2
            libgui
            libheif
            libincfs-utils
            libjpeg
            liblog
            libm
            libnativewindow
            libpiex
            libpng
            libsfntly
            libskia
            libtiff
            libui
            libutils
            libvulkan
            libwebp-decode
            libwebp-encode
            libwuffs_mirror_release_c
            libxml2
            libz
            vendor_common_inc
    )
endif ()
#################### libmmi ####################

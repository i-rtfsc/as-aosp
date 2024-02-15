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

project(mmi_diag)

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

#################### mmi ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/mmi.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/mmi_handle.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/module.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/module_local.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/input.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/draw.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/config.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/func_map.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/input_listener.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/input_listener_touch.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/input_listener_key.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/controller.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/mmi_cfg.cpp
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
    add_executable(mmi ${SOURCE_FILES})

    target_include_directories(mmi
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi
    )

    # Add c/c++ flags
    target_compile_options(mmi PRIVATE
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(mmi
            libEGL
            libGLESv2
            libarect
            libbase
            libbinder
            libbinder_ndk
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
            libhidlbase
            libimage_io
            libincfs-utils
            libjpeg
            libjpegdecoder
            libjpegencoder
            liblog
            libm
            libmmi
            libmmi_headers
            libnativewindow
            libperfetto_client_experimental
            libpiex
            libpng
            libsfntly
            libskia
            libtiff
            libui
            libultrahdr
            libutils
            libvulkan
            libwebp-decode
            libwebp-encode
            libwuffs_mirror_release_c
            libxml2
            libz
            vendor.qti.hardware.factory-V1-ndk
            vendor.qti.hardware.factory-1.0
            vendor.qti.hardware.factory-1.1
            vendor_common_inc
    )
endif ()
#################### mmi ####################

#################### mmi_diag ####################
set(FILES
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/diag_main.cpp
        ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi/diag_handle.cpp
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
    add_executable(mmi_diag ${SOURCE_FILES})

    target_include_directories(mmi_diag
            PUBLIC
            ${ANDROID_ROOT}/vendor/qcom/proprietary/commonsys/fastmmi/mmi
    )

    # Add c/c++ flags
    target_compile_options(mmi_diag PRIVATE
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(mmi_diag
            libbase
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdiag_headers_system
            libdiag_system
            libdl
            libhidlbase
            liblog
            libm
            libmmi
            libmmi_headers
            libutils
            vendor.qti.hardware.factory-V1-ndk
            vendor.qti.hardware.factory-1.0
            vendor.qti.hardware.factory-1.1
            vendor_common_inc
    )
endif ()
#################### mmi_diag ####################

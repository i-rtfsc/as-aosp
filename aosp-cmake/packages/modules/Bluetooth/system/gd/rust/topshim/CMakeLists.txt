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

project(libbt_topshim_cxx)

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

#################### libbt_topshim_cxx ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/btav/btav_shim.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/btav_sink/btav_sink_shim.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/btif/btif_shim.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/controller/controller_shim.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/gatt/gatt_ble_advertiser_shim.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/gatt/gatt_ble_scanner_shim.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/gatt/gatt_shim.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/hfp/hfp_shim.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/libbt_topshim_bridge_code/gen/gensrcs/packages/modules/Bluetooth/system/gd/rust/topshim/src/btif.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/libbt_topshim_bridge_code/gen/gensrcs/packages/modules/Bluetooth/system/gd/rust/topshim/src/controller.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/libbt_topshim_bridge_code/gen/gensrcs/packages/modules/Bluetooth/system/gd/rust/topshim/src/profiles/a2dp.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/libbt_topshim_bridge_code/gen/gensrcs/packages/modules/Bluetooth/system/gd/rust/topshim/src/profiles/avrcp.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/libbt_topshim_bridge_code/gen/gensrcs/packages/modules/Bluetooth/system/gd/rust/topshim/src/profiles/gatt.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/libbt_topshim_bridge_code/gen/gensrcs/packages/modules/Bluetooth/system/gd/rust/topshim/src/profiles/hfp.cc
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
    add_library(libbt_topshim_cxx SHARED ${SOURCE_FILES})

    target_include_directories(libbt_topshim_cxx
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/libbt_topshim_bridge_code/gen/gensrcs/packages/modules/Bluetooth/system/gd/rust/topshim
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/topshim/libbt_topshim_bridge_header/gen/gensrcs/packages/modules/Bluetooth/system/gd/rust/topshim
    )

    # Add c/c++ flags
    target_compile_options(libbt_topshim_cxx PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libbt_topshim_cxx
            libbase
            libc
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libcutils
            libdl
            libevent
            liblog
            libm
            libmodpb64
    )
endif ()
#################### libbt_topshim_cxx ####################

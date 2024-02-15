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

project(lib-bt-packets-avrcp)

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

#################### lib-bt-packets-avrcp ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/avrcp_browse_packet.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/avrcp_packet.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/avrcp_reject_packet.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/capabilities_packet.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/change_path.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/general_reject_packet.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/get_current_player_application_setting_value.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/get_element_attributes_packet.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/get_folder_items.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/get_item_attributes.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/get_play_status_packet.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/get_total_number_of_items.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/list_player_application_setting_attributes.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/list_player_application_setting_values.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/pass_through_packet.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/play_item.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/register_notification_packet.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/set_absolute_volume.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/set_addressed_player.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/set_browsed_player.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/set_player_application_setting_value.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp/vendor_packet.cc
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
    add_library(lib-bt-packets-avrcp SHARED ${SOURCE_FILES})

    target_include_directories(lib-bt-packets-avrcp
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/packet/avrcp
    )

    # Add c/c++ flags
    target_compile_options(lib-bt-packets-avrcp PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(lib-bt-packets-avrcp
            avrcp_headers
            lib-bt-packets-base
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
#################### lib-bt-packets-avrcp ####################

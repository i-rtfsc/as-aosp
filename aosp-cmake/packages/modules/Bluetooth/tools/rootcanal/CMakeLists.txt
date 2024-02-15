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

project(librootcanal_config)

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

#################### librootcanal_config ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/librootcanal_config/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/Bluetooth/tools/rootcanal/config.pb.cc
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
    add_library(librootcanal_config SHARED ${SOURCE_FILES})

    target_include_directories(librootcanal_config
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/librootcanal_config/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/Bluetooth/tools/rootcanal
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/librootcanal_config/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(librootcanal_config PRIVATE
            -Werror
    )

    target_link_libraries(librootcanal_config
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### librootcanal_config ####################

#################### h4_packetizer_lib ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/hci/h4_packetizer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/hci/h4_parser.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/hci/hci_protocol.cc
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
    add_library(h4_packetizer_lib SHARED ${SOURCE_FILES})

    target_include_directories(h4_packetizer_lib
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal
            ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/include
    )

    # Add c/c++ flags
    target_compile_options(h4_packetizer_lib PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -Wno-unused-result
    )

    target_link_libraries(h4_packetizer_lib
            jni_headers
            libbase_headers
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### h4_packetizer_lib ####################

#################### libscriptedbeaconpayload-protos-lite ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/libscriptedbeaconpayload-protos-lite/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/Bluetooth/tools/rootcanal/model/devices/scripted_beacon_ble_payload.pb.cc
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
    add_library(libscriptedbeaconpayload-protos-lite SHARED ${SOURCE_FILES})

    target_include_directories(libscriptedbeaconpayload-protos-lite
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/libscriptedbeaconpayload-protos-lite/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/Bluetooth/tools/rootcanal
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/libscriptedbeaconpayload-protos-lite/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libscriptedbeaconpayload-protos-lite PRIVATE
            -Werror
    )

    target_link_libraries(libscriptedbeaconpayload-protos-lite
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
            libprotobuf-cpp-lite
            libc
            libdl
            libm
    )
endif ()
#################### libscriptedbeaconpayload-protos-lite ####################

#################### libbt-rootcanal ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/bit_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/byte_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/byte_observer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/fragmenting_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/iterator.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/packet_view.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/raw_builder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/view.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/lib/crypto/crypto.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/lib/hci/address.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/lib/hci/class_of_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/lib/hci/pcap_filter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/controller/acl_connection.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/controller/acl_connection_handler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/controller/controller_properties.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/controller/dual_mode_controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/controller/isochronous_connection_handler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/controller/le_advertiser.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/controller/link_layer_controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/controller/sco_connection.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/devices/baseband_sniffer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/devices/beacon.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/devices/beacon_swarm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/devices/device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/devices/hci_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/devices/link_layer_socket_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/devices/scripted_beacon.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/devices/sniffer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/hci/h4_data_channel_packetizer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/hci/h4_packetizer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/hci/h4_parser.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/hci/hci_protocol.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/hci/hci_sniffer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/hci/hci_socket_transport.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/setup/async_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/setup/device_boutique.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/setup/phy_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/setup/phy_layer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/setup/test_channel_transport.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/setup/test_command_handler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/model/setup/test_model.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/net/posix/posix_async_socket.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/net/posix/posix_async_socket_connector.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/net/posix/posix_async_socket_server.cc
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
    add_library(libbt-rootcanal SHARED ${SOURCE_FILES})

    target_include_directories(libbt-rootcanal
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal
            ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/include
            ${ANDROID_ROOT}/packages/modules/Bluetooth/tools/rootcanal/rust/include
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/librootcanal_config/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/Bluetooth/tools/rootcanal
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/librootcanal_config/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/libscriptedbeaconpayload-protos-lite/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/Bluetooth/tools/rootcanal
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/libscriptedbeaconpayload-protos-lite/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/rootcanal_bredr_bb_packets_cxx_gen/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/rootcanal_hci_packets_cxx_gen/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/tools/rootcanal/rootcanal_link_layer_packets_cxx_gen/gen
    )

    # Add c/c++ flags
    target_compile_options(libbt-rootcanal PRIVATE
            -DLOG_NDEBUG=1
            -Wno-unused-result
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libbt-rootcanal
            jni_headers
            jni_headers
            libbase
            libbase
            libbase_headers
            libbase_headers
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcrypto
            libdl
            liblog
            libm
            librootcanal_config
            librootcanal_config
            librootcanal_rs
            librootcanal_rs
            libscriptedbeaconpayload-protos-lite
            libc
            libdl
            libm
    )
endif ()
#################### libbt-rootcanal ####################

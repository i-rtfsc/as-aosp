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

project(bluetooth_stack_with_facade)

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

#################### libbluetooth_gd ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/att/att_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/common/audit_log.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/common/metric_id_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/common/stop_watch.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/common/strings.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/crypto_toolbox/aes.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/crypto_toolbox/aes_cmac.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/crypto_toolbox/crypto_toolbox.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/dumpsys/filter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/dumpsys/internal/filter_internal.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/dumpsys/reflection_schema.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/snoop_logger.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/snoop_logger_socket.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/snoop_logger_socket_thread.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/syscall_wrapper_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/address.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/class_of_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/acl_connection.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/acl_fragmenter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/acl_scheduler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/classic_acl_connection.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/le_acl_connection.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/round_robin_scheduler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/distance_measurement_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/hci_layer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/hci_metrics_logging.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/le_address_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/le_advertising_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/le_scanning_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/le_scanning_reassembler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/link_key.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/msft.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/remote_name_request.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/uuid.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/vendor_specific_event_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/iso/internal/iso_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/iso/iso_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/iso/iso_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/dynamic_channel_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/dynamic_channel_service.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/fixed_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/fixed_channel_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/fixed_channel_service.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/dumpsys_helper.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/dynamic_channel_service_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/fixed_channel_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/fixed_channel_service_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/link.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/link_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/signalling_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/l2cap_classic_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/dynamic_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/fcs.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/basic_mode_channel_data_controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/data_pipeline_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/dynamic_channel_allocator.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/dynamic_channel_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/enhanced_retransmission_mode_channel_data_controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/le_credit_based_channel_data_controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/receiver.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/scheduler_fifo.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/sender.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/dynamic_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/dynamic_channel_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/dynamic_channel_service.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/fixed_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/fixed_channel_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/fixed_channel_service.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/dynamic_channel_service_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/fixed_channel_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/fixed_channel_service_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/link.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/link_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/signalling_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/l2cap_le_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/link_options.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/logging/log_redaction.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/metrics/counter_metrics.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/metrics/metrics_state.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/metrics/utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/connectability.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/discoverability.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/inquiry.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/name_db.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/page.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/scan.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/handler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/system_properties_common.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/bit_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/byte_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/byte_observer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/fragmenting_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/iterator.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/packet_view.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/raw_builder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/view.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/channel/security_manager_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/pairing/classic_pairing_handler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/record/security_record_storage.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/ecc/multprecision.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/ecc/p_256_ecc_pp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/ecdh_keys.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/facade_configuration_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/internal/security_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/l2cap_security_module_interface.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/pairing_handler_le.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/pairing_handler_le_legacy.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/pairing_handler_le_secure_connections.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/security_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/security_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/shim/dumpsys.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/shim/dumpsys_args.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/classic_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/config_cache.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/config_cache_helper.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/le_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/legacy_config_file.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/mutation.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/mutation_entry.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/storage_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/sysprops/sysprops_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/stack_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/linux_generic/attribution_processor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/linux_generic/cmd_evt_classification.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/linux_generic/hci_processor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/linux_generic/wakelock_processor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/alarm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/files.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/reactive_semaphore.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/reactor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/repeating_alarm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/thread.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/wakelock_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/android/activity_attribution.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/hci_hal_android_hidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/android/metrics.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/android/parameter_provider.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/android/system_properties.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/android/wakelock_native.cc
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
    add_library(libbluetooth_gd SHARED ${SOURCE_FILES})

    target_include_directories(libbluetooth_gd
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/proto/libbluetooth-protos/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/Bluetooth/system/gd/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/proto/libbluetooth-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/shim/libbt_init_flags_bridge_header/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/dumpsys/bundler/BluetoothGeneratedBundlerSchema_h_bfbs/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/BluetoothGeneratedDumpsysDataSchema_h/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/BluetoothGeneratedPackets_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libbluetooth_gd PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -Wno-unused-result
    )

    target_link_libraries(libbluetooth_gd
            android.hardware.bluetooth-V1-ndk
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            jni_headers
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-protos
            libbluetooth_rust_interop
            libbt-platform-protos-lite
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libevent
            libflatbuffers-cpp
            libhidlbase
            liblog
            libm
            libmodpb64
            libstatslog_bt
            libutils
    )
endif ()
#################### libbluetooth_gd ####################

#################### bluetooth_stack_with_facade ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/facade/acl_manager_facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/facade/controller_facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/facade/facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/facade/le_acl_manager_facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/facade/le_advertising_manager_facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/facade/le_initiator_address_facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/facade/le_scanning_manager_facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/iso/facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/facade/facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/shim/facade/facade.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/facade/facade_main.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/facade/grpc_root_server.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/facade/read_only_property_server.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/grpc/grpc_module.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/common.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/common.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hal/hal_facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hal/hal_facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/acl_manager_facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/acl_manager_facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/controller_facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/controller_facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/hci_facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/hci_facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/le_acl_manager_facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/le_acl_manager_facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/le_advertising_manager_facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/le_advertising_manager_facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/le_initiator_address_facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/le_initiator_address_facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/le_scanning_manager_facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/hci/le_scanning_manager_facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/iso/facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/iso/facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/l2cap/classic/facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/l2cap/classic/facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/l2cap/le/facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/l2cap/le/facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/neighbor/facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/neighbor/facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/rootservice.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/rootservice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/security/facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/security/facade.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/shim/facade.grpc.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/BlueberryFacadeGeneratedStub_cc/gen/blueberry/facade/shim/facade.pb.cc
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
    add_executable(bluetooth_stack_with_facade ${SOURCE_FILES})

    target_include_directories(bluetooth_stack_with_facade
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/proto/libbluetooth-protos/android_${OUT_ARCH_CPU}_static_asan/gen/proto/packages/modules/Bluetooth/system/gd/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/proto/libbluetooth-protos/android_${OUT_ARCH_CPU}_static_asan/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/shim/libbt_init_flags_bridge_header/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/dumpsys/bundler/BluetoothGeneratedBundlerSchema_h_bfbs/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/BluetoothGeneratedDumpsysDataSchema_h/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/BluetoothGeneratedPackets_h/gen
    )

    # Add c/c++ flags
    target_compile_options(bluetooth_stack_with_facade PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -Wno-unused-result
    )

    target_link_libraries(bluetooth_stack_with_facade
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            breakpad_client
            jni_headers
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-protos
            libbluetooth_gd
            libbluetooth_rust_interop
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.asan
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libevent
            libflatbuffers-cpp
            libgrpc++
            libgrpc++_unsecure
            libgrpc_wrap
            libhidlbase
            liblog
            libm
            libmodpb64
            libprotobuf-cpp-full
            libstatslog_bt
            libunwindstack
            libutils
    )
endif ()
#################### bluetooth_stack_with_facade ####################

#################### libbluetooth_gd_fuzzing ####################
set(FILES
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/att/att_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/common/audit_log.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/common/metric_id_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/common/stop_watch.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/common/strings.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/crypto_toolbox/aes.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/crypto_toolbox/aes_cmac.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/crypto_toolbox/crypto_toolbox.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/dumpsys/filter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/dumpsys/internal/filter_internal.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/dumpsys/reflection_schema.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/snoop_logger.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/snoop_logger_socket.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/snoop_logger_socket_thread.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/syscall_wrapper_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/address.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/class_of_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/acl_connection.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/acl_fragmenter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/acl_scheduler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/classic_acl_connection.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/le_acl_connection.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/acl_manager/round_robin_scheduler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/distance_measurement_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/hci_layer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/hci_metrics_logging.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/le_address_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/le_advertising_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/le_scanning_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/le_scanning_reassembler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/link_key.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/msft.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/remote_name_request.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/uuid.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hci/vendor_specific_event_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/iso/internal/iso_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/iso/iso_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/iso/iso_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/dynamic_channel_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/dynamic_channel_service.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/fixed_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/fixed_channel_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/fixed_channel_service.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/dumpsys_helper.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/dynamic_channel_service_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/fixed_channel_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/fixed_channel_service_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/link.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/link_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/internal/signalling_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/classic/l2cap_classic_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/dynamic_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/fcs.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/basic_mode_channel_data_controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/data_pipeline_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/dynamic_channel_allocator.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/dynamic_channel_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/enhanced_retransmission_mode_channel_data_controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/le_credit_based_channel_data_controller.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/receiver.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/scheduler_fifo.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/internal/sender.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/dynamic_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/dynamic_channel_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/dynamic_channel_service.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/fixed_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/fixed_channel_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/fixed_channel_service.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/dynamic_channel_service_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/fixed_channel_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/fixed_channel_service_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/link.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/link_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/internal/signalling_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/l2cap_le_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/l2cap/le/link_options.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/logging/log_redaction.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/metrics/counter_metrics.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/metrics/metrics_state.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/metrics/utils.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/connectability.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/discoverability.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/inquiry.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/name_db.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/page.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/neighbor/scan.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/handler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/system_properties_common.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/bit_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/byte_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/byte_observer.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/fragmenting_inserter.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/iterator.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/packet_view.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/raw_builder.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/packet/view.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/channel/security_manager_channel.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/pairing/classic_pairing_handler.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/record/security_record_storage.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/ecc/multprecision.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/ecc/p_256_ecc_pp.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/ecdh_keys.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/facade_configuration_api.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/internal/security_manager_impl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/l2cap_security_module_interface.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/pairing_handler_le.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/pairing_handler_le_legacy.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/pairing_handler_le_secure_connections.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/security_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/security/security_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/shim/dumpsys.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/shim/dumpsys_args.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/classic_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/config_cache.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/config_cache_helper.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/le_device.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/legacy_config_file.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/mutation.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/mutation_entry.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/storage/storage_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/sysprops/sysprops_module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/module.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/stack_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/fake_timer/fake_timerfd.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/linux_generic/attribution_processor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/linux_generic/cmd_evt_classification.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/linux_generic/hci_processor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/linux_generic/wakelock_processor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/alarm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/files.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/reactive_semaphore.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/reactor.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/repeating_alarm.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/thread.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/linux_generic/wakelock_manager.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/btaa/android/activity_attribution.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/hal/hci_hal_android_hidl.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/android/metrics.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/android/parameter_provider.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/android/system_properties.cc
        ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd/os/android/wakelock_native.cc
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
    add_library(libbluetooth_gd_fuzzing SHARED ${SOURCE_FILES})

    target_include_directories(libbluetooth_gd_fuzzing
            PUBLIC
            ${ANDROID_ROOT}/packages/modules/Bluetooth/system/gd
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/proto/libbluetooth-protos/android_${OUT_ARCH_CPU}_static/gen/proto/packages/modules/Bluetooth/system/gd/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/proto/libbluetooth-protos/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/rust/shim/libbt_init_flags_bridge_header/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/dumpsys/bundler/BluetoothGeneratedBundlerSchema_h_bfbs/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/BluetoothGeneratedDumpsysDataSchema_h/gen
            ${OUT_INTERMEDIATES_ROOT}/packages/modules/Bluetooth/system/gd/BluetoothGeneratedPackets_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libbluetooth_gd_fuzzing PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
            -DLOG_NDEBUG=1
            -Wno-unused-result
    )

    target_link_libraries(libbluetooth_gd_fuzzing
            android.hardware.bluetooth-V1-ndk
            android.hardware.bluetooth-1.0
            android.hardware.bluetooth-1.1
            android.system.suspend-V1-ndk
            android.system.suspend.control-V1-ndk
            jni_headers
            libbase
            libbinder_ndk
            libbluetooth-dumpsys
            libbluetooth-protos
            libbluetooth_rust_interop
            libbt-platform-protos-lite
            libbt_shim_bridge
            libbt_shim_ffi
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libcutils
            libdl
            libevent
            libflatbuffers-cpp
            libhidlbase
            liblog
            libm
            libmodpb64
            libstatslog_bt
            libutils
    )
endif ()
#################### libbluetooth_gd_fuzzing ####################

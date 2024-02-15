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

project(libchrome-crypto)

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

#################### libchrome-crypto ####################
set(FILES
        ${ANDROID_ROOT}/external/libchrome/crypto/random.cc
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
    add_library(libchrome-crypto SHARED ${SOURCE_FILES})

    target_include_directories(libchrome-crypto
            PUBLIC
            ${ANDROID_ROOT}/external/libchrome
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libchrome-include/gen/gensrcs/external/libchrome
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_jni_registration_headers/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libchrome-crypto-include/gen/gensrcs/external/libchrome
    )

    # Add c/c++ flags
    target_compile_options(libchrome-crypto PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libchrome-crypto
            jni_headers
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libchrome-crypto ####################

#################### libchrome ####################
set(FILES
        ${ANDROID_ROOT}/external/libchrome/base/at_exit.cc
        ${ANDROID_ROOT}/external/libchrome/base/barrier_closure.cc
        ${ANDROID_ROOT}/external/libchrome/base/base64.cc
        ${ANDROID_ROOT}/external/libchrome/base/base64url.cc
        ${ANDROID_ROOT}/external/libchrome/base/base_paths.cc
        ${ANDROID_ROOT}/external/libchrome/base/base_paths_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/base_switches.cc
        ${ANDROID_ROOT}/external/libchrome/base/big_endian.cc
        ${ANDROID_ROOT}/external/libchrome/base/build_time.cc
        ${ANDROID_ROOT}/external/libchrome/base/callback_helpers.cc
        ${ANDROID_ROOT}/external/libchrome/base/callback_internal.cc
        ${ANDROID_ROOT}/external/libchrome/base/command_line.cc
        ${ANDROID_ROOT}/external/libchrome/base/cpu.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/activity_tracker.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/alias.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/crash_logging.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/debugger.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/debugger_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/dump_without_crashing.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/proc_maps_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/profiler.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/stack_trace.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/task_annotator.cc
        ${ANDROID_ROOT}/external/libchrome/base/environment.cc
        ${ANDROID_ROOT}/external/libchrome/base/feature_list.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_descriptor_watcher_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_enumerator.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_enumerator_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_path.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_path_constants.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_path_watcher.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_tracing.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_util.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_util_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/important_file_writer.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/memory_mapped_file.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/memory_mapped_file_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/scoped_file.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/scoped_temp_dir.cc
        ${ANDROID_ROOT}/external/libchrome/base/guid.cc
        ${ANDROID_ROOT}/external/libchrome/base/hash.cc
        ${ANDROID_ROOT}/external/libchrome/base/json/json_file_value_serializer.cc
        ${ANDROID_ROOT}/external/libchrome/base/json/json_parser.cc
        ${ANDROID_ROOT}/external/libchrome/base/json/json_reader.cc
        ${ANDROID_ROOT}/external/libchrome/base/json/json_string_value_serializer.cc
        ${ANDROID_ROOT}/external/libchrome/base/json/json_value_converter.cc
        ${ANDROID_ROOT}/external/libchrome/base/json/json_writer.cc
        ${ANDROID_ROOT}/external/libchrome/base/json/string_escape.cc
        ${ANDROID_ROOT}/external/libchrome/base/lazy_instance_helpers.cc
        ${ANDROID_ROOT}/external/libchrome/base/location.cc
        ${ANDROID_ROOT}/external/libchrome/base/logging.cc
        ${ANDROID_ROOT}/external/libchrome/base/md5.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/aligned_memory.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/platform_shared_memory_region.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/read_only_shared_memory_region.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/ref_counted.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/ref_counted_memory.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/shared_memory_handle.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/shared_memory_helper.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/shared_memory_mapping.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/unsafe_shared_memory_region.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/weak_ptr.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/writable_shared_memory_region.cc
        ${ANDROID_ROOT}/external/libchrome/base/message_loop/incoming_task_queue.cc
        ${ANDROID_ROOT}/external/libchrome/base/message_loop/message_loop.cc
        ${ANDROID_ROOT}/external/libchrome/base/message_loop/message_loop_current.cc
        ${ANDROID_ROOT}/external/libchrome/base/message_loop/message_loop_task_runner.cc
        ${ANDROID_ROOT}/external/libchrome/base/message_loop/message_pump.cc
        ${ANDROID_ROOT}/external/libchrome/base/message_loop/message_pump_default.cc
        ${ANDROID_ROOT}/external/libchrome/base/message_loop/message_pump_libevent.cc
        ${ANDROID_ROOT}/external/libchrome/base/message_loop/watchable_io_message_pump_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/bucket_ranges.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/dummy_histogram.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/field_trial.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/field_trial_param_associator.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/histogram.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/histogram_base.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/histogram_functions.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/histogram_samples.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/histogram_snapshot_manager.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/metrics_hashes.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/persistent_histogram_allocator.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/persistent_memory_allocator.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/persistent_sample_map.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/sample_map.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/sample_vector.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/sparse_histogram.cc
        ${ANDROID_ROOT}/external/libchrome/base/metrics/statistics_recorder.cc
        ${ANDROID_ROOT}/external/libchrome/base/native_library.cc
        ${ANDROID_ROOT}/external/libchrome/base/native_library_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/observer_list_threadsafe.cc
        ${ANDROID_ROOT}/external/libchrome/base/path_service.cc
        ${ANDROID_ROOT}/external/libchrome/base/pending_task.cc
        ${ANDROID_ROOT}/external/libchrome/base/pickle.cc
        ${ANDROID_ROOT}/external/libchrome/base/posix/file_descriptor_shuffle.cc
        ${ANDROID_ROOT}/external/libchrome/base/posix/global_descriptors.cc
        ${ANDROID_ROOT}/external/libchrome/base/posix/safe_strerror.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/kill.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/kill_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/launch.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/launch_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/memory.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_handle.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_handle_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_iterator.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_metrics.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_metrics_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/rand_util.cc
        ${ANDROID_ROOT}/external/libchrome/base/rand_util_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/run_loop.cc
        ${ANDROID_ROOT}/external/libchrome/base/scoped_native_library.cc
        ${ANDROID_ROOT}/external/libchrome/base/sequence_checker_impl.cc
        ${ANDROID_ROOT}/external/libchrome/base/sequence_token.cc
        ${ANDROID_ROOT}/external/libchrome/base/sequenced_task_runner.cc
        ${ANDROID_ROOT}/external/libchrome/base/sha1.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/nullable_string16.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/pattern.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/safe_sprintf.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/string16.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/string_number_conversions.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/string_piece.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/string_split.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/string_util.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/string_util_constants.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/stringprintf.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/utf_string_conversion_utils.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/utf_string_conversions.cc
        ${ANDROID_ROOT}/external/libchrome/base/sync_socket_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/synchronization/atomic_flag.cc
        ${ANDROID_ROOT}/external/libchrome/base/synchronization/condition_variable_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/synchronization/lock.cc
        ${ANDROID_ROOT}/external/libchrome/base/synchronization/lock_impl_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/synchronization/waitable_event_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/sys_info.cc
        ${ANDROID_ROOT}/external/libchrome/base/sys_info_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/task/cancelable_task_tracker.cc
        ${ANDROID_ROOT}/external/libchrome/base/task_runner.cc
        ${ANDROID_ROOT}/external/libchrome/base/task_scheduler/scheduler_lock_impl.cc
        ${ANDROID_ROOT}/external/libchrome/base/task_scheduler/scoped_set_task_priority_for_current_thread.cc
        ${ANDROID_ROOT}/external/libchrome/base/task_scheduler/sequence.cc
        ${ANDROID_ROOT}/external/libchrome/base/task_scheduler/sequence_sort_key.cc
        ${ANDROID_ROOT}/external/libchrome/base/task_scheduler/task.cc
        ${ANDROID_ROOT}/external/libchrome/base/task_scheduler/task_traits.cc
        ${ANDROID_ROOT}/external/libchrome/base/third_party/dynamic_annotations/dynamic_annotations.c
        ${ANDROID_ROOT}/external/libchrome/base/third_party/icu/icu_utf.cc
        ${ANDROID_ROOT}/external/libchrome/base/third_party/nspr/prtime.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/platform_thread_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/post_task_and_reply_impl.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/scoped_blocking_call.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/sequence_local_storage_map.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/sequence_local_storage_slot.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/sequenced_task_runner_handle.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/simple_thread.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/thread.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/thread_checker_impl.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/thread_collision_warner.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/thread_id_name_manager.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/thread_local_storage.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/thread_local_storage_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/thread_restrictions.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/thread_task_runner_handle.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/clock.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/default_clock.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/default_tick_clock.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/tick_clock.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/time.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/time_conversion_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/time_exploded_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/time_now_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/time_override.cc
        ${ANDROID_ROOT}/external/libchrome/base/timer/elapsed_timer.cc
        ${ANDROID_ROOT}/external/libchrome/base/timer/timer.cc
        ${ANDROID_ROOT}/external/libchrome/base/unguessable_token.cc
        ${ANDROID_ROOT}/external/libchrome/base/value_iterators.cc
        ${ANDROID_ROOT}/external/libchrome/base/values.cc
        ${ANDROID_ROOT}/external/libchrome/base/version.cc
        ${ANDROID_ROOT}/external/libchrome/base/vlog.cc
        ${ANDROID_ROOT}/external/libchrome/device/bluetooth/bluetooth_advertisement.cc
        ${ANDROID_ROOT}/external/libchrome/device/bluetooth/bluetooth_uuid.cc
        ${ANDROID_ROOT}/external/libchrome/device/bluetooth/bluez/bluetooth_service_attribute_value_bluez.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/insets.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/insets_f.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/point.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/point_conversions.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/point_f.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/rect.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/rect_f.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/size.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/size_conversions.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/size_f.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/vector2d.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/geometry/vector2d_f.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/range/range.cc
        ${ANDROID_ROOT}/external/libchrome/ui/gfx/range/range_f.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_path_watcher_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/files/file_util_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/shared_memory_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/posix/unix_domain_socket.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/internal_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/memory_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_handle_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_info_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_iterator_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/process/process_metrics_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/strings/sys_string_conversions_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/sys_info_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/platform_thread_internal_posix.cc
        ${ANDROID_ROOT}/external/libchrome/base/threading/platform_thread_linux.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/android_hardware_buffer_compat.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/build_info.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/content_uri_utils.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/java_exception_reporter.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/jni_android.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/jni_array.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/jni_string.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/path_utils.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/scoped_java_ref.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/scoped_hardware_buffer_handle.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/sys_utils.cc
        ${ANDROID_ROOT}/external/libchrome/base/base_paths_android.cc
        ${ANDROID_ROOT}/external/libchrome/base/debug/stack_trace_android.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/platform_shared_memory_region_android.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/shared_memory_android.cc
        ${ANDROID_ROOT}/external/libchrome/base/memory/shared_memory_handle_android.cc
        ${ANDROID_ROOT}/external/libchrome/base/os_compat_android.cc
        ${ANDROID_ROOT}/external/libchrome/base/sys_info_android.cc
        ${ANDROID_ROOT}/external/libchrome/base/time/time_android.cc
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
    add_library(libchrome SHARED ${SOURCE_FILES})

    target_include_directories(libchrome
            PUBLIC
            ${ANDROID_ROOT}/external/libchrome
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libchrome-include/gen/gensrcs/external/libchrome
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_jni_headers/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_jni_registration_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(libchrome PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libchrome
            jni_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libevent
            libgtest_prod_headers
            liblog
            libm
            libmodpb64
    )
endif ()
#################### libchrome ####################

#################### libmojo ####################
set(FILES
        ${ANDROID_ROOT}/external/libchrome/mojo/core/broker_host.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/broker_posix.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/channel.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/channel_posix.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/configuration.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/connection_params.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/core.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/data_pipe_consumer_dispatcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/data_pipe_control_message.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/data_pipe_producer_dispatcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/dispatcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/entrypoints.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/handle_table.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/invitation_dispatcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/message_pipe_dispatcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/mojo_core.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/node_channel.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/node_controller.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/platform_handle_dispatcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/platform_handle_in_transit.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/platform_handle_utils.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/platform_shared_memory_mapping.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/request_context.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/scoped_process_handle.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/shared_buffer_dispatcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/user_message_impl.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/watch.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/watcher_dispatcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/watcher_set.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/embedder/embedder.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/embedder/scoped_ipc_support.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/ports/event.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/ports/message_queue.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/ports/name.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/ports/node.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/ports/port.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/ports/port_locker.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/ports/port_ref.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/core/ports/user_message.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/c/system/thunks.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/big_buffer.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/big_buffer_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/big_string_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/file_info_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/file_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/file_path_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/process_id_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/read_only_buffer_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/ref_counted_memory_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/shared_memory_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/string16_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/text_direction_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/thread_priority_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/time_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/unguessable_token_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/base/values_mojom_traits.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/array_internal.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/associated_binding.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/associated_group.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/associated_group_controller.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/associated_interface_ptr.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/associated_interface_ptr_state.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/binding_state.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/buffer.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/connector.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/control_message_handler.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/control_message_proxy.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/filter_chain.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/fixed_buffer.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/interface_endpoint_client.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/interface_ptr_state.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/message.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/message_dumper.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/message_header_validator.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/message_internal.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/multiplex_router.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/native_struct_serialization.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/pipe_control_message_handler.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/pipe_control_message_proxy.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/scoped_interface_endpoint_handle.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/sequence_local_sync_event_watcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/serialization_context.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/sync_call_restrictions.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/sync_event_watcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/sync_handle_registry.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/sync_handle_watcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/task_runner_helper.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/unserialized_message_context.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/validation_context.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/validation_errors.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/bindings/lib/validation_util.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/platform/named_platform_channel.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/platform/named_platform_channel_posix.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/platform/platform_channel.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/platform/platform_channel_endpoint.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/platform/platform_channel_server_endpoint.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/platform/platform_handle.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/platform/socket_utils_posix.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/buffer.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/data_pipe_drainer.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/data_pipe_utils.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/file_data_pipe_producer.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/handle_signal_tracker.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/invitation.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/isolated_connection.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/message_pipe.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/platform_handle.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/scope_to_message_pipe.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/simple_watcher.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/string_data_pipe_producer.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/trap.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/wait.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/cpp/system/wait_set.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/java/system/base_run_loop.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/java/system/core_impl.cc
        ${ANDROID_ROOT}/external/libchrome/mojo/public/java/system/watcher_impl.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/build_info.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/java_exception_reporter.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/jni_android.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/jni_array.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/jni_string.cc
        ${ANDROID_ROOT}/external/libchrome/base/android/scoped_java_ref.cc
        ${ANDROID_ROOT}/external/libchrome/ipc/ipc_message.cc
        ${ANDROID_ROOT}/external/libchrome/ipc/ipc_message_attachment.cc
        ${ANDROID_ROOT}/external/libchrome/ipc/ipc_message_attachment_set.cc
        ${ANDROID_ROOT}/external/libchrome/ipc/ipc_message_utils.cc
        ${ANDROID_ROOT}/external/libchrome/ipc/ipc_mojo_handle_attachment.cc
        ${ANDROID_ROOT}/external/libchrome/ipc/ipc_mojo_message_helper.cc
        ${ANDROID_ROOT}/external/libchrome/ipc/ipc_mojo_param_traits.cc
        ${ANDROID_ROOT}/external/libchrome/ipc/ipc_platform_file_attachment_posix.cc
        ${ANDROID_ROOT}/external/libchrome/ipc/native_handle_type_converters.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/ipc/ipc.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/ipc/ipc.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/interfaces/bindings/interface_control_messages.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/interfaces/bindings/interface_control_messages.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/interfaces/bindings/native_struct.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/interfaces/bindings/native_struct.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/interfaces/bindings/pipe_control_messages.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/interfaces/bindings/pipe_control_messages.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/big_buffer.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/big_buffer.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/big_string.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/big_string.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/file.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/file.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/file_error.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/file_error.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/file_info.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/file_info.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/file_path.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/file_path.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/process_id.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/process_id.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/read_only_buffer.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/read_only_buffer.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/ref_counted_memory.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/ref_counted_memory.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/shared_memory.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/shared_memory.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/string16.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/string16.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/text_direction.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/text_direction.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/thread_priority.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/thread_priority.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/time.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/time.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/unguessable_token.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/unguessable_token.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/values.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/mojo/public/mojom/base/values.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/ui/gfx/geometry/mojo/geometry.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/ui/gfx/geometry/mojo/geometry.mojom-shared.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/ui/gfx/range/mojo/range.mojom.cc
        ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_srcs/gen/ui/gfx/range/mojo/range.mojom-shared.cc
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
    add_library(libmojo SHARED ${SOURCE_FILES})

    target_include_directories(libmojo
            PUBLIC
            ${ANDROID_ROOT}/external/libchrome
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libchrome-include/gen/gensrcs/external/libchrome
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_jni_registration_headers/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libchrome-crypto-include/gen/gensrcs/external/libchrome
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_jni_headers/gen
            ${OUT_INTERMEDIATES_ROOT}/external/libchrome/libmojo_mojom_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(libmojo PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libmojo
            jni_headers
            libc
            libc++
            libc++demangle
            libchrome
            libchrome-crypto
            libclang_rt.builtins
            libdl
            libevent
            liblog
            libm
    )
endif ()
#################### libmojo ####################

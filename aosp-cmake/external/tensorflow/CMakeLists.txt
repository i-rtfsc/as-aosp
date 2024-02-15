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

project(TfliteNnapiDelegateTests_static)

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

#################### TfliteNnapiDelegateTests_static ####################
set(FILES
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/nnapi/nnapi_delegate_test.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/test_delegate_providers.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/test_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/acceleration_test_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/kernels/acceleration_test_util_internal.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/nnapi/acceleration_test_list.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/lite/delegates/nnapi/acceleration_test_util.cc
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
    add_library(TfliteNnapiDelegateTests_static SHARED ${SOURCE_FILES})

    target_include_directories(TfliteNnapiDelegateTests_static
            PUBLIC
            ${ANDROID_ROOT}/external/tensorflow
            ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/libtflite_mutable_schema/gen
            ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(TfliteNnapiDelegateTests_static PRIVATE
            -Werror
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(TfliteNnapiDelegateTests_static
            gemmlowp_headers
            libandroid
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libeigen
            libgmock_ndk
            libgtest_ndk_c++
            liblog
            libm
            libneuralnetworks
            libtflite_schema_headers
            libtflite_static
            tensorflow_abseil
    )
endif ()
#################### TfliteNnapiDelegateTests_static ####################

#################### tflite_configuration_proto ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tflite_configuration_proto/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/lite/experimental/acceleration/configuration/configuration.pb.cc
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
    add_library(tflite_configuration_proto SHARED ${SOURCE_FILES})

    target_include_directories(tflite_configuration_proto
            PUBLIC
            ${ANDROID_ROOT}/external/tensorflow
            ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tflite_configuration_proto/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(tflite_configuration_proto PRIVATE
            -Werror
    )

    target_link_libraries(tflite_configuration_proto
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### tflite_configuration_proto ####################

#################### libtensorflow_core_mobile_srcs_only_runtime ####################
set(FILES
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/c_api.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/c_api_function.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/checkpoint_reader.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/env.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/kernels.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/kernels_experimental.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/logging.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/ops.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/tf_buffer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/tf_datatype.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/tf_shape.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/tf_status.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/tf_status_helper.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/tf_tensor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/tf_tstring.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/eager/abstract_tensor_handle.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/eager/c_api.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/eager/c_api_debug.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/c/eager/immediate_execution_tensor_handle.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/accumulate_n_optimizer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/all_to_all.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/allocator_retry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/base_collective_executor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/bfc_allocator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/buf_rendezvous.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/build_graph_options.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/collective_executor_mgr.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/collective_param_resolver_local.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/collective_rma_local.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/collective_test_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/collective_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/colocation_graph.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/composite_device.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/constant_folding.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/control_flow_deps_to_chains.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/copy_tensor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/cost_measurement_registry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/cost_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/costmodel_manager.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/debugger_state_interface.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/device_mgr.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/device_propagation.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/device_resolver_local.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/device_set.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/direct_session.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/dynamic_device_mgr.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eval_const_tensor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/executor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/executor_factory.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/forward_type_inference.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/function.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/function_body.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/function_def_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/function_optimization_registry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/function_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/gradients.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/graph_constructor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/graph_def_builder_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/graph_execution_state.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/graph_optimizer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/graph_runner.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/graph_view.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/hierarchical_tree_broadcaster.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/immutable_executor_state.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/inline_function_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/input_colocation_exemption_registry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/inspecting_placer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/isolate_placer_inspection_required_ops_pass.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/local_device.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/local_session_selection.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/lower_case_op.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/lower_function_call_inline_policy.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/lower_function_call_op.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/lower_functional_ops.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/lower_if_op.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/lower_while_op.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/memory_types.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/mkl_cpu_allocator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/mkl_layout_pass.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/mkl_tfconversion_pass.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/no_op_cost_measurement.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/node_file_writer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/null_request_cost_accessor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/optimization_registry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/parallel_concat_optimizer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/partitioning_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/permuter.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/placer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/placer_inspection_required_ops_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/pool_allocator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/process_function_library_runtime.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/process_state.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/process_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/propagator_debug_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/propagator_state.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/quantize_training.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/renamed_device.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/rendezvous_mgr.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/rendezvous_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/replicate_per_replica_nodes.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/request_cost.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/request_cost_accessor_registry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/ring_alg.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/ring_gatherer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/ring_reducer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/scoped_allocator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/scoped_allocator_mgr.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/session.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/session_factory.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/session_options.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/session_state.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/shape_refiner.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/simple_propagator_state.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/single_threaded_cpu_device.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/single_threaded_executor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/stats_publisher_interface.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/step_stats_collector.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/threadpool_device.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/threadpool_device_factory.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/device/device_event_mgr.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/device/device_id_manager.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/device/device_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/attr_builder.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/context.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/context_distributed_manager.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/core.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/custom_device.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/custom_device_op_handler.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/eager_executor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/eager_op_rewrite_registry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/eager_operation.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/execute.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/execute_node.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/kernel_and_device.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/mkl_eager_op_rewrite.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/placement_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/shape_inference.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/tensor_handle.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/common_runtime/eager/tensor_handle_data.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/attr_value.pb_text.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/attr_value_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/cancellation.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/collective.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/common_shape_fns.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/dataset.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/device.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/device_base.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/device_factory.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/full_type_inference_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/full_type_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/function.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/function_handle_cache.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/graph_def_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/graph_to_functiondef.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/kernel_def_builder.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/kernel_def_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/load_library.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/local_rendezvous.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/logging.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/lookup_interface.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/memory_types.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/metrics.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/model.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/node_def_builder.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/node_def_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/node_properties.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/op.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/op_def_builder.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/op_def_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/op_kernel.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/op_segment.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/reader_base.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/ref_var.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/rendezvous.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/resource_mgr.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/resource_var.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/run_handler.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/run_handler_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/shape_inference.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/tensor_slice.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/tensor_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/types.pb_text.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/versions.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/allocator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/allocator_registry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/bfloat16.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/cpu_allocator_impl.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/kernel_shape_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/log_memory.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/ops_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/resource_handle.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/resource_handle.pb_text.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/tensor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/tensor.pb_text.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/tensor_shape.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/tensor_shape.pb_text.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/tracking_allocator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/typed_allocator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/types.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/variant.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/variant_op_registry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/framework/variant_tensor_data.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/algorithm.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/collective_order.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/colors.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/control_flow.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/costmodel.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/edgeset.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/graph.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/graph_def_builder.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/graph_node_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/graph_partition.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/node_builder.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/optimizer_cse.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/subgraph.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/tensor_id.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/validate.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/graph/while_context.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/padding.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/tensor_format.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/batch_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/bcast.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/command_line_flags.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/determinism.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/device_name_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/dump_graph.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/einsum_op_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/env_var.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/equal_graph_def.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/example_proto_fast_parsing.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/example_proto_helper.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/guarded_philox_random.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/matmul_autotune.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/mirror_pad_mode.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/port.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/ragged_to_dense_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/saved_tensor_slice_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/stat_summarizer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/strided_slice_op.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/tensor_slice_reader.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/tensor_slice_reader_cache.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/tensor_slice_set.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/tensor_slice_writer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/use_cudnn.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/work_sharder.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/sparse/group_iterator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/sparse/sparse_tensor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/tensor_bundle/naming.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/util/tensor_bundle/tensor_bundle.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/abi.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/coding.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/cpu_info.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/denormal.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/env.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/errors.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/file_system.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/file_system_helper.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/hash.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/numbers.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/path.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/protobuf.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/protobuf_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/scanner.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/setround.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/status.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/statusor.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/str_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/strcat.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/stringprintf.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/tensor_coding.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/tracing.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/threadpool.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/base64.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/cpu_feature_guard.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/platform_strings.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/random.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/profile_utils/android_armv7a_cpu_utils_helper.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/platform/profile_utils/cpu_utils.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/profiler/lib/scoped_memory_debug_annotation.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/tfrt/fallback/op_kernel_runner.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/env.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/env_time.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/load_library.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/port.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/posix_file_system.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/resource.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/logging.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/subprocess.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/tsl/platform/default/unbounded_work_queue.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/block.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/block_builder.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/buffered_inputstream.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/cache.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/compression.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/format.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/inputbuffer.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/inputstream_interface.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/iterator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/random_inputstream.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/record_reader.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/table.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/table_builder.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/two_level_iterator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/zlib_compression_options.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/io/zlib_inputstream.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/strings/proto_serialization.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/strings/ordered_code.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/monitoring/collection_registry.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/core/arena.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/hash/crc32c.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/hash/crc32c_accelerate.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/histogram/histogram.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/random/distribution_sampler.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/random/simple_philox.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/random/weighted_picker.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/strings/proto_text_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/nccl/collective_communicator.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/example/feature_util.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/core/lib/wav/wav_io.cc
        ${ANDROID_ROOT}/external/tensorflow/tensorflow/compiler/jit/defs.cc
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
    add_library(libtensorflow_core_mobile_srcs_only_runtime SHARED ${SOURCE_FILES})

    target_include_directories(libtensorflow_core_mobile_srcs_only_runtime
            PUBLIC
            ${ANDROID_ROOT}/external/tensorflow
            ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto
            ${ANDROID_ROOT}/external/tensorflow/third_party/absl/abseil-cpp
    )

    # Add c/c++ flags
    target_compile_options(libtensorflow_core_mobile_srcs_only_runtime PRIVATE
            -Werror
            -Wno-unused-const-variable
            -Wno-unused-but-set-variable
            -Wno-unused-parameter
    )

    target_link_libraries(libtensorflow_core_mobile_srcs_only_runtime
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libdoubleconversion
            libeigen
            libm
            libtextclassifier_hash_headers
            libz
            tensorflow_abseil
            tensorflow_core_proto_cpp_lite
            tensorflow_headers
    )
endif ()
#################### libtensorflow_core_mobile_srcs_only_runtime ####################

#################### tensorflow_core_proto_cpp_lite ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/allocation_description.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/api_def.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/attr_value.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/cost_graph.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/dataset_metadata.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/dataset_options.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/device_attributes.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/full_type.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/function.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/graph.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/graph_transfer_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/kernel_def.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/log_memory.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/model.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/node_def.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/op_def.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/reader_base.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/resource_handle.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/step_stats.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/summary.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/tensor.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/tensor_description.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/tensor_shape.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/tensor_slice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/types.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/variable.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/framework/versions.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/example/example.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/example/feature.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/example/example_parser_configuration.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/profiler/protobuf/xplane.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/profiler/profiler_options.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/util/saved_tensor_slice.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/util/event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/bfc_memory_map.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/config.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/coordination_config.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/coordination_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/cluster.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/debug.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/device_filters.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/device_properties.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/data_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/debug_event.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/eager_service.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/graph_debug_info.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/queue_runner.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/rewriter_config.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/tensor_bundle.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/saver.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/verifier_config.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/error_codes.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/composite_tensor_variant.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/master.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/meta_graph.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/named_tensor.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/remote_tensor_handle.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/saved_model.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/saved_object_graph.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/status.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/struct.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/worker.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/tensorflow_server.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/trackable_object_graph.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/protobuf/transport_options.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/core/grappler/costs/op_performance_data.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/compiler/xla/pjrt/distributed/protocol.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/tensorflow/compiler/xla/stream_executor/dnn.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/descriptor.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/any.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/source_context.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/type.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/compiler/plugin.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/api.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/duration.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/empty.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/field_mask.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/struct.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/timestamp.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto/google/protobuf/wrappers.pb.cc
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
    add_library(tensorflow_core_proto_cpp_lite SHARED ${SOURCE_FILES})

    target_include_directories(tensorflow_core_proto_cpp_lite
            PUBLIC
            ${ANDROID_ROOT}/external/tensorflow
            ${OUT_INTERMEDIATES_ROOT}/external/tensorflow/tensorflow_core_proto_cpp_lite/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(tensorflow_core_proto_cpp_lite PRIVATE
            -Werror
    )

    target_link_libraries(tensorflow_core_proto_cpp_lite
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### tensorflow_core_proto_cpp_lite ####################

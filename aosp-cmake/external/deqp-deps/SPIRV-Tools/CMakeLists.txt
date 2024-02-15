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

project(deqp_spirv-tools)

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

#################### deqp_spirv-tools ####################
set(FILES
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/assembly_grammar.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/binary.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/diagnostic.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/disassemble.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/enum_string_mapping.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/ext_inst.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/extensions.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/libspirv.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/name_mapper.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opcode.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/operand.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/parsed_operand.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/pch_source.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/print.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/software_version.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/spirv_endian.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/spirv_fuzzer_options.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/spirv_optimizer_options.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/spirv_reducer_options.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/spirv_target_env.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/spirv_validator_options.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/table.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/text.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/text_handler.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/util/bit_vector.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/util/parse_number.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/util/string_utils.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/util/timer.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/basic_block.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/construct.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/function.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/instruction.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_adjacency.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_annotation.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_arithmetics.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_atomics.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_barriers.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_bitwise.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_builtins.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_capability.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_cfg.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_composites.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_constants.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_conversion.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_debug.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_decorations.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_derivatives.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_execution_limitations.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_extensions.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_function.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_id.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_image.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_instruction.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_interfaces.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_layout.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_literals.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_logicals.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_memory.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_memory_semantics.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_misc.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_mode_setting.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_non_uniform.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_primitives.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_ray_query.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_ray_tracing.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_scopes.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_small_type_uses.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validate_type.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/val/validation_state.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/aggressive_dead_code_elim_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/amd_ext_to_khr.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/basic_block.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/block_merge_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/block_merge_util.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/build_module.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/ccp_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/cfg.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/cfg_cleanup_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/code_sink.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/combine_access_chains.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/compact_ids_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/composite.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/const_folding_rules.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/constants.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/control_dependence.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/convert_to_half_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/convert_to_sampled_image_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/copy_prop_arrays.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/dataflow.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/dead_branch_elim_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/dead_insert_elim_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/dead_variable_elimination.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/debug_info_manager.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/decoration_manager.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/def_use_manager.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/desc_sroa.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/desc_sroa_util.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/dominator_analysis.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/dominator_tree.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/eliminate_dead_constant_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/eliminate_dead_functions_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/eliminate_dead_functions_util.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/eliminate_dead_input_components_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/eliminate_dead_members_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/feature_manager.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/fix_func_call_arguments.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/fix_storage_class.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/flatten_decoration_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/fold.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/fold_spec_constant_op_and_composite_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/folding_rules.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/freeze_spec_constant_value_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/function.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/graphics_robust_access_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/if_conversion.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/inline_exhaustive_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/inline_opaque_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/inline_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/inst_bindless_check_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/inst_buff_addr_check_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/inst_debug_printf_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/instruction.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/instruction_list.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/instrument_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/interface_var_sroa.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/interp_fixup_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/ir_context.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/ir_loader.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/licm_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/local_access_chain_convert_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/local_redundancy_elimination.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/local_single_block_elim_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/local_single_store_elim_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_dependence.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_dependence_helpers.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_descriptor.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_fission.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_fusion.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_fusion_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_peeling.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_unroller.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_unswitch_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/loop_utils.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/mem_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/merge_return_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/module.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/optimizer.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/pass_manager.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/pch_source_opt.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/private_to_local_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/propagator.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/reduce_load_size.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/redundancy_elimination.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/register_pressure.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/relax_float_ops_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/remove_dontinline_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/remove_duplicates_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/remove_unused_interface_variables_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/replace_desc_array_access_using_var_index.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/replace_invalid_opc.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/scalar_analysis.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/scalar_analysis_simplification.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/scalar_replacement_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/set_spec_constant_default_value_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/simplification_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/spread_volatile_semantics.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/ssa_rewrite_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/strength_reduction_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/strip_debug_info_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/strip_nonsemantic_info_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/struct_cfg_analysis.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/type_manager.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/types.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/unify_const_pass.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/upgrade_memory_model.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/value_number_table.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/vector_dce.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/workaround1209.cpp
        ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/source/opt/wrap_opkill.cpp
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
    add_library(deqp_spirv-tools SHARED ${SOURCE_FILES})

    target_include_directories(deqp_spirv-tools
            PUBLIC
            ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools/include
            ${ANDROID_ROOT}/external/deqp-deps/SPIRV-Tools
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/SPIRV-Tools/deqp_spvtools_generate_grammar_tables/gen
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/SPIRV-Tools/deqp_spvtools_generate_language_headers/gen
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/SPIRV-Tools/deqp_spvtools_update_build_version/gen
            ${OUT_INTERMEDIATES_ROOT}/external/deqp-deps/SPIRV-Tools/deqp_spvtools_generate_registry_tables/gen
    )

    # Add c/c++ flags
    target_compile_options(deqp_spirv-tools PRIVATE
            -DAMBER_ENABLE_VK_DEBUGGING=0
            -Wall
            -Werror
    )

    target_link_libraries(deqp_spirv-tools
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### deqp_spirv-tools ####################

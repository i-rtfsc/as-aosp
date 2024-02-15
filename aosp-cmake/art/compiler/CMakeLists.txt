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

project(libart-compiler)

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

#################### libart-compiler ####################
set(FILES
        ${ANDROID_ROOT}/art/compiler/debug/elf_debug_writer.cc
        ${ANDROID_ROOT}/art/compiler/dex/inline_method_analyser.cc
        ${ANDROID_ROOT}/art/compiler/driver/compiler_options.cc
        ${ANDROID_ROOT}/art/compiler/driver/dex_compilation_unit.cc
        ${ANDROID_ROOT}/art/compiler/jit/jit_compiler.cc
        ${ANDROID_ROOT}/art/compiler/jit/jit_logger.cc
        ${ANDROID_ROOT}/art/compiler/jni/quick/calling_convention.cc
        ${ANDROID_ROOT}/art/compiler/jni/quick/jni_compiler.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/block_builder.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/block_namer.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/bounds_check_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/builder.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/cha_guard_optimization.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_utils.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_sinking.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/constant_folding.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/constructor_fence_redundancy_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/data_type.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/dead_code_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/escape.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/execution_subgraph.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/graph_checker.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/graph_visualizer.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/gvn.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/induction_var_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/induction_var_range.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/inliner.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_builder.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_simplifier.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/intrinsic_objects.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/intrinsics.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/licm.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/linear_order.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/load_store_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/load_store_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/locations.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/loop_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/loop_optimization.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/nodes.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/optimization.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/optimizing_compiler.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/parallel_move_resolver.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/prepare_for_register_allocation.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/reference_type_propagation.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/register_allocation_resolver.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/register_allocator.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/register_allocator_graph_color.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/register_allocator_linear_scan.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/select_generator.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/scheduler.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/sharpening.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/side_effects_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/ssa_builder.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/ssa_liveness_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/ssa_phi_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/stack_map_stream.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/superblock_cloner.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/write_barrier_elimination.cc
        ${ANDROID_ROOT}/art/compiler/trampolines/trampoline_compiler.cc
        ${ANDROID_ROOT}/art/compiler/utils/assembler.cc
        ${ANDROID_ROOT}/art/compiler/utils/jni_macro_assembler.cc
        ${ANDROID_ROOT}/art/compiler/compiler.cc
        ${ANDROID_ROOT}/art/compiler/jni/quick/arm/calling_convention_arm.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_vector_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/critical_native_abi_fixup_arm.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_simplifier_arm.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_simplifier_shared.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/intrinsics_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/nodes_shared.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/scheduler_arm.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm/assembler_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm/constants_arm.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm/jni_macro_assembler_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm/managed_register_arm.cc
        ${ANDROID_ROOT}/art/compiler/jni/quick/arm64/calling_convention_arm64.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_arm64.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_vector_arm64_neon.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_vector_arm64_sve.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/scheduler_arm64.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_simplifier_arm64.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/intrinsics_arm64.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm64/assembler_arm64.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm64/jni_macro_assembler_arm64.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm64/managed_register_arm64.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/linker/linker_patch.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/optimizing/locations.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/optimizing/nodes.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/optimizing/optimizing_compiler_stats.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/utils/arm/constants_arm.operator_out.cc
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
    add_library(libart-compiler SHARED ${SOURCE_FILES})

    target_include_directories(libart-compiler
            PUBLIC
            ${ANDROID_ROOT}/art/compiler
            ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libart-compiler PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libart-compiler
            art_cmdlineparser_headers
            art_disassembler_headers
            libart
            libartbase
            libartpalette
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfile
            libdl
            libelffile
            liblzma
            libm
            libnativehelper_header_only
            libprofile
            libvixl
    )
endif ()
#################### libart-compiler ####################

#################### libartd-compiler ####################
set(FILES
        ${ANDROID_ROOT}/art/compiler/debug/elf_debug_writer.cc
        ${ANDROID_ROOT}/art/compiler/dex/inline_method_analyser.cc
        ${ANDROID_ROOT}/art/compiler/driver/compiler_options.cc
        ${ANDROID_ROOT}/art/compiler/driver/dex_compilation_unit.cc
        ${ANDROID_ROOT}/art/compiler/jit/jit_compiler.cc
        ${ANDROID_ROOT}/art/compiler/jit/jit_logger.cc
        ${ANDROID_ROOT}/art/compiler/jni/quick/calling_convention.cc
        ${ANDROID_ROOT}/art/compiler/jni/quick/jni_compiler.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/block_builder.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/block_namer.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/bounds_check_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/builder.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/cha_guard_optimization.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_utils.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_sinking.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/constant_folding.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/constructor_fence_redundancy_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/data_type.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/dead_code_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/escape.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/execution_subgraph.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/graph_checker.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/graph_visualizer.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/gvn.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/induction_var_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/induction_var_range.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/inliner.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_builder.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_simplifier.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/intrinsic_objects.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/intrinsics.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/licm.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/linear_order.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/load_store_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/load_store_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/locations.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/loop_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/loop_optimization.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/nodes.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/optimization.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/optimizing_compiler.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/parallel_move_resolver.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/prepare_for_register_allocation.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/reference_type_propagation.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/register_allocation_resolver.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/register_allocator.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/register_allocator_graph_color.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/register_allocator_linear_scan.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/select_generator.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/scheduler.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/sharpening.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/side_effects_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/ssa_builder.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/ssa_liveness_analysis.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/ssa_phi_elimination.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/stack_map_stream.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/superblock_cloner.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/write_barrier_elimination.cc
        ${ANDROID_ROOT}/art/compiler/trampolines/trampoline_compiler.cc
        ${ANDROID_ROOT}/art/compiler/utils/assembler.cc
        ${ANDROID_ROOT}/art/compiler/utils/jni_macro_assembler.cc
        ${ANDROID_ROOT}/art/compiler/compiler.cc
        ${ANDROID_ROOT}/art/compiler/jni/quick/arm/calling_convention_arm.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_vector_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/critical_native_abi_fixup_arm.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_simplifier_arm.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_simplifier_shared.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/intrinsics_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/nodes_shared.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/scheduler_arm.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm/assembler_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm/constants_arm.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm/jni_macro_assembler_arm_vixl.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm/managed_register_arm.cc
        ${ANDROID_ROOT}/art/compiler/jni/quick/arm64/calling_convention_arm64.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_arm64.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_vector_arm64_neon.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/code_generator_vector_arm64_sve.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/scheduler_arm64.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/instruction_simplifier_arm64.cc
        ${ANDROID_ROOT}/art/compiler/optimizing/intrinsics_arm64.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm64/assembler_arm64.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm64/jni_macro_assembler_arm64.cc
        ${ANDROID_ROOT}/art/compiler/utils/arm64/managed_register_arm64.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/linker/linker_patch.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/optimizing/locations.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/optimizing/nodes.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/optimizing/optimizing_compiler_stats.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs/art/compiler/utils/arm/constants_arm.operator_out.cc
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
    add_library(libartd-compiler SHARED ${SOURCE_FILES})

    target_include_directories(libartd-compiler
            PUBLIC
            ${ANDROID_ROOT}/art/compiler
            ${OUT_INTERMEDIATES_ROOT}/art/compiler/art_compiler_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libartd-compiler PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(libartd-compiler
            art_cmdlineparser_headers
            art_disassembler_headers
            libartbased
            libartd
            libartpalette
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfiled
            libdl
            libelffiled
            liblzma
            libm
            libnativehelper_header_only
            libprofiled
            libvixld
    )
endif ()
#################### libartd-compiler ####################

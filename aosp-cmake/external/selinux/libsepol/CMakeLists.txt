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

project(libsepol)

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

#################### libsepol ####################
set(FILES
        ${ANDROID_ROOT}/external/selinux/libsepol/src/assertion.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/avrule_block.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/avtab.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/boolean_record.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/booleans.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/conditional.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/constraint.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/context.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/context_record.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/debug.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/ebitmap.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/expand.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/handle.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/hashtab.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/hierarchy.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/iface_record.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/interfaces.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/kernel_to_cil.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/kernel_to_common.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/kernel_to_conf.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/link.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/mls.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/module.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/module_to_cil.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/node_record.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/nodes.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/optimize.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/polcaps.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/policydb.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/policydb_convert.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/policydb_public.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/policydb_validate.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/port_record.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/ports.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/services.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/sidtab.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/symtab.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/user_record.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/users.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/util.c
        ${ANDROID_ROOT}/external/selinux/libsepol/src/write.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/android.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_binary.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_build_ast.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_copy_ast.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_find.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_fqn.c
        ${OUT_INTERMEDIATES_ROOT}/external/selinux/libsepol/libsepol/android_${OUT_ARCH_CPU}_static/gen/lex/external/selinux/libsepol/cil/src/cil_lexer.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_list.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_log.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_mem.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_parser.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_policy.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_post.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_reset_ast.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_resolve_ast.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_stack.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_strpool.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_symtab.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_tree.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_verify.c
        ${ANDROID_ROOT}/external/selinux/libsepol/cil/src/cil_write_ast.c
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
    add_library(libsepol SHARED ${SOURCE_FILES})

    target_include_directories(libsepol
            PUBLIC
            ${ANDROID_ROOT}/external/selinux/libsepol/cil/include
            ${ANDROID_ROOT}/external/selinux/libsepol/include
            ${ANDROID_ROOT}/external/selinux/libsepol/cil/src
            ${ANDROID_ROOT}/external/selinux/libsepol/src
            ${ANDROID_ROOT}/external/selinux/libsepol
    )

    # Add c/c++ flags
    target_compile_options(libsepol PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libsepol
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libsepol ####################

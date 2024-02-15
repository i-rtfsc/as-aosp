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

project(libdexfile_support_static)

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

#################### libdexfile_support_static ####################
set(FILES
        ${ANDROID_ROOT}/art/libdexfile/external/dex_file_supp.cc
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
    add_library(libdexfile_support_static SHARED ${SOURCE_FILES})

    target_include_directories(libdexfile_support_static
            PUBLIC
            ${ANDROID_ROOT}/art/libdexfile
            ${ANDROID_ROOT}/art/libdexfile/external/include
    )

    # Add c/c++ flags
    target_compile_options(libdexfile_support_static PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libdexfile_support_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfile_external_headers
            libdl
            libm
    )
endif ()
#################### libdexfile_support_static ####################

#################### libdexfile_support ####################
set(FILES
        ${ANDROID_ROOT}/art/libdexfile/external/dex_file_supp.cc
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
    add_library(libdexfile_support SHARED ${SOURCE_FILES})

    target_include_directories(libdexfile_support
            PUBLIC
            ${ANDROID_ROOT}/art/libdexfile
            ${ANDROID_ROOT}/art/libdexfile/external/include
    )

    # Add c/c++ flags
    target_compile_options(libdexfile_support PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libdexfile_support
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfile_external_headers
            libdl
            liblog
            libm
    )
endif ()
#################### libdexfile_support ####################

#################### libdexfile ####################
set(FILES
        ${ANDROID_ROOT}/art/libdexfile/dex/art_dex_file_loader.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/compact_dex_file.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/compact_offset_table.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/descriptors_names.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_exception_helpers.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_layout.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_loader.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_tracking_registrar.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_verifier.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_instruction.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/modifiers.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/primitive.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/signature.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/standard_dex_file.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/type_lookup_table.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/utf.cc
        ${ANDROID_ROOT}/art/libdexfile/external/dex_file_ext.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/dex_file.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/dex_file_layout.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/dex_instruction.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/dex_instruction_utils.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/invoke_type.operator_out.cc
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
    add_library(libdexfile SHARED ${SOURCE_FILES})

    target_include_directories(libdexfile
            PUBLIC
            ${ANDROID_ROOT}/art/libdexfile
            ${ANDROID_ROOT}/art/libdexfile/external/include
            ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libdexfile PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libdexfile
            jni_headers
            libartbase
            libartpalette
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdexfile_external_headers
            libdl
            liblog
            libm
            libwinpthread
            libz
            libziparchive
    )
endif ()
#################### libdexfile ####################

#################### libdexfiled ####################
set(FILES
        ${ANDROID_ROOT}/art/libdexfile/dex/art_dex_file_loader.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/compact_dex_file.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/compact_offset_table.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/descriptors_names.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_exception_helpers.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_layout.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_loader.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_tracking_registrar.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_file_verifier.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/dex_instruction.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/modifiers.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/primitive.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/signature.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/standard_dex_file.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/type_lookup_table.cc
        ${ANDROID_ROOT}/art/libdexfile/dex/utf.cc
        ${ANDROID_ROOT}/art/libdexfile/external/dex_file_ext.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/dex_file.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/dex_file_layout.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/dex_instruction.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/dex_instruction_utils.operator_out.cc
        ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs/art/libdexfile/dex/invoke_type.operator_out.cc
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
    add_library(libdexfiled SHARED ${SOURCE_FILES})

    target_include_directories(libdexfiled
            PUBLIC
            ${ANDROID_ROOT}/art/libdexfile
            ${ANDROID_ROOT}/art/libdexfile/external/include
            ${OUT_INTERMEDIATES_ROOT}/art/libdexfile/dexfile_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libdexfiled PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(libdexfiled
            jni_headers
            libartbased
            libartpalette
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdexfile_external_headers
            libdl
            liblog
            libm
            libwinpthread
            libz
            libziparchive
    )
endif ()
#################### libdexfiled ####################

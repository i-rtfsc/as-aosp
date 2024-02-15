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

project(dex2oat)

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

#################### libart-dex2oat ####################
set(FILES
        ${ANDROID_ROOT}/art/dex2oat/dex/quick_compiler_callbacks.cc
        ${ANDROID_ROOT}/art/dex2oat/dex/verification_results.cc
        ${ANDROID_ROOT}/art/dex2oat/driver/compiled_method.cc
        ${ANDROID_ROOT}/art/dex2oat/driver/compiled_method_storage.cc
        ${ANDROID_ROOT}/art/dex2oat/driver/compiler_driver.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/code_info_table_deduper.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/elf_writer.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/elf_writer_quick.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/image_writer.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/multi_oat_relative_patcher.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/oat_writer.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/relative_patcher.cc
        ${ANDROID_ROOT}/art/dex2oat/utils/swap_space.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/arm/relative_patcher_arm_base.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/arm/relative_patcher_thumb2.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/arm64/relative_patcher_arm64.cc
        ${OUT_INTERMEDIATES_ROOT}/art/dex2oat/art_dex2oat_operator_srcs/gen/gensrcs/art/dex2oat/linker/image_writer.operator_out.cc
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
    add_library(libart-dex2oat SHARED ${SOURCE_FILES})

    target_include_directories(libart-dex2oat
            PUBLIC
            ${ANDROID_ROOT}/art/dex2oat
            ${OUT_INTERMEDIATES_ROOT}/art/dex2oat/art_dex2oat_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libart-dex2oat PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libart-dex2oat
            libart
            libart-compiler
            libart-dexlayout
            libartpalette
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libelffile
            liblog
            liblz4
            libm
            libprofile
            libz
    )
endif ()
#################### libart-dex2oat ####################

#################### libartd-dex2oat ####################
set(FILES
        ${ANDROID_ROOT}/art/dex2oat/dex/quick_compiler_callbacks.cc
        ${ANDROID_ROOT}/art/dex2oat/dex/verification_results.cc
        ${ANDROID_ROOT}/art/dex2oat/driver/compiled_method.cc
        ${ANDROID_ROOT}/art/dex2oat/driver/compiled_method_storage.cc
        ${ANDROID_ROOT}/art/dex2oat/driver/compiler_driver.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/code_info_table_deduper.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/elf_writer.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/elf_writer_quick.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/image_writer.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/multi_oat_relative_patcher.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/oat_writer.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/relative_patcher.cc
        ${ANDROID_ROOT}/art/dex2oat/utils/swap_space.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/arm/relative_patcher_arm_base.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/arm/relative_patcher_thumb2.cc
        ${ANDROID_ROOT}/art/dex2oat/linker/arm64/relative_patcher_arm64.cc
        ${OUT_INTERMEDIATES_ROOT}/art/dex2oat/art_dex2oat_operator_srcs/gen/gensrcs/art/dex2oat/linker/image_writer.operator_out.cc
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
    add_library(libartd-dex2oat SHARED ${SOURCE_FILES})

    target_include_directories(libartd-dex2oat
            PUBLIC
            ${ANDROID_ROOT}/art/dex2oat
            ${OUT_INTERMEDIATES_ROOT}/art/dex2oat/art_dex2oat_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libartd-dex2oat PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(libartd-dex2oat
            libartd
            libartd-compiler
            libartd-dexlayout
            libartpalette
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcrypto
            libdl
            libelffiled
            liblog
            liblz4
            libm
            libprofiled
            libz
    )
endif ()
#################### libartd-dex2oat ####################

#################### dex2oatd ####################
set(FILES
        ${ANDROID_ROOT}/art/dex2oat/dex2oat_options.cc
        ${ANDROID_ROOT}/art/dex2oat/dex2oat.cc
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
    add_executable(dex2oatd ${SOURCE_FILES})

    target_include_directories(dex2oatd
            PUBLIC
            ${ANDROID_ROOT}/art/dex2oat
    )

    # Add c/c++ flags
    target_compile_options(dex2oatd PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(dex2oatd
            art_cmdlineparser_headers
            libc++_static
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdex2oatd_static
    )
endif ()
#################### dex2oatd ####################

#################### dex2oat ####################
set(FILES
        ${ANDROID_ROOT}/art/dex2oat/dex2oat_options.cc
        ${ANDROID_ROOT}/art/dex2oat/dex2oat.cc
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
    add_executable(dex2oat ${SOURCE_FILES})

    target_include_directories(dex2oat
            PUBLIC
            ${ANDROID_ROOT}/art/dex2oat
    )

    # Add c/c++ flags
    target_compile_options(dex2oat PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(dex2oat
            art_cmdlineparser_headers
            libc++_static
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdex2oat_static
    )
endif ()
#################### dex2oat ####################

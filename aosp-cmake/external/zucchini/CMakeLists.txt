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

project(zucchini)

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

#################### zucchini ####################
set(FILES
        ${ANDROID_ROOT}/external/zucchini/main_utils.cc
        ${ANDROID_ROOT}/external/zucchini/zucchini_commands.cc
        ${ANDROID_ROOT}/external/zucchini/zucchini_integration.cc
        ${ANDROID_ROOT}/external/zucchini/zucchini_main_aosp.cc
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
    add_executable(zucchini ${SOURCE_FILES})

    target_include_directories(zucchini
            PUBLIC
            ${ANDROID_ROOT}/external/zucchini/aosp/include
            ${ANDROID_ROOT}/external/zucchini/aosp/include/components
    )

    # Add c/c++ flags
    target_compile_options(zucchini PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(zucchini
            libbase
            libc
            libc++
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libzucchini
    )
endif ()
#################### zucchini ####################

#################### libzucchini ####################
set(FILES
        ${ANDROID_ROOT}/external/zucchini/abs32_utils.cc
        ${ANDROID_ROOT}/external/zucchini/address_translator.cc
        ${ANDROID_ROOT}/external/zucchini/arm_utils.cc
        ${ANDROID_ROOT}/external/zucchini/binary_data_histogram.cc
        ${ANDROID_ROOT}/external/zucchini/buffer_sink.cc
        ${ANDROID_ROOT}/external/zucchini/buffer_source.cc
        ${ANDROID_ROOT}/external/zucchini/crc32.cc
        ${ANDROID_ROOT}/external/zucchini/disassembler.cc
        ${ANDROID_ROOT}/external/zucchini/disassembler_dex.cc
        ${ANDROID_ROOT}/external/zucchini/disassembler_elf.cc
        ${ANDROID_ROOT}/external/zucchini/disassembler_no_op.cc
        ${ANDROID_ROOT}/external/zucchini/disassembler_win32.cc
        ${ANDROID_ROOT}/external/zucchini/disassembler_ztf.cc
        ${ANDROID_ROOT}/external/zucchini/element_detection.cc
        ${ANDROID_ROOT}/external/zucchini/encoded_view.cc
        ${ANDROID_ROOT}/external/zucchini/ensemble_matcher.cc
        ${ANDROID_ROOT}/external/zucchini/equivalence_map.cc
        ${ANDROID_ROOT}/external/zucchini/heuristic_ensemble_matcher.cc
        ${ANDROID_ROOT}/external/zucchini/image_index.cc
        ${ANDROID_ROOT}/external/zucchini/imposed_ensemble_matcher.cc
        ${ANDROID_ROOT}/external/zucchini/io_utils.cc
        ${ANDROID_ROOT}/external/zucchini/mapped_file.cc
        ${ANDROID_ROOT}/external/zucchini/patch_reader.cc
        ${ANDROID_ROOT}/external/zucchini/patch_writer.cc
        ${ANDROID_ROOT}/external/zucchini/reference_bytes_mixer.cc
        ${ANDROID_ROOT}/external/zucchini/reference_set.cc
        ${ANDROID_ROOT}/external/zucchini/rel32_finder.cc
        ${ANDROID_ROOT}/external/zucchini/rel32_utils.cc
        ${ANDROID_ROOT}/external/zucchini/reloc_elf.cc
        ${ANDROID_ROOT}/external/zucchini/reloc_win32.cc
        ${ANDROID_ROOT}/external/zucchini/target_pool.cc
        ${ANDROID_ROOT}/external/zucchini/targets_affinity.cc
        ${ANDROID_ROOT}/external/zucchini/zucchini_apply.cc
        ${ANDROID_ROOT}/external/zucchini/zucchini_gen.cc
        ${ANDROID_ROOT}/external/zucchini/zucchini_tools.cc
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
    add_library(libzucchini SHARED ${SOURCE_FILES})

    target_include_directories(libzucchini
            PUBLIC
            ${ANDROID_ROOT}/external/zucchini/aosp/include/components
            ${ANDROID_ROOT}/external/zucchini/aosp/include
    )

    # Add c/c++ flags
    target_compile_options(libzucchini PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libzucchini
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libchrome
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### libzucchini ####################

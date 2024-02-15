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

project(libopencsd_decoder)

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

#################### libopencsd_decoder ####################
set(FILES
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ocsd_code_follower.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ocsd_dcd_tree.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ocsd_error.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ocsd_error_logger.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ocsd_gen_elem_list.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ocsd_gen_elem_stack.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ocsd_lib_dcd_register.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ocsd_msg_logger.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ocsd_version.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/trc_component.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/trc_core_arch_map.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/trc_frame_deformatter.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/trc_gen_elem.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/trc_printable_elem.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/trc_ret_stack.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv3/trc_cmp_cfg_etmv3.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv3/trc_pkt_decode_etmv3.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv3/trc_pkt_elem_etmv3.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv3/trc_pkt_proc_etmv3.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv3/trc_pkt_proc_etmv3_impl.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv4/trc_cmp_cfg_etmv4.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv4/trc_etmv4_stack_elem.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv4/trc_pkt_decode_etmv4i.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv4/trc_pkt_elem_etmv4i.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/etmv4/trc_pkt_proc_etmv4i.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ete/trc_cmp_cfg_ete.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/i_dec/trc_i_decode.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/i_dec/trc_idec_arminst.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/mem_acc/trc_mem_acc_base.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/mem_acc/trc_mem_acc_bufptr.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/mem_acc/trc_mem_acc_cache.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/mem_acc/trc_mem_acc_cb.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/mem_acc/trc_mem_acc_file.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/mem_acc/trc_mem_acc_mapper.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/pkt_printers/raw_frame_printer.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/pkt_printers/trc_print_fact.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ptm/trc_cmp_cfg_ptm.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ptm/trc_pkt_decode_ptm.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ptm/trc_pkt_elem_ptm.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/ptm/trc_pkt_proc_ptm.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/stm/trc_pkt_decode_stm.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/stm/trc_pkt_elem_stm.cpp
        ${ANDROID_ROOT}/external/OpenCSD/decoder/source/stm/trc_pkt_proc_stm.cpp
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
    add_library(libopencsd_decoder SHARED ${SOURCE_FILES})

    target_include_directories(libopencsd_decoder
            PUBLIC
            ${ANDROID_ROOT}/external/OpenCSD/decoder/include
            ${ANDROID_ROOT}/external/OpenCSD
    )

    # Add c/c++ flags
    target_compile_options(libopencsd_decoder PRIVATE
            -Werror
    )

    target_link_libraries(libopencsd_decoder
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libopencsd_decoder ####################

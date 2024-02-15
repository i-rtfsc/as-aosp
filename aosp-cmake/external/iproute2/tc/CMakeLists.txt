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

project(tc)

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

#################### tc ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/iproute2/tc/tc/android_${OUT_ARCH_CPU}/gen/lex/external/iproute2/tc/emp_ematch.c
        ${OUT_INTERMEDIATES_ROOT}/external/iproute2/tc/emp_ematch.yacc.c/gen/emp_ematch.yacc.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_action.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_bpf.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_connmark.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_csum.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_ematch.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_estimator.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_gact.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_ife.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_mirred.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_pedit.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_police.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_sample.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_simple.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_skbedit.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_skbmod.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_tunnel_key.c
        ${ANDROID_ROOT}/external/iproute2/tc/m_vlan.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_basic.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_bpf.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_cgroup.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_flow.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_flower.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_fw.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_matchall.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_route.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_rsvp.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_tcindex.c
        ${ANDROID_ROOT}/external/iproute2/tc/f_u32.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_cbq.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_clsact.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_codel.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_drr.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_dsmark.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_fifo.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_fq.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_fq_codel.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_hfsc.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_hhf.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_htb.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_ingress.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_mqprio.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_multiq.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_netem.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_pie.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_prio.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_qfq.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_rr.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_sfb.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_tbf.c
        ${ANDROID_ROOT}/external/iproute2/tc/static-syms.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_cbq.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_class.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_core.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_exec.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_estimator.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_filter.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_monitor.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_qdisc.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_red.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_stab.c
        ${ANDROID_ROOT}/external/iproute2/tc/tc_util.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_choke.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_gred.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_red.c
        ${ANDROID_ROOT}/external/iproute2/tc/q_sfq.c
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
    add_executable(tc ${SOURCE_FILES})

    target_include_directories(tc
            PUBLIC
            ${ANDROID_ROOT}/external/iproute2/tc
            ${OUT_INTERMEDIATES_ROOT}/external/iproute2/tc/emp_ematch.yacc.h/gen
            ${OUT_INTERMEDIATES_ROOT}/external/iproute2/tc/static-syms.h/gen
    )

    # Add c/c++ flags
    target_compile_options(tc PRIVATE
            -Wall
            -Wno-unused-parameter
            -Werror
    )

    target_link_libraries(tc
            iproute2_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libiprouteutil
            libm
            libnetlink
    )
endif ()
#################### tc ####################

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

project(libltp_mem)

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

#################### libltp_mem ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/testcases/kernel/mem/lib/mem.c
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
    add_library(libltp_mem SHARED ${SOURCE_FILES})

    target_include_directories(libltp_mem
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp/testcases/kernel/include
            ${ANDROID_ROOT}/external/ltp/testcases/kernel/mem/include
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_mem PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_mem
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_mem ####################

#################### libltp_controllers ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/testcases/kernel/controllers/libcontrollers/libcontrollers.c
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
    add_library(libltp_controllers SHARED ${SOURCE_FILES})

    target_include_directories(libltp_controllers
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_controllers PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_controllers
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_controllers ####################

#################### libltp_hugetlb ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/testcases/kernel/mem/hugetlb/lib/hugetlb.c
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
    add_library(libltp_hugetlb SHARED ${SOURCE_FILES})

    target_include_directories(libltp_hugetlb
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp/testcases/kernel/include
            ${ANDROID_ROOT}/external/ltp/testcases/kernel/mem/include
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_hugetlb PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_hugetlb
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_hugetlb ####################

#################### libltp_ltpuinput ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/libs/libltpuinput/tst_uinput.c
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
    add_library(libltp_ltpuinput SHARED ${SOURCE_FILES})

    target_include_directories(libltp_ltpuinput
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_ltpuinput PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_ltpuinput
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_ltpuinput ####################

#################### libltp_ltp ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/lib/cloner.c
        ${ANDROID_ROOT}/external/ltp/lib/get_path.c
        ${ANDROID_ROOT}/external/ltp/lib/parse_opts.c
        ${ANDROID_ROOT}/external/ltp/lib/random_range.c
        ${ANDROID_ROOT}/external/ltp/lib/safe_file_ops.c
        ${ANDROID_ROOT}/external/ltp/lib/safe_macros.c
        ${ANDROID_ROOT}/external/ltp/lib/safe_net.c
        ${ANDROID_ROOT}/external/ltp/lib/safe_pthread.c
        ${ANDROID_ROOT}/external/ltp/lib/safe_stdio.c
        ${ANDROID_ROOT}/external/ltp/lib/self_exec.c
        ${ANDROID_ROOT}/external/ltp/lib/tlibio.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_af_alg.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_ansi_color.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_arch.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_assert.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_bool_expr.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_buffers.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_capability.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_cgroup.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_checkpoint.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_checksum.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_clocks.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_clone.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_cmd.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_coredump.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_cpu.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_crypto.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_device.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_dir_is_empty.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_epoll.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_fill_file.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_fill_fs.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_fips.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_fs_has_free.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_fs_link_count.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_fs_setup.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_fs_type.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_get_bad_addr.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_hugepage.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_ioctl.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_kconfig.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_kernel.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_kvercmp.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_lockdown.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_memutils.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_mkfs.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_module.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_net.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_netdevice.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_parse_opts.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_path_has_mnt_flags.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_pid.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_process_state.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_rand_data.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_res.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_resource.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_rtctime.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_rtnetlink.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_safe_file_at.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_safe_io_uring.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_safe_macros.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_safe_sysv_ipc.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_safe_timerfd.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_sig.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_sig_proc.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_status.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_supported_fs_types.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_sys_conf.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_taint.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_test.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_thread_state.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_timer.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_timer_test.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_tmpdir.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_uid.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_virt.c
        ${ANDROID_ROOT}/external/ltp/lib/tst_wallclock.c
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
    add_library(libltp_ltp SHARED ${SOURCE_FILES})

    target_include_directories(libltp_ltp
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp/lib
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_ltp PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_ltp
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_ltp ####################

#################### libltp_cpu_set ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/testcases/kernel/controllers/cpuset/cpuset_lib/cpuinfo.c
        ${ANDROID_ROOT}/external/ltp/testcases/kernel/controllers/cpuset/cpuset_lib/libbitmask.c
        ${ANDROID_ROOT}/external/ltp/testcases/kernel/controllers/cpuset/cpuset_lib/libcpuset.c
        ${ANDROID_ROOT}/external/ltp/testcases/kernel/controllers/cpuset/cpuset_lib/meminfo.c
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
    add_library(libltp_cpu_set SHARED ${SOURCE_FILES})

    target_include_directories(libltp_cpu_set
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_cpu_set PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_cpu_set
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_cpu_set ####################

#################### libltp_ltpnuma ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/libs/libltpnuma/tst_numa.c
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
    add_library(libltp_ltpnuma SHARED ${SOURCE_FILES})

    target_include_directories(libltp_ltpnuma
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_ltpnuma PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_ltpnuma
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_ltpnuma ####################

#################### libltp_ltpswap ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/libs/libltpswap/libswap.c
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
    add_library(libltp_ltpswap SHARED ${SOURCE_FILES})

    target_include_directories(libltp_ltpswap
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_ltpswap PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_ltpswap
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_ltpswap ####################

#################### libltp_ltpsigwait ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/libs/libltpsigwait/sigwait.c
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
    add_library(libltp_ltpsigwait SHARED ${SOURCE_FILES})

    target_include_directories(libltp_ltpsigwait
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_ltpsigwait PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_ltpsigwait
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_ltpsigwait ####################

#################### libltp_ltpnewipc ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/libs/libltpnewipc/libnewipc.c
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
    add_library(libltp_ltpnewipc SHARED ${SOURCE_FILES})

    target_include_directories(libltp_ltpnewipc
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_ltpnewipc PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_ltpnewipc
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_ltpnewipc ####################

#################### libltp_ltpvdso ####################
set(FILES
        ${ANDROID_ROOT}/external/ltp/libs/libltpvdso/parse_vdso.c
        ${ANDROID_ROOT}/external/ltp/libs/libltpvdso/vdso_helpers.c
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
    add_library(libltp_ltpvdso SHARED ${SOURCE_FILES})

    target_include_directories(libltp_ltpvdso
            PUBLIC
            ${ANDROID_ROOT}/external/ltp/android/include
            ${ANDROID_ROOT}/external/ltp/include
            ${ANDROID_ROOT}/external/ltp/include/old
            ${ANDROID_ROOT}/external/ltp
            ${OUT_INTERMEDIATES_ROOT}/external/ltp/ltp_syscalls_h/gen
    )

    # Add c/c++ flags
    target_compile_options(libltp_ltpvdso PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libltp_ltpvdso
            libasyncio
            libc
            libc++_static
            libc++demangle
            libcap
            libclang_rt.builtins
            libdl
            libm
            libselinux
    )
endif ()
#################### libltp_ltpvdso ####################

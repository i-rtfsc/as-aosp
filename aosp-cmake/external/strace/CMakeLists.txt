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

project(strace)

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

#################### strace ####################
set(FILES
        ${ANDROID_ROOT}/external/strace/access.c
        ${ANDROID_ROOT}/external/strace/affinity.c
        ${ANDROID_ROOT}/external/strace/aio.c
        ${ANDROID_ROOT}/external/strace/basic_filters.c
        ${ANDROID_ROOT}/external/strace/bind.c
        ${ANDROID_ROOT}/external/strace/bjm.c
        ${ANDROID_ROOT}/external/strace/block.c
        ${ANDROID_ROOT}/external/strace/bpf.c
        ${ANDROID_ROOT}/external/strace/bpf_filter.c
        ${ANDROID_ROOT}/external/strace/bpf_seccomp_filter.c
        ${ANDROID_ROOT}/external/strace/bpf_sock_filter.c
        ${ANDROID_ROOT}/external/strace/btrfs.c
        ${ANDROID_ROOT}/external/strace/cacheflush.c
        ${ANDROID_ROOT}/external/strace/capability.c
        ${ANDROID_ROOT}/external/strace/chdir.c
        ${ANDROID_ROOT}/external/strace/chmod.c
        ${ANDROID_ROOT}/external/strace/clone.c
        ${ANDROID_ROOT}/external/strace/copy_file_range.c
        ${ANDROID_ROOT}/external/strace/count.c
        ${ANDROID_ROOT}/external/strace/delay.c
        ${ANDROID_ROOT}/external/strace/desc.c
        ${ANDROID_ROOT}/external/strace/dirent.c
        ${ANDROID_ROOT}/external/strace/dirent64.c
        ${ANDROID_ROOT}/external/strace/dm.c
        ${ANDROID_ROOT}/external/strace/dyxlat.c
        ${ANDROID_ROOT}/external/strace/epoll.c
        ${ANDROID_ROOT}/external/strace/error_prints.c
        ${ANDROID_ROOT}/external/strace/evdev.c
        ${ANDROID_ROOT}/external/strace/evdev_mpers.c
        ${ANDROID_ROOT}/external/strace/eventfd.c
        ${ANDROID_ROOT}/external/strace/execve.c
        ${ANDROID_ROOT}/external/strace/fadvise.c
        ${ANDROID_ROOT}/external/strace/fallocate.c
        ${ANDROID_ROOT}/external/strace/fanotify.c
        ${ANDROID_ROOT}/external/strace/fchownat.c
        ${ANDROID_ROOT}/external/strace/fcntl.c
        ${ANDROID_ROOT}/external/strace/fetch_bpf_fprog.c
        ${ANDROID_ROOT}/external/strace/fetch_indirect_syscall_args.c
        ${ANDROID_ROOT}/external/strace/fetch_struct_flock.c
        ${ANDROID_ROOT}/external/strace/fetch_struct_keyctl_kdf_params.c
        ${ANDROID_ROOT}/external/strace/fetch_struct_mmsghdr.c
        ${ANDROID_ROOT}/external/strace/fetch_struct_msghdr.c
        ${ANDROID_ROOT}/external/strace/fetch_struct_stat.c
        ${ANDROID_ROOT}/external/strace/fetch_struct_stat64.c
        ${ANDROID_ROOT}/external/strace/fetch_struct_statfs.c
        ${ANDROID_ROOT}/external/strace/fetch_struct_xfs_quotastat.c
        ${ANDROID_ROOT}/external/strace/file_handle.c
        ${ANDROID_ROOT}/external/strace/file_ioctl.c
        ${ANDROID_ROOT}/external/strace/filter_qualify.c
        ${ANDROID_ROOT}/external/strace/flock.c
        ${ANDROID_ROOT}/external/strace/fstatfs.c
        ${ANDROID_ROOT}/external/strace/fstatfs64.c
        ${ANDROID_ROOT}/external/strace/fs_x_ioctl.c
        ${ANDROID_ROOT}/external/strace/futex.c
        ${ANDROID_ROOT}/external/strace/getcpu.c
        ${ANDROID_ROOT}/external/strace/getcwd.c
        ${ANDROID_ROOT}/external/strace/getpagesize.c
        ${ANDROID_ROOT}/external/strace/getrandom.c
        ${ANDROID_ROOT}/external/strace/get_robust_list.c
        ${ANDROID_ROOT}/external/strace/hdio.c
        ${ANDROID_ROOT}/external/strace/hostname.c
        ${ANDROID_ROOT}/external/strace/inotify.c
        ${ANDROID_ROOT}/external/strace/inotify_ioctl.c
        ${ANDROID_ROOT}/external/strace/io.c
        ${ANDROID_ROOT}/external/strace/ioctl.c
        ${ANDROID_ROOT}/external/strace/ioperm.c
        ${ANDROID_ROOT}/external/strace/iopl.c
        ${ANDROID_ROOT}/external/strace/ioprio.c
        ${ANDROID_ROOT}/external/strace/ipc.c
        ${ANDROID_ROOT}/external/strace/ipc_msg.c
        ${ANDROID_ROOT}/external/strace/ipc_msgctl.c
        ${ANDROID_ROOT}/external/strace/ipc_sem.c
        ${ANDROID_ROOT}/external/strace/ipc_shm.c
        ${ANDROID_ROOT}/external/strace/ipc_shmctl.c
        ${ANDROID_ROOT}/external/strace/kcmp.c
        ${ANDROID_ROOT}/external/strace/kexec.c
        ${ANDROID_ROOT}/external/strace/keyctl.c
        ${ANDROID_ROOT}/external/strace/kvm.c
        ${ANDROID_ROOT}/external/strace/ldt.c
        ${ANDROID_ROOT}/external/strace/link.c
        ${ANDROID_ROOT}/external/strace/listen.c
        ${ANDROID_ROOT}/external/strace/lookup_dcookie.c
        ${ANDROID_ROOT}/external/strace/loop.c
        ${ANDROID_ROOT}/external/strace/lseek.c
        ${ANDROID_ROOT}/external/strace/mem.c
        ${ANDROID_ROOT}/external/strace/membarrier.c
        ${ANDROID_ROOT}/external/strace/memfd_create.c
        ${ANDROID_ROOT}/external/strace/mknod.c
        ${ANDROID_ROOT}/external/strace/mmap_cache.c
        ${ANDROID_ROOT}/external/strace/mmap_notify.c
        ${ANDROID_ROOT}/external/strace/mmsghdr.c
        ${ANDROID_ROOT}/external/strace/mount.c
        ${ANDROID_ROOT}/external/strace/mq.c
        ${ANDROID_ROOT}/external/strace/msghdr.c
        ${ANDROID_ROOT}/external/strace/mtd.c
        ${ANDROID_ROOT}/external/strace/nbd_ioctl.c
        ${ANDROID_ROOT}/external/strace/net.c
        ${ANDROID_ROOT}/external/strace/netlink.c
        ${ANDROID_ROOT}/external/strace/netlink_crypto.c
        ${ANDROID_ROOT}/external/strace/netlink_inet_diag.c
        ${ANDROID_ROOT}/external/strace/netlink_kobject_uevent.c
        ${ANDROID_ROOT}/external/strace/netlink_netfilter.c
        ${ANDROID_ROOT}/external/strace/netlink_netlink_diag.c
        ${ANDROID_ROOT}/external/strace/netlink_packet_diag.c
        ${ANDROID_ROOT}/external/strace/netlink_route.c
        ${ANDROID_ROOT}/external/strace/netlink_selinux.c
        ${ANDROID_ROOT}/external/strace/netlink_smc_diag.c
        ${ANDROID_ROOT}/external/strace/netlink_sock_diag.c
        ${ANDROID_ROOT}/external/strace/netlink_unix_diag.c
        ${ANDROID_ROOT}/external/strace/nlattr.c
        ${ANDROID_ROOT}/external/strace/nsfs.c
        ${ANDROID_ROOT}/external/strace/numa.c
        ${ANDROID_ROOT}/external/strace/number_set.c
        ${ANDROID_ROOT}/external/strace/oldstat.c
        ${ANDROID_ROOT}/external/strace/open.c
        ${ANDROID_ROOT}/external/strace/pathtrace.c
        ${ANDROID_ROOT}/external/strace/perf.c
        ${ANDROID_ROOT}/external/strace/perf_ioctl.c
        ${ANDROID_ROOT}/external/strace/personality.c
        ${ANDROID_ROOT}/external/strace/pkeys.c
        ${ANDROID_ROOT}/external/strace/poll.c
        ${ANDROID_ROOT}/external/strace/prctl.c
        ${ANDROID_ROOT}/external/strace/print_aio_sigset.c
        ${ANDROID_ROOT}/external/strace/print_dev_t.c
        ${ANDROID_ROOT}/external/strace/print_group_req.c
        ${ANDROID_ROOT}/external/strace/print_ifindex.c
        ${ANDROID_ROOT}/external/strace/print_mac.c
        ${ANDROID_ROOT}/external/strace/print_mq_attr.c
        ${ANDROID_ROOT}/external/strace/print_msgbuf.c
        ${ANDROID_ROOT}/external/strace/print_sigevent.c
        ${ANDROID_ROOT}/external/strace/print_sg_req_info.c
        ${ANDROID_ROOT}/external/strace/print_statfs.c
        ${ANDROID_ROOT}/external/strace/print_struct_stat.c
        ${ANDROID_ROOT}/external/strace/print_time.c
        ${ANDROID_ROOT}/external/strace/print_timespec.c
        ${ANDROID_ROOT}/external/strace/print_timeval.c
        ${ANDROID_ROOT}/external/strace/print_timex.c
        ${ANDROID_ROOT}/external/strace/printmode.c
        ${ANDROID_ROOT}/external/strace/printrusage.c
        ${ANDROID_ROOT}/external/strace/printsiginfo.c
        ${ANDROID_ROOT}/external/strace/process.c
        ${ANDROID_ROOT}/external/strace/process_vm.c
        ${ANDROID_ROOT}/external/strace/ptp.c
        ${ANDROID_ROOT}/external/strace/quota.c
        ${ANDROID_ROOT}/external/strace/readahead.c
        ${ANDROID_ROOT}/external/strace/readlink.c
        ${ANDROID_ROOT}/external/strace/reboot.c
        ${ANDROID_ROOT}/external/strace/renameat.c
        ${ANDROID_ROOT}/external/strace/resource.c
        ${ANDROID_ROOT}/external/strace/retval.c
        ${ANDROID_ROOT}/external/strace/riscv.c
        ${ANDROID_ROOT}/external/strace/rtc.c
        ${ANDROID_ROOT}/external/strace/rtnl_addr.c
        ${ANDROID_ROOT}/external/strace/rtnl_addrlabel.c
        ${ANDROID_ROOT}/external/strace/rtnl_dcb.c
        ${ANDROID_ROOT}/external/strace/rtnl_link.c
        ${ANDROID_ROOT}/external/strace/rtnl_mdb.c
        ${ANDROID_ROOT}/external/strace/rtnl_neigh.c
        ${ANDROID_ROOT}/external/strace/rtnl_neightbl.c
        ${ANDROID_ROOT}/external/strace/rtnl_netconf.c
        ${ANDROID_ROOT}/external/strace/rtnl_nsid.c
        ${ANDROID_ROOT}/external/strace/rtnl_route.c
        ${ANDROID_ROOT}/external/strace/rtnl_rule.c
        ${ANDROID_ROOT}/external/strace/rtnl_tc.c
        ${ANDROID_ROOT}/external/strace/rtnl_tc_action.c
        ${ANDROID_ROOT}/external/strace/rt_sigframe.c
        ${ANDROID_ROOT}/external/strace/rt_sigreturn.c
        ${ANDROID_ROOT}/external/strace/sched.c
        ${ANDROID_ROOT}/external/strace/scsi.c
        ${ANDROID_ROOT}/external/strace/seccomp.c
        ${ANDROID_ROOT}/external/strace/sendfile.c
        ${ANDROID_ROOT}/external/strace/sg_io_v3.c
        ${ANDROID_ROOT}/external/strace/sg_io_v4.c
        ${ANDROID_ROOT}/external/strace/shutdown.c
        ${ANDROID_ROOT}/external/strace/sigaltstack.c
        ${ANDROID_ROOT}/external/strace/signal.c
        ${ANDROID_ROOT}/external/strace/signalfd.c
        ${ANDROID_ROOT}/external/strace/sigreturn.c
        ${ANDROID_ROOT}/external/strace/sock.c
        ${ANDROID_ROOT}/external/strace/sockaddr.c
        ${ANDROID_ROOT}/external/strace/socketcall.c
        ${ANDROID_ROOT}/external/strace/socketutils.c
        ${ANDROID_ROOT}/external/strace/sram_alloc.c
        ${ANDROID_ROOT}/external/strace/stat.c
        ${ANDROID_ROOT}/external/strace/stat64.c
        ${ANDROID_ROOT}/external/strace/statfs.c
        ${ANDROID_ROOT}/external/strace/statfs64.c
        ${ANDROID_ROOT}/external/strace/statx.c
        ${ANDROID_ROOT}/external/strace/strace.c
        ${ANDROID_ROOT}/external/strace/string_to_uint.c
        ${ANDROID_ROOT}/external/strace/swapon.c
        ${ANDROID_ROOT}/external/strace/sync_file_range.c
        ${ANDROID_ROOT}/external/strace/sync_file_range2.c
        ${ANDROID_ROOT}/external/strace/syscall.c
        ${ANDROID_ROOT}/external/strace/sysctl.c
        ${ANDROID_ROOT}/external/strace/sysinfo.c
        ${ANDROID_ROOT}/external/strace/syslog.c
        ${ANDROID_ROOT}/external/strace/sysmips.c
        ${ANDROID_ROOT}/external/strace/term.c
        ${ANDROID_ROOT}/external/strace/time.c
        ${ANDROID_ROOT}/external/strace/times.c
        ${ANDROID_ROOT}/external/strace/truncate.c
        ${ANDROID_ROOT}/external/strace/ubi.c
        ${ANDROID_ROOT}/external/strace/ucopy.c
        ${ANDROID_ROOT}/external/strace/uid16.c
        ${ANDROID_ROOT}/external/strace/uid.c
        ${ANDROID_ROOT}/external/strace/umask.c
        ${ANDROID_ROOT}/external/strace/umount.c
        ${ANDROID_ROOT}/external/strace/uname.c
        ${ANDROID_ROOT}/external/strace/upeek.c
        ${ANDROID_ROOT}/external/strace/upoke.c
        ${ANDROID_ROOT}/external/strace/userfaultfd.c
        ${ANDROID_ROOT}/external/strace/ustat.c
        ${ANDROID_ROOT}/external/strace/util.c
        ${ANDROID_ROOT}/external/strace/utime.c
        ${ANDROID_ROOT}/external/strace/utimes.c
        ${ANDROID_ROOT}/external/strace/v4l2.c
        ${ANDROID_ROOT}/external/strace/wait.c
        ${ANDROID_ROOT}/external/strace/xattr.c
        ${ANDROID_ROOT}/external/strace/xlat.c
        ${ANDROID_ROOT}/external/strace/xmalloc.c
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
    add_executable(strace ${SOURCE_FILES})

    target_include_directories(strace
            PUBLIC
            ${ANDROID_ROOT}/external/strace/linux/aarch64
            ${ANDROID_ROOT}/external/strace/linux
            ${ANDROID_ROOT}/external/strace
    )

    # Add c/c++ flags
    target_compile_options(strace PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(strace
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### strace ####################

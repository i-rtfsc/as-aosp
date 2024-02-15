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

project(toybox-static)

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

#################### toybox_vendor ####################
set(FILES
        ${ANDROID_ROOT}/external/toybox/lib/args.c
        ${ANDROID_ROOT}/external/toybox/lib/commas.c
        ${ANDROID_ROOT}/external/toybox/lib/dirtree.c
        ${ANDROID_ROOT}/external/toybox/lib/env.c
        ${ANDROID_ROOT}/external/toybox/lib/lib.c
        ${ANDROID_ROOT}/external/toybox/lib/llist.c
        ${ANDROID_ROOT}/external/toybox/lib/net.c
        ${ANDROID_ROOT}/external/toybox/lib/portability.c
        ${ANDROID_ROOT}/external/toybox/lib/tty.c
        ${ANDROID_ROOT}/external/toybox/lib/utf8.c
        ${ANDROID_ROOT}/external/toybox/lib/xwrap.c
        ${ANDROID_ROOT}/external/toybox/main.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/gzip.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/hostname.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/md5sum.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/mktemp.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/seq.c
        ${ANDROID_ROOT}/external/toybox/toys/net/microcom.c
        ${ANDROID_ROOT}/external/toybox/toys/other/dos2unix.c
        ${ANDROID_ROOT}/external/toybox/toys/other/readlink.c
        ${ANDROID_ROOT}/external/toybox/toys/other/setsid.c
        ${ANDROID_ROOT}/external/toybox/toys/other/stat.c
        ${ANDROID_ROOT}/external/toybox/toys/other/timeout.c
        ${ANDROID_ROOT}/external/toybox/toys/other/truncate.c
        ${ANDROID_ROOT}/external/toybox/toys/other/which.c
        ${ANDROID_ROOT}/external/toybox/toys/other/xxd.c
        ${ANDROID_ROOT}/external/toybox/toys/other/yes.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/dd.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/diff.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/expr.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/getopt.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/tr.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/basename.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cat.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/chmod.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cmp.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/comm.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cp.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cpio.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cut.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/date.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/dirname.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/du.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/echo.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/env.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/file.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/find.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/getconf.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/grep.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/head.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/id.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ln.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ls.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/mkdir.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/nl.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/od.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/paste.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/patch.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/printf.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/pwd.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/rm.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/rmdir.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/sed.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/sleep.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/sort.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tail.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tar.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tee.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/test.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/touch.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/true.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uname.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uniq.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/wc.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/xargs.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ps.c
        ${ANDROID_ROOT}/external/toybox/toys/other/taskset.c
        ${ANDROID_ROOT}/external/toybox/toys/android/getenforce.c
        ${ANDROID_ROOT}/external/toybox/toys/android/load_policy.c
        ${ANDROID_ROOT}/external/toybox/toys/android/log.c
        ${ANDROID_ROOT}/external/toybox/toys/android/restorecon.c
        ${ANDROID_ROOT}/external/toybox/toys/android/runcon.c
        ${ANDROID_ROOT}/external/toybox/toys/android/sendevent.c
        ${ANDROID_ROOT}/external/toybox/toys/android/setenforce.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/dmesg.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/killall.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/mknod.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/mount.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/pidof.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/sync.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/umount.c
        ${ANDROID_ROOT}/external/toybox/toys/net/ifconfig.c
        ${ANDROID_ROOT}/external/toybox/toys/net/netcat.c
        ${ANDROID_ROOT}/external/toybox/toys/net/netstat.c
        ${ANDROID_ROOT}/external/toybox/toys/net/ping.c
        ${ANDROID_ROOT}/external/toybox/toys/net/rfkill.c
        ${ANDROID_ROOT}/external/toybox/toys/net/tunctl.c
        ${ANDROID_ROOT}/external/toybox/toys/other/acpi.c
        ${ANDROID_ROOT}/external/toybox/toys/other/base64.c
        ${ANDROID_ROOT}/external/toybox/toys/other/blkdiscard.c
        ${ANDROID_ROOT}/external/toybox/toys/other/blkid.c
        ${ANDROID_ROOT}/external/toybox/toys/other/blockdev.c
        ${ANDROID_ROOT}/external/toybox/toys/other/chcon.c
        ${ANDROID_ROOT}/external/toybox/toys/other/chroot.c
        ${ANDROID_ROOT}/external/toybox/toys/other/chrt.c
        ${ANDROID_ROOT}/external/toybox/toys/other/clear.c
        ${ANDROID_ROOT}/external/toybox/toys/other/devmem.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fallocate.c
        ${ANDROID_ROOT}/external/toybox/toys/other/flock.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fmt.c
        ${ANDROID_ROOT}/external/toybox/toys/other/free.c
        ${ANDROID_ROOT}/external/toybox/toys/other/freeramdisk.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fsfreeze.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fsync.c
        ${ANDROID_ROOT}/external/toybox/toys/other/help.c
        ${ANDROID_ROOT}/external/toybox/toys/other/hwclock.c
        ${ANDROID_ROOT}/external/toybox/toys/other/i2ctools.c
        ${ANDROID_ROOT}/external/toybox/toys/other/inotifyd.c
        ${ANDROID_ROOT}/external/toybox/toys/other/insmod.c
        ${ANDROID_ROOT}/external/toybox/toys/other/ionice.c
        ${ANDROID_ROOT}/external/toybox/toys/other/losetup.c
        ${ANDROID_ROOT}/external/toybox/toys/other/lsattr.c
        ${ANDROID_ROOT}/external/toybox/toys/other/lsmod.c
        ${ANDROID_ROOT}/external/toybox/toys/other/lsusb.c
        ${ANDROID_ROOT}/external/toybox/toys/other/makedevs.c
        ${ANDROID_ROOT}/external/toybox/toys/other/mkswap.c
        ${ANDROID_ROOT}/external/toybox/toys/other/modinfo.c
        ${ANDROID_ROOT}/external/toybox/toys/other/mountpoint.c
        ${ANDROID_ROOT}/external/toybox/toys/other/nbd_client.c
        ${ANDROID_ROOT}/external/toybox/toys/other/nsenter.c
        ${ANDROID_ROOT}/external/toybox/toys/other/partprobe.c
        ${ANDROID_ROOT}/external/toybox/toys/other/pivot_root.c
        ${ANDROID_ROOT}/external/toybox/toys/other/pmap.c
        ${ANDROID_ROOT}/external/toybox/toys/other/printenv.c
        ${ANDROID_ROOT}/external/toybox/toys/other/pwdx.c
        ${ANDROID_ROOT}/external/toybox/toys/other/readelf.c
        ${ANDROID_ROOT}/external/toybox/toys/other/rev.c
        ${ANDROID_ROOT}/external/toybox/toys/other/rmmod.c
        ${ANDROID_ROOT}/external/toybox/toys/other/rtcwake.c
        ${ANDROID_ROOT}/external/toybox/toys/other/setfattr.c
        ${ANDROID_ROOT}/external/toybox/toys/other/swapoff.c
        ${ANDROID_ROOT}/external/toybox/toys/other/swapon.c
        ${ANDROID_ROOT}/external/toybox/toys/other/sysctl.c
        ${ANDROID_ROOT}/external/toybox/toys/other/tac.c
        ${ANDROID_ROOT}/external/toybox/toys/other/uclampset.c
        ${ANDROID_ROOT}/external/toybox/toys/other/uptime.c
        ${ANDROID_ROOT}/external/toybox/toys/other/usleep.c
        ${ANDROID_ROOT}/external/toybox/toys/other/uuidgen.c
        ${ANDROID_ROOT}/external/toybox/toys/other/vconfig.c
        ${ANDROID_ROOT}/external/toybox/toys/other/vmstat.c
        ${ANDROID_ROOT}/external/toybox/toys/other/watch.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/brctl.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/getfattr.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/lsof.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/modprobe.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/more.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/stty.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/traceroute.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/vi.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cal.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/chgrp.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cksum.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/df.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/expand.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/false.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/iconv.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/kill.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/logger.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/mkfifo.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/nice.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/nohup.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/renice.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/split.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/strings.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/time.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tty.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ulimit.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/unlink.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uudecode.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uuencode.c
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
    add_executable(toybox_vendor ${SOURCE_FILES})

    target_include_directories(toybox_vendor
            PUBLIC
            ${ANDROID_ROOT}/external/toybox/android/device
            ${ANDROID_ROOT}/external/toybox
    )

    # Add c/c++ flags
    target_compile_options(toybox_vendor PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(toybox_vendor
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libprocessgroup
            libselinux
            libz
            libc
            libdl
            libm
    )
endif ()
#################### toybox_vendor ####################

#################### toybox ####################
set(FILES
        ${ANDROID_ROOT}/external/toybox/lib/args.c
        ${ANDROID_ROOT}/external/toybox/lib/commas.c
        ${ANDROID_ROOT}/external/toybox/lib/dirtree.c
        ${ANDROID_ROOT}/external/toybox/lib/env.c
        ${ANDROID_ROOT}/external/toybox/lib/lib.c
        ${ANDROID_ROOT}/external/toybox/lib/llist.c
        ${ANDROID_ROOT}/external/toybox/lib/net.c
        ${ANDROID_ROOT}/external/toybox/lib/portability.c
        ${ANDROID_ROOT}/external/toybox/lib/tty.c
        ${ANDROID_ROOT}/external/toybox/lib/utf8.c
        ${ANDROID_ROOT}/external/toybox/lib/xwrap.c
        ${ANDROID_ROOT}/external/toybox/main.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/gzip.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/hostname.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/md5sum.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/mktemp.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/seq.c
        ${ANDROID_ROOT}/external/toybox/toys/net/microcom.c
        ${ANDROID_ROOT}/external/toybox/toys/other/dos2unix.c
        ${ANDROID_ROOT}/external/toybox/toys/other/readlink.c
        ${ANDROID_ROOT}/external/toybox/toys/other/setsid.c
        ${ANDROID_ROOT}/external/toybox/toys/other/stat.c
        ${ANDROID_ROOT}/external/toybox/toys/other/timeout.c
        ${ANDROID_ROOT}/external/toybox/toys/other/truncate.c
        ${ANDROID_ROOT}/external/toybox/toys/other/which.c
        ${ANDROID_ROOT}/external/toybox/toys/other/xxd.c
        ${ANDROID_ROOT}/external/toybox/toys/other/yes.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/dd.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/diff.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/expr.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/getopt.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/tr.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/basename.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cat.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/chmod.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cmp.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/comm.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cp.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cpio.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cut.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/date.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/dirname.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/du.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/echo.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/env.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/file.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/find.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/getconf.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/grep.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/head.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/id.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ln.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ls.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/mkdir.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/nl.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/od.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/paste.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/patch.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/printf.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/pwd.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/rm.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/rmdir.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/sed.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/sleep.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/sort.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tail.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tar.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tee.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/test.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/touch.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/true.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uname.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uniq.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/wc.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/xargs.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ps.c
        ${ANDROID_ROOT}/external/toybox/toys/other/taskset.c
        ${ANDROID_ROOT}/external/toybox/toys/android/getenforce.c
        ${ANDROID_ROOT}/external/toybox/toys/android/load_policy.c
        ${ANDROID_ROOT}/external/toybox/toys/android/log.c
        ${ANDROID_ROOT}/external/toybox/toys/android/restorecon.c
        ${ANDROID_ROOT}/external/toybox/toys/android/runcon.c
        ${ANDROID_ROOT}/external/toybox/toys/android/sendevent.c
        ${ANDROID_ROOT}/external/toybox/toys/android/setenforce.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/dmesg.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/killall.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/mknod.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/mount.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/pidof.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/sync.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/umount.c
        ${ANDROID_ROOT}/external/toybox/toys/net/ifconfig.c
        ${ANDROID_ROOT}/external/toybox/toys/net/netcat.c
        ${ANDROID_ROOT}/external/toybox/toys/net/netstat.c
        ${ANDROID_ROOT}/external/toybox/toys/net/ping.c
        ${ANDROID_ROOT}/external/toybox/toys/net/rfkill.c
        ${ANDROID_ROOT}/external/toybox/toys/net/tunctl.c
        ${ANDROID_ROOT}/external/toybox/toys/other/acpi.c
        ${ANDROID_ROOT}/external/toybox/toys/other/base64.c
        ${ANDROID_ROOT}/external/toybox/toys/other/blkdiscard.c
        ${ANDROID_ROOT}/external/toybox/toys/other/blkid.c
        ${ANDROID_ROOT}/external/toybox/toys/other/blockdev.c
        ${ANDROID_ROOT}/external/toybox/toys/other/chcon.c
        ${ANDROID_ROOT}/external/toybox/toys/other/chroot.c
        ${ANDROID_ROOT}/external/toybox/toys/other/chrt.c
        ${ANDROID_ROOT}/external/toybox/toys/other/clear.c
        ${ANDROID_ROOT}/external/toybox/toys/other/devmem.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fallocate.c
        ${ANDROID_ROOT}/external/toybox/toys/other/flock.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fmt.c
        ${ANDROID_ROOT}/external/toybox/toys/other/free.c
        ${ANDROID_ROOT}/external/toybox/toys/other/freeramdisk.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fsfreeze.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fsync.c
        ${ANDROID_ROOT}/external/toybox/toys/other/help.c
        ${ANDROID_ROOT}/external/toybox/toys/other/hwclock.c
        ${ANDROID_ROOT}/external/toybox/toys/other/i2ctools.c
        ${ANDROID_ROOT}/external/toybox/toys/other/inotifyd.c
        ${ANDROID_ROOT}/external/toybox/toys/other/insmod.c
        ${ANDROID_ROOT}/external/toybox/toys/other/ionice.c
        ${ANDROID_ROOT}/external/toybox/toys/other/losetup.c
        ${ANDROID_ROOT}/external/toybox/toys/other/lsattr.c
        ${ANDROID_ROOT}/external/toybox/toys/other/lsmod.c
        ${ANDROID_ROOT}/external/toybox/toys/other/lsusb.c
        ${ANDROID_ROOT}/external/toybox/toys/other/makedevs.c
        ${ANDROID_ROOT}/external/toybox/toys/other/mkswap.c
        ${ANDROID_ROOT}/external/toybox/toys/other/modinfo.c
        ${ANDROID_ROOT}/external/toybox/toys/other/mountpoint.c
        ${ANDROID_ROOT}/external/toybox/toys/other/nbd_client.c
        ${ANDROID_ROOT}/external/toybox/toys/other/nsenter.c
        ${ANDROID_ROOT}/external/toybox/toys/other/partprobe.c
        ${ANDROID_ROOT}/external/toybox/toys/other/pivot_root.c
        ${ANDROID_ROOT}/external/toybox/toys/other/pmap.c
        ${ANDROID_ROOT}/external/toybox/toys/other/printenv.c
        ${ANDROID_ROOT}/external/toybox/toys/other/pwdx.c
        ${ANDROID_ROOT}/external/toybox/toys/other/readelf.c
        ${ANDROID_ROOT}/external/toybox/toys/other/rev.c
        ${ANDROID_ROOT}/external/toybox/toys/other/rmmod.c
        ${ANDROID_ROOT}/external/toybox/toys/other/rtcwake.c
        ${ANDROID_ROOT}/external/toybox/toys/other/setfattr.c
        ${ANDROID_ROOT}/external/toybox/toys/other/swapoff.c
        ${ANDROID_ROOT}/external/toybox/toys/other/swapon.c
        ${ANDROID_ROOT}/external/toybox/toys/other/sysctl.c
        ${ANDROID_ROOT}/external/toybox/toys/other/tac.c
        ${ANDROID_ROOT}/external/toybox/toys/other/uclampset.c
        ${ANDROID_ROOT}/external/toybox/toys/other/uptime.c
        ${ANDROID_ROOT}/external/toybox/toys/other/usleep.c
        ${ANDROID_ROOT}/external/toybox/toys/other/uuidgen.c
        ${ANDROID_ROOT}/external/toybox/toys/other/vconfig.c
        ${ANDROID_ROOT}/external/toybox/toys/other/vmstat.c
        ${ANDROID_ROOT}/external/toybox/toys/other/watch.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/brctl.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/getfattr.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/lsof.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/modprobe.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/more.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/stty.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/traceroute.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/vi.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cal.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/chgrp.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cksum.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/df.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/expand.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/false.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/iconv.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/kill.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/logger.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/mkfifo.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/nice.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/nohup.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/renice.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/split.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/strings.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/time.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tty.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ulimit.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/unlink.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uudecode.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uuencode.c
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
    add_executable(toybox ${SOURCE_FILES})

    target_include_directories(toybox
            PUBLIC
            ${ANDROID_ROOT}/external/toybox/android/device
            ${ANDROID_ROOT}/external/toybox
    )

    # Add c/c++ flags
    target_compile_options(toybox PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(toybox
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            liblog
            libm
            libprocessgroup
            libselinux
            libz
    )
endif ()
#################### toybox ####################

#################### toybox-static ####################
set(FILES
        ${ANDROID_ROOT}/external/toybox/lib/args.c
        ${ANDROID_ROOT}/external/toybox/lib/commas.c
        ${ANDROID_ROOT}/external/toybox/lib/dirtree.c
        ${ANDROID_ROOT}/external/toybox/lib/env.c
        ${ANDROID_ROOT}/external/toybox/lib/lib.c
        ${ANDROID_ROOT}/external/toybox/lib/llist.c
        ${ANDROID_ROOT}/external/toybox/lib/net.c
        ${ANDROID_ROOT}/external/toybox/lib/portability.c
        ${ANDROID_ROOT}/external/toybox/lib/tty.c
        ${ANDROID_ROOT}/external/toybox/lib/utf8.c
        ${ANDROID_ROOT}/external/toybox/lib/xwrap.c
        ${ANDROID_ROOT}/external/toybox/main.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/gzip.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/hostname.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/md5sum.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/mktemp.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/seq.c
        ${ANDROID_ROOT}/external/toybox/toys/net/microcom.c
        ${ANDROID_ROOT}/external/toybox/toys/other/dos2unix.c
        ${ANDROID_ROOT}/external/toybox/toys/other/readlink.c
        ${ANDROID_ROOT}/external/toybox/toys/other/setsid.c
        ${ANDROID_ROOT}/external/toybox/toys/other/stat.c
        ${ANDROID_ROOT}/external/toybox/toys/other/timeout.c
        ${ANDROID_ROOT}/external/toybox/toys/other/truncate.c
        ${ANDROID_ROOT}/external/toybox/toys/other/which.c
        ${ANDROID_ROOT}/external/toybox/toys/other/xxd.c
        ${ANDROID_ROOT}/external/toybox/toys/other/yes.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/dd.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/diff.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/expr.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/getopt.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/tr.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/basename.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cat.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/chmod.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cmp.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/comm.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cp.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cpio.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cut.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/date.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/dirname.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/du.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/echo.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/env.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/file.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/find.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/getconf.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/grep.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/head.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/id.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ln.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ls.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/mkdir.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/nl.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/od.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/paste.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/patch.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/printf.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/pwd.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/rm.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/rmdir.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/sed.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/sleep.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/sort.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tail.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tar.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tee.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/test.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/touch.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/true.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uname.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uniq.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/wc.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/xargs.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ps.c
        ${ANDROID_ROOT}/external/toybox/toys/other/taskset.c
        ${ANDROID_ROOT}/external/toybox/toys/android/getenforce.c
        ${ANDROID_ROOT}/external/toybox/toys/android/load_policy.c
        ${ANDROID_ROOT}/external/toybox/toys/android/log.c
        ${ANDROID_ROOT}/external/toybox/toys/android/restorecon.c
        ${ANDROID_ROOT}/external/toybox/toys/android/runcon.c
        ${ANDROID_ROOT}/external/toybox/toys/android/sendevent.c
        ${ANDROID_ROOT}/external/toybox/toys/android/setenforce.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/dmesg.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/killall.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/mknod.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/mount.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/pidof.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/sync.c
        ${ANDROID_ROOT}/external/toybox/toys/lsb/umount.c
        ${ANDROID_ROOT}/external/toybox/toys/net/ifconfig.c
        ${ANDROID_ROOT}/external/toybox/toys/net/netcat.c
        ${ANDROID_ROOT}/external/toybox/toys/net/netstat.c
        ${ANDROID_ROOT}/external/toybox/toys/net/ping.c
        ${ANDROID_ROOT}/external/toybox/toys/net/rfkill.c
        ${ANDROID_ROOT}/external/toybox/toys/net/tunctl.c
        ${ANDROID_ROOT}/external/toybox/toys/other/acpi.c
        ${ANDROID_ROOT}/external/toybox/toys/other/base64.c
        ${ANDROID_ROOT}/external/toybox/toys/other/blkdiscard.c
        ${ANDROID_ROOT}/external/toybox/toys/other/blkid.c
        ${ANDROID_ROOT}/external/toybox/toys/other/blockdev.c
        ${ANDROID_ROOT}/external/toybox/toys/other/chcon.c
        ${ANDROID_ROOT}/external/toybox/toys/other/chroot.c
        ${ANDROID_ROOT}/external/toybox/toys/other/chrt.c
        ${ANDROID_ROOT}/external/toybox/toys/other/clear.c
        ${ANDROID_ROOT}/external/toybox/toys/other/devmem.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fallocate.c
        ${ANDROID_ROOT}/external/toybox/toys/other/flock.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fmt.c
        ${ANDROID_ROOT}/external/toybox/toys/other/free.c
        ${ANDROID_ROOT}/external/toybox/toys/other/freeramdisk.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fsfreeze.c
        ${ANDROID_ROOT}/external/toybox/toys/other/fsync.c
        ${ANDROID_ROOT}/external/toybox/toys/other/help.c
        ${ANDROID_ROOT}/external/toybox/toys/other/hwclock.c
        ${ANDROID_ROOT}/external/toybox/toys/other/i2ctools.c
        ${ANDROID_ROOT}/external/toybox/toys/other/inotifyd.c
        ${ANDROID_ROOT}/external/toybox/toys/other/insmod.c
        ${ANDROID_ROOT}/external/toybox/toys/other/ionice.c
        ${ANDROID_ROOT}/external/toybox/toys/other/losetup.c
        ${ANDROID_ROOT}/external/toybox/toys/other/lsattr.c
        ${ANDROID_ROOT}/external/toybox/toys/other/lsmod.c
        ${ANDROID_ROOT}/external/toybox/toys/other/lsusb.c
        ${ANDROID_ROOT}/external/toybox/toys/other/makedevs.c
        ${ANDROID_ROOT}/external/toybox/toys/other/mkswap.c
        ${ANDROID_ROOT}/external/toybox/toys/other/modinfo.c
        ${ANDROID_ROOT}/external/toybox/toys/other/mountpoint.c
        ${ANDROID_ROOT}/external/toybox/toys/other/nbd_client.c
        ${ANDROID_ROOT}/external/toybox/toys/other/nsenter.c
        ${ANDROID_ROOT}/external/toybox/toys/other/partprobe.c
        ${ANDROID_ROOT}/external/toybox/toys/other/pivot_root.c
        ${ANDROID_ROOT}/external/toybox/toys/other/pmap.c
        ${ANDROID_ROOT}/external/toybox/toys/other/printenv.c
        ${ANDROID_ROOT}/external/toybox/toys/other/pwdx.c
        ${ANDROID_ROOT}/external/toybox/toys/other/readelf.c
        ${ANDROID_ROOT}/external/toybox/toys/other/rev.c
        ${ANDROID_ROOT}/external/toybox/toys/other/rmmod.c
        ${ANDROID_ROOT}/external/toybox/toys/other/rtcwake.c
        ${ANDROID_ROOT}/external/toybox/toys/other/setfattr.c
        ${ANDROID_ROOT}/external/toybox/toys/other/swapoff.c
        ${ANDROID_ROOT}/external/toybox/toys/other/swapon.c
        ${ANDROID_ROOT}/external/toybox/toys/other/sysctl.c
        ${ANDROID_ROOT}/external/toybox/toys/other/tac.c
        ${ANDROID_ROOT}/external/toybox/toys/other/uclampset.c
        ${ANDROID_ROOT}/external/toybox/toys/other/uptime.c
        ${ANDROID_ROOT}/external/toybox/toys/other/usleep.c
        ${ANDROID_ROOT}/external/toybox/toys/other/uuidgen.c
        ${ANDROID_ROOT}/external/toybox/toys/other/vconfig.c
        ${ANDROID_ROOT}/external/toybox/toys/other/vmstat.c
        ${ANDROID_ROOT}/external/toybox/toys/other/watch.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/brctl.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/getfattr.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/lsof.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/modprobe.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/more.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/stty.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/traceroute.c
        ${ANDROID_ROOT}/external/toybox/toys/pending/vi.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cal.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/chgrp.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/cksum.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/df.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/expand.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/false.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/iconv.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/kill.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/logger.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/mkfifo.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/nice.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/nohup.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/renice.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/split.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/strings.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/time.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/tty.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/ulimit.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/unlink.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uudecode.c
        ${ANDROID_ROOT}/external/toybox/toys/posix/uuencode.c
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
    add_executable(toybox-static ${SOURCE_FILES})

    target_include_directories(toybox-static
            PUBLIC
            ${ANDROID_ROOT}/external/toybox/android/device
            ${ANDROID_ROOT}/external/toybox
    )

    # Add c/c++ flags
    target_compile_options(toybox-static PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(toybox-static
            libbase
            libc
            libc++_static
            libc++demangle
            libcgrouprc
            libcgrouprc_format
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto_static
            libdl
            liblog
            libm
            libprocessgroup
            libselinux
            libunwind
            libz
    )
endif ()
#################### toybox-static ####################

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

project(trace-cmd)

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

#################### libtracecmd ####################
set(FILES
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/test.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-blk-hack.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-compress.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-compress-zlib.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-filter-hash.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-ftrace.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-hash.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-hooks.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-input.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-msg.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-output.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-perf.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-plugin.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-recorder.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-timesync.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-timesync-kvm.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-timesync-ptp.c
        ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/trace-util.c
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
    add_library(libtracecmd SHARED ${SOURCE_FILES})

    target_include_directories(libtracecmd
            PUBLIC
            ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/include
            ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/include/private
            ${ANDROID_ROOT}/external/trace-cmd/include/trace-cmd
            ${ANDROID_ROOT}/external/trace-cmd/tracecmd/include
            ${ANDROID_ROOT}/external/trace-cmd/include
            ${ANDROID_ROOT}/external/trace-cmd
            ${OUT_INTERMEDIATES_ROOT}/external/trace-cmd/tc_version_header/gen
    )

    # Add c/c++ flags
    target_compile_options(libtracecmd PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-but-set-variable
    )

    target_link_libraries(libtracecmd
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libtraceevent
            libtracefs
            libz
    )
endif ()
#################### libtracecmd ####################

#################### trace-cmd ####################
set(FILES
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-agent.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-check-events.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-clear.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-cmd.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-convert.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-dump.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-hist.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-list.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-listen.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-mem.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-profile.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-read.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-record.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-restore.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-setup-guest.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-show.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-snapshot.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-split.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-stack.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-stat.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-stream.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-usage.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-vm.c
        ${ANDROID_ROOT}/external/trace-cmd/tracecmd/trace-vsock.c
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
    add_executable(trace-cmd ${SOURCE_FILES})

    target_include_directories(trace-cmd
            PUBLIC
            ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/include/private
            ${ANDROID_ROOT}/external/trace-cmd/include/trace-cmd
            ${ANDROID_ROOT}/external/trace-cmd/tracecmd/include
            ${ANDROID_ROOT}/external/trace-cmd/include
            ${ANDROID_ROOT}/external/trace-cmd
            ${ANDROID_ROOT}/external/trace-cmd/lib/trace-cmd/include
            ${OUT_INTERMEDIATES_ROOT}/external/trace-cmd/tc_version_header/gen
    )

    # Add c/c++ flags
    target_compile_options(trace-cmd PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(trace-cmd
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libtracecmd
            libtraceevent
            libtracefs
            libunwind
    )
endif ()
#################### trace-cmd ####################

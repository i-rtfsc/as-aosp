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

project(btt)

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

#################### blktrace ####################
set(FILES
        ${ANDROID_ROOT}/external/blktrace/blktrace.c
        ${ANDROID_ROOT}/external/blktrace/act_mask.c
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
    add_executable(blktrace ${SOURCE_FILES})

    target_include_directories(blktrace
            PUBLIC
            ${ANDROID_ROOT}/external/blktrace
            ${ANDROID_ROOT}/external/blktrace/btt
    )

    # Add c/c++ flags
    target_compile_options(blktrace PRIVATE
            -Werror
            -Wall
            -Wno-unused-function
    )

    target_link_libraries(blktrace
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### blktrace ####################

#################### blkrawverify ####################
set(FILES
        ${ANDROID_ROOT}/external/blktrace/blkrawverify.c
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
    add_executable(blkrawverify ${SOURCE_FILES})

    target_include_directories(blkrawverify
            PUBLIC
            ${ANDROID_ROOT}/external/blktrace
            ${ANDROID_ROOT}/external/blktrace/btt
    )

    # Add c/c++ flags
    target_compile_options(blkrawverify PRIVATE
            -Werror
            -Wall
            -Wno-unused-function
    )

    target_link_libraries(blkrawverify
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### blkrawverify ####################

#################### blkparse ####################
set(FILES
        ${ANDROID_ROOT}/external/blktrace/blkparse.c
        ${ANDROID_ROOT}/external/blktrace/blkparse_fmt.c
        ${ANDROID_ROOT}/external/blktrace/rbtree.c
        ${ANDROID_ROOT}/external/blktrace/act_mask.c
        ${ANDROID_ROOT}/external/blktrace/strverscmp.c
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
    add_executable(blkparse ${SOURCE_FILES})

    target_include_directories(blkparse
            PUBLIC
            ${ANDROID_ROOT}/external/blktrace
            ${ANDROID_ROOT}/external/blktrace/btt
    )

    # Add c/c++ flags
    target_compile_options(blkparse PRIVATE
            -Werror
            -Wall
            -Wno-unused-function
    )

    target_link_libraries(blkparse
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### blkparse ####################

#################### verify_blkparse ####################
set(FILES
        ${ANDROID_ROOT}/external/blktrace/verify_blkparse.c
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
    add_executable(verify_blkparse ${SOURCE_FILES})

    target_include_directories(verify_blkparse
            PUBLIC
            ${ANDROID_ROOT}/external/blktrace
            ${ANDROID_ROOT}/external/blktrace/btt
    )

    # Add c/c++ flags
    target_compile_options(verify_blkparse PRIVATE
            -Werror
            -Wall
            -Wno-unused-function
    )

    target_link_libraries(verify_blkparse
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### verify_blkparse ####################

#################### blkiomon ####################
set(FILES
        ${ANDROID_ROOT}/external/blktrace/blkiomon.c
        ${ANDROID_ROOT}/external/blktrace/rbtree.c
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
    add_executable(blkiomon ${SOURCE_FILES})

    target_include_directories(blkiomon
            PUBLIC
            ${ANDROID_ROOT}/external/blktrace
            ${ANDROID_ROOT}/external/blktrace/btt
    )

    # Add c/c++ flags
    target_compile_options(blkiomon PRIVATE
            -Werror
            -Wall
            -Wno-unused-function
    )

    target_link_libraries(blkiomon
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### blkiomon ####################

#################### btt ####################
set(FILES
        ${ANDROID_ROOT}/external/blktrace/btt/args.c
        ${ANDROID_ROOT}/external/blktrace/btt/bt_timeline.c
        ${ANDROID_ROOT}/external/blktrace/btt/devmap.c
        ${ANDROID_ROOT}/external/blktrace/btt/devs.c
        ${ANDROID_ROOT}/external/blktrace/btt/dip_rb.c
        ${ANDROID_ROOT}/external/blktrace/btt/iostat.c
        ${ANDROID_ROOT}/external/blktrace/btt/latency.c
        ${ANDROID_ROOT}/external/blktrace/btt/misc.c
        ${ANDROID_ROOT}/external/blktrace/btt/output.c
        ${ANDROID_ROOT}/external/blktrace/btt/proc.c
        ${ANDROID_ROOT}/external/blktrace/btt/seek.c
        ${ANDROID_ROOT}/external/blktrace/btt/trace.c
        ${ANDROID_ROOT}/external/blktrace/btt/trace_complete.c
        ${ANDROID_ROOT}/external/blktrace/btt/trace_im.c
        ${ANDROID_ROOT}/external/blktrace/btt/trace_issue.c
        ${ANDROID_ROOT}/external/blktrace/btt/trace_queue.c
        ${ANDROID_ROOT}/external/blktrace/btt/trace_remap.c
        ${ANDROID_ROOT}/external/blktrace/btt/trace_requeue.c
        ${ANDROID_ROOT}/external/blktrace/rbtree.c
        ${ANDROID_ROOT}/external/blktrace/btt/mmap.c
        ${ANDROID_ROOT}/external/blktrace/btt/trace_plug.c
        ${ANDROID_ROOT}/external/blktrace/btt/bno_dump.c
        ${ANDROID_ROOT}/external/blktrace/btt/unplug_hist.c
        ${ANDROID_ROOT}/external/blktrace/btt/q2d.c
        ${ANDROID_ROOT}/external/blktrace/btt/aqd.c
        ${ANDROID_ROOT}/external/blktrace/btt/plat.c
        ${ANDROID_ROOT}/external/blktrace/btt/p_live.c
        ${ANDROID_ROOT}/external/blktrace/btt/rstats.c
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
    add_executable(btt ${SOURCE_FILES})

    target_include_directories(btt
            PUBLIC
            ${ANDROID_ROOT}/external/blktrace
            ${ANDROID_ROOT}/external/blktrace/btt
    )

    # Add c/c++ flags
    target_compile_options(btt PRIVATE
            -Werror
            -Wall
            -Wno-unused-function
    )

    target_link_libraries(btt
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### btt ####################

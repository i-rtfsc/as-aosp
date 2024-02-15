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

project(cpufreq)

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

#################### fsdist.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/fsdist.bpf.c
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
    add_library(fsdist.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(fsdist.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(fsdist.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(fsdist.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### fsdist.bpf.o ####################

#################### runqlat.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/runqlat.bpf.c
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
    add_library(runqlat.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(runqlat.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(runqlat.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(runqlat.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### runqlat.bpf.o ####################

#################### fsdist ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/fsdist.c
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
    add_executable(fsdist ${SOURCE_FILES})

    target_include_directories(fsdist
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/fsdist.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(fsdist PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(fsdist
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### fsdist ####################

#################### gethostlatency ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/gethostlatency.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/uprobe_helpers.c
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
    add_executable(gethostlatency ${SOURCE_FILES})

    target_include_directories(gethostlatency
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/gethostlatency.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(gethostlatency PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(gethostlatency
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### gethostlatency ####################

#################### offcputime.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/offcputime.bpf.c
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
    add_library(offcputime.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(offcputime.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(offcputime.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(offcputime.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### offcputime.bpf.o ####################

#################### drsnoop ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/drsnoop.c
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
    add_executable(drsnoop ${SOURCE_FILES})

    target_include_directories(drsnoop
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/drsnoop.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(drsnoop PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(drsnoop
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### drsnoop ####################

#################### funclatency ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/funclatency.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/uprobe_helpers.c
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
    add_executable(funclatency ${SOURCE_FILES})

    target_include_directories(funclatency
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/funclatency.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(funclatency PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(funclatency
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### funclatency ####################

#################### bindsnoop.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/bindsnoop.bpf.c
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
    add_library(bindsnoop.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(bindsnoop.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(bindsnoop.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(bindsnoop.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### bindsnoop.bpf.o ####################

#################### runqlen.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/runqlen.bpf.c
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
    add_library(runqlen.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(runqlen.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(runqlen.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(runqlen.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### runqlen.bpf.o ####################

#################### solisten ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/solisten.c
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
    add_executable(solisten ${SOURCE_FILES})

    target_include_directories(solisten
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/solisten.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(solisten PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(solisten
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### solisten ####################

#################### filelife.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/filelife.bpf.c
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
    add_library(filelife.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(filelife.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(filelife.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(filelife.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### filelife.bpf.o ####################

#################### llcstat.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/llcstat.bpf.c
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
    add_library(llcstat.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(llcstat.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(llcstat.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(llcstat.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### llcstat.bpf.o ####################

#################### fsslower ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/fsslower.c
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
    add_executable(fsslower ${SOURCE_FILES})

    target_include_directories(fsslower
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/fsslower.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(fsslower PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(fsslower
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### fsslower ####################

#################### bindsnoop ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/bindsnoop.c
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
    add_executable(bindsnoop ${SOURCE_FILES})

    target_include_directories(bindsnoop
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/bindsnoop.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(bindsnoop PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(bindsnoop
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### bindsnoop ####################

#################### runqlat ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/runqlat.c
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
    add_executable(runqlat ${SOURCE_FILES})

    target_include_directories(runqlat
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/runqlat.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(runqlat PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(runqlat
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### runqlat ####################

#################### cpudist.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/cpudist.bpf.c
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
    add_library(cpudist.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(cpudist.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(cpudist.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(cpudist.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### cpudist.bpf.o ####################

#################### softirqs.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/softirqs.bpf.c
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
    add_library(softirqs.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(softirqs.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(softirqs.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(softirqs.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### softirqs.bpf.o ####################

#################### filelife ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/filelife.c
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
    add_executable(filelife ${SOURCE_FILES})

    target_include_directories(filelife
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/filelife.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(filelife PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(filelife
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### filelife ####################

#################### solisten.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/solisten.bpf.c
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
    add_library(solisten.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(solisten.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(solisten.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(solisten.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### solisten.bpf.o ####################

#################### cpudist ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/cpudist.c
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
    add_executable(cpudist ${SOURCE_FILES})

    target_include_directories(cpudist
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/cpudist.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(cpudist PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(cpudist
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### cpudist ####################

#################### tcprtt.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/tcprtt.bpf.c
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
    add_library(tcprtt.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(tcprtt.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(tcprtt.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(tcprtt.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### tcprtt.bpf.o ####################

#################### tcpconnect.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/tcpconnect.bpf.c
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
    add_library(tcpconnect.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(tcpconnect.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(tcpconnect.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(tcpconnect.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### tcpconnect.bpf.o ####################

#################### tcprtt ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/tcprtt.c
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
    add_executable(tcprtt ${SOURCE_FILES})

    target_include_directories(tcprtt
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/tcprtt.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(tcprtt PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(tcprtt
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### tcprtt ####################

#################### filetop.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/filetop.bpf.c
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
    add_library(filetop.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(filetop.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(filetop.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(filetop.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### filetop.bpf.o ####################

#################### llcstat ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/llcstat.c
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
    add_executable(llcstat ${SOURCE_FILES})

    target_include_directories(llcstat
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/llcstat.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(llcstat PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(llcstat
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### llcstat ####################

#################### runqlen ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/runqlen.c
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
    add_executable(runqlen ${SOURCE_FILES})

    target_include_directories(runqlen
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/runqlen.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(runqlen PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(runqlen
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### runqlen ####################

#################### filetop ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/filetop.c
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
    add_executable(filetop ${SOURCE_FILES})

    target_include_directories(filetop
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/filetop.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(filetop PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(filetop
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### filetop ####################

#################### drsnoop.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/drsnoop.bpf.c
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
    add_library(drsnoop.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(drsnoop.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(drsnoop.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(drsnoop.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### drsnoop.bpf.o ####################

#################### runqslower.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/runqslower.bpf.c
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
    add_library(runqslower.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(runqslower.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(runqslower.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(runqslower.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### runqslower.bpf.o ####################

#################### runqslower ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/runqslower.c
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
    add_executable(runqslower ${SOURCE_FILES})

    target_include_directories(runqslower
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/runqslower.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(runqslower PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(runqslower
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### runqslower ####################

#################### offcputime ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/offcputime.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/uprobe_helpers.c
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
    add_executable(offcputime ${SOURCE_FILES})

    target_include_directories(offcputime
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/offcputime.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(offcputime PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(offcputime
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### offcputime ####################

#################### softirqs ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/softirqs.c
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
    add_executable(softirqs ${SOURCE_FILES})

    target_include_directories(softirqs
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/softirqs.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(softirqs PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(softirqs
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### softirqs ####################

#################### fsslower.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/fsslower.bpf.c
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
    add_library(fsslower.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(fsslower.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(fsslower.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(fsslower.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### fsslower.bpf.o ####################

#################### ksnoop.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/ksnoop.bpf.c
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
    add_library(ksnoop.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(ksnoop.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(ksnoop.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(ksnoop.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### ksnoop.bpf.o ####################

#################### tcpconnect ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/map_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/tcpconnect.c
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
    add_executable(tcpconnect ${SOURCE_FILES})

    target_include_directories(tcpconnect
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/tcpconnect.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(tcpconnect PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(tcpconnect
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### tcpconnect ####################

#################### vfsstat.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/vfsstat.bpf.c
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
    add_library(vfsstat.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(vfsstat.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(vfsstat.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(vfsstat.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### vfsstat.bpf.o ####################

#################### cpufreq.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/cpufreq.bpf.c
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
    add_library(cpufreq.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(cpufreq.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(cpufreq.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(cpufreq.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### cpufreq.bpf.o ####################

#################### funclatency.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/funclatency.bpf.c
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
    add_library(funclatency.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(funclatency.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(funclatency.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(funclatency.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### funclatency.bpf.o ####################

#################### gethostlatency.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/gethostlatency.bpf.c
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
    add_library(gethostlatency.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(gethostlatency.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(gethostlatency.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(gethostlatency.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### gethostlatency.bpf.o ####################

#################### vfsstat ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/vfsstat.c
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
    add_executable(vfsstat ${SOURCE_FILES})

    target_include_directories(vfsstat
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/vfsstat.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(vfsstat PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(vfsstat
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### vfsstat ####################

#################### oomkill ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/oomkill.c
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
    add_executable(oomkill ${SOURCE_FILES})

    target_include_directories(oomkill
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/oomkill.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(oomkill PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(oomkill
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### oomkill ####################

#################### oomkill.bpf.o ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/oomkill.bpf.c
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
    add_library(oomkill.bpf.o SHARED ${SOURCE_FILES})

    target_include_directories(oomkill.bpf.o
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/arm64
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
    )

    # Add c/c++ flags
    target_compile_options(oomkill.bpf.o PRIVATE
            -Werror
            -Wno-unused-command-line-argument
            -Wno-unused-parameter
    )

    target_link_libraries(oomkill.bpf.o
            bpf_prog_headers
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### oomkill.bpf.o ####################

#################### ksnoop ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/ksnoop.c
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
    add_executable(ksnoop ${SOURCE_FILES})

    target_include_directories(ksnoop
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/ksnoop.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(ksnoop PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(ksnoop
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### ksnoop ####################

#################### cpufreq ####################
set(FILES
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/android/argp.cpp
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/trace_helpers.c
        ${ANDROID_ROOT}/external/bcc/libbpf-tools/cpufreq.c
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
    add_executable(cpufreq ${SOURCE_FILES})

    target_include_directories(cpufreq
            PUBLIC
            ${ANDROID_ROOT}/external/bcc/libbpf-tools/android
            ${ANDROID_ROOT}/external/bcc/libbpf-tools
            ${OUT_INTERMEDIATES_ROOT}/external/bcc/libbpf-tools/cpufreq.skel.h/gen
    )

    # Add c/c++ flags
    target_compile_options(cpufreq PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(cpufreq
            libbpf
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libelf
            libm
            libz
    )
endif ()
#################### cpufreq ####################

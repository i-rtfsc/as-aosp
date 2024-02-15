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

project(libsimpleperf_report_fuzzer)

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

#################### simpleperf ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/simpleperf/main.cpp
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
    add_executable(simpleperf ${SOURCE_FILES})

    target_include_directories(simpleperf
            PUBLIC
            ${ANDROID_ROOT}/system/extras/simpleperf
    )

    # Add c/c++ flags
    target_compile_options(simpleperf PRIVATE
            -Werror
    )

    target_link_libraries(simpleperf
            libLLVM_android
            libbase
            libbuildversion
            libc
            libc++
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdexfile_support
            libdl
            libevent
            liblog
            liblzma
            libm
            libopencsd_decoder
            libprocinfo
            libprotobuf-cpp-lite
            librustc_demangle_static
            libsimpleperf
            libsimpleperf_etm_decoder
            libsimpleperf_regex
            libunwindstack
            libziparchive
    )
endif ()
#################### simpleperf ####################

#################### simpleperf_ndk ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/simpleperf/main.cpp
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
    add_executable(simpleperf_ndk ${SOURCE_FILES})

    target_include_directories(simpleperf_ndk
            PUBLIC
            ${ANDROID_ROOT}/system/extras/simpleperf
    )

    # Add c/c++ flags
    target_compile_options(simpleperf_ndk PRIVATE
            -Werror
    )

    target_link_libraries(simpleperf_ndk
            libLLVMBitReader
            libLLVMCore
            libLLVMMC
            libLLVMMCParser
            libLLVMObject
            libLLVMSupport
            libbase
            libbuildversion
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdexfile_static
            libdl
            libevent
            liblog
            liblzma
            libm
            libopencsd_decoder
            libprocinfo
            libprotobuf-cpp-lite
            librustc_demangle_static
            libsimpleperf
            libsimpleperf_etm_decoder
            libsimpleperf_regex
            libunwind
            libunwindstack
            libutils
            libwinpthread
            libz
            libziparchive
    )
endif ()
#################### simpleperf_ndk ####################

#################### libsimpleperf_etm_decoder ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/simpleperf/ETMDecoder.cpp
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
    add_library(libsimpleperf_etm_decoder SHARED ${SOURCE_FILES})

    target_include_directories(libsimpleperf_etm_decoder
            PUBLIC
            ${ANDROID_ROOT}/system/extras/simpleperf
    )

    # Add c/c++ flags
    target_compile_options(libsimpleperf_etm_decoder PRIVATE
            -Werror
    )

    target_link_libraries(libsimpleperf_etm_decoder
            libLLVMBitReader
            libLLVMCore
            libLLVMMC
            libLLVMMCParser
            libLLVMObject
            libLLVMSupport
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            liblzma
            libm
            libopencsd_decoder
            libwinpthread
            libz
            libziparchive
    )
endif ()
#################### libsimpleperf_etm_decoder ####################

#################### libsimpleperf_report_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/simpleperf/libsimpleperf_report_fuzzer.cpp
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
    add_executable(libsimpleperf_report_fuzzer ${SOURCE_FILES})

    target_include_directories(libsimpleperf_report_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/extras/simpleperf
    )

    # Add c/c++ flags
    target_compile_options(libsimpleperf_report_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libsimpleperf_report_fuzzer
            libLLVMBitReader
            libLLVMCore
            libLLVMMC
            libLLVMMCParser
            libLLVMObject
            libLLVMSupport
            libbase
            libbuildversion
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdexfile_static
            libdl
            libevent
            liblog
            liblzma
            libm
            libopencsd_decoder
            libprocinfo
            libprotobuf-cpp-lite
            librustc_demangle_static
            libsimpleperf
            libsimpleperf_etm_decoder
            libsimpleperf_regex
            libunwindstack
            libutils
            libz
            libziparchive
    )
endif ()
#################### libsimpleperf_report_fuzzer ####################

#################### libsimpleperf_profcollect ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/simpleperf/profcollect.cpp
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
    add_library(libsimpleperf_profcollect SHARED ${SOURCE_FILES})

    target_include_directories(libsimpleperf_profcollect
            PUBLIC
            ${ANDROID_ROOT}/system/extras/simpleperf
    )

    # Add c/c++ flags
    target_compile_options(libsimpleperf_profcollect PRIVATE
            -Werror
    )

    target_link_libraries(libsimpleperf_profcollect
            libLLVM_android
            libbase
            libbuildversion
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdexfile_support
            libdl
            libevent
            liblog
            liblzma
            libm
            libopencsd_decoder
            libpower
            libprocinfo
            libprotobuf-cpp-lite
            librustc_demangle_static
            libsimpleperf
            libsimpleperf_etm_decoder
            libsimpleperf_regex
            libunwindstack
            libziparchive
    )
endif ()
#################### libsimpleperf_profcollect ####################

#################### libsimpleperf_record ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/simpleperf/record_lib_interface.cpp
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
    add_library(libsimpleperf_record SHARED ${SOURCE_FILES})

    target_include_directories(libsimpleperf_record
            PUBLIC
            ${ANDROID_ROOT}/system/extras/simpleperf/include
            ${ANDROID_ROOT}/system/extras/simpleperf
    )

    # Add c/c++ flags
    target_compile_options(libsimpleperf_record PRIVATE
            -Werror
    )

    target_link_libraries(libsimpleperf_record
            libLLVMBitReader
            libLLVMCore
            libLLVMMC
            libLLVMMCParser
            libLLVMObject
            libLLVMSupport
            libbase
            libbuildversion
            libc
            libc++
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdexfile_static
            libdl
            libevent
            liblog
            liblzma
            libm
            libopencsd_decoder
            libprocinfo
            libprotobuf-cpp-lite
            librustc_demangle_static
            libsimpleperf
            libsimpleperf_etm_decoder
            libsimpleperf_regex
            libunwindstack
            libutils
            libz
            libziparchive
    )
endif ()
#################### libsimpleperf_record ####################

#################### libsimpleperf_regex ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/simpleperf/RegEx.cpp
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
    add_library(libsimpleperf_regex SHARED ${SOURCE_FILES})

    target_include_directories(libsimpleperf_regex
            PUBLIC
            ${ANDROID_ROOT}/system/extras/simpleperf
    )

    # Add c/c++ flags
    target_compile_options(libsimpleperf_regex PRIVATE
            -Werror
    )

    target_link_libraries(libsimpleperf_regex
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libsimpleperf_regex ####################

#################### libsimpleperf ####################
set(FILES
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_dumprecord.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_help.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_inject.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_kmem.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_merge.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_report.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_report_sample.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/extras/simpleperf/libsimpleperf/android_${OUT_ARCH_CPU}_static/gen/proto/system/extras/simpleperf/cmd_report_sample.pb.cc
        ${ANDROID_ROOT}/system/extras/simpleperf/command.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/dso.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/extras/simpleperf/libsimpleperf/android_${OUT_ARCH_CPU}_static/gen/proto/system/extras/simpleperf/etm_branch_list.pb.cc
        ${ANDROID_ROOT}/system/extras/simpleperf/ETMBranchListFile.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/event_attr.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/event_type.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/kallsyms.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/perf_regs.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/read_apk.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/read_elf.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/read_symbol_map.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/record.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/RecordFilter.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/extras/simpleperf/libsimpleperf/android_${OUT_ARCH_CPU}_static/gen/proto/system/extras/simpleperf/record_file.pb.cc
        ${ANDROID_ROOT}/system/extras/simpleperf/record_file_reader.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/record_file_writer.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/report_utils.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/thread_tree.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/tracing.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/utils.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/CallChainJoiner.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_api.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_debug_unwind.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_list.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_monitor.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_record.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_stat.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_trace_sched.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/environment.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/ETMRecorder.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/event_fd.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/event_selection_set.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/IOEventLoop.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/JITDebugReader.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/MapRecordReader.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/OfflineUnwinder.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/ProbeEvents.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/read_dex_file.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/RecordReadThread.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/workload.cpp
        ${ANDROID_ROOT}/system/extras/simpleperf/cmd_boot_record.cpp
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
    add_library(libsimpleperf SHARED ${SOURCE_FILES})

    target_include_directories(libsimpleperf
            PUBLIC
            ${ANDROID_ROOT}/system/extras/simpleperf
            ${OUT_INTERMEDIATES_ROOT}/system/extras/simpleperf/libsimpleperf/android_${OUT_ARCH_CPU}_static/gen/proto/system/extras/simpleperf
            ${OUT_INTERMEDIATES_ROOT}/system/extras/simpleperf/libsimpleperf/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libsimpleperf PRIVATE
            -Werror
    )

    target_link_libraries(libsimpleperf
            libLLVMBitReader
            libLLVMCore
            libLLVMMC
            libLLVMMCParser
            libLLVMObject
            libLLVMSupport
            libbase
            libbuildversion
            libc
            libc++_static
            libc++demangle
            libc++fs
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdexfile_support
            libdl
            libevent
            liblog
            liblzma
            libm
            libopencsd_decoder
            libprocinfo
            libprotobuf-cpp-lite
            librustc_demangle_static
            libsimpleperf_etm_decoder
            libsimpleperf_regex
            libunwindstack
            libutils
            libwinpthread
            libz
            libziparchive
    )
endif ()
#################### libsimpleperf ####################

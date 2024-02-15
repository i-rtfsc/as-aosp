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

project(libgoogle-benchmark)

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

#################### libgoogle-benchmark ####################
set(FILES
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark_api_internal.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark_name.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark_register.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark_runner.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/colorprint.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/commandlineflags.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/complexity.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/console_reporter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/counter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/csv_reporter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/json_reporter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/reporter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/sleep.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/statistics.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/string_util.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/sysinfo.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/timers.cc
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
    add_library(libgoogle-benchmark SHARED ${SOURCE_FILES})

    target_include_directories(libgoogle-benchmark
            PUBLIC
            ${ANDROID_ROOT}/external/google-benchmark/include
            ${ANDROID_ROOT}/external/google-benchmark
    )

    # Add c/c++ flags
    target_compile_options(libgoogle-benchmark PRIVATE
            -Werror
    )

    target_link_libraries(libgoogle-benchmark
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libgoogle-benchmark ####################

#################### libgoogle-benchmark-main ####################
set(FILES
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark_api_internal.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark_main.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark_name.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark_register.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/benchmark_runner.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/colorprint.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/commandlineflags.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/complexity.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/console_reporter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/counter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/csv_reporter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/json_reporter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/reporter.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/sleep.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/statistics.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/string_util.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/sysinfo.cc
        ${ANDROID_ROOT}/external/google-benchmark/src/timers.cc
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
    add_library(libgoogle-benchmark-main SHARED ${SOURCE_FILES})

    target_include_directories(libgoogle-benchmark-main
            PUBLIC
            ${ANDROID_ROOT}/external/google-benchmark/include
            ${ANDROID_ROOT}/external/google-benchmark
    )

    # Add c/c++ flags
    target_compile_options(libgoogle-benchmark-main PRIVATE
            -Werror
    )

    target_link_libraries(libgoogle-benchmark-main
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libgoogle-benchmark-main ####################

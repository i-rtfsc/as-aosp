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

project(bench_noop_static)

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

#################### bench_cxa_atexit ####################
set(FILES
        ${ANDROID_ROOT}/bionic/benchmarks/spawn/bench_cxa_atexit.cpp
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
    add_executable(bench_cxa_atexit ${SOURCE_FILES})

    target_include_directories(bench_cxa_atexit
            PUBLIC
            ${ANDROID_ROOT}/bionic/benchmarks/spawn
    )

    # Add c/c++ flags
    target_compile_options(bench_cxa_atexit PRIVATE
            -Werror
    )

    target_link_libraries(bench_cxa_atexit
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bench_cxa_atexit ####################

#################### bench_noop ####################
set(FILES
        ${ANDROID_ROOT}/bionic/benchmarks/spawn/noop.cpp
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
    add_executable(bench_noop ${SOURCE_FILES})

    target_include_directories(bench_noop
            PUBLIC
            ${ANDROID_ROOT}/bionic/benchmarks/spawn
    )

    # Add c/c++ flags
    target_compile_options(bench_noop PRIVATE
            -Werror
    )

    target_link_libraries(bench_noop
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bench_noop ####################

#################### libbionic_spawn_benchmark ####################
set(FILES
        ${ANDROID_ROOT}/bionic/benchmarks/spawn/spawn_benchmark.cpp
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
    add_library(libbionic_spawn_benchmark SHARED ${SOURCE_FILES})

    target_include_directories(libbionic_spawn_benchmark
            PUBLIC
            ${ANDROID_ROOT}/bionic/benchmarks/spawn/include
            ${ANDROID_ROOT}/bionic/benchmarks/spawn
    )

    # Add c/c++ flags
    target_compile_options(libbionic_spawn_benchmark PRIVATE
            -Werror
    )

    target_link_libraries(libbionic_spawn_benchmark
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libgoogle-benchmark
            liblog
            libm
    )
endif ()
#################### libbionic_spawn_benchmark ####################

#################### bench_noop_nostl ####################
set(FILES
        ${ANDROID_ROOT}/bionic/benchmarks/spawn/noop.cpp
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
    add_executable(bench_noop_nostl ${SOURCE_FILES})

    target_include_directories(bench_noop_nostl
            PUBLIC
            ${ANDROID_ROOT}/bionic/benchmarks/spawn
    )

    # Add c/c++ flags
    target_compile_options(bench_noop_nostl PRIVATE
            -Werror
    )

    target_link_libraries(bench_noop_nostl
            libc
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bench_noop_nostl ####################

#################### bench_noop_static ####################
set(FILES
        ${ANDROID_ROOT}/bionic/benchmarks/spawn/noop.cpp
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
    add_executable(bench_noop_static ${SOURCE_FILES})

    target_include_directories(bench_noop_static
            PUBLIC
            ${ANDROID_ROOT}/bionic/benchmarks/spawn
    )

    # Add c/c++ flags
    target_compile_options(bench_noop_static PRIVATE
            -Werror
    )

    target_link_libraries(bench_noop_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libunwind
    )
endif ()
#################### bench_noop_static ####################

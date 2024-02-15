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

project(afl-fuzz)

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

#################### afl-showmap ####################
set(FILES
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-showmap.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-common.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-sharedmem.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-forkserver.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-performance.c
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
    add_executable(afl-showmap ${SOURCE_FILES})

    target_include_directories(afl-showmap
            PUBLIC
            ${ANDROID_ROOT}/external/AFLplusplus/include
            ${ANDROID_ROOT}/external/AFLplusplus/instrumentation
            ${ANDROID_ROOT}/external/AFLplusplus
    )

    # Add c/c++ flags
    target_compile_options(afl-showmap PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
            -DDEBUG_BUILD
    )

    target_link_libraries(afl-showmap
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libunwind
    )
endif ()
#################### afl-showmap ####################

#################### aflpp_driver ####################
set(FILES
        ${ANDROID_ROOT}/external/AFLplusplus/utils/aflpp_driver/aflpp_driver.c
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
    add_library(aflpp_driver SHARED ${SOURCE_FILES})

    target_include_directories(aflpp_driver
            PUBLIC
            ${ANDROID_ROOT}/external/AFLplusplus/include
            ${ANDROID_ROOT}/external/AFLplusplus/instrumentation
            ${ANDROID_ROOT}/external/AFLplusplus
    )

    # Add c/c++ flags
    target_compile_options(aflpp_driver PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
            -DDEBUG_BUILD
    )

    target_link_libraries(aflpp_driver
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### aflpp_driver ####################

#################### afl-fuzz ####################
set(FILES
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-bitmap.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-cmplog.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-extras.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-init.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-mutators.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-one.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-python.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-queue.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-redqueen.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-run.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-state.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-stats.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz-statsd.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-fuzz.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-common.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-sharedmem.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-forkserver.c
        ${ANDROID_ROOT}/external/AFLplusplus/src/afl-performance.c
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
    add_executable(afl-fuzz ${SOURCE_FILES})

    target_include_directories(afl-fuzz
            PUBLIC
            ${ANDROID_ROOT}/external/AFLplusplus/include
            ${ANDROID_ROOT}/external/AFLplusplus/instrumentation
            ${ANDROID_ROOT}/external/AFLplusplus
    )

    # Add c/c++ flags
    target_compile_options(afl-fuzz PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
            -DDEBUG_BUILD
    )

    target_link_libraries(afl-fuzz
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### afl-fuzz ####################

#################### afl-compiler-rt ####################
set(FILES
        ${ANDROID_ROOT}/external/AFLplusplus/instrumentation/afl-compiler-rt.o.c
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
    add_library(afl-compiler-rt SHARED ${SOURCE_FILES})

    target_include_directories(afl-compiler-rt
            PUBLIC
            ${ANDROID_ROOT}/external/AFLplusplus/include
            ${ANDROID_ROOT}/external/AFLplusplus/instrumentation
            ${ANDROID_ROOT}/external/AFLplusplus
    )

    # Add c/c++ flags
    target_compile_options(afl-compiler-rt PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
            -DDEBUG_BUILD
    )

    target_link_libraries(afl-compiler-rt
            libc
            libc++
            libc++demangle
            libdl
            libm
    )
endif ()
#################### afl-compiler-rt ####################

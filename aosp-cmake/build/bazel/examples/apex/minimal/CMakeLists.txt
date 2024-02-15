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

project(build.bazel.examples.apex.minimal_dummy_cc_lib)

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

#################### build.bazel.examples.apex.minimal_dummy_cc_lib ####################
set(FILES
        ${ANDROID_ROOT}/build/bazel/examples/apex/minimal/dummy_cc_lib.cc
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
    add_library(build.bazel.examples.apex.minimal_dummy_cc_lib SHARED ${SOURCE_FILES})

    target_include_directories(build.bazel.examples.apex.minimal_dummy_cc_lib
            PUBLIC
            ${ANDROID_ROOT}/build/bazel/examples/apex/minimal
    )

    # Add c/c++ flags
    target_compile_options(build.bazel.examples.apex.minimal_dummy_cc_lib PRIVATE
            -Werror
    )

    target_link_libraries(build.bazel.examples.apex.minimal_dummy_cc_lib
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### build.bazel.examples.apex.minimal_dummy_cc_lib ####################

#################### build.bazel.examples.apex.minimal_dummy_cc_lib_3 ####################
set(FILES
        ${ANDROID_ROOT}/build/bazel/examples/apex/minimal/dummy_cc_lib_3.cc
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
    add_library(build.bazel.examples.apex.minimal_dummy_cc_lib_3 SHARED ${SOURCE_FILES})

    target_include_directories(build.bazel.examples.apex.minimal_dummy_cc_lib_3
            PUBLIC
            ${ANDROID_ROOT}/build/bazel/examples/apex/minimal
    )

    # Add c/c++ flags
    target_compile_options(build.bazel.examples.apex.minimal_dummy_cc_lib_3 PRIVATE
            -Werror
    )

    target_link_libraries(build.bazel.examples.apex.minimal_dummy_cc_lib_3
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### build.bazel.examples.apex.minimal_dummy_cc_lib_3 ####################

#################### build.bazel.examples.apex.minimal_dummy_cc_lib_2 ####################
set(FILES
        ${ANDROID_ROOT}/build/bazel/examples/apex/minimal/dummy_cc_lib_2.cc
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
    add_library(build.bazel.examples.apex.minimal_dummy_cc_lib_2 SHARED ${SOURCE_FILES})

    target_include_directories(build.bazel.examples.apex.minimal_dummy_cc_lib_2
            PUBLIC
            ${ANDROID_ROOT}/build/bazel/examples/apex/minimal
    )

    # Add c/c++ flags
    target_compile_options(build.bazel.examples.apex.minimal_dummy_cc_lib_2 PRIVATE
            -Werror
    )

    target_link_libraries(build.bazel.examples.apex.minimal_dummy_cc_lib_2
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### build.bazel.examples.apex.minimal_dummy_cc_lib_2 ####################

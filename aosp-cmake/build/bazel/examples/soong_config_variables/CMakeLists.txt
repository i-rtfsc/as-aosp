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

project(bp2build_adbd)

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

#################### bp2build_bar ####################
set(FILES
        ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables/main.cpp
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
    add_library(bp2build_bar SHARED ${SOURCE_FILES})

    target_include_directories(bp2build_bar
            PUBLIC
            ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables
    )

    # Add c/c++ flags
    target_compile_options(bp2build_bar PRIVATE
            -Werror
    )

    target_link_libraries(bp2build_bar
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bp2build_bar ####################

#################### bp2build_foo ####################
set(FILES
        ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables/main.cpp
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
    add_library(bp2build_foo SHARED ${SOURCE_FILES})

    target_include_directories(bp2build_foo
            PUBLIC
            ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables
    )

    # Add c/c++ flags
    target_compile_options(bp2build_foo PRIVATE
            -Werror
    )

    target_link_libraries(bp2build_foo
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bp2build_foo ####################

#################### bp2build_baz ####################
set(FILES
        ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables/main.cpp
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
    add_library(bp2build_baz SHARED ${SOURCE_FILES})

    target_include_directories(bp2build_baz
            PUBLIC
            ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables
    )

    # Add c/c++ flags
    target_compile_options(bp2build_baz PRIVATE
            -Werror
    )

    target_link_libraries(bp2build_baz
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bp2build_baz ####################

#################### bp2build_quux ####################
set(FILES
        ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables/main.cpp
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
    add_library(bp2build_quux SHARED ${SOURCE_FILES})

    target_include_directories(bp2build_quux
            PUBLIC
            ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables
    )

    # Add c/c++ flags
    target_compile_options(bp2build_quux PRIVATE
            -Werror
    )

    target_link_libraries(bp2build_quux
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bp2build_quux ####################

#################### bp2build_qux ####################
set(FILES
        ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables/main.cpp
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
    add_library(bp2build_qux SHARED ${SOURCE_FILES})

    target_include_directories(bp2build_qux
            PUBLIC
            ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables
    )

    # Add c/c++ flags
    target_compile_options(bp2build_qux PRIVATE
            -Werror
    )

    target_link_libraries(bp2build_qux
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bp2build_qux ####################

#################### bp2build_adbd ####################
set(FILES
        ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables/main.cpp
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
    add_executable(bp2build_adbd ${SOURCE_FILES})

    target_include_directories(bp2build_adbd
            PUBLIC
            ${ANDROID_ROOT}/build/bazel/examples/soong_config_variables
    )

    # Add c/c++ flags
    target_compile_options(bp2build_adbd PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(bp2build_adbd
            bp2build_bar
            bp2build_foo
            bp2build_quux
            bp2build_qux
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### bp2build_adbd ####################

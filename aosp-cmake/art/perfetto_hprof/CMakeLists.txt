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

project(libperfetto_hprof)

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

#################### libperfetto_hprof ####################
set(FILES
        ${ANDROID_ROOT}/art/perfetto_hprof/perfetto_hprof.cc
        ${OUT_INTERMEDIATES_ROOT}/art/perfetto_hprof/art_perfetto_hprof_operator_srcs/gen/gensrcs/art/perfetto_hprof/perfetto_hprof.operator_out.cc
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
    add_library(libperfetto_hprof SHARED ${SOURCE_FILES})

    target_include_directories(libperfetto_hprof
            PUBLIC
            ${ANDROID_ROOT}/art/perfetto_hprof
            ${OUT_INTERMEDIATES_ROOT}/art/perfetto_hprof/art_perfetto_hprof_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libperfetto_hprof PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libperfetto_hprof
            libart
            libartbase
            libartpalette
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfile
            libdl
            liblog
            libm
            libnativehelper_header_only
            libperfetto_client_experimental
            perfetto_trace_protos
    )
endif ()
#################### libperfetto_hprof ####################

#################### libperfetto_hprofd ####################
set(FILES
        ${ANDROID_ROOT}/art/perfetto_hprof/perfetto_hprof.cc
        ${OUT_INTERMEDIATES_ROOT}/art/perfetto_hprof/art_perfetto_hprof_operator_srcs/gen/gensrcs/art/perfetto_hprof/perfetto_hprof.operator_out.cc
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
    add_library(libperfetto_hprofd SHARED ${SOURCE_FILES})

    target_include_directories(libperfetto_hprofd
            PUBLIC
            ${ANDROID_ROOT}/art/perfetto_hprof
            ${OUT_INTERMEDIATES_ROOT}/art/perfetto_hprof/art_perfetto_hprof_operator_srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(libperfetto_hprofd PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(libperfetto_hprofd
            libartbased
            libartd
            libartpalette
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfiled
            libdl
            liblog
            libm
            libnativehelper_header_only
            libperfetto_client_experimental
            perfetto_trace_protos
    )
endif ()
#################### libperfetto_hprofd ####################

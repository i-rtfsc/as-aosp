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

project(slicer_ndk)

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

#################### slicer_ndk ####################
set(FILES
        ${ANDROID_ROOT}/tools/dexter/slicer/bytecode_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/code_ir.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/common.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/control_flow_graph.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/debuginfo_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_bytecode.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_format.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_ir.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_ir_builder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_utf8.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/instrumentation.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/reader.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/tryblocks_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/writer.cc
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
    add_library(slicer_ndk SHARED ${SOURCE_FILES})

    target_include_directories(slicer_ndk
            PUBLIC
            ${ANDROID_ROOT}/tools/dexter/slicer/export
            ${ANDROID_ROOT}/tools/dexter/slicer
    )

    # Add c/c++ flags
    target_compile_options(slicer_ndk PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(slicer_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
    )
endif ()
#################### slicer_ndk ####################

#################### slicer_ndk_no_rtti ####################
set(FILES
        ${ANDROID_ROOT}/tools/dexter/slicer/bytecode_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/code_ir.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/common.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/control_flow_graph.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/debuginfo_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_bytecode.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_format.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_ir.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_ir_builder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_utf8.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/instrumentation.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/reader.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/tryblocks_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/writer.cc
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
    add_library(slicer_ndk_no_rtti SHARED ${SOURCE_FILES})

    target_include_directories(slicer_ndk_no_rtti
            PUBLIC
            ${ANDROID_ROOT}/tools/dexter/slicer/export
            ${ANDROID_ROOT}/tools/dexter/slicer
    )

    # Add c/c++ flags
    target_compile_options(slicer_ndk_no_rtti PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(slicer_ndk_no_rtti
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
    )
endif ()
#################### slicer_ndk_no_rtti ####################

#################### slicer_no_rtti ####################
set(FILES
        ${ANDROID_ROOT}/tools/dexter/slicer/bytecode_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/code_ir.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/common.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/control_flow_graph.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/debuginfo_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_bytecode.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_format.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_ir.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_ir_builder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_utf8.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/instrumentation.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/reader.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/tryblocks_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/writer.cc
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
    add_library(slicer_no_rtti SHARED ${SOURCE_FILES})

    target_include_directories(slicer_no_rtti
            PUBLIC
            ${ANDROID_ROOT}/tools/dexter/slicer/export
            ${ANDROID_ROOT}/tools/dexter/slicer
    )

    # Add c/c++ flags
    target_compile_options(slicer_no_rtti PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(slicer_no_rtti
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
    )
endif ()
#################### slicer_no_rtti ####################

#################### slicer ####################
set(FILES
        ${ANDROID_ROOT}/tools/dexter/slicer/bytecode_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/code_ir.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/common.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/control_flow_graph.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/debuginfo_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_bytecode.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_format.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_ir.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_ir_builder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/dex_utf8.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/instrumentation.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/reader.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/tryblocks_encoder.cc
        ${ANDROID_ROOT}/tools/dexter/slicer/writer.cc
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
    add_library(slicer SHARED ${SOURCE_FILES})

    target_include_directories(slicer
            PUBLIC
            ${ANDROID_ROOT}/tools/dexter/slicer/export
            ${ANDROID_ROOT}/tools/dexter/slicer
    )

    # Add c/c++ flags
    target_compile_options(slicer PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(slicer
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
    )
endif ()
#################### slicer ####################

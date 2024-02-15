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

project(libtracefast-interpreter)

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

#################### libtracefast-interpreter ####################
set(FILES
        ${ANDROID_ROOT}/art/tools/tracefast-plugin/tracefast.cc
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
    add_library(libtracefast-interpreter SHARED ${SOURCE_FILES})

    target_include_directories(libtracefast-interpreter
            PUBLIC
            ${ANDROID_ROOT}/art/tools/tracefast-plugin
    )

    # Add c/c++ flags
    target_compile_options(libtracefast-interpreter PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libtracefast-interpreter
            libart
            libartbase
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnativehelper_header_only
    )
endif ()
#################### libtracefast-interpreter ####################

#################### libtracefast-trampolined ####################
set(FILES
        ${ANDROID_ROOT}/art/tools/tracefast-plugin/tracefast.cc
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
    add_library(libtracefast-trampolined SHARED ${SOURCE_FILES})

    target_include_directories(libtracefast-trampolined
            PUBLIC
            ${ANDROID_ROOT}/art/tools/tracefast-plugin
    )

    # Add c/c++ flags
    target_compile_options(libtracefast-trampolined PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(libtracefast-trampolined
            libartbased
            libartd
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnativehelper_header_only
    )
endif ()
#################### libtracefast-trampolined ####################

#################### libtracefast-trampoline ####################
set(FILES
        ${ANDROID_ROOT}/art/tools/tracefast-plugin/tracefast.cc
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
    add_library(libtracefast-trampoline SHARED ${SOURCE_FILES})

    target_include_directories(libtracefast-trampoline
            PUBLIC
            ${ANDROID_ROOT}/art/tools/tracefast-plugin
    )

    # Add c/c++ flags
    target_compile_options(libtracefast-trampoline PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libtracefast-trampoline
            libart
            libartbase
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnativehelper_header_only
    )
endif ()
#################### libtracefast-trampoline ####################

#################### libtracefast-interpreterd ####################
set(FILES
        ${ANDROID_ROOT}/art/tools/tracefast-plugin/tracefast.cc
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
    add_library(libtracefast-interpreterd SHARED ${SOURCE_FILES})

    target_include_directories(libtracefast-interpreterd
            PUBLIC
            ${ANDROID_ROOT}/art/tools/tracefast-plugin
    )

    # Add c/c++ flags
    target_compile_options(libtracefast-interpreterd PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(libtracefast-interpreterd
            libartbased
            libartd
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libnativehelper_header_only
    )
endif ()
#################### libtracefast-interpreterd ####################

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

project(libopenjdkjvmtid)

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

#################### libopenjdkjvmtid ####################
set(FILES
        ${ANDROID_ROOT}/art/openjdkjvmti/alloc_manager.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/deopt_manager.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/events.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/object_tagging.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/OpenjdkJvmTi.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_allocator.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_breakpoint.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_class.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_class_definition.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_class_loader.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_ddms.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_dump.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_extension.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_field.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_heap.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_jni.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_logging.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_method.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_monitor.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_object.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_phase.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_properties.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_search.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_stack.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_redefine.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_thread.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_threadgroup.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_timers.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/transform.cc
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
    add_library(libopenjdkjvmtid SHARED ${SOURCE_FILES})

    target_include_directories(libopenjdkjvmtid
            PUBLIC
            ${ANDROID_ROOT}/art/openjdkjvmti
            ${ANDROID_ROOT}/art/openjdkjvmti/include
    )

    # Add c/c++ flags
    target_compile_options(libopenjdkjvmtid PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(libopenjdkjvmtid
            art_cmdlineparser_headers
            libartbased
            libartd
            libartd-compiler
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfiled
            libdl
            libm
            libnativehelper_header_only
            libopenjdkjvmti_headers
    )
endif ()
#################### libopenjdkjvmtid ####################

#################### libopenjdkjvmti ####################
set(FILES
        ${ANDROID_ROOT}/art/openjdkjvmti/alloc_manager.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/deopt_manager.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/events.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/object_tagging.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/OpenjdkJvmTi.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_allocator.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_breakpoint.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_class.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_class_definition.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_class_loader.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_ddms.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_dump.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_extension.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_field.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_heap.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_jni.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_logging.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_method.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_monitor.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_object.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_phase.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_properties.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_search.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_stack.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_redefine.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_thread.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_threadgroup.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/ti_timers.cc
        ${ANDROID_ROOT}/art/openjdkjvmti/transform.cc
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
    add_library(libopenjdkjvmti SHARED ${SOURCE_FILES})

    target_include_directories(libopenjdkjvmti
            PUBLIC
            ${ANDROID_ROOT}/art/openjdkjvmti
            ${ANDROID_ROOT}/art/openjdkjvmti/include
    )

    # Add c/c++ flags
    target_compile_options(libopenjdkjvmti PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libopenjdkjvmti
            art_cmdlineparser_headers
            libart
            libart-compiler
            libartbase
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdexfile
            libdl
            libm
            libnativehelper_header_only
            libopenjdkjvmti_headers
    )
endif ()
#################### libopenjdkjvmti ####################

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

project(kernel_log_monitor)

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

#################### libcuttlefish_kernel_log_monitor_utils ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/kernel_log_monitor/utils.cc
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
    add_library(libcuttlefish_kernel_log_monitor_utils SHARED ${SOURCE_FILES})

    target_include_directories(libcuttlefish_kernel_log_monitor_utils
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/kernel_log_monitor
    )

    # Add c/c++ flags
    target_compile_options(libcuttlefish_kernel_log_monitor_utils PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libcuttlefish_kernel_log_monitor_utils
            cuttlefish_common_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcuttlefish_fs
            libcuttlefish_utils
            libdl
            libjsoncpp
            libm
    )
endif ()
#################### libcuttlefish_kernel_log_monitor_utils ####################

#################### kernel_log_monitor ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/kernel_log_monitor/main.cc
        ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/kernel_log_monitor/kernel_log_server.cc
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
    add_executable(kernel_log_monitor ${SOURCE_FILES})

    target_include_directories(kernel_log_monitor
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/host/commands/kernel_log_monitor
    )

    # Add c/c++ flags
    target_compile_options(kernel_log_monitor PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(kernel_log_monitor
            cuttlefish_common_headers
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcuttlefish_fs
            libcuttlefish_host_config
            libcuttlefish_kernel_log_monitor_utils
            libcuttlefish_utils
            libdl
            libext2_blkid
            libgflags
            libjsoncpp
            libm
    )
endif ()
#################### kernel_log_monitor ####################

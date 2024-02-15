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

project(liblog_for_runtime_apex)

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

#################### liblog_for_runtime_apex ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/liblog/log_event_list.cpp
        ${ANDROID_ROOT}/system/logging/liblog/log_event_write.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logger_name.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logger_read.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logger_write.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logprint.cpp
        ${ANDROID_ROOT}/system/logging/liblog/properties.cpp
        ${ANDROID_ROOT}/system/logging/liblog/event_tag_map.cpp
        ${ANDROID_ROOT}/system/logging/liblog/log_time.cpp
        ${ANDROID_ROOT}/system/logging/liblog/pmsg_reader.cpp
        ${ANDROID_ROOT}/system/logging/liblog/pmsg_writer.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logd_reader.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logd_writer.cpp
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
    add_library(liblog_for_runtime_apex SHARED ${SOURCE_FILES})

    target_include_directories(liblog_for_runtime_apex
            PUBLIC
            ${ANDROID_ROOT}/system/logging/liblog
            ${ANDROID_ROOT}/system/logging/liblog/include
    )

    # Add c/c++ flags
    target_compile_options(liblog_for_runtime_apex PRIVATE
            -Wall
            -Werror
            -Wextra
            -DANDROID_DEBUGGABLE=0
    )

    target_link_libraries(liblog_for_runtime_apex
            libbase_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils_headers
            libdl
            liblog_headers
            libm
            libwinpthread
    )
endif ()
#################### liblog_for_runtime_apex ####################

#################### liblog ####################
set(FILES
        ${ANDROID_ROOT}/system/logging/liblog/log_event_list.cpp
        ${ANDROID_ROOT}/system/logging/liblog/log_event_write.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logger_name.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logger_read.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logger_write.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logprint.cpp
        ${ANDROID_ROOT}/system/logging/liblog/properties.cpp
        ${ANDROID_ROOT}/system/logging/liblog/event_tag_map.cpp
        ${ANDROID_ROOT}/system/logging/liblog/log_time.cpp
        ${ANDROID_ROOT}/system/logging/liblog/pmsg_reader.cpp
        ${ANDROID_ROOT}/system/logging/liblog/pmsg_writer.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logd_reader.cpp
        ${ANDROID_ROOT}/system/logging/liblog/logd_writer.cpp
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
    add_library(liblog SHARED ${SOURCE_FILES})

    target_include_directories(liblog
            PUBLIC
            ${ANDROID_ROOT}/system/logging/liblog
            ${ANDROID_ROOT}/system/logging/liblog/include
    )

    # Add c/c++ flags
    target_compile_options(liblog PRIVATE
            -Wall
            -Werror
            -Wextra
            -DANDROID_DEBUGGABLE=0
    )

    target_link_libraries(liblog
            libbase_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils_headers
            libdl
            liblog_headers
            libm
            libwinpthread
    )
endif ()
#################### liblog ####################

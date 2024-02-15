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

project(incidentd)

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

#################### incidentd ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/Broadcaster.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/FdBuffer.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/IncidentService.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/Privacy.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/PrivacyFilter.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/Reporter.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/Section.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/Throttler.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/WorkDirectory.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/incidentd_util.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/main.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/proto_util.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src/report_directory.cpp
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/incidentd/incidentd/android_${OUT_ARCH_CPU}/gen/proto/frameworks/base/cmds/incidentd/src/report_file.pb.cc
        ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/incidentd/incidentd_section_list/gen/section_list.cpp
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
    add_executable(incidentd ${SOURCE_FILES})

    target_include_directories(incidentd
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/incidentd/src
            ${ANDROID_ROOT}/frameworks/base/cmds/incidentd
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/incidentd/incidentd/android_${OUT_ARCH_CPU}/gen/proto/frameworks/base/cmds/incidentd
            ${OUT_INTERMEDIATES_ROOT}/frameworks/base/cmds/incidentd/incidentd/android_${OUT_ARCH_CPU}/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(incidentd PRIVATE
            -Wall
            -Werror
            -Wno-unused-variable
    )

    target_link_libraries(incidentd
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdebuggerd_client
            libdl
            libdumputils
            libincidentcompanion
            libincidentpriv
            liblog
            libm
            libplatformprotos
            libprotobuf-cpp-lite
            libprotoutil
            libservices
            libutils
    )
endif ()
#################### incidentd ####################

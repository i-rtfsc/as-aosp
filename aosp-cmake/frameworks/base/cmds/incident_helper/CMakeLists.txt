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

project(incident_helper)

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

#################### incident_helper ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers/BatteryTypeParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers/CpuFreqParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers/CpuInfoParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers/EventLogTagsParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers/KernelWakesParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers/PageTypeInfoParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers/ProcrankParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers/PsParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers/SystemPropertiesParser.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/TextParserBase.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/ih_util.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/main.cpp
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
    add_executable(incident_helper ${SOURCE_FILES})

    target_include_directories(incident_helper
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src
            ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper/src/parsers
            ${ANDROID_ROOT}/frameworks/base/cmds/incident_helper
    )

    # Add c/c++ flags
    target_compile_options(incident_helper PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(incident_helper
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libprotoutil
            libutils
    )
endif ()
#################### incident_helper ####################

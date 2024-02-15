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

project(odrefreshd)

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

#################### libodrstatslog ####################
set(FILES
        ${ANDROID_ROOT}/art/odrefresh/odr_metrics_record.cc
        ${ANDROID_ROOT}/art/odrefresh/odr_statslog_android.cc
        ${OUT_INTERMEDIATES_ROOT}/art/odrefresh/statslog_odrefresh.cpp/gen/statslog_odrefresh.cpp
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
    add_library(libodrstatslog SHARED ${SOURCE_FILES})

    target_include_directories(libodrstatslog
            PUBLIC
            ${ANDROID_ROOT}/art/odrefresh/include
            ${ANDROID_ROOT}/art/odrefresh
            ${OUT_INTERMEDIATES_ROOT}/art/odrefresh/statslog_odrefresh.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/art/odrefresh/statslog_odrefresh.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libodrstatslog PRIVATE
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libodrstatslog
            libart_headers
            libartbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libstatssocket
    )
endif ()
#################### libodrstatslog ####################

#################### odrefreshd ####################
set(FILES
        ${ANDROID_ROOT}/art/odrefresh/odrefresh.cc
        ${ANDROID_ROOT}/art/odrefresh/odr_common.cc
        ${ANDROID_ROOT}/art/odrefresh/odr_compilation_log.cc
        ${ANDROID_ROOT}/art/odrefresh/odr_fs_utils.cc
        ${ANDROID_ROOT}/art/odrefresh/odr_metrics.cc
        ${ANDROID_ROOT}/art/odrefresh/odr_metrics_record.cc
        ${ANDROID_ROOT}/art/odrefresh/odrefresh_main.cc
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex-info-list-tinyxml/gen/cpp/com_android_apex.cpp
        ${OUT_INTERMEDIATES_ROOT}/system/apex/apexd/apex-info-list-tinyxml/gen/cpp/com_android_apex_enums.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/odrefresh/art-apex-cache-info/gen/cpp/com_android_art.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/odrefresh/art-apex-cache-info/gen/cpp/com_android_art_enums.cpp
        ${OUT_INTERMEDIATES_ROOT}/art/odrefresh/art-odrefresh-operator-srcs/gen/gensrcs/art/odrefresh/odr_metrics.operator_out.cc
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
    add_executable(odrefreshd ${SOURCE_FILES})

    target_include_directories(odrefreshd
            PUBLIC
            ${ANDROID_ROOT}/art/odrefresh/include
            ${ANDROID_ROOT}/art/odrefresh
            ${OUT_INTERMEDIATES_ROOT}/art/odrefresh/art-apex-cache-info/gen/cpp/include
            ${OUT_INTERMEDIATES_ROOT}/art/odrefresh/art-odrefresh-operator-srcs/gen/gensrcs
    )

    # Add c/c++ flags
    target_compile_options(odrefreshd PRIVATE
            -Wall
            -Werror
            -Wextra
            -DVIXL_DEBUG
            -UNDEBUG
    )

    target_link_libraries(odrefreshd
            dexoptanalyzer_headers
            libartbased
            libartd
            libartpalette
            libarttools
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdexfile
            libdl
            liblog
            libm
            libmodules-utils-build
    )
endif ()
#################### odrefreshd ####################

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

project(libexpresslog)

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

#################### libexpresslog ####################
set(FILES
        ${ANDROID_ROOT}/system/core/libstats/expresslog/Counter.cpp
        ${ANDROID_ROOT}/system/core/libstats/expresslog/Histogram.cpp
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
    add_library(libexpresslog SHARED ${SOURCE_FILES})

    target_include_directories(libexpresslog
            PUBLIC
            ${ANDROID_ROOT}/system/core/libstats/expresslog/include
            ${ANDROID_ROOT}/system/core/libstats/expresslog
            ${OUT_INTERMEDIATES_ROOT}/system/core/libstats/expresslog/statslog_express.h/gen
    )

    # Add c/c++ flags
    target_compile_options(libexpresslog PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libexpresslog
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            libstatslog_express
            libstatssocket
            libtextclassifier_hash_headers
            libtextclassifier_hash_static
    )
endif ()
#################### libexpresslog ####################

#################### libstatslog_express ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/system/core/libstats/expresslog/statslog_express.cpp/gen/statslog_express.cpp
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
    add_library(libstatslog_express SHARED ${SOURCE_FILES})

    target_include_directories(libstatslog_express
            PUBLIC
            ${ANDROID_ROOT}/system/core/libstats/expresslog
            ${OUT_INTERMEDIATES_ROOT}/system/core/libstats/expresslog/statslog_express.cpp/gen
            ${OUT_INTERMEDIATES_ROOT}/system/core/libstats/expresslog/statslog_express.h/gen
    )

    target_link_libraries(libstatslog_express
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libstatssocket
    )
endif ()
#################### libstatslog_express ####################

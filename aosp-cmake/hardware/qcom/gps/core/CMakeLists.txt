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

project(libloc_core)

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

#################### libloc_core ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/gps/core/LocApiBase.cpp
        ${ANDROID_ROOT}/hardware/qcom/gps/core/LocAdapterBase.cpp
        ${ANDROID_ROOT}/hardware/qcom/gps/core/ContextBase.cpp
        ${ANDROID_ROOT}/hardware/qcom/gps/core/LocContext.cpp
        ${ANDROID_ROOT}/hardware/qcom/gps/core/loc_core_log.cpp
        ${ANDROID_ROOT}/hardware/qcom/gps/core/data-items/DataItemsFactoryProxy.cpp
        ${ANDROID_ROOT}/hardware/qcom/gps/core/data-items/DataItemConcreteTypes.cpp
        ${ANDROID_ROOT}/hardware/qcom/gps/core/SystemStatusOsObserver.cpp
        ${ANDROID_ROOT}/hardware/qcom/gps/core/SystemStatus.cpp
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
    add_library(libloc_core SHARED ${SOURCE_FILES})

    target_include_directories(libloc_core
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/gps/core/data-items
            ${ANDROID_ROOT}/hardware/qcom/gps/core/observer
            ${ANDROID_ROOT}/hardware/qcom/gps/core
    )

    # Add c/c++ flags
    target_compile_options(libloc_core PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libloc_core
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libgps.utils
            libgps.utils_headers
            libloc_pla_headers
            liblocation_api_headers
            liblog
            libm
            libutils
            libutils_headers
            libc
            libdl
            libm
    )
endif ()
#################### libloc_core ####################

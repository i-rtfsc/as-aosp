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

project(libp61-jcop-kit)

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

#################### libp61-jcop-kit ####################
set(FILES
        ${ANDROID_ROOT}/external/libese/third_party/NXPNFC_P61_JCOP_Kit/src/AlaLib.cpp
        ${ANDROID_ROOT}/external/libese/third_party/NXPNFC_P61_JCOP_Kit/src/JcopOsDownload.cpp
        ${ANDROID_ROOT}/external/libese/third_party/NXPNFC_P61_JCOP_Kit/src/JcDnld.cpp
        ${ANDROID_ROOT}/external/libese/third_party/NXPNFC_P61_JCOP_Kit/src/Ala.cpp
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
    add_library(libp61-jcop-kit SHARED ${SOURCE_FILES})

    target_include_directories(libp61-jcop-kit
            PUBLIC
            ${ANDROID_ROOT}/external/libese/third_party/NXPNFC_P61_JCOP_Kit/include
            ${ANDROID_ROOT}/external/libese/third_party/NXPNFC_P61_JCOP_Kit/inc
            ${ANDROID_ROOT}/external/libese/third_party/NXPNFC_P61_JCOP_Kit
    )

    # Add c/c++ flags
    target_compile_options(libp61-jcop-kit PRIVATE
            -Werror
            -Wall
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libp61-jcop-kit
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libp61-jcop-kit ####################

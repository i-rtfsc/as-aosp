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

project(android.hardware.secure_element.thales.libse)

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

#################### android.hardware.secure_element.thales.libse ####################
set(FILES
        ${ANDROID_ROOT}/hardware/st/secure_element2/secure_element/esehal/src/checksum.c
        ${ANDROID_ROOT}/hardware/st/secure_element2/secure_element/esehal/src/iso7816_t1.c
        ${ANDROID_ROOT}/hardware/st/secure_element2/secure_element/esehal/src/libse-gto.c
        ${ANDROID_ROOT}/hardware/st/secure_element2/secure_element/esehal/src/spi.c
        ${ANDROID_ROOT}/hardware/st/secure_element2/secure_element/esehal/src/transport.c
        ${ANDROID_ROOT}/hardware/st/secure_element2/secure_element/esehal/src/log.c
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
    add_library(android.hardware.secure_element.thales.libse SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.secure_element.thales.libse
            PUBLIC
            ${ANDROID_ROOT}/hardware/st/secure_element2/secure_element/esehal/src
            ${ANDROID_ROOT}/hardware/st/secure_element2/secure_element/esehal
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.secure_element.thales.libse PRIVATE
            -DENABLE_DEBUG=1
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(android.hardware.secure_element.thales.libse
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.secure_element.thales.libse ####################

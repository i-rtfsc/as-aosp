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

project(ese_spi_st)

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

#################### ese_spi_st ####################
set(FILES
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/SpiLayerDriver.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/SpiLayerInterface.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/SpiLayerComm.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/StEseApi.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/T1protocol.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/Atp.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/Iso13239CRC.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/Tpdu.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/Utils.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/ese_config.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/config.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/android_logmsg.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/DataMgmt.cc
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
    add_library(ese_spi_st SHARED ${SOURCE_FILES})

    target_include_directories(ese_spi_st
            PUBLIC
            ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib
            ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver
    )

    # Add c/c++ flags
    target_compile_options(ese_spi_st PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(ese_spi_st
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### ese_spi_st ####################

#################### ese_spi_st_static ####################
set(FILES
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/SpiLayerDriver.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/SpiLayerInterface.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/SpiLayerComm.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/StEseApi.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/T1protocol.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/Atp.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/Iso13239CRC.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/Tpdu.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/Utils.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/ese_config.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/config.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/android_logmsg.cc
        ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib/DataMgmt.cc
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
    add_library(ese_spi_st_static SHARED ${SOURCE_FILES})

    target_include_directories(ese_spi_st_static
            PUBLIC
            ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver
            ${ANDROID_ROOT}/hardware/st/secure_element/ese-spi-driver/utils-lib
    )

    # Add c/c++ flags
    target_compile_options(ese_spi_st_static PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(ese_spi_st_static
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
    )
endif ()
#################### ese_spi_st_static ####################

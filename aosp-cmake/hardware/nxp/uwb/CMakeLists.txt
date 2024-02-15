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

project(uwb_uci.helios)

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

#################### uwb_uci.helios ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/fwd/sr1xx/phNxpUciHal_fwd.cc
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/hal/phNxpUciHal.cc
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/hal/phNxpUciHal_ext.cc
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/log/phNxpLog.cc
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/tml/phDal4Uwb_messageQueueLib.cc
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/tml/phOsalUwb_Timer.cc
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/tml/phTmlUwb.cc
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/tml/phTmlUwb_spi.cc
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/utils/phNxpConfig.cc
        ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/utils/phNxpUciHal_utils.cc
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
    add_library(uwb_uci.helios SHARED ${SOURCE_FILES})

    target_include_directories(uwb_uci.helios
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/fwd/sr1xx
            ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/inc
            ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/inc/common
            ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/hal
            ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/log
            ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/tml
            ${ANDROID_ROOT}/hardware/nxp/uwb/halimpl/utils
            ${ANDROID_ROOT}/hardware/nxp/uwb/extns/inc
            ${ANDROID_ROOT}/hardware/nxp/uwb
    )

    # Add c/c++ flags
    target_compile_options(uwb_uci.helios PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(uwb_uci.helios
            android.hardware.uwb-V1-ndk
            libbase
            libc++
            libc++demangle
            libc
            libchrome
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### uwb_uci.helios ####################

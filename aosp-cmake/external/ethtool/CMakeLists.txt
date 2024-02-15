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

project(ethtool)

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

#################### ethtool ####################
set(FILES
        ${ANDROID_ROOT}/external/ethtool/amd8111e.c
        ${ANDROID_ROOT}/external/ethtool/at76c50x-usb.c
        ${ANDROID_ROOT}/external/ethtool/de2104x.c
        ${ANDROID_ROOT}/external/ethtool/dsa.c
        ${ANDROID_ROOT}/external/ethtool/e100.c
        ${ANDROID_ROOT}/external/ethtool/e1000.c
        ${ANDROID_ROOT}/external/ethtool/et131x.c
        ${ANDROID_ROOT}/external/ethtool/ethtool.c
        ${ANDROID_ROOT}/external/ethtool/fec.c
        ${ANDROID_ROOT}/external/ethtool/fec_8xx.c
        ${ANDROID_ROOT}/external/ethtool/fjes.c
        ${ANDROID_ROOT}/external/ethtool/ibm_emac.c
        ${ANDROID_ROOT}/external/ethtool/igb.c
        ${ANDROID_ROOT}/external/ethtool/ixgb.c
        ${ANDROID_ROOT}/external/ethtool/ixgbe.c
        ${ANDROID_ROOT}/external/ethtool/ixgbevf.c
        ${ANDROID_ROOT}/external/ethtool/lan78xx.c
        ${ANDROID_ROOT}/external/ethtool/marvell.c
        ${ANDROID_ROOT}/external/ethtool/natsemi.c
        ${ANDROID_ROOT}/external/ethtool/pcnet32.c
        ${ANDROID_ROOT}/external/ethtool/qsfp.c
        ${ANDROID_ROOT}/external/ethtool/realtek.c
        ${ANDROID_ROOT}/external/ethtool/rxclass.c
        ${ANDROID_ROOT}/external/ethtool/sfc.c
        ${ANDROID_ROOT}/external/ethtool/sff-common.c
        ${ANDROID_ROOT}/external/ethtool/sfpdiag.c
        ${ANDROID_ROOT}/external/ethtool/sfpid.c
        ${ANDROID_ROOT}/external/ethtool/smsc911x.c
        ${ANDROID_ROOT}/external/ethtool/stmmac.c
        ${ANDROID_ROOT}/external/ethtool/tg3.c
        ${ANDROID_ROOT}/external/ethtool/tse.c
        ${ANDROID_ROOT}/external/ethtool/vioc.c
        ${ANDROID_ROOT}/external/ethtool/vmxnet3.c
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
    add_executable(ethtool ${SOURCE_FILES})

    target_include_directories(ethtool
            PUBLIC
            ${ANDROID_ROOT}/external/ethtool
    )

    # Add c/c++ flags
    target_compile_options(ethtool PRIVATE
            -Werror
    )

    target_link_libraries(ethtool
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### ethtool ####################

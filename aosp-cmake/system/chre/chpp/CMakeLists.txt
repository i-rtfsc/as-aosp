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

project(chre_chpp_linux)

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

#################### chre_chpp_linux ####################
set(FILES
        ${ANDROID_ROOT}/system/chre/chpp/clients/discovery.c
        ${ANDROID_ROOT}/system/chre/chpp/clients/loopback.c
        ${ANDROID_ROOT}/system/chre/chpp/clients/timesync.c
        ${ANDROID_ROOT}/system/chre/chpp/clients/gnss.c
        ${ANDROID_ROOT}/system/chre/chpp/clients/wifi.c
        ${ANDROID_ROOT}/system/chre/chpp/clients/wwan.c
        ${ANDROID_ROOT}/system/chre/chpp/common/gnss_convert.c
        ${ANDROID_ROOT}/system/chre/chpp/common/wifi_convert.c
        ${ANDROID_ROOT}/system/chre/chpp/common/wifi_utils.c
        ${ANDROID_ROOT}/system/chre/chpp/common/wwan_convert.c
        ${ANDROID_ROOT}/system/chre/chpp/platform/linux/services/platform_gnss.c
        ${ANDROID_ROOT}/system/chre/chpp/services/gnss.c
        ${ANDROID_ROOT}/system/chre/chpp/services/wifi.c
        ${ANDROID_ROOT}/system/chre/chpp/services/wwan.c
        ${ANDROID_ROOT}/system/chre/chpp/transport.c
        ${ANDROID_ROOT}/system/chre/chpp/app.c
        ${ANDROID_ROOT}/system/chre/chpp/clients.c
        ${ANDROID_ROOT}/system/chre/chpp/platform/pal_api.c
        ${ANDROID_ROOT}/system/chre/chpp/platform/linux/memory.c
        ${ANDROID_ROOT}/system/chre/chpp/platform/linux/notifier.c
        ${ANDROID_ROOT}/system/chre/chpp/platform/shared/crc.c
        ${ANDROID_ROOT}/system/chre/chpp/services.c
        ${ANDROID_ROOT}/system/chre/chpp/services/discovery.c
        ${ANDROID_ROOT}/system/chre/chpp/services/loopback.c
        ${ANDROID_ROOT}/system/chre/chpp/services/nonhandle.c
        ${ANDROID_ROOT}/system/chre/chpp/services/timesync.c
        ${ANDROID_ROOT}/system/chre/chpp/platform/linux/link.c
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
    add_library(chre_chpp_linux SHARED ${SOURCE_FILES})

    target_include_directories(chre_chpp_linux
            PUBLIC
            ${ANDROID_ROOT}/system/chre/chpp/platform/linux/include
            ${ANDROID_ROOT}/system/chre/chpp/include
            ${ANDROID_ROOT}/system/chre/chpp
    )

    # Add c/c++ flags
    target_compile_options(chre_chpp_linux PRIVATE
            -Wall
            -Werror
            -Wextra
            -DCHPP_DEBUG_ASSERT_ENABLED
    )

    target_link_libraries(chre_chpp_linux
            chre_api
            chre_pal
            chre_pal_linux
            libc++_static
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### chre_chpp_linux ####################

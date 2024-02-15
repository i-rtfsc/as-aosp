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

project(libwifi-hal-common)

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

#################### libwifi-hal-common ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/opt/net/wifi/libwifi_hal/wifi_hal_common.cpp
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
    add_library(libwifi-hal-common SHARED ${SOURCE_FILES})

    target_include_directories(libwifi-hal-common
            PUBLIC
            ${ANDROID_ROOT}/frameworks/opt/net/wifi/libwifi_hal/include
            ${ANDROID_ROOT}/frameworks/opt/net/wifi/libwifi_hal
    )

    # Add c/c++ flags
    target_compile_options(libwifi-hal-common PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libwifi-hal-common
            libbase
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils_headers
            libdl
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libwifi-hal-common ####################

#################### libwifi-hal-fallback ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/opt/net/wifi/libwifi_hal/wifi_hal_fallback.cpp
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
    add_library(libwifi-hal-fallback SHARED ${SOURCE_FILES})

    target_include_directories(libwifi-hal-fallback
            PUBLIC
            ${ANDROID_ROOT}/frameworks/opt/net/wifi/libwifi_hal
    )

    # Add c/c++ flags
    target_compile_options(libwifi-hal-fallback PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libwifi-hal-fallback
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libhardware_legacy_headers
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libwifi-hal-fallback ####################

#################### libwifi-hal ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/opt/net/wifi/libwifi_hal/driver_tool.cpp
        ${ANDROID_ROOT}/frameworks/opt/net/wifi/libwifi_hal/hal_tool.cpp
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
    add_library(libwifi-hal SHARED ${SOURCE_FILES})

    target_include_directories(libwifi-hal
            PUBLIC
            ${ANDROID_ROOT}/frameworks/opt/net/wifi/libwifi_hal/include
            ${ANDROID_ROOT}/frameworks/opt/net/wifi/libwifi_hal
    )

    # Add c/c++ flags
    target_compile_options(libwifi-hal PRIVATE
            -Wall
            -Werror
            -Wextra
            -Wno-unused-function
            -Wno-unused-parameter
    )

    target_link_libraries(libwifi-hal
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcld80211
            libcrypto
            libcutils
            libdl
            libhardware_legacy_headers
            liblog
            libm
            libnl
            libpasn
            libutils
            libwifi-hal-common
            libwifi-hal-qcom
            libc
            libdl
            libm
    )
endif ()
#################### libwifi-hal ####################

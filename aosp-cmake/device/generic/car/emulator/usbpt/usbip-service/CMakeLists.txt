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

project(usbip_service)

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

#################### usbip_service ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/car/emulator/usbpt/usbip-service/UsbIpService.cpp
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
    add_executable(usbip_service ${SOURCE_FILES})

    target_include_directories(usbip_service
            PUBLIC
            ${ANDROID_ROOT}/device/generic/car/emulator/usbpt/usbip-service
    )

    # Add c/c++ flags
    target_compile_options(usbip_service PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(usbip_service
            libbase
            libbase
            libc++
            libc++
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libcutils
            libcutils
            libdl
            liblog
            liblog
            libm
            libutils
            libutils
            usbip_utils
            libc
            libdl
            libm
    )
endif ()
#################### usbip_service ####################

#################### usbip_utils ####################
set(FILES
        ${ANDROID_ROOT}/device/generic/car/emulator/usbpt/usbip-service/UsbIpUtils.cpp
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
    add_library(usbip_utils SHARED ${SOURCE_FILES})

    target_include_directories(usbip_utils
            PUBLIC
            ${ANDROID_ROOT}/device/generic/car/emulator/usbpt/usbip-service
    )

    # Add c/c++ flags
    target_compile_options(usbip_utils PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(usbip_utils
            libbase
            libbase
            libc++
            libc++
            libc
            libclang_rt.builtins
            libclang_rt.builtins
            libcutils
            libcutils
            libdl
            liblog
            liblog
            libm
            libutils
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### usbip_utils ####################

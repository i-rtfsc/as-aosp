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

project(vendor.qti.hardware.display.demura-service)

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

#################### libfilefinder ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/oem_services/file_finder_oem_extension.cpp
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
    add_library(libfilefinder SHARED ${SOURCE_FILES})

    target_include_directories(libfilefinder
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/oem_services
    )

    # Add c/c++ flags
    target_compile_options(libfilefinder PRIVATE
            -Wall
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libfilefinder
            display_headers
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdisplaydebug
            libdl
            liblog
            libm
            libutils
            qti_kernel_headers
            vendor.qti.hardware.display.demura-V1-ndk
            libc
            libdl
            libm
    )
endif ()
#################### libfilefinder ####################

#################### vendor.qti.hardware.display.demura-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/oem_services/demura_file_finder.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/oem_services/service.cpp
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
    add_executable(vendor.qti.hardware.display.demura-service ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.display.demura-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/oem_services
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.display.demura-service PRIVATE
            -Wall
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(vendor.qti.hardware.display.demura-service
            display_headers
            libbinder
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdisplaydebug
            libdl
            libfilefinder
            liblog
            libm
            libsdmutils
            libutils
            qti_kernel_headers
            vendor.qti.hardware.display.demura-V1-ndk
            libc
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.display.demura-service ####################

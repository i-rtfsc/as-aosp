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

project(libese-hw-nxp-pn80t-spidev)

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

#################### libese-hw-nxp-pn80t-spidev ####################
set(FILES
        ${ANDROID_ROOT}/external/libese/libese-hw/nxp/pn80t/linux_spidev.c
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
    add_library(libese-hw-nxp-pn80t-spidev SHARED ${SOURCE_FILES})

    target_include_directories(libese-hw-nxp-pn80t-spidev
            PUBLIC
            ${ANDROID_ROOT}/external/libese/libese-hw/nxp/include
            ${ANDROID_ROOT}/external/libese/libese-hw/nxp
    )

    # Add c/c++ flags
    target_compile_options(libese-hw-nxp-pn80t-spidev PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libese-hw-nxp-pn80t-spidev
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libese
            libese-api-headers
            libese-hw-nxp-pn80t-common
            libese-sysdeps
            libese-teq1
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libese-hw-nxp-pn80t-spidev ####################

#################### libese-hw-nxp-pn80t-nq-nci ####################
set(FILES
        ${ANDROID_ROOT}/external/libese/libese-hw/nxp/pn80t/nq_nci.c
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
    add_library(libese-hw-nxp-pn80t-nq-nci SHARED ${SOURCE_FILES})

    target_include_directories(libese-hw-nxp-pn80t-nq-nci
            PUBLIC
            ${ANDROID_ROOT}/external/libese/libese-hw/nxp/include
            ${ANDROID_ROOT}/external/libese/libese-hw/nxp
    )

    # Add c/c++ flags
    target_compile_options(libese-hw-nxp-pn80t-nq-nci PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(libese-hw-nxp-pn80t-nq-nci
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libese
            libese-api-headers
            libese-hw-nxp-pn80t-common
            libese-sysdeps
            libese-teq1
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libese-hw-nxp-pn80t-nq-nci ####################

#################### libese-hw-nxp-pn80t-common ####################
set(FILES
        ${ANDROID_ROOT}/external/libese/libese-hw/nxp/pn80t/common.c
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
    add_library(libese-hw-nxp-pn80t-common SHARED ${SOURCE_FILES})

    target_include_directories(libese-hw-nxp-pn80t-common
            PUBLIC
            ${ANDROID_ROOT}/external/libese/libese-hw/nxp/include
            ${ANDROID_ROOT}/external/libese/libese-hw/nxp
    )

    # Add c/c++ flags
    target_compile_options(libese-hw-nxp-pn80t-common PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(libese-hw-nxp-pn80t-common
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libese
            libese-sysdeps
            libese-teq1
            liblog
            libm
            libc
            libdl
            libm
    )
endif ()
#################### libese-hw-nxp-pn80t-common ####################

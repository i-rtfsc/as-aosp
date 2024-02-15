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

project(sn100nfc_utils_ringbuffer_fuzzer)

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

#################### libsn100nfcutils ####################
set(FILES
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/utils/config.cc
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/utils/ringbuffer.cc
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
    add_library(libsn100nfcutils SHARED ${SOURCE_FILES})

    target_include_directories(libsn100nfcutils
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/utils/include
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/utils
    )

    # Add c/c++ flags
    target_compile_options(libsn100nfcutils PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libsn100nfcutils
            libbase
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libsn100nfcutils ####################

#################### sn100nfc_utils_ringbuffer_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/utils/test/ringbuffer_fuzzer/ringbuffer_fuzzer.cpp
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
    add_executable(sn100nfc_utils_ringbuffer_fuzzer ${SOURCE_FILES})

    target_include_directories(sn100nfc_utils_ringbuffer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/utils
            ${ANDROID_ROOT}/vendor/nxp/opensource/commonsys/external/libnfc-nci/SN100x/utils/include
    )

    target_link_libraries(sn100nfc_utils_ringbuffer_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libsn100nfcutils
    )
endif ()
#################### sn100nfc_utils_ringbuffer_fuzzer ####################

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

project(android.hardware.secure_element@1.1-service)

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

#################### android.hardware.secure_element-1.1-service-disabled ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/1.1/NxpEseService.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/1.1/SecureElement.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl/NxpEse.cpp
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
    add_executable(android.hardware.secure_element-1.1-service-disabled ${SOURCE_FILES})

    target_include_directories(android.hardware.secure_element-1.1-service-disabled
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/common/include
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/common
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/inc
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/pal
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/ls_client/inc
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.secure_element-1.1-service-disabled PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.secure_element-1.1-service-disabled
            android.hardware.secure_element-1.0
            android.hardware.secure_element-1.1
            ese_spi_nxp
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            ls_client
            vendor.nxp.nxpese-1.0
            vendor.nxp.nxpnfc-1.0
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.secure_element-1.1-service-disabled ####################

#################### android.hardware.secure_element-1.0-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/1.0/NxpEseService.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/1.0/SecureElement.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl/NxpEse.cpp
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
    add_executable(android.hardware.secure_element-1.0-service ${SOURCE_FILES})

    target_include_directories(android.hardware.secure_element-1.0-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/common/include
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/common
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/inc
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/pal
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/ls_client/inc
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.secure_element-1.0-service PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.secure_element-1.0-service
            android.hardware.secure_element-1.0
            ese_spi_nxp
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            ls_client
            vendor.nxp.nxpese-1.0
            vendor.nxp.nxpnfc-1.0
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.secure_element-1.0-service ####################

#################### android.hardware.secure_element-1.0-service-disabled ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/1.0/NxpEseService.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/1.0/SecureElement.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl/NxpEse.cpp
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
    add_executable(android.hardware.secure_element-1.0-service-disabled ${SOURCE_FILES})

    target_include_directories(android.hardware.secure_element-1.0-service-disabled
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/common/include
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/common
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/inc
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/pal
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/ls_client/inc
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.secure_element-1.0-service-disabled PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.secure_element-1.0-service-disabled
            android.hardware.secure_element-1.0
            ese_spi_nxp
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            ls_client
            vendor.nxp.nxpese-1.0
            vendor.nxp.nxpnfc-1.0
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.secure_element-1.0-service-disabled ####################

#################### ls_client ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/ls_client/src/LsLib.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/ls_client/src/LsClient.cpp
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
    add_library(ls_client SHARED ${SOURCE_FILES})

    target_include_directories(ls_client
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/ls_client/inc
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/common/include
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/common
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/inc
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/pal
    )

    # Add c/c++ flags
    target_compile_options(ls_client PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(ls_client
            ese_spi_nxp
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcrypto
            libcutils
            libdl
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### ls_client ####################

#################### ese_spi_nxp ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/lib/phNxpEseDataMgr.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/lib/phNxpEseProto7816_3.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/lib/phNxpEse_Api.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/pal/phNxpEsePal.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/pal/spi/phNxpEsePal_spi.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/spm/phNxpEse_Spm.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/utils/ese_config.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/utils/config.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/utils/ringbuffer.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/src/adaptation/NfcAdaptation.cpp
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
    add_library(ese_spi_nxp SHARED ${SOURCE_FILES})

    target_include_directories(ese_spi_nxp
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/common/include
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/common
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/inc
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/pal
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/lib
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/pal/spi
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/utils
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/src/include
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x
    )

    # Add c/c++ flags
    target_compile_options(ese_spi_nxp PRIVATE
            -Wall
            -Werror
            -Wall
            -Werror
    )

    target_link_libraries(ese_spi_nxp
            android.hardware.nfc-1.0
            android.hardware.nfc-1.1
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            vendor.nxp.nxpese-1.0
            vendor.nxp.nxpnfc-1.0
            libc
            libdl
            libm
    )
endif ()
#################### ese_spi_nxp ####################

#################### android.hardware.secure_element-1.1-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/1.1/NxpEseService.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/1.1/SecureElement.cpp
        ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl/NxpEse.cpp
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
    add_executable(android.hardware.secure_element-1.1-service ${SOURCE_FILES})

    target_include_directories(android.hardware.secure_element-1.1-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/extns/impl
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/common/include
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/common
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/inc
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/libese-spi/p73/pal
            ${ANDROID_ROOT}/hardware/nxp/secure_element/pn8x/ls_client/inc
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.secure_element-1.1-service PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.secure_element-1.1-service
            android.hardware.secure_element-1.0
            android.hardware.secure_element-1.1
            ese_spi_nxp
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            ls_client
            vendor.nxp.nxpese-1.0
            vendor.nxp.nxpnfc-1.0
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.secure_element-1.1-service ####################

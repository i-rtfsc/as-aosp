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

project(android.hardware.keymaster@4.0-service.trusty)

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

#################### libtrusty_ipc ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/keymaster/ipc/trusty_keymaster_ipc.cpp
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
    add_library(libtrusty_ipc SHARED ${SOURCE_FILES})

    target_include_directories(libtrusty_ipc
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/keymaster/include
            ${ANDROID_ROOT}/system/core/trusty/keymaster
    )

    # Add c/c++ flags
    target_compile_options(libtrusty_ipc PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libtrusty_ipc
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            libhardware
            libkeymaster_messages
            liblog
            libm
            libtrusty
            libutils
            libxml2
            libc
            libdl
            libm
    )
endif ()
#################### libtrusty_ipc ####################

#################### android.hardware.security.keymint-service.trusty ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/keymaster/TrustyKeymaster.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/ipc/trusty_keymaster_ipc.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/keymint/TrustyKeyMintDevice.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/keymint/TrustyKeyMintOperation.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/keymint/TrustyRemotelyProvisionedComponentDevice.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/keymint/TrustySecureClock.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/keymint/TrustySharedSecret.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/keymint/service.cpp
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
    add_executable(android.hardware.security.keymint-service.trusty ${SOURCE_FILES})

    target_include_directories(android.hardware.security.keymint-service.trusty
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/keymaster/include
            ${ANDROID_ROOT}/system/core/trusty/keymaster
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.security.keymint-service.trusty PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(android.hardware.security.keymint-service.trusty
            android.hardware.security.keymint-V3-ndk
            android.hardware.security.rkp-V3-ndk
            android.hardware.security.secureclock-V1-ndk
            android.hardware.security.sharedsecret-V1-ndk
            lib_android_keymaster_keymint_utils
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libhardware
            libkeymaster_messages
            libkeymint
            liblog
            libm
            libtrusty
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.security.keymint-service.trusty ####################

#################### android.hardware.keymaster-3.0-service.trusty ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/keymaster/3.0/service.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/3.0/TrustyKeymaster3Device.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/ipc/trusty_keymaster_ipc.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/TrustyKeymaster.cpp
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
    add_executable(android.hardware.keymaster-3.0-service.trusty ${SOURCE_FILES})

    target_include_directories(android.hardware.keymaster-3.0-service.trusty
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/keymaster/include
            ${ANDROID_ROOT}/system/core/trusty/keymaster
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.keymaster-3.0-service.trusty PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.keymaster-3.0-service.trusty
            android.hardware.keymaster-3.0
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            libkeymaster3device
            libkeymaster_messages
            liblog
            libm
            libtrusty
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.keymaster-3.0-service.trusty ####################

#################### android.hardware.keymaster-4.0-service.trusty ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/keymaster/4.0/service.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/4.0/TrustyKeymaster4Device.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/ipc/trusty_keymaster_ipc.cpp
        ${ANDROID_ROOT}/system/core/trusty/keymaster/TrustyKeymaster.cpp
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
    add_executable(android.hardware.keymaster-4.0-service.trusty ${SOURCE_FILES})

    target_include_directories(android.hardware.keymaster-4.0-service.trusty
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/keymaster/include
            ${ANDROID_ROOT}/system/core/trusty/keymaster
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.keymaster-4.0-service.trusty PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.keymaster-4.0-service.trusty
            android.hardware.keymaster-4.0
            libbase
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libhardware
            libhidlbase
            libkeymaster4
            libkeymaster_messages
            liblog
            libm
            libtrusty
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.keymaster-4.0-service.trusty ####################

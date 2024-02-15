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

project(android.hardware.confirmationui-service.trusty)

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

#################### android.hardware.confirmationui.not-so-secure-input ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/confirmationui/NotSoSecureInput.cpp
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
    add_library(android.hardware.confirmationui.not-so-secure-input SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.confirmationui.not-so-secure-input
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/confirmationui
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.confirmationui.not-so-secure-input PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.confirmationui.not-so-secure-input
            libbase
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcrypto
            libdl
            libm
            libteeui_hal_support
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.confirmationui.not-so-secure-input ####################

#################### android.hardware.confirmationui-service.trusty_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/confirmationui/fuzzer.cpp
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
    add_executable(android.hardware.confirmationui-service.trusty_fuzzer ${SOURCE_FILES})

    target_include_directories(android.hardware.confirmationui-service.trusty_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/confirmationui
            ${ANDROID_ROOT}/system/core/trusty/confirmationui/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.confirmationui-service.trusty_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.confirmationui-service.trusty_fuzzer
            android.hardware.confirmationui-V1-ndk
            android.hardware.confirmationui-lib.trusty
            android.hardware.confirmationui.not-so-secure-input
            libbase
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.confirmationui-service.trusty_fuzzer ####################

#################### android.hardware.confirmationui-lib.trusty ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/confirmationui/TrustyApp.cpp
        ${ANDROID_ROOT}/system/core/trusty/confirmationui/TrustyConfirmationUI.cpp
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
    add_library(android.hardware.confirmationui-lib.trusty SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.confirmationui-lib.trusty
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/confirmationui/include
            ${ANDROID_ROOT}/system/core/trusty/confirmationui
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.confirmationui-lib.trusty PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.confirmationui-lib.trusty
            android.hardware.confirmationui-V1-ndk
            android.hardware.security.keymint-V3-ndk
            libbase
            libbinder_ndk
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            libdmabufheap
            libm
            libteeui_hal_support
            libtrusty
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.confirmationui-lib.trusty ####################

#################### android.hardware.confirmationui-service.trusty ####################
set(FILES
        ${ANDROID_ROOT}/system/core/trusty/confirmationui/service.cpp
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
    add_executable(android.hardware.confirmationui-service.trusty ${SOURCE_FILES})

    target_include_directories(android.hardware.confirmationui-service.trusty
            PUBLIC
            ${ANDROID_ROOT}/system/core/trusty/confirmationui
            ${ANDROID_ROOT}/system/core/trusty/confirmationui/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.confirmationui-service.trusty PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.confirmationui-service.trusty
            android.hardware.confirmationui-V1-ndk
            android.hardware.confirmationui-lib.trusty
            android.hardware.confirmationui.not-so-secure-input
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libhidlbase
            libm
            libteeui_hal_support
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.confirmationui-service.trusty ####################

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

project(android.hardware.health-service.cuttlefish_recovery)

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

#################### android.hardware.health-service.cuttlefish ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/health/health-aidl.cpp
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
    add_executable(android.hardware.health-service.cuttlefish ${SOURCE_FILES})

    target_include_directories(android.hardware.health-service.cuttlefish
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/health
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-service.cuttlefish PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.health-service.cuttlefish
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            libbase
            libbatterymonitor
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealth_aidl_impl
            libhealthloop
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.health-service.cuttlefish ####################

#################### android.hardware.health-service.cuttlefish_recovery ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/health/health-aidl.cpp
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
    add_executable(android.hardware.health-service.cuttlefish_recovery ${SOURCE_FILES})

    target_include_directories(android.hardware.health-service.cuttlefish_recovery
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/health
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-service.cuttlefish_recovery PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.health-service.cuttlefish_recovery
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            libbase
            libbatterymonitor
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealth_aidl_impl
            libhealthloop
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.health-service.cuttlefish_recovery ####################

#################### android.hardware.health-2.1-impl-cuttlefish ####################
set(FILES
        ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/health/health-hidl.cpp
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
    add_library(android.hardware.health-2.1-impl-cuttlefish SHARED ${SOURCE_FILES})

    target_include_directories(android.hardware.health-2.1-impl-cuttlefish
            PUBLIC
            ${ANDROID_ROOT}/device/google/cuttlefish/guest/hals/health
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-2.1-impl-cuttlefish PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.health-2.1-impl-cuttlefish
            android.hardware.health-1.0-convert
            android.hardware.health-2.0
            android.hardware.health-2.1
            libbase
            libbatterymonitor
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealth2impl
            libhealthloop
            libhidlbase
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.health-2.1-impl-cuttlefish ####################

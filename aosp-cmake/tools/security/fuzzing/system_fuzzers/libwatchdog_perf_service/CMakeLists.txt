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

project(libwatchdog_perf_service_fuzzer)

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

#################### libwatchdog_perf_service_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/tools/security/fuzzing/system_fuzzers/libwatchdog_perf_service/libwatchdog_perf_service_fuzzer.cpp
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
    add_executable(libwatchdog_perf_service_fuzzer ${SOURCE_FILES})

    target_include_directories(libwatchdog_perf_service_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/tools/security/fuzzing/system_fuzzers/libwatchdog_perf_service
    )

    # Add c/c++ flags
    target_compile_options(libwatchdog_perf_service_fuzzer PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wno-unused-variable
            -Wall
            -Werror
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libwatchdog_perf_service_fuzzer
            VehicleHalUtils
            android-automotive-large-parcelable-lib
            android.automotive.watchdog-V3-ndk
            android.automotive.watchdog.internal-V3-ndk
            android.hardware.automotive.vehicle-V2-ndk
            android.hardware.automotive.vehicle.property-V2-ndk
            android.hardware.automotive.vehicle-2.0
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libgtest
            libgtest_prod_headers
            libhidlbase
            liblog
            libm
            libmath
            libprocessgroup
            libtinyxml2
            libutils
            libvhalclient
            libwatchdog_perf_service
            libwatchdog_properties
            packagemanager_aidl-cpp
    )
endif ()
#################### libwatchdog_perf_service_fuzzer ####################

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

project(android.hardware.health-service.example_recovery)

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

#################### fuzz_libhealth_aidl_impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/ChargerUtils.cpp
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/health-convert.cpp
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/HalHealthLoop.cpp
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/Health.cpp
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/LinkedCallback.cpp
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
    add_library(fuzz_libhealth_aidl_impl SHARED ${SOURCE_FILES})

    target_include_directories(fuzz_libhealth_aidl_impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/include
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(fuzz_libhealth_aidl_impl PRIVATE
            -Werror
    )

    target_link_libraries(fuzz_libhealth_aidl_impl
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            libbase
            libbatterymonitor
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealthd_charger_ui
            libhealthd_draw
            libhealthloop
            libhidlbase
            liblog
            libm
            libminui
            libpng
            libsuspend
            libutils
    )
endif ()
#################### fuzz_libhealth_aidl_impl ####################

#################### libhealth_aidl_impl ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/health-convert.cpp
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/HalHealthLoop.cpp
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/Health.cpp
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/LinkedCallback.cpp
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
    add_library(libhealth_aidl_impl SHARED ${SOURCE_FILES})

    target_include_directories(libhealth_aidl_impl
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/include
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(libhealth_aidl_impl PRIVATE
            -Werror
    )

    target_link_libraries(libhealth_aidl_impl
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            libbase
            libbatterymonitor
            libbinder_ndk
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhealthd_charger_ui
            libhealthd_draw
            libhealthloop
            libhidlbase
            liblog
            libm
            libminui
            libpng
            libsuspend
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libhealth_aidl_impl ####################

#################### android.hardware.health-service.aidl_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/fuzzer.cpp
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
    add_executable(android.hardware.health-service.aidl_fuzzer ${SOURCE_FILES})

    target_include_directories(android.hardware.health-service.aidl_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-service.aidl_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.health-service.aidl_fuzzer
            android.hardware.health-V2-ndk
            android.hardware.health-translate-ndk
            fuzz_libhealth_aidl_impl
            libbase
            libbatterymonitor
            libbinder
            libbinder_ndk
            libbinder_random_parcel
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libhealthd_charger_ui
            libhealthd_draw
            libhealthloop
            libhidlbase
            liblog
            libm
            libminui
            libpng
            libsuspend
            libutils
    )
endif ()
#################### android.hardware.health-service.aidl_fuzzer ####################

#################### android.hardware.health-service.example ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/main.cpp
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
    add_executable(android.hardware.health-service.example ${SOURCE_FILES})

    target_include_directories(android.hardware.health-service.example
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-service.example PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.health-service.example
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
            libhealthd_charger_ui
            libhealthd_draw
            libhealthloop
            libhidlbase
            liblog
            libm
            libminui
            libpng
            libsuspend
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.health-service.example ####################

#################### android.hardware.health-service.example_recovery ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/main.cpp
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
    add_executable(android.hardware.health-service.example_recovery ${SOURCE_FILES})

    target_include_directories(android.hardware.health-service.example_recovery
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default
            ${ANDROID_ROOT}/hardware/interfaces/health/aidl/default/include
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.health-service.example_recovery PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.health-service.example_recovery
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
#################### android.hardware.health-service.example_recovery ####################

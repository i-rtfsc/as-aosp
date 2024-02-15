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

project(android.hardware.tv.hdmi.cec-service)

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

#################### android.hardware.tv.hdmi.cec-service_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/tv/hdmi/cec/aidl/default/fuzzer.cpp
        ${ANDROID_ROOT}/hardware/interfaces/tv/hdmi/cec/aidl/default/HdmiCecMock.cpp
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
    add_executable(android.hardware.tv.hdmi.cec-service_fuzzer ${SOURCE_FILES})

    target_include_directories(android.hardware.tv.hdmi.cec-service_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/tv/hdmi/cec/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.tv.hdmi.cec-service_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(android.hardware.tv.hdmi.cec-service_fuzzer
            android.hardware.tv.hdmi.cec-V1-ndk
            libbase
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
            liblog
            libm
            libutils
    )
endif ()
#################### android.hardware.tv.hdmi.cec-service_fuzzer ####################

#################### android.hardware.tv.hdmi.cec-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/interfaces/tv/hdmi/cec/aidl/default/serviceMock.cpp
        ${ANDROID_ROOT}/hardware/interfaces/tv/hdmi/cec/aidl/default/HdmiCecMock.cpp
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
    add_executable(android.hardware.tv.hdmi.cec-service ${SOURCE_FILES})

    target_include_directories(android.hardware.tv.hdmi.cec-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/interfaces/tv/hdmi/cec/aidl/default
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.tv.hdmi.cec-service PRIVATE
            -Werror
            -Wall
            -Wextra
    )

    target_link_libraries(android.hardware.tv.hdmi.cec-service
            android.hardware.tv.hdmi.cec-V1-ndk
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhardware
            libhidlbase
            liblog
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.tv.hdmi.cec-service ####################

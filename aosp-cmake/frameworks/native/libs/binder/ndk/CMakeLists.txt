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

project(libbinder_ndk)

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

#################### libbinder_ndk ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/ibinder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/ibinder_jni.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/libbinder.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/parcel.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/parcel_jni.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/process.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/stability.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/status.cpp
        ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/service_manager.cpp
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
    add_library(libbinder_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libbinder_ndk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_cpp
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_ndk
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk/include_platform
            ${ANDROID_ROOT}/frameworks/native/libs/binder/ndk
    )

    # Add c/c++ flags
    target_compile_options(libbinder_ndk PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libbinder_ndk
            jni_headers
            libandroid_runtime_lazy
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libbinder_ndk ####################

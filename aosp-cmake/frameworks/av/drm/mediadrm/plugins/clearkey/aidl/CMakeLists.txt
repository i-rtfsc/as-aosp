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

project(android.hardware.drm-service-lazy.clearkey)

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

#################### android.hardware.drm-service.clearkey.aidl_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/CreatePluginFactories.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/CryptoPlugin.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/DrmFactory.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/DrmPlugin.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/fuzzer.cpp
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
    add_executable(android.hardware.drm-service.clearkey.aidl_fuzzer ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-service.clearkey.aidl_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/include
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-service.clearkey.aidl_fuzzer PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.drm-service.clearkey.aidl_fuzzer
            android.hardware.common-V2-ndk
            android.hardware.drm-V1-ndk
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
            libclearkeybase_fuzz
            libcrypto
            libcutils
            libdl
            libjsoncpp
            liblog
            libm
            libprotobuf-cpp-lite
            libutils
    )
endif ()
#################### android.hardware.drm-service.clearkey.aidl_fuzzer ####################

#################### android.hardware.drm-service.clearkey ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/CreatePluginFactories.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/CryptoPlugin.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/DrmFactory.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/DrmPlugin.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/Service.cpp
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
    add_executable(android.hardware.drm-service.clearkey ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-service.clearkey
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/include
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-service.clearkey PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.drm-service.clearkey
            android.hardware.common-V2-ndk
            android.hardware.drm-V1-ndk
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libclearkeybase
            libcrypto
            libdl
            libjsoncpp
            liblog
            libm
            libprotobuf-cpp-lite
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.drm-service.clearkey ####################

#################### android.hardware.drm-service-lazy.clearkey ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/CreatePluginFactories.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/CryptoPlugin.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/DrmFactory.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/DrmPlugin.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/ServiceLazy.cpp
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
    add_executable(android.hardware.drm-service-lazy.clearkey ${SOURCE_FILES})

    target_include_directories(android.hardware.drm-service-lazy.clearkey
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl/include
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/aidl
    )

    # Add c/c++ flags
    target_compile_options(android.hardware.drm-service-lazy.clearkey PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(android.hardware.drm-service-lazy.clearkey
            android.hardware.common-V2-ndk
            android.hardware.drm-V1-ndk
            libbase
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libclearkeybase
            libcrypto
            libdl
            libjsoncpp
            liblog
            libm
            libprotobuf-cpp-lite
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### android.hardware.drm-service-lazy.clearkey ####################

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

project(libstagefright_mpeg2support)

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

#################### libstagefright_mpeg2support ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/AnotherPacketSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/ATSParser.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/CasManager.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/ESQueue.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/HlsSampleDecryptor.cpp
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
    add_library(libstagefright_mpeg2support SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_mpeg2support
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/include
            ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/include/mpeg2ts
            ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_mpeg2support PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_mpeg2support
            android.hardware.cas.native-1.0
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libaudioclient_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcrypto
            libdl
            libhidlmemory
            libm
            libmedia_datasource_headers
            libstagefright_foundation_headers
            libstagefright_headers
            libstagefright_metadatautils
            media_ndk_headers
    )
endif ()
#################### libstagefright_mpeg2support ####################

#################### libstagefright_mpeg2support_nocrypto ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/AnotherPacketSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/ATSParser.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/CasManager.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/ESQueue.cpp
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
    add_library(libstagefright_mpeg2support_nocrypto SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_mpeg2support_nocrypto
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/include
            ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts/include/mpeg2ts
            ${ANDROID_ROOT}/frameworks/av/media/module/mpeg2ts
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_mpeg2support_nocrypto PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_mpeg2support_nocrypto
            android.hardware.cas.native-1.0
            android.hidl.allocator-1.0
            android.hidl.memory-1.0
            libaudioclient_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libhidlmemory
            libm
            libmedia_datasource_headers
            libstagefright_foundation_headers
            libstagefright_headers
            libstagefright_metadatautils
            media_ndk_headers
    )
endif ()
#################### libstagefright_mpeg2support_nocrypto ####################

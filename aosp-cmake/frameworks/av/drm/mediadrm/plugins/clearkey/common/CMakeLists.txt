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

project(libclearkeycommon)

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

#################### libclearkeycommon ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/ClearKeyUUID.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/Utils.cpp
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
    add_library(libclearkeycommon SHARED ${SOURCE_FILES})

    target_include_directories(libclearkeycommon
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/include
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common
    )

    # Add c/c++ flags
    target_compile_options(libclearkeycommon PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libclearkeycommon
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libclearkeycommon ####################

#################### libclearkeybase_fuzz ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/AesCtrDecryptor.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/Base64.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/Buffer.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/ClearKeyUUID.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/DeviceFiles.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/InitDataParser.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/JsonWebKey.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/MemoryFileSystem.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/Session.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/SessionLibrary.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/Utils.cpp
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
    add_library(libclearkeybase_fuzz SHARED ${SOURCE_FILES})

    target_include_directories(libclearkeybase_fuzz
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/include
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/include/clearkeydrm
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common_fuzz/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/av/drm/mediadrm/plugins/clearkey/common
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common_fuzz/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libclearkeybase_fuzz PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libclearkeybase_fuzz
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libclearkeydevicefiles-protos.common_fuzz
            libcrypto
            libdl
            libjsmn
            libm
            libutils
    )
endif ()
#################### libclearkeybase_fuzz ####################

#################### libclearkeybase ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/AesCtrDecryptor.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/Base64.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/Buffer.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/ClearKeyUUID.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/DeviceFiles.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/InitDataParser.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/JsonWebKey.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/MemoryFileSystem.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/Session.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/SessionLibrary.cpp
        ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/Utils.cpp
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
    add_library(libclearkeybase SHARED ${SOURCE_FILES})

    target_include_directories(libclearkeybase
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/include
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/include/clearkeydrm
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/frameworks/av/drm/mediadrm/plugins/clearkey/common
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libclearkeybase PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libclearkeybase
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libclearkeydevicefiles-protos.common
            libcrypto
            libdl
            libjsmn
            libm
            libutils
            libc
            libdl
            libm
    )
endif ()
#################### libclearkeybase ####################

#################### libclearkeydevicefiles-protos.common ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/frameworks/av/drm/mediadrm/plugins/clearkey/common/protos/DeviceFiles.pb.cc
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
    add_library(libclearkeydevicefiles-protos.common SHARED ${SOURCE_FILES})

    target_include_directories(libclearkeydevicefiles-protos.common
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto/frameworks/av/drm/mediadrm/plugins/clearkey/common
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common/android_vendor.34_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libclearkeydevicefiles-protos.common PRIVATE
            -Werror
    )

    target_link_libraries(libclearkeydevicefiles-protos.common
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
            libc
            libdl
            libm
    )
endif ()
#################### libclearkeydevicefiles-protos.common ####################

#################### libclearkeydevicefiles-protos.common_fuzz ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common_fuzz/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/av/drm/mediadrm/plugins/clearkey/common/protos/DeviceFiles.pb.cc
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
    add_library(libclearkeydevicefiles-protos.common_fuzz SHARED ${SOURCE_FILES})

    target_include_directories(libclearkeydevicefiles-protos.common_fuzz
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common_fuzz/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/av/drm/mediadrm/plugins/clearkey/common
            ${OUT_INTERMEDIATES_ROOT}/frameworks/av/drm/mediadrm/plugins/clearkey/common/libclearkeydevicefiles-protos.common_fuzz/android_${OUT_ARCH_CPU}_static/gen/proto
    )

    # Add c/c++ flags
    target_compile_options(libclearkeydevicefiles-protos.common_fuzz PRIVATE
            -Werror
    )

    target_link_libraries(libclearkeydevicefiles-protos.common_fuzz
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libprotobuf-cpp-lite
    )
endif ()
#################### libclearkeydevicefiles-protos.common_fuzz ####################

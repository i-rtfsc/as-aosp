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

project(libandroidfw_fuzzer_lib)

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

#################### libandroidfw_fuzzer_lib ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/CursorWindow.cpp
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
    add_library(libandroidfw_fuzzer_lib SHARED ${SOURCE_FILES})

    target_include_directories(libandroidfw_fuzzer_lib
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/libs/androidfw/include
            ${ANDROID_ROOT}/frameworks/base/libs/androidfw
    )

    # Add c/c++ flags
    target_compile_options(libandroidfw_fuzzer_lib PRIVATE
            -Werror
    )

    target_link_libraries(libandroidfw_fuzzer_lib
            libandroidfw
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libutils
            libziparchive
    )
endif ()
#################### libandroidfw_fuzzer_lib ####################

#################### libandroidfw ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ApkAssets.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ApkParsing.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/Asset.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/AssetDir.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/AssetManager.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/AssetManager2.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/AssetsProvider.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/AttributeResolution.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/BigBuffer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ChunkIterator.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ConfigDescription.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/Idmap.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/LoadedArsc.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/Locale.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/LocaleData.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/misc.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ObbFile.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/PosixUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ResourceTimer.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ResourceTypes.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ResourceUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/StreamingZipInflater.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/StringPool.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/TypeWrappers.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/Util.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ZipFileRO.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/ZipUtils.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/BackupData.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/BackupHelpers.cpp
        ${ANDROID_ROOT}/frameworks/base/libs/androidfw/CursorWindow.cpp
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
    add_library(libandroidfw SHARED ${SOURCE_FILES})

    target_include_directories(libandroidfw
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/libs/androidfw/include
            ${ANDROID_ROOT}/frameworks/base/libs/androidfw
    )

    # Add c/c++ flags
    target_compile_options(libandroidfw PRIVATE
            -Werror
    )

    target_link_libraries(libandroidfw
            libbase
            libbinder
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libincfs
            libincfs-utils
            liblog
            libm
            libutils
            libwinpthread
            libz
            libziparchive
            libziparchive_for_incfs
    )
endif ()
#################### libandroidfw ####################

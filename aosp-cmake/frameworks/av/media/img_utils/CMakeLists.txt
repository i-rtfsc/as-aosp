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

project(libimg_utils)

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

#################### libimg_utils ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/EndianUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/FileInput.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/FileOutput.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/SortedEntryVector.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/Input.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/Output.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/Orderable.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/TiffIfd.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/TiffWritable.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/TiffWriter.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/TiffEntry.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/TiffEntryImpl.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/ByteArrayOutput.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/DngUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/img_utils/src/StripSource.cpp
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
    add_library(libimg_utils SHARED ${SOURCE_FILES})

    target_include_directories(libimg_utils
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/img_utils/include
            ${ANDROID_ROOT}/frameworks/av/media/img_utils
    )

    # Add c/c++ flags
    target_compile_options(libimg_utils PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libimg_utils
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### libimg_utils ####################

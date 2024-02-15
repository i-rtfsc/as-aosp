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

project(libkmsxx)

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

#################### libkmsxx ####################
set(FILES
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/atomicreq.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/blob.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/card.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/connector.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/crtc.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/dmabufframebuffer.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/drmobject.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/drmpropobject.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/dumbframebuffer.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/encoder.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/extframebuffer.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/framebuffer.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/helpers.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/mode_cvt.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/modedb.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/modedb_cea.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/modedb_dmt.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/pixelformats.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/plane.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/property.cpp
        ${ANDROID_ROOT}/external/libkmsxx/kms++/src/videomode.cpp
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
    add_library(libkmsxx SHARED ${SOURCE_FILES})

    target_include_directories(libkmsxx
            PUBLIC
            ${ANDROID_ROOT}/external/libkmsxx/kms++/inc
            ${ANDROID_ROOT}/external/libkmsxx
    )

    # Add c/c++ flags
    target_compile_options(libkmsxx PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
    )

    target_link_libraries(libkmsxx
            libbase
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libdrm
            libm
    )
endif ()
#################### libkmsxx ####################

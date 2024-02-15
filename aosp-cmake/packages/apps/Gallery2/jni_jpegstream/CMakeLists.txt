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

project(libjni_jpegstream)

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

#################### libjni_jpegstream ####################
set(FILES
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream/src/inputstream_wrapper.cpp
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream/src/jpegstream.cpp
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream/src/jerr_hook.cpp
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream/src/jpeg_hook.cpp
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream/src/jpeg_writer.cpp
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream/src/jpeg_reader.cpp
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream/src/outputstream_wrapper.cpp
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream/src/stream_wrapper.cpp
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
    add_library(libjni_jpegstream SHARED ${SOURCE_FILES})

    target_include_directories(libjni_jpegstream
            PUBLIC
            ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream/src
            ${ANDROID_ROOT}/packages/apps/Gallery2/jni_jpegstream
    )

    # Add c/c++ flags
    target_compile_options(libjni_jpegstream PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libjni_jpegstream
            jni_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libjpeg_static_ndk
            liblog
            libm
    )
endif ()
#################### libjni_jpegstream ####################

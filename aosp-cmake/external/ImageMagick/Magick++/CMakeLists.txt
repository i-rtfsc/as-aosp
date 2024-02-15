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

project(Magick++)

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

#################### Magick++ ####################
set(FILES
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Blob.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/BlobRef.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/CoderInfo.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Color.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Drawable.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Exception.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Functions.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Geometry.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Image.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/ImageRef.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Montage.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Options.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Pixels.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/ResourceLimits.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/STL.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/SecurityPolicy.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Statistic.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/Thread.cpp
        ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib/TypeMetric.cpp
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
    add_library(Magick++ SHARED ${SOURCE_FILES})

    target_include_directories(Magick++
            PUBLIC
            ${ANDROID_ROOT}/external/ImageMagick/Magick++/lib
            ${ANDROID_ROOT}/external/ImageMagick/Magick++
    )

    # Add c/c++ flags
    target_compile_options(Magick++ PRIVATE
            -Wall
            -Werror
            -Wno-unused-private-field
    )

    target_link_libraries(Magick++
            Magick_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### Magick++ ####################

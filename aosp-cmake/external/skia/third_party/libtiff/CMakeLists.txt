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

project(libtiff)

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

#################### libtiff ####################
set(FILES
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_acorn.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_aux.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_close.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_codec.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_color.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_compress.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_dir.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_dirinfo.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_dirread.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_dirwrite.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_dumpmode.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_error.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_extension.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_fax3.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_fax3sm.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_flush.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_getimage.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_jbig.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_jpeg.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_luv.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_lzw.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_next.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_ojpeg.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_open.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_packbits.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_pixarlog.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_predict.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_print.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_read.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_strip.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_swab.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_thunder.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_tile.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_version.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_warning.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_write.c
        ${ANDROID_ROOT}/external/skia/third_party/libtiff/tif_zip.c
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
    add_library(libtiff SHARED ${SOURCE_FILES})

    target_include_directories(libtiff
            PUBLIC
            ${ANDROID_ROOT}/external/skia/third_party/libtiff
    )

    # Add c/c++ flags
    target_compile_options(libtiff PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
            -Wno-unused-variable
    )

    target_link_libraries(libtiff
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libjpeg
            libm
            libutils
            libz
    )
endif ()
#################### libtiff ####################

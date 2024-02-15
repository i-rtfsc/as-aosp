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

project(MagickCore)

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

#################### MagickCore ####################
set(FILES
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/accelerate.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/animate.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/annotate.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/artifact.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/attribute.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/blob.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/cache-view.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/cache.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/channel.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/cipher.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/client.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/coder.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/color.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/colormap.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/colorspace.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/compare.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/composite.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/compress.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/configure.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/constitute.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/decorate.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/delegate.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/deprecate.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/display.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/distort.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/distribute-cache.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/draw.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/effect.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/enhance.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/exception.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/feature.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/fourier.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/fx.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/gem.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/geometry.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/histogram.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/identify.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/image-view.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/image.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/layer.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/linked-list.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/list.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/locale.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/log.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/magic.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/magick.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/matrix.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/memory.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/mime.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/module.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/monitor.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/montage.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/morphology.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/nt-base.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/nt-feature.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/opencl.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/option.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/paint.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/pixel.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/policy.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/prepress.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/profile.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/property.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/quantize.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/quantum-export.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/quantum-import.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/quantum.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/random.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/registry.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/resample.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/resize.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/resource.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/segment.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/semaphore.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/shear.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/signature.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/splay-tree.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/static.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/statistic.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/stream.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/string.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/thread.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/threshold.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/timer.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/token.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/transform.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/type.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/utility.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/version.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/vision.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/visual-effects.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/vms.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/widget.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/xml-tree.c
        ${ANDROID_ROOT}/external/ImageMagick/MagickCore/xwindow.c
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
    add_library(MagickCore SHARED ${SOURCE_FILES})

    target_include_directories(MagickCore
            PUBLIC
            ${ANDROID_ROOT}/external/ImageMagick/MagickCore
    )

    # Add c/c++ flags
    target_compile_options(MagickCore PRIVATE
            -Wall
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(MagickCore
            Magick_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### MagickCore ####################

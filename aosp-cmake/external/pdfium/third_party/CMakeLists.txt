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

project(libpdfium-skia_shared)

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

#################### libpdfium-skia_shared ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/third_party/skia_shared/SkFloatToDecimal.cpp
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
    add_library(libpdfium-skia_shared SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-skia_shared
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/third_party
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-skia_shared PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-skia_shared
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-skia_shared ####################

#################### libpdfium-agg ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/third_party/agg23/agg_curves.cpp
        ${ANDROID_ROOT}/external/pdfium/third_party/agg23/agg_path_storage.cpp
        ${ANDROID_ROOT}/external/pdfium/third_party/agg23/agg_rasterizer_scanline_aa.cpp
        ${ANDROID_ROOT}/external/pdfium/third_party/agg23/agg_vcgen_dash.cpp
        ${ANDROID_ROOT}/external/pdfium/third_party/agg23/agg_vcgen_stroke.cpp
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
    add_library(libpdfium-agg SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-agg
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/third_party
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-agg PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(libpdfium-agg
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-agg ####################

#################### libpdfium-libopenjpeg2 ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/bio.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/cio.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/dwt.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/event.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/function_list.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/image.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/invert.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/j2k.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/jp2.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/mct.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/mqc.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/openjpeg.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/opj_clock.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/pi.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/sparse_array.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/t1.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/t2.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/tcd.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/tgt.c
        ${ANDROID_ROOT}/external/pdfium/third_party/libopenjpeg20/thread.c
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
    add_library(libpdfium-libopenjpeg2 SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-libopenjpeg2
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/third_party
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-libopenjpeg2 PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-libopenjpeg2
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-libopenjpeg2 ####################

#################### libpdfium-lcms2 ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsalpha.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmscam02.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmscgats.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmscnvrt.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmserr.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsgamma.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsgmt.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmshalf.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsintrp.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsio0.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsio1.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmslut.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsmd5.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsmtrx.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsnamed.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsopt.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmspack.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmspcs.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsplugin.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsps2.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmssamp.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmssm.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmstypes.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsvirt.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmswtpnt.c
        ${ANDROID_ROOT}/external/pdfium/third_party/lcms/src/cmsxform.c
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
    add_library(libpdfium-lcms2 SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-lcms2
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/third_party
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-lcms2 PRIVATE
            -Werror
            -Wno-unused-parameter
            -Wno-unused-function
    )

    target_link_libraries(libpdfium-lcms2
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-lcms2 ####################

#################### libpdfium-third_party-base ####################
set(FILES
        ${ANDROID_ROOT}/external/pdfium/third_party/base/debug/alias.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/address_space_randomization.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/oom_callback.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/page_allocator.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/partition_alloc.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/partition_bucket.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/partition_oom.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/partition_page.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/partition_root_base.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/random.cc
        ${ANDROID_ROOT}/external/pdfium/third_party/base/allocator/partition_allocator/spin_lock.cc
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
    add_library(libpdfium-third_party-base SHARED ${SOURCE_FILES})

    target_include_directories(libpdfium-third_party-base
            PUBLIC
            ${ANDROID_ROOT}/external/pdfium/third_party
    )

    # Add c/c++ flags
    target_compile_options(libpdfium-third_party-base PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpdfium-third_party-base
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            pdfium-headers
            pdfium-third-party-headers
    )
endif ()
#################### libpdfium-third_party-base ####################

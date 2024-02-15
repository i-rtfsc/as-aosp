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

project(xnnpack_neonfp16arith_prod_microkernels)

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

#################### xnnpack_neonfp16arith_prod_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-avgpool/9p8x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-avgpool/9x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up8x25-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x3-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x4-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x9-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool-cw/neonfp16arith-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7p7x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/1x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/6x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-ibilinear-chw/gen/neonfp16arith-p8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-ibilinear/gen/neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/1x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/6x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-maxpool/9p8x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-pavgpool/9p8x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-pavgpool/9x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-prelu/gen/neonfp16arith-2x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x40.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-rmax/neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-spmm/gen/32x1-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vadd-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vaddc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmaxc-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmin-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vminc-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmul-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmulc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vrsubc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsqrdiff-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsqrdiffc-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsub-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsubc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vclamp/gen/vclamp-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-velu/gen/velu-neonfp16arith-rr1-p3-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vhswish/gen/vhswish-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vlrelu/gen/vlrelu-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vmulcaddc/gen/c8-minmax-neonfp16arith-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndd-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndne-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndu-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndz-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1fma-x40.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vunary/gen/vabs-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vunary/gen/vneg-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vunary/gen/vsqr-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-conv-hwc2chw/3x3s2p1c3x4-neonfp16arith-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-2x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3s2p1-minmax-neonfp16arith-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vdiv-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vdivc-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vrdivc-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsqrt/gen/neonfp16arith-sqrt-x8.c
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
    add_library(xnnpack_neonfp16arith_prod_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neonfp16arith_prod_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neonfp16arith_prod_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neonfp16arith_prod_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neonfp16arith_prod_microkernels ####################

#################### xnnpack_neonfma_prod_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x3-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x9-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x25-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/neonfma-p8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/neonfma-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x1-minmax-neonfma-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-lut16-p3-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-p6-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c4-minmax-neonfma-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2recps-x16.c
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
    add_library(xnnpack_neonfma_prod_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neonfma_prod_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neonfma_prod_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neonfma_prod_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neonfma_prod_microkernels ####################

#################### xnnpack_params ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/params.c
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
    add_library(xnnpack_params SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_params
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_params PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_params
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
    )
endif ()
#################### xnnpack_params ####################

#################### xnnpack_qs8_mobilenet_v1 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/qs8-mobilenet-v1.cc
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
    add_library(xnnpack_qs8_mobilenet_v1 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_qs8_mobilenet_v1
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_qs8_mobilenet_v1 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_qs8_mobilenet_v1
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_qs8_mobilenet_v1 ####################

#################### xnnpack_mobilenet_v3_small_fp32 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp32-mobilenet-v3-small.cc
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
    add_library(xnnpack_mobilenet_v3_small_fp32 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v3_small_fp32
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v3_small_fp32 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v3_small_fp32
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v3_small_fp32 ####################

#################### xnnpack_mutex ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/mutex.c
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
    add_library(xnnpack_mutex SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mutex
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mutex PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_mutex
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            libm
            libpthreadpool
            xnnpack_logging
    )
endif ()
#################### xnnpack_mutex ####################

#################### xnnpack_neon_prod_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neon-int16-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/4x-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/9p8x-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/9x-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-avgpool/9p8x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-avgpool/9x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc2chw/3x3s2p1c3x4-neon-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x3-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x9-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x25-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neon-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool-cw/neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool/7p7x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool/7x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x2-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neon-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/neon-p8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x2-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neon-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-maxpool/9p8x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-pavgpool/9p8x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-pavgpool/9x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-2x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-rmax/neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x1-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmaxc-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmin-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vminc-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiff-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiffc-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vclamp/gen/vclamp-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-lut16-p3-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vhswish/gen/vhswish-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vlrelu/gen/vlrelu-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c4-minmax-neon-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut64-p2-nr2recps-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vabs-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vneg-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vsqr-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x25-minmax-fp32-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x3-minmax-fp32-neon-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x25-minmax-rndnu-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x9-minmax-rndnu-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x25-minmax-rndnu-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-rndnu-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-rndnu-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-neon-ld64-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-neon-ld64-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-rndnu-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-rndnu-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-avgpool/9p8x-minmax-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-avgpool/9x-minmax-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up8x25-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up16x9-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-rndnu-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-rndnu-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-neon-ld64-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-neon-ld64-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-rndnu-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-rndnu-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-ibilinear/gen/neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-ibilinear/gen/neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-maxpool/9p8x-minmax-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-vclamp/neon-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-ibilinear/gen/neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-ibilinear/gen/neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-maxpool/9p8x-minmax-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-rmax/neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-vclamp/neon-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-fill/neon-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-pad/neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/16x16-reuse-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/xm-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x2-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x3-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x4-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/8x8-reuse-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-packx/x4-neon-st4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-reuse-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-unpool/neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/xm-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x2-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x3-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x4-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc2chw/3x3s2p1c3x4-neonfma-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neonfma-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x2-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x2-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x2-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neon-sqrt-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-neon-tbx128x4-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/4x4-aarch64-tbl.c
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
    add_library(xnnpack_neon_prod_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neon_prod_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neon_prod_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neon_prod_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neon_prod_microkernels ####################

#################### xnnpack_mobilenet_v1_fp32_sparse ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp32-sparse-mobilenet-v1.cc
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
    add_library(xnnpack_mobilenet_v1_fp32_sparse SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v1_fp32_sparse
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v1_fp32_sparse PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v1_fp32_sparse
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v1_fp32_sparse ####################

#################### xnnpack_allocator ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/allocator.c
        ${ANDROID_ROOT}/external/XNNPACK/src/memory.c
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
    add_library(xnnpack_allocator SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_allocator
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_allocator PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_allocator
            fp16_headers
            fxdiv_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            libm
            libpthreadpool
    )
endif ()
#################### xnnpack_allocator ####################

#################### xnnpack_qs8_mobilenet_v2 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/qs8-mobilenet-v2.cc
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
    add_library(xnnpack_qs8_mobilenet_v2 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_qs8_mobilenet_v2
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_qs8_mobilenet_v2 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_qs8_mobilenet_v2
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_qs8_mobilenet_v2 ####################

#################### xnnpack_neonv8_prod_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neonv8-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neonv8-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x25-minmax-fp32-neonv8-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x3-minmax-fp32-neonv8-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neonv8-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neonv8-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x16-minmax-fp32-neonv8-mlal-lane.c
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
    add_library(xnnpack_neonv8_prod_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neonv8_prod_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neonv8_prod_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neonv8_prod_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neonv8_prod_microkernels ####################

#################### xnnpack_tables ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/tables/exp2-k-over-64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/tables/exp2-k-over-2048.c
        ${ANDROID_ROOT}/external/XNNPACK/src/tables/exp2minus-k-over-4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/tables/exp2minus-k-over-8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/tables/exp2minus-k-over-16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/tables/exp2minus-k-over-64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/tables/exp2minus-k-over-2048.c
        ${ANDROID_ROOT}/external/XNNPACK/src/tables/vlog.c
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
    add_library(xnnpack_tables SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_tables
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_tables PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_tables
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### xnnpack_tables ####################

#################### xnnpack_packing ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/packing.c
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
    add_library(xnnpack_packing SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_packing
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_packing PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_packing
            fp16_headers
            fxdiv_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
    )
endif ()
#################### xnnpack_packing ####################

#################### xnnpack_operators ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/operator-delete.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operator-run.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operator-utils.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/argmax-pooling-nhwc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/average-pooling-nhwc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/binary-elementwise-nd.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/channel-shuffle-nc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/constant-pad-nd.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/convolution-nchw.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/convolution-nhwc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/deconvolution-nhwc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/fully-connected-nc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/global-average-pooling-ncw.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/global-average-pooling-nwc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/lut-elementwise-nc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/max-pooling-nhwc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/prelu-nc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/resize-bilinear-nchw.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/resize-bilinear-nhwc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/softmax-nc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/transpose-nd.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/unary-elementwise-nc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/unpooling-nhwc.c
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
    add_library(xnnpack_operators SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_operators
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_operators PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_operators
            fp16_headers
            fxdiv_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_indirection
            xnnpack_logging
            xnnpack_microparams_init
            xnnpack_normalization
            xnnpack_packing
            xnnpack_post_operation
    )
endif ()
#################### xnnpack_operators ####################

#################### xnnpack_scalar_bench_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-bfly4/gen/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-bfly4/gen/scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-bfly4/gen/scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-bfly4/gen/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-bfly4/samples1-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-fftr/gen/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-fftr/gen/scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-fftr/gen/scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-fftr/gen/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-vsquareabs/gen/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-vsquareabs/gen/scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-vsquareabs/gen/scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-vsquareabs/gen/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/4x-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/9p8x-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/9x-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-avgpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-avgpool/9x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/3x3s2p0p1c3x4-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/3x3s2p1c3x4-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc2chw/3x3s2p1c3x4-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x3-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x3-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x3-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x3-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x4-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x4-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x9-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x9-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x9-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x9-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x25-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x25-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x25-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x25-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x3-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x3-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x3-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x3-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x4-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x4-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x9-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x9-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x9-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x9-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x25-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x25-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x25-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up2x25-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-1x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-1x1-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-1x1-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-2x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-3x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-4x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-5x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-6x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-scalar-1x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-scalar-1x1-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-scalar-1x1-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-scalar-2x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-scalar-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-scalar-3x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-scalar-4x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-1x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-1x1-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-1x1-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-1x1-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-2x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-2x1-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-3x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-3x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-1x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-1x1-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-1x1-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-1x1-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-2x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-2x1-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-3x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-3x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-scalar-bitcast-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-scalar-bitcast-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-scalar-bitcast-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-scalar-bitcast-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-scalar-fabsf-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-scalar-fabsf-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-scalar-fabsf-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-scalar-fabsf-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool-cw/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool/7p7x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool/7x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/1x4inc-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/2x4inc-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x4inc-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/2x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/2x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/2x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x2-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x2-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/scalar-p1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/scalar-p2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/scalar-p4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/scalar-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/scalar-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/2x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/2x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/2x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x2-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x2-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-maxpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-pavgpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-pavgpool/9x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ppmm/gen/2x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ppmm/gen/3x3-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ppmm/gen/4x2-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ppmm/gen/4x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/scalar-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/scalar-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-fmagic-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-fmagic-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-fmagic-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-fmagic-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-imagic-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-imagic-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-imagic-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-imagic-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-lrintf-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-lrintf-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-lrintf-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-lrintf-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-fmagic-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-fmagic-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-fmagic-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-fmagic-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-imagic-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-imagic-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-imagic-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-imagic-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-lrintf-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-lrintf-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-lrintf-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-lrintf-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-lut64-p2-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-lut64-p2-x2-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-lut64-p2-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-lut64-p2-x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-lut64-p2-x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-lut64-p2-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-p5-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-p5-x2-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-p5-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-p5-x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-p5-x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-p5-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-rmax/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/1x1-minmax-scalar-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/1x1-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/2x1-minmax-scalar-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/2x1-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x1-minmax-scalar-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x1-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x1-minmax-scalar-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x1-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x2-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmaxc-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmaxc-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmaxc-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmaxc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmin-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmin-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmin-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmin-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vminc-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vminc-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vminc-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vminc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiff-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiff-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiff-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiff-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiffc-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiffc-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiffc-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiffc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-relu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-relu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-relu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-relu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vclamp/gen/vclamp-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vclamp/gen/vclamp-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vclamp/gen/vclamp-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-lut16-p3-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-lut16-p3-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-lut16-p3-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-lut16-p3-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-lut16-p3-x5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-lut16-p3-x6.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-p6-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-p6-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-p6-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-p6-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-p6-x5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-p6-x6.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vhswish/gen/vhswish-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vhswish/gen/vhswish-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vhswish/gen/vhswish-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vlrelu/gen/vlrelu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vlrelu/gen/vlrelu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vlrelu/gen/vlrelu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c1-minmax-scalar-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c2-minmax-scalar-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c4-minmax-scalar-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrelu/gen/vrelu-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrelu/gen/vrelu-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrelu/gen/vrelu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrelu/gen/vrelu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-scalar-libm-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-scalar-libm-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-scalar-libm-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-scalar-libm-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-scalar-libm-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-scalar-libm-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-scalar-libm-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-scalar-libm-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-scalar-libm-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-scalar-libm-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-scalar-libm-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-scalar-libm-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-lut64-p2-div-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-lut64-p2-div-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-lut64-p2-div-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-lut2048-p1-div-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-lut2048-p1-div-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-lut2048-p1-div-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-p5-div-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-p5-div-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-p5-div-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/scalar-sqrt-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/scalar-sqrt-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/scalar-sqrt-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vabs-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vabs-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vabs-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vneg-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vneg-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vneg-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vsqr-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vsqr-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vsqr-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f32-f16-scalar-bitcast.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f32-f16-scalar-fabsf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f32-scalar-rr2-lut64-p2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f32-scalar-rr2-lut2048-p1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f32-scalar-rr2-p5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-scalar-rr2-lut4-p4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-scalar-rr2-lut8-p3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-scalar-rr2-lut8-p4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-scalar-rr2-lut16-p3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-scalar-rr2-lut16-p4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-scalar-rr2-p5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-scalar-rr2-p6.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundd-scalar-addsub.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundd-scalar-cvt.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundd-scalar-floor.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundne-scalar-addsub.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundne-scalar-nearbyint.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundne-scalar-rint.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundu-scalar-addsub.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundu-scalar-ceil.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundu-scalar-cvt.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundz-scalar-addsub.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundz-scalar-cvt.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundz-scalar-trunc.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-scalar-rr2-lut64-p2-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-scalar-rr2-lut2048-p1-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-scalar-rr2-p5-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-bitmanip.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-clz-binsearch.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-clz-newton.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-cvti32-sqrt-lrint.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-cvti64-sqrt-lrint.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-cvti64-sqrtf-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-cvtu32-sqrt-lrint.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-cvtu32-sqrtf-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-hashemian.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u32-scalar-tflm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u64-scalar-cvtu32-sqrt-cvtsatu32f64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u64-scalar-cvtu32-sqrt-llrint.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-u64-scalar-cvtu64-sqrt-llrint.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x3-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x9-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x9-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x25-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x25-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up2x3-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up2x3-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up2x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up2x9-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up2x9-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up2x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up2x25-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up2x25-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up4x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up4x9-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up4x9-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up4x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up4x25-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up4x25-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up1x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up1x9-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up1x9-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up1x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up1x25-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up1x25-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up2x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up2x9-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up2x9-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up2x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up2x25-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up2x25-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up4x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up4x9-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up4x9-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up4x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up4x25-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up4x25-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-fmagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-fmagic-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-fmagic-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-imagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-imagic-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-imagic-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-lrintf-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-lrintf-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-lrintf-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-fmagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-fmagic-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-fmagic-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-imagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-imagic-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-imagic-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-lrintf-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-lrintf-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-lrintf-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/gemmlowp-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/rndna-scalar-signed64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/rndna-scalar-unsigned32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/rndna-scalar-unsigned64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/rndnu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vcvt/gen/vcvt-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vcvt/gen/vcvt-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vcvt/gen/vcvt-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-scalar-andxor-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-scalar-andxor-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-scalar-andxor-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-scalar-select-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-scalar-select-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-scalar-select-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-avgpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-avgpool/9x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up1x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up1x9-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up1x9-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up1x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up1x25-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up1x25-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up2x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up2x9-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up2x9-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up2x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up2x25-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up2x25-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up4x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up4x9-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up4x9-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up4x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up4x25-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up4x25-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-fmagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-fmagic-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-fmagic-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-imagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-imagic-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-imagic-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-lrintf-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-lrintf-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-lrintf-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-fmagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-fmagic-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-fmagic-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-imagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-imagic-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-imagic-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-lrintf-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-lrintf-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-lrintf-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x2-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x2-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x4-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x4-minmax-fp32-scalar-imagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x4-minmax-fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-requantization/fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-requantization/fp32-scalar-lrintf.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-requantization/gemmlowp-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-requantization/rndna-scalar-signed64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-requantization/rndna-scalar-unsigned32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-requantization/rndna-scalar-unsigned64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vcvt/gen/vcvt-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vcvt/gen/vcvt-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vcvt/gen/vcvt-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-scalar-andxor-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-scalar-andxor-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-scalar-andxor-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-scalar-select-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-scalar-select-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-scalar-select-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-ibilinear/gen/scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-ibilinear/gen/scalar-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-ibilinear/gen/scalar-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-maxpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-vclamp/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-rmaxabs/gen/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-rmaxabs/gen/scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-rmaxabs/gen/scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-rmaxabs/gen/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-vlshift/gen/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-vlshift/gen/scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-vlshift/gen/scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-vlshift/gen/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-ibilinear/gen/scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-ibilinear/gen/scalar-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-ibilinear/gen/scalar-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-lut32norm/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-maxpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-rmax/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-vclamp/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u32-filterbank-accumulate/gen/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u32-filterbank-subtract/scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u32-vlog/gen/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u32-vlog/gen/scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u32-vlog/gen/scalar-x3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u32-vlog/gen/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u64-u32-vsqrtshift/scalar-cvtu32-sqrt-cvtu32f64-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-copy/memcpy.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-fill/scalar-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-pad/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-transpose/1x1-memcpy.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-scalar-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/1x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/1x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/2x1-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/2x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/2x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/4x1-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/4x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/4x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/xm-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x3-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/1x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/1x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/2x1-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/2x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/2x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x1-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/gen/1x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/gen/1x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/gen/2x1-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/gen/2x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/gen/2x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/gen/4x1-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/gen/4x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/gen/4x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-packx/x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-packx/x3-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-packx/x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/1x2-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/1x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/1x4-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/1x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x1-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x1-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x4-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x1-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x1-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x2-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-unpool/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/xm-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x3-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/1x2-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/1x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/2x1-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/2x1-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/2x2-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/2x2-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/4x1-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/4x1-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/4x2-scalar-float.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x64-transposec/gen/4x2-scalar-int.c
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
    add_library(xnnpack_scalar_bench_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_scalar_bench_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_scalar_bench_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_scalar_bench_microkernels
            fp16_headers
            fxdiv_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_scalar_bench_microkernels ####################

#################### xnnpack_mobilenet_v3_small_fp32_sparse ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp32-sparse-mobilenet-v3-small.cc
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
    add_library(xnnpack_mobilenet_v3_small_fp32_sparse SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v3_small_fp32_sparse
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v3_small_fp32_sparse PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v3_small_fp32_sparse
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v3_small_fp32_sparse ####################

#################### xnnpack_mobilenet_v3_small_fp16 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp16-mobilenet-v3-small.cc
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
    add_library(xnnpack_mobilenet_v3_small_fp16 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v3_small_fp16
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v3_small_fp16 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v3_small_fp16
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v3_small_fp16 ####################

#################### xnnpack_normalization ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/normalization.c
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
    add_library(xnnpack_normalization SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_normalization
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_normalization PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_normalization
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            libm
            libpthreadpool
            xnnpack_logging
    )
endif ()
#################### xnnpack_normalization ####################

#################### xnnpack_indirection ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/indirection.c
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
    add_library(xnnpack_indirection SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_indirection
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_indirection PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_indirection
            fp16_headers
            fxdiv_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
    )
endif ()
#################### xnnpack_indirection ####################

#################### libXNNPACK ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/init.c
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
    add_library(libXNNPACK SHARED ${SOURCE_FILES})

    target_include_directories(libXNNPACK
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(libXNNPACK PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libXNNPACK
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libcpuinfo
            libdl
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_jit
            xnnpack_logging
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_operators
            xnnpack_packing
            xnnpack_params
            xnnpack_prod_microkernels
            xnnpack_subgraph
    )
endif ()
#################### libXNNPACK ####################

#################### xnnpack_mobilenet_v3_large_fp16 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp16-mobilenet-v3-large.cc
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
    add_library(xnnpack_mobilenet_v3_large_fp16 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v3_large_fp16
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v3_large_fp16 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v3_large_fp16
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v3_large_fp16 ####################

#################### xnnpack_mobilenet_v2_fp16 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp16-mobilenet-v2.cc
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
    add_library(xnnpack_mobilenet_v2_fp16 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v2_fp16
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v2_fp16 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v2_fp16
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v2_fp16 ####################

#################### xnnpack_mobilenet_v2_fp32_sparse ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp32-sparse-mobilenet-v2.cc
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
    add_library(xnnpack_mobilenet_v2_fp32_sparse SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v2_fp32_sparse
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v2_fp32_sparse PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v2_fp32_sparse
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v2_fp32_sparse ####################

#################### xnnpack_neonv8_bench_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neonv8-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neonv8-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neonv8-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neonv8-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neonv8-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neonv8-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-neonv8-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-neonv8-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-neonv8-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-neonv8-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-neonv8-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f32-qs8-neonv8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f32-qu8-neonv8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundd-neonv8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundne-neonv8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundu-neonv8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundz-neonv8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x3-minmax-fp32-neonv8-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x9-minmax-fp32-neonv8-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x9-minmax-fp32-neonv8-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x25-minmax-fp32-neonv8-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x25-minmax-fp32-neonv8-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x3-minmax-fp32-neonv8-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x3-minmax-fp32-neonv8-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neonv8-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neonv8-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neonv8-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neonv8-mul8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neonv8-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neonv8-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neonv8-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neonv8-mul8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up24x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up24x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up32x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up32x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4s2-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c4-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c4-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c4-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c4s2-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4s2-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c4-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c4-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c4-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c4s2-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x8-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x8-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x16-minmax-fp32-neonv8-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up24x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up24x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up32x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up32x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-neonv8-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-neonv8-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-neonv8-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-neonv8-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-neonv8-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-neonv8-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-neonv8-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-neonv8-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4s2-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4s2-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-fp32-neonv8-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4s2-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-fp32-neonv8-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2s4-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-fp32-neonv8-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-fp32-neonv8-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-fp32-neonv8-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4s2-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c8-minmax-fp32-neonv8-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-neonv8-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-neonv8-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-neonv8-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-neonv8-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-neonv8-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-neonv8-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up8x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up8x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up16x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up16x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up24x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up24x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up32x9-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up32x25-minmax-fp32-neonv8-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-neonv8-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-neonv8-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-neonv8-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-neonv8-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-neonv8-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-neonv8-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-neonv8-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-neonv8-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x16-minmax-fp32-neonv8-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-neonv8-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-neonv8-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-neonv8-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-neonv8-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-neonv8-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-neonv8-ld128-x16.c
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
    add_library(xnnpack_neonv8_bench_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neonv8_bench_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neonv8_bench_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neonv8_bench_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neonv8_bench_microkernels ####################

#################### xnnpack_qc8_mobilenet_v2 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/qc8-mobilenet-v2.cc
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
    add_library(xnnpack_qc8_mobilenet_v2 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_qc8_mobilenet_v2
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_qc8_mobilenet_v2 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_qc8_mobilenet_v2
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_qc8_mobilenet_v2 ####################

#################### xnnpack_neon_bench_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-bfly4/samples1-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-vsquareabs/gen/neon-mlal-ld128-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-vsquareabs/gen/neon-mlal-ld128-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-vsquareabs/gen/neon-mlal-ld128-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/cs16-vsquareabs/gen/neon-mlal-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neon-int16-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neon-int16-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neon-int16-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neon-int16-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neon-int32-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neon-int32-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neon-int32-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neon-int32-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/4x-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/9p8x-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/9x-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-avgpool/9p8x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-avgpool/9x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p0p1c3x4-neon-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p0p1c3x4-neon-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p0p1c3x8-neon-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p0p1c3x8-neon-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p1c3x4-neon-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p1c3x4-neon-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p1c3x8-neon-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p1c3x8-neon-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc2chw/3x3s2p1c3x4-neon-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x3-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x3-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x4-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x9-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x9-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x25-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x25-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x3-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x3-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x4-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x9-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x9-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x25-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x25-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-5x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neon-6x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neon-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neon-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neon-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neon-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neon-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neon-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neon-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neon-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-1x4-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-2x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-3x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-4x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neon-5x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-1x4-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-2x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-3x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neon-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-neon-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool-cw/neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool/7p7x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool/7x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/1x8inc-minmax-neon-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/1x8inc-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/1x8s4inc-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8inc-minmax-neon-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8inc-minmax-neon-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8inc-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8inc-minmax-neon-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8s4inc-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/5x8inc-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8inc-minmax-neon-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8inc-minmax-neon-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8inc-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8inc-minmax-neon-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8s4inc-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/8x8s4inc-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x8-minmax-neon-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x8s4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x2-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neon-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neon-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neon-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8s4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/5x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x2-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8-minmax-neon-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8-minmax-neon-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8-minmax-neon-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8s4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/8x8s4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/neon-p4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/neon-p8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/neon-p16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x8-minmax-neon-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x8s4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x2-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x4-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neon-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neon-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neon-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8s4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x2-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8-minmax-neon-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8-minmax-neon-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8-minmax-neon-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8-minmax-neon-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8s4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/8x8s4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-maxpool/9p8x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-pavgpool/9p8x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-pavgpool/9x-minmax-neon-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ppmm/gen/4x8-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ppmm/gen/8x8-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-1x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-1x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-2x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-2x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-4x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/neon-4x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neon-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neon-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x8-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x12-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x12-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x16-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x16-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x20-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x20-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-lut64-p2-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x8-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x12-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x12-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x16-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x16-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x20-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x20-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neon-rr2-p5-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-rmax/neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x1-minmax-neon-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x1-minmax-neon-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x1-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x1-minmax-neon-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x1-minmax-neon-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x1-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/12x1-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/16x1-minmax-neon-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/16x1-minmax-neon-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/16x1-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x1-minmax-neon-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x1-minmax-neon-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x1-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmaxc-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmaxc-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmin-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmin-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vminc-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vminc-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiff-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiff-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiffc-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiffc-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vclamp/gen/vclamp-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vclamp/gen/vclamp-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-lut16-p3-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-lut16-p3-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-lut16-p3-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-lut16-p3-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-lut16-p3-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-lut16-p3-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-p6-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-p6-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-p6-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-p6-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-p6-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neon-rr2-p6-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vhswish/gen/vhswish-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vhswish/gen/vhswish-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vhswish/gen/vhswish-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vlrelu/gen/vlrelu-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vlrelu/gen/vlrelu-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c4-minmax-neon-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c8-minmax-neon-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrelu/gen/vrelu-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrelu/gen/vrelu-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut64-p2-nr2recps-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut64-p2-nr2recps-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut64-p2-nr2recps-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut64-p2-nr2recps-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut64-p2-nr2recps-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut64-p2-nr2recps-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut2048-p1-nr2recps-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut2048-p1-nr2recps-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut2048-p1-nr2recps-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut2048-p1-nr2recps-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut2048-p1-nr2recps-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-lut2048-p1-nr2recps-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-p5-nr2recps-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-p5-nr2recps-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-p5-nr2recps-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-p5-nr2recps-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-p5-nr2recps-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neon-rr2-p5-nr2recps-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vabs-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vabs-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vneg-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vneg-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vsqr-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vsqr-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f16-f32-neon-int16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f16-f32-neon-int32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f32-f16-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f32-qs8-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f32-qu8-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-neon-rr2-lut16-p3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-neon-rr2-p6.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundd-neon-addsub.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundd-neon-cvt.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundne-neon-addsub.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundu-neon-addsub.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundu-neon-cvt.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundz-neon-addsub.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/roundz-neon-cvt.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neon-rr2-lut64-p2-nr2recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neon-rr2-lut2048-p1-nr2recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neon-rr2-p5-nr2recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-neon-nr1rsqrts.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-neon-nr2rsqrts.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-neon-nr3rsqrts.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x3-minmax-fp32-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x9-minmax-fp32-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x9-minmax-fp32-neon-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x25-minmax-fp32-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x25-minmax-fp32-neon-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up8x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x3-minmax-fp32-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x3-minmax-fp32-neon-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neon-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neon-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neon-mul8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neon-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neon-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neon-mul8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up16x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up24x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up24x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up32x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up32x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2-minmax-fp32-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4s2-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c8-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2-minmax-fp32-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c4-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c4-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c4-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c4s2-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x8c8-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/3x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2-minmax-fp32-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4s2-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c8-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2-minmax-fp32-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c4-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c4-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c4-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c4s2-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x8c8-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/3x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x8-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x16-minmax-fp32-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x9-minmax-rndnu-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x9-minmax-rndnu-neon-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x9-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x25-minmax-rndnu-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x25-minmax-rndnu-neon-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up8x25-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x9-minmax-rndnu-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x9-minmax-rndnu-neon-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x9-minmax-rndnu-neon-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x9-minmax-rndnu-neon-mul8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x9-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x25-minmax-rndnu-neon-mla8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x25-minmax-rndnu-neon-mla8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x25-minmax-rndnu-neon-mul8-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x25-minmax-rndnu-neon-mul8-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up16x25-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up24x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up24x9-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up24x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up24x25-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up32x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up32x9-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up32x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up32x25-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-neon-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-neon-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-neon-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-rndnu-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-rndnu-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-rndnu-neon-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-rndnu-neon-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-neon-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-neon-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-rndnu-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-rndnu-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-rndnu-neon-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-rndnu-neon-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-fp32-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4s2-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c8-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-fp32-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4s2-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c8-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x8c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x16c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x8c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/3x16c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/6x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/6x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/6x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/6x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-fp32-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4s2-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c8-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-fp32-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2s4-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-fp32-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-fp32-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-fp32-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4s2-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c8-minmax-fp32-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x8c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x16c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x8c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/3x16c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16-minmax-rndnu-neon-mull-addw-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2-minmax-rndnu-neon-mlal-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2-minmax-rndnu-neon-mull-ld4r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2s4-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c2s4-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4-minmax-rndnu-neon-mlal-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4-minmax-rndnu-neon-mlal-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4-minmax-rndnu-neon-mlal-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4-minmax-rndnu-neon-mull-dup.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4-minmax-rndnu-neon-mull-ld1r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4-minmax-rndnu-neon-mull-ld2r.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4s2-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4s2-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c8-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c8-minmax-rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c16-minmax-rndnu-neon-mlal.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/6x8-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/6x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/6x16-minmax-rndnu-neon-mlal-lane-prfm.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/6x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/fp32-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/gemmlowp-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/rndna-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/rndnu-neon-mull.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-requantization/rndnu-neon-qdmulh.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-neon-ld64-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-neon-ld64-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-neon-ld128-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-neon-ld64-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-neon-ld64-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-neon-ld128-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vcvt/gen/vcvt-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vcvt/gen/vcvt-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vlrelu/gen/vlrelu-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-rndnu-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-rndnu-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-rndnu-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-rndnu-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-rndnu-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-rndnu-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-avgpool/9p8x-minmax-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-avgpool/9x-minmax-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up8x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up8x9-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up8x9-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up8x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up8x25-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up8x25-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up16x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up16x9-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up16x9-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up16x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up16x25-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up16x25-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up24x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up24x9-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up24x9-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up24x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up24x25-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up24x25-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up32x9-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up32x9-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up32x9-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up32x25-minmax-fp32-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up32x25-minmax-rndnu-neon-mul8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up32x25-minmax-rndnu-neon-mul16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-neon-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-neon-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-neon-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-rndnu-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-rndnu-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-rndnu-neon-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-rndnu-neon-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-neon-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-neon-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-rndnu-neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-rndnu-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-rndnu-neon-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-rndnu-neon-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/6x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/6x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x8-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x16-minmax-fp32-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/6x8-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/6x16-minmax-rndnu-neon-mlal-lane.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-requantization/fp32-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-requantization/gemmlowp-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-requantization/rndna-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-neon-ld64-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-neon-ld64-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vcvt/gen/vcvt-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vcvt/gen/vcvt-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vcvt/gen/vcvt-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vlrelu/gen/vlrelu-neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-rndnu-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-rndnu-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-rndnu-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-rndnu-neon-ld64-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-rndnu-neon-ld64-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-rndnu-neon-ld128-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-ibilinear/gen/neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-ibilinear/gen/neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-maxpool/2p2x-minmax-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-maxpool/4p3x-minmax-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-maxpool/9p8x-minmax-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-vclamp/neon-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-rmaxabs/gen/neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-rmaxabs/gen/neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-rmaxabs/gen/neon-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-rmaxabs/gen/neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-vlshift/gen/neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-vlshift/gen/neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-vlshift/gen/neon-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-vlshift/gen/neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-shift12-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-shift12-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-shift12-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-shift12-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-shift15-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-shift15-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-shift15-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-shift15-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s16-window/gen/neon-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-ibilinear/gen/neon-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-ibilinear/gen/neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-maxpool/9p8x-minmax-neon-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-rmax/neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-vclamp/neon-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u32-filterbank-accumulate/gen/neon-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u32-filterbank-accumulate/gen/neon-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-fill/neon-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-pad/neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/8x8-multi-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/8x8-multi-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/8x8-multi-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/8x8-reuse-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/8x8-reuse-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/8x8-reuse-multi-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/8x8-reuse-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/16x16-reuse-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/16x16-reuse-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/16x16-reuse-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/xm-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x2-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x3-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x4-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x4-multi-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x4-multi-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x4-multi-multi-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x4-multi-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x4-reuse-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x4-reuse-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x4-reuse-multi-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/4x4-reuse-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/8x8-multi-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/8x8-multi-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/8x8-multi-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/8x8-reuse-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/8x8-reuse-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/8x8-reuse-multi-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/8x8-reuse-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/2x2-neon-tbl.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-packx/x4-neon-st4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-multi-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-multi-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-multi-multi-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-multi-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-reuse-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-reuse-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-reuse-multi-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x2-reuse-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-multi-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-multi-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-multi-multi-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-multi-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-reuse-dec-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-reuse-mov-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-reuse-multi-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/4x4-reuse-switch-zip-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-unpool/neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/xm-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x2-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x3-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x4-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/1x4c8-minmax-neonfma-zip.c
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/2x4c8-minmax-neonfma-zip.c
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/3x4c8-minmax-neonfma-zip.c
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/4x4c8-minmax-neonfma-zip.c
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/5x4c8-minmax-neonfma-zip.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p0p1c3x4-neonfma-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p0p1c3x4-neonfma-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p0p1c3x8-neonfma-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p0p1c3x8-neonfma-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p1c3x4-neonfma-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p1c3x4-neonfma-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p1c3x8-neonfma-2x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/gen/3x3s2p1c3x8-neonfma-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc2chw/3x3s2p1c3x4-neonfma-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-5x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-neonfma-6x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neonfma-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neonfma-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neonfma-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neonfma-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neonfma-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neonfma-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neonfma-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-neonfma-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-1x4-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-2x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-3x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-4x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-neonfma-5x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-1x4-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-2x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-3x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-neonfma-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/1x8inc-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8inc-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8inc-minmax-neonfma-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/5x8inc-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8inc-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8inc-minmax-neonfma-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x2-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neonfma-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/5x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x2-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8-minmax-neonfma-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x2-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x4-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neonfma-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x2-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8-minmax-neonfma-lane-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8-minmax-neonfma-lane-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x2-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x2-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/12x2-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/12x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/16x2-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/16x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x2-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-minmax-neon-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-minmax-neon-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-div-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-div-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-div-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-div-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-div-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-div-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-div-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-div-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-div-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-div-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-div-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-div-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-div-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-div-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-div-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-div-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-div-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-div-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neon-sqrt-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neon-sqrt-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-lut64-p2-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-lut2048-p1-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-p5-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-lut64-p2-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-lut2048-p1-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-p5-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-neon-tbx128x4-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-neon-tbx128x4-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-neon-tbx128x4-x48.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-neon-tbx128x4-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x24-transposec/4x4-aarch64-tbl.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/4x4-aarch64-tbl.c
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
    add_library(xnnpack_neon_bench_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neon_bench_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neon_bench_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neon_bench_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neon_bench_microkernels ####################

#################### xnnpack_im2col ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/im2col.c
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
    add_library(xnnpack_im2col SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_im2col
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_im2col PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_im2col
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### xnnpack_im2col ####################

#################### xnnpack_subgraph ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/memory-planner.c
        ${ANDROID_ROOT}/external/XNNPACK/src/runtime.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/abs.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/add2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/argmax-pooling-2d.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/average-pooling-2d.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/bankers-rounding.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/ceiling.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/clamp.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/concatenate.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/convert.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/convolution-2d.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/deconvolution-2d.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/depth-to-space.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/depthwise-convolution-2d.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/divide.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/elu.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/even-split.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/floor.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/fully-connected.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/global-average-pooling.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/hardswish.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/leaky-relu.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/max-pooling-2d.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/maximum2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/minimum2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/multiply2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/negate.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/prelu.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/sigmoid.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/softmax.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/square-root.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/square.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/squared-difference.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/static-constant-pad.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/static-reshape.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/static-resize-bilinear-2d.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/static-transpose.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/subtract.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/unpooling-2d.c
        ${ANDROID_ROOT}/external/XNNPACK/src/subgraph/validation.c
        ${ANDROID_ROOT}/external/XNNPACK/src/tensor.c
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
    add_library(xnnpack_subgraph SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_subgraph
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_subgraph PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_subgraph
            fp16_headers
            fxdiv_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_logging
            xnnpack_microparams_init
            xnnpack_normalization
            xnnpack_packing
            xnnpack_post_operation
    )
endif ()
#################### xnnpack_subgraph ####################

#################### xnnpack_neonfp16_prod_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neonfp16-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-neonfp16-x16.c
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
    add_library(xnnpack_neonfp16_prod_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neonfp16_prod_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neonfp16_prod_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neonfp16_prod_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neonfp16_prod_microkernels ####################

#################### xnnpack_qc8_mobilenet_v1 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/qc8-mobilenet-v1.cc
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
    add_library(xnnpack_qc8_mobilenet_v1 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_qc8_mobilenet_v1
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_qc8_mobilenet_v1 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_qc8_mobilenet_v1
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_qc8_mobilenet_v1 ####################

#################### xnnpack_neondot_bench_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/6x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/8x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/8x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/6x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/8x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/8x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/6x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/6x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/8x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/8x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/6x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/6x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/8x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/8x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x32c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x32c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/3x32c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/5x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/5x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/6x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/6x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/8x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/8x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x32c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x32c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/3x32c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/5x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/5x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/6x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/6x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/8x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/8x16c4-minmax-rndnu-neondot.c
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
    add_library(xnnpack_neondot_bench_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neondot_bench_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neondot_bench_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neondot_bench_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neondot_bench_microkernels ####################

#################### xnnpack_cache ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/cache.c
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
    add_library(xnnpack_cache SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_cache
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_cache PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_cache
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_logging
            xnnpack_mutex
    )
endif ()
#################### xnnpack_cache ####################

#################### xnnpack_neonfp16_bench_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neonfp16-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-neonfp16-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-neonfp16-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-neonfp16-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f16-f32-neonfp16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/cvt-f32-f16-neonfp16.c
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
    add_library(xnnpack_neonfp16_bench_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neonfp16_bench_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neonfp16_bench_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neonfp16_bench_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neonfp16_bench_microkernels ####################

#################### xnnpack_neondot_prod_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/4x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x8c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/4x16c4-minmax-fp32-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/4x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/4x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/4x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x16c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x8c4-minmax-rndnu-neondot.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/4x16c4-minmax-rndnu-neondot.c
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
    add_library(xnnpack_neondot_prod_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neondot_prod_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neondot_prod_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neondot_prod_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neondot_prod_microkernels ####################

#################### xnnpack_mobilenet_v1_fp16 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp16-mobilenet-v1.cc
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
    add_library(xnnpack_mobilenet_v1_fp16 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v1_fp16
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v1_fp16 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v1_fp16
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v1_fp16 ####################

#################### xnnpack_post_operation ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/operators/post-operation.c
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
    add_library(xnnpack_post_operation SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_post_operation
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_post_operation PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_post_operation
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_params
    )
endif ()
#################### xnnpack_post_operation ####################

#################### xnnpack_mobilenet_v1_fp32 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp32-mobilenet-v1.cc
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
    add_library(xnnpack_mobilenet_v1_fp32 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v1_fp32
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v1_fp32 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v1_fp32
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v1_fp32 ####################

#################### xnnpack_neonfp16arith_bench_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-avgpool/9p8x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-avgpool/9x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up8x3-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up8x3-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up8x4-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up8x4-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up8x9-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up8x9-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up8x25-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up8x25-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x3-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x3-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x4-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x4-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x9-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x9-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x25-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up16x25-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up32x3-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up32x3-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up32x4-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up32x4-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up32x9-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up32x9-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up32x25-minmax-neonfp16arith-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv/gen/up32x25-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool-cw/neonfp16arith-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool-cw/neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7p7x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7p7x-minmax-neonfp16arith-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7p7x-minmax-neonfp16arith-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7p7x-minmax-neonfp16arith-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7x-minmax-neonfp16arith-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7x-minmax-neonfp16arith-c24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gavgpool/gen/7x-minmax-neonfp16arith-c32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen-inc/1x8inc-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen-inc/1x16inc-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen-inc/4x8inc-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen-inc/4x16inc-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen-inc/6x8inc-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen-inc/6x16inc-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen-inc/8x8inc-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen-inc/8x16inc-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/1x8-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/1x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/4x8-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/4x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/6x8-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/6x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/8x8-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-gemm/gen/8x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-ibilinear-chw/gen/neonfp16arith-p4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-ibilinear-chw/gen/neonfp16arith-p8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-ibilinear-chw/gen/neonfp16arith-p16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-ibilinear/gen/neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-ibilinear/gen/neonfp16arith-c16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/1x8-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/1x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/4x8-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/4x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/6x8-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/6x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/8x8-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-igemm/gen/8x16-minmax-neonfp16arith-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-maxpool/9p8x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-pavgpool/9p8x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-pavgpool/9x-minmax-neonfp16arith-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-prelu/gen/neonfp16arith-2x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-prelu/gen/neonfp16arith-2x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x32-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x32-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x40-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x40-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x40.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x48-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x48-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x48.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x64-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x64-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x72-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x72.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x80-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x80-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x80.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x96-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x96-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x96-acc6.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-raddstoreexpminusmax/gen/neonfp16arith-rr2-p2-x96.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-rmax/neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-spmm/gen/8x1-minmax-neonfp16arith-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-spmm/gen/8x1-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-spmm/gen/16x1-minmax-neonfp16arith-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-spmm/gen/16x1-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-spmm/gen/24x1-minmax-neonfp16arith-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-spmm/gen/24x1-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-spmm/gen/32x1-minmax-neonfp16arith-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-spmm/gen/32x1-minmax-neonfp16arith.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vadd-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vadd-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vaddc-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vaddc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmaxc-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmaxc-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmin-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmin-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vminc-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vminc-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmul-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmul-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmulc-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vmulc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vrsubc-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vrsubc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsqrdiff-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsqrdiff-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsqrdiffc-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsqrdiffc-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsub-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsub-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsubc-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vsubc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vclamp/gen/vclamp-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vclamp/gen/vclamp-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-velu/gen/velu-neonfp16arith-rr1-p3-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-velu/gen/velu-neonfp16arith-rr1-p3-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vhswish/gen/vhswish-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vhswish/gen/vhswish-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vlrelu/gen/vlrelu-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vlrelu/gen/vlrelu-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vmulcaddc/gen/c8-minmax-neonfp16arith-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vmulcaddc/gen/c16-minmax-neonfp16arith-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndd-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndd-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndne-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndne-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndu-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndu-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndz-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vrnd/gen/vrndz-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1fma-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1fma-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1fma-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1fma-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1fma-x40.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1fma-x48.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1fma-x56.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1fma-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1recps-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1recps-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1recps-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1recps-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1recps-x40.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1recps-x48.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1recps-x56.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-nr1recps-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vunary/gen/vabs-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vunary/gen/vabs-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vunary/gen/vneg-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vunary/gen/vneg-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vunary/gen/vsqr-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vunary/gen/vsqr-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/exp-f16-neonfp16arith-rr2-p3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f16-neonfp16arith-rr1-p2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f16-neonfp16arith-rr1-p3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f16-neonfp16arith-rr2-p2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f16-neonfp16arith-rr2-p3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f16-neonfp16arith-rr1-p3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f16-neonfp16arith-rr2-p3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr2-p2-nr1fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr2-p2-nr1recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr2-p2-recpe.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr2-p3-nr1fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr2-p3-nr1recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr2-p3-recpe.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-conv-hwc2chw/3x3s2p1c3x4-neonfp16arith-2x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-1x8-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-1x8-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-1x8-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-1x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-2x8-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-2x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-3x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-4x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-5x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3p1-minmax-neonfp16arith-6x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3s2p1-minmax-neonfp16arith-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3s2p1-minmax-neonfp16arith-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3s2p1-minmax-neonfp16arith-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3s2p1-minmax-neonfp16arith-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3s2p1-minmax-neonfp16arith-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3s2p1-minmax-neonfp16arith-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3s2p1-minmax-neonfp16arith-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/3x3s2p1-minmax-neonfp16arith-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-1x4-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-2x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-3x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-4x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-4x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5p2-minmax-neonfp16arith-5x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-1x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-1x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-1x4-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-1x4-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-1x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-2x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-2x4-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-3x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-dwconv2d-chw/gen/5x5s2p2-minmax-neonfp16arith-3x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vdiv-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vdiv-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vdivc-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vdivc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vrdivc-minmax-neonfp16arith-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vbinary/gen/vrdivc-minmax-neonfp16arith-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-div-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-div-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-div-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-div-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-div-x40.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-div-x48.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-div-x56.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsigmoid/gen/vsigmoid-neonfp16arith-rr2-p2-div-x64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsqrt/gen/neonfp16arith-sqrt-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-vsqrt/gen/neonfp16arith-sqrt-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr1-p2-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr1-p3-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr2-p2-div.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f16-neonfp16arith-rr2-p3-div.c
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
    add_library(xnnpack_neonfp16arith_bench_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neonfp16arith_bench_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neonfp16arith_bench_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neonfp16arith_bench_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neonfp16arith_bench_microkernels ####################

#################### xnnpack_neonfma_bench_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/1x4c8-minmax-neonfma-shland.c
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/2x4c8-minmax-neonfma-shland.c
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/3x4c8-minmax-neonfma-shland.c
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/4x4c8-minmax-neonfma-shland.c
        ${ANDROID_ROOT}/external/XNNPACK/src/bf16-gemm/gen/5x4c8-minmax-neonfma-shland.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x3-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x3-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x4-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x9-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x9-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x25-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up4x25-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x3-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x3-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x4-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x9-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x9-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x25-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up8x25-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x3-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x3-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x3-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x3-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x4-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x4-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x4-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x9-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x9-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x9-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x9-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x25-minmax-neon-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x25-minmax-neon.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x25-minmax-neonfma-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up16x25-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/1x8inc-minmax-neonfma-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/1x8s4inc-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8inc-minmax-neonfma-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8inc-minmax-neonfma-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/4x8s4inc-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8inc-minmax-neonfma-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8inc-minmax-neonfma-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/6x8s4inc-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen-inc/8x8s4inc-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x8-minmax-neonfma-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neonfma-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8-minmax-neonfma-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8-minmax-neonfma-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8-minmax-neonfma-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/6x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/8x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/neonfma-p4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/neonfma-p8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/neonfma-p16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/neonfma-c4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/neonfma-c8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x8-minmax-neonfma-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neonfma-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8-minmax-neonfma-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8-minmax-neonfma-dup-ld64.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8-minmax-neonfma-dup-ld128.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/6x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/8x8s4-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ppmm/gen/4x8-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ppmm/gen/8x8-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x8-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x12-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x12-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x16-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x16-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x20-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x20-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-lut64-p2-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x8-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x12-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x12-acc3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x16-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x16-acc4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x20-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x20-acc5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/neonfma-rr1-p5-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x1-minmax-neonfma-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x1-minmax-neonfma-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/4x1-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x1-minmax-neonfma-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x1-minmax-neonfma-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x1-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/12x1-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/16x1-minmax-neonfma-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/16x1-minmax-neonfma-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/16x1-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x1-minmax-neonfma-pipelined.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x1-minmax-neonfma-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/32x1-minmax-neonfma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-lut16-p3-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-lut16-p3-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-lut16-p3-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-lut16-p3-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-lut16-p3-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-lut16-p3-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-p6-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-p6-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-p6-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-p6-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-p6-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-neonfma-rr1-p6-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c4-minmax-neonfma-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c8-minmax-neonfma-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr1recps1fma-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr1recps1fma-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr1recps1fma-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr1recps1fma-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr1recps1fma-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr1recps1fma-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2fma-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2fma-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2fma-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2fma-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2fma-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2fma-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2recps-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2recps-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2recps-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2recps-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2recps-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut64-p2-nr2recps-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr1recps1fma-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr1recps1fma-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr1recps1fma-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr1recps1fma-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr1recps1fma-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr1recps1fma-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2fma-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2fma-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2fma-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2fma-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2fma-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2fma-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2recps-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2recps-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2recps-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2recps-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2recps-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-lut2048-p1-nr2recps-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr1recps1fma-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr1recps1fma-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr1recps1fma-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr1recps1fma-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr1recps1fma-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr1recps1fma-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2fma-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2fma-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2fma-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2fma-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2fma-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2fma-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2recps-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2recps-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2recps-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2recps-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2recps-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-neonfma-rr1-p5-nr2recps-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x28.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x36.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr1rsqrts1fma1adj-x40.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x12.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x20.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x24.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x28.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x32.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x36.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/neonfma-nr2fma1adj-x40.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/exp-f32-neonfma-rr2-lut64-p2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/exp-f32-neonfma-rr2-p5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f32-neonfma-rr2-lut64-p2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f32-neonfma-rr2-lut2048-p1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expminus-f32-neonfma-rr2-p5.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-neonfma-rr1-lut16-p3.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/expm1minus-f32-neonfma-rr1-p6.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-lut64-p2-nr1recps1fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-lut64-p2-nr2fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-lut64-p2-nr2recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-lut2048-p1-nr1recps1fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-lut2048-p1-nr2fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-lut2048-p1-nr2recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-p5-nr1recps1fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-p5-nr2fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr1-p5-nr2recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-lut64-p2-nr1recps1fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-lut64-p2-nr2fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-lut64-p2-nr2recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-lut2048-p1-nr1recps1fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-lut2048-p1-nr2fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-lut2048-p1-nr2recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-p5-nr1recps1fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-p5-nr2fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sigmoid-f32-neonfma-rr2-p5-nr2recps.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-neonfma-nr1fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-neonfma-nr1rsqrts1fma1adj.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-neonfma-nr2fma.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-neonfma-nr2fma1adj.c
        ${ANDROID_ROOT}/external/XNNPACK/src/math/sqrt-neonfma-nr3fma.c
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
    add_library(xnnpack_neonfma_bench_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_neonfma_bench_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_neonfma_bench_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_neonfma_bench_microkernels
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_neonfma_bench_microkernels ####################

#################### xnnpack_microparams_init ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/microparams-init.c
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
    add_library(xnnpack_microparams_init SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_microparams_init
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_microparams_init PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_microparams_init
            fp16_headers
            fxdiv_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
    )
endif ()
#################### xnnpack_microparams_init ####################

#################### xnnpack_logging ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/datatype-strings.c
        ${ANDROID_ROOT}/external/XNNPACK/src/log.c
        ${ANDROID_ROOT}/external/XNNPACK/src/node-type.c
        ${ANDROID_ROOT}/external/XNNPACK/src/operator-strings.c
        ${ANDROID_ROOT}/external/XNNPACK/src/ukernel-strings.c
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
    add_library(xnnpack_logging SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_logging
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_logging PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_logging
            fp16_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            liblog
            libm
            libpthreadpool
    )
endif ()
#################### xnnpack_logging ####################

#################### xnnpack_qu8_mobilenet_v2 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/qu8-mobilenet-v2.cc
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
    add_library(xnnpack_qu8_mobilenet_v2 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_qu8_mobilenet_v2
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_qu8_mobilenet_v2 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_qu8_mobilenet_v2
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_qu8_mobilenet_v2 ####################

#################### xnnpack_scalar_prod_microkernels ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-lut32norm/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-copy/memcpy.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-transpose/1x1-memcpy.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-lut/gen/lut-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f16-f32-vcvt/gen/vcvt-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/4x-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/9p8x-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-argmaxpool/9x-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-avgpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-avgpool/9x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/3x3s2p0p1c3x4-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc/3x3s2p1c3x4-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-conv-hwc2chw/3x3s2p1c3x4-scalar-1x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x3-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x3-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x4-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x4-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x9-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x9-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x25-minmax-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv/gen/up1x25-scalar-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3p1-minmax-scalar-4x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/3x3s2p1-minmax-scalar-2x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5p2-minmax-scalar-2x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-dwconv2d-chw/gen/5x5s2p2-minmax-scalar-2x1-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-f16-vcvt/gen/vcvt-scalar-fabsf-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool-cw/scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool/7p7x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gavgpool/7x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/1x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x2-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/gen/4x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear-chw/gen/scalar-p4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-ibilinear/gen/scalar-c2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/1x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x2-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x4-relu-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/gen/4x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-maxpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-pavgpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-pavgpool/9x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-prelu/gen/scalar-2x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qs8-vcvt/gen/vcvt-scalar-imagic-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-qu8-vcvt/gen/vcvt-scalar-imagic-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-raddstoreexpminusmax/gen/scalar-rr2-p5-x4-acc2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-rmax/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x1-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x2-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-spmm/gen/8x4-minmax-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vadd-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vaddc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdiv-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vdivc-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmaxc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmin-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vminc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmul-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vmulc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrdivc-minmax-scalar-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vrsubc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiff-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsqrdiffc-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsub-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vbinary/gen/vsubc-minmax-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vclamp/gen/vclamp-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-velu/gen/velu-scalar-rr2-lut16-p3-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vhswish/gen/vhswish-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vlrelu/gen/vlrelu-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vmulcaddc/gen/c1-minmax-scalar-2x.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrelu/gen/vrelu-scalar-x8.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndd-scalar-libm-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndne-scalar-libm-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndu-scalar-libm-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vrnd/gen/vrndz-scalar-libm-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsigmoid/gen/vsigmoid-scalar-rr2-lut64-p2-div-x2.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vsqrt/gen/scalar-sqrt-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vabs-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vneg-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-vunary/gen/vsqr-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x3-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-dwconv/gen/up1x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-gemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qc8-igemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up1x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-dwconv/gen/up1x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-f32-vcvt/gen/vcvt-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7p7x-minmax-fp32-scalar-imagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gavgpool/gen/7x-minmax-fp32-scalar-imagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-gemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-igemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vadd/gen/minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vaddc/gen/minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmul/gen/minmax-fp32-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qs8-vmulc/gen/minmax-fp32-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-avgpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-avgpool/9x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up1x9-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-dwconv/gen/up1x25-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-f32-vcvt/gen/vcvt-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7p7x-minmax-fp32-scalar-imagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gavgpool/gen/7x-minmax-fp32-scalar-imagic-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-gemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/1x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-igemm/gen/2x2-minmax-fp32-scalar-fmagic.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vadd/gen/minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vaddc/gen/minmax-scalar-x1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmul/gen/minmax-fp32-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/qu8-vmulc/gen/minmax-fp32-scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-ibilinear/gen/scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-maxpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/s8-vclamp/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-ibilinear/gen/scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-maxpool/9p8x-minmax-scalar-c1.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-rmax/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/u8-vclamp/scalar-x4.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-fill/scalar-x16.c
        ${ANDROID_ROOT}/external/XNNPACK/src/xx-pad/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-transposec/gen/2x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/xm-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x3-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x8-zip/x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x16-transposec/gen/2x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-packx/x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-packx/x3-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-packx/x4-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-transposec/gen/2x4-scalar-int.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-unpool/scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/xm-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x2-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x3-scalar.c
        ${ANDROID_ROOT}/external/XNNPACK/src/x32-zip/x4-scalar.c
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
    add_library(xnnpack_scalar_prod_microkernels SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_scalar_prod_microkernels
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_scalar_prod_microkernels PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_scalar_prod_microkernels
            fp16_headers
            fxdiv_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libpthreadpool
            xnnpack_tables
    )
endif ()
#################### xnnpack_scalar_prod_microkernels ####################

#################### xnnpack_jit ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/src/jit/aarch32-assembler.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/jit/aarch64-assembler.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/jit/assembler.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/upto6x8-aarch64-neonfma-cortex-a75.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/1x8-aarch64-neonfma-cortex-a75.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/4x8-aarch64-neonfma-cortex-a75.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-gemm/6x8-aarch64-neonfma-ld128.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/upto6x8-aarch64-neonfma-cortex-a75.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/1x8-aarch64-neonfma-cortex-a75.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/4x8-aarch64-neonfma-cortex-a75.cc
        ${ANDROID_ROOT}/external/XNNPACK/src/f32-igemm/6x8-aarch64-neonfma-ld128.cc
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
    add_library(xnnpack_jit SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_jit
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/include
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_jit PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(xnnpack_jit
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclog
            libdl
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_logging
            xnnpack_params
    )
endif ()
#################### xnnpack_jit ####################

#################### xnnpack_qu8_mobilenet_v1 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/qu8-mobilenet-v1.cc
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
    add_library(xnnpack_qu8_mobilenet_v1 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_qu8_mobilenet_v1
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_qu8_mobilenet_v1 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_qu8_mobilenet_v1
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_qu8_mobilenet_v1 ####################

#################### xnnpack_mobilenet_v3_large_fp32_sparse ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp32-sparse-mobilenet-v3-large.cc
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
    add_library(xnnpack_mobilenet_v3_large_fp32_sparse SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v3_large_fp32_sparse
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v3_large_fp32_sparse PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v3_large_fp32_sparse
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v3_large_fp32_sparse ####################

#################### xnnpack_mobilenet_v2_fp32 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp32-mobilenet-v2.cc
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
    add_library(xnnpack_mobilenet_v2_fp32 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v2_fp32
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v2_fp32 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v2_fp32
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v2_fp32 ####################

#################### xnnpack_mobilenet_v3_large_fp32 ####################
set(FILES
        ${ANDROID_ROOT}/external/XNNPACK/models/fp32-mobilenet-v3-large.cc
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
    add_library(xnnpack_mobilenet_v3_large_fp32 SHARED ${SOURCE_FILES})

    target_include_directories(xnnpack_mobilenet_v3_large_fp32
            PUBLIC
            ${ANDROID_ROOT}/external/XNNPACK/bench
            ${ANDROID_ROOT}/external/XNNPACK/models
            ${ANDROID_ROOT}/external/XNNPACK/test
            ${ANDROID_ROOT}/external/XNNPACK/src
            ${ANDROID_ROOT}/external/XNNPACK
            ${ANDROID_ROOT}/external/XNNPACK/include
    )

    # Add c/c++ flags
    target_compile_options(xnnpack_mobilenet_v3_large_fp32 PRIVATE
            -Werror
            -Wno-unused-function
            -Wno-unused-parameter
            -Wno-unused-private-field
    )

    target_link_libraries(xnnpack_mobilenet_v3_large_fp32
            fp16_headers
            libXNNPACK
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgmock
            liblog
            libm
            libpthreadpool
            xnnpack_allocator
            xnnpack_cache
            xnnpack_microparams_init
            xnnpack_mutex
            xnnpack_normalization
            xnnpack_params
    )
endif ()
#################### xnnpack_mobilenet_v3_large_fp32 ####################

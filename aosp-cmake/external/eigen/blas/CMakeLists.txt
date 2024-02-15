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

project(libF77blas)

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

#################### libF77blas ####################
set(FILES
        ${ANDROID_ROOT}/external/eigen/blas/single.cpp
        ${ANDROID_ROOT}/external/eigen/blas/double.cpp
        ${ANDROID_ROOT}/external/eigen/blas/complex_single.cpp
        ${ANDROID_ROOT}/external/eigen/blas/complex_double.cpp
        ${ANDROID_ROOT}/external/eigen/blas/xerbla.cpp
        ${ANDROID_ROOT}/external/eigen/blas/f2c/complexdots.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/srotm.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/srotmg.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/drotm.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/drotmg.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/lsame.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/dspmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/ssbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/chbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/sspmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/zhbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/chpmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/dsbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/zhpmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/dtbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/stbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/ctbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/ztbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/d_cnjg.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/r_cnjg.c
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
    add_library(libF77blas SHARED ${SOURCE_FILES})

    target_include_directories(libF77blas
            PUBLIC
            ${ANDROID_ROOT}/external/eigen/blas
    )

    # Add c/c++ flags
    target_compile_options(libF77blas PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libF77blas
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libF77blas ####################

#################### libF77blasV8 ####################
set(FILES
        ${ANDROID_ROOT}/external/eigen/blas/single.cpp
        ${ANDROID_ROOT}/external/eigen/blas/double.cpp
        ${ANDROID_ROOT}/external/eigen/blas/complex_single.cpp
        ${ANDROID_ROOT}/external/eigen/blas/complex_double.cpp
        ${ANDROID_ROOT}/external/eigen/blas/xerbla.cpp
        ${ANDROID_ROOT}/external/eigen/blas/f2c/complexdots.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/srotm.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/srotmg.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/drotm.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/drotmg.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/lsame.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/dspmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/ssbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/chbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/sspmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/zhbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/chpmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/dsbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/zhpmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/dtbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/stbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/ctbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/ztbmv.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/d_cnjg.c
        ${ANDROID_ROOT}/external/eigen/blas/f2c/r_cnjg.c
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
    add_library(libF77blasV8 SHARED ${SOURCE_FILES})

    target_include_directories(libF77blasV8
            PUBLIC
            ${ANDROID_ROOT}/external/eigen/blas
    )

    # Add c/c++ flags
    target_compile_options(libF77blasV8 PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libF77blasV8
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libF77blasV8 ####################

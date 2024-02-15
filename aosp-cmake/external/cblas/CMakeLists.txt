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

project(libblasV8)

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

#################### libblasV8 ####################
set(FILES
        ${ANDROID_ROOT}/external/cblas/src/cblas_caxpy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ccopy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cdotc_sub.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cdotu_sub.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgerc.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgeru.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cher.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cher2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cher2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cherk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chpr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chpr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_csscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cswap.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_csymm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_csyr2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_csyrk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctbsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctpsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctrmm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctrmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctrsm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctrsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dasum.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_daxpy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dcopy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ddot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dgbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dgemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dgemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dger.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dnrm2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_drot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_drotg.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_drotm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_drotmg.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsdot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dspmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dspr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dspr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dswap.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsymm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsymv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsyr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsyr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsyr2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsyrk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtbsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtpsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtrmm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtrmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtrsm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtrsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dzasum.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dznrm2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_globals.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_icamax.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_idamax.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_isamax.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_izamax.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sasum.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_saxpy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_scasum.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_scnrm2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_scopy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sdot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sdsdot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sgbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sgemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sgemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sger.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_snrm2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_srot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_srotg.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_srotm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_srotmg.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sspmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sspr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sspr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sswap.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssymm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssymv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssyr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssyr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssyr2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssyrk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_stbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_stbsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_stpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_stpsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_strmm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_strmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_strsm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_strsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_xerbla.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zaxpy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zcopy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zdotc_sub.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zdotu_sub.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zdscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgerc.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgeru.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zher.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zher2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zher2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zherk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhpr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhpr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zswap.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zsymm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zsyr2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zsyrk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztbsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztpsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztrmm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztrmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztrsm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztrsv.c
        ${ANDROID_ROOT}/external/cblas/src/xerbla.c
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
    add_library(libblasV8 SHARED ${SOURCE_FILES})

    target_include_directories(libblasV8
            PUBLIC
            ${ANDROID_ROOT}/external/cblas/include
            ${ANDROID_ROOT}/external/cblas
    )

    # Add c/c++ flags
    target_compile_options(libblasV8 PRIVATE
            -Werror
    )

    target_link_libraries(libblasV8
            libF77blasV8
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libblasV8 ####################

#################### libblas ####################
set(FILES
        ${ANDROID_ROOT}/external/cblas/src/cblas_caxpy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ccopy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cdotc_sub.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cdotu_sub.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgerc.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cgeru.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cher.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cher2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cher2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cherk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chpr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_chpr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_csscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_cswap.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_csymm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_csyr2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_csyrk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctbsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctpsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctrmm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctrmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctrsm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ctrsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dasum.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_daxpy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dcopy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ddot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dgbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dgemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dgemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dger.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dnrm2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_drot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_drotg.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_drotm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_drotmg.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsdot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dspmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dspr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dspr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dswap.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsymm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsymv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsyr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsyr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsyr2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dsyrk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtbsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtpsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtrmm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtrmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtrsm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dtrsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dzasum.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_dznrm2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_globals.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_icamax.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_idamax.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_isamax.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_izamax.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sasum.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_saxpy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_scasum.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_scnrm2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_scopy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sdot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sdsdot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sgbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sgemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sgemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sger.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_snrm2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_srot.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_srotg.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_srotm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_srotmg.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sspmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sspr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sspr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_sswap.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssymm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssymv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssyr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssyr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssyr2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ssyrk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_stbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_stbsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_stpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_stpsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_strmm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_strmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_strsm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_strsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_xerbla.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zaxpy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zcopy.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zdotc_sub.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zdotu_sub.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zdscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgerc.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zgeru.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhemm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhemv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zher.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zher2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zher2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zherk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhpr.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zhpr2.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zscal.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zswap.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zsymm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zsyr2k.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_zsyrk.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztbmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztbsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztpmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztpsv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztrmm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztrmv.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztrsm.c
        ${ANDROID_ROOT}/external/cblas/src/cblas_ztrsv.c
        ${ANDROID_ROOT}/external/cblas/src/xerbla.c
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
    add_library(libblas SHARED ${SOURCE_FILES})

    target_include_directories(libblas
            PUBLIC
            ${ANDROID_ROOT}/external/cblas/include
            ${ANDROID_ROOT}/external/cblas
    )

    # Add c/c++ flags
    target_compile_options(libblas PRIVATE
            -Werror
    )

    target_link_libraries(libblas
            libF77blas
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libblas ####################

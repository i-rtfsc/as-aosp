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

project(libfdlibm)

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

#################### libfdlibm ####################
set(FILES
        ${ANDROID_ROOT}/external/fdlibm/k_standard.c
        ${ANDROID_ROOT}/external/fdlibm/k_rem_pio2.c
        ${ANDROID_ROOT}/external/fdlibm/k_cos.c
        ${ANDROID_ROOT}/external/fdlibm/k_sin.c
        ${ANDROID_ROOT}/external/fdlibm/k_tan.c
        ${ANDROID_ROOT}/external/fdlibm/e_acos.c
        ${ANDROID_ROOT}/external/fdlibm/e_acosh.c
        ${ANDROID_ROOT}/external/fdlibm/e_asin.c
        ${ANDROID_ROOT}/external/fdlibm/e_atan2.c
        ${ANDROID_ROOT}/external/fdlibm/e_atanh.c
        ${ANDROID_ROOT}/external/fdlibm/e_cosh.c
        ${ANDROID_ROOT}/external/fdlibm/e_exp.c
        ${ANDROID_ROOT}/external/fdlibm/e_fmod.c
        ${ANDROID_ROOT}/external/fdlibm/e_gamma.c
        ${ANDROID_ROOT}/external/fdlibm/e_gamma_r.c
        ${ANDROID_ROOT}/external/fdlibm/e_hypot.c
        ${ANDROID_ROOT}/external/fdlibm/e_j0.c
        ${ANDROID_ROOT}/external/fdlibm/e_j1.c
        ${ANDROID_ROOT}/external/fdlibm/e_jn.c
        ${ANDROID_ROOT}/external/fdlibm/e_lgamma.c
        ${ANDROID_ROOT}/external/fdlibm/e_lgamma_r.c
        ${ANDROID_ROOT}/external/fdlibm/e_log.c
        ${ANDROID_ROOT}/external/fdlibm/e_log10.c
        ${ANDROID_ROOT}/external/fdlibm/e_pow.c
        ${ANDROID_ROOT}/external/fdlibm/e_rem_pio2.c
        ${ANDROID_ROOT}/external/fdlibm/e_remainder.c
        ${ANDROID_ROOT}/external/fdlibm/e_scalb.c
        ${ANDROID_ROOT}/external/fdlibm/e_sinh.c
        ${ANDROID_ROOT}/external/fdlibm/e_sqrt.c
        ${ANDROID_ROOT}/external/fdlibm/w_acos.c
        ${ANDROID_ROOT}/external/fdlibm/w_acosh.c
        ${ANDROID_ROOT}/external/fdlibm/w_asin.c
        ${ANDROID_ROOT}/external/fdlibm/w_atan2.c
        ${ANDROID_ROOT}/external/fdlibm/w_atanh.c
        ${ANDROID_ROOT}/external/fdlibm/w_cosh.c
        ${ANDROID_ROOT}/external/fdlibm/w_exp.c
        ${ANDROID_ROOT}/external/fdlibm/w_fmod.c
        ${ANDROID_ROOT}/external/fdlibm/w_gamma.c
        ${ANDROID_ROOT}/external/fdlibm/w_gamma_r.c
        ${ANDROID_ROOT}/external/fdlibm/w_hypot.c
        ${ANDROID_ROOT}/external/fdlibm/w_j0.c
        ${ANDROID_ROOT}/external/fdlibm/w_j1.c
        ${ANDROID_ROOT}/external/fdlibm/w_jn.c
        ${ANDROID_ROOT}/external/fdlibm/w_lgamma.c
        ${ANDROID_ROOT}/external/fdlibm/w_lgamma_r.c
        ${ANDROID_ROOT}/external/fdlibm/w_log.c
        ${ANDROID_ROOT}/external/fdlibm/w_log10.c
        ${ANDROID_ROOT}/external/fdlibm/w_pow.c
        ${ANDROID_ROOT}/external/fdlibm/w_remainder.c
        ${ANDROID_ROOT}/external/fdlibm/w_scalb.c
        ${ANDROID_ROOT}/external/fdlibm/w_sinh.c
        ${ANDROID_ROOT}/external/fdlibm/w_sqrt.c
        ${ANDROID_ROOT}/external/fdlibm/s_asinh.c
        ${ANDROID_ROOT}/external/fdlibm/s_atan.c
        ${ANDROID_ROOT}/external/fdlibm/s_cbrt.c
        ${ANDROID_ROOT}/external/fdlibm/s_ceil.c
        ${ANDROID_ROOT}/external/fdlibm/s_copysign.c
        ${ANDROID_ROOT}/external/fdlibm/s_cos.c
        ${ANDROID_ROOT}/external/fdlibm/s_erf.c
        ${ANDROID_ROOT}/external/fdlibm/s_expm1.c
        ${ANDROID_ROOT}/external/fdlibm/s_fabs.c
        ${ANDROID_ROOT}/external/fdlibm/s_finite.c
        ${ANDROID_ROOT}/external/fdlibm/s_floor.c
        ${ANDROID_ROOT}/external/fdlibm/s_frexp.c
        ${ANDROID_ROOT}/external/fdlibm/s_ilogb.c
        ${ANDROID_ROOT}/external/fdlibm/s_isnan.c
        ${ANDROID_ROOT}/external/fdlibm/s_ldexp.c
        ${ANDROID_ROOT}/external/fdlibm/s_lib_version.c
        ${ANDROID_ROOT}/external/fdlibm/s_log1p.c
        ${ANDROID_ROOT}/external/fdlibm/s_logb.c
        ${ANDROID_ROOT}/external/fdlibm/s_matherr.c
        ${ANDROID_ROOT}/external/fdlibm/s_modf.c
        ${ANDROID_ROOT}/external/fdlibm/s_nextafter.c
        ${ANDROID_ROOT}/external/fdlibm/s_rint.c
        ${ANDROID_ROOT}/external/fdlibm/s_scalbn.c
        ${ANDROID_ROOT}/external/fdlibm/s_signgam.c
        ${ANDROID_ROOT}/external/fdlibm/s_significand.c
        ${ANDROID_ROOT}/external/fdlibm/s_sin.c
        ${ANDROID_ROOT}/external/fdlibm/s_tan.c
        ${ANDROID_ROOT}/external/fdlibm/s_tanh.c
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
    add_library(libfdlibm SHARED ${SOURCE_FILES})

    target_include_directories(libfdlibm
            PUBLIC
            ${ANDROID_ROOT}/external/fdlibm
    )

    # Add c/c++ flags
    target_compile_options(libfdlibm PRIVATE
            -Werror
    )

    target_link_libraries(libfdlibm
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libfdlibm ####################

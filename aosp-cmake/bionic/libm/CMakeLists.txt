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

project(libm)

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

#################### libm ####################
set(FILES
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/bsdsrc/b_tgamma.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/catrig.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/catrigf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_acos.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_acosf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_acosh.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_acoshf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_asin.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_asinf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_atan2.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_atan2f.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_atanh.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_atanhf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_cosh.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_coshf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_fmod.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_fmodf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_gamma.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_gammaf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_gammaf_r.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_gamma_r.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_hypot.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_hypotf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_j0.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_j0f.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_j1.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_j1f.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_jn.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_jnf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_lgamma.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_lgammaf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_lgammaf_r.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_lgamma_r.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_log10.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_log10f.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_remainder.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_remainderf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_rem_pio2.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_rem_pio2f.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_scalb.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_scalbf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_sinh.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_sinhf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/k_cos.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/k_cosf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/k_exp.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/k_expf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/k_rem_pio2.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/k_sin.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/k_sinf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/k_tan.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/k_tanf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_asinh.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_asinhf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_atan.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_atanf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_carg.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cargf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cargl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cbrt.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cbrtf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_ccosh.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_ccoshf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cexp.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cexpf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cimag.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cimagf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cimagl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_clog.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_clogf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_conj.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_conjf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_conjl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cos.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cospi.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cpow.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cpowf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cpowl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cproj.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cprojf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cprojl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_creal.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_crealf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_creall.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_csinh.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_csinhf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_csqrt.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_csqrtf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_ctanh.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_ctanhf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_erf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_erff.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_expm1.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_expm1f.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_fdim.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_finite.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_finitef.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_frexp.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_frexpf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_ilogb.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_ilogbf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_log1p.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_log1pf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_logb.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_logbf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_modf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_modff.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_nan.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_nearbyint.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_nextafter.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_nextafterf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_remquo.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_remquof.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_scalbln.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_scalbn.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_scalbnf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_signgam.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_significand.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_significandf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_sin.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_sinpi.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_sincos.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_tan.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_tanf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_tanh.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_tanhf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_tgammaf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/w_cabs.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/w_cabsf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/w_cabsl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/w_drem.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/w_dremf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-netbsd/lib/libm/complex/ccoshl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-netbsd/lib/libm/complex/ccosl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-netbsd/lib/libm/complex/cephes_subrl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-netbsd/lib/libm/complex/cexpl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-netbsd/lib/libm/complex/csinhl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-netbsd/lib/libm/complex/csinl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-netbsd/lib/libm/complex/ctanhl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-netbsd/lib/libm/complex/ctanl.c
        ${ANDROID_ROOT}/bionic/libm/digittoint.c
        ${ANDROID_ROOT}/bionic/libm/significandl.c
        ${ANDROID_ROOT}/bionic/libm/fake_long_double.c
        ${ANDROID_ROOT}/bionic/libm/builtins.cpp
        ${ANDROID_ROOT}/bionic/libm/signbit.cpp
        ${ANDROID_ROOT}/bionic/libm/arm64/fenv.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/catrigl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_acosl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_acoshl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_asinl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_atan2l.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_atanhl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_fmodl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_hypotl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_lgammal.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_remainderl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_sqrtl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_asinhl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_atanl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cbrtl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_ceill.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_clogl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_coshl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_cosl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_csqrtl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_floorl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_fmal.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_fmaxl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_fminl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_modfl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_frexpl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_ilogbl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_llrintl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_llroundl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_logbl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_lrintl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_lroundl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_nextafterl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_nexttoward.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_nexttowardf.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_remquol.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_rintl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_roundl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_scalbnl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_sincosl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/e_sinhl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_sinl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_tanhl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_tanl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src/s_truncl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/invtrig.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/e_lgammal_r.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/e_powl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/k_cosl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/k_sinl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/k_tanl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/s_erfl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/s_exp2l.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/s_expl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/s_logl.c
        ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128/s_nanl.c
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
    add_library(libm SHARED ${SOURCE_FILES})

    target_include_directories(libm
            PUBLIC
            ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/ld128
            ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/android/include
            ${ANDROID_ROOT}/bionic/libm/upstream-freebsd/lib/msun/src
            ${ANDROID_ROOT}/bionic/libm
    )

    # Add c/c++ flags
    target_compile_options(libm PRIVATE
            -Wall
            -Werror
            -Wno-unused-const-variable
            -Wno-unused-variable
    )

    target_link_libraries(libm
            libarm-optimized-routines-math
            libc
            libclang_rt.builtins
    )
endif ()
#################### libm ####################

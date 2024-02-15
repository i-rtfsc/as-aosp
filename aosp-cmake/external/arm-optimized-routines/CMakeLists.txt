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

project(libarm-optimized-routines-math)

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

#################### libarm-optimized-routines-math ####################
set(FILES
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/cosf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/exp.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/exp2.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/exp2f.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/exp2f_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/exp_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/expf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/log.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/log2.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/log2_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/log2f.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/log2f_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/log_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/logf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/logf_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/math_err.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/math_errf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/pow.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/pow_log_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/powf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/powf_log2_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_cos.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_cosf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_exp.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_exp2f.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_exp2f_1u.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_expf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_expf_1u.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_log.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_logf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_pow.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_powf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_sin.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/s_sinf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/sincosf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/sincosf_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/sinf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_cos.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_cosf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_exp.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_exp2f.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_exp2f_1u.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_exp_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_expf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_expf_1u.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_log.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_log_data.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_logf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_pow.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_powf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_sin.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/v_sinf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_cos.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_cosf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_exp.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_exp2f.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_exp2f_1u.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_expf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_expf_1u.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_log.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_logf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_pow.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_powf.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_sin.c
        ${ANDROID_ROOT}/external/arm-optimized-routines/math/vn_sinf.c
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
    add_library(libarm-optimized-routines-math SHARED ${SOURCE_FILES})

    target_include_directories(libarm-optimized-routines-math
            PUBLIC
            ${ANDROID_ROOT}/external/arm-optimized-routines/math/include
            ${ANDROID_ROOT}/external/arm-optimized-routines
    )

    # Add c/c++ flags
    target_compile_options(libarm-optimized-routines-math PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libarm-optimized-routines-math
            libc_headers
            libclang_rt.builtins
    )
endif ()
#################### libarm-optimized-routines-math ####################

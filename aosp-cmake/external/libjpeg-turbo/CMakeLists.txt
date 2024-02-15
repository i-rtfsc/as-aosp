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

project(tjbench)

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

#################### tjbench ####################
set(FILES
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdatadst-tj.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdatasrc-tj.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/rdbmp.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/rdppm.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/tjbench.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/tjutil.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/transupp.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/turbojpeg.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/wrbmp.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/wrppm.c
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
    add_executable(tjbench ${SOURCE_FILES})

    target_include_directories(tjbench
            PUBLIC
            ${ANDROID_ROOT}/external/libjpeg-turbo
    )

    # Add c/c++ flags
    target_compile_options(tjbench PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(tjbench
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libjpeg
            libm
    )
endif ()
#################### tjbench ####################

#################### libjpeg ####################
set(FILES
        ${ANDROID_ROOT}/external/libjpeg-turbo/jaricom.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcapimin.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcapistd.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcarith.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jccoefct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jccolor.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcdctmgr.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jchuff.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcicc.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcinit.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcmainct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcmarker.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcmaster.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcomapi.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcparam.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcphuff.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcprepct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcsample.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jctrans.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdapimin.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdapistd.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdarith.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdatadst.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdatasrc.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdcoefct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdcolor.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jddctmgr.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdhuff.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdicc.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdinput.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdmainct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdmarker.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdmaster.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdmerge.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdphuff.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdpostct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdsample.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdtrans.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jerror.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jfdctflt.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jfdctfst.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jfdctint.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jidctflt.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jidctfst.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jidctint.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jidctred.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jmemmgr.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jmemnobs.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jpeg_nbits_table.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jquant1.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jquant2.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jutils.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/aarch64/jchuff-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/aarch64/jsimd.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jccolor-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jcgray-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jcphuff-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jcsample-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jdcolor-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jdmerge-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jdsample-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jfdctfst-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jfdctint-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jidctfst-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jidctint-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jidctred-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jquanti-neon.c
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
    add_library(libjpeg SHARED ${SOURCE_FILES})

    target_include_directories(libjpeg
            PUBLIC
            ${ANDROID_ROOT}/external/libjpeg-turbo
            ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm
    )

    # Add c/c++ flags
    target_compile_options(libjpeg PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libjpeg
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
    )
endif ()
#################### libjpeg ####################

#################### libjpeg_static_ndk ####################
set(FILES
        ${ANDROID_ROOT}/external/libjpeg-turbo/jaricom.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcapimin.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcapistd.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcarith.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jccoefct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jccolor.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcdctmgr.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jchuff.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcicc.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcinit.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcmainct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcmarker.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcmaster.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcomapi.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcparam.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcphuff.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcprepct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jcsample.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jctrans.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdapimin.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdapistd.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdarith.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdatadst.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdatasrc.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdcoefct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdcolor.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jddctmgr.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdhuff.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdicc.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdinput.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdmainct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdmarker.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdmaster.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdmerge.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdphuff.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdpostct.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdsample.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jdtrans.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jerror.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jfdctflt.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jfdctfst.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jfdctint.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jidctflt.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jidctfst.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jidctint.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jidctred.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jmemmgr.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jmemnobs.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jpeg_nbits_table.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jquant1.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jquant2.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/jutils.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/aarch64/jchuff-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/aarch64/jsimd.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jccolor-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jcgray-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jcphuff-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jcsample-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jdcolor-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jdmerge-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jdsample-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jfdctfst-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jfdctint-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jidctfst-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jidctint-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jidctred-neon.c
        ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm/jquanti-neon.c
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
    add_library(libjpeg_static_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libjpeg_static_ndk
            PUBLIC
            ${ANDROID_ROOT}/external/libjpeg-turbo
            ${ANDROID_ROOT}/external/libjpeg-turbo/simd/arm
    )

    # Add c/c++ flags
    target_compile_options(libjpeg_static_ndk PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libjpeg_static_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libjpeg_static_ndk ####################

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

project(libpng_read_fuzzer)

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

#################### libpng_read_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/external/libpng/contrib/oss-fuzz/libpng_read_fuzzer.cc
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
    add_executable(libpng_read_fuzzer ${SOURCE_FILES})

    target_include_directories(libpng_read_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/external/libpng
    )

    # Add c/c++ flags
    target_compile_options(libpng_read_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(libpng_read_fuzzer
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libm
            libpng
            libz
    )
endif ()
#################### libpng_read_fuzzer ####################

#################### libpng ####################
set(FILES
        ${ANDROID_ROOT}/external/libpng/png.c
        ${ANDROID_ROOT}/external/libpng/pngerror.c
        ${ANDROID_ROOT}/external/libpng/pngget.c
        ${ANDROID_ROOT}/external/libpng/pngmem.c
        ${ANDROID_ROOT}/external/libpng/pngpread.c
        ${ANDROID_ROOT}/external/libpng/pngread.c
        ${ANDROID_ROOT}/external/libpng/pngrio.c
        ${ANDROID_ROOT}/external/libpng/pngrtran.c
        ${ANDROID_ROOT}/external/libpng/pngrutil.c
        ${ANDROID_ROOT}/external/libpng/pngset.c
        ${ANDROID_ROOT}/external/libpng/pngtrans.c
        ${ANDROID_ROOT}/external/libpng/pngwio.c
        ${ANDROID_ROOT}/external/libpng/pngwrite.c
        ${ANDROID_ROOT}/external/libpng/pngwtran.c
        ${ANDROID_ROOT}/external/libpng/pngwutil.c
        ${ANDROID_ROOT}/external/libpng/arm/arm_init.c
        ${ANDROID_ROOT}/external/libpng/arm/filter_neon_intrinsics.c
        ${ANDROID_ROOT}/external/libpng/arm/palette_neon_intrinsics.c
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
    add_library(libpng SHARED ${SOURCE_FILES})

    target_include_directories(libpng
            PUBLIC
            ${ANDROID_ROOT}/external/libpng
    )

    # Add c/c++ flags
    target_compile_options(libpng PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpng
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwinpthread
            libz
    )
endif ()
#################### libpng ####################

#################### libpng_ndk ####################
set(FILES
        ${ANDROID_ROOT}/external/libpng/png.c
        ${ANDROID_ROOT}/external/libpng/pngerror.c
        ${ANDROID_ROOT}/external/libpng/pngget.c
        ${ANDROID_ROOT}/external/libpng/pngmem.c
        ${ANDROID_ROOT}/external/libpng/pngpread.c
        ${ANDROID_ROOT}/external/libpng/pngread.c
        ${ANDROID_ROOT}/external/libpng/pngrio.c
        ${ANDROID_ROOT}/external/libpng/pngrtran.c
        ${ANDROID_ROOT}/external/libpng/pngrutil.c
        ${ANDROID_ROOT}/external/libpng/pngset.c
        ${ANDROID_ROOT}/external/libpng/pngtrans.c
        ${ANDROID_ROOT}/external/libpng/pngwio.c
        ${ANDROID_ROOT}/external/libpng/pngwrite.c
        ${ANDROID_ROOT}/external/libpng/pngwtran.c
        ${ANDROID_ROOT}/external/libpng/pngwutil.c
        ${ANDROID_ROOT}/external/libpng/arm/arm_init.c
        ${ANDROID_ROOT}/external/libpng/arm/filter_neon_intrinsics.c
        ${ANDROID_ROOT}/external/libpng/arm/palette_neon_intrinsics.c
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
    add_library(libpng_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libpng_ndk
            PUBLIC
            ${ANDROID_ROOT}/external/libpng
    )

    # Add c/c++ flags
    target_compile_options(libpng_ndk PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libpng_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libz
    )
endif ()
#################### libpng_ndk ####################

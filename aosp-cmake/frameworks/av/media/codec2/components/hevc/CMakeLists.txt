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

project(libcodec2_soft_hevcenc)

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

#################### libcodec2_soft_hevcenc ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/components/hevc/C2SoftHevcEnc.cpp
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
    add_library(libcodec2_soft_hevcenc SHARED ${SOURCE_FILES})

    target_include_directories(libcodec2_soft_hevcenc
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/components/hevc
    )

    # Add c/c++ flags
    target_compile_options(libcodec2_soft_hevcenc PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libcodec2_soft_hevcenc
            libarect_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcodec2
            libcodec2_soft_common
            libcodec2_vndk
            libcutils
            libdl
            libhevcenc
            liblog
            libm
            libnativewindow_headers
            libsfplugin_ccodec_utils
            libstagefright_foundation
            libutils
    )
endif ()
#################### libcodec2_soft_hevcenc ####################

#################### libcodec2_soft_hevcdec ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/codec2/components/hevc/C2SoftHevcDec.cpp
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
    add_library(libcodec2_soft_hevcdec SHARED ${SOURCE_FILES})

    target_include_directories(libcodec2_soft_hevcdec
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/codec2/components/hevc
    )

    # Add c/c++ flags
    target_compile_options(libcodec2_soft_hevcdec PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libcodec2_soft_hevcdec
            libarect_headers
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcodec2
            libcodec2_soft_common
            libcodec2_vndk
            libcutils
            libdl
            libhevcdec
            liblog
            libm
            libnativewindow_headers
            libsfplugin_ccodec_utils
            libstagefright_foundation
            libutils
    )
endif ()
#################### libcodec2_soft_hevcdec ####################

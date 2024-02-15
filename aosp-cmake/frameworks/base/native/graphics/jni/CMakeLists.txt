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

project(imagedecoder_fuzzer)

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

#################### libjnigraphics ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/native/graphics/jni/imagedecoder.cpp
        ${ANDROID_ROOT}/frameworks/base/native/graphics/jni/aassetstreamadaptor.cpp
        ${ANDROID_ROOT}/frameworks/base/native/graphics/jni/bitmap.cpp
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
    add_library(libjnigraphics SHARED ${SOURCE_FILES})

    target_include_directories(libjnigraphics
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/native/graphics/jni
    )

    # Add c/c++ flags
    target_compile_options(libjnigraphics PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libjnigraphics
            jni_headers
            libandroid
            libarect
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libhwui
            libhwui_internal_headers
            liblog
            libm
            libnativewindow_headers
    )
endif ()
#################### libjnigraphics ####################

#################### imagedecoder_png_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/native/graphics/jni/fuzz/fuzz_imagedecoder.cpp
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
    add_executable(imagedecoder_png_fuzzer ${SOURCE_FILES})

    target_include_directories(imagedecoder_png_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/native/graphics/jni
    )

    # Add c/c++ flags
    target_compile_options(imagedecoder_png_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(imagedecoder_png_fuzzer
            jni_headers
            libarect
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libjnigraphics
            libm
            libutils
            libz
    )
endif ()
#################### imagedecoder_png_fuzzer ####################

#################### imagedecoder_fuzzer ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/native/graphics/jni/fuzz/fuzz_imagedecoder.cpp
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
    add_executable(imagedecoder_fuzzer ${SOURCE_FILES})

    target_include_directories(imagedecoder_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/native/graphics/jni
    )

    # Add c/c++ flags
    target_compile_options(imagedecoder_fuzzer PRIVATE
            -Werror
    )

    target_link_libraries(imagedecoder_fuzzer
            jni_headers
            libarect
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libdl
            libjnigraphics
            libm
            libutils
    )
endif ()
#################### imagedecoder_fuzzer ####################

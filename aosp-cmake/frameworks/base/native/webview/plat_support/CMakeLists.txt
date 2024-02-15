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

project(libwebviewchromium_plat_support)

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

#################### libwebviewchromium_plat_support ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/native/webview/plat_support/draw_functor.cpp
        ${ANDROID_ROOT}/frameworks/base/native/webview/plat_support/draw_gl_functor.cpp
        ${ANDROID_ROOT}/frameworks/base/native/webview/plat_support/functor_utils.cpp
        ${ANDROID_ROOT}/frameworks/base/native/webview/plat_support/jni_entry_point.cpp
        ${ANDROID_ROOT}/frameworks/base/native/webview/plat_support/graphics_utils.cpp
        ${ANDROID_ROOT}/frameworks/base/native/webview/plat_support/graphic_buffer_impl.cpp
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
    add_library(libwebviewchromium_plat_support SHARED ${SOURCE_FILES})

    target_include_directories(libwebviewchromium_plat_support
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/native/webview/plat_support
    )

    # Add c/c++ flags
    target_compile_options(libwebviewchromium_plat_support PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libwebviewchromium_plat_support
            libandroidfw
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libhwui
            libhwui_internal_headers
            liblog
            libm
            libnativehelper
            libui
            libutils
            libvulkan
    )
endif ()
#################### libwebviewchromium_plat_support ####################

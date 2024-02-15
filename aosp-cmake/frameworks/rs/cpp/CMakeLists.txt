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

project(libRSDispatch)

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

#################### libRSDispatch ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/cpp/rsDispatch.cpp
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
    add_library(libRSDispatch SHARED ${SOURCE_FILES})

    target_include_directories(libRSDispatch
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs/cpp
    )

    # Add c/c++ flags
    target_compile_options(libRSDispatch PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libRSDispatch
            jni_headers
            libc
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libRSDispatch ####################

#################### libRScpp_static ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/cpp/RenderScript.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/BaseObj.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Element.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Type.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Allocation.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Script.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/ScriptC.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/ScriptIntrinsics.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/ScriptIntrinsicBLAS.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Sampler.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/rsCppUtils.cpp
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
    add_library(libRScpp_static SHARED ${SOURCE_FILES})

    target_include_directories(libRScpp_static
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs/cpp
    )

    # Add c/c++ flags
    target_compile_options(libRScpp_static PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libRScpp_static
            jni_headers
            libRSDispatch
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
            rs-headers
    )
endif ()
#################### libRScpp_static ####################

#################### libRScpp ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/rs/cpp/RenderScript.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/BaseObj.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Element.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Type.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Allocation.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Script.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/ScriptC.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/ScriptIntrinsics.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/ScriptIntrinsicBLAS.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/Sampler.cpp
        ${ANDROID_ROOT}/frameworks/rs/cpp/rsCppUtils.cpp
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
    add_library(libRScpp SHARED ${SOURCE_FILES})

    target_include_directories(libRScpp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/rs/cpp
    )

    # Add c/c++ flags
    target_compile_options(libRScpp PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libRScpp
            jni_headers
            libRSDispatch
            libarect_headers
            libbase_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libgui
            liblog
            libm
            libnativebase_headers
            libnativewindow_headers
            libutils
            rs-headers
    )
endif ()
#################### libRScpp ####################

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

project(libstagefright_omx_utils)

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

#################### libstagefright_omx_utils ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/OMXUtils.cpp
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
    add_library(libstagefright_omx_utils SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_omx_utils
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_omx_utils PRIVATE
            -Wall
            -Werror
    )

    target_link_libraries(libstagefright_omx_utils
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libmedia_omx
            libstagefright_foundation
            media_plugin_headers
    )
endif ()
#################### libstagefright_omx_utils ####################

#################### libstagefright_omx ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/OMXStore.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/OMXNodeInstance.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/OMXUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/OmxGraphicBufferSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/Omx.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/OmxStore.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/WGraphicBufferSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/WOmxNode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/WOmxObserver.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/WOmxBufferSource.cpp
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
    add_library(libstagefright_omx SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_omx
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_omx PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libstagefright_omx
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.media.omx-1.0
            libbase
            libbase_headers
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_standalone
            libcutils
            libdl
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_omx
            libstagefright_bufferqueue_helper
            libstagefright_foundation
            libstagefright_xmlparser
            libui
            libutils
            libvndksupport
            media_plugin_headers
    )
endif ()
#################### libstagefright_omx ####################

#################### libstagefright_omx_ext ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/OMXStore.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/OMXNodeInstance.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/OMXUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/OmxGraphicBufferSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/Omx.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/OmxStore.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/WGraphicBufferSource.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/WOmxNode.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/WOmxObserver.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/1.0/WOmxBufferSource.cpp
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
    add_library(libstagefright_omx_ext SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_omx_ext
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_omx_ext PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libstagefright_omx_ext
            android.hardware.graphics.bufferqueue-1.0
            android.hardware.media.omx-1.0
            libbase
            libbase_headers
            libbinder
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            libhidlbase
            libhidlmemory
            liblog
            libm
            libmedia_omx
            libstagefright_bufferqueue_helper
            libstagefright_foundation_ext
            libstagefright_xmlparser
            libui
            libutils
            libvndksupport
            media_plugin_headers
            libc
            libdl
            libm
    )
endif ()
#################### libstagefright_omx_ext ####################

#################### libstagefright_softomx ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/SimpleSoftOMXComponent.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/SoftOMXComponent.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/SoftVideoDecoderOMXComponent.cpp
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/SoftVideoEncoderOMXComponent.cpp
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
    add_library(libstagefright_softomx SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_softomx
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_softomx PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libstagefright_softomx
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libstagefright_foundation
            libui
            libutils
            media_plugin_headers
    )
endif ()
#################### libstagefright_softomx ####################

#################### libstagefright_softomx_plugin ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/SoftOMXPlugin.cpp
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
    add_library(libstagefright_softomx_plugin SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_softomx_plugin
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx/include
            ${ANDROID_ROOT}/frameworks/av/media/libstagefright/omx
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_softomx_plugin PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libstagefright_softomx_plugin
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libstagefright_foundation
            libstagefright_softomx
            libutils
            media_plugin_headers
    )
endif ()
#################### libstagefright_softomx_plugin ####################

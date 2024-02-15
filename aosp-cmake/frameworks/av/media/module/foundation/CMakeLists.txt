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

project(libstagefright_foundation)

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

#################### libstagefright_foundation ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AAtomizer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ABitReader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ABuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ADebug.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AHandler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ALooper.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ALooperRoster.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AMessage.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AString.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AStringUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AudioPresentationInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ByteUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ColorUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ColorUtils_fill.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/FoundationUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaBufferBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaBufferGroup.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaDefs.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaKeys.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MetaData.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MetaDataBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/OpusHeader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/avc_utils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/base64.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/hexdump.cpp
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
    add_library(libstagefright_foundation SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_foundation
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation/include/media/stagefright/foundation
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation/include
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_foundation PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_foundation
            av-headers
            libbinder
            libbinder_headers
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libstagefright_foundation_headers
            libutils
            media_ndk_headers
            media_plugin_headers
    )
endif ()
#################### libstagefright_foundation ####################

#################### libstagefright_foundation_without_imemory ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AAtomizer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ABitReader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ABuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ADebug.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AHandler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ALooper.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ALooperRoster.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AMessage.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AString.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AStringUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AudioPresentationInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ByteUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ColorUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ColorUtils_fill.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/FoundationUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaBufferBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaBufferGroup.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaDefs.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaKeys.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MetaData.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MetaDataBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/OpusHeader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/avc_utils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/base64.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/hexdump.cpp
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
    add_library(libstagefright_foundation_without_imemory SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_foundation_without_imemory
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation/include/media/stagefright/foundation
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation/include
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_foundation_without_imemory PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_foundation_without_imemory
            av-headers
            libbinder
            libbinder_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libstagefright_foundation_headers
            libutils
            media_ndk_headers
            media_plugin_headers
    )
endif ()
#################### libstagefright_foundation_without_imemory ####################

#################### libstagefright_foundation_ext ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AAtomizer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ABitReader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ABuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ADebug.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AHandler.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ALooper.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ALooperRoster.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AMessage.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AString.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AStringUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/AudioPresentationInfo.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ByteUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ColorUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ColorUtils_fill.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/FoundationUtils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaBuffer.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaBufferBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaBufferGroup.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaDefs.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MediaKeys.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MetaData.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/MetaDataBase.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/OpusHeader.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/avc_utils.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/base64.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/hexdump.cpp
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
    add_library(libstagefright_foundation_ext SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_foundation_ext
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation/include/media/stagefright/foundation
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation/include
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_foundation_ext PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_foundation_ext
            av-headers
            libbinder_headers
            libc++
            libc++_static
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdl
            liblog
            libm
            libstagefright_foundation_headers
            libutils
            media_ndk_headers
            media_plugin_headers
            libc
            libdl
            libm
    )
endif ()
#################### libstagefright_foundation_ext ####################

#################### libstagefright_foundation_colorutils_ndk ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ColorUtils_ndk.cpp
        ${ANDROID_ROOT}/frameworks/av/media/module/foundation/ColorUtils_fill.cpp
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
    add_library(libstagefright_foundation_colorutils_ndk SHARED ${SOURCE_FILES})

    target_include_directories(libstagefright_foundation_colorutils_ndk
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation/include/media/stagefright/foundation
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation
            ${ANDROID_ROOT}/frameworks/av/media/module/foundation/include
    )

    # Add c/c++ flags
    target_compile_options(libstagefright_foundation_colorutils_ndk PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libstagefright_foundation_colorutils_ndk
            libbinder_headers
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libdl
            liblog
            libm
            libstagefright_foundation_headers
            libutils
            media_ndk_headers
            media_plugin_headers
    )
endif ()
#################### libstagefright_foundation_colorutils_ndk ####################

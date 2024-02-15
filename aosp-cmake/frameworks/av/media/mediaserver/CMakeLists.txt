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

project(mediaserver64)

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

#################### mediaserverwrapper ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/mediaserver/mediaserverwrapper.cpp
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
    add_executable(mediaserverwrapper ${SOURCE_FILES})

    target_include_directories(mediaserverwrapper
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/mediaserver
    )

    # Add c/c++ flags
    target_compile_options(mediaserverwrapper PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(mediaserverwrapper
            android.hardware.media.omx-1.0
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libfmq
            libhidlbase
            libicu
            liblog
            libm
            libmediaplayerservice
            libregistermsext
            libresourcemanagerservice
            libutils
    )
endif ()
#################### mediaserverwrapper ####################

#################### libregistermsext ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/mediaserver/register.cpp
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
    add_library(libregistermsext SHARED ${SOURCE_FILES})

    target_include_directories(libregistermsext
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/mediaserver
    )

    # Add c/c++ flags
    target_compile_options(libregistermsext PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libregistermsext
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
    )
endif ()
#################### libregistermsext ####################

#################### mediaserver64 ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/av/media/mediaserver/main_mediaserver.cpp
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
    add_executable(mediaserver64 ${SOURCE_FILES})

    target_include_directories(mediaserver64
            PUBLIC
            ${ANDROID_ROOT}/frameworks/av/media/mediaserver
    )

    # Add c/c++ flags
    target_compile_options(mediaserver64 PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(mediaserver64
            android.hardware.media.omx-1.0
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libfmq
            libhidlbase
            libicu
            liblog
            libm
            libmediaplayerservice
            libregistermsext
            libresourcemanagerservice
            libutils
    )
endif ()
#################### mediaserver64 ####################

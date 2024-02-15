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

project(libaudioroute)

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

#################### libaudioroute ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_route/audio_route.c
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
    add_library(libaudioroute SHARED ${SOURCE_FILES})

    target_include_directories(libaudioroute
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_route/include
            ${ANDROID_ROOT}/system/media/audio_route
    )

    # Add c/c++ flags
    target_compile_options(libaudioroute PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudioroute
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libexpat
            liblog
            libm
            libtinyalsa
            libutils
    )
endif ()
#################### libaudioroute ####################

#################### libaudioroutev2 ####################
set(FILES
        ${ANDROID_ROOT}/system/media/audio_route/audio_route.c
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
    add_library(libaudioroutev2 SHARED ${SOURCE_FILES})

    target_include_directories(libaudioroutev2
            PUBLIC
            ${ANDROID_ROOT}/system/media/audio_route/include
            ${ANDROID_ROOT}/system/media/audio_route
    )

    # Add c/c++ flags
    target_compile_options(libaudioroutev2 PRIVATE
            -Werror
            -Wall
    )

    target_link_libraries(libaudioroutev2
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libexpat
            liblog
            libm
            libtinyalsav2
            libutils
    )
endif ()
#################### libaudioroutev2 ####################

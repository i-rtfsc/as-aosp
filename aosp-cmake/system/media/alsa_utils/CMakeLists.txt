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

project(libalsautilsv2)

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

#################### libalsautilsv2 ####################
set(FILES
        ${ANDROID_ROOT}/system/media/alsa_utils/alsa_device_profile.c
        ${ANDROID_ROOT}/system/media/alsa_utils/alsa_device_proxy.c
        ${ANDROID_ROOT}/system/media/alsa_utils/alsa_logging.c
        ${ANDROID_ROOT}/system/media/alsa_utils/alsa_format.c
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
    add_library(libalsautilsv2 SHARED ${SOURCE_FILES})

    target_include_directories(libalsautilsv2
            PUBLIC
            ${ANDROID_ROOT}/system/media/alsa_utils/include
            ${ANDROID_ROOT}/system/media/alsa_utils
    )

    # Add c/c++ flags
    target_compile_options(libalsautilsv2 PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libalsautilsv2
            libaudio_system_headers
            libaudioutils
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libtinyalsav2
            libc
            libdl
            libm
    )
endif ()
#################### libalsautilsv2 ####################

#################### libalsautils ####################
set(FILES
        ${ANDROID_ROOT}/system/media/alsa_utils/alsa_device_profile.c
        ${ANDROID_ROOT}/system/media/alsa_utils/alsa_device_proxy.c
        ${ANDROID_ROOT}/system/media/alsa_utils/alsa_logging.c
        ${ANDROID_ROOT}/system/media/alsa_utils/alsa_format.c
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
    add_library(libalsautils SHARED ${SOURCE_FILES})

    target_include_directories(libalsautils
            PUBLIC
            ${ANDROID_ROOT}/system/media/alsa_utils/include
            ${ANDROID_ROOT}/system/media/alsa_utils
    )

    # Add c/c++ flags
    target_compile_options(libalsautils PRIVATE
            -Werror
            -Wall
            -Wno-unused-parameter
    )

    target_link_libraries(libalsautils
            libaudio_system_headers
            libaudioutils
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libtinyalsa
            libc
            libdl
            libm
    )
endif ()
#################### libalsautils ####################

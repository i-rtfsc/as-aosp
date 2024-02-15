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

project(bootanimation)

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

#################### bootanimation ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/cmds/bootanimation/BootAnimationUtil.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/bootanimation/bootanimation_main.cpp
        ${ANDROID_ROOT}/frameworks/base/cmds/bootanimation/audioplay.cpp
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
    add_executable(bootanimation ${SOURCE_FILES})

    target_include_directories(bootanimation
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/bootanimation
    )

    # Add c/c++ flags
    target_compile_options(bootanimation PRIVATE
            -DFLYME_GRAPHICS_EXTEND
            -Wall
            -Werror
    )

    target_link_libraries(bootanimation
            jni_headers
            libOpenSLES
            libandroidfw
            libbase
            libbinder
            libbootanimation
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            liblog
            libm
            libutils
    )
endif ()
#################### bootanimation ####################

#################### libbootanimation ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/base/cmds/bootanimation/BootAnimation.cpp
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
    add_library(libbootanimation SHARED ${SOURCE_FILES})

    target_include_directories(libbootanimation
            PUBLIC
            ${ANDROID_ROOT}/frameworks/base/cmds/bootanimation
    )

    # Add c/c++ flags
    target_compile_options(libbootanimation PRIVATE
            -DFLYME_GRAPHICS_EXTEND
            -Wall
            -Werror
    )

    target_link_libraries(libbootanimation
            libEGL
            libGLESv2
            libandroidfw
            libbase
            libbinder
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libgui
            libjnigraphics
            liblog
            libm
            libui
            libutils
    )
endif ()
#################### libbootanimation ####################

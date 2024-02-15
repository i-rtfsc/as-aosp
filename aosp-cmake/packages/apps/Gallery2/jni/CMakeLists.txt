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

project(libjni_filtershow_filters)

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

#################### libjni_filtershow_filters ####################
set(FILES
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/gradient.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/saturated.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/exposure.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/edge.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/contrast.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/hue.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/shadows.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/highlight.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/hsv.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/vibrance.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/geometry.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/negative.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/redEyeMath.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/fx.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/wbalance.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/redeye.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/bwfilter.c
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/tinyplanet.cc
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/filters/kmeans.cc
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
    add_library(libjni_filtershow_filters SHARED ${SOURCE_FILES})

    target_include_directories(libjni_filtershow_filters
            PUBLIC
            ${ANDROID_ROOT}/packages/apps/Gallery2/jni
    )

    # Add c/c++ flags
    target_compile_options(libjni_filtershow_filters PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libjni_filtershow_filters
            jni_headers
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libjnigraphics
            liblog
            libm
    )
endif ()
#################### libjni_filtershow_filters ####################

#################### libjni_eglfence ####################
set(FILES
        ${ANDROID_ROOT}/packages/apps/Gallery2/jni/jni_egl_fence.cpp
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
    add_library(libjni_eglfence SHARED ${SOURCE_FILES})

    target_include_directories(libjni_eglfence
            PUBLIC
            ${ANDROID_ROOT}/packages/apps/Gallery2/jni
    )

    # Add c/c++ flags
    target_compile_options(libjni_eglfence PRIVATE
            -Wall
            -Wextra
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libjni_eglfence
            jni_headers
            libEGL
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            liblog
            libm
    )
endif ()
#################### libjni_eglfence ####################

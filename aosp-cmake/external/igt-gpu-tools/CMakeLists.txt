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

project(igt-gpu-tools)

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

#################### igt-gpu-tools ####################
set(FILES
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/drmtest.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/gem.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/gem_msm.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_aux.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_color_encoding.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_core.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_debugfs.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_device.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_draw.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_dummy.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_dummyload.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_fb.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_kmod.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_kms.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_pm.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_rand.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_stats.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_sysfs.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/igt_syncobj.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/ion.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/ioctl_wrappers.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/sw_sync.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/i915/gem_mman.c
        ${ANDROID_ROOT}/external/igt-gpu-tools/lib/uwildmat/uwildmat.c
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
    add_library(igt-gpu-tools SHARED ${SOURCE_FILES})

    target_include_directories(igt-gpu-tools
            PUBLIC
            ${ANDROID_ROOT}/external/igt-gpu-tools/include
            ${ANDROID_ROOT}/external/igt-gpu-tools/lib
            ${ANDROID_ROOT}/external/igt-gpu-tools/lib/stubs/drm
            ${ANDROID_ROOT}/external/igt-gpu-tools/prebuilt-intermediates
            ${ANDROID_ROOT}/external/igt-gpu-tools
    )

    # Add c/c++ flags
    target_compile_options(igt-gpu-tools PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(igt-gpu-tools
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libdrm
            libelf
            libion
            libkmod
            liblog
            libm
    )
endif ()
#################### igt-gpu-tools ####################

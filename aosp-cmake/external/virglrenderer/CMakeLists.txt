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

project(libvirglrenderer)

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

#################### libvirglrenderer ####################
set(FILES
        ${ANDROID_ROOT}/external/virglrenderer/prebuilt-intermediates/src/u_format_table.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/cso_cache/cso_cache.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/cso_cache/cso_hash.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_build.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_dump.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_info.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_iterate.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_parse.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_sanity.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_scan.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_strings.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_text.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/tgsi/tgsi_util.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/util/u_debug_describe.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/util/u_format.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/util/u_hash_table.c
        ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/util/u_texture.c
        ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util/anon_file.c
        ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util/bitscan.c
        ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util/hash_table.c
        ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util/os_file.c
        ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util/os_misc.c
        ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util/ralloc.c
        ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util/u_cpu_detect.c
        ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util/u_debug.c
        ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util/u_math.c
        ${ANDROID_ROOT}/external/virglrenderer/src/iov.c
        ${ANDROID_ROOT}/external/virglrenderer/src/virgl_context.c
        ${ANDROID_ROOT}/external/virglrenderer/src/virglrenderer.c
        ${ANDROID_ROOT}/external/virglrenderer/src/virgl_resource.c
        ${ANDROID_ROOT}/external/virglrenderer/src/virgl_util.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_blitter.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_debug.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_decode.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_formats.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_object.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_renderer.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_shader.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_tweaks.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_winsys.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_winsys_egl.c
        ${ANDROID_ROOT}/external/virglrenderer/src/vrend_winsys_gbm.c
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
    add_library(libvirglrenderer SHARED ${SOURCE_FILES})

    target_include_directories(libvirglrenderer
            PUBLIC
            ${ANDROID_ROOT}/external/virglrenderer/prebuilt-intermediates
            ${ANDROID_ROOT}/external/virglrenderer/src
            ${ANDROID_ROOT}/external/virglrenderer/src/drm
            ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary
            ${ANDROID_ROOT}/external/virglrenderer/src/gallium/auxiliary/util
            ${ANDROID_ROOT}/external/virglrenderer/src/gallium/include
            ${ANDROID_ROOT}/external/virglrenderer/src/mesa
            ${ANDROID_ROOT}/external/virglrenderer/src/mesa/compat
            ${ANDROID_ROOT}/external/virglrenderer/src/mesa/pipe
            ${ANDROID_ROOT}/external/virglrenderer/src/mesa/util
            ${ANDROID_ROOT}/external/virglrenderer/src/venus
            ${ANDROID_ROOT}/external/virglrenderer
    )

    # Add c/c++ flags
    target_compile_options(libvirglrenderer PRIVATE
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(libvirglrenderer
            libc
            libc++
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libcutils
            libdl
            libdrm
            libdrm_headers
            libepoxy
            libgbm
            liblog
            libm
    )
endif ()
#################### libvirglrenderer ####################

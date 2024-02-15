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

project(libwayland_client_static)

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

#################### libwayland_client_static ####################
set(FILES
        ${ANDROID_ROOT}/external/wayland/src/connection.c
        ${ANDROID_ROOT}/external/wayland/src/wayland-client.c
        ${ANDROID_ROOT}/external/wayland/src/wayland-os.c
        ${ANDROID_ROOT}/external/wayland/src/wayland-util.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland/wayland_core_protocol_sources_static/gen/wayland-protocol-private.c
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
    add_library(libwayland_client_static SHARED ${SOURCE_FILES})

    target_include_directories(libwayland_client_static
            PUBLIC
            ${ANDROID_ROOT}/external/wayland/src
            ${ANDROID_ROOT}/external/wayland/egl
            ${ANDROID_ROOT}/external/wayland
            ${OUT_INTERMEDIATES_ROOT}/external/wayland/wayland_core_protocol_sources_static/gen
            ${OUT_INTERMEDIATES_ROOT}/external/wayland/wayland_core_client_protocol_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(libwayland_client_static PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libwayland_client_static
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libffi
            libm
    )
endif ()
#################### libwayland_client_static ####################

#################### libwayland_server ####################
set(FILES
        ${ANDROID_ROOT}/external/wayland/src/connection.c
        ${ANDROID_ROOT}/external/wayland/src/wayland-os.c
        ${ANDROID_ROOT}/external/wayland/src/wayland-util.c
        ${ANDROID_ROOT}/external/wayland/src/wayland-server.c
        ${ANDROID_ROOT}/external/wayland/src/wayland-shm.c
        ${ANDROID_ROOT}/external/wayland/src/event-loop.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland/wayland_core_protocol_sources_static/gen/wayland-protocol-private.c
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
    add_library(libwayland_server SHARED ${SOURCE_FILES})

    target_include_directories(libwayland_server
            PUBLIC
            ${ANDROID_ROOT}/external/wayland/src
            ${ANDROID_ROOT}/external/wayland
            ${OUT_INTERMEDIATES_ROOT}/external/wayland/wayland_core_protocol_sources_static/gen
            ${OUT_INTERMEDIATES_ROOT}/external/wayland/wayland_core_server_protocol_headers/gen
    )

    # Add c/c++ flags
    target_compile_options(libwayland_server PRIVATE
            -Werror
            -Wall
            -Wextra
            -Wno-unused-parameter
            -Wno-unused-variable
    )

    target_link_libraries(libwayland_server
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libffi
            libm
    )
endif ()
#################### libwayland_server ####################

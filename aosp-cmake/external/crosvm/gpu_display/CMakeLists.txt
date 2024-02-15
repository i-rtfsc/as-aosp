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

project(libwayland_crosvm_gpu_display_extension_server_protocols)

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

#################### libwayland_crosvm_gpu_display_extension_server_protocols ####################
set(FILES
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/presentation-time.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/viewporter.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-shell.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/fullscreen-shell-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/idle-inhibit-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/input-method-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/input-timestamps-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/keyboard-shortcuts-inhibit-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/linux-dmabuf-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/linux-explicit-synchronization-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/pointer-constraints-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/pointer-gestures-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/primary-selection-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/relative-pointer-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/tablet-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/tablet-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/text-input-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-decoration-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-foreign-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-foreign-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-output-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-shell-unstable-v6.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xwayland-keyboard-grab-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/text-input-unstable-v3.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/alpha-compositing-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/aura-shell.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/color-space-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/cursor-shapes-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/extended-drag-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/gaming-input-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/keyboard-configuration-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/keyboard-extension-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/notification-shell-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/remote-shell-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/remote-shell-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/secure-output-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/stylus-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/stylus-tools-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/text-input-extension-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/touchpad-haptics-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/vsync-feedback-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/virtio-gpu-metadata-v1.c
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
    add_library(libwayland_crosvm_gpu_display_extension_server_protocols SHARED ${SOURCE_FILES})

    target_include_directories(libwayland_crosvm_gpu_display_extension_server_protocols
            PUBLIC
            ${ANDROID_ROOT}/external/crosvm/gpu_display
            ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen
            ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_server_protocol_headers/gen/wayland_protocol_codegen
    )

    # Add c/c++ flags
    target_compile_options(libwayland_crosvm_gpu_display_extension_server_protocols PRIVATE
            -Wall
            -Wextra
            -Werror
    )

    target_link_libraries(libwayland_crosvm_gpu_display_extension_server_protocols
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libm
            libwayland_server
    )
endif ()
#################### libwayland_crosvm_gpu_display_extension_server_protocols ####################

#################### libdisplay_wl ####################
set(FILES
        ${ANDROID_ROOT}/external/crosvm/gpu_display/src/display_wl.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/presentation-time.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/viewporter.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-shell.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/fullscreen-shell-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/idle-inhibit-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/input-method-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/input-timestamps-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/keyboard-shortcuts-inhibit-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/linux-dmabuf-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/linux-explicit-synchronization-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/pointer-constraints-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/pointer-gestures-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/primary-selection-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/relative-pointer-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/tablet-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/tablet-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/text-input-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-decoration-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-foreign-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-foreign-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-output-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xdg-shell-unstable-v6.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/xwayland-keyboard-grab-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/text-input-unstable-v3.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/alpha-compositing-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/aura-shell.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/color-space-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/cursor-shapes-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/extended-drag-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/gaming-input-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/keyboard-configuration-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/keyboard-extension-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/notification-shell-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/remote-shell-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/remote-shell-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/secure-output-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/stylus-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/stylus-tools-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/text-input-extension-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/touchpad-haptics-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/vsync-feedback-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen/virtio-gpu-metadata-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/presentation-time.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/viewporter.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/xdg-shell.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/fullscreen-shell-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/idle-inhibit-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/input-method-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/input-timestamps-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/keyboard-shortcuts-inhibit-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/linux-dmabuf-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/linux-explicit-synchronization-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/pointer-constraints-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/pointer-gestures-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/primary-selection-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/relative-pointer-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/tablet-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/tablet-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/text-input-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/xdg-decoration-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/xdg-foreign-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/xdg-foreign-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/xdg-output-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/xdg-shell-unstable-v6.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/xwayland-keyboard-grab-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/text-input-unstable-v3.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/alpha-compositing-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/aura-shell.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/color-space-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/cursor-shapes-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/extended-drag-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/gaming-input-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/keyboard-configuration-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/keyboard-extension-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/notification-shell-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/remote-shell-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/remote-shell-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/secure-output-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/stylus-unstable-v2.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/stylus-tools-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/text-input-extension-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/touchpad-haptics-unstable-v1.c
        ${OUT_INTERMEDIATES_ROOT}/external/wayland-protocols/wayland_extension_protocol_sources/gen/wayland_protocol_codegen/vsync-feedback-unstable-v1.c
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
    add_library(libdisplay_wl SHARED ${SOURCE_FILES})

    target_include_directories(libdisplay_wl
            PUBLIC
            ${ANDROID_ROOT}/external/crosvm/gpu_display
            ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_protocol_sources/gen/wayland_protocol_codegen
            ${OUT_INTERMEDIATES_ROOT}/external/crosvm/gpu_display/gpu_display_client_protocol_headers/gen/wayland_protocol_codegen
    )

    # Add c/c++ flags
    target_compile_options(libdisplay_wl PRIVATE
            -Werror
    )

    target_link_libraries(libdisplay_wl
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libdl
            libffi
            libm
            libwayland_client
            libwayland_client_static
    )
endif ()
#################### libdisplay_wl ####################

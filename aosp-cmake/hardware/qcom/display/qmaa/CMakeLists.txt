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

project(vendor.qti.hardware.display.composer-service)

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

#################### vendor.qti.hardware.display.composer-service ####################
set(FILES
        ${ANDROID_ROOT}/hardware/qcom/display/composer/AidlComposer.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/AidlComposerClient.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/AidlComposerHandleImporter.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/DisplayConfigAIDL.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/QtiComposer3Client.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/cpuhint.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/gl_color_convert.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/gl_color_convert_impl.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/gl_common.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/gl_layer_stitch.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/gl_layer_stitch_impl.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_buffer_allocator.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_buffer_sync_handler.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_callbacks.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_color_manager.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_color_mode_stc.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_common.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_debugger.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_display.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_display_builtin.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_display_pluggable.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_display_pluggable_test.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_display_virtual.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_display_virtual_dpu.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_display_virtual_factory.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_display_virtual_gpu.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_layers.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_session.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_session_services.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_socket_handler.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/hwc_tonemapper.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/ipc_impl.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/composer/service.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/qmaa/QtiQmaaComposer.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/qmaa/QtiQmaaComposerClient.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/qmaa/QtiQmaaComposerHandleImporter.cpp
        ${ANDROID_ROOT}/hardware/qcom/display/qmaa/service.cpp
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
    add_executable(vendor.qti.hardware.display.composer-service ${SOURCE_FILES})

    target_include_directories(vendor.qti.hardware.display.composer-service
            PUBLIC
            ${ANDROID_ROOT}/hardware/qcom/display/qmaa
    )

    # Add c/c++ flags
    target_compile_options(vendor.qti.hardware.display.composer-service PRIVATE
            -Wall
            -Werror
            -Wno-unused-parameter
    )

    target_link_libraries(vendor.qti.hardware.display.composer-service
            android.hardware.common-V2-ndk
            android.hardware.graphics.allocator-V2-ndk
            android.hardware.graphics.composer3-V2-ndk
            android.hardware.graphics.mapper-4.0
            device_kernel_headers
            display_headers
            libEGL
            libGLESv2
            libGLESv3
            libaidlcommonsupport
            libbinder
            libbinder_ndk
            libc++
            libc++demangle
            libc
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdisplayconfig.qti
            libdisplaydebug
            libdl
            libdrm
            libfmq
            libgpu_tonemapper
            libgralloc.qti
            libgralloctypes
            libgrallocutils
            libhardware
            libhardware_legacy
            libhidlbase
            libhistogram
            liblog
            libm
            libqservice
            libsdmcore
            libsdmutils
            libsync
            libui
            libutils
            qti_display_kernel_headers
            qti_kernel_headers
            vendor.qti.hardware.display.composer3-V1-ndk
            vendor.qti.hardware.display.config-V8-ndk
            vendor.qti.hardware.display.demura-V1-ndk
            vendor.qti.hardware.display.mapper-4.0
            vendor.qti.hardware.display.mapperextensions-1.1
            vendor.qti.hardware.display.mapperextensions-1.2
            vendor.qti.hardware.display.mapperextensions-1.3
            libc
            libdl
            libm
    )
endif ()
#################### vendor.qti.hardware.display.composer-service ####################

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

project(libqtisurfaceflingerextension)

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

#################### libqtisurfaceflingerextension ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiNullExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiNullDisplaySurfaceExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiSurfaceFlingerExtensionFactory.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiAidlComposerHalExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiFeatureManager.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiDolphinWrapper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiFramebufferSurfaceExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiHidlComposerHalExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiHWComposerExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiPhaseOffsetsExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiPowerAdvisorExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiSurfaceFlingerExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiVirtualDisplaySurfaceExtension.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension/QtiWorkDurationsExtension.cpp
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
    add_library(libqtisurfaceflingerextension SHARED ${SOURCE_FILES})

    target_include_directories(libqtisurfaceflingerextension
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension
    )

    # Add c/c++ flags
    target_compile_options(libqtisurfaceflingerextension PRIVATE
            -DFLYME_GRAPHICS_EXTEND
            -DFLYME_SUPPORT_LTPO_V2
            -Wall
            -Werror
            -Wextra
            -DQTI_DISPLAY_EXTENSION
            -Wno-unused-parameter
            -Wno-unused-parameter
            -DQTI_COMPOSER3_EXTENSIONS
    )

    target_link_libraries(libqtisurfaceflingerextension
            android.hardware.common-V2-ndk
            android.hardware.common.fmq-V1-ndk
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            android.hardware.configstore-1.1
            android.hardware.graphics.allocator-2.0
            android.hardware.graphics.allocator-3.0
            android.hardware.graphics.common-1.2
            android.hardware.graphics.composer3-V2-ndk
            android.hardware.graphics.composer3-command-buffer
            android.hardware.graphics.composer-2.1
            android.hardware.graphics.composer-2.1-command-buffer
            android.hardware.graphics.composer-2.2
            android.hardware.graphics.composer-2.2-command-buffer
            android.hardware.graphics.composer-2.3
            android.hardware.graphics.composer-2.3-command-buffer
            android.hardware.graphics.composer-2.4
            android.hardware.graphics.composer-2.4-command-buffer
            android.hardware.power-V4-cpp
            display_intf_headers
            libEGL
            libGLESv1_CM
            libGLESv2
            libSurfaceFlingerProp
            libaidlcommonsupport
            libarect
            libbase
            libbinder
            libbinder_ndk
            libc
            libc++_static
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcompositionengine
            libcutils
            libdisplayconfig.system.qti
            libdl
            libem
            libfmq
            libframetimeline
            libgui
            libgui_aidl_static
            libhidlbase
            libinput
            liblayers_proto
            liblog
            libm
            libnativewindow
            libperfetto_client_experimental
            libpng
            libpowermanager
            libprocessgroup
            libprotobuf-cpp-lite
            librenderengine
            libruntime_log
            libscheduler
            libserviceutils
            libshaders
            libskia_renderengine
            libsurfaceflinger_headers
            libsync
            libtiff
            libtimestats
            libtonemap
            libui
            libutils
            libvulkan
            libz
            server_configurable_flags
            vendor.qti.hardware.display.composer3-V1-ndk
            vendor.qti.hardware.display.composer-3.0
            vendor.qti.hardware.display.composer-3.1
            vendor.qti.hardware.display.config-V7-ndk
    )
endif ()
#################### libqtisurfaceflingerextension ####################

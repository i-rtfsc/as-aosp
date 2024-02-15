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

project(libcompositionengine)

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

#################### libcompositionengine ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/planner/CachedSet.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/planner/Flattener.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/planner/LayerState.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/planner/Planner.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/planner/Predictor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/planner/TexturePool.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/ClientCompositionRequestCache.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/CompositionEngine.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/Display.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/DisplayColorProfile.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/DisplaySurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/DumpHelpers.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/HwcAsyncWorker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/HwcBufferCache.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/LayerFECompositionState.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/Output.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/OutputCompositionState.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/OutputLayer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/OutputLayerCompositionState.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/src/RenderSurface.cpp
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
    add_library(libcompositionengine SHARED ${SOURCE_FILES})

    target_include_directories(libcompositionengine
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/include
    )

    # Add c/c++ flags
    target_compile_options(libcompositionengine PRIVATE
            -DQTI_DISPLAY_EXTENSION
            -Wno-unused-parameter
            -Wno-unused-parameter
            -DQTI_COMPOSER3_EXTENSIONS
            -DFLYME_GRAPHICS_EXTEND
            -DFLYME_COLOR_MANAGEMENT
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libcompositionengine
            android.hardware.graphics.allocator-2.0
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
            android.hardware.power-1.0
            android.hardware.power-1.3
            display_intf_headers
            libEGL
            libGLESv2
            libSurfaceFlingerProp
            libaidlcommonsupport
            libarect
            libbase
            libbinder_ndk
            libc
            libc++
            libc++_static
            libc++demangle
            libcgrouprc
            libcgrouprc_format
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcutils
            libdisplayconfig.system.qti
            libdl
            libgui
            libjsoncpp
            liblayers_proto
            liblog
            libm
            libmath
            libnativewindow
            libperfetto_client_experimental
            libpng
            libprocessgroup
            libprotobuf-cpp-lite
            libqticompositionengineextension
            librenderengine
            libruntime_log
            libskia_renderengine
            libsurfaceflinger_headers
            libtiff
            libtimestats
            libtonemap
            libui
            libutils
            libvulkan
            libz
            vendor.qti.hardware.display.composer3-V1-ndk
            vendor.qti.hardware.display.composer-3.0
            vendor.qti.hardware.display.composer-3.1
            vendor.qti.hardware.display.config-V7-ndk
    )
endif ()
#################### libcompositionengine ####################

#################### libcompositionengine_mocks ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/CompositionEngine.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/Display.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/DisplayColorProfile.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/DisplaySurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/LayerFE.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/NativeWindow.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/Output.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/OutputLayer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/mock/RenderSurface.cpp
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
    add_library(libcompositionengine_mocks SHARED ${SOURCE_FILES})

    target_include_directories(libcompositionengine_mocks
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/include
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine
    )

    # Add c/c++ flags
    target_compile_options(libcompositionengine_mocks PRIVATE
            -DFLYME_GRAPHICS_EXTEND
            -DFLYME_COLOR_MANAGEMENT
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(libcompositionengine_mocks
            android.hardware.graphics.allocator-2.0
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
            android.hardware.power-1.0
            android.hardware.power-1.3
            libEGL
            libGLESv2
            libSurfaceFlingerProp
            libaidlcommonsupport
            libarect
            libbase
            libc
            libc++
            libc++_static
            libc++demangle
            libcgrouprc
            libcgrouprc_format
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcompositionengine
            libcutils
            libdl
            libgmock
            libgtest
            libgui
            libjsoncpp
            liblayers_proto
            liblog
            libm
            libmath
            libnativewindow
            libperfetto_client_experimental
            libpng
            libprocessgroup
            libprotobuf-cpp-lite
            libqticompositionengineextension
            librenderengine
            libruntime_log
            libskia_renderengine
            libsurfaceflinger_headers
            libtiff
            libtimestats
            libtonemap
            libui
            libutils
            libvulkan
            libz
    )
endif ()
#################### libcompositionengine_mocks ####################

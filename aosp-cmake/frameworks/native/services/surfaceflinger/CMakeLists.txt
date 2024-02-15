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

project(surfaceflinger)

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

#################### libSurfaceFlingerProp ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/SurfaceFlingerProperties.cpp
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
    add_library(libSurfaceFlingerProp SHARED ${SOURCE_FILES})

    target_include_directories(libSurfaceFlingerProp
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/sysprop/libSurfaceFlingerProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(libSurfaceFlingerProp PRIVATE
            -Werror
    )

    target_link_libraries(libSurfaceFlingerProp
            android.hardware.configstore-utils
            android.hardware.configstore-1.0
            android.hardware.configstore-1.1
            android.hardware.graphics.common-1.2
            libSurfaceFlingerProperties
            libc
            libc++
            libc++demangle
            libclang_rt.builtins
            libdl
            libhidlbase
            liblog
            libm
            libui
            libutils
    )
endif ()
#################### libSurfaceFlingerProp ####################

#################### surfaceflinger ####################
set(FILES
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/BackgroundExecutor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Client.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/ClientCache.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Display/DisplaySnapshot.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayDevice.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayHardware/AidlComposerHal.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayHardware/ComposerHal.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayHardware/FramebufferSurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayHardware/HWC2.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayHardware/HWComposer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayHardware/HidlComposerHal.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayHardware/PowerAdvisor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayHardware/VirtualDisplaySurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/DisplayRenderArea.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Effects/Daltonizer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/EventLog/EventLog.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrontEnd/LayerCreationArgs.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrontEnd/LayerHandle.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrontEnd/LayerSnapshot.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrontEnd/LayerSnapshotBuilder.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrontEnd/LayerHierarchy.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrontEnd/LayerLifecycleManager.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrontEnd/RequestedLayerState.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrontEnd/TransactionHandler.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FlagManager.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FpsReporter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrameTracer/FrameTracer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrameTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/HdrLayerInfoReporter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/WindowInfosListenerInvoker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Layer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/LayerFE.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/LayerProtoHelper.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/LayerRenderArea.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/LayerVector.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/NativeWindowSurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/RefreshRateOverlay.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/RegionSamplingThread.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/RenderArea.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/EventThread.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/FrameRateOverrideMappings.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/OneShotTimer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/LayerHistory.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/LayerInfo.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/MessageQueue.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/RefreshRateSelector.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/Scheduler.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/VSyncDispatchTimerQueue.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/VSyncPredictor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/VSyncReactor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/VsyncConfiguration.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/VsyncModulator.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/VsyncSchedule.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/ScreenCaptureOutput.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/StartPropertySetThread.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/SurfaceFlinger.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/SurfaceFlingerDefaultFactory.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Tracing/LayerTracing.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Tracing/TransactionTracing.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Tracing/TransactionProtoParser.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TransactionCallbackInvoker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TunnelModeEnabledReporter.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/main_surfaceflinger.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/SurfaceFlingerFactory.cpp
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
    add_executable(surfaceflinger ${SOURCE_FILES})

    target_include_directories(surfaceflinger
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/CompositionEngine/include
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/FrameTimeline
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/Scheduler/include
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/QtiExtension
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TimeStats
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/timestatsproto/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/timestatsproto/libtimestats_proto/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/native/services/surfaceflinger/TimeStats/timestatsproto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/timestatsproto/libtimestats_proto/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/timestatsproto/libtimestats_proto/android_${OUT_ARCH_CPU}_shared/gen/proto/frameworks/native/services/surfaceflinger/TimeStats/timestatsproto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/TimeStats/timestatsproto/libtimestats_proto/android_${OUT_ARCH_CPU}_shared/gen/proto
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/layerproto/include
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/layerproto/liblayers_proto/android_${OUT_ARCH_CPU}_static/gen/proto/frameworks/native/services/surfaceflinger/layerproto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/layerproto/liblayers_proto/android_${OUT_ARCH_CPU}_static/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/layerproto/liblayers_proto/android_${OUT_ARCH_CPU}_shared/gen/proto/frameworks/native/services/surfaceflinger/layerproto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/layerproto/liblayers_proto/android_${OUT_ARCH_CPU}_shared/gen/proto
            ${OUT_INTERMEDIATES_ROOT}/frameworks/native/services/surfaceflinger/sysprop/libSurfaceFlingerProperties/android_${OUT_ARCH_CPU}_static/gen/sysprop/include
    )

    # Add c/c++ flags
    target_compile_options(surfaceflinger PRIVATE
            -DQTI_DISPLAY_EXTENSION
            -Wno-unused-parameter
            -Wno-unused-parameter
            -DQTI_COMPOSER3_EXTENSIONS
            -Wall
            -Werror
            -Wextra
    )

    target_link_libraries(surfaceflinger
            android.frameworks.displayservice-1.0
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
            libc++
            libc++demangle
            libclang_rt.builtins
            libclang_rt.ubsan_minimal
            libcompositionengine
            libcutils
            libdisplayconfig.system.qti
            libdisplayservicehidl
            libdl
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
            libqticompositionengineextension
            libqtisurfaceflingerextension
            librenderengine
            libscheduler
            libserviceutils
            libshaders
            libskia_renderengine
            libsync
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
#################### surfaceflinger ####################

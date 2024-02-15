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

project(surfaceflinger_displayhardware_fuzzer)

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

#################### surfaceflinger_fuzzer ####################
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
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerHalController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockComposer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockHWC2.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPower.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPowerHintSession.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerAdvisor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockEventThread.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTimeline.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTracer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockNativeWindowSurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockTimeStats.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVsyncController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncDispatch.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/system/window/MockNativeWindow.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer/surfaceflinger_fuzzer.cpp
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
    add_executable(surfaceflinger_fuzzer ${SOURCE_FILES})

    target_include_directories(surfaceflinger_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(surfaceflinger_fuzzer PRIVATE
            -DQTI_DISPLAY_EXTENSION
            -Wno-unused-parameter
            -Wno-unused-parameter
            -DQTI_COMPOSER3_EXTENSIONS
            -Wall
            -Werror
            -Wextra
            -Wno-unused-result
            -Wno-unused-function
    )

    target_link_libraries(surfaceflinger_fuzzer
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
            android.hardware.graphics.composer-2.1-resources
            android.hardware.graphics.composer-2.2
            android.hardware.graphics.composer-2.2-command-buffer
            android.hardware.graphics.composer-2.3
            android.hardware.graphics.composer-2.3-command-buffer
            android.hardware.graphics.composer-2.4
            android.hardware.graphics.composer-2.4-command-buffer
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
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
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine
            libcompositionengine_mocks
            libcutils
            libdisplayconfig.system.qti
            libdl
            libem
            libfmq
            libframetimeline
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
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
            libprotobuf-cpp-full
            libprotobuf-cpp-lite
            libprotoutil
            libqticompositionengineextension
            libqtisurfaceflingerextension
            librenderengine
            librenderengine_mocks
            libscheduler
            libserviceutils
            libshaders
            libskia_renderengine
            libstatspull
            libstatssocket
            libsurfaceflinger_headers
            libsync
            libtiff
            libtimestats
            libtimestats_proto
            libtonemap
            libui
            libui_fuzzableDataspaces_headers
            libui_headers
            libutils
            libvulkan
            libz
            perfetto_trace_protos
            server_configurable_flags
            vendor.qti.hardware.display.composer3-V1-ndk
            vendor.qti.hardware.display.composer-3.0
            vendor.qti.hardware.display.composer-3.1
            vendor.qti.hardware.display.config-V7-ndk
    )
endif ()
#################### surfaceflinger_fuzzer ####################

#################### surfaceflinger_layer_fuzzer ####################
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
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerHalController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockComposer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockHWC2.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPower.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPowerHintSession.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerAdvisor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockEventThread.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTimeline.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTracer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockNativeWindowSurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockTimeStats.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVsyncController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncDispatch.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/system/window/MockNativeWindow.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer/surfaceflinger_layer_fuzzer.cpp
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
    add_executable(surfaceflinger_layer_fuzzer ${SOURCE_FILES})

    target_include_directories(surfaceflinger_layer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(surfaceflinger_layer_fuzzer PRIVATE
            -DQTI_DISPLAY_EXTENSION
            -Wno-unused-parameter
            -Wno-unused-parameter
            -DQTI_COMPOSER3_EXTENSIONS
            -Wall
            -Werror
            -Wextra
            -Wno-unused-result
            -Wno-unused-function
    )

    target_link_libraries(surfaceflinger_layer_fuzzer
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
            android.hardware.graphics.composer-2.1-resources
            android.hardware.graphics.composer-2.2
            android.hardware.graphics.composer-2.2-command-buffer
            android.hardware.graphics.composer-2.3
            android.hardware.graphics.composer-2.3-command-buffer
            android.hardware.graphics.composer-2.4
            android.hardware.graphics.composer-2.4-command-buffer
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
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
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine
            libcompositionengine_mocks
            libcutils
            libdisplayconfig.system.qti
            libdl
            libem
            libfmq
            libframetimeline
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_headers
            libgui_mocks
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
            libprotobuf-cpp-full
            libprotobuf-cpp-lite
            libprotoutil
            libqticompositionengineextension
            libqtisurfaceflingerextension
            librenderengine
            librenderengine_mocks
            libscheduler
            libserviceutils
            libshaders
            libskia_renderengine
            libstatspull
            libstatssocket
            libsurfaceflinger_headers
            libsync
            libtiff
            libtimestats
            libtimestats_proto
            libtonemap
            libui
            libui_fuzzableDataspaces_headers
            libui_headers
            libutils
            libvulkan
            libz
            perfetto_trace_protos
            server_configurable_flags
            vendor.qti.hardware.display.composer3-V1-ndk
            vendor.qti.hardware.display.composer-3.0
            vendor.qti.hardware.display.composer-3.1
            vendor.qti.hardware.display.config-V7-ndk
    )
endif ()
#################### surfaceflinger_layer_fuzzer ####################

#################### surfaceflinger_frametracer_fuzzer ####################
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
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerHalController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockComposer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockHWC2.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPower.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPowerHintSession.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerAdvisor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockEventThread.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTimeline.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTracer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockNativeWindowSurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockTimeStats.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVsyncController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncDispatch.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/system/window/MockNativeWindow.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer/surfaceflinger_frametracer_fuzzer.cpp
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
    add_executable(surfaceflinger_frametracer_fuzzer ${SOURCE_FILES})

    target_include_directories(surfaceflinger_frametracer_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(surfaceflinger_frametracer_fuzzer PRIVATE
            -DQTI_DISPLAY_EXTENSION
            -Wno-unused-parameter
            -Wno-unused-parameter
            -DQTI_COMPOSER3_EXTENSIONS
            -Wall
            -Werror
            -Wextra
            -Wno-unused-result
            -Wno-unused-function
    )

    target_link_libraries(surfaceflinger_frametracer_fuzzer
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
            android.hardware.graphics.composer-2.1-resources
            android.hardware.graphics.composer-2.2
            android.hardware.graphics.composer-2.2-command-buffer
            android.hardware.graphics.composer-2.3
            android.hardware.graphics.composer-2.3-command-buffer
            android.hardware.graphics.composer-2.4
            android.hardware.graphics.composer-2.4-command-buffer
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
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
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine
            libcompositionengine_mocks
            libcutils
            libdisplayconfig.system.qti
            libdl
            libem
            libfmq
            libframetimeline
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
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
            libprotobuf-cpp-full
            libprotobuf-cpp-lite
            libprotoutil
            libqticompositionengineextension
            libqtisurfaceflingerextension
            librenderengine
            librenderengine_mocks
            libscheduler
            libserviceutils
            libshaders
            libskia_renderengine
            libstatspull
            libstatssocket
            libsurfaceflinger_headers
            libsync
            libtiff
            libtimestats
            libtimestats_proto
            libtonemap
            libui
            libui_fuzzableDataspaces_headers
            libui_headers
            libutils
            libvulkan
            libz
            perfetto_trace_protos
            server_configurable_flags
            vendor.qti.hardware.display.composer3-V1-ndk
            vendor.qti.hardware.display.composer-3.0
            vendor.qti.hardware.display.composer-3.1
            vendor.qti.hardware.display.config-V7-ndk
    )
endif ()
#################### surfaceflinger_frametracer_fuzzer ####################

#################### surfaceflinger_scheduler_fuzzer ####################
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
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerHalController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockComposer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockHWC2.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPower.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPowerHintSession.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerAdvisor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockEventThread.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTimeline.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTracer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockNativeWindowSurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockTimeStats.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVsyncController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncDispatch.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/system/window/MockNativeWindow.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer/surfaceflinger_scheduler_fuzzer.cpp
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
    add_executable(surfaceflinger_scheduler_fuzzer ${SOURCE_FILES})

    target_include_directories(surfaceflinger_scheduler_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(surfaceflinger_scheduler_fuzzer PRIVATE
            -DQTI_DISPLAY_EXTENSION
            -Wno-unused-parameter
            -Wno-unused-parameter
            -DQTI_COMPOSER3_EXTENSIONS
            -Wall
            -Werror
            -Wextra
            -Wno-unused-result
            -Wno-unused-function
    )

    target_link_libraries(surfaceflinger_scheduler_fuzzer
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
            android.hardware.graphics.composer-2.1-resources
            android.hardware.graphics.composer-2.2
            android.hardware.graphics.composer-2.2-command-buffer
            android.hardware.graphics.composer-2.3
            android.hardware.graphics.composer-2.3-command-buffer
            android.hardware.graphics.composer-2.4
            android.hardware.graphics.composer-2.4-command-buffer
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
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
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine
            libcompositionengine_mocks
            libcutils
            libdisplayconfig.system.qti
            libdl
            libem
            libfmq
            libframetimeline
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
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
            libprotobuf-cpp-full
            libprotobuf-cpp-lite
            libprotoutil
            libqticompositionengineextension
            libqtisurfaceflingerextension
            librenderengine
            librenderengine_mocks
            libscheduler
            libserviceutils
            libshaders
            libskia_renderengine
            libstatspull
            libstatssocket
            libsurfaceflinger_headers
            libsync
            libtiff
            libtimestats
            libtimestats_proto
            libtonemap
            libui
            libui_fuzzableDataspaces_headers
            libui_headers
            libutils
            libvulkan
            libz
            perfetto_trace_protos
            server_configurable_flags
            vendor.qti.hardware.display.composer3-V1-ndk
            vendor.qti.hardware.display.composer-3.0
            vendor.qti.hardware.display.composer-3.1
            vendor.qti.hardware.display.config-V7-ndk
    )
endif ()
#################### surfaceflinger_scheduler_fuzzer ####################

#################### surfaceflinger_displayhardware_fuzzer ####################
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
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerHalController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockComposer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockHWC2.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPower.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockIPowerHintSession.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/DisplayHardware/MockPowerAdvisor.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockEventThread.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTimeline.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockFrameTracer.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockNativeWindowSurface.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockTimeStats.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVsyncController.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncDispatch.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/MockVSyncTracker.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/tests/unittests/mock/system/window/MockNativeWindow.cpp
        ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer/surfaceflinger_displayhardware_fuzzer.cpp
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
    add_executable(surfaceflinger_displayhardware_fuzzer ${SOURCE_FILES})

    target_include_directories(surfaceflinger_displayhardware_fuzzer
            PUBLIC
            ${ANDROID_ROOT}/frameworks/native/services/surfaceflinger/fuzzer
    )

    # Add c/c++ flags
    target_compile_options(surfaceflinger_displayhardware_fuzzer PRIVATE
            -DQTI_DISPLAY_EXTENSION
            -Wno-unused-parameter
            -Wno-unused-parameter
            -DQTI_COMPOSER3_EXTENSIONS
            -Wall
            -Werror
            -Wextra
            -Wno-unused-result
            -Wno-unused-function
    )

    target_link_libraries(surfaceflinger_displayhardware_fuzzer
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
            android.hardware.graphics.composer-2.1-resources
            android.hardware.graphics.composer-2.2
            android.hardware.graphics.composer-2.2-command-buffer
            android.hardware.graphics.composer-2.3
            android.hardware.graphics.composer-2.3-command-buffer
            android.hardware.graphics.composer-2.4
            android.hardware.graphics.composer-2.4-command-buffer
            android.hardware.graphics.composer-2.4-hal
            android.hardware.graphics.mapper-2.0
            android.hardware.graphics.mapper-3.0
            android.hardware.graphics.mapper-4.0
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
            libclang_rt.fuzzer
            libclang_rt.ubsan_standalone
            libcompositionengine
            libcompositionengine_mocks
            libcutils
            libdisplayconfig.system.qti
            libdl
            libem
            libfmq
            libframetimeline
            libgmock
            libgmock_main
            libgmock_main_ndk
            libgmock_ndk
            libgtest_ndk_c++
            libgui
            libgui_aidl_static
            libgui_mocks
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
            libprotobuf-cpp-full
            libprotobuf-cpp-lite
            libprotoutil
            libqticompositionengineextension
            libqtisurfaceflingerextension
            librenderengine
            librenderengine_mocks
            libscheduler
            libserviceutils
            libshaders
            libskia_renderengine
            libstatspull
            libstatssocket
            libsurfaceflinger_headers
            libsync
            libtiff
            libtimestats
            libtimestats_proto
            libtonemap
            libui
            libui_fuzzableDataspaces_headers
            libui_headers
            libutils
            libvulkan
            libz
            perfetto_trace_protos
            server_configurable_flags
            vendor.qti.hardware.display.composer3-V1-ndk
            vendor.qti.hardware.display.composer-3.0
            vendor.qti.hardware.display.composer-3.1
            vendor.qti.hardware.display.config-V7-ndk
    )
endif ()
#################### surfaceflinger_displayhardware_fuzzer ####################
